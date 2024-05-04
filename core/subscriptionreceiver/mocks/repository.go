// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	subscriptionreceiver "github.com/goto/siren/core/subscriptionreceiver"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

type Repository_Expecter struct {
	mock *mock.Mock
}

func (_m *Repository) EXPECT() *Repository_Expecter {
	return &Repository_Expecter{mock: &_m.Mock}
}

// BulkCreate provides a mock function with given fields: _a0, _a1
func (_m *Repository) BulkCreate(_a0 context.Context, _a1 []subscriptionreceiver.Relation) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for BulkCreate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []subscriptionreceiver.Relation) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_BulkCreate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BulkCreate'
type Repository_BulkCreate_Call struct {
	*mock.Call
}

// BulkCreate is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 []subscriptionreceiver.Relation
func (_e *Repository_Expecter) BulkCreate(_a0 interface{}, _a1 interface{}) *Repository_BulkCreate_Call {
	return &Repository_BulkCreate_Call{Call: _e.mock.On("BulkCreate", _a0, _a1)}
}

func (_c *Repository_BulkCreate_Call) Run(run func(_a0 context.Context, _a1 []subscriptionreceiver.Relation)) *Repository_BulkCreate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]subscriptionreceiver.Relation))
	})
	return _c
}

func (_c *Repository_BulkCreate_Call) Return(_a0 error) *Repository_BulkCreate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_BulkCreate_Call) RunAndReturn(run func(context.Context, []subscriptionreceiver.Relation) error) *Repository_BulkCreate_Call {
	_c.Call.Return(run)
	return _c
}

// BulkDelete provides a mock function with given fields: ctx, flt
func (_m *Repository) BulkDelete(ctx context.Context, flt subscriptionreceiver.DeleteFilter) error {
	ret := _m.Called(ctx, flt)

	if len(ret) == 0 {
		panic("no return value specified for BulkDelete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, subscriptionreceiver.DeleteFilter) error); ok {
		r0 = rf(ctx, flt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_BulkDelete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BulkDelete'
type Repository_BulkDelete_Call struct {
	*mock.Call
}

// BulkDelete is a helper method to define mock.On call
//   - ctx context.Context
//   - flt subscriptionreceiver.DeleteFilter
func (_e *Repository_Expecter) BulkDelete(ctx interface{}, flt interface{}) *Repository_BulkDelete_Call {
	return &Repository_BulkDelete_Call{Call: _e.mock.On("BulkDelete", ctx, flt)}
}

func (_c *Repository_BulkDelete_Call) Run(run func(ctx context.Context, flt subscriptionreceiver.DeleteFilter)) *Repository_BulkDelete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(subscriptionreceiver.DeleteFilter))
	})
	return _c
}

func (_c *Repository_BulkDelete_Call) Return(_a0 error) *Repository_BulkDelete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_BulkDelete_Call) RunAndReturn(run func(context.Context, subscriptionreceiver.DeleteFilter) error) *Repository_BulkDelete_Call {
	_c.Call.Return(run)
	return _c
}

// BulkSoftDelete provides a mock function with given fields: ctx, flt
func (_m *Repository) BulkSoftDelete(ctx context.Context, flt subscriptionreceiver.DeleteFilter) error {
	ret := _m.Called(ctx, flt)

	if len(ret) == 0 {
		panic("no return value specified for BulkSoftDelete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, subscriptionreceiver.DeleteFilter) error); ok {
		r0 = rf(ctx, flt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_BulkSoftDelete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BulkSoftDelete'
type Repository_BulkSoftDelete_Call struct {
	*mock.Call
}

// BulkSoftDelete is a helper method to define mock.On call
//   - ctx context.Context
//   - flt subscriptionreceiver.DeleteFilter
func (_e *Repository_Expecter) BulkSoftDelete(ctx interface{}, flt interface{}) *Repository_BulkSoftDelete_Call {
	return &Repository_BulkSoftDelete_Call{Call: _e.mock.On("BulkSoftDelete", ctx, flt)}
}

func (_c *Repository_BulkSoftDelete_Call) Run(run func(ctx context.Context, flt subscriptionreceiver.DeleteFilter)) *Repository_BulkSoftDelete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(subscriptionreceiver.DeleteFilter))
	})
	return _c
}

func (_c *Repository_BulkSoftDelete_Call) Return(_a0 error) *Repository_BulkSoftDelete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_BulkSoftDelete_Call) RunAndReturn(run func(context.Context, subscriptionreceiver.DeleteFilter) error) *Repository_BulkSoftDelete_Call {
	_c.Call.Return(run)
	return _c
}

// BulkUpsert provides a mock function with given fields: _a0, _a1
func (_m *Repository) BulkUpsert(_a0 context.Context, _a1 []subscriptionreceiver.Relation) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for BulkUpsert")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []subscriptionreceiver.Relation) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_BulkUpsert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BulkUpsert'
type Repository_BulkUpsert_Call struct {
	*mock.Call
}

// BulkUpsert is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 []subscriptionreceiver.Relation
func (_e *Repository_Expecter) BulkUpsert(_a0 interface{}, _a1 interface{}) *Repository_BulkUpsert_Call {
	return &Repository_BulkUpsert_Call{Call: _e.mock.On("BulkUpsert", _a0, _a1)}
}

func (_c *Repository_BulkUpsert_Call) Run(run func(_a0 context.Context, _a1 []subscriptionreceiver.Relation)) *Repository_BulkUpsert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]subscriptionreceiver.Relation))
	})
	return _c
}

func (_c *Repository_BulkUpsert_Call) Return(_a0 error) *Repository_BulkUpsert_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_BulkUpsert_Call) RunAndReturn(run func(context.Context, []subscriptionreceiver.Relation) error) *Repository_BulkUpsert_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: _a0, _a1
func (_m *Repository) List(_a0 context.Context, _a1 subscriptionreceiver.Filter) ([]subscriptionreceiver.Relation, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []subscriptionreceiver.Relation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, subscriptionreceiver.Filter) ([]subscriptionreceiver.Relation, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, subscriptionreceiver.Filter) []subscriptionreceiver.Relation); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]subscriptionreceiver.Relation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, subscriptionreceiver.Filter) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type Repository_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 subscriptionreceiver.Filter
func (_e *Repository_Expecter) List(_a0 interface{}, _a1 interface{}) *Repository_List_Call {
	return &Repository_List_Call{Call: _e.mock.On("List", _a0, _a1)}
}

func (_c *Repository_List_Call) Run(run func(_a0 context.Context, _a1 subscriptionreceiver.Filter)) *Repository_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(subscriptionreceiver.Filter))
	})
	return _c
}

func (_c *Repository_List_Call) Return(_a0 []subscriptionreceiver.Relation, _a1 error) *Repository_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_List_Call) RunAndReturn(run func(context.Context, subscriptionreceiver.Filter) ([]subscriptionreceiver.Relation, error)) *Repository_List_Call {
	_c.Call.Return(run)
	return _c
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
