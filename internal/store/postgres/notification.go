package postgres

import (
	"context"

	"github.com/goto/siren/core/notification"
	"github.com/goto/siren/internal/store/model"
	"github.com/goto/siren/pkg/pgc"
)

const notificationInsertQuery = `
INSERT INTO notifications (namespace_id, type, data, labels, valid_duration, template, unique_key, created_at)
    VALUES ($1, $2, $3, $4, $5, $6, $7, now())
RETURNING *
`

// NotificationRepository talks to the store to read or insert data
type NotificationRepository struct {
	client *pgc.Client
}

// NewNotificationRepository returns NotificationRepository struct
func NewNotificationRepository(client *pgc.Client) *NotificationRepository {
	return &NotificationRepository{
		client: client,
	}
}

func (r *NotificationRepository) Create(ctx context.Context, n notification.Notification) (notification.Notification, error) {
	nModel := new(model.Notification)
	nModel.FromDomain(n)

	var newNModel model.Notification
	if err := r.client.QueryRowxContext(ctx, notificationInsertQuery,
		nModel.NamespaceID,
		nModel.Type,
		nModel.Data,
		nModel.Labels,
		nModel.ValidDuration,
		nModel.Template,
		nModel.UniqueKey,
	).StructScan(&newNModel); err != nil {
		return notification.Notification{}, err
	}

	return newNModel.ToDomain(), nil
}
