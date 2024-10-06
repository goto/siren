package notification

import (
	"context"
	"time"

	"github.com/goto/siren/core/subscription"
	"github.com/goto/siren/pkg/errors"
)

const (
	ValidDurationRequestKey string = "valid_duration"

	RouterReceiver      string = "receiver"
	RouterSubscriber    string = "subscriber"
	RouterDirectChannel string = "direct_channel"

	TypeAlert string = "alert"
	TypeEvent string = "event"

	DispatchKindBulkNotification   = "bulk_notification"
	DispatchKindSingleNotification = "single_notification"
)

type Repository interface {
	Transactor
	BulkCreate(context.Context, []Notification) ([]Notification, error)
	Create(context.Context, Notification) (Notification, error)
	List(context.Context, Filter) ([]Notification, error)
}

type Transactor interface {
	WithTransaction(ctx context.Context) context.Context
	Rollback(ctx context.Context, err error) error
	Commit(ctx context.Context) error
}

// Notification is a model of notification
type Notification struct {
	ID                string                   `json:"id"`
	NamespaceID       uint64                   `json:"namespace_id"`
	Type              string                   `json:"type"`
	Data              map[string]any           `json:"data"`
	Labels            map[string]string        `json:"labels"`
	ValidDuration     time.Duration            `json:"valid_duration"`
	Template          string                   `json:"template"`
	UniqueKey         string                   `json:"unique_key"`
	ReceiverSelectors []map[string]interface{} `json:"receiver_selectors"`
	CreatedAt         time.Time                `json:"created_at"`

	// won't be stored in notification table, only to propagate this to notification_subscriber
	AlertIDs []int64
}

func (n *Notification) EnrichID(id string) {
	if n == nil {
		return
	}
	n.ID = id

	if len(n.Data) == 0 {
		n.Data = map[string]any{}
	}

	n.Data["id"] = id
}

func (n Notification) Validate(routerKind string) error {
	switch routerKind {
	case RouterReceiver:
		if len(n.ReceiverSelectors) == 0 {
			return errors.ErrInvalid.WithMsgf("notification type receiver should have receiver_selectors: %v", n)
		}
	case RouterSubscriber:
		if len(n.Labels) == 0 {
			return errors.ErrInvalid.WithMsgf("notification type subscriber should have labels: %v", n)
		}
	case RouterDirectChannel:
		if len(n.ReceiverSelectors) == 0 {
			return errors.ErrInvalid.WithMsgf("notification type direct channel should have receiver_selectors: %v", n)
		}
		if n.Type != "direct_channel" {
			return errors.ErrInvalid.WithMsgf("invalid notification type for direct channel: %v", n)
		}
		// Additional checks for direct channel can be added here if needed
	default:
		return errors.ErrInvalid.WithMsgf("invalid router kind: %v", routerKind)
	}
	return nil
}

func (n Notification) MetaMessage(receiverView subscription.ReceiverView) MetaMessage {
	m := MetaMessage{
		ReceiverID:       receiverView.ID,
		ReceiverType:     receiverView.Type,
		NotificationType: n.Type,
		ReceiverConfigs:  receiverView.Configurations,
		Data:             n.Data,
		ValidDuration:    n.ValidDuration,
		Template:         n.Template,
		Labels:           n.Labels,
		UniqueKey:        n.UniqueKey,
	}

	if receiverView.SubscriptionID != 0 {
		m.SubscriptionIDs = []uint64{receiverView.SubscriptionID}
	}

	if n.ID != "" {
		m.NotificationIDs = []string{n.ID}
	}

	return m
}
