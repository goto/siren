package cortexv1plugin

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	texttemplate "text/template"
	"time"

	"github.com/goto/siren/core/alert"
	"github.com/goto/siren/core/namespace"
	"github.com/goto/siren/core/provider"
	"github.com/goto/siren/core/rule"
	"github.com/goto/siren/core/template"
	"github.com/goto/siren/pkg/errors"
	"github.com/goto/siren/pkg/httpclient"
	"github.com/goto/siren/plugins/receivers/base"
	"github.com/grafana/cortex-tools/pkg/client"
	"github.com/grafana/cortex-tools/pkg/rules/rwrulefmt"
	"github.com/hashicorp/go-hclog"
	"github.com/imdario/mergo"
	"github.com/mcuadros/go-defaults"
	"github.com/mitchellh/mapstructure"
	promconfig "github.com/prometheus/alertmanager/config"
	"github.com/prometheus/prometheus/pkg/rulefmt"
	"gopkg.in/yaml.v3"
)

type CortexCaller interface {
	CreateAlertmanagerConfig(ctx context.Context, cfg string, templates map[string]string) error
	CreateRuleGroup(ctx context.Context, namespace string, rg rwrulefmt.RuleGroup) error
	DeleteRuleGroup(ctx context.Context, namespace, groupName string) error
	GetRuleGroup(ctx context.Context, namespace, groupName string) (*rwrulefmt.RuleGroup, error)
	ListRules(ctx context.Context, namespace string) (map[string][]rwrulefmt.RuleGroup, error)
}

// Service is a service layer of cortex provider plugin
type PluginService struct {
	base.UnimplementedService

	logger         hclog.Logger
	cfg            Config
	helperTemplate string
	configYaml     string
	cortexClient   CortexCaller
	httpClient     *httpclient.Client
}

// NewPluginService returns cortex service provider plugin struct
func NewPluginService(logger hclog.Logger, opts ...ServiceOption) *PluginService {
	s := &PluginService{
		logger:         logger,
		helperTemplate: HelperTemplateString,
		configYaml:     ConfigYamlString,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *PluginService) SetConfig(ctx context.Context, configRaw string) error {
	var defaultConfig = new(Config)
	defaults.SetDefaults(defaultConfig)

	var cfg Config
	var configJSON = json.RawMessage(configRaw)

	if err := json.Unmarshal(configJSON, &cfg); err != nil {
		return errors.ErrInvalid.WithMsgf("invalid config json").WithCausef(err.Error())
	}

	if err := mergo.Merge(&cfg, defaultConfig); err != nil {
		return errors.ErrInvalid.WithMsgf("failed to merge default config with the input config").WithCausef(err.Error())
	}

	if cfg.WebhookBaseAPI == "" {
		return errors.New("Cortex webhook base api string in config cannot be empty")
	}

	s.cfg = cfg
	s.httpClient = httpclient.New(cfg.HTTPClient)

	s.logger.Debug(fmt.Sprintf("running cortex plugin with this config: %+v", cfg))
	return nil
}

// TransformToAlerts is a function to transform alert body in hook API to []*alert.Alert
func (s *PluginService) TransformToAlerts(ctx context.Context, providerID uint64, namespaceID uint64, body map[string]any) ([]alert.Alert, int, error) {
	var groupAlert = &GroupAlert{}
	if err := mapstructure.Decode(body, groupAlert); err != nil {
		return nil, 0, err
	}

	var (
		alerts        = make([]alert.Alert, 0)
		badAlertCount = 0
		firingLen     = 0
	)

	for _, item := range groupAlert.Alerts {

		if err := item.Validate(); err != nil {
			s.logger.Error(fmt.Sprintf("invalid alerts: %s", err.Error()), "group key", groupAlert.GroupKey, "alert detail", item)
			badAlertCount++
			continue
		}

		if item.Status == "firing" {
			firingLen++
		}

		severity := item.Labels["severity"]
		if item.Status == "resolved" {
			severity = item.Status
		}

		startsAt, err := time.Parse(time.RFC3339Nano, item.StartsAt)
		if err != nil {
			badAlertCount++
			break
		}

		alrt := alert.Alert{
			ProviderID:   providerID,
			NamespaceID:  namespaceID,
			ResourceName: item.Annotations["resource"],
			MetricName:   item.Annotations["metric_name"],
			MetricValue:  item.Annotations["metric_value"],
			Severity:     severity,
			Rule:         item.Annotations["template"],
			TriggeredAt:  startsAt,

			GroupKey:     groupAlert.GroupKey,
			Status:       item.Status,
			Annotations:  item.Annotations,
			Labels:       item.Labels,
			GeneratorURL: item.GeneratorURL,
			Fingerprint:  item.Fingerprint,
		}

		alerts = append(alerts, alrt)
	}

	if badAlertCount > 0 {
		s.logger.Error("parameters are missing for alert", "group key", groupAlert.GroupKey, "alert count", badAlertCount)
		return alerts, firingLen, nil
	}

	return alerts, firingLen, nil
}

// SyncRuntimeConfig synchronizes runtime configuration of provider
func (s *PluginService) SyncRuntimeConfig(ctx context.Context, namespaceID uint64, namespaceURN string, namespaceLabels map[string]string, prov provider.Provider) (map[string]string, error) {
	webhookURL := fmt.Sprintf("%s/%d/%d", s.cfg.WebhookBaseAPI, prov.ID, namespaceID)

	tmplConfig := TemplateConfig{
		GroupWaitDuration:      s.cfg.GroupWaitDuration,
		GroupIntervalDuration:  s.cfg.GroupIntervalDuration,
		RepeatIntervalDuration: s.cfg.RepeatIntervalDuration,
		WebhookURL:             webhookURL,
	}

	cfg, err := s.generateAlertmanagerConfig(tmplConfig)
	if err != nil {
		return nil, err
	}
	templates := map[string]string{
		"helper.tmpl": s.helperTemplate,
	}

	cortexClient, err := s.getCortexClient(prov.Host, namespaceURN)
	if err != nil {
		return nil, err
	}

	if err = cortexClient.CreateAlertmanagerConfig(ctx, cfg, templates); err != nil {
		return nil, err
	}
	return make(map[string]string), nil
}

// UpsertRule manages upsert logic to cortex ruler. Cortex client API granularity is on the rule-group.
// This function has a logic to work with rule-level granurality and adapt it to cortex logic.
func (s *PluginService) UpsertRule(ctx context.Context, ns namespace.Namespace, prov provider.Provider, rl *rule.Rule, templateToUpdate *template.Template) error {
	inputValues := make(map[string]string)
	for _, v := range rl.Variables {
		inputValues[v.Name] = v.Value
	}

	renderedRule, err := template.RenderWithEnrichedDefault(templateToUpdate.Body, templateToUpdate.Variables, inputValues)
	if err != nil {
		return err
	}

	cortexClient, err := s.getCortexClient(prov.Host, ns.URN)
	if err != nil {
		return err
	}

	var upsertedRuleNodes []rulefmt.RuleNode
	if err := yaml.Unmarshal([]byte(renderedRule), &upsertedRuleNodes); err != nil {
		return errors.ErrInvalid.WithMsgf("cannot parse upserted rule").WithCausef(err.Error())
	}

	cortexRuleGroup, err := cortexClient.GetRuleGroup(ctx, rl.Namespace, rl.GroupName)
	if err != nil {
		if errors.Is(err, client.ErrResourceNotFound) {
			cortexRuleGroup = &rwrulefmt.RuleGroup{}
		} else {
			return errors.ErrInvalid.WithMsgf("cannot get rule group from cortex when upserting rules").WithCausef(err.Error())
		}
	}

	newRuleNodes, err := mergeRuleNodes(cortexRuleGroup.Rules, upsertedRuleNodes, rl.Enabled)
	if err != nil {
		return err
	}

	if len(newRuleNodes) == 0 {
		if err := cortexClient.DeleteRuleGroup(ctx, rl.Namespace, rl.GroupName); err != nil {
			if err.Error() == "requested resource not found" {
				return nil
			}
			return fmt.Errorf("error calling cortex: %w", err)
		}
		return nil
	}

	cortexRuleGroup = &rwrulefmt.RuleGroup{
		RuleGroup: rulefmt.RuleGroup{
			Name:  rl.GroupName,
			Rules: newRuleNodes,
		},
	}
	if err := cortexClient.CreateRuleGroup(ctx, rl.Namespace, *cortexRuleGroup); err != nil {
		return fmt.Errorf("error calling cortex: %w", err)
	}
	return nil
}

func mergeRuleNodes(ruleNodes []rulefmt.RuleNode, newRuleNodes []rulefmt.RuleNode, enabled bool) ([]rulefmt.RuleNode, error) {
	for _, nrn := range newRuleNodes {
		var action string = "insert"
		var idxCount = 0
		for _, ruleNode := range ruleNodes {
			if ruleNode.Alert.Value == nrn.Alert.Value {
				if !enabled {
					action = "delete"
					break
				}
				action = "update"
				break
			}
			idxCount++
		}

		switch action {
		case "delete":
			if idxCount >= len(ruleNodes) || idxCount < 0 {
				return nil, errors.New("something wrong when comparing rule node")
			}
			ruleNodes = append(ruleNodes[:idxCount], ruleNodes[idxCount+1:]...)
		case "update":
			ruleNodes[idxCount] = nrn
		default:
			if !enabled {
				return ruleNodes, nil
			}
			ruleNodes = append(ruleNodes, nrn)
		}
	}

	return ruleNodes, nil
}

func (s *PluginService) generateAlertmanagerConfig(tmplConfig TemplateConfig) (string, error) {
	delims := texttemplate.New("alertmanagerConfigTemplate").Delims("[[", "]]")
	parse, err := delims.Parse(s.configYaml)
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer
	err = parse.Execute(&tpl, tmplConfig)
	if err != nil {
		// it is unlikely that the code returns error here
		return "", err
	}
	configStr := tpl.String()
	_, err = promconfig.Load(configStr)
	if err != nil {
		return "", err
	}
	return configStr, nil
}

func (s *PluginService) getCortexClient(address string, tenant string) (CortexCaller, error) {
	if s.cortexClient != nil {
		return s.cortexClient, nil
	}
	cortexClient, err := client.New(client.Config{
		Address: address,
		ID:      tenant,
	})
	if err != nil {
		return nil, err
	}

	cortexClient.Client = *s.httpClient.HTTP()

	return cortexClient, nil
}
