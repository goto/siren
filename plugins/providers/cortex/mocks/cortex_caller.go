// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	rwrulefmt "github.com/grafana/cortex-tools/pkg/rules/rwrulefmt"
)

// CortexCaller is an autogenerated mock type for the CortexCaller type
type CortexCaller struct {
	mock.Mock
}

type CortexCaller_Expecter struct {
	mock *mock.Mock
}

func (_m *CortexCaller) EXPECT() *CortexCaller_Expecter {
	return &CortexCaller_Expecter{mock: &_m.Mock}
}

// CreateAlertmanagerConfig provides a mock function with given fields: ctx, cfg, templates
func (_m *CortexCaller) CreateAlertmanagerConfig(ctx context.Context, cfg string, templates map[string]string) error {
	ret := _m.Called(ctx, cfg, templates)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]string) error); ok {
		r0 = rf(ctx, cfg, templates)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CortexCaller_CreateAlertmanagerConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateAlertmanagerConfig'
type CortexCaller_CreateAlertmanagerConfig_Call struct {
	*mock.Call
}

// CreateAlertmanagerConfig is a helper method to define mock.On call
//  - ctx context.Context
//  - cfg string
//  - templates map[string]string
func (_e *CortexCaller_Expecter) CreateAlertmanagerConfig(ctx interface{}, cfg interface{}, templates interface{}) *CortexCaller_CreateAlertmanagerConfig_Call {
	return &CortexCaller_CreateAlertmanagerConfig_Call{Call: _e.mock.On("CreateAlertmanagerConfig", ctx, cfg, templates)}
}

func (_c *CortexCaller_CreateAlertmanagerConfig_Call) Run(run func(ctx context.Context, cfg string, templates map[string]string)) *CortexCaller_CreateAlertmanagerConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(map[string]string))
	})
	return _c
}

func (_c *CortexCaller_CreateAlertmanagerConfig_Call) Return(_a0 error) *CortexCaller_CreateAlertmanagerConfig_Call {
	_c.Call.Return(_a0)
	return _c
}

// CreateRuleGroup provides a mock function with given fields: ctx, namespace, rg
func (_m *CortexCaller) CreateRuleGroup(ctx context.Context, namespace string, rg rwrulefmt.RuleGroup) error {
	ret := _m.Called(ctx, namespace, rg)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, rwrulefmt.RuleGroup) error); ok {
		r0 = rf(ctx, namespace, rg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CortexCaller_CreateRuleGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateRuleGroup'
type CortexCaller_CreateRuleGroup_Call struct {
	*mock.Call
}

// CreateRuleGroup is a helper method to define mock.On call
//  - ctx context.Context
//  - namespace string
//  - rg rwrulefmt.RuleGroup
func (_e *CortexCaller_Expecter) CreateRuleGroup(ctx interface{}, namespace interface{}, rg interface{}) *CortexCaller_CreateRuleGroup_Call {
	return &CortexCaller_CreateRuleGroup_Call{Call: _e.mock.On("CreateRuleGroup", ctx, namespace, rg)}
}

func (_c *CortexCaller_CreateRuleGroup_Call) Run(run func(ctx context.Context, namespace string, rg rwrulefmt.RuleGroup)) *CortexCaller_CreateRuleGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(rwrulefmt.RuleGroup))
	})
	return _c
}

func (_c *CortexCaller_CreateRuleGroup_Call) Return(_a0 error) *CortexCaller_CreateRuleGroup_Call {
	_c.Call.Return(_a0)
	return _c
}

// DeleteRuleGroup provides a mock function with given fields: ctx, namespace, groupName
func (_m *CortexCaller) DeleteRuleGroup(ctx context.Context, namespace string, groupName string) error {
	ret := _m.Called(ctx, namespace, groupName)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, namespace, groupName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CortexCaller_DeleteRuleGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteRuleGroup'
type CortexCaller_DeleteRuleGroup_Call struct {
	*mock.Call
}

// DeleteRuleGroup is a helper method to define mock.On call
//  - ctx context.Context
//  - namespace string
//  - groupName string
func (_e *CortexCaller_Expecter) DeleteRuleGroup(ctx interface{}, namespace interface{}, groupName interface{}) *CortexCaller_DeleteRuleGroup_Call {
	return &CortexCaller_DeleteRuleGroup_Call{Call: _e.mock.On("DeleteRuleGroup", ctx, namespace, groupName)}
}

func (_c *CortexCaller_DeleteRuleGroup_Call) Run(run func(ctx context.Context, namespace string, groupName string)) *CortexCaller_DeleteRuleGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *CortexCaller_DeleteRuleGroup_Call) Return(_a0 error) *CortexCaller_DeleteRuleGroup_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetRuleGroup provides a mock function with given fields: ctx, namespace, groupName
func (_m *CortexCaller) GetRuleGroup(ctx context.Context, namespace string, groupName string) (*rwrulefmt.RuleGroup, error) {
	ret := _m.Called(ctx, namespace, groupName)

	var r0 *rwrulefmt.RuleGroup
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *rwrulefmt.RuleGroup); ok {
		r0 = rf(ctx, namespace, groupName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rwrulefmt.RuleGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, namespace, groupName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CortexCaller_GetRuleGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRuleGroup'
type CortexCaller_GetRuleGroup_Call struct {
	*mock.Call
}

// GetRuleGroup is a helper method to define mock.On call
//  - ctx context.Context
//  - namespace string
//  - groupName string
func (_e *CortexCaller_Expecter) GetRuleGroup(ctx interface{}, namespace interface{}, groupName interface{}) *CortexCaller_GetRuleGroup_Call {
	return &CortexCaller_GetRuleGroup_Call{Call: _e.mock.On("GetRuleGroup", ctx, namespace, groupName)}
}

func (_c *CortexCaller_GetRuleGroup_Call) Run(run func(ctx context.Context, namespace string, groupName string)) *CortexCaller_GetRuleGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *CortexCaller_GetRuleGroup_Call) Return(_a0 *rwrulefmt.RuleGroup, _a1 error) *CortexCaller_GetRuleGroup_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ListRules provides a mock function with given fields: ctx, namespace
func (_m *CortexCaller) ListRules(ctx context.Context, namespace string) (map[string][]rwrulefmt.RuleGroup, error) {
	ret := _m.Called(ctx, namespace)

	var r0 map[string][]rwrulefmt.RuleGroup
	if rf, ok := ret.Get(0).(func(context.Context, string) map[string][]rwrulefmt.RuleGroup); ok {
		r0 = rf(ctx, namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]rwrulefmt.RuleGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CortexCaller_ListRules_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListRules'
type CortexCaller_ListRules_Call struct {
	*mock.Call
}

// ListRules is a helper method to define mock.On call
//  - ctx context.Context
//  - namespace string
func (_e *CortexCaller_Expecter) ListRules(ctx interface{}, namespace interface{}) *CortexCaller_ListRules_Call {
	return &CortexCaller_ListRules_Call{Call: _e.mock.On("ListRules", ctx, namespace)}
}

func (_c *CortexCaller_ListRules_Call) Run(run func(ctx context.Context, namespace string)) *CortexCaller_ListRules_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *CortexCaller_ListRules_Call) Return(_a0 map[string][]rwrulefmt.RuleGroup, _a1 error) *CortexCaller_ListRules_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewCortexCaller interface {
	mock.TestingT
	Cleanup(func())
}

// NewCortexCaller creates a new instance of CortexCaller. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCortexCaller(t mockConstructorTestingTNewCortexCaller) *CortexCaller {
	mock := &CortexCaller{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
