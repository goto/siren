// Code generated by mockery v2.33.2. DO NOT EDIT.

package mocks

import (
	context "context"

	subscription "github.com/goto/siren/core/subscription"
	mock "github.com/stretchr/testify/mock"
)

// SubscriptionService is an autogenerated mock type for the SubscriptionService type
type SubscriptionService struct {
	mock.Mock
}

type SubscriptionService_Expecter struct {
	mock *mock.Mock
}

func (_m *SubscriptionService) EXPECT() *SubscriptionService_Expecter {
	return &SubscriptionService_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *SubscriptionService) Create(_a0 context.Context, _a1 *subscription.Subscription) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *subscription.Subscription) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubscriptionService_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type SubscriptionService_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *subscription.Subscription
func (_e *SubscriptionService_Expecter) Create(_a0 interface{}, _a1 interface{}) *SubscriptionService_Create_Call {
	return &SubscriptionService_Create_Call{Call: _e.mock.On("Create", _a0, _a1)}
}

func (_c *SubscriptionService_Create_Call) Run(run func(_a0 context.Context, _a1 *subscription.Subscription)) *SubscriptionService_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*subscription.Subscription))
	})
	return _c
}

func (_c *SubscriptionService_Create_Call) Return(_a0 error) *SubscriptionService_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SubscriptionService_Create_Call) RunAndReturn(run func(context.Context, *subscription.Subscription) error) *SubscriptionService_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *SubscriptionService) Delete(_a0 context.Context, _a1 uint64) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubscriptionService_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type SubscriptionService_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 uint64
func (_e *SubscriptionService_Expecter) Delete(_a0 interface{}, _a1 interface{}) *SubscriptionService_Delete_Call {
	return &SubscriptionService_Delete_Call{Call: _e.mock.On("Delete", _a0, _a1)}
}

func (_c *SubscriptionService_Delete_Call) Run(run func(_a0 context.Context, _a1 uint64)) *SubscriptionService_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *SubscriptionService_Delete_Call) Return(_a0 error) *SubscriptionService_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SubscriptionService_Delete_Call) RunAndReturn(run func(context.Context, uint64) error) *SubscriptionService_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *SubscriptionService) Get(_a0 context.Context, _a1 uint64) (*subscription.Subscription, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *subscription.Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*subscription.Subscription, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *subscription.Subscription); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*subscription.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubscriptionService_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type SubscriptionService_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 uint64
func (_e *SubscriptionService_Expecter) Get(_a0 interface{}, _a1 interface{}) *SubscriptionService_Get_Call {
	return &SubscriptionService_Get_Call{Call: _e.mock.On("Get", _a0, _a1)}
}

func (_c *SubscriptionService_Get_Call) Run(run func(_a0 context.Context, _a1 uint64)) *SubscriptionService_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *SubscriptionService_Get_Call) Return(_a0 *subscription.Subscription, _a1 error) *SubscriptionService_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SubscriptionService_Get_Call) RunAndReturn(run func(context.Context, uint64) (*subscription.Subscription, error)) *SubscriptionService_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: _a0, _a1
func (_m *SubscriptionService) List(_a0 context.Context, _a1 subscription.Filter) ([]subscription.Subscription, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []subscription.Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, subscription.Filter) ([]subscription.Subscription, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, subscription.Filter) []subscription.Subscription); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]subscription.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, subscription.Filter) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubscriptionService_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type SubscriptionService_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 subscription.Filter
func (_e *SubscriptionService_Expecter) List(_a0 interface{}, _a1 interface{}) *SubscriptionService_List_Call {
	return &SubscriptionService_List_Call{Call: _e.mock.On("List", _a0, _a1)}
}

func (_c *SubscriptionService_List_Call) Run(run func(_a0 context.Context, _a1 subscription.Filter)) *SubscriptionService_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(subscription.Filter))
	})
	return _c
}

func (_c *SubscriptionService_List_Call) Return(_a0 []subscription.Subscription, _a1 error) *SubscriptionService_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SubscriptionService_List_Call) RunAndReturn(run func(context.Context, subscription.Filter) ([]subscription.Subscription, error)) *SubscriptionService_List_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *SubscriptionService) Update(_a0 context.Context, _a1 *subscription.Subscription) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *subscription.Subscription) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubscriptionService_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type SubscriptionService_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *subscription.Subscription
func (_e *SubscriptionService_Expecter) Update(_a0 interface{}, _a1 interface{}) *SubscriptionService_Update_Call {
	return &SubscriptionService_Update_Call{Call: _e.mock.On("Update", _a0, _a1)}
}

func (_c *SubscriptionService_Update_Call) Run(run func(_a0 context.Context, _a1 *subscription.Subscription)) *SubscriptionService_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*subscription.Subscription))
	})
	return _c
}

func (_c *SubscriptionService_Update_Call) Return(_a0 error) *SubscriptionService_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SubscriptionService_Update_Call) RunAndReturn(run func(context.Context, *subscription.Subscription) error) *SubscriptionService_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewSubscriptionService creates a new instance of SubscriptionService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSubscriptionService(t interface {
	mock.TestingT
	Cleanup(func())
}) *SubscriptionService {
	mock := &SubscriptionService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
