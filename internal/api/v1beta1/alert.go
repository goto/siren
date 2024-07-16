package v1beta1

import (
	"context"

	"github.com/goto/siren/core/alert"
	"github.com/goto/siren/internal/api"
	sirenv1beta1 "github.com/goto/siren/proto/gotocompany/siren/v1beta1"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *GRPCServer) ListAlerts(ctx context.Context, req *sirenv1beta1.ListAlertsRequest) (*sirenv1beta1.ListAlertsResponse, error) {
	alerts, err := s.alertService.List(ctx, alert.Filter{
		ResourceName: req.GetResourceName(),
		ProviderID:   req.GetProviderId(),
		StartTime:    int64(req.GetStartTime()),
		EndTime:      int64(req.GetEndTime()),
		// SilenceID:    req.GetSilenced(),
	})
	if err != nil {
		return nil, api.GenerateRPCErr(s.logger, err)
	}

	items := []*sirenv1beta1.Alert{}
	for _, alert := range alerts {
		item := &sirenv1beta1.Alert{
			Id:            alert.ID,
			ProviderId:    alert.ProviderID,
			ResourceName:  alert.ResourceName,
			MetricName:    alert.MetricName,
			MetricValue:   alert.MetricValue,
			Severity:      alert.Severity,
			Rule:          alert.Rule,
			TriggeredAt:   timestamppb.New(alert.TriggeredAt),
			SilenceStatus: alert.SilenceStatus,
		}
		items = append(items, item)
	}
	return &sirenv1beta1.ListAlertsResponse{
		Alerts: items,
	}, nil
}

func (s *GRPCServer) CreateAlerts(ctx context.Context, req *sirenv1beta1.CreateAlertsRequest) (*sirenv1beta1.CreateAlertsResponse, error) {
	items, err := s.createAlerts(ctx, req.GetProviderType(), req.GetProviderId(), 0, req.GetBody().AsMap())
	if err != nil {
		return nil, api.GenerateRPCErr(s.logger, err)
	}

	return &sirenv1beta1.CreateAlertsResponse{
		Alerts: items,
	}, nil
}

func (s *GRPCServer) CreateAlertsWithNamespace(ctx context.Context, req *sirenv1beta1.CreateAlertsWithNamespaceRequest) (*sirenv1beta1.CreateAlertsWithNamespaceResponse, error) {
	var namespaceID uint64 = 0
	if !s.cfg.UseGlobalSubscription {
		namespaceID = req.GetNamespaceId()
	}
	if s.cfg.WithDebugRequest {
		reqJSON, err := protojson.Marshal(req)
		if err != nil {
			s.logger.Debug("cannot marshal CreateAlertsWithNamespace req to json", "err", err.Error())
		}
		s.logger.Debug("incoming create alert with namespace", "json", string(reqJSON))
	}
	items, err := s.createAlerts(ctx, req.GetProviderType(), req.GetProviderId(), namespaceID, req.GetBody().AsMap())
	if err != nil {
		return nil, api.GenerateRPCErr(s.logger, err)
	}

	return &sirenv1beta1.CreateAlertsWithNamespaceResponse{
		Alerts: items,
	}, nil
}

func (s *GRPCServer) createAlerts(ctx context.Context, providerType string, providerID uint64, namespaceID uint64, body map[string]any) ([]*sirenv1beta1.Alert, error) {
	createdAlerts, err := s.alertService.CreateAlerts(ctx, providerType, providerID, namespaceID, body)
	if err != nil {
		return nil, err
	}

	items := []*sirenv1beta1.Alert{}
	for _, item := range createdAlerts {
		alertHistoryItem := &sirenv1beta1.Alert{
			Id:           item.ID,
			ProviderId:   item.ProviderID,
			ResourceName: item.ResourceName,
			MetricName:   item.MetricName,
			MetricValue:  item.MetricValue,
			Severity:     item.Severity,
			Rule:         item.Rule,
			TriggeredAt:  timestamppb.New(item.TriggeredAt),
		}
		items = append(items, alertHistoryItem)
	}

	return items, nil
}
