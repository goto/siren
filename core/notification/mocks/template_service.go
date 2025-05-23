// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	template "github.com/goto/siren/core/template"
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

// GetByName provides a mock function with given fields: ctx, name
func (_m *TemplateService) GetByName(ctx context.Context, name string) (*template.Template, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetByName")
	}

	var r0 *template.Template
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*template.Template, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *template.Template); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*template.Template)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
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
//   - ctx context.Context
//   - name string
func (_e *TemplateService_Expecter) GetByName(ctx interface{}, name interface{}) *TemplateService_GetByName_Call {
	return &TemplateService_GetByName_Call{Call: _e.mock.On("GetByName", ctx, name)}
}

func (_c *TemplateService_GetByName_Call) Run(run func(ctx context.Context, name string)) *TemplateService_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *TemplateService_GetByName_Call) Return(_a0 *template.Template, _a1 error) *TemplateService_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TemplateService_GetByName_Call) RunAndReturn(run func(context.Context, string) (*template.Template, error)) *TemplateService_GetByName_Call {
	_c.Call.Return(run)
	return _c
}

// NewTemplateService creates a new instance of TemplateService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTemplateService(t interface {
	mock.TestingT
	Cleanup(func())
}) *TemplateService {
	mock := &TemplateService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
