// Code generated by mockery v2.33.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// ConfigResolver is an autogenerated mock type for the ConfigResolver type
type ConfigResolver struct {
	mock.Mock
}

type ConfigResolver_Expecter struct {
	mock *mock.Mock
}

func (_m *ConfigResolver) EXPECT() *ConfigResolver_Expecter {
	return &ConfigResolver_Expecter{mock: &_m.Mock}
}

// BuildData provides a mock function with given fields: ctx, configs
func (_m *ConfigResolver) BuildData(ctx context.Context, configs map[string]interface{}) (map[string]interface{}, error) {
	ret := _m.Called(ctx, configs)

	var r0 map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}) (map[string]interface{}, error)); ok {
		return rf(ctx, configs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}) map[string]interface{}); ok {
		r0 = rf(ctx, configs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, map[string]interface{}) error); ok {
		r1 = rf(ctx, configs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConfigResolver_BuildData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BuildData'
type ConfigResolver_BuildData_Call struct {
	*mock.Call
}

// BuildData is a helper method to define mock.On call
//   - ctx context.Context
//   - configs map[string]interface{}
func (_e *ConfigResolver_Expecter) BuildData(ctx interface{}, configs interface{}) *ConfigResolver_BuildData_Call {
	return &ConfigResolver_BuildData_Call{Call: _e.mock.On("BuildData", ctx, configs)}
}

func (_c *ConfigResolver_BuildData_Call) Run(run func(ctx context.Context, configs map[string]interface{})) *ConfigResolver_BuildData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(map[string]interface{}))
	})
	return _c
}

func (_c *ConfigResolver_BuildData_Call) Return(_a0 map[string]interface{}, _a1 error) *ConfigResolver_BuildData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ConfigResolver_BuildData_Call) RunAndReturn(run func(context.Context, map[string]interface{}) (map[string]interface{}, error)) *ConfigResolver_BuildData_Call {
	_c.Call.Return(run)
	return _c
}

// PostHookDBTransformConfigs provides a mock function with given fields: ctx, configs
func (_m *ConfigResolver) PostHookDBTransformConfigs(ctx context.Context, configs map[string]interface{}) (map[string]interface{}, error) {
	ret := _m.Called(ctx, configs)

	var r0 map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}) (map[string]interface{}, error)); ok {
		return rf(ctx, configs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}) map[string]interface{}); ok {
		r0 = rf(ctx, configs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, map[string]interface{}) error); ok {
		r1 = rf(ctx, configs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConfigResolver_PostHookDBTransformConfigs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PostHookDBTransformConfigs'
type ConfigResolver_PostHookDBTransformConfigs_Call struct {
	*mock.Call
}

// PostHookDBTransformConfigs is a helper method to define mock.On call
//   - ctx context.Context
//   - configs map[string]interface{}
func (_e *ConfigResolver_Expecter) PostHookDBTransformConfigs(ctx interface{}, configs interface{}) *ConfigResolver_PostHookDBTransformConfigs_Call {
	return &ConfigResolver_PostHookDBTransformConfigs_Call{Call: _e.mock.On("PostHookDBTransformConfigs", ctx, configs)}
}

func (_c *ConfigResolver_PostHookDBTransformConfigs_Call) Run(run func(ctx context.Context, configs map[string]interface{})) *ConfigResolver_PostHookDBTransformConfigs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(map[string]interface{}))
	})
	return _c
}

func (_c *ConfigResolver_PostHookDBTransformConfigs_Call) Return(_a0 map[string]interface{}, _a1 error) *ConfigResolver_PostHookDBTransformConfigs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ConfigResolver_PostHookDBTransformConfigs_Call) RunAndReturn(run func(context.Context, map[string]interface{}) (map[string]interface{}, error)) *ConfigResolver_PostHookDBTransformConfigs_Call {
	_c.Call.Return(run)
	return _c
}

// PreHookDBTransformConfigs provides a mock function with given fields: ctx, configs
func (_m *ConfigResolver) PreHookDBTransformConfigs(ctx context.Context, configs map[string]interface{}) (map[string]interface{}, error) {
	ret := _m.Called(ctx, configs)

	var r0 map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}) (map[string]interface{}, error)); ok {
		return rf(ctx, configs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}) map[string]interface{}); ok {
		r0 = rf(ctx, configs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, map[string]interface{}) error); ok {
		r1 = rf(ctx, configs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConfigResolver_PreHookDBTransformConfigs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PreHookDBTransformConfigs'
type ConfigResolver_PreHookDBTransformConfigs_Call struct {
	*mock.Call
}

// PreHookDBTransformConfigs is a helper method to define mock.On call
//   - ctx context.Context
//   - configs map[string]interface{}
func (_e *ConfigResolver_Expecter) PreHookDBTransformConfigs(ctx interface{}, configs interface{}) *ConfigResolver_PreHookDBTransformConfigs_Call {
	return &ConfigResolver_PreHookDBTransformConfigs_Call{Call: _e.mock.On("PreHookDBTransformConfigs", ctx, configs)}
}

func (_c *ConfigResolver_PreHookDBTransformConfigs_Call) Run(run func(ctx context.Context, configs map[string]interface{})) *ConfigResolver_PreHookDBTransformConfigs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(map[string]interface{}))
	})
	return _c
}

func (_c *ConfigResolver_PreHookDBTransformConfigs_Call) Return(_a0 map[string]interface{}, _a1 error) *ConfigResolver_PreHookDBTransformConfigs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ConfigResolver_PreHookDBTransformConfigs_Call) RunAndReturn(run func(context.Context, map[string]interface{}) (map[string]interface{}, error)) *ConfigResolver_PreHookDBTransformConfigs_Call {
	_c.Call.Return(run)
	return _c
}

// NewConfigResolver creates a new instance of ConfigResolver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConfigResolver(t interface {
	mock.TestingT
	Cleanup(func())
}) *ConfigResolver {
	mock := &ConfigResolver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
