package larkchannel

import (
	"context"
	"fmt"

	"github.com/goto/siren/core/notification"
	"github.com/goto/siren/pkg/errors"
	"github.com/goto/siren/plugins/receivers/base"
	"github.com/goto/siren/plugins/receivers/lark"
	"github.com/mitchellh/mapstructure"
)

// PluginService is a plugin service layer for lark channel type
type PluginService struct {
	base.UnimplementedService
	larkPluginService *lark.PluginService
}

// NewPluginService returns lark channel plugin service struct. This service implement [receiver.Resolver] and [notification.Notifier] interface.
func NewPluginService(cfg lark.AppConfig, cryptoClient lark.Encryptor, opts ...lark.ServiceOption) *PluginService {
	return &PluginService{
		larkPluginService: lark.NewPluginService(cfg, cryptoClient, opts...),
	}
}

func (s *PluginService) PreHookDBTransformConfigs(ctx context.Context, configurations map[string]any) (map[string]any, error) {
	receiverConfig := &ReceiverConfig{}
	if err := mapstructure.Decode(configurations, receiverConfig); err != nil {
		return nil, fmt.Errorf("failed to transform configurations to receiver config: %w", err)
	}

	if err := receiverConfig.Validate(); err != nil {
		return nil, errors.ErrInvalid.WithMsgf(err.Error())
	}

	return configurations, nil
}

// PostHookTransformConfigs do transformation in post-hook service lifecycle
func (s *PluginService) PostHookDBTransformConfigs(ctx context.Context, configurations map[string]any) (map[string]any, error) {
	transformedConfigs, err := s.larkPluginService.PostHookDBTransformConfigs(ctx, configurations)
	// if lark_channel is not expaneded, it is okay to have lark config empty
	if err != nil && !errors.Is(err, errors.ErrInvalid) {
		return nil, fmt.Errorf("lark channel post hook db failed: %w", err)
	}

	var mergedConfigs = map[string]any{}
	for k, v := range configurations {
		if value, ok := transformedConfigs[k]; ok {
			mergedConfigs[k] = value
		} else {
			mergedConfigs[k] = v
		}
	}

	return mergedConfigs, nil
}

func (s *PluginService) PreHookQueueTransformConfigs(ctx context.Context, notificationConfigMap map[string]any) (map[string]any, error) {
	transformedConfigs, err := s.larkPluginService.PreHookQueueTransformConfigs(ctx, notificationConfigMap)
	if err != nil {
		return nil, err
	}

	var mergedConfigs = map[string]any{}
	for k, v := range notificationConfigMap {
		if value, ok := transformedConfigs[k]; ok {
			mergedConfigs[k] = value
		} else {
			mergedConfigs[k] = v
		}
	}

	return mergedConfigs, nil
}

func (s *PluginService) PostHookQueueTransformConfigs(ctx context.Context, notificationConfigMap map[string]any) (map[string]any, error) {
	transformedConfigs, err := s.larkPluginService.PostHookQueueTransformConfigs(ctx, notificationConfigMap)
	if err != nil {
		return nil, err
	}

	var mergedConfigs = map[string]any{}
	for k, v := range notificationConfigMap {
		if value, ok := transformedConfigs[k]; ok {
			mergedConfigs[k] = value
		} else {
			mergedConfigs[k] = v
		}
	}

	return mergedConfigs, nil
}

func (s *PluginService) Send(ctx context.Context, notificationMessage notification.Message) (bool, error) {
	return s.larkPluginService.Send(ctx, notificationMessage)
}

func (s *PluginService) GetSystemDefaultTemplate() string {
	return s.larkPluginService.GetSystemDefaultTemplate()
}
