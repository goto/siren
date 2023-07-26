package receiver

import "context"

// ConfigResolver is an interface for the receiver to resolve all configs and function
// related to a specific receiver type. Receiver plugin needs to implement
// this interface for all configs and functionality resolution.
//
//go:generate mockery --name=ConfigResolver -r --case underscore --with-expecter --structname ConfigResolver --filename config_resolver.go --output=./mocks
type ConfigResolver interface {
	BuildData(ctx context.Context, configs map[string]any) (map[string]any, error)
	PreHookDBTransformConfigs(ctx context.Context, configs map[string]any, parentID uint64) (map[string]any, error)
	PostHookDBTransformConfigs(ctx context.Context, configs map[string]any) (map[string]any, error)
}
