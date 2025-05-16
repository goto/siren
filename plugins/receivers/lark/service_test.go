package lark_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/goto/salt/log"
	"github.com/goto/siren/core/notification"
	"github.com/goto/siren/pkg/errors"
	"github.com/goto/siren/pkg/retry"
	"github.com/goto/siren/pkg/secret"
	"github.com/goto/siren/plugins/receivers/lark"
	"github.com/goto/siren/plugins/receivers/lark/mocks"
	"github.com/stretchr/testify/mock"
)

func TestService_BuildData(t *testing.T) {
	type testCase struct {
		Description string
		Setup       func(sc *mocks.LarkCaller, e *mocks.Encryptor)
		Confs       map[string]any
		Err         error
	}
	var (
		ctx       = context.TODO()
		testCases = []testCase{
			{
				Description: "should return error if configuration is invalid",
				Setup:       func(sc *mocks.LarkCaller, e *mocks.Encryptor) {},
				Confs:       make(map[string]any),
				Err:         errors.New("invalid lark receiver config, workspace: , token: "),
			},
			{
				Description: "should return error if failed to get workspace channels with lark client",
				Setup: func(sc *mocks.LarkCaller, e *mocks.Encryptor) {
					sc.EXPECT().GetWorkspaceChannels(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("secret.MaskableString"), mock.AnythingOfType("secret.MaskableString")).Return(nil, errors.New("some error"))
				},
				Confs: map[string]any{
					"client_id":     secret.MaskableString("key"),
					"client_secret": secret.MaskableString("key"),
				},
				Err: errors.New("could not get channels: some error"),
			},
			{
				Description: "should return nil error if success populating receiver.Receiver",
				Setup: func(sc *mocks.LarkCaller, e *mocks.Encryptor) {
					sc.EXPECT().GetWorkspaceChannels(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("secret.MaskableString"), mock.AnythingOfType("secret.MaskableString")).Return([]lark.Channel{
						{
							ID:   "id",
							Name: "name",
						},
					}, nil)
				},
				Confs: map[string]any{
					"client_id":     secret.MaskableString("key"),
					"client_secret": secret.MaskableString("key"),
				},
			},
		}
	)

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			var (
				larkClientMock = new(mocks.LarkCaller)
				encryptorMock  = new(mocks.Encryptor)
			)

			svc := lark.NewPluginService(lark.AppConfig{}, log.NewNoop(), encryptorMock, lark.WithLarkClient(larkClientMock))

			tc.Setup(larkClientMock, encryptorMock)

			_, err := svc.BuildData(ctx, tc.Confs)
			if tc.Err != err {
				if tc.Err.Error() != err.Error() {
					t.Fatalf("got error %s, expected was %s", err.Error(), tc.Err.Error())
				}
			}

			larkClientMock.AssertExpectations(t)
			encryptorMock.AssertExpectations(t)
		})
	}
}

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
			name: "should return error if failed to decode notification detail",
			notificationMessage: notification.Message{
				Details: map[string]any{
					"text": make(chan bool),
				},
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

			s := lark.NewPluginService(lark.AppConfig{}, log.NewNoop(), nil, lark.WithLarkClient(mockLarkClient))

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

func TestService_PreHookQueueTransformConfigs(t *testing.T) {
	tests := []struct {
		name                  string
		setup                 func(*mocks.Encryptor)
		notificationConfigMap map[string]any
		want                  map[string]any
		wantErr               bool
	}{
		{
			name:                  "should return error if failed to parse configmap to notification config",
			notificationConfigMap: nil,
			wantErr:               true,
		},
		{
			name: "should return error if validate notification config failed",
			notificationConfigMap: map[string]any{
				"client_id": 123,
			},
			wantErr: true,
		},
		{
			name: "should return error if lark cred encryption failed",
			notificationConfigMap: map[string]any{
				"client_id": secret.MaskableString("a token"),
			},
			setup: func(e *mocks.Encryptor) {
				e.EXPECT().Encrypt(mock.AnythingOfType("secret.MaskableString")).Return("", errors.New("some error"))
			},
			wantErr: true,
		},
		{
			name: "should return encrypted lark cred if succeed",

			notificationConfigMap: map[string]any{
				"client_secret": secret.MaskableString("a token"),
				"client_id":     secret.MaskableString("a token"),
				"channel_name":  "channel",
				"channel_type":  "",
			},
			setup: func(e *mocks.Encryptor) {
				e.EXPECT().Encrypt(mock.AnythingOfType("secret.MaskableString")).Return(secret.MaskableString("maskable-token"), nil)
			},
			want: map[string]any{
				"client_secret":  secret.MaskableString("maskable-token"),
				"client_id":      secret.MaskableString("maskable-token"),
				"valid_duration": time.Duration(0),
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

			s := lark.NewPluginService(lark.AppConfig{}, log.NewNoop(), mockEncryptor)
			got, err := s.PreHookQueueTransformConfigs(context.TODO(), tt.notificationConfigMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.PreHookQueueTransformConfigs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.PreHookQueueTransformConfigs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_PostHookQueueTransformConfigs(t *testing.T) {
	tests := []struct {
		name                  string
		setup                 func(*mocks.Encryptor)
		notificationConfigMap map[string]any
		want                  map[string]any
		wantErr               bool
	}{
		{
			name:                  "should return error if failed to parse configmap to notification config",
			notificationConfigMap: nil,
			wantErr:               true,
		},
		{
			name: "should return error if validate notification config failed",
			notificationConfigMap: map[string]any{
				"client_id": 123,
			},
			wantErr: true,
		},
		{
			name: "should return error if lark cred decryption failed",
			notificationConfigMap: map[string]any{
				"client_id": secret.MaskableString("a token"),
			},
			setup: func(e *mocks.Encryptor) {
				e.EXPECT().Decrypt(mock.AnythingOfType("secret.MaskableString")).Return("", errors.New("some error"))
			},
			wantErr: true,
		},
		{
			name: "should return decrypted lark cred if succeed",

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
				"client_id":      secret.MaskableString("maskable-token"),
				"client_secret":  secret.MaskableString("maskable-token"),
				"channel_name":   "channel",
				"channel_type":   "",
				"valid_duration": time.Duration(0),
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

			s := lark.NewPluginService(lark.AppConfig{}, log.NewNoop(), mockEncryptor)
			got, err := s.PostHookQueueTransformConfigs(context.TODO(), tt.notificationConfigMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.PostHookQueueTransformConfigs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.PostHookQueueTransformConfigs() = %v, want %v", got, tt.want)
			}
		})
	}
}
