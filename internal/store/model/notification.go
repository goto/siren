package model

import (
	"database/sql"
	"time"

	"github.com/goto/siren/core/notification"
	"github.com/goto/siren/pkg/pgc"
)

type Notification struct {
	ID                string               `db:"id"`
	NamespaceID       sql.NullInt64        `db:"namespace_id"`
	Type              string               `db:"type"`
	Data              pgc.StringAnyMap     `db:"data"`
	Labels            pgc.StringStringMap  `db:"labels"`
	ValidDuration     pgc.TimeDuration     `db:"valid_duration"`
	UniqueKey         sql.NullString       `db:"unique_key"`
	Template          sql.NullString       `db:"template"`
	CreatedAt         time.Time            `db:"created_at"`
	ReceiverSelectors pgc.ListStringAnyMap `db:"receiver_selectors"`
}

func (n *Notification) FromDomain(d notification.Notification) {
	n.ID = d.ID
	n.Type = d.Type
	n.Data = d.Data
	n.Labels = d.Labels
	n.ValidDuration = pgc.TimeDuration(d.ValidDuration)

	if d.NamespaceID == 0 {
		n.NamespaceID = sql.NullInt64{Valid: false}
	} else {
		n.NamespaceID = sql.NullInt64{Int64: int64(d.NamespaceID), Valid: true}
	}

	if d.Template == "" {
		n.Template = sql.NullString{Valid: false}
	} else {
		n.Template = sql.NullString{String: d.Template, Valid: true}
	}

	if d.UniqueKey == "" {
		n.UniqueKey = sql.NullString{Valid: false}
	} else {
		n.UniqueKey = sql.NullString{String: d.UniqueKey, Valid: true}
	}

	n.CreatedAt = d.CreatedAt
	n.ReceiverSelectors = d.ReceiverSelectors
}

func (n *Notification) ToDomain() *notification.Notification {
	receiverSelectors := n.ReceiverSelectors
	if receiverSelectors == nil {
		receiverSelectors = []map[string]any{}
	}
	return &notification.Notification{
		ID:                n.ID,
		NamespaceID:       uint64(n.NamespaceID.Int64),
		Type:              n.Type,
		Data:              n.Data,
		Labels:            n.Labels,
		ValidDuration:     time.Duration(n.ValidDuration),
		Template:          n.Template.String,
		UniqueKey:         n.UniqueKey.String,
		CreatedAt:         n.CreatedAt,
		ReceiverSelectors: receiverSelectors,
	}
}
