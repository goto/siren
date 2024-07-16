package subscription

import (
	context "context"
	"fmt"
	"time"

	"github.com/goto/siren/core/receiver"
	"github.com/goto/siren/core/silence"
	"github.com/goto/siren/core/subscriptionreceiver"
)

type Transactor interface {
	WithTransaction(ctx context.Context) context.Context
	Rollback(ctx context.Context, err error) error
	Commit(ctx context.Context) error
}

type Repository interface {
	Transactor
	List(context.Context, Filter) ([]Subscription, error)
	Create(context.Context, *Subscription) error
	Get(context.Context, uint64) (*Subscription, error)
	Update(context.Context, *Subscription) error
	Delete(context.Context, uint64) error
	MatchLabelsFetchReceivers(ctx context.Context, flt Filter) ([]ReceiverView, error)
}

type ReceiverView struct {
	ID             uint64            `json:"id"` // receiver_id
	Name           string            `json:"name"`
	Labels         map[string]string `json:"labels"`
	Type           string            `json:"type"`
	Configurations map[string]any    `json:"configurations"`
	ParentID       uint64            `json:"parent_id"`
	CreatedAt      time.Time         `json:"created_at"`
	UpdatedAt      time.Time         `json:"updated_at"`
	SubscriptionID uint64            `json:"subscription_id"`
	Match          map[string]string `json:"match"`
}

func (rcv *ReceiverView) FromReceiver(r receiver.Receiver) {
	rcv.ID = r.ID
	rcv.Name = r.Name
	rcv.Labels = r.Labels
	rcv.Type = r.Type
	rcv.Configurations = r.Configurations
	rcv.ParentID = r.ParentID
	rcv.CreatedAt = r.CreatedAt
	rcv.UpdatedAt = r.UpdatedAt

}

type Receiver struct {
	ID                         uint64            `json:"id"`
	Configuration              map[string]any    `json:"configuration"`
	SubscriptionReceiverLabels map[string]string `json:"subscription_receiver_labels"`

	// Type won't be exposed to the end-user, this is used to add more details for notification purposes
	Type string
}

type Subscription struct {
	ID        uint64            `json:"id"`
	URN       string            `json:"urn"`
	Namespace uint64            `json:"namespace"`
	Receivers []Receiver        `json:"receivers"`
	Match     map[string]string `json:"match"`
	Metadata  map[string]any    `json:"metadata"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	CreatedBy string            `json:"created_by"`
	UpdatedBy string            `json:"updated_by"`

	// for v1 api cases
	ReceiversRelation []subscriptionreceiver.Relation `json:"receivers_relation"`
}

func (s Subscription) ReceiversAsMap() map[uint64]Receiver {
	var m = make(map[uint64]Receiver)
	for _, rcv := range s.Receivers {
		m[rcv.ID] = rcv
	}
	return m
}

func (s Subscription) SilenceReceivers(silences []silence.Silence) (map[uint64][]silence.Silence, []Receiver, error) {
	var (
		nonSilencedReceiversMap = map[uint64]Receiver{}
		silencedReceiversMap    = map[uint64][]silence.Silence{}
	)

	if len(silences) == 0 {
		return nil, s.Receivers, nil
	}

	// evaluate all receivers of subscribers with all matched silences
	for _, sil := range silences {
		for _, rcv := range s.Receivers {
			isSilenced, err := sil.EvaluateSubscriptionRule(rcv)
			if err != nil {
				return nil, nil, fmt.Errorf("error evaluating subscription receiver %v: %w", rcv, err)
			}

			if isSilenced {
				if len(silencedReceiversMap) == 0 {
					silencedReceiversMap = make(map[uint64][]silence.Silence)
				}
				silencedReceiversMap[rcv.ID] = append(silencedReceiversMap[rcv.ID], sil)
			} else {
				nonSilencedReceiversMap[rcv.ID] = rcv
			}
		}
	}

	var nonSilencedReceivers []Receiver
	for k, v := range nonSilencedReceiversMap {
		// remove if non silenced receivers are part of silenced receivers
		if _, ok := silencedReceiversMap[k]; !ok {
			nonSilencedReceivers = append(nonSilencedReceivers, v)
		}
	}

	return silencedReceiversMap, nonSilencedReceivers, nil
}
