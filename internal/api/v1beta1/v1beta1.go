package v1beta1

import (
	"github.com/goto/salt/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"

	"github.com/goto/siren/internal/api"
	sirenv1beta1 "github.com/goto/siren/proto/gotocompany/siren/v1beta1"
)

// GRPCServerOption provides ability to configure the grpc initialization
type GRPCServerOption func(*GRPCServer)

func WithGlobalSubscription(useGlobalSubscription bool) GRPCServerOption {
	return func(s *GRPCServer) {
		s.cfg.UseGlobalSubscription = useGlobalSubscription
	}
}

func WithDebugRequest(debugRequest bool) GRPCServerOption {
	return func(s *GRPCServer) {
		s.cfg.WithDebugRequest = debugRequest
	}
}

type GRPCServer struct {
	logger  log.Logger
	cfg     api.Config
	headers api.HeadersConfig
	sirenv1beta1.UnimplementedSirenServiceServer
	templateService             api.TemplateService
	ruleService                 api.RuleService
	alertService                api.AlertService
	providerService             api.ProviderService
	namespaceService            api.NamespaceService
	receiverService             api.ReceiverService
	subscriptionService         api.SubscriptionService
	subscriptionReceiverService api.SubscriptionReceiverService
	notificationService         api.NotificationService
	silenceService              api.SilenceService

	metricBulkNotificationsCount            metric.Int64Gauge
	metricNotificationReceiverSelectorCount metric.Int64Gauge
}

func NewGRPCServer(
	logger log.Logger,
	headers api.HeadersConfig,
	apiDeps *api.Deps,
	opts ...GRPCServerOption) (*GRPCServer, error) {

	metricBulkNotificationsCount, err := otel.Meter("github.com/goto/siren/internal/api").
		Int64Gauge("api.bulknotifications.notifications")
	if err != nil {
		return nil, err
	}
	metricNotificationReceiverSelectorCount, err := otel.Meter("github.com/goto/siren/internal/api").
		Int64Gauge("api.notification.receiverselectors")
	if err != nil {
		return nil, err
	}

	s := &GRPCServer{
		headers:                     headers,
		logger:                      logger,
		templateService:             apiDeps.TemplateService,
		ruleService:                 apiDeps.RuleService,
		alertService:                apiDeps.AlertService,
		providerService:             apiDeps.ProviderService,
		namespaceService:            apiDeps.NamespaceService,
		receiverService:             apiDeps.ReceiverService,
		subscriptionService:         apiDeps.SubscriptionService,
		subscriptionReceiverService: apiDeps.SubscriptionReceiverService,
		notificationService:         apiDeps.NotificationService,
		silenceService:              apiDeps.SilenceService,

		metricBulkNotificationsCount:            metricBulkNotificationsCount,
		metricNotificationReceiverSelectorCount: metricNotificationReceiverSelectorCount,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s, nil
}
