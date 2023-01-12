package notification_test

import (
	"context"
	"testing"

	"github.com/odpf/salt/log"
	"github.com/stretchr/testify/mock"

	"github.com/odpf/siren/core/idempotency"
	"github.com/odpf/siren/core/notification"
	"github.com/odpf/siren/core/notification/mocks"
	"github.com/odpf/siren/core/receiver"
	"github.com/odpf/siren/core/subscription"
	"github.com/odpf/siren/pkg/errors"
)

const testPluginType = "test"

func TestService_DispatchToReceiver(t *testing.T) {
	testCases := []struct {
		name    string
		setup   func(*mocks.ReceiverService, *mocks.Queuer, *mocks.Notifier)
		n       notification.Notification
		wantErr bool
	}{
		{
			name: "should return error if failed to transform notification to messages",
			setup: func(rs *mocks.ReceiverService, q *mocks.Queuer, _ *mocks.Notifier) {
				q.EXPECT().Type().Return("postgresql")
				rs.EXPECT().Get(mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("uint64"), mock.AnythingOfType("receiver.GetOption")).Return(&receiver.Receiver{}, nil)
			},
			n: notification.Notification{
				ValidDurationString: "xxx",
			},
			wantErr: true,
		},
		{
			name: "should return error if there is an error when fetching receiver",
			setup: func(rs *mocks.ReceiverService, q *mocks.Queuer, _ *mocks.Notifier) {
				q.EXPECT().Type().Return("postgresql")
				rs.EXPECT().Get(mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("uint64"), mock.AnythingOfType("receiver.GetOption")).Return(nil, errors.New("some error"))
			},
			wantErr: true,
		},
		{
			name: "should return error if prehook transform config return error",
			setup: func(rs *mocks.ReceiverService, q *mocks.Queuer, n *mocks.Notifier) {
				q.EXPECT().Type().Return("postgresql")
				rs.EXPECT().Get(mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("uint64"), mock.AnythingOfType("receiver.GetOption")).Return(&receiver.Receiver{
					Type: testPluginType,
					Configurations: map[string]interface{}{
						"key": "value",
					},
				}, nil)
				n.EXPECT().PreHookQueueTransformConfigs(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("map[string]interface {}")).Return(nil, errors.New("invalid config"))
			},
			wantErr: true,
		},
		{
			name: "should return error if enqueue error",
			setup: func(rs *mocks.ReceiverService, q *mocks.Queuer, n *mocks.Notifier) {
				q.EXPECT().Type().Return("postgresql")
				rs.EXPECT().Get(mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("uint64"), mock.AnythingOfType("receiver.GetOption")).Return(&receiver.Receiver{
					Type: testPluginType,
					Configurations: map[string]interface{}{
						"key": "value",
					},
				}, nil)
				n.EXPECT().PreHookQueueTransformConfigs(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("map[string]interface {}")).Return(map[string]interface{}{
					"key": "value",
				}, nil)
				q.EXPECT().Enqueue(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("notification.Message")).Return(errors.New("some error"))
			},
			wantErr: true,
		},
		{
			name: "should return no error if enqueue success",
			setup: func(rs *mocks.ReceiverService, q *mocks.Queuer, n *mocks.Notifier) {
				q.EXPECT().Type().Return("postgresql")
				rs.EXPECT().Get(mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("uint64"), mock.AnythingOfType("receiver.GetOption")).Return(&receiver.Receiver{
					Type: testPluginType,
					Configurations: map[string]interface{}{
						"key": "value",
					},
				}, nil)
				n.EXPECT().PreHookQueueTransformConfigs(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("map[string]interface {}")).Return(map[string]interface{}{
					"key": "value",
				}, nil)
				q.EXPECT().Enqueue(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("notification.Message")).Return(nil)
			},
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var (
				mockReceiverService = new(mocks.ReceiverService)
				mockQueuer          = new(mocks.Queuer)
				mockNotifier        = new(mocks.Notifier)
			)

			if tc.setup != nil {
				tc.setup(mockReceiverService, mockQueuer, mockNotifier)
			}

			ns := notification.NewService(
				log.NewNoop(), mockQueuer, nil, mockReceiverService, nil, map[string]notification.Notifier{
					testPluginType: mockNotifier,
				})

			if err := ns.DispatchToReceiver(context.Background(), tc.n, 1); (err != nil) != tc.wantErr {
				t.Errorf("NotificationService.Dispatch() error = %v, wantErr %v", err, tc.wantErr)
			}

			mockReceiverService.AssertExpectations(t)
			mockQueuer.AssertExpectations(t)
			mockNotifier.AssertExpectations(t)
		})
	}
}

func TestService_DispatchToSubscribers(t *testing.T) {
	testCases := []struct {
		name    string
		setup   func(*mocks.SubscriptionService, *mocks.Queuer, *mocks.Notifier)
		n       notification.Notification
		wantErr bool
	}{
		{
			name: "should return error if there is an error when matching labels",
			setup: func(ss *mocks.SubscriptionService, q *mocks.Queuer, _ *mocks.Notifier) {
				q.EXPECT().Type().Return("postgresql")
				ss.EXPECT().MatchByLabels(mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("uint64"), mock.AnythingOfType("map[string]string")).Return(nil, errors.New("some error"))
			},
			wantErr: true,
		},
		{
			name: "should return error if receiver type of a receiver is unknown",
			setup: func(ss *mocks.SubscriptionService, q *mocks.Queuer, _ *mocks.Notifier) {
				q.EXPECT().Type().Return("postgresql")
				ss.EXPECT().
					MatchByLabels(mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("uint64"), mock.AnythingOfType("map[string]string")).
					Return([]subscription.Subscription{
						{
							Receivers: []subscription.Receiver{
								{
									Type: "random",
								},
							},
						},
					}, nil)
			},
			wantErr: true,
		},
		{
			name: "should return error if there is no matching subscription",
			setup: func(ss *mocks.SubscriptionService, q *mocks.Queuer, _ *mocks.Notifier) {
				q.EXPECT().Type().Return("postgresql")
				ss.EXPECT().MatchByLabels(mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("uint64"), mock.AnythingOfType("map[string]string")).Return(nil, nil)
			},
			wantErr: true,
		},
		{
			name: "should return error if failed to transform notification to messages",
			setup: func(ss *mocks.SubscriptionService, q *mocks.Queuer, _ *mocks.Notifier) {
				q.EXPECT().Type().Return("postgresql")
				ss.EXPECT().
					MatchByLabels(mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("uint64"), mock.AnythingOfType("map[string]string")).
					Return([]subscription.Subscription{
						{
							Receivers: []subscription.Receiver{
								{
									Type: testPluginType,
								},
							},
						},
					}, nil)
			},
			n: notification.Notification{
				ValidDurationString: "xxx",
			},
			wantErr: true,
		},
		{
			name: "should return error if receiver config is invalid",
			setup: func(ss *mocks.SubscriptionService, q *mocks.Queuer, n *mocks.Notifier) {
				q.EXPECT().Type().Return("postgresql")
				ss.EXPECT().
					MatchByLabels(mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("uint64"), mock.AnythingOfType("map[string]string")).
					Return([]subscription.Subscription{
						{
							Receivers: []subscription.Receiver{
								{
									Type: testPluginType,
									Configuration: map[string]interface{}{
										"key": "value",
									},
								},
							},
						},
					}, nil)
				n.EXPECT().PreHookQueueTransformConfigs(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("map[string]interface {}")).Return(nil, errors.New("invalid config"))
			},
			wantErr: true,
		},
		{
			name: "should return error if enqueue error",
			setup: func(ss *mocks.SubscriptionService, q *mocks.Queuer, n *mocks.Notifier) {
				q.EXPECT().Type().Return("postgresql")
				ss.EXPECT().
					MatchByLabels(mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("uint64"), mock.AnythingOfType("map[string]string")).
					Return([]subscription.Subscription{
						{
							Receivers: []subscription.Receiver{
								{
									Type: testPluginType,
									Configuration: map[string]interface{}{
										"key": "value",
									},
								},
							},
						},
					}, nil)
				n.EXPECT().PreHookQueueTransformConfigs(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("map[string]interface {}")).Return(map[string]interface{}{
					"key": "value",
				}, nil)
				q.EXPECT().Enqueue(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("notification.Message")).Return(errors.New("some error"))
			},
			wantErr: true,
		},
		{
			name: "should return no error if enqueue success",
			setup: func(ss *mocks.SubscriptionService, q *mocks.Queuer, n *mocks.Notifier) {
				q.EXPECT().Type().Return("postgresql")
				ss.EXPECT().
					MatchByLabels(mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("uint64"), mock.AnythingOfType("map[string]string")).
					Return([]subscription.Subscription{
						{
							Receivers: []subscription.Receiver{
								{
									Type: testPluginType,
									Configuration: map[string]interface{}{
										"key": "value",
									},
								},
							},
						},
					}, nil)
				n.EXPECT().PreHookQueueTransformConfigs(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("map[string]interface {}")).Return(map[string]interface{}{
					"key": "value",
				}, nil)
				q.EXPECT().Enqueue(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("notification.Message")).Return(nil)
			},
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var (
				mockSubscriptionService = new(mocks.SubscriptionService)
				mockQueuer              = new(mocks.Queuer)
				mockNotifier            = new(mocks.Notifier)
			)

			if tc.setup != nil {
				tc.setup(mockSubscriptionService, mockQueuer, mockNotifier)
			}

			ns := notification.NewService(
				log.NewNoop(), mockQueuer, nil, nil, mockSubscriptionService, map[string]notification.Notifier{
					testPluginType: mockNotifier,
				})

			if err := ns.DispatchToSubscribers(context.Background(), 1, tc.n); (err != nil) != tc.wantErr {
				t.Errorf("NotificationService.Dispatch() error = %v, wantErr %v", err, tc.wantErr)
			}

			mockSubscriptionService.AssertExpectations(t)
			mockQueuer.AssertExpectations(t)
			mockNotifier.AssertExpectations(t)
		})
	}
}

func TestService_CheckAndInsertIdempotency(t *testing.T) {
	var (
		scope = "test-scope"
		key   = "test-key"
	)
	testCases := []struct {
		name    string
		setup   func(*mocks.IdempotencyRepository)
		scope   string
		key     string
		wantErr bool
	}{
		{
			name: "should return error if idempotency exist and success",
			setup: func(ir *mocks.IdempotencyRepository) {
				ir.EXPECT().InsertOnConflictReturning(mock.AnythingOfType("*context.emptyCtx"), scope, key).Return(nil, errors.ErrConflict)
			},
			scope:   scope,
			key:     key,
			wantErr: true,
		},
		{
			name: "should return error if repository returning some error",
			setup: func(ir *mocks.IdempotencyRepository) {
				ir.EXPECT().InsertOnConflictReturning(mock.AnythingOfType("*context.emptyCtx"), scope, key).Return(nil, errors.New("some error"))
			},
			scope:   scope,
			key:     key,
			wantErr: true,
		},
		{
			name: "should return id and nil error if no idempotency exists",
			setup: func(ir *mocks.IdempotencyRepository) {
				ir.EXPECT().InsertOnConflictReturning(mock.AnythingOfType("*context.emptyCtx"), scope, key).Return(&idempotency.Idempotency{
					ID: 1,
				}, nil)
			},
			scope:   scope,
			key:     key,
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockIdempotencyRepository := new(mocks.IdempotencyRepository)

			if tc.setup != nil {
				tc.setup(mockIdempotencyRepository)
			}

			ns := notification.NewService(
				log.NewNoop(), nil, mockIdempotencyRepository, nil, nil, nil)

			_, err := ns.CheckAndInsertIdempotency(context.Background(), tc.scope, tc.key)

			if (err != nil) != tc.wantErr {
				t.Errorf("NotificationService.CheckAndInsertIdempotency() error = %v, wantErr %v", err, tc.wantErr)
			}

			mockIdempotencyRepository.AssertExpectations(t)
		})
	}
}