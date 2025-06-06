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
	"github.com/stretchr/testify/mock"
)

func TestRouterReceiverService_PrepareMetaMessage(t *testing.T) {
	tests := []struct {
		name       string
		setup      func(*mocks.ReceiverService, *mocks.Notifier)
		n          notification.Notification
		want       []notification.MetaMessage
		want1      []log.Notification
		wantErrStr string
	}{
		{
			name: "should return error if number of receiver selector is more than threshold",
			n: notification.Notification{
				ReceiverSelectors: []map[string]any{
					{
						"k1": "v1",
					},
					{
						"k2": "v2",
					},
					{
						"k3": "v3",
					},
				},
			},
			wantErrStr: "number of receiver selectors should be less than or equal threshold 2",
		},
		{
			name: "should return error if receiver selector value is not string",
			n: notification.Notification{
				ReceiverSelectors: []map[string]any{
					{
						"k1": map[string]any{},
					},
					{
						"k2": "v2",
					},
				},
			},
			wantErrStr: "receiver selector value of 'k1' should be a string",
		},
		{
			name: "should return error if receiver service return error",
			n:    notification.Notification{},
			setup: func(rs *mocks.ReceiverService, n *mocks.Notifier) {
				rs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("receiver.Filter")).Return(nil, errors.New("some error"))
			},
			wantErrStr: "some error",
		},
		{
			name: "should return error if receiver not found",
			n:    notification.Notification{},
			setup: func(rs *mocks.ReceiverService, n *mocks.Notifier) {
				rs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("receiver.Filter")).Return(nil, nil)
			},
			wantErrStr: "requested entity not found",
		},
		{
			name: "should return error if meta message result more than max messages receiver flow",
			n:    notification.Notification{},
			setup: func(rs *mocks.ReceiverService, n *mocks.Notifier) {
				rs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("receiver.Filter")).Return([]receiver.Receiver{
					{
						ID: 1,
					},
					{
						ID: 2,
					},
					{
						ID: 3,
					},
				}, nil)
			},
			wantErrStr: "sending 3 messages exceed max messages receiver flow threshold 2. this will spam and broadcast to 3 channel. found 0 receiver selectors passed, you might want to check your receiver selectors configuration",
		},
		// Config Override
		{
			name: "should return error if receiver selectors is more than one but there is a config override",
			n: notification.Notification{
				ReceiverSelectors: []map[string]any{
					{
						"k1": "v1",
						"config": map[string]any{
							"k3": "v3",
						},
					},
					{
						"k2": "v2",
					},
				},
			},
			setup: func(rs *mocks.ReceiverService, n *mocks.Notifier) {
				rs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("receiver.Filter")).Return([]receiver.Receiver{
					{
						ID: 1,
					},
					{
						ID: 2,
					},
					{
						ID: 3,
					},
				}, nil)
			},
			wantErrStr: "config override could only be used with one selector",
		},
		{
			name: "should return error if config override i receiver selectors is not a map",
			n: notification.Notification{
				ReceiverSelectors: []map[string]any{
					{
						"k1":     "v1",
						"config": 123,
					},
					{
						"k2": "v2",
					},
				},
			},
			setup: func(rs *mocks.ReceiverService, n *mocks.Notifier) {
				rs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("receiver.Filter")).Return([]receiver.Receiver{
					{
						ID: 1,
					},
					{
						ID: 2,
					},
					{
						ID: 3,
					},
				}, nil)
			},
			wantErrStr: "config should be in map and follow notification config",
		},
		{
			name: "should return error if config override is being used to more than 1 evaluated receivers",
			n: notification.Notification{
				ReceiverSelectors: []map[string]any{
					{
						"k1": "v1",
						"config": map[string]any{
							"k3": "v3",
						},
					},
				},
			},
			setup: func(rs *mocks.ReceiverService, n *mocks.Notifier) {
				rs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("receiver.Filter")).Return([]receiver.Receiver{
					{
						ID: 1,
					},
					{
						ID: 2,
					},
					{
						ID: 3,
					},
				}, nil)
			},
			wantErrStr: "config override could only be used to 1 receiver, but got 3 receiver",
		},
		{
			name: "should return no error if succeed",
			n: notification.Notification{
				ID:          "test-notification-id",
				NamespaceID: 123,
			},
			setup: func(rs *mocks.ReceiverService, n *mocks.Notifier) {
				rs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("receiver.Filter")).Return([]receiver.Receiver{
					{
						ID: 1,
						Configurations: map[string]interface{}{
							"token":        "token1",
							"workspace":    "workspace1",
							"channel_name": "channel1",
						},
					},
					{
						ID: 2,
						Configurations: map[string]interface{}{
							"token":        "token2",
							"workspace":    "workspace2",
							"channel_name": "channel2",
						},
					},
				}, nil)
			},
			want: []notification.MetaMessage{
				{
					ReceiverID:      1,
					NotificationIDs: []string{"test-notification-id"},
					ReceiverConfigs: map[string]interface{}{
						"token":        "token1",
						"workspace":    "workspace1",
						"channel_name": "channel1",
					},
				},
				{
					ReceiverID:      2,
					NotificationIDs: []string{"test-notification-id"},
					ReceiverConfigs: map[string]interface{}{
						"token":        "token2",
						"workspace":    "workspace2",
						"channel_name": "channel2",
					},
				},
			},
			want1: []log.Notification{
				{
					ReceiverID:     1,
					NotificationID: "test-notification-id",
					NamespaceID:    123,
				},
				{
					ReceiverID:     2,
					NotificationID: "test-notification-id",
					NamespaceID:    123,
				},
			},
		},
		{
			name: "should return no error if succeed with config override feature",
			n: notification.Notification{
				ID: "test-notification-id",
				ReceiverSelectors: []map[string]any{
					{
						"k1": "v1",
						"config": map[string]any{
							"channel_name": "channel1",
							"channel_type": "xyz",
						},
					},
				},
				NamespaceID: 123,
			},
			setup: func(rs *mocks.ReceiverService, n *mocks.Notifier) {
				rs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("receiver.Filter")).Return([]receiver.Receiver{
					{
						ID: 1,
						Configurations: map[string]interface{}{
							"token":     "token1",
							"workspace": "workspace1",
						},
					},
				}, nil)
			},
			want: []notification.MetaMessage{
				{
					ReceiverID:      1,
					NotificationIDs: []string{"test-notification-id"},
					ReceiverConfigs: map[string]interface{}{
						"token":        "token1",
						"workspace":    "workspace1",
						"channel_name": "channel1",
						"channel_type": "xyz",
					},
				},
			},
			want1: []log.Notification{
				{
					ReceiverID:     1,
					NotificationID: "test-notification-id",
					NamespaceID:    123,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				mockReceiverService = new(mocks.ReceiverService)
				mockNotifier        = new(mocks.Notifier)
				mockTemplateService = new(mocks.TemplateService)
			)
			s := notification.NewRouterReceiverService(
				notification.Deps{
					Cfg: notification.Config{
						MaxNumReceiverSelectors: 2,
						MaxMessagesReceiverFlow: 2,
					},
					Logger:          saltlog.NewNoop(),
					ReceiverService: mockReceiverService,
					TemplateService: mockTemplateService,
				},
			)

			if tt.setup != nil {
				tt.setup(mockReceiverService, mockNotifier)
			}

			got, got1, err := s.PrepareMetaMessages(context.TODO(), tt.n)
			if err != nil {
				if err.Error() != tt.wantErrStr {
					t.Errorf("RouterSubscriberService.PrepareMetaMessage() error = %v, wantErr %v", err, tt.wantErrStr)
					return
				}
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("RouterSubscriberService.PrepareMetaMessage() diff = %v", diff)
			}
			if diff := cmp.Diff(got1, tt.want1); diff != "" {
				t.Errorf("RouterSubscriberService.PrepareMetaMessage() diff = %v", diff)
			}
		})
	}
}
