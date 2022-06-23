// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	context "context"

	provider "github.com/odpf/siren/core/provider"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// ProviderRepository is an autogenerated mock type for the Repository type
type ProviderRepository struct {
	mock.Mock
}

type ProviderRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *ProviderRepository) EXPECT() *ProviderRepository_Expecter {
	return &ProviderRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *ProviderRepository) Create(_a0 context.Context, _a1 *provider.Provider) (uint64, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, *provider.Provider) uint64); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *provider.Provider) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProviderRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type ProviderRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *provider.Provider
func (_e *ProviderRepository_Expecter) Create(_a0 interface{}, _a1 interface{}) *ProviderRepository_Create_Call {
	return &ProviderRepository_Create_Call{Call: _e.mock.On("Create", _a0, _a1)}
}

func (_c *ProviderRepository_Create_Call) Run(run func(_a0 context.Context, _a1 *provider.Provider)) *ProviderRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*provider.Provider))
	})
	return _c
}

func (_c *ProviderRepository_Create_Call) Return(_a0 uint64, _a1 error) *ProviderRepository_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *ProviderRepository) Delete(_a0 context.Context, _a1 uint64) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProviderRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type ProviderRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 uint64
func (_e *ProviderRepository_Expecter) Delete(_a0 interface{}, _a1 interface{}) *ProviderRepository_Delete_Call {
	return &ProviderRepository_Delete_Call{Call: _e.mock.On("Delete", _a0, _a1)}
}

func (_c *ProviderRepository_Delete_Call) Run(run func(_a0 context.Context, _a1 uint64)) *ProviderRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *ProviderRepository_Delete_Call) Return(_a0 error) *ProviderRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *ProviderRepository) Get(_a0 context.Context, _a1 uint64) (*provider.Provider, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *provider.Provider
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *provider.Provider); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*provider.Provider)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProviderRepository_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type ProviderRepository_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 uint64
func (_e *ProviderRepository_Expecter) Get(_a0 interface{}, _a1 interface{}) *ProviderRepository_Get_Call {
	return &ProviderRepository_Get_Call{Call: _e.mock.On("Get", _a0, _a1)}
}

func (_c *ProviderRepository_Get_Call) Run(run func(_a0 context.Context, _a1 uint64)) *ProviderRepository_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *ProviderRepository_Get_Call) Return(_a0 *provider.Provider, _a1 error) *ProviderRepository_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// List provides a mock function with given fields: _a0, _a1
func (_m *ProviderRepository) List(_a0 context.Context, _a1 provider.Filter) ([]*provider.Provider, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*provider.Provider
	if rf, ok := ret.Get(0).(func(context.Context, provider.Filter) []*provider.Provider); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*provider.Provider)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, provider.Filter) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProviderRepository_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type ProviderRepository_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 provider.Filter
func (_e *ProviderRepository_Expecter) List(_a0 interface{}, _a1 interface{}) *ProviderRepository_List_Call {
	return &ProviderRepository_List_Call{Call: _e.mock.On("List", _a0, _a1)}
}

func (_c *ProviderRepository_List_Call) Run(run func(_a0 context.Context, _a1 provider.Filter)) *ProviderRepository_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(provider.Filter))
	})
	return _c
}

func (_c *ProviderRepository_List_Call) Return(_a0 []*provider.Provider, _a1 error) *ProviderRepository_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *ProviderRepository) Update(_a0 context.Context, _a1 *provider.Provider) (uint64, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, *provider.Provider) uint64); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *provider.Provider) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProviderRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type ProviderRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *provider.Provider
func (_e *ProviderRepository_Expecter) Update(_a0 interface{}, _a1 interface{}) *ProviderRepository_Update_Call {
	return &ProviderRepository_Update_Call{Call: _e.mock.On("Update", _a0, _a1)}
}

func (_c *ProviderRepository_Update_Call) Run(run func(_a0 context.Context, _a1 *provider.Provider)) *ProviderRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*provider.Provider))
	})
	return _c
}

func (_c *ProviderRepository_Update_Call) Return(_a0 uint64, _a1 error) *ProviderRepository_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// NewProviderRepository creates a new instance of ProviderRepository. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewProviderRepository(t testing.TB) *ProviderRepository {
	mock := &ProviderRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
