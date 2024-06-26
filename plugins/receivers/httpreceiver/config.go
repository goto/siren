package httpreceiver

import (
	"fmt"

	"github.com/goto/siren/pkg/httpclient"
	"github.com/goto/siren/pkg/retry"
)

// AppConfig is a config loaded when siren is started
type AppConfig struct {
	Retry      retry.Config      `mapstructure:"retry" yaml:"retry"`
	HTTPClient httpclient.Config `mapstructure:"http_client" yaml:"http_client"`
}
type ReceiverConfig struct {
	URL string `mapstructure:"url"`
}

func (c *ReceiverConfig) Validate() error {
	if c.URL == "" {
		return fmt.Errorf("invalid http receiver config, url: %s", c.URL)
	}
	return nil
}

func (c *ReceiverConfig) AsMap() map[string]any {
	return map[string]any{
		"url": c.URL,
	}
}

type NotificationConfig struct {
	ReceiverConfig `mapstructure:",squash"`
}

func (c *NotificationConfig) AsMap() map[string]any {
	return c.ReceiverConfig.AsMap()
}
