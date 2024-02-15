// Code generated by mockery v2.40.3. DO NOT EDIT.

package mocks

import (
	model "github.com/enbility/spine-go/model"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// PendingRequestsInterface is an autogenerated mock type for the PendingRequestsInterface type
type PendingRequestsInterface struct {
	mock.Mock
}

type PendingRequestsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *PendingRequestsInterface) EXPECT() *PendingRequestsInterface_Expecter {
	return &PendingRequestsInterface_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: ski, counter, maxDelay
func (_m *PendingRequestsInterface) Add(ski string, counter model.MsgCounterType, maxDelay time.Duration) {
	_m.Called(ski, counter, maxDelay)
}

// PendingRequestsInterface_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type PendingRequestsInterface_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - ski string
//   - counter model.MsgCounterType
//   - maxDelay time.Duration
func (_e *PendingRequestsInterface_Expecter) Add(ski interface{}, counter interface{}, maxDelay interface{}) *PendingRequestsInterface_Add_Call {
	return &PendingRequestsInterface_Add_Call{Call: _e.mock.On("Add", ski, counter, maxDelay)}
}

func (_c *PendingRequestsInterface_Add_Call) Run(run func(ski string, counter model.MsgCounterType, maxDelay time.Duration)) *PendingRequestsInterface_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(model.MsgCounterType), args[2].(time.Duration))
	})
	return _c
}

func (_c *PendingRequestsInterface_Add_Call) Return() *PendingRequestsInterface_Add_Call {
	_c.Call.Return()
	return _c
}

func (_c *PendingRequestsInterface_Add_Call) RunAndReturn(run func(string, model.MsgCounterType, time.Duration)) *PendingRequestsInterface_Add_Call {
	_c.Call.Return(run)
	return _c
}

// GetData provides a mock function with given fields: ski, counter
func (_m *PendingRequestsInterface) GetData(ski string, counter model.MsgCounterType) (interface{}, *model.ErrorType) {
	ret := _m.Called(ski, counter)

	if len(ret) == 0 {
		panic("no return value specified for GetData")
	}

	var r0 interface{}
	var r1 *model.ErrorType
	if rf, ok := ret.Get(0).(func(string, model.MsgCounterType) (interface{}, *model.ErrorType)); ok {
		return rf(ski, counter)
	}
	if rf, ok := ret.Get(0).(func(string, model.MsgCounterType) interface{}); ok {
		r0 = rf(ski, counter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string, model.MsgCounterType) *model.ErrorType); ok {
		r1 = rf(ski, counter)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.ErrorType)
		}
	}

	return r0, r1
}

// PendingRequestsInterface_GetData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetData'
type PendingRequestsInterface_GetData_Call struct {
	*mock.Call
}

// GetData is a helper method to define mock.On call
//   - ski string
//   - counter model.MsgCounterType
func (_e *PendingRequestsInterface_Expecter) GetData(ski interface{}, counter interface{}) *PendingRequestsInterface_GetData_Call {
	return &PendingRequestsInterface_GetData_Call{Call: _e.mock.On("GetData", ski, counter)}
}

func (_c *PendingRequestsInterface_GetData_Call) Run(run func(ski string, counter model.MsgCounterType)) *PendingRequestsInterface_GetData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(model.MsgCounterType))
	})
	return _c
}

func (_c *PendingRequestsInterface_GetData_Call) Return(_a0 interface{}, _a1 *model.ErrorType) *PendingRequestsInterface_GetData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PendingRequestsInterface_GetData_Call) RunAndReturn(run func(string, model.MsgCounterType) (interface{}, *model.ErrorType)) *PendingRequestsInterface_GetData_Call {
	_c.Call.Return(run)
	return _c
}

// Remove provides a mock function with given fields: ski, counter
func (_m *PendingRequestsInterface) Remove(ski string, counter model.MsgCounterType) *model.ErrorType {
	ret := _m.Called(ski, counter)

	if len(ret) == 0 {
		panic("no return value specified for Remove")
	}

	var r0 *model.ErrorType
	if rf, ok := ret.Get(0).(func(string, model.MsgCounterType) *model.ErrorType); ok {
		r0 = rf(ski, counter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ErrorType)
		}
	}

	return r0
}

// PendingRequestsInterface_Remove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Remove'
type PendingRequestsInterface_Remove_Call struct {
	*mock.Call
}

// Remove is a helper method to define mock.On call
//   - ski string
//   - counter model.MsgCounterType
func (_e *PendingRequestsInterface_Expecter) Remove(ski interface{}, counter interface{}) *PendingRequestsInterface_Remove_Call {
	return &PendingRequestsInterface_Remove_Call{Call: _e.mock.On("Remove", ski, counter)}
}

func (_c *PendingRequestsInterface_Remove_Call) Run(run func(ski string, counter model.MsgCounterType)) *PendingRequestsInterface_Remove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(model.MsgCounterType))
	})
	return _c
}

func (_c *PendingRequestsInterface_Remove_Call) Return(_a0 *model.ErrorType) *PendingRequestsInterface_Remove_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PendingRequestsInterface_Remove_Call) RunAndReturn(run func(string, model.MsgCounterType) *model.ErrorType) *PendingRequestsInterface_Remove_Call {
	_c.Call.Return(run)
	return _c
}

// SetData provides a mock function with given fields: ski, counter, data
func (_m *PendingRequestsInterface) SetData(ski string, counter model.MsgCounterType, data interface{}) *model.ErrorType {
	ret := _m.Called(ski, counter, data)

	if len(ret) == 0 {
		panic("no return value specified for SetData")
	}

	var r0 *model.ErrorType
	if rf, ok := ret.Get(0).(func(string, model.MsgCounterType, interface{}) *model.ErrorType); ok {
		r0 = rf(ski, counter, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ErrorType)
		}
	}

	return r0
}

// PendingRequestsInterface_SetData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetData'
type PendingRequestsInterface_SetData_Call struct {
	*mock.Call
}

// SetData is a helper method to define mock.On call
//   - ski string
//   - counter model.MsgCounterType
//   - data interface{}
func (_e *PendingRequestsInterface_Expecter) SetData(ski interface{}, counter interface{}, data interface{}) *PendingRequestsInterface_SetData_Call {
	return &PendingRequestsInterface_SetData_Call{Call: _e.mock.On("SetData", ski, counter, data)}
}

func (_c *PendingRequestsInterface_SetData_Call) Run(run func(ski string, counter model.MsgCounterType, data interface{})) *PendingRequestsInterface_SetData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(model.MsgCounterType), args[2].(interface{}))
	})
	return _c
}

func (_c *PendingRequestsInterface_SetData_Call) Return(_a0 *model.ErrorType) *PendingRequestsInterface_SetData_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PendingRequestsInterface_SetData_Call) RunAndReturn(run func(string, model.MsgCounterType, interface{}) *model.ErrorType) *PendingRequestsInterface_SetData_Call {
	_c.Call.Return(run)
	return _c
}

// SetResult provides a mock function with given fields: ski, counter, errorResult
func (_m *PendingRequestsInterface) SetResult(ski string, counter model.MsgCounterType, errorResult *model.ErrorType) *model.ErrorType {
	ret := _m.Called(ski, counter, errorResult)

	if len(ret) == 0 {
		panic("no return value specified for SetResult")
	}

	var r0 *model.ErrorType
	if rf, ok := ret.Get(0).(func(string, model.MsgCounterType, *model.ErrorType) *model.ErrorType); ok {
		r0 = rf(ski, counter, errorResult)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ErrorType)
		}
	}

	return r0
}

// PendingRequestsInterface_SetResult_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetResult'
type PendingRequestsInterface_SetResult_Call struct {
	*mock.Call
}

// SetResult is a helper method to define mock.On call
//   - ski string
//   - counter model.MsgCounterType
//   - errorResult *model.ErrorType
func (_e *PendingRequestsInterface_Expecter) SetResult(ski interface{}, counter interface{}, errorResult interface{}) *PendingRequestsInterface_SetResult_Call {
	return &PendingRequestsInterface_SetResult_Call{Call: _e.mock.On("SetResult", ski, counter, errorResult)}
}

func (_c *PendingRequestsInterface_SetResult_Call) Run(run func(ski string, counter model.MsgCounterType, errorResult *model.ErrorType)) *PendingRequestsInterface_SetResult_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(model.MsgCounterType), args[2].(*model.ErrorType))
	})
	return _c
}

func (_c *PendingRequestsInterface_SetResult_Call) Return(_a0 *model.ErrorType) *PendingRequestsInterface_SetResult_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PendingRequestsInterface_SetResult_Call) RunAndReturn(run func(string, model.MsgCounterType, *model.ErrorType) *model.ErrorType) *PendingRequestsInterface_SetResult_Call {
	_c.Call.Return(run)
	return _c
}

// NewPendingRequestsInterface creates a new instance of PendingRequestsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPendingRequestsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *PendingRequestsInterface {
	mock := &PendingRequestsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
