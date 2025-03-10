// Code generated by mockery. DO NOT EDIT.

package redisotel

import mock "github.com/stretchr/testify/mock"

// Mockoption is an autogenerated mock type for the option type
type Mockoption struct {
	mock.Mock
}

type Mockoption_Expecter struct {
	mock *mock.Mock
}

func (_m *Mockoption) EXPECT() *Mockoption_Expecter {
	return &Mockoption_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: conf
func (_m *Mockoption) Execute(conf *config) {
	_m.Called(conf)
}

// Mockoption_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type Mockoption_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - conf *config
func (_e *Mockoption_Expecter) Execute(conf interface{}) *Mockoption_Execute_Call {
	return &Mockoption_Execute_Call{Call: _e.mock.On("Execute", conf)}
}

func (_c *Mockoption_Execute_Call) Run(run func(conf *config)) *Mockoption_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*config))
	})
	return _c
}

func (_c *Mockoption_Execute_Call) Return() *Mockoption_Execute_Call {
	_c.Call.Return()
	return _c
}

func (_c *Mockoption_Execute_Call) RunAndReturn(run func(*config)) *Mockoption_Execute_Call {
	_c.Run(run)
	return _c
}

// NewMockoption creates a new instance of Mockoption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockoption(t interface {
	mock.TestingT
	Cleanup(func())
}) *Mockoption {
	mock := &Mockoption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
