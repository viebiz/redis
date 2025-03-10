// Code generated by mockery. DO NOT EDIT.

package redis

import mock "github.com/stretchr/testify/mock"

// MockLimiter is an autogenerated mock type for the Limiter type
type MockLimiter struct {
	mock.Mock
}

type MockLimiter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockLimiter) EXPECT() *MockLimiter_Expecter {
	return &MockLimiter_Expecter{mock: &_m.Mock}
}

// Allow provides a mock function with no fields
func (_m *MockLimiter) Allow() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Allow")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockLimiter_Allow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Allow'
type MockLimiter_Allow_Call struct {
	*mock.Call
}

// Allow is a helper method to define mock.On call
func (_e *MockLimiter_Expecter) Allow() *MockLimiter_Allow_Call {
	return &MockLimiter_Allow_Call{Call: _e.mock.On("Allow")}
}

func (_c *MockLimiter_Allow_Call) Run(run func()) *MockLimiter_Allow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockLimiter_Allow_Call) Return(_a0 error) *MockLimiter_Allow_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLimiter_Allow_Call) RunAndReturn(run func() error) *MockLimiter_Allow_Call {
	_c.Call.Return(run)
	return _c
}

// ReportResult provides a mock function with given fields: result
func (_m *MockLimiter) ReportResult(result error) {
	_m.Called(result)
}

// MockLimiter_ReportResult_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReportResult'
type MockLimiter_ReportResult_Call struct {
	*mock.Call
}

// ReportResult is a helper method to define mock.On call
//   - result error
func (_e *MockLimiter_Expecter) ReportResult(result interface{}) *MockLimiter_ReportResult_Call {
	return &MockLimiter_ReportResult_Call{Call: _e.mock.On("ReportResult", result)}
}

func (_c *MockLimiter_ReportResult_Call) Run(run func(result error)) *MockLimiter_ReportResult_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(error))
	})
	return _c
}

func (_c *MockLimiter_ReportResult_Call) Return() *MockLimiter_ReportResult_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLimiter_ReportResult_Call) RunAndReturn(run func(error)) *MockLimiter_ReportResult_Call {
	_c.Run(run)
	return _c
}

// NewMockLimiter creates a new instance of MockLimiter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockLimiter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockLimiter {
	mock := &MockLimiter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
