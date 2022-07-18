// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	receiver "github.com/odpf/siren/core/receiver"
	mock "github.com/stretchr/testify/mock"
)

// TypeService is an autogenerated mock type for the TypeService type
type TypeService struct {
	mock.Mock
}

type TypeService_Expecter struct {
	mock *mock.Mock
}

func (_m *TypeService) EXPECT() *TypeService_Expecter {
	return &TypeService_Expecter{mock: &_m.Mock}
}

// Decrypt provides a mock function with given fields: r
func (_m *TypeService) Decrypt(r *receiver.Receiver) error {
	ret := _m.Called(r)

	var r0 error
	if rf, ok := ret.Get(0).(func(*receiver.Receiver) error); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TypeService_Decrypt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Decrypt'
type TypeService_Decrypt_Call struct {
	*mock.Call
}

// Decrypt is a helper method to define mock.On call
//  - r *receiver.Receiver
func (_e *TypeService_Expecter) Decrypt(r interface{}) *TypeService_Decrypt_Call {
	return &TypeService_Decrypt_Call{Call: _e.mock.On("Decrypt", r)}
}

func (_c *TypeService_Decrypt_Call) Run(run func(r *receiver.Receiver)) *TypeService_Decrypt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*receiver.Receiver))
	})
	return _c
}

func (_c *TypeService_Decrypt_Call) Return(_a0 error) *TypeService_Decrypt_Call {
	_c.Call.Return(_a0)
	return _c
}

// Encrypt provides a mock function with given fields: r
func (_m *TypeService) Encrypt(r *receiver.Receiver) error {
	ret := _m.Called(r)

	var r0 error
	if rf, ok := ret.Get(0).(func(*receiver.Receiver) error); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TypeService_Encrypt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Encrypt'
type TypeService_Encrypt_Call struct {
	*mock.Call
}

// Encrypt is a helper method to define mock.On call
//  - r *receiver.Receiver
func (_e *TypeService_Expecter) Encrypt(r interface{}) *TypeService_Encrypt_Call {
	return &TypeService_Encrypt_Call{Call: _e.mock.On("Encrypt", r)}
}

func (_c *TypeService_Encrypt_Call) Run(run func(r *receiver.Receiver)) *TypeService_Encrypt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*receiver.Receiver))
	})
	return _c
}

func (_c *TypeService_Encrypt_Call) Return(_a0 error) *TypeService_Encrypt_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetSubscriptionConfig provides a mock function with given fields: subsConfs, receiverConfs
func (_m *TypeService) GetSubscriptionConfig(subsConfs map[string]string, receiverConfs receiver.Configurations) (map[string]string, error) {
	ret := _m.Called(subsConfs, receiverConfs)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(map[string]string, receiver.Configurations) map[string]string); ok {
		r0 = rf(subsConfs, receiverConfs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]string, receiver.Configurations) error); ok {
		r1 = rf(subsConfs, receiverConfs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TypeService_GetSubscriptionConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSubscriptionConfig'
type TypeService_GetSubscriptionConfig_Call struct {
	*mock.Call
}

// GetSubscriptionConfig is a helper method to define mock.On call
//  - subsConfs map[string]string
//  - receiverConfs receiver.Configurations
func (_e *TypeService_Expecter) GetSubscriptionConfig(subsConfs interface{}, receiverConfs interface{}) *TypeService_GetSubscriptionConfig_Call {
	return &TypeService_GetSubscriptionConfig_Call{Call: _e.mock.On("GetSubscriptionConfig", subsConfs, receiverConfs)}
}

func (_c *TypeService_GetSubscriptionConfig_Call) Run(run func(subsConfs map[string]string, receiverConfs receiver.Configurations)) *TypeService_GetSubscriptionConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(map[string]string), args[1].(receiver.Configurations))
	})
	return _c
}

func (_c *TypeService_GetSubscriptionConfig_Call) Return(_a0 map[string]string, _a1 error) *TypeService_GetSubscriptionConfig_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Notify provides a mock function with given fields: ctx, rcv, payloadMessage
func (_m *TypeService) Notify(ctx context.Context, rcv *receiver.Receiver, payloadMessage receiver.NotificationMessage) error {
	ret := _m.Called(ctx, rcv, payloadMessage)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *receiver.Receiver, receiver.NotificationMessage) error); ok {
		r0 = rf(ctx, rcv, payloadMessage)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TypeService_Notify_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Notify'
type TypeService_Notify_Call struct {
	*mock.Call
}

// Notify is a helper method to define mock.On call
//  - ctx context.Context
//  - rcv *receiver.Receiver
//  - payloadMessage receiver.NotificationMessage
func (_e *TypeService_Expecter) Notify(ctx interface{}, rcv interface{}, payloadMessage interface{}) *TypeService_Notify_Call {
	return &TypeService_Notify_Call{Call: _e.mock.On("Notify", ctx, rcv, payloadMessage)}
}

func (_c *TypeService_Notify_Call) Run(run func(ctx context.Context, rcv *receiver.Receiver, payloadMessage receiver.NotificationMessage)) *TypeService_Notify_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*receiver.Receiver), args[2].(receiver.NotificationMessage))
	})
	return _c
}

func (_c *TypeService_Notify_Call) Return(_a0 error) *TypeService_Notify_Call {
	_c.Call.Return(_a0)
	return _c
}

// PopulateReceiver provides a mock function with given fields: ctx, rcv
func (_m *TypeService) PopulateReceiver(ctx context.Context, rcv *receiver.Receiver) (*receiver.Receiver, error) {
	ret := _m.Called(ctx, rcv)

	var r0 *receiver.Receiver
	if rf, ok := ret.Get(0).(func(context.Context, *receiver.Receiver) *receiver.Receiver); ok {
		r0 = rf(ctx, rcv)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*receiver.Receiver)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *receiver.Receiver) error); ok {
		r1 = rf(ctx, rcv)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TypeService_PopulateReceiver_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PopulateReceiver'
type TypeService_PopulateReceiver_Call struct {
	*mock.Call
}

// PopulateReceiver is a helper method to define mock.On call
//  - ctx context.Context
//  - rcv *receiver.Receiver
func (_e *TypeService_Expecter) PopulateReceiver(ctx interface{}, rcv interface{}) *TypeService_PopulateReceiver_Call {
	return &TypeService_PopulateReceiver_Call{Call: _e.mock.On("PopulateReceiver", ctx, rcv)}
}

func (_c *TypeService_PopulateReceiver_Call) Run(run func(ctx context.Context, rcv *receiver.Receiver)) *TypeService_PopulateReceiver_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*receiver.Receiver))
	})
	return _c
}

func (_c *TypeService_PopulateReceiver_Call) Return(_a0 *receiver.Receiver, _a1 error) *TypeService_PopulateReceiver_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ValidateConfiguration provides a mock function with given fields: rcv
func (_m *TypeService) ValidateConfiguration(rcv *receiver.Receiver) error {
	ret := _m.Called(rcv)

	var r0 error
	if rf, ok := ret.Get(0).(func(*receiver.Receiver) error); ok {
		r0 = rf(rcv)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TypeService_ValidateConfiguration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ValidateConfiguration'
type TypeService_ValidateConfiguration_Call struct {
	*mock.Call
}

// ValidateConfiguration is a helper method to define mock.On call
//  - rcv *receiver.Receiver
func (_e *TypeService_Expecter) ValidateConfiguration(rcv interface{}) *TypeService_ValidateConfiguration_Call {
	return &TypeService_ValidateConfiguration_Call{Call: _e.mock.On("ValidateConfiguration", rcv)}
}

func (_c *TypeService_ValidateConfiguration_Call) Run(run func(rcv *receiver.Receiver)) *TypeService_ValidateConfiguration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*receiver.Receiver))
	})
	return _c
}

func (_c *TypeService_ValidateConfiguration_Call) Return(_a0 error) *TypeService_ValidateConfiguration_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewTypeService interface {
	mock.TestingT
	Cleanup(func())
}

// NewTypeService creates a new instance of TypeService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTypeService(t mockConstructorTestingTNewTypeService) *TypeService {
	mock := &TypeService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
