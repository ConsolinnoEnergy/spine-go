// Code generated by mockery v2.40.3. DO NOT EDIT.

package mocks

import (
	model "github.com/enbility/spine-go/model"
	mock "github.com/stretchr/testify/mock"
)

// FunctionDataInterface is an autogenerated mock type for the FunctionDataInterface type
type FunctionDataInterface struct {
	mock.Mock
}

type FunctionDataInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *FunctionDataInterface) EXPECT() *FunctionDataInterface_Expecter {
	return &FunctionDataInterface_Expecter{mock: &_m.Mock}
}

// DataCopyAny provides a mock function with given fields:
func (_m *FunctionDataInterface) DataCopyAny() interface{} {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DataCopyAny")
	}

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// FunctionDataInterface_DataCopyAny_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DataCopyAny'
type FunctionDataInterface_DataCopyAny_Call struct {
	*mock.Call
}

// DataCopyAny is a helper method to define mock.On call
func (_e *FunctionDataInterface_Expecter) DataCopyAny() *FunctionDataInterface_DataCopyAny_Call {
	return &FunctionDataInterface_DataCopyAny_Call{Call: _e.mock.On("DataCopyAny")}
}

func (_c *FunctionDataInterface_DataCopyAny_Call) Run(run func()) *FunctionDataInterface_DataCopyAny_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FunctionDataInterface_DataCopyAny_Call) Return(_a0 interface{}) *FunctionDataInterface_DataCopyAny_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FunctionDataInterface_DataCopyAny_Call) RunAndReturn(run func() interface{}) *FunctionDataInterface_DataCopyAny_Call {
	_c.Call.Return(run)
	return _c
}

// FunctionType provides a mock function with given fields:
func (_m *FunctionDataInterface) FunctionType() model.FunctionType {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FunctionType")
	}

	var r0 model.FunctionType
	if rf, ok := ret.Get(0).(func() model.FunctionType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(model.FunctionType)
	}

	return r0
}

// FunctionDataInterface_FunctionType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FunctionType'
type FunctionDataInterface_FunctionType_Call struct {
	*mock.Call
}

// FunctionType is a helper method to define mock.On call
func (_e *FunctionDataInterface_Expecter) FunctionType() *FunctionDataInterface_FunctionType_Call {
	return &FunctionDataInterface_FunctionType_Call{Call: _e.mock.On("FunctionType")}
}

func (_c *FunctionDataInterface_FunctionType_Call) Run(run func()) *FunctionDataInterface_FunctionType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FunctionDataInterface_FunctionType_Call) Return(_a0 model.FunctionType) *FunctionDataInterface_FunctionType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FunctionDataInterface_FunctionType_Call) RunAndReturn(run func() model.FunctionType) *FunctionDataInterface_FunctionType_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateDataAny provides a mock function with given fields: data, filterPartial, filterDelete
func (_m *FunctionDataInterface) UpdateDataAny(data interface{}, filterPartial *model.FilterType, filterDelete *model.FilterType) *model.ErrorType {
	ret := _m.Called(data, filterPartial, filterDelete)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDataAny")
	}

	var r0 *model.ErrorType
	if rf, ok := ret.Get(0).(func(interface{}, *model.FilterType, *model.FilterType) *model.ErrorType); ok {
		r0 = rf(data, filterPartial, filterDelete)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ErrorType)
		}
	}

	return r0
}

// FunctionDataInterface_UpdateDataAny_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateDataAny'
type FunctionDataInterface_UpdateDataAny_Call struct {
	*mock.Call
}

// UpdateDataAny is a helper method to define mock.On call
//   - data interface{}
//   - filterPartial *model.FilterType
//   - filterDelete *model.FilterType
func (_e *FunctionDataInterface_Expecter) UpdateDataAny(data interface{}, filterPartial interface{}, filterDelete interface{}) *FunctionDataInterface_UpdateDataAny_Call {
	return &FunctionDataInterface_UpdateDataAny_Call{Call: _e.mock.On("UpdateDataAny", data, filterPartial, filterDelete)}
}

func (_c *FunctionDataInterface_UpdateDataAny_Call) Run(run func(data interface{}, filterPartial *model.FilterType, filterDelete *model.FilterType)) *FunctionDataInterface_UpdateDataAny_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}), args[1].(*model.FilterType), args[2].(*model.FilterType))
	})
	return _c
}

func (_c *FunctionDataInterface_UpdateDataAny_Call) Return(_a0 *model.ErrorType) *FunctionDataInterface_UpdateDataAny_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FunctionDataInterface_UpdateDataAny_Call) RunAndReturn(run func(interface{}, *model.FilterType, *model.FilterType) *model.ErrorType) *FunctionDataInterface_UpdateDataAny_Call {
	_c.Call.Return(run)
	return _c
}

// NewFunctionDataInterface creates a new instance of FunctionDataInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFunctionDataInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *FunctionDataInterface {
	mock := &FunctionDataInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
