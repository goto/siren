// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	receiver "github.com/goto/siren/core/receiver"
	mock "github.com/stretchr/testify/mock"
)

// ReceiverService is an autogenerated mock type for the ReceiverService type
type ReceiverService struct {
	mock.Mock
}

type ReceiverService_Expecter struct {
	mock *mock.Mock
}

func (_m *ReceiverService) EXPECT() *ReceiverService_Expecter {
	return &ReceiverService_Expecter{mock: &_m.Mock}
}

// List provides a mock function with given fields: ctx, flt
func (_m *ReceiverService) List(ctx context.Context, flt receiver.Filter) ([]receiver.Receiver, error) {
	ret := _m.Called(ctx, flt)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []receiver.Receiver
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, receiver.Filter) ([]receiver.Receiver, error)); ok {
		return rf(ctx, flt)
	}
	if rf, ok := ret.Get(0).(func(context.Context, receiver.Filter) []receiver.Receiver); ok {
		r0 = rf(ctx, flt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]receiver.Receiver)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, receiver.Filter) error); ok {
		r1 = rf(ctx, flt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReceiverService_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type ReceiverService_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - flt receiver.Filter
func (_e *ReceiverService_Expecter) List(ctx interface{}, flt interface{}) *ReceiverService_List_Call {
	return &ReceiverService_List_Call{Call: _e.mock.On("List", ctx, flt)}
}

func (_c *ReceiverService_List_Call) Run(run func(ctx context.Context, flt receiver.Filter)) *ReceiverService_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(receiver.Filter))
	})
	return _c
}

func (_c *ReceiverService_List_Call) Return(_a0 []receiver.Receiver, _a1 error) *ReceiverService_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ReceiverService_List_Call) RunAndReturn(run func(context.Context, receiver.Filter) ([]receiver.Receiver, error)) *ReceiverService_List_Call {
	_c.Call.Return(run)
	return _c
}

// NewReceiverService creates a new instance of ReceiverService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewReceiverService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ReceiverService {
	mock := &ReceiverService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
