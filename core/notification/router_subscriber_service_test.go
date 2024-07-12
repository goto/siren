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
	"github.com/goto/siren/core/subscription"
	"github.com/stretchr/testify/mock"
)

func TestRouterSubscriberService_PrepareMetaMessage(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(*mocks.SubscriptionService, *mocks.Notifier)
		n       notification.Notification
		want    []notification.MetaMessage
		want1   []log.Notification
		wantErr bool
	}{
		{
			name: "should return error if subscription service match by labels return error",
			setup: func(ss1 *mocks.SubscriptionService, n *mocks.Notifier) {
				ss1.EXPECT().MatchByLabelsV2(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64"), mock.AnythingOfType("map[string]string")).Return(nil, errors.New("some error"))
			},
			wantErr: true,
		},
		{
			name: "should return error if no matching subscriptions",
			setup: func(ss1 *mocks.SubscriptionService, n *mocks.Notifier) {
				ss1.EXPECT().MatchByLabelsV2(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64"), mock.AnythingOfType("map[string]string")).Return(nil, nil)
			},
			wantErr: true,
		},
		{
			name: "should return metamessages and notification logs if succeed",
			setup: func(ss1 *mocks.SubscriptionService, n *mocks.Notifier) {
				ss1.EXPECT().MatchByLabelsV2(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64"), mock.AnythingOfType("map[string]string")).Return([]subscription.ReceiverView{
					{

						ID:   1,
						Name: "rcv 1",
						Labels: map[string]string{
							"k1": "k2",
						},
						Type:           receiver.TypeSlackChannel,
						SubscriptionID: 123,
					},
				}, nil)
			},
			want: []notification.MetaMessage{
				{
					ReceiverID:      1,
					SubscriptionIDs: []uint64{123},
					ReceiverType:    "slack_channel",
				},
			},
			want1: []log.Notification{
				{
					SubscriptionID: 123,
					ReceiverID:     1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				mockSubscriptionService = new(mocks.SubscriptionService)
				mockNotifier            = new(mocks.Notifier)
				mockTemplateService     = new(mocks.TemplateService)
			)
			s := notification.NewRouterSubscriberService(
				notification.Deps{
					Logger:              saltlog.NewNoop(),
					SubscriptionService: mockSubscriptionService,
					TemplateService:     mockTemplateService,
				},
			)

			if tt.setup != nil {
				tt.setup(mockSubscriptionService, mockNotifier)
			}

			got, got1, err := s.PrepareMetaMessages(context.TODO(), tt.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("RouterSubscriberService.PrepareMetaMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
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
