package base

import (
	"context"

	"github.com/goto/siren/core/notification"
	"github.com/goto/siren/plugins"
)

// UnimplementedService is a base receiver plugin service layer
type UnimplementedService struct{}

func (s *UnimplementedService) PreHookDBTransformConfigs(ctx context.Context, configurations map[string]any) (map[string]any, error) {
	return configurations, nil
}

func (s *UnimplementedService) PostHookDBTransformConfigs(ctx context.Context, configurations map[string]any) (map[string]any, error) {
	return configurations, nil
}

func (s *UnimplementedService) BuildData(ctx context.Context, configurations map[string]any) (map[string]any, error) {
	return map[string]any{}, nil
}

func (s *UnimplementedService) PreHookQueueTransformConfigs(ctx context.Context, notificationConfigMap map[string]any) (map[string]any, error) {
	return notificationConfigMap, nil
}

func (s *UnimplementedService) PostHookQueueTransformConfigs(ctx context.Context, notificationConfigMap map[string]any) (map[string]any, error) {
	return notificationConfigMap, nil
}

func (s *UnimplementedService) GetSystemDefaultTemplate() string {
	return ""
}

func (s *UnimplementedService) PostProcessMessage(mm notification.MetaMessage, m *notification.Message) *notification.Message {
	return m
}

func (s *UnimplementedService) Send(ctx context.Context, notificationMessage notification.Message) (bool, error) {
	return false, plugins.ErrNotImplemented
}
