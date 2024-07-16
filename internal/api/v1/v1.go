package v1

import (
	"github.com/goto/salt/log"

	"github.com/goto/siren/internal/api"
	sirenv1 "github.com/goto/siren/proto/gotocompany/siren/v1"
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
	sirenv1.UnimplementedSirenServiceServer
	logger                      log.Logger
	cfg                         api.Config
	headers                     api.HeadersConfig
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
}

func NewGRPCServer(
	logger log.Logger,
	headers api.HeadersConfig,
	apiDeps *api.Deps,
	opts ...GRPCServerOption) *GRPCServer {

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
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}
