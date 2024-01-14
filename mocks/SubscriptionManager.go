// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	api "github.com/enbility/spine-go/api"
	mock "github.com/stretchr/testify/mock"

	model "github.com/enbility/spine-go/model"
)

// SubscriptionManager is an autogenerated mock type for the SubscriptionManager type
type SubscriptionManager struct {
	mock.Mock
}

type SubscriptionManager_Expecter struct {
	mock *mock.Mock
}

func (_m *SubscriptionManager) EXPECT() *SubscriptionManager_Expecter {
	return &SubscriptionManager_Expecter{mock: &_m.Mock}
}

// AddSubscription provides a mock function with given fields: remoteDevice, data
func (_m *SubscriptionManager) AddSubscription(remoteDevice api.DeviceRemote, data model.SubscriptionManagementRequestCallType) error {
	ret := _m.Called(remoteDevice, data)

	if len(ret) == 0 {
		panic("no return value specified for AddSubscription")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(api.DeviceRemote, model.SubscriptionManagementRequestCallType) error); ok {
		r0 = rf(remoteDevice, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubscriptionManager_AddSubscription_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddSubscription'
type SubscriptionManager_AddSubscription_Call struct {
	*mock.Call
}

// AddSubscription is a helper method to define mock.On call
//   - remoteDevice api.DeviceRemote
//   - data model.SubscriptionManagementRequestCallType
func (_e *SubscriptionManager_Expecter) AddSubscription(remoteDevice interface{}, data interface{}) *SubscriptionManager_AddSubscription_Call {
	return &SubscriptionManager_AddSubscription_Call{Call: _e.mock.On("AddSubscription", remoteDevice, data)}
}

func (_c *SubscriptionManager_AddSubscription_Call) Run(run func(remoteDevice api.DeviceRemote, data model.SubscriptionManagementRequestCallType)) *SubscriptionManager_AddSubscription_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.DeviceRemote), args[1].(model.SubscriptionManagementRequestCallType))
	})
	return _c
}

func (_c *SubscriptionManager_AddSubscription_Call) Return(_a0 error) *SubscriptionManager_AddSubscription_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SubscriptionManager_AddSubscription_Call) RunAndReturn(run func(api.DeviceRemote, model.SubscriptionManagementRequestCallType) error) *SubscriptionManager_AddSubscription_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveSubscription provides a mock function with given fields: data, remoteDevice
func (_m *SubscriptionManager) RemoveSubscription(data model.SubscriptionManagementDeleteCallType, remoteDevice api.DeviceRemote) error {
	ret := _m.Called(data, remoteDevice)

	if len(ret) == 0 {
		panic("no return value specified for RemoveSubscription")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(model.SubscriptionManagementDeleteCallType, api.DeviceRemote) error); ok {
		r0 = rf(data, remoteDevice)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubscriptionManager_RemoveSubscription_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveSubscription'
type SubscriptionManager_RemoveSubscription_Call struct {
	*mock.Call
}

// RemoveSubscription is a helper method to define mock.On call
//   - data model.SubscriptionManagementDeleteCallType
//   - remoteDevice api.DeviceRemote
func (_e *SubscriptionManager_Expecter) RemoveSubscription(data interface{}, remoteDevice interface{}) *SubscriptionManager_RemoveSubscription_Call {
	return &SubscriptionManager_RemoveSubscription_Call{Call: _e.mock.On("RemoveSubscription", data, remoteDevice)}
}

func (_c *SubscriptionManager_RemoveSubscription_Call) Run(run func(data model.SubscriptionManagementDeleteCallType, remoteDevice api.DeviceRemote)) *SubscriptionManager_RemoveSubscription_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.SubscriptionManagementDeleteCallType), args[1].(api.DeviceRemote))
	})
	return _c
}

func (_c *SubscriptionManager_RemoveSubscription_Call) Return(_a0 error) *SubscriptionManager_RemoveSubscription_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SubscriptionManager_RemoveSubscription_Call) RunAndReturn(run func(model.SubscriptionManagementDeleteCallType, api.DeviceRemote) error) *SubscriptionManager_RemoveSubscription_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveSubscriptionsForDevice provides a mock function with given fields: remoteDevice
func (_m *SubscriptionManager) RemoveSubscriptionsForDevice(remoteDevice api.DeviceRemote) {
	_m.Called(remoteDevice)
}

// SubscriptionManager_RemoveSubscriptionsForDevice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveSubscriptionsForDevice'
type SubscriptionManager_RemoveSubscriptionsForDevice_Call struct {
	*mock.Call
}

// RemoveSubscriptionsForDevice is a helper method to define mock.On call
//   - remoteDevice api.DeviceRemote
func (_e *SubscriptionManager_Expecter) RemoveSubscriptionsForDevice(remoteDevice interface{}) *SubscriptionManager_RemoveSubscriptionsForDevice_Call {
	return &SubscriptionManager_RemoveSubscriptionsForDevice_Call{Call: _e.mock.On("RemoveSubscriptionsForDevice", remoteDevice)}
}

func (_c *SubscriptionManager_RemoveSubscriptionsForDevice_Call) Run(run func(remoteDevice api.DeviceRemote)) *SubscriptionManager_RemoveSubscriptionsForDevice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.DeviceRemote))
	})
	return _c
}

func (_c *SubscriptionManager_RemoveSubscriptionsForDevice_Call) Return() *SubscriptionManager_RemoveSubscriptionsForDevice_Call {
	_c.Call.Return()
	return _c
}

func (_c *SubscriptionManager_RemoveSubscriptionsForDevice_Call) RunAndReturn(run func(api.DeviceRemote)) *SubscriptionManager_RemoveSubscriptionsForDevice_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveSubscriptionsForEntity provides a mock function with given fields: remoteEntity
func (_m *SubscriptionManager) RemoveSubscriptionsForEntity(remoteEntity api.EntityRemote) {
	_m.Called(remoteEntity)
}

// SubscriptionManager_RemoveSubscriptionsForEntity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveSubscriptionsForEntity'
type SubscriptionManager_RemoveSubscriptionsForEntity_Call struct {
	*mock.Call
}

// RemoveSubscriptionsForEntity is a helper method to define mock.On call
//   - remoteEntity api.EntityRemote
func (_e *SubscriptionManager_Expecter) RemoveSubscriptionsForEntity(remoteEntity interface{}) *SubscriptionManager_RemoveSubscriptionsForEntity_Call {
	return &SubscriptionManager_RemoveSubscriptionsForEntity_Call{Call: _e.mock.On("RemoveSubscriptionsForEntity", remoteEntity)}
}

func (_c *SubscriptionManager_RemoveSubscriptionsForEntity_Call) Run(run func(remoteEntity api.EntityRemote)) *SubscriptionManager_RemoveSubscriptionsForEntity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.EntityRemote))
	})
	return _c
}

func (_c *SubscriptionManager_RemoveSubscriptionsForEntity_Call) Return() *SubscriptionManager_RemoveSubscriptionsForEntity_Call {
	_c.Call.Return()
	return _c
}

func (_c *SubscriptionManager_RemoveSubscriptionsForEntity_Call) RunAndReturn(run func(api.EntityRemote)) *SubscriptionManager_RemoveSubscriptionsForEntity_Call {
	_c.Call.Return(run)
	return _c
}

// Subscriptions provides a mock function with given fields: remoteDevice
func (_m *SubscriptionManager) Subscriptions(remoteDevice api.DeviceRemote) []*api.SubscriptionEntry {
	ret := _m.Called(remoteDevice)

	if len(ret) == 0 {
		panic("no return value specified for Subscriptions")
	}

	var r0 []*api.SubscriptionEntry
	if rf, ok := ret.Get(0).(func(api.DeviceRemote) []*api.SubscriptionEntry); ok {
		r0 = rf(remoteDevice)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*api.SubscriptionEntry)
		}
	}

	return r0
}

// SubscriptionManager_Subscriptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Subscriptions'
type SubscriptionManager_Subscriptions_Call struct {
	*mock.Call
}

// Subscriptions is a helper method to define mock.On call
//   - remoteDevice api.DeviceRemote
func (_e *SubscriptionManager_Expecter) Subscriptions(remoteDevice interface{}) *SubscriptionManager_Subscriptions_Call {
	return &SubscriptionManager_Subscriptions_Call{Call: _e.mock.On("Subscriptions", remoteDevice)}
}

func (_c *SubscriptionManager_Subscriptions_Call) Run(run func(remoteDevice api.DeviceRemote)) *SubscriptionManager_Subscriptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.DeviceRemote))
	})
	return _c
}

func (_c *SubscriptionManager_Subscriptions_Call) Return(_a0 []*api.SubscriptionEntry) *SubscriptionManager_Subscriptions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SubscriptionManager_Subscriptions_Call) RunAndReturn(run func(api.DeviceRemote) []*api.SubscriptionEntry) *SubscriptionManager_Subscriptions_Call {
	_c.Call.Return(run)
	return _c
}

// SubscriptionsOnFeature provides a mock function with given fields: featureAddress
func (_m *SubscriptionManager) SubscriptionsOnFeature(featureAddress model.FeatureAddressType) []*api.SubscriptionEntry {
	ret := _m.Called(featureAddress)

	if len(ret) == 0 {
		panic("no return value specified for SubscriptionsOnFeature")
	}

	var r0 []*api.SubscriptionEntry
	if rf, ok := ret.Get(0).(func(model.FeatureAddressType) []*api.SubscriptionEntry); ok {
		r0 = rf(featureAddress)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*api.SubscriptionEntry)
		}
	}

	return r0
}

// SubscriptionManager_SubscriptionsOnFeature_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SubscriptionsOnFeature'
type SubscriptionManager_SubscriptionsOnFeature_Call struct {
	*mock.Call
}

// SubscriptionsOnFeature is a helper method to define mock.On call
//   - featureAddress model.FeatureAddressType
func (_e *SubscriptionManager_Expecter) SubscriptionsOnFeature(featureAddress interface{}) *SubscriptionManager_SubscriptionsOnFeature_Call {
	return &SubscriptionManager_SubscriptionsOnFeature_Call{Call: _e.mock.On("SubscriptionsOnFeature", featureAddress)}
}

func (_c *SubscriptionManager_SubscriptionsOnFeature_Call) Run(run func(featureAddress model.FeatureAddressType)) *SubscriptionManager_SubscriptionsOnFeature_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.FeatureAddressType))
	})
	return _c
}

func (_c *SubscriptionManager_SubscriptionsOnFeature_Call) Return(_a0 []*api.SubscriptionEntry) *SubscriptionManager_SubscriptionsOnFeature_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SubscriptionManager_SubscriptionsOnFeature_Call) RunAndReturn(run func(model.FeatureAddressType) []*api.SubscriptionEntry) *SubscriptionManager_SubscriptionsOnFeature_Call {
	_c.Call.Return(run)
	return _c
}

// NewSubscriptionManager creates a new instance of SubscriptionManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSubscriptionManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *SubscriptionManager {
	mock := &SubscriptionManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
