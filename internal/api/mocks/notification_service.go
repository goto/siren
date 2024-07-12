// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	notification "github.com/goto/siren/core/notification"
	mock "github.com/stretchr/testify/mock"

	time "time"
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

// CheckIdempotency provides a mock function with given fields: ctx, scope, key
func (_m *NotificationService) CheckIdempotency(ctx context.Context, scope string, key string) (string, error) {
	ret := _m.Called(ctx, scope, key)

	if len(ret) == 0 {
		panic("no return value specified for CheckIdempotency")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (string, error)); ok {
		return rf(ctx, scope, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, scope, key)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, scope, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NotificationService_CheckIdempotency_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckIdempotency'
type NotificationService_CheckIdempotency_Call struct {
	*mock.Call
}

// CheckIdempotency is a helper method to define mock.On call
//   - ctx context.Context
//   - scope string
//   - key string
func (_e *NotificationService_Expecter) CheckIdempotency(ctx interface{}, scope interface{}, key interface{}) *NotificationService_CheckIdempotency_Call {
	return &NotificationService_CheckIdempotency_Call{Call: _e.mock.On("CheckIdempotency", ctx, scope, key)}
}

func (_c *NotificationService_CheckIdempotency_Call) Run(run func(ctx context.Context, scope string, key string)) *NotificationService_CheckIdempotency_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *NotificationService_CheckIdempotency_Call) Return(_a0 string, _a1 error) *NotificationService_CheckIdempotency_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NotificationService_CheckIdempotency_Call) RunAndReturn(run func(context.Context, string, string) (string, error)) *NotificationService_CheckIdempotency_Call {
	_c.Call.Return(run)
	return _c
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

// InsertIdempotency provides a mock function with given fields: ctx, scope, key, notificationID
func (_m *NotificationService) InsertIdempotency(ctx context.Context, scope string, key string, notificationID string) error {
	ret := _m.Called(ctx, scope, key, notificationID)

	if len(ret) == 0 {
		panic("no return value specified for InsertIdempotency")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, scope, key, notificationID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NotificationService_InsertIdempotency_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InsertIdempotency'
type NotificationService_InsertIdempotency_Call struct {
	*mock.Call
}

// InsertIdempotency is a helper method to define mock.On call
//   - ctx context.Context
//   - scope string
//   - key string
//   - notificationID string
func (_e *NotificationService_Expecter) InsertIdempotency(ctx interface{}, scope interface{}, key interface{}, notificationID interface{}) *NotificationService_InsertIdempotency_Call {
	return &NotificationService_InsertIdempotency_Call{Call: _e.mock.On("InsertIdempotency", ctx, scope, key, notificationID)}
}

func (_c *NotificationService_InsertIdempotency_Call) Run(run func(ctx context.Context, scope string, key string, notificationID string)) *NotificationService_InsertIdempotency_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *NotificationService_InsertIdempotency_Call) Return(_a0 error) *NotificationService_InsertIdempotency_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *NotificationService_InsertIdempotency_Call) RunAndReturn(run func(context.Context, string, string, string) error) *NotificationService_InsertIdempotency_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, flt
func (_m *NotificationService) List(ctx context.Context, flt notification.Filter) ([]notification.Notification, error) {
	ret := _m.Called(ctx, flt)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []notification.Notification
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, notification.Filter) ([]notification.Notification, error)); ok {
		return rf(ctx, flt)
	}
	if rf, ok := ret.Get(0).(func(context.Context, notification.Filter) []notification.Notification); ok {
		r0 = rf(ctx, flt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]notification.Notification)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, notification.Filter) error); ok {
		r1 = rf(ctx, flt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NotificationService_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type NotificationService_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - flt notification.Filter
func (_e *NotificationService_Expecter) List(ctx interface{}, flt interface{}) *NotificationService_List_Call {
	return &NotificationService_List_Call{Call: _e.mock.On("List", ctx, flt)}
}

func (_c *NotificationService_List_Call) Run(run func(ctx context.Context, flt notification.Filter)) *NotificationService_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(notification.Filter))
	})
	return _c
}

func (_c *NotificationService_List_Call) Return(_a0 []notification.Notification, _a1 error) *NotificationService_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NotificationService_List_Call) RunAndReturn(run func(context.Context, notification.Filter) ([]notification.Notification, error)) *NotificationService_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListNotificationMessages provides a mock function with given fields: ctx, notificationID
func (_m *NotificationService) ListNotificationMessages(ctx context.Context, notificationID string) ([]notification.Message, error) {
	ret := _m.Called(ctx, notificationID)

	if len(ret) == 0 {
		panic("no return value specified for ListNotificationMessages")
	}

	var r0 []notification.Message
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]notification.Message, error)); ok {
		return rf(ctx, notificationID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []notification.Message); ok {
		r0 = rf(ctx, notificationID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]notification.Message)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, notificationID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NotificationService_ListNotificationMessages_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListNotificationMessages'
type NotificationService_ListNotificationMessages_Call struct {
	*mock.Call
}

// ListNotificationMessages is a helper method to define mock.On call
//   - ctx context.Context
//   - notificationID string
func (_e *NotificationService_Expecter) ListNotificationMessages(ctx interface{}, notificationID interface{}) *NotificationService_ListNotificationMessages_Call {
	return &NotificationService_ListNotificationMessages_Call{Call: _e.mock.On("ListNotificationMessages", ctx, notificationID)}
}

func (_c *NotificationService_ListNotificationMessages_Call) Run(run func(ctx context.Context, notificationID string)) *NotificationService_ListNotificationMessages_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *NotificationService_ListNotificationMessages_Call) Return(_a0 []notification.Message, _a1 error) *NotificationService_ListNotificationMessages_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NotificationService_ListNotificationMessages_Call) RunAndReturn(run func(context.Context, string) ([]notification.Message, error)) *NotificationService_ListNotificationMessages_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveIdempotencies provides a mock function with given fields: ctx, TTL
func (_m *NotificationService) RemoveIdempotencies(ctx context.Context, TTL time.Duration) error {
	ret := _m.Called(ctx, TTL)

	if len(ret) == 0 {
		panic("no return value specified for RemoveIdempotencies")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Duration) error); ok {
		r0 = rf(ctx, TTL)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NotificationService_RemoveIdempotencies_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveIdempotencies'
type NotificationService_RemoveIdempotencies_Call struct {
	*mock.Call
}

// RemoveIdempotencies is a helper method to define mock.On call
//   - ctx context.Context
//   - TTL time.Duration
func (_e *NotificationService_Expecter) RemoveIdempotencies(ctx interface{}, TTL interface{}) *NotificationService_RemoveIdempotencies_Call {
	return &NotificationService_RemoveIdempotencies_Call{Call: _e.mock.On("RemoveIdempotencies", ctx, TTL)}
}

func (_c *NotificationService_RemoveIdempotencies_Call) Run(run func(ctx context.Context, TTL time.Duration)) *NotificationService_RemoveIdempotencies_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(time.Duration))
	})
	return _c
}

func (_c *NotificationService_RemoveIdempotencies_Call) Return(_a0 error) *NotificationService_RemoveIdempotencies_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *NotificationService_RemoveIdempotencies_Call) RunAndReturn(run func(context.Context, time.Duration) error) *NotificationService_RemoveIdempotencies_Call {
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
