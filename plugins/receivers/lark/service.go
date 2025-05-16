package lark

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/goto/salt/log"
	"github.com/goto/siren/core/notification"
	"github.com/goto/siren/pkg/errors"
	"github.com/goto/siren/pkg/httpclient"
	"github.com/goto/siren/pkg/retry"
	"github.com/goto/siren/plugins/receivers/base"
	"github.com/mitchellh/mapstructure"
)

const (
	TypeChannelChannel = "channel"
	TypeChannelUser    = "user"

	defaultChannelType = TypeChannelChannel
)

// PluginService is a plugin service layer for lark
type PluginService struct {
	base.UnimplementedService
	client       LarkCaller
	cryptoClient Encryptor
	httpClient   *httpclient.Client
	retrier      retry.Runner
}

// NewPluginService returns lark plugin service struct. This service implement [receiver.Resolver] and [notification.Notifier] interface.
func NewPluginService(cfg AppConfig, logger log.Logger, cryptoClient Encryptor, opts ...ServiceOption) *PluginService {
	s := &PluginService{}

	for _, opt := range opts {
		opt(s)
	}

	s.cryptoClient = cryptoClient

	if s.httpClient == nil {
		s.httpClient = httpclient.New(cfg.HTTPClient)
	}

	if s.client == nil {
		s.client = NewClient(cfg, logger, ClientWithHTTPClient(s.httpClient), ClientWithRetrier(s.retrier))
	}

	return s
}

func (s *PluginService) PreHookDBTransformConfigs(ctx context.Context, configurations map[string]any) (map[string]any, error) {
	larkCredentialConfig := &ReceiverConfig{}
	if err := mapstructure.Decode(configurations, larkCredentialConfig); err != nil {
		return nil, fmt.Errorf("failed to transform configurations to pre transform config: %w", err)
	}

	if err := larkCredentialConfig.Validate(); err != nil {
		return nil, errors.ErrInvalid.WithMsgf(err.Error())
	}

	encryptedClientId, err := s.cryptoClient.Encrypt(larkCredentialConfig.ClientID)
	if err != nil {
		return nil, fmt.Errorf("lark clientId encryption failed: %w", err)
	}
	encryptedClientSecret, err := s.cryptoClient.Encrypt(larkCredentialConfig.ClientSecret)
	if err != nil {
		return nil, fmt.Errorf("lark clientSecret encryption failed: %w", err)
	}

	receiverConfig := ReceiverConfig{
		ClientID:     encryptedClientId,
		ClientSecret: encryptedClientSecret,
	}

	return receiverConfig.AsMap(), nil
}

func (s *PluginService) PostHookDBTransformConfigs(ctx context.Context, configurations map[string]any) (map[string]any, error) {
	receiverConfig := &ReceiverConfig{}
	if err := mapstructure.Decode(configurations, receiverConfig); err != nil {
		return nil, fmt.Errorf("failed to transform configurations to notification config: %w", err)
	}
	if err := receiverConfig.Validate(); err != nil {
		return nil, err
	}

	clientId, err := s.cryptoClient.Decrypt(receiverConfig.ClientID)
	if err != nil {
		return nil, fmt.Errorf("lark clientId decryption failed: %w", err)
	}
	clientSecret, err := s.cryptoClient.Decrypt(receiverConfig.ClientSecret)
	if err != nil {
		return nil, fmt.Errorf("lark clientSecret decryption failed: %w", err)
	}

	receiverConfig.ClientID = clientId
	receiverConfig.ClientSecret = clientSecret

	return receiverConfig.AsMap(), nil
}

func (s *PluginService) PreHookQueueTransformConfigs(ctx context.Context, configurations map[string]any) (map[string]any, error) {
	larkCredentialConfig := &ReceiverConfig{}
	if err := mapstructure.Decode(configurations, larkCredentialConfig); err != nil {
		return nil, fmt.Errorf("failed to transform configurations to pre transform config: %w", err)
	}

	if err := larkCredentialConfig.Validate(); err != nil {
		return nil, errors.ErrInvalid.WithMsgf(err.Error())
	}

	encryptedClientId, err := s.cryptoClient.Encrypt(larkCredentialConfig.ClientID)
	if err != nil {
		return nil, fmt.Errorf("lark clientId encryption failed: %w", err)
	}
	encryptedClientSecret, err := s.cryptoClient.Encrypt(larkCredentialConfig.ClientSecret)
	if err != nil {
		return nil, fmt.Errorf("lark clientSecret encryption failed: %w", err)
	}

	receiverConfig := ReceiverConfig{
		ClientID:     encryptedClientId,
		ClientSecret: encryptedClientSecret,
	}

	return receiverConfig.AsMap(), nil
}
func (s *PluginService) PostHookQueueTransformConfigs(ctx context.Context, notificationConfigMap map[string]any) (map[string]any, error) {
	notificationConfig := &NotificationConfig{}
	if err := mapstructure.Decode(notificationConfigMap, notificationConfig); err != nil {
		return nil, fmt.Errorf("failed to transform configurations to notification config: %w", err)
	}

	if err := notificationConfig.Validate(); err != nil {
		return nil, err
	}

	clientId, err := s.cryptoClient.Decrypt(notificationConfig.ClientID)
	if err != nil {
		return nil, fmt.Errorf("lark clientId decryption failed: %w", err)
	}
	clientSecret, err := s.cryptoClient.Decrypt(notificationConfig.ClientSecret)
	if err != nil {
		return nil, fmt.Errorf("lark clientSecret decryption failed: %w", err)
	}

	notificationConfig.ClientID = clientId
	notificationConfig.ClientSecret = clientSecret

	return notificationConfig.AsMap(), nil
}

// BuildData populates receiver data field based on config
func (s *PluginService) BuildData(ctx context.Context, configurations map[string]any) (map[string]any, error) {
	receiverConfig := &ReceiverConfig{}
	if err := mapstructure.Decode(configurations, receiverConfig); err != nil {
		return nil, fmt.Errorf("failed to transform configurations to receiver config: %w", err)
	}

	if err := receiverConfig.Validate(); err != nil {
		return nil, err
	}

	channels, err := s.client.GetWorkspaceChannels(
		ctx,
		receiverConfig.ClientID,
		receiverConfig.ClientSecret,
	)
	if err != nil {
		return nil, fmt.Errorf("could not get channels: %w", err)
	}

	data, err := json.Marshal(channels)
	if err != nil {
		// this is very unlikely to return error since we have an explicitly defined type of channels
		return nil, fmt.Errorf("invalid channels: %w", err)
	}

	receiverData := ReceiverData{
		Channels: string(data),
	}

	return receiverData.AsMap(), nil
}

func (s *PluginService) Send(ctx context.Context, notificationMessage notification.Message) (bool, error) {
	notificationConfig := &NotificationConfig{}
	if err := mapstructure.Decode(notificationMessage.Configs, notificationConfig); err != nil {
		return false, err
	}

	larkMessage := &Message{}
	if err := mapstructure.Decode(notificationMessage.Details, &larkMessage); err != nil {
		return false, err
	}

	if notificationConfig.ChannelType == "" {
		notificationConfig.ChannelType = defaultChannelType
	}
	if notificationConfig.ChannelName != "" {
		larkMessage.Channel = notificationConfig.ChannelName
	}

	if err := s.client.Notify(ctx, *notificationConfig, *larkMessage); err != nil {
		if errors.As(err, new(retry.RetryableError)) {
			return true, err
		} else {
			return false, err
		}
	}

	return false, nil
}

func (s *PluginService) GetSystemDefaultTemplate() string {
	return defaultAlertTemplateBody
}
