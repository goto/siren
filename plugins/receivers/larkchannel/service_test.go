package larkchannel_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/goto/salt/log"
	"github.com/goto/siren/core/notification"
	"github.com/goto/siren/pkg/errors"
	"github.com/goto/siren/pkg/retry"
	"github.com/goto/siren/pkg/secret"
	"github.com/goto/siren/plugins/receivers/lark"
	"github.com/goto/siren/plugins/receivers/lark/mocks"
	"github.com/goto/siren/plugins/receivers/larkchannel"
	"github.com/stretchr/testify/mock"
)

func TestService_Send(t *testing.T) {
	tests := []struct {
		name                string
		setup               func(*mocks.LarkCaller)
		notificationMessage notification.Message
		wantRetryable       bool
		wantErr             bool
	}{
		{
			name: "should return error if failed to decode notification config",
			notificationMessage: notification.Message{
				Configs: map[string]any{
					"client_id": true,
				},
				Details: map[string]any{},
			},
			wantErr: true,
		},
		{
			name: "should return error and not retryable if notify return error",
			setup: func(sc *mocks.LarkCaller) {
				sc.EXPECT().Notify(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("lark.NotificationConfig"), mock.AnythingOfType("lark.Message")).Return(errors.New("some error"))
			},
			notificationMessage: notification.Message{
				Configs: map[string]any{
					"client_id": true,
				},
				Details: map[string]any{
					"message": "hello",
				},
			},
			wantRetryable: false,
			wantErr:       true,
		},
		{
			name: "should return error and retryable if notify return retryable error",
			setup: func(sc *mocks.LarkCaller) {
				sc.EXPECT().Notify(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("lark.NotificationConfig"), mock.AnythingOfType("lark.Message")).Return(retry.RetryableError{Err: errors.New("some error")})
			},
			notificationMessage: notification.Message{
				Configs: map[string]any{
					"client_id": "123123",
				},
				Details: map[string]any{
					"message": "hello",
				},
			},
			wantRetryable: true,
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLarkClient := new(mocks.LarkCaller)

			if tt.setup != nil {
				tt.setup(mockLarkClient)
			}

			s := larkchannel.NewPluginService(lark.AppConfig{}, log.NewNoop(), nil, lark.WithLarkClient(mockLarkClient))

			got, err := s.Send(context.TODO(), tt.notificationMessage)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.Publish() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.wantRetryable {
				t.Errorf("Service.Publish() = %v, want %v", got, tt.wantRetryable)
			}
		})
	}
}

func TestService_PostHookDBTransformConfigs(t *testing.T) {
	tests := []struct {
		name                  string
		setup                 func(*mocks.Encryptor)
		notificationConfigMap map[string]any
		want                  map[string]any
		wantErr               bool
	}{
		{
			name: "should return error if failed to parse configmap to notification config",
			notificationConfigMap: map[string]any{
				"client_id": 123,
			},
			wantErr: true,
		},
		{
			name: "should return error if validate notification config failed",
			notificationConfigMap: map[string]any{
				"client_id": 123,
			},
			wantErr: true,
		},
		{
			name: "should return error if lark token decryption failed",
			notificationConfigMap: map[string]any{
				"client_secret": 123,
			},
			setup: func(e *mocks.Encryptor) {
				e.EXPECT().Decrypt(mock.AnythingOfType("secret.MaskableString")).Return("", errors.New("some error"))
			},
			wantErr: true,
		},
		{
			name: "should return decrypted lark token if succeed",

			notificationConfigMap: map[string]any{
				"client_id":     secret.MaskableString("a token"),
				"client_secret": secret.MaskableString("a token"),
				"channel_name":  "channel",
				"channel_type":  "",
			},
			setup: func(e *mocks.Encryptor) {
				e.EXPECT().Decrypt(mock.AnythingOfType("secret.MaskableString")).Return(secret.MaskableString("maskable-token"), nil)
			},
			want: map[string]any{
				"client_id":     secret.MaskableString("maskable-token"),
				"client_secret": secret.MaskableString("maskable-token"),
				"channel_name":  "channel",
				"channel_type":  "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				mockEncryptor = new(mocks.Encryptor)
			)

			if tt.setup != nil {
				tt.setup(mockEncryptor)
			}

			s := larkchannel.NewPluginService(lark.AppConfig{}, log.NewNoop(), mockEncryptor)
			got, err := s.PostHookDBTransformConfigs(context.TODO(), tt.notificationConfigMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.PostHookDBTransformConfigs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.PostHookDBTransformConfigs() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPluginService_PreHookDBTransformConfigs(t *testing.T) {
	tests := []struct {
		name           string
		setup          func(*mocks.Encryptor)
		configurations map[string]any
		wantErr        bool
	}{
		{
			name:    "should return error if channel_name is missing",
			wantErr: true,
		},
		{
			name: "shouldd return non error if channel_name is not missing",
			configurations: map[string]any{
				"channel_name": "a-channel",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &larkchannel.PluginService{}
			_, err := s.PreHookDBTransformConfigs(context.TODO(), tt.configurations)
			if (err != nil) != tt.wantErr {
				t.Errorf("PluginService.PreHookDBTransformConfigs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
