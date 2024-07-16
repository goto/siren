package notification

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	saltlog "github.com/goto/salt/log"
	"github.com/mitchellh/hashstructure/v2"
	"golang.org/x/exp/maps"

	"github.com/goto/siren/core/log"
	"github.com/goto/siren/core/receiver"
	"github.com/goto/siren/core/silence"
	"github.com/goto/siren/core/subscription"
	"github.com/goto/siren/core/template"
	"github.com/goto/siren/pkg/errors"
	"github.com/goto/siren/pkg/structure"
)

type Router interface {
	PrepareMetaMessages(ctx context.Context, n Notification) (metaMessages []MetaMessage, notificationLogs []log.Notification, err error)
}

type Dispatcher interface {
	Dispatch(ctx context.Context, ns []Notification) ([]string, error)
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
	deps            Deps
	routerMap       map[string]Router
	notifierPlugins map[string]Notifier
}

type Deps struct {
	Cfg                   Config
	Logger                saltlog.Logger
	Repository            Repository
	Q                     Queuer
	IdempotencyRepository IdempotencyRepository
	AlertRepository       AlertRepository
	LogService            LogService
	ReceiverService       ReceiverService
	TemplateService       TemplateService
	SubscriptionService   SubscriptionService
	SilenceService        SilenceService
}

// NewService creates a new notification service
func NewService(
	deps Deps,
	routerMap map[string]Router,
	notifierPlugins map[string]Notifier,
) *Service {
	return &Service{
		deps:            deps,
		routerMap:       routerMap,
		notifierPlugins: notifierPlugins,
	}
}

func (s *Service) Dispatch(ctx context.Context, ns []Notification) ([]string, error) {
	ctx = s.deps.Repository.WithTransaction(ctx)

	insertedNotifications, err := s.deps.Repository.BulkCreate(ctx, ns)
	if err != nil {
		return nil, err
	}

	ids, err := s.dispatchInternal(ctx, insertedNotifications)
	if err != nil {
		if err := s.deps.Repository.Rollback(ctx, err); err != nil {
			return nil, err
		}
		return nil, err
	}

	if err := s.deps.Repository.Commit(ctx); err != nil {
		return nil, err
	}

	return ids, nil
}

func (s *Service) CheckIdempotency(ctx context.Context, scope, key string) (string, error) {
	idempt, err := s.deps.IdempotencyRepository.Check(ctx, scope, key)
	if err != nil {
		return "", err
	}

	return idempt.NotificationID, nil
}

func (s *Service) InsertIdempotency(ctx context.Context, scope, key, notificationID string) error {
	if _, err := s.deps.IdempotencyRepository.Create(ctx, scope, key, notificationID); err != nil {
		return err
	}

	return nil
}

func (s *Service) RemoveIdempotencies(ctx context.Context, TTL time.Duration) error {
	return s.deps.IdempotencyRepository.Delete(ctx, IdempotencyFilter{
		TTL: TTL,
	})
}

func (s *Service) ListNotificationMessages(ctx context.Context, notificationID string) ([]Message, error) {
	messages, err := s.deps.Q.ListMessages(ctx, notificationID)
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
	notifications, err := s.deps.Repository.List(ctx, flt)
	if err != nil {
		return nil, err
	}

	return notifications, err
}

func (s *Service) getRouter(notificationRouterKind string) (Router, error) {
	selectedRouter, exist := s.routerMap[notificationRouterKind]
	if !exist {
		return nil, errors.ErrInvalid.WithMsgf("unsupported notification router kind: %q", notificationRouterKind)
	}
	return selectedRouter, nil
}

func (s *Service) dispatchInternal(ctx context.Context, ns []Notification) (notificationIDs []string, err error) {
	var (
		metaMessages     []MetaMessage
		notificationLogs []log.Notification
	)

	for _, n := range ns {
		var flow string
		if len(n.ReceiverSelectors) != 0 {
			flow = RouterReceiver
		} else if len(n.Labels) != 0 {
			flow = RouterSubscriber
		} else {
			return nil, errors.ErrInvalid.WithMsgf("no receivers or labels found, unknown flow")
		}

		// NOTE: never invalid cause we have checked above
		if err := n.Validate(flow); err != nil {
			return nil, err
		}

		// TODO: test if flow is not recognized
		router, err := s.getRouter(flow)
		if err != nil {
			return nil, err
		}

		generatedMetaMessages, generatedNotificationLogs, err := router.PrepareMetaMessages(ctx, n)
		if err != nil {
			if errors.Is(err, ErrRouteSubscriberNoMatchFound) {
				errMessage := fmt.Sprintf("not matching any subscription for notification: %v", n)
				nJson, err := json.MarshalIndent(n, "", "  ")
				if err == nil {
					errMessage = fmt.Sprintf("not matching any subscription for notification: %s", string(nJson))
				}
				s.deps.Logger.Warn(errMessage)
				continue
			}
			return nil, err
		}

		metaMessages = append(metaMessages, generatedMetaMessages...)
		notificationLogs = append(notificationLogs, generatedNotificationLogs...)
	}

	messages, err := s.PrepareMessages(ctx, metaMessages)
	if err != nil {
		return nil, err
	}

	if len(messages) == 0 {
		s.deps.Logger.Info("no messages to process")
		return nil, ErrNoMessage
	}

	if err := s.deps.LogService.LogNotifications(ctx, notificationLogs...); err != nil {
		return nil, fmt.Errorf("failed logging notifications: %w", err)
	}

	if err := s.deps.Q.Enqueue(ctx, messages...); err != nil {
		return nil, fmt.Errorf("failed enqueuing messages: %w", err)
	}

	for _, n := range ns {
		notificationIDs = append(notificationIDs, n.ID)
	}

	return notificationIDs, nil
}

func (s *Service) PrepareMessages(ctx context.Context, metaMessages []MetaMessage) ([]Message, error) {
	if len(metaMessages) == 0 {
		return []Message{}, nil
	}

	reducedMetaMessages, err := ReduceMetaMessages(metaMessages, s.deps.Cfg.GroupBy)
	if err != nil {
		return nil, err
	}

	messages, err := s.RenderMessages(ctx, reducedMetaMessages)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (s *Service) RenderMessages(ctx context.Context, metaMessages []MetaMessage) (messages []Message, err error) {
	for _, mm := range metaMessages {
		notifierPlugin, err := s.getNotifierPlugin(mm.ReceiverType)
		if err != nil {
			return nil, err
		}

		message, err := InitMessageByMetaMessage(
			ctx,
			notifierPlugin,
			s.deps.TemplateService,
			mm,
			InitWithExpiryDuration(mm.ValidDuration),
		)
		if err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}
	return messages, nil
}

func (s *Service) getNotifierPlugin(receiverType string) (Notifier, error) {
	notifierPlugin, exist := s.notifierPlugins[receiverType]
	if !exist {
		return nil, errors.ErrInvalid.WithMsgf("unsupported receiver type: %q", receiverType)
	}
	return notifierPlugin, nil
}

func ReduceMetaMessages(metaMessages []MetaMessage, groupBy []string) ([]MetaMessage, error) {
	var (
		hashedMetaMessagesMap = map[uint64]MetaMessage{}
		metaMessagesResult    = []MetaMessage{}
	)
	for _, mm := range metaMessages {
		// exclude alert from reducer
		if mm.NotificationType == TypeAlert {
			metaMessagesResult = append(metaMessagesResult, mm)
			continue
		}

		groupedLabels := structure.BuildGroupLabels(mm.Labels, groupBy)
		groupedLabels["_receiver.ID"] = fmt.Sprintf("%d", mm.ReceiverID)
		groupedLabels["_notification.template"] = mm.Template

		hash, err := hashstructure.Hash(groupedLabels, hashstructure.FormatV2, nil)
		if err != nil {
			return nil, fmt.Errorf("cannot get hash from metamessage %v", mm)
		}

		if _, ok := hashedMetaMessagesMap[hash]; !ok {
			if mm.MergedLabels == nil {
				mm.MergedLabels = map[string][]string{}
				for k, v := range mm.Labels {
					mm.MergedLabels[k] = append(mm.MergedLabels[k], v)
				}
			}
			hashedMetaMessagesMap[hash] = mm

		} else {
			hashedMetaMessagesMap[hash] = MergeMetaMessage(mm, hashedMetaMessagesMap[hash])
		}

	}
	return append(metaMessagesResult, maps.Values(hashedMetaMessagesMap)...), nil
}

func MergeMetaMessage(from MetaMessage, to MetaMessage) MetaMessage {
	var output = to
	for k, v := range from.Labels {
		output.MergedLabels[k] = append(output.MergedLabels[k], v)
	}
	output.NotificationIDs = append(output.NotificationIDs, from.NotificationIDs...)
	output.SubscriptionIDs = append(output.SubscriptionIDs, from.SubscriptionIDs...)
	return output
}
