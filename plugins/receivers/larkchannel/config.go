package larkchannel

import (
	"fmt"

	"github.com/goto/siren/plugins/receivers/lark"
)

// ReceiverConfig is a stored config for a lark receiver
type ReceiverConfig struct {
	LarkReceiverConfig lark.ReceiverConfig `mapstructure:",squash"`
	ChannelName        string              `json:"channel_name" mapstructure:"channel_name"`
	ChannelType        string              `json:"channel_type" mapstructure:"channel_type"`
}

func (c *ReceiverConfig) Validate() error {
	if c.ChannelName == "" {
		return fmt.Errorf("invalid lark_channel receiver config, channel_name can't be empty")
	}
	return nil
}

func (c *ReceiverConfig) AsMap() map[string]any {
	return map[string]any{
		"client_id":      c.LarkReceiverConfig.ClientID,
		"client_secret":  c.LarkReceiverConfig.ClientSecret,
		"channel_name":   c.ChannelName,
		"channel_type":   c.ChannelType,
		"valid_duration": c.LarkReceiverConfig.ValidDuration,
	}
}

// NotificationConfig has all configs needed to send notification
type NotificationConfig struct {
	ReceiverConfig `mapstructure:",squash"`
}

// Validate validates whether notification config contains required fields or not
// channel_name is not mandatory because in NotifyToReceiver flow, channel_name
// is being passed from the request (not from the config)
func (c *NotificationConfig) Validate() error {
	if err := c.ReceiverConfig.LarkReceiverConfig.Validate(); err != nil {
		return err
	}
	if err := c.ReceiverConfig.Validate(); err != nil {
		return err
	}
	return nil
}

func (c *NotificationConfig) AsMap() map[string]any {
	notificationMap := make(map[string]any)

	for k, v := range c.ReceiverConfig.AsMap() {
		notificationMap[k] = v
	}

	return notificationMap
}
