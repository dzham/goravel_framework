// Code generated by mockery. DO NOT EDIT.

package testing

import (
	testing "github.com/goravel/framework/contracts/testing"
	mock "github.com/stretchr/testify/mock"
)

// Testing is an autogenerated mock type for the Testing type
type Testing struct {
	mock.Mock
}

type Testing_Expecter struct {
	mock *mock.Mock
}

func (_m *Testing) EXPECT() *Testing_Expecter {
	return &Testing_Expecter{mock: &_m.Mock}
}

// Docker provides a mock function with no fields
func (_m *Testing) Docker() testing.Docker {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Docker")
	}

	var r0 testing.Docker
	if rf, ok := ret.Get(0).(func() testing.Docker); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(testing.Docker)
		}
	}

	return r0
}

// Testing_Docker_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Docker'
type Testing_Docker_Call struct {
	*mock.Call
}

// Docker is a helper method to define mock.On call
func (_e *Testing_Expecter) Docker() *Testing_Docker_Call {
	return &Testing_Docker_Call{Call: _e.mock.On("Docker")}
}

func (_c *Testing_Docker_Call) Run(run func()) *Testing_Docker_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Testing_Docker_Call) Return(_a0 testing.Docker) *Testing_Docker_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Testing_Docker_Call) RunAndReturn(run func() testing.Docker) *Testing_Docker_Call {
	_c.Call.Return(run)
	return _c
}

// NewTesting creates a new instance of Testing. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTesting(t interface {
	mock.TestingT
	Cleanup(func())
}) *Testing {
	mock := &Testing{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
