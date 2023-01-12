// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	template "github.com/odpf/siren/core/template"
	mock "github.com/stretchr/testify/mock"
)

// TemplateService is an autogenerated mock type for the TemplateService type
type TemplateService struct {
	mock.Mock
}

type TemplateService_Expecter struct {
	mock *mock.Mock
}

func (_m *TemplateService) EXPECT() *TemplateService_Expecter {
	return &TemplateService_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *TemplateService) Delete(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TemplateService_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type TemplateService_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *TemplateService_Expecter) Delete(_a0 interface{}, _a1 interface{}) *TemplateService_Delete_Call {
	return &TemplateService_Delete_Call{Call: _e.mock.On("Delete", _a0, _a1)}
}

func (_c *TemplateService_Delete_Call) Run(run func(_a0 context.Context, _a1 string)) *TemplateService_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *TemplateService_Delete_Call) Return(_a0 error) *TemplateService_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetByName provides a mock function with given fields: _a0, _a1
func (_m *TemplateService) GetByName(_a0 context.Context, _a1 string) (*template.Template, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *template.Template
	if rf, ok := ret.Get(0).(func(context.Context, string) *template.Template); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*template.Template)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TemplateService_GetByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByName'
type TemplateService_GetByName_Call struct {
	*mock.Call
}

// GetByName is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *TemplateService_Expecter) GetByName(_a0 interface{}, _a1 interface{}) *TemplateService_GetByName_Call {
	return &TemplateService_GetByName_Call{Call: _e.mock.On("GetByName", _a0, _a1)}
}

func (_c *TemplateService_GetByName_Call) Run(run func(_a0 context.Context, _a1 string)) *TemplateService_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *TemplateService_GetByName_Call) Return(_a0 *template.Template, _a1 error) *TemplateService_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// List provides a mock function with given fields: _a0, _a1
func (_m *TemplateService) List(_a0 context.Context, _a1 template.Filter) ([]template.Template, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []template.Template
	if rf, ok := ret.Get(0).(func(context.Context, template.Filter) []template.Template); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]template.Template)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, template.Filter) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TemplateService_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type TemplateService_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 template.Filter
func (_e *TemplateService_Expecter) List(_a0 interface{}, _a1 interface{}) *TemplateService_List_Call {
	return &TemplateService_List_Call{Call: _e.mock.On("List", _a0, _a1)}
}

func (_c *TemplateService_List_Call) Run(run func(_a0 context.Context, _a1 template.Filter)) *TemplateService_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(template.Filter))
	})
	return _c
}

func (_c *TemplateService_List_Call) Return(_a0 []template.Template, _a1 error) *TemplateService_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Render provides a mock function with given fields: _a0, _a1, _a2
func (_m *TemplateService) Render(_a0 context.Context, _a1 string, _a2 map[string]string) (string, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]string) string); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, map[string]string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TemplateService_Render_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Render'
type TemplateService_Render_Call struct {
	*mock.Call
}

// Render is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
//   - _a2 map[string]string
func (_e *TemplateService_Expecter) Render(_a0 interface{}, _a1 interface{}, _a2 interface{}) *TemplateService_Render_Call {
	return &TemplateService_Render_Call{Call: _e.mock.On("Render", _a0, _a1, _a2)}
}

func (_c *TemplateService_Render_Call) Run(run func(_a0 context.Context, _a1 string, _a2 map[string]string)) *TemplateService_Render_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(map[string]string))
	})
	return _c
}

func (_c *TemplateService_Render_Call) Return(_a0 string, _a1 error) *TemplateService_Render_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Upsert provides a mock function with given fields: _a0, _a1
func (_m *TemplateService) Upsert(_a0 context.Context, _a1 *template.Template) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *template.Template) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TemplateService_Upsert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Upsert'
type TemplateService_Upsert_Call struct {
	*mock.Call
}

// Upsert is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *template.Template
func (_e *TemplateService_Expecter) Upsert(_a0 interface{}, _a1 interface{}) *TemplateService_Upsert_Call {
	return &TemplateService_Upsert_Call{Call: _e.mock.On("Upsert", _a0, _a1)}
}

func (_c *TemplateService_Upsert_Call) Run(run func(_a0 context.Context, _a1 *template.Template)) *TemplateService_Upsert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*template.Template))
	})
	return _c
}

func (_c *TemplateService_Upsert_Call) Return(_a0 error) *TemplateService_Upsert_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewTemplateService interface {
	mock.TestingT
	Cleanup(func())
}

// NewTemplateService creates a new instance of TemplateService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTemplateService(t mockConstructorTestingTNewTemplateService) *TemplateService {
	mock := &TemplateService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}