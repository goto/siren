package notification_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	saltlog "github.com/goto/salt/log"
	"github.com/goto/siren/core/log"
	"github.com/goto/siren/core/notification"
	"github.com/goto/siren/core/notification/mocks"
	"github.com/goto/siren/core/receiver"
	"github.com/goto/siren/core/template"
	"github.com/stretchr/testify/mock"
)

const (
	testType = "test"
)

func TestService_DispatchFailure(t *testing.T) {
	tests := []struct {
		name       string
		n          []notification.Notification
		setup      func([]notification.Notification, *mocks.Repository, *mocks.AlertRepository, *mocks.LogService, *mocks.Queuer, *mocks.Router, *mocks.Notifier)
		wantErrStr string
	}{
		{
			name: "should return error if bulk create return error",
			n: []notification.Notification{
				{
					Type: notification.TypeAlert,
					Labels: map[string]string{
						"k1": "v1",
					},
				},
			},
			setup: func(n []notification.Notification, r *mocks.Repository, _ *mocks.AlertRepository, _ *mocks.LogService, _ *mocks.Queuer, _ *mocks.Router, _ *mocks.Notifier) {
				r.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*errors.errorString")).Return(nil)
				r.EXPECT().BulkCreate(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("[]notification.Notification")).Return(nil, errors.New("some error"))
			},
			wantErrStr: "some error",
		},
		{
			name: "should return error if receiver selectors and labels are empty",
			n: []notification.Notification{
				{
					Type: notification.TypeAlert,
				},
			},
			setup: func(n []notification.Notification, r *mocks.Repository, _ *mocks.AlertRepository, _ *mocks.LogService, _ *mocks.Queuer, _ *mocks.Router, _ *mocks.Notifier) {
				r.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("errors.Error")).Return(nil)
				r.EXPECT().BulkCreate(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("[]notification.Notification")).Return([]notification.Notification{
					{
						Type: notification.TypeAlert,
					},
				}, nil)
			},
			wantErrStr: "no receivers or labels found, unknown flow",
		},
		{
			name: "should return error if router prepare meta messages return generic error",
			n: []notification.Notification{
				{
					Type: notification.TypeAlert,
					Labels: map[string]string{
						"k1": "v1",
					},
				},
			},
			setup: func(n []notification.Notification, r *mocks.Repository, _ *mocks.AlertRepository, _ *mocks.LogService, _ *mocks.Queuer, ro *mocks.Router, _ *mocks.Notifier) {
				r.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*errors.errorString")).Return(nil)
				r.EXPECT().BulkCreate(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("[]notification.Notification")).Return([]notification.Notification{
					{
						Type: notification.TypeAlert,
						Labels: map[string]string{
							"k1": "v1",
						},
					},
				}, nil)
				ro.EXPECT().PrepareMetaMessages(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return(nil, nil, errors.New("some error"))
			},
			wantErrStr: "some error",
		},
		{
			name: "should pass if router prepare meta messages return ErrRouteSubscriberNoMatchFound and return ErrNoMessages",
			n: []notification.Notification{
				{
					Type: notification.TypeAlert,
					Labels: map[string]string{
						"k1": "v1",
					},
				},
			},
			setup: func(n []notification.Notification, r *mocks.Repository, _ *mocks.AlertRepository, _ *mocks.LogService, _ *mocks.Queuer, ro *mocks.Router, _ *mocks.Notifier) {
				r.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*errors.errorString")).Return(nil)
				r.EXPECT().BulkCreate(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("[]notification.Notification")).Return([]notification.Notification{
					{
						Type: notification.TypeAlert,
						Labels: map[string]string{
							"k1": "v1",
						},
					},
				}, nil)
				ro.EXPECT().PrepareMetaMessages(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return(nil, nil, notification.ErrRouteSubscriberNoMatchFound)
			},
			wantErrStr: notification.ErrNoMessage.Error(),
		},
		{
			name: "should return error if prepare messages return error",
			n: []notification.Notification{
				{
					Type: notification.TypeAlert,
					Labels: map[string]string{
						"k1": "v1",
					},
				},
			},
			setup: func(n []notification.Notification, r *mocks.Repository, _ *mocks.AlertRepository, _ *mocks.LogService, _ *mocks.Queuer, ro *mocks.Router, no *mocks.Notifier) {
				r.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*errors.errorString")).Return(nil)
				r.EXPECT().BulkCreate(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("[]notification.Notification")).Return([]notification.Notification{
					{
						Type: notification.TypeAlert,
						Labels: map[string]string{
							"k1": "v1",
						},
					},
				}, nil)
				ro.EXPECT().PrepareMetaMessages(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return([]notification.MetaMessage{
					{
						ReceiverID:      1,
						SubscriptionIDs: []uint64{1, 2},
						ReceiverType:    receiver.TypeSlack,
					},
				}, []log.Notification{}, nil)
				no.EXPECT().PreHookQueueTransformConfigs(mock.AnythingOfType("context.todoCtx"), mock.Anything).Return(nil, errors.New("some error"))

			},
			wantErrStr: "some error",
		},
		{
			name: "should return error if template not found",
			n: []notification.Notification{
				{
					Type: notification.TypeAlert,
					Labels: map[string]string{
						"k1": "v1",
					},
				},
			},
			setup: func(n []notification.Notification, r *mocks.Repository, _ *mocks.AlertRepository, l *mocks.LogService, _ *mocks.Queuer, ro *mocks.Router, no *mocks.Notifier) {
				r.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("errors.Error")).Return(nil)
				r.EXPECT().BulkCreate(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("[]notification.Notification")).Return([]notification.Notification{
					{
						Type: notification.TypeAlert,
						Labels: map[string]string{
							"k1": "v1",
						},
					},
				}, nil)
				ro.EXPECT().PrepareMetaMessages(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return([]notification.MetaMessage{
					{
						ReceiverID:      1,
						SubscriptionIDs: []uint64{1, 2},
						ReceiverType:    receiver.TypeSlack,
					},
				}, []log.Notification{
					{
						ID: "1",
					},
				}, nil)
				no.EXPECT().PreHookQueueTransformConfigs(mock.AnythingOfType("context.todoCtx"), mock.Anything).Return(map[string]any{}, nil)
				no.EXPECT().GetSystemDefaultTemplate().Return("system-default")
			},
			wantErrStr: "found no template, template is mandatory",
		},
		{
			name: "should return error if log notifications return error",
			n: []notification.Notification{
				{
					Type: notification.TypeAlert,
					Labels: map[string]string{
						"k1": "v1",
					},
				},
			},
			setup: func(n []notification.Notification, r *mocks.Repository, _ *mocks.AlertRepository, l *mocks.LogService, _ *mocks.Queuer, ro *mocks.Router, no *mocks.Notifier) {
				r.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*fmt.wrapError")).Return(nil)
				r.EXPECT().BulkCreate(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("[]notification.Notification")).Return([]notification.Notification{
					{
						Type: notification.TypeAlert,
						Labels: map[string]string{
							"k1": "v1",
						},
					},
				}, nil)
				ro.EXPECT().PrepareMetaMessages(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return([]notification.MetaMessage{
					{
						ReceiverID:      1,
						SubscriptionIDs: []uint64{1, 2},
						ReceiverType:    receiver.TypeSlack,
						Template:        template.ReservedName_SystemDefault,
					},
				}, []log.Notification{
					{
						ID: "1",
					},
				}, nil)
				no.EXPECT().GetSystemDefaultTemplate().Return("")
				no.EXPECT().PreHookQueueTransformConfigs(mock.AnythingOfType("context.todoCtx"), mock.Anything).Return(map[string]any{}, nil)
				l.EXPECT().LogNotifications(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("log.Notification")).Return(errors.New("some error"))
			},
			wantErrStr: "failed logging notifications: some error",
		},
		{
			name: "should return error if enqueue return error",
			n: []notification.Notification{
				{
					Type: notification.TypeAlert,
					Labels: map[string]string{
						"k1": "v1",
					},
				},
			},
			setup: func(n []notification.Notification, r *mocks.Repository, _ *mocks.AlertRepository, l *mocks.LogService, q *mocks.Queuer, ro *mocks.Router, no *mocks.Notifier) {
				r.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*fmt.wrapError")).Return(nil)
				r.EXPECT().BulkCreate(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("[]notification.Notification")).Return([]notification.Notification{
					{
						Type: notification.TypeAlert,
						Labels: map[string]string{
							"k1": "v1",
						},
					},
				}, nil)
				ro.EXPECT().PrepareMetaMessages(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return([]notification.MetaMessage{
					{
						ReceiverID:      1,
						SubscriptionIDs: []uint64{1, 2},
						ReceiverType:    receiver.TypeSlack,
						Template:        template.ReservedName_SystemDefault,
					},
				}, []log.Notification{}, nil)
				no.EXPECT().GetSystemDefaultTemplate().Return("")
				no.EXPECT().PreHookQueueTransformConfigs(mock.AnythingOfType("context.todoCtx"), mock.Anything).Return(map[string]any{}, nil)
				l.EXPECT().LogNotifications(mock.AnythingOfType("context.todoCtx")).Return(nil)
				q.EXPECT().Enqueue(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Message")).Return(errors.New("some error"))
			},
			wantErrStr: "failed enqueuing messages: some error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				mockQueuer          = new(mocks.Queuer)
				mockRepository      = new(mocks.Repository)
				mockAlertRepository = new(mocks.AlertRepository)
				mockLogService      = new(mocks.LogService)
				mockRouter          = new(mocks.Router)
				mockNotifier        = new(mocks.Notifier)
			)

			if tt.setup != nil {
				mockRepository.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
				tt.setup(tt.n, mockRepository, mockAlertRepository, mockLogService, mockQueuer, mockRouter, mockNotifier)
			}

			s := notification.NewService(
				notification.Deps{
					Logger:          saltlog.NewNoop(),
					Repository:      mockRepository,
					Q:               mockQueuer,
					AlertRepository: mockAlertRepository,
					LogService:      mockLogService,
				},
				map[string]notification.Router{
					testType:                      mockRouter,
					notification.RouterSubscriber: mockRouter,
				},
				map[string]notification.Notifier{
					receiver.TypeSlack: mockNotifier,
				},
			)
			if _, err := s.Dispatch(context.TODO(), tt.n); err != nil {
				if err.Error() != tt.wantErrStr {
					t.Errorf("Service.DispatchFailure() error = %v, wantErr %v", err, tt.wantErrStr)
				}
			} else {
				t.Errorf("Service.DispatchFailure() error = %v, wantErr %v", err, tt.wantErrStr)
			}
		})
	}
}

func TestService_DispatchSuccess(t *testing.T) {
	tests := []struct {
		name       string
		n          []notification.Notification
		setup      func([]notification.Notification, *mocks.Repository, *mocks.AlertRepository, *mocks.LogService, *mocks.Queuer, *mocks.Router, *mocks.Notifier)
		wantResult []string
	}{
		{
			name: "should succesfully enqueue messages",
			n: []notification.Notification{
				{
					ID:   "1",
					Type: notification.TypeAlert,
					Labels: map[string]string{
						"k1": "v1",
					},
				},
				{
					ID:   "2",
					Type: notification.TypeAlert,
					Labels: map[string]string{
						"k2": "v2",
					},
				},
			},
			setup: func(n []notification.Notification, r *mocks.Repository, _ *mocks.AlertRepository, l *mocks.LogService, q *mocks.Queuer, ro *mocks.Router, no *mocks.Notifier) {
				r.EXPECT().Commit(mock.AnythingOfType("context.todoCtx")).Return(nil)
				r.EXPECT().BulkCreate(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("[]notification.Notification")).Return([]notification.Notification{
					{
						ID:   "1",
						Type: notification.TypeAlert,
						Labels: map[string]string{
							"k1": "v1",
						},
					},
					{
						ID:   "2",
						Type: notification.TypeAlert,
						Labels: map[string]string{
							"k2": "v2",
						},
					},
				}, nil)
				ro.EXPECT().PrepareMetaMessages(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return([]notification.MetaMessage{
					{
						ReceiverID:      1,
						SubscriptionIDs: []uint64{1, 2},
						ReceiverType:    receiver.TypeSlack,
						Template:        template.ReservedName_SystemDefault,
					}, {
						ReceiverID:      2,
						SubscriptionIDs: []uint64{1, 2},
						ReceiverType:    receiver.TypeSlack,
						Template:        template.ReservedName_SystemDefault,
					},
				}, []log.Notification{}, nil)
				no.EXPECT().GetSystemDefaultTemplate().Return("")
				no.EXPECT().PreHookQueueTransformConfigs(mock.AnythingOfType("context.todoCtx"), mock.Anything).Return(map[string]any{}, nil)
				l.EXPECT().LogNotifications(mock.AnythingOfType("context.todoCtx")).Return(nil)
				q.EXPECT().Enqueue(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Message"), mock.AnythingOfType("notification.Message")).Return(nil)
			},
			wantResult: []string{"1", "2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				mockQueuer          = new(mocks.Queuer)
				mockRepository      = new(mocks.Repository)
				mockAlertRepository = new(mocks.AlertRepository)
				mockLogService      = new(mocks.LogService)
				mockRouter          = new(mocks.Router)
				mockNotifier        = new(mocks.Notifier)
			)

			if tt.setup != nil {
				mockRepository.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
				tt.setup(tt.n, mockRepository, mockAlertRepository, mockLogService, mockQueuer, mockRouter, mockNotifier)
			}

			s := notification.NewService(
				notification.Deps{
					Logger:          saltlog.NewNoop(),
					Repository:      mockRepository,
					Q:               mockQueuer,
					AlertRepository: mockAlertRepository,
					LogService:      mockLogService,
				},
				map[string]notification.Router{
					testType:                      mockRouter,
					notification.RouterSubscriber: mockRouter,
				},
				map[string]notification.Notifier{
					receiver.TypeSlack: mockNotifier,
				},
			)
			if ids, err := s.Dispatch(context.TODO(), tt.n); err != nil {
				t.Errorf("Service.DispatchSuccess() error = %v", err)
			} else {
				if diff := cmp.Diff(tt.wantResult, ids); diff != "" {
					t.Errorf("Service.DispatchSuccess() diff = %v", diff)
				}
			}
		})
	}
}
