package plugins

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/goto/siren/core/alert"
	"github.com/goto/siren/core/namespace"
	"github.com/goto/siren/core/provider"
	"github.com/goto/siren/core/rule"
	"github.com/goto/siren/core/template"
	sirenproviderv1beta1 "github.com/goto/siren/proto/gotocompany/siren/provider/v1beta1"
	sirenv1beta1 "github.com/goto/siren/proto/gotocompany/siren/v1beta1"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/structpb"
)

type ProviderV1beta1 interface {
	SyncRuntimeConfig(ctx context.Context, namespaceID uint64, namespaceURN string, namespaceLabels map[string]string, prov provider.Provider) (map[string]string, error)
	UpsertRule(ctx context.Context, ns namespace.Namespace, prov provider.Provider, rl *rule.Rule, templateToUpdate *template.Template) error
	SetConfig(ctx context.Context, configRaw string) error
	TransformToAlerts(ctx context.Context, providerID uint64, namespaceID uint64, body map[string]any) ([]alert.Alert, int, error)
}

type GRPCClient struct {
	client sirenproviderv1beta1.ProviderServiceClient
}

func NewProviderClient(c sirenproviderv1beta1.ProviderServiceClient) *GRPCClient {
	return &GRPCClient{
		client: c,
	}
}

func (c *GRPCClient) SyncRuntimeConfig(ctx context.Context, namespaceID uint64, namespaceURN string, namespaceLabels map[string]string, prov provider.Provider) (map[string]string, error) {
	protoProv, err := prov.ToV1beta1Proto()
	if err != nil {
		return nil, err
	}
	resp, err := c.client.SyncRuntimeConfig(ctx, &sirenproviderv1beta1.SyncRuntimeConfigRequest{
		NamespaceId:  fmt.Sprintf("%d", namespaceID),
		NamespaceUrn: namespaceURN,
		Provider:     protoProv,
		Labels:       namespaceLabels,
	})
	if err != nil {
		return nil, err
	}
	return resp.GetLabels(), nil
}

func (c *GRPCClient) UpsertRule(ctx context.Context, ns namespace.Namespace, prov provider.Provider, rl *rule.Rule, templateToUpdate *template.Template) error {
	protoProv, err := prov.ToV1beta1Proto()
	if err != nil {
		return err
	}
	nsProto, err := ns.ToV1beta1Proto()
	if err != nil {
		return err
	}
	if _, err := c.client.UpsertRule(ctx, &sirenproviderv1beta1.UpsertRuleRequest{
		Namespace: nsProto,
		Provider:  protoProv,
		Rule:      rl.ToV1beta1Proto(),
		Template:  templateToUpdate.ToV1beta1Proto(),
	}); err != nil {
		return err
	}
	return nil
}

func (c *GRPCClient) SetConfig(ctx context.Context, configRaw string) error {
	if _, err := c.client.SetConfig(ctx, &sirenproviderv1beta1.SetConfigRequest{
		ConfigRaw: configRaw,
	}); err != nil {
		return err
	}
	return nil
}

func (c *GRPCClient) TransformToAlerts(ctx context.Context, providerID uint64, namespaceID uint64, body map[string]any) ([]alert.Alert, int, error) {
	bodyPB, err := structpb.NewStruct(body)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to transform body to structpb: %s", err.Error())
	}
	resp, err := c.client.TransformToAlerts(ctx, &sirenproviderv1beta1.TransformToAlertsRequest{
		ProviderId:  fmt.Sprintf("%d", providerID),
		NamespaceId: fmt.Sprintf("%d", namespaceID),
		Body:        bodyPB,
	})
	if err != nil {
		return nil, 0, err
	}

	var alerts []alert.Alert
	for _, alertPB := range resp.GetAlerts() {
		alerts = append(alerts, *alert.FromV1beta1Proto(alertPB))
	}
	return alerts, int(resp.GetFiringNum()), nil
}

type GRPCServer struct {
	sirenproviderv1beta1.UnimplementedProviderServiceServer
	service ProviderV1beta1
}

func NewProviderServer(service ProviderV1beta1) *GRPCServer {
	return &GRPCServer{
		service: service,
	}
}

func (s *GRPCServer) SyncRuntimeConfig(ctx context.Context, req *sirenproviderv1beta1.SyncRuntimeConfigRequest) (*sirenproviderv1beta1.SyncRuntimeConfigResponse, error) {
	prov := provider.Provider{}
	if req.GetProvider() != nil {
		grpcProvider := req.GetProvider()

		prov.ID = grpcProvider.GetId()
		prov.URN = grpcProvider.GetUrn()
		prov.Host = grpcProvider.GetHost()
		prov.Name = grpcProvider.GetName()
		prov.Type = grpcProvider.GetType()
		prov.Credentials = grpcProvider.GetCredentials().AsMap()
		prov.Labels = grpcProvider.GetLabels()
		prov.CreatedAt = grpcProvider.GetCreatedAt().AsTime()
		prov.UpdatedAt = grpcProvider.GetUpdatedAt().AsTime()
	}

	namespaceIDUint64, err := strconv.ParseUint(req.GetNamespaceId(), 10, 64)
	if err != nil {
		return nil, errors.New("error parsing namespace ID")
	}

	encrichedLabels, err := s.service.SyncRuntimeConfig(ctx, namespaceIDUint64, req.GetNamespaceUrn(), req.GetLabels(), prov)
	if err != nil {
		return nil, err
	}

	return &sirenproviderv1beta1.SyncRuntimeConfigResponse{
		Labels: encrichedLabels,
	}, nil
}

func (s *GRPCServer) UpsertRule(ctx context.Context, req *sirenproviderv1beta1.UpsertRuleRequest) (*sirenproviderv1beta1.UpsertRuleResponse, error) {
	prov := provider.Provider{}
	if req.GetProvider() != nil {
		grpcProvider := req.GetProvider()

		prov = provider.FromV1beta1Proto(grpcProvider)
	}

	ns := namespace.Namespace{}
	if req.GetNamespace() != nil {
		grpcNamespace := req.GetNamespace()

		ns = namespace.FromV1beta1Proto(grpcNamespace)
	}

	rl := rule.Rule{}
	if req.GetRule() != nil {
		grpcRule := req.GetRule()

		rl = rule.FromV1beta1Proto(grpcRule)
	}

	tmplt := template.Template{}
	if req.GetTemplate() != nil {
		grpcTemplate := req.GetTemplate()

		tmplt = template.FromV1beta1Proto(grpcTemplate)
	}
	return &sirenproviderv1beta1.UpsertRuleResponse{}, s.service.UpsertRule(ctx, ns, prov, &rl, &tmplt)
}

func (s *GRPCServer) SetConfig(ctx context.Context, req *sirenproviderv1beta1.SetConfigRequest) (*sirenproviderv1beta1.SetConfigResponse, error) {
	return &sirenproviderv1beta1.SetConfigResponse{}, s.service.SetConfig(ctx, req.GetConfigRaw())
}

func (s *GRPCServer) TransformToAlerts(ctx context.Context, req *sirenproviderv1beta1.TransformToAlertsRequest) (*sirenproviderv1beta1.TransformToAlertsResponse, error) {
	providerIDUint64, err := strconv.ParseUint(req.GetProviderId(), 10, 64)
	if err != nil {
		return nil, errors.New("error parsing provider ID")
	}
	namespaceIDUint64, err := strconv.ParseUint(req.GetNamespaceId(), 10, 64)
	if err != nil {
		return nil, errors.New("error parsing namespace ID")
	}
	alerts, firingNum, err := s.service.TransformToAlerts(ctx, providerIDUint64, namespaceIDUint64, req.GetBody().AsMap())
	if err != nil {
		return nil, err
	}

	var alertsPB = make([]*sirenv1beta1.Alert, 0)
	for _, al := range alerts {
		alertsPB = append(alertsPB, al.ToV1beta1Proto())
	}

	return &sirenproviderv1beta1.TransformToAlertsResponse{
		Alerts:    alertsPB,
		FiringNum: uint64(firingNum),
	}, nil
}

type ProviderV1beta1GRPCPlugin struct {
	plugin.NetRPCUnsupportedPlugin
	GRPCProvider func() sirenproviderv1beta1.ProviderServiceServer
}

func (c *ProviderV1beta1GRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	sirenproviderv1beta1.RegisterProviderServiceServer(s, c.GRPCProvider())
	return nil
}

func (c *ProviderV1beta1GRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, cl *grpc.ClientConn) (interface{}, error) {
	return NewProviderClient(sirenproviderv1beta1.NewProviderServiceClient(cl)), nil
}
