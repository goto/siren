package lark

import (
	"fmt"
	"time"

	"github.com/goto/siren/pkg/errors"
	"github.com/goto/siren/pkg/httpclient"
	"github.com/goto/siren/pkg/retry"
	"github.com/goto/siren/pkg/secret"
)

// AppConfig is a config loaded when siren is started
type AppConfig struct {
	APIHost       string            `mapstructure:"api_host"`
	Retry         retry.Config      `mapstructure:"retry"`
	HTTPClient    httpclient.Config `mapstructure:"http_client"`
	ValidDuration time.Duration     `mapstructure:"valid_duration" yaml:"valid_duration"`
}

// LarkCredentialConfig is config that needs to be passed when a new lark
// receiver is being added
type LarkCredentialConfig struct {
	ClientID     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
}

func (c *LarkCredentialConfig) Validate() error {
	if c.ClientID != "" && c.ClientSecret != "" {
		return nil
	}
	return fmt.Errorf("invalid lark credentials, client_id: %s, client_secret: %s", c.ClientID, c.ClientSecret)
}

// ReceiverConfig is a stored config for a lark receiver
type ReceiverConfig struct {
	ClientID      secret.MaskableString `mapstructure:"client_id"`
	ClientSecret  secret.MaskableString `mapstructure:"client_secret"`
	ValidDuration time.Duration         `mapstructure:"valid_duration" yaml:"valid_duration"`
}

func (c *ReceiverConfig) Validate() error {
	if c.ClientID != "" && c.ClientSecret != "" {
		return nil
	}
	return errors.ErrInvalid.WithMsgf("invalid lark receiver config, workspace: %s, token: %s", c.ClientID, c.ClientSecret)
}

func (c *ReceiverConfig) AsMap() map[string]any {
	return map[string]any{
		"client_id":      c.ClientID,
		"client_secret":  c.ClientSecret,
		"valid_duration": c.ValidDuration,
	}
}

// ReceiverData is a stored data for a lark receiver
type ReceiverData struct {
	Channels string `json:"channels" mapstructure:"channels"`
}

func (c *ReceiverData) AsMap() map[string]any {
	return map[string]any{
		"channels": c.Channels,
	}
}

// SubscriptionConfig is a stored config for a subscription of a lark receiver
type SubscriptionConfig struct {
	ChannelName string `json:"channel_name" mapstructure:"channel_name"`
	ChannelType string `json:"channel_type" mapstructure:"channel_type"`
}

func (c *SubscriptionConfig) AsMap() map[string]any {
	return map[string]any{
		"channel_name": c.ChannelName,
		"channel_type": c.ChannelType,
	}
}

// NotificationConfig has all configs needed to send notification
type NotificationConfig struct {
	ReceiverConfig     `mapstructure:",squash"`
	SubscriptionConfig `mapstructure:",squash"`
}

// Validate validates whether notification config contains required fields or not
// channel_name is not mandatory because in NotifyToReceiver flow, channel_name
// is being passed from the request (not from the config)
func (c *NotificationConfig) Validate() error {
	if c.ClientID != "" && c.ClientSecret != "" {
		return nil
	}
	return fmt.Errorf("invalid lark notification config, workspace: %s, token: %s, channel_name: %s", c.ClientSecret, c.ClientID, c.ChannelName)
}

func (c *NotificationConfig) AsMap() map[string]any {
	notificationMap := make(map[string]any)

	for k, v := range c.ReceiverConfig.AsMap() {
		notificationMap[k] = v
	}

	for k, v := range c.SubscriptionConfig.AsMap() {
		notificationMap[k] = v
	}

	return notificationMap
}
