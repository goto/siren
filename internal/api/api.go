package api

import (
	"context"
	"time"

	"github.com/goto/salt/log"
	"github.com/goto/siren/core/alert"
	"github.com/goto/siren/core/namespace"
	"github.com/goto/siren/core/notification"
	"github.com/goto/siren/core/provider"
	"github.com/goto/siren/core/receiver"
	"github.com/goto/siren/core/rule"
	"github.com/goto/siren/core/silence"
	"github.com/goto/siren/core/subscription"
	"github.com/goto/siren/core/subscriptionreceiver"
	"github.com/goto/siren/core/template"
	"github.com/goto/siren/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AlertService interface {
	CreateAlerts(ctx context.Context, providerType string, providerID uint64, namespaceID uint64, body map[string]any) ([]alert.Alert, error)
	List(context.Context, alert.Filter) ([]alert.Alert, error)
}

type NamespaceService interface {
	List(context.Context) ([]namespace.Namespace, error)
	Create(context.Context, *namespace.Namespace) error
	Get(context.Context, uint64) (*namespace.Namespace, error)
	Update(context.Context, *namespace.Namespace) error
	Delete(context.Context, uint64) error
}

type ProviderService interface {
	List(context.Context, provider.Filter) ([]provider.Provider, error)
	Create(context.Context, *provider.Provider) error
	Get(context.Context, uint64) (*provider.Provider, error)
	Update(context.Context, *provider.Provider) error
	Delete(context.Context, uint64) error
}

type ReceiverService interface {
	List(ctx context.Context, flt receiver.Filter) ([]receiver.Receiver, error)
	Create(ctx context.Context, rcv *receiver.Receiver) error
	Get(ctx context.Context, id uint64, gopts ...receiver.GetOption) (*receiver.Receiver, error)
	Update(ctx context.Context, rcv *receiver.Receiver) error
	Delete(ctx context.Context, id uint64) error
}

type RuleService interface {
	Upsert(context.Context, *rule.Rule) error
	List(context.Context, rule.Filter) ([]rule.Rule, error)
}

type SubscriptionService interface {
	ListV2(context.Context, subscription.Filter) ([]subscription.Subscription, error)
	CreateV2(context.Context, *subscription.Subscription) error
	GetV2(context.Context, uint64) (*subscription.Subscription, error)
	UpdateV2(context.Context, *subscription.Subscription) error
	DeleteV2(context.Context, uint64) error
}

type SubscriptionReceiverService interface {
	List(ctx context.Context, flt subscriptionreceiver.Filter) ([]subscriptionreceiver.Relation, error)
	BulkUpsert(ctx context.Context, rels []subscriptionreceiver.Relation) error
	Update(ctx context.Context, rel *subscriptionreceiver.Relation) error
	BulkSoftDelete(ctx context.Context, flt subscriptionreceiver.DeleteFilter) error
}

type TemplateService interface {
	Upsert(context.Context, *template.Template) error
	List(context.Context, template.Filter) ([]template.Template, error)
	GetByName(context.Context, string) (*template.Template, error)
	Delete(context.Context, string) error
	Render(context.Context, string, map[string]string) (string, error)
}

type NotificationService interface {
	Dispatch(context.Context, []notification.Notification) ([]string, error)
	RemoveIdempotencies(ctx context.Context, TTL time.Duration) error
	CheckIdempotency(ctx context.Context, scope, key string) (string, error)
	InsertIdempotency(ctx context.Context, scope, key, notificationID string) error
	ListNotificationMessages(ctx context.Context, notificationID string) ([]notification.Message, error)
	List(ctx context.Context, flt notification.Filter) ([]notification.Notification, error)
}

type SilenceService interface {
	Create(ctx context.Context, sil silence.Silence) (string, error)
	List(ctx context.Context, filter silence.Filter) ([]silence.Silence, error)
	Get(ctx context.Context, id string) (silence.Silence, error)
	Delete(ctx context.Context, id string) error
}

type Deps struct {
	TemplateService             TemplateService
	RuleService                 RuleService
	AlertService                AlertService
	ProviderService             ProviderService
	NamespaceService            NamespaceService
	ReceiverService             ReceiverService
	SubscriptionService         SubscriptionService
	SubscriptionReceiverService SubscriptionReceiverService
	NotificationService         NotificationService
	SilenceService              SilenceService
}

func GenerateRPCErr(logger log.Logger, e error) error {
	var err = errors.E(e)

	var code codes.Code
	switch {
	case errors.Is(err, errors.ErrNotFound):
		code = codes.NotFound

	case errors.Is(err, errors.ErrConflict):
		code = codes.AlreadyExists

	case errors.Is(err, errors.ErrInvalid):
		code = codes.InvalidArgument

	default:
		// TODO This will create 2 logs, grpc log and
		// the error detail (Message & Cause) log
		// there might be a better approach to solve this
		code = codes.Internal
		logger.Error(errors.Verbose(err).Error())
	}

	return status.Error(code, err.Error())
}
