// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	provider "github.com/goto/siren/core/provider"
)

// ConfigSyncer is an autogenerated mock type for the ConfigSyncer type
type ConfigSyncer struct {
	mock.Mock
}

type ConfigSyncer_Expecter struct {
	mock *mock.Mock
}

func (_m *ConfigSyncer) EXPECT() *ConfigSyncer_Expecter {
	return &ConfigSyncer_Expecter{mock: &_m.Mock}
}

// SyncRuntimeConfig provides a mock function with given fields: ctx, namespaceID, namespaceURN, namespaceLabels, prov
func (_m *ConfigSyncer) SyncRuntimeConfig(ctx context.Context, namespaceID uint64, namespaceURN string, namespaceLabels map[string]string, prov provider.Provider) (map[string]string, error) {
	ret := _m.Called(ctx, namespaceID, namespaceURN, namespaceLabels, prov)

	if len(ret) == 0 {
		panic("no return value specified for SyncRuntimeConfig")
	}

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, string, map[string]string, provider.Provider) (map[string]string, error)); ok {
		return rf(ctx, namespaceID, namespaceURN, namespaceLabels, prov)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, string, map[string]string, provider.Provider) map[string]string); ok {
		r0 = rf(ctx, namespaceID, namespaceURN, namespaceLabels, prov)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, string, map[string]string, provider.Provider) error); ok {
		r1 = rf(ctx, namespaceID, namespaceURN, namespaceLabels, prov)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConfigSyncer_SyncRuntimeConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SyncRuntimeConfig'
type ConfigSyncer_SyncRuntimeConfig_Call struct {
	*mock.Call
}

// SyncRuntimeConfig is a helper method to define mock.On call
//   - ctx context.Context
//   - namespaceID uint64
//   - namespaceURN string
//   - namespaceLabels map[string]string
//   - prov provider.Provider
func (_e *ConfigSyncer_Expecter) SyncRuntimeConfig(ctx interface{}, namespaceID interface{}, namespaceURN interface{}, namespaceLabels interface{}, prov interface{}) *ConfigSyncer_SyncRuntimeConfig_Call {
	return &ConfigSyncer_SyncRuntimeConfig_Call{Call: _e.mock.On("SyncRuntimeConfig", ctx, namespaceID, namespaceURN, namespaceLabels, prov)}
}

func (_c *ConfigSyncer_SyncRuntimeConfig_Call) Run(run func(ctx context.Context, namespaceID uint64, namespaceURN string, namespaceLabels map[string]string, prov provider.Provider)) *ConfigSyncer_SyncRuntimeConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(string), args[3].(map[string]string), args[4].(provider.Provider))
	})
	return _c
}

func (_c *ConfigSyncer_SyncRuntimeConfig_Call) Return(_a0 map[string]string, _a1 error) *ConfigSyncer_SyncRuntimeConfig_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ConfigSyncer_SyncRuntimeConfig_Call) RunAndReturn(run func(context.Context, uint64, string, map[string]string, provider.Provider) (map[string]string, error)) *ConfigSyncer_SyncRuntimeConfig_Call {
	_c.Call.Return(run)
	return _c
}

// NewConfigSyncer creates a new instance of ConfigSyncer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConfigSyncer(t interface {
	mock.TestingT
	Cleanup(func())
}) *ConfigSyncer {
	mock := &ConfigSyncer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
