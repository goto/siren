// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"

	rule "github.com/goto/siren/core/rule"
	mock "github.com/stretchr/testify/mock"
)

// RuleService is an autogenerated mock type for the RuleService type
type RuleService struct {
	mock.Mock
}

type RuleService_Expecter struct {
	mock *mock.Mock
}

func (_m *RuleService) EXPECT() *RuleService_Expecter {
	return &RuleService_Expecter{mock: &_m.Mock}
}

// List provides a mock function with given fields: _a0, _a1
func (_m *RuleService) List(_a0 context.Context, _a1 rule.Filter) ([]rule.Rule, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []rule.Rule
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, rule.Filter) ([]rule.Rule, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, rule.Filter) []rule.Rule); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]rule.Rule)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, rule.Filter) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RuleService_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type RuleService_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 rule.Filter
func (_e *RuleService_Expecter) List(_a0 interface{}, _a1 interface{}) *RuleService_List_Call {
	return &RuleService_List_Call{Call: _e.mock.On("List", _a0, _a1)}
}

func (_c *RuleService_List_Call) Run(run func(_a0 context.Context, _a1 rule.Filter)) *RuleService_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(rule.Filter))
	})
	return _c
}

func (_c *RuleService_List_Call) Return(_a0 []rule.Rule, _a1 error) *RuleService_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RuleService_List_Call) RunAndReturn(run func(context.Context, rule.Filter) ([]rule.Rule, error)) *RuleService_List_Call {
	_c.Call.Return(run)
	return _c
}

// Upsert provides a mock function with given fields: _a0, _a1
func (_m *RuleService) Upsert(_a0 context.Context, _a1 *rule.Rule) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Upsert")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *rule.Rule) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RuleService_Upsert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Upsert'
type RuleService_Upsert_Call struct {
	*mock.Call
}

// Upsert is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *rule.Rule
func (_e *RuleService_Expecter) Upsert(_a0 interface{}, _a1 interface{}) *RuleService_Upsert_Call {
	return &RuleService_Upsert_Call{Call: _e.mock.On("Upsert", _a0, _a1)}
}

func (_c *RuleService_Upsert_Call) Run(run func(_a0 context.Context, _a1 *rule.Rule)) *RuleService_Upsert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*rule.Rule))
	})
	return _c
}

func (_c *RuleService_Upsert_Call) Return(_a0 error) *RuleService_Upsert_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RuleService_Upsert_Call) RunAndReturn(run func(context.Context, *rule.Rule) error) *RuleService_Upsert_Call {
	_c.Call.Return(run)
	return _c
}

// NewRuleService creates a new instance of RuleService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRuleService(t interface {
	mock.TestingT
	Cleanup(func())
}) *RuleService {
	mock := &RuleService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
