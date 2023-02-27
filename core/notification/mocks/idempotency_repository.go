// Code generated by mockery v2.21.3. DO NOT EDIT.

package mocks

import (
	context "context"

	notification "github.com/goto/siren/core/notification"
	mock "github.com/stretchr/testify/mock"
)

// IdempotencyRepository is an autogenerated mock type for the IdempotencyRepository type
type IdempotencyRepository struct {
	mock.Mock
}

type IdempotencyRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *IdempotencyRepository) EXPECT() *IdempotencyRepository_Expecter {
	return &IdempotencyRepository_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *IdempotencyRepository) Delete(_a0 context.Context, _a1 notification.IdempotencyFilter) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, notification.IdempotencyFilter) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IdempotencyRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type IdempotencyRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 notification.IdempotencyFilter
func (_e *IdempotencyRepository_Expecter) Delete(_a0 interface{}, _a1 interface{}) *IdempotencyRepository_Delete_Call {
	return &IdempotencyRepository_Delete_Call{Call: _e.mock.On("Delete", _a0, _a1)}
}

func (_c *IdempotencyRepository_Delete_Call) Run(run func(_a0 context.Context, _a1 notification.IdempotencyFilter)) *IdempotencyRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(notification.IdempotencyFilter))
	})
	return _c
}

func (_c *IdempotencyRepository_Delete_Call) Return(_a0 error) *IdempotencyRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IdempotencyRepository_Delete_Call) RunAndReturn(run func(context.Context, notification.IdempotencyFilter) error) *IdempotencyRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// InsertOnConflictReturning provides a mock function with given fields: _a0, _a1, _a2
func (_m *IdempotencyRepository) InsertOnConflictReturning(_a0 context.Context, _a1 string, _a2 string) (*notification.Idempotency, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *notification.Idempotency
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*notification.Idempotency, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *notification.Idempotency); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*notification.Idempotency)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IdempotencyRepository_InsertOnConflictReturning_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InsertOnConflictReturning'
type IdempotencyRepository_InsertOnConflictReturning_Call struct {
	*mock.Call
}

// InsertOnConflictReturning is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
//   - _a2 string
func (_e *IdempotencyRepository_Expecter) InsertOnConflictReturning(_a0 interface{}, _a1 interface{}, _a2 interface{}) *IdempotencyRepository_InsertOnConflictReturning_Call {
	return &IdempotencyRepository_InsertOnConflictReturning_Call{Call: _e.mock.On("InsertOnConflictReturning", _a0, _a1, _a2)}
}

func (_c *IdempotencyRepository_InsertOnConflictReturning_Call) Run(run func(_a0 context.Context, _a1 string, _a2 string)) *IdempotencyRepository_InsertOnConflictReturning_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *IdempotencyRepository_InsertOnConflictReturning_Call) Return(_a0 *notification.Idempotency, _a1 error) *IdempotencyRepository_InsertOnConflictReturning_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IdempotencyRepository_InsertOnConflictReturning_Call) RunAndReturn(run func(context.Context, string, string) (*notification.Idempotency, error)) *IdempotencyRepository_InsertOnConflictReturning_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateSuccess provides a mock function with given fields: _a0, _a1, _a2
func (_m *IdempotencyRepository) UpdateSuccess(_a0 context.Context, _a1 uint64, _a2 bool) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, bool) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IdempotencyRepository_UpdateSuccess_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateSuccess'
type IdempotencyRepository_UpdateSuccess_Call struct {
	*mock.Call
}

// UpdateSuccess is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 uint64
//   - _a2 bool
func (_e *IdempotencyRepository_Expecter) UpdateSuccess(_a0 interface{}, _a1 interface{}, _a2 interface{}) *IdempotencyRepository_UpdateSuccess_Call {
	return &IdempotencyRepository_UpdateSuccess_Call{Call: _e.mock.On("UpdateSuccess", _a0, _a1, _a2)}
}

func (_c *IdempotencyRepository_UpdateSuccess_Call) Run(run func(_a0 context.Context, _a1 uint64, _a2 bool)) *IdempotencyRepository_UpdateSuccess_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(bool))
	})
	return _c
}

func (_c *IdempotencyRepository_UpdateSuccess_Call) Return(_a0 error) *IdempotencyRepository_UpdateSuccess_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IdempotencyRepository_UpdateSuccess_Call) RunAndReturn(run func(context.Context, uint64, bool) error) *IdempotencyRepository_UpdateSuccess_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewIdempotencyRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewIdempotencyRepository creates a new instance of IdempotencyRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIdempotencyRepository(t mockConstructorTestingTNewIdempotencyRepository) *IdempotencyRepository {
	mock := &IdempotencyRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
