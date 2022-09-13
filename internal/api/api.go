package api

import (
	"context"

	"github.com/odpf/siren/core/alert"
	"github.com/odpf/siren/core/namespace"
	"github.com/odpf/siren/core/provider"
	"github.com/odpf/siren/core/receiver"
	"github.com/odpf/siren/core/rule"
	"github.com/odpf/siren/core/subscription"
	"github.com/odpf/siren/core/template"
)

//go:generate mockery --name=AlertService -r --case underscore --with-expecter --structname AlertService --filename alert_service.go --output=./mocks
type AlertService interface {
	Create(context.Context, []*alert.Alert) ([]alert.Alert, error)
	List(context.Context, alert.Filter) ([]alert.Alert, error)
}

//go:generate mockery --name=NamespaceService -r --case underscore --with-expecter --structname NamespaceService --filename namespace_service.go --output=./mocks
type NamespaceService interface {
	List(context.Context) ([]namespace.Namespace, error)
	Create(context.Context, *namespace.Namespace) error
	Get(context.Context, uint64) (*namespace.Namespace, error)
	Update(context.Context, *namespace.Namespace) error
	Delete(context.Context, uint64) error
}

//go:generate mockery --name=ProviderService -r --case underscore --with-expecter --structname ProviderService --filename provider_service.go --output=./mocks
type ProviderService interface {
	List(context.Context, provider.Filter) ([]provider.Provider, error)
	Create(context.Context, *provider.Provider) error
	Get(context.Context, uint64) (*provider.Provider, error)
	Update(context.Context, *provider.Provider) error
	Delete(context.Context, uint64) error
}

//go:generate mockery --name=ReceiverService -r --case underscore --with-expecter --structname ReceiverService --filename receiver_service.go --output=./mocks
type ReceiverService interface {
	List(ctx context.Context, flt receiver.Filter) ([]receiver.Receiver, error)
	Create(ctx context.Context, rcv *receiver.Receiver) error
	Get(ctx context.Context, id uint64) (*receiver.Receiver, error)
	Update(ctx context.Context, rcv *receiver.Receiver) error
	Delete(ctx context.Context, id uint64) error
	Notify(ctx context.Context, id uint64, payloadMessage receiver.NotificationMessage) error
	GetSubscriptionConfig(subsConfs map[string]string, rcv *receiver.Receiver) (map[string]string, error)
}

//go:generate mockery --name=RuleService -r --case underscore --with-expecter --structname RuleService --filename rule_service.go --output=./mocks
type RuleService interface {
	Upsert(context.Context, *rule.Rule) error
	List(context.Context, rule.Filter) ([]rule.Rule, error)
}

//go:generate mockery --name=SubscriptionService -r --case underscore --with-expecter --structname SubscriptionService --filename subscription_service.go --output=./mocks
type SubscriptionService interface {
	List(context.Context, subscription.Filter) ([]subscription.Subscription, error)
	Create(context.Context, *subscription.Subscription) error
	Get(context.Context, uint64) (*subscription.Subscription, error)
	Update(context.Context, *subscription.Subscription) error
	Delete(context.Context, uint64) error
}

//go:generate mockery --name=TemplateService -r --case underscore --with-expecter --structname TemplateService --filename template_service.go --output=./mocks
type TemplateService interface {
	Upsert(context.Context, *template.Template) error
	List(context.Context, template.Filter) ([]template.Template, error)
	GetByName(context.Context, string) (*template.Template, error)
	Delete(context.Context, string) error
	Render(context.Context, string, map[string]string) (string, error)
}

type Deps struct {
	TemplateService     TemplateService
	RuleService         RuleService
	AlertService        AlertService
	ProviderService     ProviderService
	NamespaceService    NamespaceService
	ReceiverService     ReceiverService
	SubscriptionService SubscriptionService
}