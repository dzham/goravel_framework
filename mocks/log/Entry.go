// Code generated by mockery. DO NOT EDIT.

package log

import (
	context "context"

	log "github.com/goravel/framework/contracts/log"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Entry is an autogenerated mock type for the Entry type
type Entry struct {
	mock.Mock
}

type Entry_Expecter struct {
	mock *mock.Mock
}

func (_m *Entry) EXPECT() *Entry_Expecter {
	return &Entry_Expecter{mock: &_m.Mock}
}

// Code provides a mock function with given fields:
func (_m *Entry) Code() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Code")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Entry_Code_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Code'
type Entry_Code_Call struct {
	*mock.Call
}

// Code is a helper method to define mock.On call
func (_e *Entry_Expecter) Code() *Entry_Code_Call {
	return &Entry_Code_Call{Call: _e.mock.On("Code")}
}

func (_c *Entry_Code_Call) Run(run func()) *Entry_Code_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Entry_Code_Call) Return(_a0 string) *Entry_Code_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Entry_Code_Call) RunAndReturn(run func() string) *Entry_Code_Call {
	_c.Call.Return(run)
	return _c
}

// Context provides a mock function with given fields:
func (_m *Entry) Context() context.Context {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Context")
	}

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// Entry_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type Entry_Context_Call struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *Entry_Expecter) Context() *Entry_Context_Call {
	return &Entry_Context_Call{Call: _e.mock.On("Context")}
}

func (_c *Entry_Context_Call) Run(run func()) *Entry_Context_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Entry_Context_Call) Return(_a0 context.Context) *Entry_Context_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Entry_Context_Call) RunAndReturn(run func() context.Context) *Entry_Context_Call {
	_c.Call.Return(run)
	return _c
}

// Data provides a mock function with given fields:
func (_m *Entry) Data() log.Data {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Data")
	}

	var r0 log.Data
	if rf, ok := ret.Get(0).(func() log.Data); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(log.Data)
		}
	}

	return r0
}

// Entry_Data_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Data'
type Entry_Data_Call struct {
	*mock.Call
}

// Data is a helper method to define mock.On call
func (_e *Entry_Expecter) Data() *Entry_Data_Call {
	return &Entry_Data_Call{Call: _e.mock.On("Data")}
}

func (_c *Entry_Data_Call) Run(run func()) *Entry_Data_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Entry_Data_Call) Return(_a0 log.Data) *Entry_Data_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Entry_Data_Call) RunAndReturn(run func() log.Data) *Entry_Data_Call {
	_c.Call.Return(run)
	return _c
}

// Domain provides a mock function with given fields:
func (_m *Entry) Domain() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Domain")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Entry_Domain_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Domain'
type Entry_Domain_Call struct {
	*mock.Call
}

// Domain is a helper method to define mock.On call
func (_e *Entry_Expecter) Domain() *Entry_Domain_Call {
	return &Entry_Domain_Call{Call: _e.mock.On("Domain")}
}

func (_c *Entry_Domain_Call) Run(run func()) *Entry_Domain_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Entry_Domain_Call) Return(_a0 string) *Entry_Domain_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Entry_Domain_Call) RunAndReturn(run func() string) *Entry_Domain_Call {
	_c.Call.Return(run)
	return _c
}

// Hint provides a mock function with given fields:
func (_m *Entry) Hint() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Hint")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Entry_Hint_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Hint'
type Entry_Hint_Call struct {
	*mock.Call
}

// Hint is a helper method to define mock.On call
func (_e *Entry_Expecter) Hint() *Entry_Hint_Call {
	return &Entry_Hint_Call{Call: _e.mock.On("Hint")}
}

func (_c *Entry_Hint_Call) Run(run func()) *Entry_Hint_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Entry_Hint_Call) Return(_a0 string) *Entry_Hint_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Entry_Hint_Call) RunAndReturn(run func() string) *Entry_Hint_Call {
	_c.Call.Return(run)
	return _c
}

// Level provides a mock function with given fields:
func (_m *Entry) Level() log.Level {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Level")
	}

	var r0 log.Level
	if rf, ok := ret.Get(0).(func() log.Level); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(log.Level)
	}

	return r0
}

// Entry_Level_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Level'
type Entry_Level_Call struct {
	*mock.Call
}

// Level is a helper method to define mock.On call
func (_e *Entry_Expecter) Level() *Entry_Level_Call {
	return &Entry_Level_Call{Call: _e.mock.On("Level")}
}

func (_c *Entry_Level_Call) Run(run func()) *Entry_Level_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Entry_Level_Call) Return(_a0 log.Level) *Entry_Level_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Entry_Level_Call) RunAndReturn(run func() log.Level) *Entry_Level_Call {
	_c.Call.Return(run)
	return _c
}

// Message provides a mock function with given fields:
func (_m *Entry) Message() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Message")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Entry_Message_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Message'
type Entry_Message_Call struct {
	*mock.Call
}

// Message is a helper method to define mock.On call
func (_e *Entry_Expecter) Message() *Entry_Message_Call {
	return &Entry_Message_Call{Call: _e.mock.On("Message")}
}

func (_c *Entry_Message_Call) Run(run func()) *Entry_Message_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Entry_Message_Call) Return(_a0 string) *Entry_Message_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Entry_Message_Call) RunAndReturn(run func() string) *Entry_Message_Call {
	_c.Call.Return(run)
	return _c
}

// Owner provides a mock function with given fields:
func (_m *Entry) Owner() any {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Owner")
	}

	var r0 any
	if rf, ok := ret.Get(0).(func() any); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(any)
		}
	}

	return r0
}

// Entry_Owner_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Owner'
type Entry_Owner_Call struct {
	*mock.Call
}

// Owner is a helper method to define mock.On call
func (_e *Entry_Expecter) Owner() *Entry_Owner_Call {
	return &Entry_Owner_Call{Call: _e.mock.On("Owner")}
}

func (_c *Entry_Owner_Call) Run(run func()) *Entry_Owner_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Entry_Owner_Call) Return(_a0 any) *Entry_Owner_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Entry_Owner_Call) RunAndReturn(run func() any) *Entry_Owner_Call {
	_c.Call.Return(run)
	return _c
}

// Request provides a mock function with given fields:
func (_m *Entry) Request() map[string]any {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Request")
	}

	var r0 map[string]any
	if rf, ok := ret.Get(0).(func() map[string]any); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]any)
		}
	}

	return r0
}

// Entry_Request_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Request'
type Entry_Request_Call struct {
	*mock.Call
}

// Request is a helper method to define mock.On call
func (_e *Entry_Expecter) Request() *Entry_Request_Call {
	return &Entry_Request_Call{Call: _e.mock.On("Request")}
}

func (_c *Entry_Request_Call) Run(run func()) *Entry_Request_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Entry_Request_Call) Return(_a0 map[string]any) *Entry_Request_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Entry_Request_Call) RunAndReturn(run func() map[string]any) *Entry_Request_Call {
	_c.Call.Return(run)
	return _c
}

// Response provides a mock function with given fields:
func (_m *Entry) Response() map[string]any {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Response")
	}

	var r0 map[string]any
	if rf, ok := ret.Get(0).(func() map[string]any); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]any)
		}
	}

	return r0
}

// Entry_Response_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Response'
type Entry_Response_Call struct {
	*mock.Call
}

// Response is a helper method to define mock.On call
func (_e *Entry_Expecter) Response() *Entry_Response_Call {
	return &Entry_Response_Call{Call: _e.mock.On("Response")}
}

func (_c *Entry_Response_Call) Run(run func()) *Entry_Response_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Entry_Response_Call) Return(_a0 map[string]any) *Entry_Response_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Entry_Response_Call) RunAndReturn(run func() map[string]any) *Entry_Response_Call {
	_c.Call.Return(run)
	return _c
}

// Tags provides a mock function with given fields:
func (_m *Entry) Tags() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Tags")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// Entry_Tags_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Tags'
type Entry_Tags_Call struct {
	*mock.Call
}

// Tags is a helper method to define mock.On call
func (_e *Entry_Expecter) Tags() *Entry_Tags_Call {
	return &Entry_Tags_Call{Call: _e.mock.On("Tags")}
}

func (_c *Entry_Tags_Call) Run(run func()) *Entry_Tags_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Entry_Tags_Call) Return(_a0 []string) *Entry_Tags_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Entry_Tags_Call) RunAndReturn(run func() []string) *Entry_Tags_Call {
	_c.Call.Return(run)
	return _c
}

// Time provides a mock function with given fields:
func (_m *Entry) Time() time.Time {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Time")
	}

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// Entry_Time_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Time'
type Entry_Time_Call struct {
	*mock.Call
}

// Time is a helper method to define mock.On call
func (_e *Entry_Expecter) Time() *Entry_Time_Call {
	return &Entry_Time_Call{Call: _e.mock.On("Time")}
}

func (_c *Entry_Time_Call) Run(run func()) *Entry_Time_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Entry_Time_Call) Return(_a0 time.Time) *Entry_Time_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Entry_Time_Call) RunAndReturn(run func() time.Time) *Entry_Time_Call {
	_c.Call.Return(run)
	return _c
}

// Trace provides a mock function with given fields:
func (_m *Entry) Trace() map[string]any {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Trace")
	}

	var r0 map[string]any
	if rf, ok := ret.Get(0).(func() map[string]any); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]any)
		}
	}

	return r0
}

// Entry_Trace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Trace'
type Entry_Trace_Call struct {
	*mock.Call
}

// Trace is a helper method to define mock.On call
func (_e *Entry_Expecter) Trace() *Entry_Trace_Call {
	return &Entry_Trace_Call{Call: _e.mock.On("Trace")}
}

func (_c *Entry_Trace_Call) Run(run func()) *Entry_Trace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Entry_Trace_Call) Return(_a0 map[string]any) *Entry_Trace_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Entry_Trace_Call) RunAndReturn(run func() map[string]any) *Entry_Trace_Call {
	_c.Call.Return(run)
	return _c
}

// User provides a mock function with given fields:
func (_m *Entry) User() any {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for User")
	}

	var r0 any
	if rf, ok := ret.Get(0).(func() any); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(any)
		}
	}

	return r0
}

// Entry_User_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'User'
type Entry_User_Call struct {
	*mock.Call
}

// User is a helper method to define mock.On call
func (_e *Entry_Expecter) User() *Entry_User_Call {
	return &Entry_User_Call{Call: _e.mock.On("User")}
}

func (_c *Entry_User_Call) Run(run func()) *Entry_User_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Entry_User_Call) Return(_a0 any) *Entry_User_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Entry_User_Call) RunAndReturn(run func() any) *Entry_User_Call {
	_c.Call.Return(run)
	return _c
}

// With provides a mock function with given fields:
func (_m *Entry) With() map[string]any {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for With")
	}

	var r0 map[string]any
	if rf, ok := ret.Get(0).(func() map[string]any); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]any)
		}
	}

	return r0
}

// Entry_With_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'With'
type Entry_With_Call struct {
	*mock.Call
}

// With is a helper method to define mock.On call
func (_e *Entry_Expecter) With() *Entry_With_Call {
	return &Entry_With_Call{Call: _e.mock.On("With")}
}

func (_c *Entry_With_Call) Run(run func()) *Entry_With_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Entry_With_Call) Return(_a0 map[string]any) *Entry_With_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Entry_With_Call) RunAndReturn(run func() map[string]any) *Entry_With_Call {
	_c.Call.Return(run)
	return _c
}

// NewEntry creates a new instance of Entry. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEntry(t interface {
	mock.TestingT
	Cleanup(func())
}) *Entry {
	mock := &Entry{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
