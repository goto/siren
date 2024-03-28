package postgresq

import (
	"context"
	"errors"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/goto/salt/db"
	"github.com/goto/salt/log"
	"github.com/goto/siren/core/notification"
	"github.com/goto/siren/pkg/pgc"
	"github.com/goto/siren/plugins/queues/postgresq/migrations"
)

const (
	MessageQueueTableName     = "message_queue"
	MessageQueueSchemaName    = "notification"
	MessageQueueTableFullName = MessageQueueSchemaName + "." + MessageQueueTableName
)

type Strategy string

const (
	StrategyDefault Strategy = "default"
	StrategyDLQ     Strategy = "dlq"
)

type Queue struct {
	logger   log.Logger
	pgClient *pgc.Client
	strategy Strategy
}

var (
	successCallbackQuery = fmt.Sprintf(`
UPDATE %s
SET updated_at = $1, status = $2, try_count = $3
WHERE id = $4
`, MessageQueueTableFullName)

	errorCallbackQuery = fmt.Sprintf(`
UPDATE %s
SET updated_at = $1, status = $2, try_count = $3, last_error = $4, retryable = $5
WHERE id = $6
`, MessageQueueTableFullName)

	queueEnqueueNamedQuery = fmt.Sprintf(`
INSERT INTO %s
	(id, notification_id, status, receiver_type, configs, details, last_error, max_tries, try_count, retryable, expired_at, created_at, updated_at)
    VALUES (:id,:notification_id,:status,:receiver_type,:configs,:details,:last_error,:max_tries,:try_count,:retryable,:expired_at,:created_at,:updated_at)
`, MessageQueueTableFullName)

	messagesListQueryBuilder = sq.Select(
		"id",
		"notification_id",
		"status",
		"receiver_type",
		"details",
		"last_error",
		"max_tries",
		"try_count",
		"retryable",
		"expired_at",
		"created_at",
		"updated_at",
	).From(MessageQueueTableFullName)
)

func getQueueDequeueQuery(batchSize int, receiverTypesList string) string {
	return fmt.Sprintf(`
UPDATE %s
SET status = '%s', updated_at = now()
WHERE id IN (
    SELECT id
    FROM %s
    WHERE (status = '%s' OR status = '%s') AND retryable IS FALSE %s AND (expired_at < now() OR expired_at IS NULL) AND try_count < max_tries AND last_error IS NULL
    ORDER BY expired_at
    FOR UPDATE SKIP LOCKED
    LIMIT %d
)
RETURNING *
`, MessageQueueTableFullName, notification.MessageStatusPending, MessageQueueTableFullName, notification.MessageStatusEnqueued, notification.MessageStatusPending, receiverTypesList, batchSize)
}

func getDLQDequeueQuery(batchSize int, receiverTypesList string) string {
	return fmt.Sprintf(`
UPDATE %s
SET status = '%s', updated_at = now()
WHERE id IN (
    SELECT id
    FROM %s
    WHERE (status = '%s' OR status = '%s') AND retryable IS TRUE  %s AND (expired_at < now() OR expired_at IS NULL) AND try_count < max_tries AND last_error IS NOT NULL
    ORDER BY expired_at
    FOR UPDATE SKIP LOCKED
    LIMIT %d
)
RETURNING *
`, MessageQueueTableFullName, notification.MessageStatusPending, MessageQueueTableFullName, notification.MessageStatusFailed, notification.MessageStatusPending, receiverTypesList, batchSize)
}

// New creates a new queue instance
func New(logger log.Logger, dbConfig db.Config, opts ...QueueOption) (*Queue, error) {
	q := &Queue{
		logger:   logger,
		strategy: StrategyDefault,
	}

	pgClient, err := pgc.NewClient(logger, dbConfig)
	if err != nil {
		return nil, fmt.Errorf("error creating postgres queue client: %w", err)
	}

	q.pgClient = pgClient

	// create schema if not exist
	_, err = pgClient.ExecContext(context.Background(), fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", MessageQueueSchemaName))
	if err != nil {
		return nil, fmt.Errorf("failed to create notification schema: %w", err)
	}

	dbConfig.URL = dbConfig.URL + fmt.Sprintf("&search_path=%s", MessageQueueSchemaName)

	if err := db.RunMigrations(dbConfig, migrations.FS, migrations.ResourcePath); err != nil {
		return nil, fmt.Errorf("error migrating postgres queue: %w", err)
	}

	for _, opt := range opts {
		opt(q)
	}

	return q, nil
}

// Dequeue pop the queue based on specific filters (receiver types or batch size) and process the messages with handlerFn
// message left in pending state that has expired or been updated long time ago means there was a failure when transforming row into a struct
func (q *Queue) Dequeue(ctx context.Context, receiverTypes []string, batchSize int, handlerFn func(context.Context, []notification.Message) error) error {
	messages := []notification.Message{}

	receiverTypesQuery := getFilterReceiverTypes(receiverTypes)

	var dequeueQuery string
	if q.strategy == StrategyDLQ {
		dequeueQuery = getDLQDequeueQuery(batchSize, receiverTypesQuery)
	} else {
		dequeueQuery = getQueueDequeueQuery(batchSize, receiverTypesQuery)
	}

	rows, err := q.pgClient.GetDB(ctx).QueryxContext(ctx, dequeueQuery)
	if err != nil {
		return err
	}
	for rows.Next() {
		msg := NotificationMessage{}
		if err := rows.StructScan(&msg); err != nil {
			q.logger.Error("failed to transform message row into struct", "strategy", q.strategy, "error", err)
			continue
		}
		msgDomain := msg.ToDomain()

		messages = append(messages, msgDomain)
	}

	if len(messages) == 0 {
		return notification.ErrNoMessage
	} else {
		q.logger.Debug(fmt.Sprintf("dequeued %d messages with batch size %d", len(messages), batchSize), "strategy", q.strategy)
		if err := handlerFn(ctx, messages); err != nil {
			return fmt.Errorf("error processing dequeued message: %w", err)
		}
	}

	return nil
}

// Enqueue pushes messages to the queue
func (q *Queue) Enqueue(ctx context.Context, ms ...notification.Message) error {
	messages := []NotificationMessage{}
	for _, m := range ms {
		message := &NotificationMessage{}
		message.FromDomain(m)

		messages = append(messages, *message)
	}

	res, err := q.pgClient.NamedExecContext(ctx, queueEnqueueNamedQuery, messages)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows affected when enqueueing messages")
	}
	return nil
}

// SuccessCallback is a callback that will be called once the message is succesfully handled by handlerFn
func (q *Queue) SuccessCallback(ctx context.Context, ms notification.Message) error {
	q.logger.Debug("marking a message as published", "strategy", q.strategy, "id", ms.ID)
	res, err := q.pgClient.ExecContext(ctx, successCallbackQuery, ms.UpdatedAt, ms.Status, ms.TryCount, ms.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows affected when marking row as published")
	}
	q.logger.Debug("marked a message as published", "strategy", q.strategy, "id", ms.ID)
	return nil
}

// ErrorCallback is a callback that will be called once the message is failed to be handled by handlerFn
func (q *Queue) ErrorCallback(ctx context.Context, ms notification.Message) error {
	q.logger.Debug("marking a message as failed with", "strategy", q.strategy, "id", ms.ID)
	res, err := q.pgClient.ExecContext(ctx, errorCallbackQuery, ms.UpdatedAt, ms.Status, ms.TryCount, ms.LastError, ms.Retryable, ms.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows affected when marking row as failed")
	}
	q.logger.Debug("marked a message as failed with", "strategy", q.strategy, "id", ms.ID)
	return nil
}

// Stop will close the db
func (q *Queue) Stop(ctx context.Context) error {
	return q.pgClient.Close()
}

func (q *Queue) ListMessages(ctx context.Context, notificationID string) ([]notification.Message, error) {
	var queryBuilder = messagesListQueryBuilder.Where("notification_id = ?", notificationID)

	query, args, err := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := q.pgClient.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	messagesDomain := []notification.Message{}
	for rows.Next() {
		var messageModel NotificationMessage
		if err := rows.StructScan(&messageModel); err != nil {
			return nil, err
		}

		messagesDomain = append(messagesDomain, messageModel.ToDomain())
	}

	return messagesDomain, nil
}

func getFilterReceiverTypes(receiverTypes []string) string {
	var receiverTypesQuery = ""
	if len(receiverTypes) > 0 {
		receiverTypesQuery = "AND receiver_type IN ("
		for _, rs := range receiverTypes {
			receiverTypesQuery += "'"
			receiverTypesQuery += rs
			receiverTypesQuery += "'"
			receiverTypesQuery += ","
		}
		receiverTypesQuery = strings.TrimSuffix(receiverTypesQuery, ",")
		receiverTypesQuery += ")"
	}
	return receiverTypesQuery
}
