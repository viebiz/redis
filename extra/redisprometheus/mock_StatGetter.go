// Code generated by mockery. DO NOT EDIT.

package redisprometheus

import (
	mock "github.com/stretchr/testify/mock"
	redis "github.com/viebiz/redis"
)

// MockStatGetter is an autogenerated mock type for the StatGetter type
type MockStatGetter struct {
	mock.Mock
}

type MockStatGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStatGetter) EXPECT() *MockStatGetter_Expecter {
	return &MockStatGetter_Expecter{mock: &_m.Mock}
}

// PoolStats provides a mock function with no fields
func (_m *MockStatGetter) PoolStats() *redis.PoolStats {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for PoolStats")
	}

	var r0 *redis.PoolStats
	if rf, ok := ret.Get(0).(func() *redis.PoolStats); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redis.PoolStats)
		}
	}

	return r0
}

// MockStatGetter_PoolStats_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PoolStats'
type MockStatGetter_PoolStats_Call struct {
	*mock.Call
}

// PoolStats is a helper method to define mock.On call
func (_e *MockStatGetter_Expecter) PoolStats() *MockStatGetter_PoolStats_Call {
	return &MockStatGetter_PoolStats_Call{Call: _e.mock.On("PoolStats")}
}

func (_c *MockStatGetter_PoolStats_Call) Run(run func()) *MockStatGetter_PoolStats_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStatGetter_PoolStats_Call) Return(_a0 *redis.PoolStats) *MockStatGetter_PoolStats_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStatGetter_PoolStats_Call) RunAndReturn(run func() *redis.PoolStats) *MockStatGetter_PoolStats_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockStatGetter creates a new instance of MockStatGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStatGetter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStatGetter {
	mock := &MockStatGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
