// Code generated by mockery v2.53.3. DO NOT EDIT.

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

// Create provides a mock function with given fields: ctx, rcv
func (_m *ReceiverService) Create(ctx context.Context, rcv *receiver.Receiver) error {
	ret := _m.Called(ctx, rcv)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *receiver.Receiver) error); ok {
		r0 = rf(ctx, rcv)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReceiverService_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type ReceiverService_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - rcv *receiver.Receiver
func (_e *ReceiverService_Expecter) Create(ctx interface{}, rcv interface{}) *ReceiverService_Create_Call {
	return &ReceiverService_Create_Call{Call: _e.mock.On("Create", ctx, rcv)}
}

func (_c *ReceiverService_Create_Call) Run(run func(ctx context.Context, rcv *receiver.Receiver)) *ReceiverService_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*receiver.Receiver))
	})
	return _c
}

func (_c *ReceiverService_Create_Call) Return(_a0 error) *ReceiverService_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ReceiverService_Create_Call) RunAndReturn(run func(context.Context, *receiver.Receiver) error) *ReceiverService_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, id
func (_m *ReceiverService) Delete(ctx context.Context, id uint64) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReceiverService_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type ReceiverService_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint64
func (_e *ReceiverService_Expecter) Delete(ctx interface{}, id interface{}) *ReceiverService_Delete_Call {
	return &ReceiverService_Delete_Call{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *ReceiverService_Delete_Call) Run(run func(ctx context.Context, id uint64)) *ReceiverService_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *ReceiverService_Delete_Call) Return(_a0 error) *ReceiverService_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ReceiverService_Delete_Call) RunAndReturn(run func(context.Context, uint64) error) *ReceiverService_Delete_Call {
	_c.Call.Return(run)
	return _c
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

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *receiver.Receiver
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, ...receiver.GetOption) (*receiver.Receiver, error)); ok {
		return rf(ctx, id, gopts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, ...receiver.GetOption) *receiver.Receiver); ok {
		r0 = rf(ctx, id, gopts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*receiver.Receiver)
		}
	}

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

func (_c *ReceiverService_Get_Call) RunAndReturn(run func(context.Context, uint64, ...receiver.GetOption) (*receiver.Receiver, error)) *ReceiverService_Get_Call {
	_c.Call.Return(run)
	return _c
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

// Update provides a mock function with given fields: ctx, rcv
func (_m *ReceiverService) Update(ctx context.Context, rcv *receiver.Receiver) error {
	ret := _m.Called(ctx, rcv)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *receiver.Receiver) error); ok {
		r0 = rf(ctx, rcv)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReceiverService_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type ReceiverService_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - rcv *receiver.Receiver
func (_e *ReceiverService_Expecter) Update(ctx interface{}, rcv interface{}) *ReceiverService_Update_Call {
	return &ReceiverService_Update_Call{Call: _e.mock.On("Update", ctx, rcv)}
}

func (_c *ReceiverService_Update_Call) Run(run func(ctx context.Context, rcv *receiver.Receiver)) *ReceiverService_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*receiver.Receiver))
	})
	return _c
}

func (_c *ReceiverService_Update_Call) Return(_a0 error) *ReceiverService_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ReceiverService_Update_Call) RunAndReturn(run func(context.Context, *receiver.Receiver) error) *ReceiverService_Update_Call {
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
