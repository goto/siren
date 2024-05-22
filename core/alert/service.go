package alert

import (
	"context"
	"fmt"
	"time"

	saltlog "github.com/goto/salt/log"
	"github.com/goto/siren/core/notification"
	"github.com/goto/siren/core/template"
	"github.com/goto/siren/pkg/errors"
	"github.com/goto/siren/pkg/structure"
)

type LogService interface {
	ListAlertIDsBySilenceID(ctx context.Context, silenceID string) ([]int64, error)
}

type NotificationService interface {
	Dispatch(ctx context.Context, n notification.Notification) (string, error)
}

// Service handles business logic
type Service struct {
	cfg                 Config
	logger              saltlog.Logger
	repository          Repository
	logService          LogService
	notificationService NotificationService
	registry            map[string]AlertTransformer
}

// NewService returns repository struct
func NewService(cfg Config, logger saltlog.Logger, repository Repository, logService LogService, notificationService NotificationService, registry map[string]AlertTransformer) *Service {
	return &Service{cfg, logger, repository, logService, notificationService, registry}
}

func (s *Service) rollback(ctx context.Context, err error) error {
	if rbErr := s.repository.Rollback(ctx, err); rbErr != nil {
		return rbErr
	} else {
		return err
	}
}

func (s *Service) CreateAlerts(ctx context.Context, providerType string, providerID uint64, namespaceID uint64, body map[string]any) ([]Alert, error) {
	ctx = s.repository.WithTransaction(ctx)
	pluginService, err := s.getProviderPluginService(providerType)
	if err != nil {
		return nil, err
	}
	alerts, firingLen, err := pluginService.TransformToAlerts(ctx, providerID, namespaceID, body)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(alerts); i++ {
		createdAlert, err := s.repository.Create(ctx, alerts[i])
		if err != nil {
			repoErr := err
			if errors.Is(err, ErrRelation) {
				repoErr = errors.ErrNotFound.WithMsgf(err.Error())
			}
			if rbErr := s.rollback(ctx, repoErr); rbErr != nil {
				return nil, rbErr
			}
			return nil, repoErr
		}
		alerts[i].ID = createdAlert.ID
	}

	if len(alerts) > 0 {
		// Publish to notification service
		ns, err := BuildNotifications(alerts, firingLen, time.Now(), s.cfg.GroupBy)
		if err != nil {
			if rbErr := s.rollback(ctx, err); rbErr != nil {
				return nil, rbErr
			}
			s.logger.Warn("failed to build notifications from alert", "err", err, "alerts", alerts)
		}

		for _, n := range ns {
			if _, err := s.notificationService.Dispatch(ctx, n); err != nil {
				s.logger.Warn("failed to send alert as notification", "err", err, "notification", n)
			}
		}
	} else {
		s.logger.Warn("failed to send alert as notification, empty created alerts")
	}

	return alerts, nil
}

func (s *Service) List(ctx context.Context, flt Filter) ([]Alert, error) {
	if flt.EndTime == 0 {
		flt.EndTime = time.Now().Unix()
	}

	if flt.SilenceID != "" {
		alertIDs, err := s.logService.ListAlertIDsBySilenceID(ctx, flt.SilenceID)
		if err != nil {
			return nil, err
		}
		flt.IDs = alertIDs
	}

	return s.repository.List(ctx, flt)
}

func (s *Service) getProviderPluginService(providerType string) (AlertTransformer, error) {
	pluginService, exist := s.registry[providerType]
	if !exist {
		return nil, errors.ErrInvalid.WithMsgf("unsupported provider type: %q", providerType)
	}
	return pluginService, nil
}

// Transform alerts and populate Data and Labels to be interpolated to the system-default template
// .Data
// - id
// - status "FIRING"/"RESOLVED"
// - resource
// - template
// - metric_value
// - metric_name
// - generator_url
// - num_alerts_firing
// - dashboard
// - playbook
// - summary
// .Labels
// - severity "WARNING"/"CRITICAL"
// - alertname
// - (others labels defined in rules)
func BuildNotifications(
	alerts []Alert,
	firingLen int,
	createdTime time.Time,
	groupByLabels []string,
) ([]notification.Notification, error) {
	if len(alerts) == 0 {
		return nil, errors.New("empty alerts")
	}

	alertsMap, err := structure.GroupByLabels(alerts, groupByLabels, func(a Alert) map[string]string { return a.Labels })
	if err != nil {
		return nil, err
	}

	var notifications []notification.Notification

	for hashKey, groupedAlerts := range alertsMap {
		sampleAlert := groupedAlerts[0]

		data := map[string]any{}

		mergedAnnotations := map[string][]string{}
		for _, a := range groupedAlerts {
			for k, v := range a.Annotations {
				mergedAnnotations[k] = append(mergedAnnotations[k], v)
			}
		}
		// make unique
		for k, v := range mergedAnnotations {
			mergedAnnotations[k] = structure.RemoveDuplicate(v)
		}
		// render annotations
		for k, vSlice := range mergedAnnotations {
			for _, v := range vSlice {
				if _, ok := data[k]; ok {
					data[k] = fmt.Sprintf("%s\n%s", data[k], v)
				} else {
					data[k] = v
				}
			}
		}

		data["status"] = sampleAlert.Status
		data["generator_url"] = sampleAlert.GeneratorURL
		data["num_alerts_firing"] = firingLen

		alertIDs := []int64{}

		for _, a := range groupedAlerts {
			alertIDs = append(alertIDs, int64(a.ID))
		}

		for k, v := range sampleAlert.Labels {
			data[k] = v
		}

		notifications = append(notifications, notification.Notification{
			NamespaceID: sampleAlert.NamespaceID,
			Type:        notification.TypeAlert,
			Data:        data,
			Labels:      sampleAlert.Labels,
			Template:    template.ReservedName_SystemDefault,
			UniqueKey:   structure.HashGroupKey(sampleAlert.GroupKey, hashKey),
			CreatedAt:   createdTime,
			AlertIDs:    alertIDs,
		})
	}

	return notifications, nil
}
