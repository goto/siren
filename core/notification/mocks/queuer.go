// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"

	notification "github.com/goto/siren/core/notification"
	mock "github.com/stretchr/testify/mock"

	queues "github.com/goto/siren/plugins/queues"
)

// Queuer is an autogenerated mock type for the Queuer type
type Queuer struct {
	mock.Mock
}

type Queuer_Expecter struct {
	mock *mock.Mock
}

func (_m *Queuer) EXPECT() *Queuer_Expecter {
	return &Queuer_Expecter{mock: &_m.Mock}
}

// Cleanup provides a mock function with given fields: ctx, filter
func (_m *Queuer) Cleanup(ctx context.Context, filter queues.FilterCleanup) error {
	ret := _m.Called(ctx, filter)

	if len(ret) == 0 {
		panic("no return value specified for Cleanup")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, queues.FilterCleanup) error); ok {
		r0 = rf(ctx, filter)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Queuer_Cleanup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Cleanup'
type Queuer_Cleanup_Call struct {
	*mock.Call
}

// Cleanup is a helper method to define mock.On call
//   - ctx context.Context
//   - filter queues.FilterCleanup
func (_e *Queuer_Expecter) Cleanup(ctx interface{}, filter interface{}) *Queuer_Cleanup_Call {
	return &Queuer_Cleanup_Call{Call: _e.mock.On("Cleanup", ctx, filter)}
}

func (_c *Queuer_Cleanup_Call) Run(run func(ctx context.Context, filter queues.FilterCleanup)) *Queuer_Cleanup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(queues.FilterCleanup))
	})
	return _c
}

func (_c *Queuer_Cleanup_Call) Return(_a0 error) *Queuer_Cleanup_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Queuer_Cleanup_Call) RunAndReturn(run func(context.Context, queues.FilterCleanup) error) *Queuer_Cleanup_Call {
	_c.Call.Return(run)
	return _c
}

// Dequeue provides a mock function with given fields: ctx, receiverTypes, batchSize, handlerFn
func (_m *Queuer) Dequeue(ctx context.Context, receiverTypes []string, batchSize int, handlerFn func(context.Context, []notification.Message) error) error {
	ret := _m.Called(ctx, receiverTypes, batchSize, handlerFn)

	if len(ret) == 0 {
		panic("no return value specified for Dequeue")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []string, int, func(context.Context, []notification.Message) error) error); ok {
		r0 = rf(ctx, receiverTypes, batchSize, handlerFn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Queuer_Dequeue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Dequeue'
type Queuer_Dequeue_Call struct {
	*mock.Call
}

// Dequeue is a helper method to define mock.On call
//   - ctx context.Context
//   - receiverTypes []string
//   - batchSize int
//   - handlerFn func(context.Context , []notification.Message) error
func (_e *Queuer_Expecter) Dequeue(ctx interface{}, receiverTypes interface{}, batchSize interface{}, handlerFn interface{}) *Queuer_Dequeue_Call {
	return &Queuer_Dequeue_Call{Call: _e.mock.On("Dequeue", ctx, receiverTypes, batchSize, handlerFn)}
}

func (_c *Queuer_Dequeue_Call) Run(run func(ctx context.Context, receiverTypes []string, batchSize int, handlerFn func(context.Context, []notification.Message) error)) *Queuer_Dequeue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string), args[2].(int), args[3].(func(context.Context, []notification.Message) error))
	})
	return _c
}

func (_c *Queuer_Dequeue_Call) Return(_a0 error) *Queuer_Dequeue_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Queuer_Dequeue_Call) RunAndReturn(run func(context.Context, []string, int, func(context.Context, []notification.Message) error) error) *Queuer_Dequeue_Call {
	_c.Call.Return(run)
	return _c
}

// Enqueue provides a mock function with given fields: ctx, ms
func (_m *Queuer) Enqueue(ctx context.Context, ms ...notification.Message) error {
	_va := make([]interface{}, len(ms))
	for _i := range ms {
		_va[_i] = ms[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Enqueue")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...notification.Message) error); ok {
		r0 = rf(ctx, ms...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Queuer_Enqueue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Enqueue'
type Queuer_Enqueue_Call struct {
	*mock.Call
}

// Enqueue is a helper method to define mock.On call
//   - ctx context.Context
//   - ms ...notification.Message
func (_e *Queuer_Expecter) Enqueue(ctx interface{}, ms ...interface{}) *Queuer_Enqueue_Call {
	return &Queuer_Enqueue_Call{Call: _e.mock.On("Enqueue",
		append([]interface{}{ctx}, ms...)...)}
}

func (_c *Queuer_Enqueue_Call) Run(run func(ctx context.Context, ms ...notification.Message)) *Queuer_Enqueue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]notification.Message, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(notification.Message)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *Queuer_Enqueue_Call) Return(_a0 error) *Queuer_Enqueue_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Queuer_Enqueue_Call) RunAndReturn(run func(context.Context, ...notification.Message) error) *Queuer_Enqueue_Call {
	_c.Call.Return(run)
	return _c
}

// ErrorCallback provides a mock function with given fields: ctx, ms
func (_m *Queuer) ErrorCallback(ctx context.Context, ms notification.Message) error {
	ret := _m.Called(ctx, ms)

	if len(ret) == 0 {
		panic("no return value specified for ErrorCallback")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, notification.Message) error); ok {
		r0 = rf(ctx, ms)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Queuer_ErrorCallback_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ErrorCallback'
type Queuer_ErrorCallback_Call struct {
	*mock.Call
}

// ErrorCallback is a helper method to define mock.On call
//   - ctx context.Context
//   - ms notification.Message
func (_e *Queuer_Expecter) ErrorCallback(ctx interface{}, ms interface{}) *Queuer_ErrorCallback_Call {
	return &Queuer_ErrorCallback_Call{Call: _e.mock.On("ErrorCallback", ctx, ms)}
}

func (_c *Queuer_ErrorCallback_Call) Run(run func(ctx context.Context, ms notification.Message)) *Queuer_ErrorCallback_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(notification.Message))
	})
	return _c
}

func (_c *Queuer_ErrorCallback_Call) Return(_a0 error) *Queuer_ErrorCallback_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Queuer_ErrorCallback_Call) RunAndReturn(run func(context.Context, notification.Message) error) *Queuer_ErrorCallback_Call {
	_c.Call.Return(run)
	return _c
}

// ListMessages provides a mock function with given fields: ctx, notificationID
func (_m *Queuer) ListMessages(ctx context.Context, notificationID string) ([]notification.Message, error) {
	ret := _m.Called(ctx, notificationID)

	if len(ret) == 0 {
		panic("no return value specified for ListMessages")
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

// Queuer_ListMessages_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListMessages'
type Queuer_ListMessages_Call struct {
	*mock.Call
}

// ListMessages is a helper method to define mock.On call
//   - ctx context.Context
//   - notificationID string
func (_e *Queuer_Expecter) ListMessages(ctx interface{}, notificationID interface{}) *Queuer_ListMessages_Call {
	return &Queuer_ListMessages_Call{Call: _e.mock.On("ListMessages", ctx, notificationID)}
}

func (_c *Queuer_ListMessages_Call) Run(run func(ctx context.Context, notificationID string)) *Queuer_ListMessages_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Queuer_ListMessages_Call) Return(_a0 []notification.Message, _a1 error) *Queuer_ListMessages_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Queuer_ListMessages_Call) RunAndReturn(run func(context.Context, string) ([]notification.Message, error)) *Queuer_ListMessages_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields: ctx
func (_m *Queuer) Stop(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Stop")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Queuer_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type Queuer_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Queuer_Expecter) Stop(ctx interface{}) *Queuer_Stop_Call {
	return &Queuer_Stop_Call{Call: _e.mock.On("Stop", ctx)}
}

func (_c *Queuer_Stop_Call) Run(run func(ctx context.Context)) *Queuer_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Queuer_Stop_Call) Return(_a0 error) *Queuer_Stop_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Queuer_Stop_Call) RunAndReturn(run func(context.Context) error) *Queuer_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// SuccessCallback provides a mock function with given fields: ctx, ms
func (_m *Queuer) SuccessCallback(ctx context.Context, ms notification.Message) error {
	ret := _m.Called(ctx, ms)

	if len(ret) == 0 {
		panic("no return value specified for SuccessCallback")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, notification.Message) error); ok {
		r0 = rf(ctx, ms)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Queuer_SuccessCallback_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SuccessCallback'
type Queuer_SuccessCallback_Call struct {
	*mock.Call
}

// SuccessCallback is a helper method to define mock.On call
//   - ctx context.Context
//   - ms notification.Message
func (_e *Queuer_Expecter) SuccessCallback(ctx interface{}, ms interface{}) *Queuer_SuccessCallback_Call {
	return &Queuer_SuccessCallback_Call{Call: _e.mock.On("SuccessCallback", ctx, ms)}
}

func (_c *Queuer_SuccessCallback_Call) Run(run func(ctx context.Context, ms notification.Message)) *Queuer_SuccessCallback_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(notification.Message))
	})
	return _c
}

func (_c *Queuer_SuccessCallback_Call) Return(_a0 error) *Queuer_SuccessCallback_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Queuer_SuccessCallback_Call) RunAndReturn(run func(context.Context, notification.Message) error) *Queuer_SuccessCallback_Call {
	_c.Call.Return(run)
	return _c
}

// NewQueuer creates a new instance of Queuer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQueuer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Queuer {
	mock := &Queuer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
