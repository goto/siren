// Code generated by mockery v2.33.2. DO NOT EDIT.

package mocks

import (
	context "context"

	log "github.com/goto/siren/core/log"
	mock "github.com/stretchr/testify/mock"
)

// LogService is an autogenerated mock type for the LogService type
type LogService struct {
	mock.Mock
}

type LogService_Expecter struct {
	mock *mock.Mock
}

func (_m *LogService) EXPECT() *LogService_Expecter {
	return &LogService_Expecter{mock: &_m.Mock}
}

// LogNotifications provides a mock function with given fields: ctx, nlogs
func (_m *LogService) LogNotifications(ctx context.Context, nlogs ...log.Notification) error {
	_va := make([]interface{}, len(nlogs))
	for _i := range nlogs {
		_va[_i] = nlogs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...log.Notification) error); ok {
		r0 = rf(ctx, nlogs...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LogService_LogNotifications_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LogNotifications'
type LogService_LogNotifications_Call struct {
	*mock.Call
}

// LogNotifications is a helper method to define mock.On call
//   - ctx context.Context
//   - nlogs ...log.Notification
func (_e *LogService_Expecter) LogNotifications(ctx interface{}, nlogs ...interface{}) *LogService_LogNotifications_Call {
	return &LogService_LogNotifications_Call{Call: _e.mock.On("LogNotifications",
		append([]interface{}{ctx}, nlogs...)...)}
}

func (_c *LogService_LogNotifications_Call) Run(run func(ctx context.Context, nlogs ...log.Notification)) *LogService_LogNotifications_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]log.Notification, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(log.Notification)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *LogService_LogNotifications_Call) Return(_a0 error) *LogService_LogNotifications_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *LogService_LogNotifications_Call) RunAndReturn(run func(context.Context, ...log.Notification) error) *LogService_LogNotifications_Call {
	_c.Call.Return(run)
	return _c
}

// NewLogService creates a new instance of LogService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLogService(t interface {
	mock.TestingT
	Cleanup(func())
}) *LogService {
	mock := &LogService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
