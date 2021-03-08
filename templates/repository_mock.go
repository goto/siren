package templates

import (
	"github.com/stretchr/testify/mock"
)

// TemplatesRepositoryMock is an autogenerated mock type for the TemplatesRepositoryMock type
type TemplatesRepositoryMock struct {
	mock.Mock
}

// Delete provides a mock function with given fields: _a0
func (_m *TemplatesRepositoryMock) Delete(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByName provides a mock function with given fields: _a0
func (_m *TemplatesRepositoryMock) GetByName(_a0 string) (*Template, error) {
	ret := _m.Called(_a0)

	var r0 *Template
	if rf, ok := ret.Get(0).(func(string) *Template); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Template)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Index provides a mock function with given fields: _a0
func (_m *TemplatesRepositoryMock) Index(_a0 string) ([]Template, error) {
	ret := _m.Called(_a0)

	var r0 []Template
	if rf, ok := ret.Get(0).(func(string) []Template); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Template)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Migrate provides a mock function with given fields:
func (_m *TemplatesRepositoryMock) Migrate() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Render provides a mock function with given fields: _a0, _a1
func (_m *TemplatesRepositoryMock) Render(_a0 string, _a1 map[string]string) (string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, map[string]string) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, map[string]string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Upsert provides a mock function with given fields: _a0
func (_m *TemplatesRepositoryMock) Upsert(_a0 *Template) (*Template, error) {
	ret := _m.Called(_a0)

	var r0 *Template
	if rf, ok := ret.Get(0).(func(*Template) *Template); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Template)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*Template) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
