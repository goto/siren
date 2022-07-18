// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Encryptor is an autogenerated mock type for the Encryptor type
type Encryptor struct {
	mock.Mock
}

type Encryptor_Expecter struct {
	mock *mock.Mock
}

func (_m *Encryptor) EXPECT() *Encryptor_Expecter {
	return &Encryptor_Expecter{mock: &_m.Mock}
}

// Decrypt provides a mock function with given fields: str
func (_m *Encryptor) Decrypt(str string) (string, error) {
	ret := _m.Called(str)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(str)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(str)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Encryptor_Decrypt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Decrypt'
type Encryptor_Decrypt_Call struct {
	*mock.Call
}

// Decrypt is a helper method to define mock.On call
//  - str string
func (_e *Encryptor_Expecter) Decrypt(str interface{}) *Encryptor_Decrypt_Call {
	return &Encryptor_Decrypt_Call{Call: _e.mock.On("Decrypt", str)}
}

func (_c *Encryptor_Decrypt_Call) Run(run func(str string)) *Encryptor_Decrypt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Encryptor_Decrypt_Call) Return(_a0 string, _a1 error) *Encryptor_Decrypt_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Encrypt provides a mock function with given fields: str
func (_m *Encryptor) Encrypt(str string) (string, error) {
	ret := _m.Called(str)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(str)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(str)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Encryptor_Encrypt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Encrypt'
type Encryptor_Encrypt_Call struct {
	*mock.Call
}

// Encrypt is a helper method to define mock.On call
//  - str string
func (_e *Encryptor_Expecter) Encrypt(str interface{}) *Encryptor_Encrypt_Call {
	return &Encryptor_Encrypt_Call{Call: _e.mock.On("Encrypt", str)}
}

func (_c *Encryptor_Encrypt_Call) Run(run func(str string)) *Encryptor_Encrypt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Encryptor_Encrypt_Call) Return(_a0 string, _a1 error) *Encryptor_Encrypt_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewEncryptor interface {
	mock.TestingT
	Cleanup(func())
}

// NewEncryptor creates a new instance of Encryptor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEncryptor(t mockConstructorTestingTNewEncryptor) *Encryptor {
	mock := &Encryptor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
