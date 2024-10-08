package lark

import (
	"github.com/goto/siren/pkg/httpclient"
	"github.com/goto/siren/pkg/retry"
)

type ServiceOption func(*PluginService)

// WithHTTPClient assigns custom http client when creating a lark service
func WithHTTPClient(httpClient *httpclient.Client) ServiceOption {
	return func(s *PluginService) {
		s.httpClient = httpClient
	}
}

// WithRetrier wraps client call with retrier
func WithRetrier(runner retry.Runner) ServiceOption {
	return func(s *PluginService) {
		// note: for now retry only happen in send message context method
		s.retrier = runner
	}
}

func WithLarkClient(client LarkCaller) ServiceOption {
	return func(s *PluginService) {
		s.client = client
	}
}
