// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"

	notification "github.com/goto/siren/core/notification"
	mock "github.com/stretchr/testify/mock"
)

// NotificationService is an autogenerated mock type for the NotificationService type
type NotificationService struct {
	mock.Mock
}

type NotificationService_Expecter struct {
	mock *mock.Mock
}

func (_m *NotificationService) EXPECT() *NotificationService_Expecter {
	return &NotificationService_Expecter{mock: &_m.Mock}
}

// Dispatch provides a mock function with given fields: _a0, _a1
func (_m *NotificationService) Dispatch(_a0 context.Context, _a1 []notification.Notification) ([]string, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Dispatch")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []notification.Notification) ([]string, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []notification.Notification) []string); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []notification.Notification) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NotificationService_Dispatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Dispatch'
type NotificationService_Dispatch_Call struct {
	*mock.Call
}

// Dispatch is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 []notification.Notification
func (_e *NotificationService_Expecter) Dispatch(_a0 interface{}, _a1 interface{}) *NotificationService_Dispatch_Call {
	return &NotificationService_Dispatch_Call{Call: _e.mock.On("Dispatch", _a0, _a1)}
}

func (_c *NotificationService_Dispatch_Call) Run(run func(_a0 context.Context, _a1 []notification.Notification)) *NotificationService_Dispatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]notification.Notification))
	})
	return _c
}

func (_c *NotificationService_Dispatch_Call) Return(_a0 []string, _a1 error) *NotificationService_Dispatch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NotificationService_Dispatch_Call) RunAndReturn(run func(context.Context, []notification.Notification) ([]string, error)) *NotificationService_Dispatch_Call {
	_c.Call.Return(run)
	return _c
}

// NewNotificationService creates a new instance of NotificationService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNotificationService(t interface {
	mock.TestingT
	Cleanup(func())
}) *NotificationService {
	mock := &NotificationService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
