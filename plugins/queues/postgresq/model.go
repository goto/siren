package postgresq

import (
	"database/sql"
	"time"

	"github.com/goto/siren/core/notification"
	"github.com/goto/siren/pkg/pgc"
)

type NotificationMessage struct {
	ID              string         `db:"id"`
	NotificationIDs pgc.ListString `db:"notification_ids"`

	Status string `db:"status"`

	ReceiverType string           `db:"receiver_type"`
	Configs      pgc.StringAnyMap `db:"configs"`
	Details      pgc.StringAnyMap `db:"details"`
	Metadata     pgc.StringAnyMap `db:"metadata"`
	LastError    sql.NullString   `db:"last_error"`

	MaxTries  int  `db:"max_tries"`
	TryCount  int  `db:"try_count"`
	Retryable bool `db:"retryable"`

	ExpiredAt sql.NullTime `db:"expired_at"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
}

func (nm *NotificationMessage) FromDomain(domainMessage notification.Message) {
	nm.ID = domainMessage.ID
	nm.NotificationIDs = domainMessage.NotificationIDs
	nm.Status = string(domainMessage.Status)
	nm.ReceiverType = domainMessage.ReceiverType
	nm.Configs = domainMessage.Configs
	nm.Details = domainMessage.Details

	nm.LastError = sql.NullString{String: domainMessage.LastError, Valid: func() bool {
		if domainMessage.LastError == "" {
			return false
		} else {
			return true
		}
	}()}
	nm.MaxTries = domainMessage.MaxTries
	nm.TryCount = domainMessage.TryCount
	nm.Retryable = domainMessage.Retryable
	nm.ExpiredAt = sql.NullTime{Time: domainMessage.ExpiredAt, Valid: func() bool {
		if domainMessage.ExpiredAt.IsZero() {
			return false
		} else {
			return true
		}
	}()}
	nm.CreatedAt = domainMessage.CreatedAt
	nm.UpdatedAt = domainMessage.UpdatedAt
}

func (nm *NotificationMessage) ToDomain() notification.Message {
	return notification.Message{
		ID:              nm.ID,
		NotificationIDs: nm.NotificationIDs,
		Status:          notification.MessageStatus(nm.Status),

		ReceiverType: nm.ReceiverType,
		Configs:      nm.Configs,
		Details:      nm.Details,
		LastError:    nm.LastError.String,

		MaxTries:  nm.MaxTries,
		TryCount:  nm.TryCount,
		Retryable: nm.Retryable,

		ExpiredAt: nm.ExpiredAt.Time,
		CreatedAt: nm.CreatedAt,
		UpdatedAt: nm.UpdatedAt,
	}
}
