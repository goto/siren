package v1beta1_test

import (
	"context"
	"testing"

	"github.com/goto/salt/log"
	"github.com/goto/siren/core/notification"
	"github.com/goto/siren/internal/api"
	"github.com/goto/siren/internal/api/mocks"
	"github.com/goto/siren/internal/api/v1beta1"
	"github.com/goto/siren/pkg/errors"
	sirenv1beta1 "github.com/goto/siren/proto/gotocompany/siren/v1beta1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/structpb"
)

func TestGRPCServer_PostNotification(t *testing.T) {
	const (
		idempotencyHeaderKey = "idempotency-key"
		notificationID       = "1234-5678-0987"
	)
	testCases := []struct {
		name           string
		idempotencyKey string
		setup          func(*mocks.NotificationService)
		errString      string
	}{
		{
			name:           "should return invalid argument if post notification return invalid argument",
			idempotencyKey: "test",
			setup: func(ns *mocks.NotificationService) {
				ns.EXPECT().CheckIdempotency(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("", errors.ErrNotFound)
				ns.EXPECT().Dispatch(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("[]notification.Notification")).Return(nil, errors.ErrInvalid)
			},
			errString: "rpc error: code = InvalidArgument desc = request is not valid",
		},
		{
			name:           "should return internal error if post notification return some error",
			idempotencyKey: "test",
			setup: func(ns *mocks.NotificationService) {
				ns.EXPECT().CheckIdempotency(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("", errors.ErrNotFound)
				ns.EXPECT().Dispatch(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("[]notification.Notification")).Return(nil, errors.New("some error"))
			},
			errString: "rpc error: code = Internal desc = some unexpected error occurred",
		},
		{
			name:           "should return invalid error if post notification return err no message",
			idempotencyKey: "test",
			setup: func(ns *mocks.NotificationService) {
				ns.EXPECT().CheckIdempotency(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("", errors.ErrNotFound)
				ns.EXPECT().Dispatch(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("[]notification.Notification")).Return(nil, notification.ErrNoMessage)
			},
			errString: "rpc error: code = InvalidArgument desc = no message sent, probably because not matching any subscription or receiver",
		},
		{
			name:           "should return success if request is idempotent",
			idempotencyKey: "test",
			setup: func(ns *mocks.NotificationService) {
				ns.EXPECT().CheckIdempotency(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(notificationID, nil)
			},
		},
		{
			name:           "should return error if idempotency checking return error",
			idempotencyKey: "test",
			setup: func(ns *mocks.NotificationService) {
				ns.EXPECT().CheckIdempotency(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("", errors.New("some error"))
			},
			errString: "rpc error: code = Internal desc = some unexpected error occurred",
		},
		{
			name:           "should return error if error inserting idempotency",
			idempotencyKey: "test",
			setup: func(ns *mocks.NotificationService) {
				ns.EXPECT().CheckIdempotency(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("", errors.ErrNotFound)
				ns.EXPECT().Dispatch(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("[]notification.Notification")).Return([]string{notificationID}, nil)
				ns.EXPECT().InsertIdempotency(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(errors.New("some error"))
			},
			errString: "rpc error: code = Internal desc = some unexpected error occurred",
		},
		{
			name:           "should return OK response if post notification succeed",
			idempotencyKey: "test",
			setup: func(ns *mocks.NotificationService) {
				ns.EXPECT().CheckIdempotency(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("", errors.ErrNotFound)
				ns.EXPECT().Dispatch(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("[]notification.Notification")).Return([]string{notificationID}, nil)
				ns.EXPECT().InsertIdempotency(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var (
				mockNotificationService = new(mocks.NotificationService)
			)

			if tc.setup != nil {
				tc.setup(mockNotificationService)
			}

			dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{
				IdempotencyKey: idempotencyHeaderKey,
			}, &api.Deps{NotificationService: mockNotificationService})
			require.NoError(t, err)

			ctx := metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{
				idempotencyHeaderKey: tc.idempotencyKey,
			}))

			// Create a dummy request with valid receivers
			req := &sirenv1beta1.PostNotificationRequest{
				Receivers: []*structpb.Struct{
					{
						Fields: map[string]*structpb.Value{
							"id": structpb.NewStringValue("123"),
						},
					},
				},
			}

			_, err = dummyGRPCServer.PostNotification(ctx, req)

			if tc.errString != "" {
				assert.EqualError(t, err, tc.errString)
			} else {
				assert.NoError(t, err)
			}

			mockNotificationService.AssertExpectations(t)
		})
	}
}

func TestGRPCServer_ListNotifications(t *testing.T) {
	ctx := context.TODO()

	t.Run("should return list of notifications with type reciever", func(t *testing.T) {
		dummyResult := []notification.Notification{
			{
				NamespaceID: 1,
				Type:        "reciever",
				Template:    "",
				Data: map[string]any{
					"data-key": "data-value",
				},
				Labels:            map[string]string{},
				ReceiverSelectors: []map[string]interface{}{},
			},
		}

		mockNotificationService := &mocks.NotificationService{}
		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{NotificationService: mockNotificationService})
		require.NoError(t, err)
		mockNotificationService.EXPECT().List(mock.AnythingOfType("context.todoCtx"), notification.Filter{Type: "reciever"}).Return(dummyResult, nil).Once()
		res, err := dummyGRPCServer.ListNotifications(ctx, &sirenv1beta1.ListNotificationsRequest{Type: "reciever"})
		assert.Nil(t, err)
		assert.Equal(t, 1, len(res.GetNotifications()))
		assert.Equal(t, "reciever", res.GetNotifications()[0].Type)
	})

	t.Run("should return error if list notifications failed", func(t *testing.T) {
		mockNotificationService := &mocks.NotificationService{}
		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{NotificationService: mockNotificationService})
		require.NoError(t, err)
		mockNotificationService.EXPECT().List(mock.AnythingOfType("context.todoCtx"), notification.Filter{Type: "alert"}).Return(nil, errors.New("internal server error")).Once()
		res, err := dummyGRPCServer.ListNotifications(ctx, &sirenv1beta1.ListNotificationsRequest{Type: "alert"})
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Internal desc = some unexpected error occurred")
	})
}