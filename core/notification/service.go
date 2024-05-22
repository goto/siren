package notification

import (
	"context"
	"fmt"
	"time"

	saltlog "github.com/goto/salt/log"

	"github.com/goto/siren/core/log"
	"github.com/goto/siren/core/receiver"
	"github.com/goto/siren/core/silence"
	"github.com/goto/siren/core/subscription"
	"github.com/goto/siren/core/template"
	"github.com/goto/siren/pkg/errors"
)

type Dispatcher interface {
	PrepareMessage(ctx context.Context, n Notification) ([]Message, []log.Notification, bool, error)
	PrepareMessageV2(ctx context.Context, n Notification) ([]Message, []log.Notification, bool, error)
}

type SubscriptionService interface {
	MatchByLabels(ctx context.Context, namespaceID uint64, labels map[string]string) ([]subscription.Subscription, error)
	MatchByLabelsV2(ctx context.Context, namespaceID uint64, labels map[string]string) ([]subscription.ReceiverView, error)
}

type ReceiverService interface {
	List(ctx context.Context, flt receiver.Filter) ([]receiver.Receiver, error)
}

type SilenceService interface {
	List(ctx context.Context, filter silence.Filter) ([]silence.Silence, error)
}

type AlertRepository interface {
	BulkUpdateSilence(context.Context, []int64, string) error
}

type LogService interface {
	LogNotifications(ctx context.Context, nlogs ...log.Notification) error
}

type TemplateService interface {
	GetByName(ctx context.Context, name string) (*template.Template, error)
}

// Service is a service for notification domain
type Service struct {
	logger                saltlog.Logger
	cfg                   Config
	q                     Queuer
	idempotencyRepository IdempotencyRepository
	alertRepository       AlertRepository
	logService            LogService
	repository            Repository
	receiverService       ReceiverService
	subscriptionService   SubscriptionService
	silenceService        SilenceService
	notifierPlugins       map[string]Notifier
	dispatcher            map[string]Dispatcher
	enableSilenceFeature  bool
}

type Deps struct {
	IdempotencyRepository     IdempotencyRepository
	AlertRepository           AlertRepository
	LogService                LogService
	ReceiverService           ReceiverService
	TemplateService           TemplateService
	SubscriptionService       SubscriptionService
	SilenceService            SilenceService
	DispatchReceiverService   Dispatcher
	DispatchSubscriberService Dispatcher
}

// NewService creates a new notification service
func NewService(
	logger saltlog.Logger,
	cfg Config,
	repository Repository,
	q Queuer,
	notifierPlugins map[string]Notifier,
	deps Deps,
	enableSilenceFeature bool,
) *Service {
	var (
		dispatchReceiverService   = deps.DispatchReceiverService
		dispatchSubscriberService = deps.DispatchSubscriberService
	)
	if deps.DispatchReceiverService == nil {
		dispatchReceiverService = NewDispatchReceiverService(DispatchReceiverConfig{
			MaxMessagesReceiverFlow: cfg.MaxMessagesReceiverFlow,
			MaxNumReceiverSelectors: cfg.MaxNumReceiverSelectors,
		}, deps.ReceiverService, deps.TemplateService, notifierPlugins)
	}
	if deps.DispatchSubscriberService == nil {
		dispatchSubscriberService = NewDispatchSubscriberService(logger, deps.SubscriptionService, deps.SilenceService, deps.TemplateService, notifierPlugins, enableSilenceFeature)
	}

	ns := &Service{
		logger:                logger,
		cfg:                   cfg,
		q:                     q,
		repository:            repository,
		idempotencyRepository: deps.IdempotencyRepository,
		alertRepository:       deps.AlertRepository,
		logService:            deps.LogService,
		receiverService:       deps.ReceiverService,
		subscriptionService:   deps.SubscriptionService,
		silenceService:        deps.SilenceService,
		dispatcher: map[string]Dispatcher{
			FlowReceiver:   dispatchReceiverService,
			FlowSubscriber: dispatchSubscriberService,
		},
		notifierPlugins:      notifierPlugins,
		enableSilenceFeature: enableSilenceFeature,
	}

	return ns
}

func (s *Service) getDispatcherFlowService(notificationFlow string) (Dispatcher, error) {
	selectedDispatcher, exist := s.dispatcher[notificationFlow]
	if !exist {
		return nil, errors.ErrInvalid.WithMsgf("unsupported notification type: %q", notificationFlow)
	}
	return selectedDispatcher, nil
}

func (s *Service) Dispatch(ctx context.Context, n Notification) (string, error) {
	ctx = s.repository.WithTransaction(ctx)
	no, err := s.repository.Create(ctx, n)
	if err != nil {
		if err := s.repository.Rollback(ctx, err); err != nil {
			return "", err
		}
		return "", err
	}

	n.EnrichID(no.ID)

	switch n.Type {
	case TypeAlert:
		if err := s.dispatchAlerts(ctx, n); err != nil {
			if err := s.repository.Rollback(ctx, err); err != nil {
				return "", err
			}
			return "", err
		}
	case TypeEvent:
		if err := s.dispatchEvents(ctx, n); err != nil {
			if err := s.repository.Rollback(ctx, err); err != nil {
				return "", err
			}
			return "", err
		}
	default:
		if err := s.repository.Rollback(ctx, err); err != nil {
			return "", err
		}
		return "", errors.ErrInternal.WithMsgf("unknown notification type")
	}

	if err := s.repository.Commit(ctx); err != nil {
		return "", err
	}

	return n.ID, nil
}

func (s *Service) dispatchByFlow(ctx context.Context, n Notification, flow string) error {
	if err := n.Validate(flow); err != nil {
		return err
	}

	dispatcherService, err := s.getDispatcherFlowService(flow)
	if err != nil {
		return err
	}

	var (
		messages         []Message
		notificationLogs []log.Notification
		hasSilenced      bool
	)
	if s.cfg.SubscriptionV2Enabled {
		messages, notificationLogs, hasSilenced, err = dispatcherService.PrepareMessageV2(ctx, n)
		if err != nil {
			return err
		}
	} else {
		messages, notificationLogs, hasSilenced, err = dispatcherService.PrepareMessage(ctx, n)
		if err != nil {
			return err
		}
	}

	if len(messages) == 0 && len(notificationLogs) == 0 {
		return fmt.Errorf("something wrong and no messages will be sent with notification: %v", n)
	}

	if err := s.logService.LogNotifications(ctx, notificationLogs...); err != nil {
		return fmt.Errorf("failed logging notifications: %w", err)
	}

	// Reliability of silence feature need to be tested more
	if s.enableSilenceFeature {
		if err := s.alertRepository.BulkUpdateSilence(ctx, n.AlertIDs, silence.Status(hasSilenced, len(messages) != 0)); err != nil {
			return fmt.Errorf("failed updating silence status: %w", err)
		}
	}

	if len(messages) == 0 {
		s.logger.Info("no messages to enqueue")
		return nil
	}

	if err := s.q.Enqueue(ctx, messages...); err != nil {
		return fmt.Errorf("failed enqueuing messages: %w", err)
	}

	return nil
}

func (s *Service) dispatchEvents(ctx context.Context, n Notification) error {
	if len(n.ReceiverSelectors) == 0 && len(n.Labels) == 0 {
		return errors.ErrInvalid.WithMsgf("no receivers found")
	}

	if len(n.ReceiverSelectors) != 0 {
		if err := s.dispatchByFlow(ctx, n, FlowReceiver); err != nil {
			return err
		}
	}

	if len(n.Labels) != 0 {
		if err := s.dispatchByFlow(ctx, n, FlowSubscriber); err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) dispatchAlerts(ctx context.Context, n Notification) error {
	if err := s.dispatchByFlow(ctx, n, FlowSubscriber); err != nil {
		return err
	}

	return nil
}

func (s *Service) CheckIdempotency(ctx context.Context, scope, key string) (string, error) {
	idempt, err := s.idempotencyRepository.Check(ctx, scope, key)
	if err != nil {
		return "", err
	}

	return idempt.NotificationID, nil
}

func (s *Service) InsertIdempotency(ctx context.Context, scope, key, notificationID string) error {
	if _, err := s.idempotencyRepository.Create(ctx, scope, key, notificationID); err != nil {
		return err
	}

	return nil
}

func (s *Service) RemoveIdempotencies(ctx context.Context, TTL time.Duration) error {
	return s.idempotencyRepository.Delete(ctx, IdempotencyFilter{
		TTL: TTL,
	})
}

func (s *Service) ListNotificationMessages(ctx context.Context, notificationID string) ([]Message, error) {
	messages, err := s.q.ListMessages(ctx, notificationID)
	if err != nil {
		return nil, err
	}

	messages = s.discardSecrets(messages)

	return messages, nil
}

// TODO might want to do smarter way to discard secrets
func (s *Service) discardSecrets(messages []Message) []Message {
	newMessages := make([]Message, 0)

	for _, msg := range messages {
		newMsg := msg
		cfg := newMsg.Configs
		// slack token
		delete(cfg, "token")
		// pagerduty service key
		delete(cfg, "service_key")
		newMsg.Configs = cfg
		newMessages = append(newMessages, newMsg)
	}

	return newMessages
}

func (s *Service) List(ctx context.Context, flt Filter) ([]Notification, error) {
	notifications, err := s.repository.List(ctx, flt)
	if err != nil {
		return nil, err
	}

	return notifications, err
}
