// Code generated by mockery 2.9.0. DO NOT EDIT.

package rules

import (
	domain "github.com/odpf/siren/domain"
	mock "github.com/stretchr/testify/mock"
)

// RuleRepositoryMock is an autogenerated mock type for the RuleRepository type
type RuleRepositoryMock struct {
	mock.Mock
}

// Get provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4
func (_m *RuleRepositoryMock) Get(_a0 string, _a1 string, _a2 string, _a3 string, _a4 uint64) ([]Rule, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4)

	var r0 []Rule
	if rf, ok := ret.Get(0).(func(string, string, string, string, uint64) []Rule); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Rule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string, uint64) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Migrate provides a mock function with given fields:
func (_m *RuleRepositoryMock) Migrate() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Upsert provides a mock function with given fields: _a0, _a1
func (_m *RuleRepositoryMock) Upsert(_a0 *Rule, _a1 domain.TemplatesService) (*Rule, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *Rule
	if rf, ok := ret.Get(0).(func(*Rule, domain.TemplatesService) *Rule); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Rule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*Rule, domain.TemplatesService) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}