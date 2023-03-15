// Code generated by mockery v2.21.3. DO NOT EDIT.

package mocks

import (
	context "context"

	log "github.com/goto/siren/core/log"
	mock "github.com/stretchr/testify/mock"
)

// NotificationLogRepository is an autogenerated mock type for the NotificationLogRepository type
type NotificationLogRepository struct {
	mock.Mock
}

type NotificationLogRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *NotificationLogRepository) EXPECT() *NotificationLogRepository_Expecter {
	return &NotificationLogRepository_Expecter{mock: &_m.Mock}
}

// BulkCreate provides a mock function with given fields: _a0, _a1
func (_m *NotificationLogRepository) BulkCreate(_a0 context.Context, _a1 []log.Notification) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []log.Notification) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NotificationLogRepository_BulkCreate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BulkCreate'
type NotificationLogRepository_BulkCreate_Call struct {
	*mock.Call
}

// BulkCreate is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 []log.Notification
func (_e *NotificationLogRepository_Expecter) BulkCreate(_a0 interface{}, _a1 interface{}) *NotificationLogRepository_BulkCreate_Call {
	return &NotificationLogRepository_BulkCreate_Call{Call: _e.mock.On("BulkCreate", _a0, _a1)}
}

func (_c *NotificationLogRepository_BulkCreate_Call) Run(run func(_a0 context.Context, _a1 []log.Notification)) *NotificationLogRepository_BulkCreate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]log.Notification))
	})
	return _c
}

func (_c *NotificationLogRepository_BulkCreate_Call) Return(_a0 error) *NotificationLogRepository_BulkCreate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *NotificationLogRepository_BulkCreate_Call) RunAndReturn(run func(context.Context, []log.Notification) error) *NotificationLogRepository_BulkCreate_Call {
	_c.Call.Return(run)
	return _c
}

// ListAlertIDsBySilenceID provides a mock function with given fields: _a0, _a1
func (_m *NotificationLogRepository) ListAlertIDsBySilenceID(_a0 context.Context, _a1 string) ([]int64, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]int64, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []int64); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NotificationLogRepository_ListAlertIDsBySilenceID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAlertIDsBySilenceID'
type NotificationLogRepository_ListAlertIDsBySilenceID_Call struct {
	*mock.Call
}

// ListAlertIDsBySilenceID is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *NotificationLogRepository_Expecter) ListAlertIDsBySilenceID(_a0 interface{}, _a1 interface{}) *NotificationLogRepository_ListAlertIDsBySilenceID_Call {
	return &NotificationLogRepository_ListAlertIDsBySilenceID_Call{Call: _e.mock.On("ListAlertIDsBySilenceID", _a0, _a1)}
}

func (_c *NotificationLogRepository_ListAlertIDsBySilenceID_Call) Run(run func(_a0 context.Context, _a1 string)) *NotificationLogRepository_ListAlertIDsBySilenceID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *NotificationLogRepository_ListAlertIDsBySilenceID_Call) Return(_a0 []int64, _a1 error) *NotificationLogRepository_ListAlertIDsBySilenceID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NotificationLogRepository_ListAlertIDsBySilenceID_Call) RunAndReturn(run func(context.Context, string) ([]int64, error)) *NotificationLogRepository_ListAlertIDsBySilenceID_Call {
	_c.Call.Return(run)
	return _c
}

// ListSubscriptionIDsBySilenceID provides a mock function with given fields: _a0, _a1
func (_m *NotificationLogRepository) ListSubscriptionIDsBySilenceID(_a0 context.Context, _a1 string) ([]int64, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]int64, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []int64); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NotificationLogRepository_ListSubscriptionIDsBySilenceID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSubscriptionIDsBySilenceID'
type NotificationLogRepository_ListSubscriptionIDsBySilenceID_Call struct {
	*mock.Call
}

// ListSubscriptionIDsBySilenceID is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *NotificationLogRepository_Expecter) ListSubscriptionIDsBySilenceID(_a0 interface{}, _a1 interface{}) *NotificationLogRepository_ListSubscriptionIDsBySilenceID_Call {
	return &NotificationLogRepository_ListSubscriptionIDsBySilenceID_Call{Call: _e.mock.On("ListSubscriptionIDsBySilenceID", _a0, _a1)}
}

func (_c *NotificationLogRepository_ListSubscriptionIDsBySilenceID_Call) Run(run func(_a0 context.Context, _a1 string)) *NotificationLogRepository_ListSubscriptionIDsBySilenceID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *NotificationLogRepository_ListSubscriptionIDsBySilenceID_Call) Return(_a0 []int64, _a1 error) *NotificationLogRepository_ListSubscriptionIDsBySilenceID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NotificationLogRepository_ListSubscriptionIDsBySilenceID_Call) RunAndReturn(run func(context.Context, string) ([]int64, error)) *NotificationLogRepository_ListSubscriptionIDsBySilenceID_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewNotificationLogRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewNotificationLogRepository creates a new instance of NotificationLogRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNotificationLogRepository(t mockConstructorTestingTNewNotificationLogRepository) *NotificationLogRepository {
	mock := &NotificationLogRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
