// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	receiver "github.com/goto/siren/core/receiver"
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

// Get provides a mock function with given fields: ctx, id, gopts
func (_m *ReceiverService) Get(ctx context.Context, id uint64, gopts ...receiver.GetOption) (*receiver.Receiver, error) {
	_va := make([]interface{}, len(gopts))
	for _i := range gopts {
		_va[_i] = gopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, id)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *receiver.Receiver
	if rf, ok := ret.Get(0).(func(context.Context, uint64, ...receiver.GetOption) *receiver.Receiver); ok {
		r0 = rf(ctx, id, gopts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*receiver.Receiver)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, ...receiver.GetOption) error); ok {
		r1 = rf(ctx, id, gopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReceiverService_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type ReceiverService_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint64
//   - gopts ...receiver.GetOption
func (_e *ReceiverService_Expecter) Get(ctx interface{}, id interface{}, gopts ...interface{}) *ReceiverService_Get_Call {
	return &ReceiverService_Get_Call{Call: _e.mock.On("Get",
		append([]interface{}{ctx, id}, gopts...)...)}
}

func (_c *ReceiverService_Get_Call) Run(run func(ctx context.Context, id uint64, gopts ...receiver.GetOption)) *ReceiverService_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]receiver.GetOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(receiver.GetOption)
			}
		}
		run(args[0].(context.Context), args[1].(uint64), variadicArgs...)
	})
	return _c
}

func (_c *ReceiverService_Get_Call) Return(_a0 *receiver.Receiver, _a1 error) *ReceiverService_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewReceiverService interface {
	mock.TestingT
	Cleanup(func())
}

// NewReceiverService creates a new instance of ReceiverService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReceiverService(t mockConstructorTestingTNewReceiverService) *ReceiverService {
	mock := &ReceiverService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
