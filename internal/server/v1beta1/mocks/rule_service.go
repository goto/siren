// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	context "context"

	rule "github.com/odpf/siren/core/rule"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
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

// Get provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5
func (_m *RuleService) Get(_a0 context.Context, _a1 string, _a2 string, _a3 string, _a4 string, _a5 uint64) ([]rule.Rule, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5)

	var r0 []rule.Rule
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, uint64) []rule.Rule); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]rule.Rule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string, uint64) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4, _a5)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RuleService_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type RuleService_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 string
//  - _a2 string
//  - _a3 string
//  - _a4 string
//  - _a5 uint64
func (_e *RuleService_Expecter) Get(_a0 interface{}, _a1 interface{}, _a2 interface{}, _a3 interface{}, _a4 interface{}, _a5 interface{}) *RuleService_Get_Call {
	return &RuleService_Get_Call{Call: _e.mock.On("Get", _a0, _a1, _a2, _a3, _a4, _a5)}
}

func (_c *RuleService_Get_Call) Run(run func(_a0 context.Context, _a1 string, _a2 string, _a3 string, _a4 string, _a5 uint64)) *RuleService_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string), args[4].(string), args[5].(uint64))
	})
	return _c
}

func (_c *RuleService_Get_Call) Return(_a0 []rule.Rule, _a1 error) *RuleService_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Migrate provides a mock function with given fields:
func (_m *RuleService) Migrate() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RuleService_Migrate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Migrate'
type RuleService_Migrate_Call struct {
	*mock.Call
}

// Migrate is a helper method to define mock.On call
func (_e *RuleService_Expecter) Migrate() *RuleService_Migrate_Call {
	return &RuleService_Migrate_Call{Call: _e.mock.On("Migrate")}
}

func (_c *RuleService_Migrate_Call) Run(run func()) *RuleService_Migrate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RuleService_Migrate_Call) Return(_a0 error) *RuleService_Migrate_Call {
	_c.Call.Return(_a0)
	return _c
}

// Upsert provides a mock function with given fields: _a0, _a1
func (_m *RuleService) Upsert(_a0 context.Context, _a1 *rule.Rule) error {
	ret := _m.Called(_a0, _a1)

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
//  - _a0 context.Context
//  - _a1 *rule.Rule
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

// NewRuleService creates a new instance of RuleService. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewRuleService(t testing.TB) *RuleService {
	mock := &RuleService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
