// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	subscription "github.com/goto/siren/core/subscription"
)

// SubscriptionService is an autogenerated mock type for the SubscriptionService type
type SubscriptionService struct {
	mock.Mock
}

type SubscriptionService_Expecter struct {
	mock *mock.Mock
}

func (_m *SubscriptionService) EXPECT() *SubscriptionService_Expecter {
	return &SubscriptionService_Expecter{mock: &_m.Mock}
}

// MatchByLabels provides a mock function with given fields: ctx, namespaceID, labels
func (_m *SubscriptionService) MatchByLabels(ctx context.Context, namespaceID uint64, labels map[string]string) ([]subscription.Subscription, error) {
	ret := _m.Called(ctx, namespaceID, labels)

	if len(ret) == 0 {
		panic("no return value specified for MatchByLabels")
	}

	var r0 []subscription.Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, map[string]string) ([]subscription.Subscription, error)); ok {
		return rf(ctx, namespaceID, labels)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, map[string]string) []subscription.Subscription); ok {
		r0 = rf(ctx, namespaceID, labels)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]subscription.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, map[string]string) error); ok {
		r1 = rf(ctx, namespaceID, labels)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubscriptionService_MatchByLabels_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MatchByLabels'
type SubscriptionService_MatchByLabels_Call struct {
	*mock.Call
}

// MatchByLabels is a helper method to define mock.On call
//   - ctx context.Context
//   - namespaceID uint64
//   - labels map[string]string
func (_e *SubscriptionService_Expecter) MatchByLabels(ctx interface{}, namespaceID interface{}, labels interface{}) *SubscriptionService_MatchByLabels_Call {
	return &SubscriptionService_MatchByLabels_Call{Call: _e.mock.On("MatchByLabels", ctx, namespaceID, labels)}
}

func (_c *SubscriptionService_MatchByLabels_Call) Run(run func(ctx context.Context, namespaceID uint64, labels map[string]string)) *SubscriptionService_MatchByLabels_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(map[string]string))
	})
	return _c
}

func (_c *SubscriptionService_MatchByLabels_Call) Return(_a0 []subscription.Subscription, _a1 error) *SubscriptionService_MatchByLabels_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SubscriptionService_MatchByLabels_Call) RunAndReturn(run func(context.Context, uint64, map[string]string) ([]subscription.Subscription, error)) *SubscriptionService_MatchByLabels_Call {
	_c.Call.Return(run)
	return _c
}

// MatchByLabelsV2 provides a mock function with given fields: ctx, namespaceID, labels
func (_m *SubscriptionService) MatchByLabelsV2(ctx context.Context, namespaceID uint64, labels map[string]string) ([]subscription.ReceiverView, error) {
	ret := _m.Called(ctx, namespaceID, labels)

	if len(ret) == 0 {
		panic("no return value specified for MatchByLabelsV2")
	}

	var r0 []subscription.ReceiverView
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, map[string]string) ([]subscription.ReceiverView, error)); ok {
		return rf(ctx, namespaceID, labels)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, map[string]string) []subscription.ReceiverView); ok {
		r0 = rf(ctx, namespaceID, labels)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]subscription.ReceiverView)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, map[string]string) error); ok {
		r1 = rf(ctx, namespaceID, labels)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubscriptionService_MatchByLabelsV2_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MatchByLabelsV2'
type SubscriptionService_MatchByLabelsV2_Call struct {
	*mock.Call
}

// MatchByLabelsV2 is a helper method to define mock.On call
//   - ctx context.Context
//   - namespaceID uint64
//   - labels map[string]string
func (_e *SubscriptionService_Expecter) MatchByLabelsV2(ctx interface{}, namespaceID interface{}, labels interface{}) *SubscriptionService_MatchByLabelsV2_Call {
	return &SubscriptionService_MatchByLabelsV2_Call{Call: _e.mock.On("MatchByLabelsV2", ctx, namespaceID, labels)}
}

func (_c *SubscriptionService_MatchByLabelsV2_Call) Run(run func(ctx context.Context, namespaceID uint64, labels map[string]string)) *SubscriptionService_MatchByLabelsV2_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(map[string]string))
	})
	return _c
}

func (_c *SubscriptionService_MatchByLabelsV2_Call) Return(_a0 []subscription.ReceiverView, _a1 error) *SubscriptionService_MatchByLabelsV2_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SubscriptionService_MatchByLabelsV2_Call) RunAndReturn(run func(context.Context, uint64, map[string]string) ([]subscription.ReceiverView, error)) *SubscriptionService_MatchByLabelsV2_Call {
	_c.Call.Return(run)
	return _c
}

// NewSubscriptionService creates a new instance of SubscriptionService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSubscriptionService(t interface {
	mock.TestingT
	Cleanup(func())
}) *SubscriptionService {
	mock := &SubscriptionService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
