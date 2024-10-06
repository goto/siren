package v1beta1

import (
	"context"
	"fmt"
	"strings"

	"github.com/goto/siren/core/notification"
	"github.com/goto/siren/core/template"
	"github.com/goto/siren/internal/api"
	"github.com/goto/siren/pkg/errors"
	sirenv1beta1 "github.com/goto/siren/proto/gotocompany/siren/v1beta1"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const notificationAPIScope = "notification_api"

func (s *GRPCServer) parseReceivers(pbSelectors []*structpb.Struct) ([]map[string]interface{}, error) {
	var receiverSelectors []map[string]interface{}

	for _, pbSelector := range pbSelectors {
		selector := make(map[string]interface{})
		for k, v := range pbSelector.AsMap() {
			if k == "config" {
				configMap, ok := v.(map[string]interface{})
				if !ok {
					return nil, errors.ErrInvalid.WithMsgf("invalid config format, expected map[string]interface{}")
				}
				selector[k] = configMap
			} else {
				selector[k] = v
			}
		}
		receiverSelectors = append(receiverSelectors, selector)
	}

	return receiverSelectors, nil
}

func (s *GRPCServer) validatePostNotificationPayload(receiverSelectors []map[string]interface{}, labels map[string]string) error {
	if len(receiverSelectors) == 0 && len(labels) == 0 {
		return errors.ErrInvalid.WithMsgf("either receivers or labels must be provided")
	}
	if len(receiverSelectors) > 0 && len(labels) > 0 {
		return errors.ErrInvalid.WithMsgf("receivers and labels cannot be used at the same time, should use either one of them")
	}

	return nil
}

func (s *GRPCServer) PostNotification(ctx context.Context, req *sirenv1beta1.PostNotificationRequest) (*sirenv1beta1.PostNotificationResponse, error) {
	idempotencyScope := api.GetHeaderString(ctx, s.headers.IdempotencyScope)
	if idempotencyScope == "" {
		idempotencyScope = notificationAPIScope
	}

	idempotencyKey := api.GetHeaderString(ctx, s.headers.IdempotencyKey)
	if idempotencyKey != "" {
		if notificationID, err := s.notificationService.CheckIdempotency(ctx, idempotencyScope, idempotencyKey); notificationID != "" {
			return &sirenv1beta1.PostNotificationResponse{
				NotificationId: notificationID,
			}, nil
		} else if errors.Is(err, errors.ErrNotFound) {
			s.logger.Debug("no idempotency found with detail", "scope", idempotencyScope, "key", idempotencyKey)
		} else {
			return nil, api.GenerateRPCErr(s.logger, fmt.Errorf("error when checking idempotency: %w", err))
		}
	}

	receiverSelectors, err := s.parseReceivers(req.GetReceivers())
	if err != nil {
		return nil, api.GenerateRPCErr(s.logger, err)
	}

	if err := s.validatePostNotificationPayload(receiverSelectors, req.GetLabels()); err != nil {
		return nil, api.GenerateRPCErr(s.logger, err)
	}

	var notificationTemplate = template.ReservedName_SystemDefault
	if req.GetTemplate() != "" {
		notificationTemplate = req.GetTemplate()
	}

	s.metricNotificationReceiverSelectorCount.Record(
		ctx, int64(len(receiverSelectors)),
		metric.WithAttributes(
			attribute.String("template", notificationTemplate),
		),
	)

	notificationIDs, err := s.notificationService.Dispatch(ctx, []notification.Notification{
		{
			Type:              notification.TypeEvent,
			Data:              req.GetData().AsMap(),
			Labels:            req.GetLabels(),
			Template:          notificationTemplate,
			ReceiverSelectors: receiverSelectors,
		},
	})
	if err != nil {
		if errors.Is(err, notification.ErrNoMessage) {
			return nil, api.GenerateRPCErr(s.logger, errors.ErrInvalid.WithMsgf(err.Error()))
		}
		return nil, api.GenerateRPCErr(s.logger, err)
	}

	if len(notificationIDs) != 1 {
		return nil, api.GenerateRPCErr(s.logger, errors.ErrInternal.WithMsgf("should send 1 notification only but got %d", len(notificationIDs)))
	}

	if idempotencyKey != "" {
		if err := s.notificationService.InsertIdempotency(ctx, idempotencyScope, idempotencyKey, notificationIDs[0]); err != nil {
			return nil, api.GenerateRPCErr(s.logger, err)
		}
	}

	return &sirenv1beta1.PostNotificationResponse{
		NotificationId: notificationIDs[0],
	}, nil
}

// func (s *GRPCServer) convertNotificationToPb(n notification.Notification) (*sirenv1beta1.Notification, error) {
// 	pbData, err := structpb.NewStruct(n.Data)
// 	if err != nil {
// 		return nil, errors.ErrInternal.WithMsgf("failed to convert notification data to protobuf: %v", err)
// 	}

// 	return &sirenv1beta1.Notification{
// 		Id:            n.ID,
// 		NamespaceId:   n.NamespaceID,
// 		Type:          n.Type,
// 		Data:          pbData,
// 		Labels:        n.Labels,
// 		ValidDuration: durationpb.New(n.ValidDuration),
// 		Template:      n.Template,
// 		CreateAt:      timestamppb.New(n.CreatedAt),
// 		UniqueKey:     n.UniqueKey,
// 	}, nil
// }

func (s *GRPCServer) PostBulkNotifications(ctx context.Context, req *sirenv1beta1.PostBulkNotificationsRequest) (*sirenv1beta1.PostBulkNotificationsResponse, error) {
	if len(req.GetNotifications()) == 0 {
		return nil, api.GenerateRPCErr(s.logger, errors.ErrInvalid.WithMsgf("no bulk notifications found"))
	}
	idempotencyScope := api.GetHeaderString(ctx, s.headers.IdempotencyScope)
	if idempotencyScope == "" {
		idempotencyScope = notificationAPIScope
	}

	idempotencyKey := api.GetHeaderString(ctx, s.headers.IdempotencyKey)
	if idempotencyKey != "" {
		if notificationIDs, err := s.notificationService.CheckIdempotency(ctx, idempotencyScope, idempotencyKey); notificationIDs != "" {
			splittedNotificationIDs := strings.Split(notificationIDs, ",")
			return &sirenv1beta1.PostBulkNotificationsResponse{
				NotificationIds: splittedNotificationIDs,
			}, nil
		} else if errors.Is(err, errors.ErrNotFound) {
			s.logger.Debug("no idempotency found with detail", "scope", idempotencyScope, "key", idempotencyKey)
		} else {
			return nil, api.GenerateRPCErr(s.logger, fmt.Errorf("error when checking idempotency: %w", err))
		}
	}

	notifications := []notification.Notification{}

	for _, nProto := range req.GetNotifications() {
		var notificationTemplate = template.ReservedName_SystemDefault
		if nProto.GetTemplate() != "" {
			notificationTemplate = nProto.GetTemplate()
		}
		notifications = append(notifications, notification.Notification{
			NamespaceID:   nProto.GetNamespaceId(),
			Type:          notification.TypeEvent,
			Data:          nProto.GetData().AsMap(),
			Labels:        nProto.GetLabels(),
			ValidDuration: nProto.GetValidDuration().AsDuration(),
			Template:      notificationTemplate,
		})
	}

	s.metricBulkNotificationsCount.Record(
		ctx, int64(len(notifications)),
	)

	notificationIDs, err := s.notificationService.Dispatch(ctx, notifications)
	if err != nil {
		if errors.Is(err, notification.ErrNoMessage) {
			return nil, api.GenerateRPCErr(s.logger, errors.ErrInvalid.WithMsgf(err.Error()))
		}
		return nil, api.GenerateRPCErr(s.logger, err)
	}

	if len(notificationIDs) == 0 {
		return &sirenv1beta1.PostBulkNotificationsResponse{
			NotificationIds: []string{},
		}, nil
	}

	if idempotencyKey != "" {
		if err := s.notificationService.InsertIdempotency(ctx, idempotencyScope, idempotencyKey, strings.Join(notificationIDs, ",")); err != nil {
			return nil, api.GenerateRPCErr(s.logger, err)
		}
	}

	return &sirenv1beta1.PostBulkNotificationsResponse{
		NotificationIds: notificationIDs,
	}, nil
}

func (s *GRPCServer) ListNotificationMessages(ctx context.Context, req *sirenv1beta1.ListNotificationMessagesRequest) (*sirenv1beta1.ListNotificationMessagesResponse, error) {
	resp, err := s.notificationService.ListNotificationMessages(ctx, req.GetNotificationId())
	if err != nil {
		return nil, api.GenerateRPCErr(s.logger, err)
	}

	items := []*sirenv1beta1.NotificationMessage{}
	for _, msg := range resp {
		item, err := msg.ToV1beta1Proto()
		if err != nil {
			return nil, api.GenerateRPCErr(s.logger, err)
		}

		items = append(items, item)
	}
	return &sirenv1beta1.ListNotificationMessagesResponse{
		Messages: items,
	}, nil
}

func (s *GRPCServer) ListNotifications(ctx context.Context, req *sirenv1beta1.ListNotificationsRequest) (*sirenv1beta1.ListNotificationsResponse, error) {
	notifications, err := s.notificationService.List(ctx, notification.Filter{
		Type:             req.GetType(),
		Template:         req.GetTemplate(),
		Labels:           req.GetLabels(),
		ReceiverSelector: req.GetReceiverSelectors(),
	})
	if err != nil {
		return nil, api.GenerateRPCErr(s.logger, err)
	}

	items := []*sirenv1beta1.Notification{}

	for _, notification := range notifications {
		data, err := structpb.NewStruct(notification.Data)
		if err != nil {
			return nil, api.GenerateRPCErr(s.logger, err)
		}

		item := &sirenv1beta1.Notification{
			Id:            notification.ID,
			NamespaceId:   notification.NamespaceID,
			Type:          notification.Type,
			Data:          data,
			Labels:        notification.Labels,
			ValidDuration: durationpb.New(notification.ValidDuration),
			Template:      notification.Template,
			CreateAt:      timestamppb.New(notification.CreatedAt),
			UniqueKey:     notification.UniqueKey,
		}
		items = append(items, item)
	}
	return &sirenv1beta1.ListNotificationsResponse{
		Notifications: items,
	}, nil
}
