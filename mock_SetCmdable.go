// Code generated by mockery. DO NOT EDIT.

package redis

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockSetCmdable is an autogenerated mock type for the SetCmdable type
type MockSetCmdable struct {
	mock.Mock
}

type MockSetCmdable_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSetCmdable) EXPECT() *MockSetCmdable_Expecter {
	return &MockSetCmdable_Expecter{mock: &_m.Mock}
}

// SAdd provides a mock function with given fields: ctx, key, members
func (_m *MockSetCmdable) SAdd(ctx context.Context, key string, members ...interface{}) *IntCmd {
	var _ca []interface{}
	_ca = append(_ca, ctx, key)
	_ca = append(_ca, members...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SAdd")
	}

	var r0 *IntCmd
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) *IntCmd); ok {
		r0 = rf(ctx, key, members...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*IntCmd)
		}
	}

	return r0
}

// MockSetCmdable_SAdd_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SAdd'
type MockSetCmdable_SAdd_Call struct {
	*mock.Call
}

// SAdd is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - members ...interface{}
func (_e *MockSetCmdable_Expecter) SAdd(ctx interface{}, key interface{}, members ...interface{}) *MockSetCmdable_SAdd_Call {
	return &MockSetCmdable_SAdd_Call{Call: _e.mock.On("SAdd",
		append([]interface{}{ctx, key}, members...)...)}
}

func (_c *MockSetCmdable_SAdd_Call) Run(run func(ctx context.Context, key string, members ...interface{})) *MockSetCmdable_SAdd_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockSetCmdable_SAdd_Call) Return(_a0 *IntCmd) *MockSetCmdable_SAdd_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSetCmdable_SAdd_Call) RunAndReturn(run func(context.Context, string, ...interface{}) *IntCmd) *MockSetCmdable_SAdd_Call {
	_c.Call.Return(run)
	return _c
}

// SCard provides a mock function with given fields: ctx, key
func (_m *MockSetCmdable) SCard(ctx context.Context, key string) *IntCmd {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for SCard")
	}

	var r0 *IntCmd
	if rf, ok := ret.Get(0).(func(context.Context, string) *IntCmd); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*IntCmd)
		}
	}

	return r0
}

// MockSetCmdable_SCard_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SCard'
type MockSetCmdable_SCard_Call struct {
	*mock.Call
}

// SCard is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
func (_e *MockSetCmdable_Expecter) SCard(ctx interface{}, key interface{}) *MockSetCmdable_SCard_Call {
	return &MockSetCmdable_SCard_Call{Call: _e.mock.On("SCard", ctx, key)}
}

func (_c *MockSetCmdable_SCard_Call) Run(run func(ctx context.Context, key string)) *MockSetCmdable_SCard_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockSetCmdable_SCard_Call) Return(_a0 *IntCmd) *MockSetCmdable_SCard_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSetCmdable_SCard_Call) RunAndReturn(run func(context.Context, string) *IntCmd) *MockSetCmdable_SCard_Call {
	_c.Call.Return(run)
	return _c
}

// SDiff provides a mock function with given fields: ctx, keys
func (_m *MockSetCmdable) SDiff(ctx context.Context, keys ...string) *StringSliceCmd {
	_va := make([]interface{}, len(keys))
	for _i := range keys {
		_va[_i] = keys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SDiff")
	}

	var r0 *StringSliceCmd
	if rf, ok := ret.Get(0).(func(context.Context, ...string) *StringSliceCmd); ok {
		r0 = rf(ctx, keys...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*StringSliceCmd)
		}
	}

	return r0
}

// MockSetCmdable_SDiff_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SDiff'
type MockSetCmdable_SDiff_Call struct {
	*mock.Call
}

// SDiff is a helper method to define mock.On call
//   - ctx context.Context
//   - keys ...string
func (_e *MockSetCmdable_Expecter) SDiff(ctx interface{}, keys ...interface{}) *MockSetCmdable_SDiff_Call {
	return &MockSetCmdable_SDiff_Call{Call: _e.mock.On("SDiff",
		append([]interface{}{ctx}, keys...)...)}
}

func (_c *MockSetCmdable_SDiff_Call) Run(run func(ctx context.Context, keys ...string)) *MockSetCmdable_SDiff_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *MockSetCmdable_SDiff_Call) Return(_a0 *StringSliceCmd) *MockSetCmdable_SDiff_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSetCmdable_SDiff_Call) RunAndReturn(run func(context.Context, ...string) *StringSliceCmd) *MockSetCmdable_SDiff_Call {
	_c.Call.Return(run)
	return _c
}

// SDiffStore provides a mock function with given fields: ctx, destination, keys
func (_m *MockSetCmdable) SDiffStore(ctx context.Context, destination string, keys ...string) *IntCmd {
	_va := make([]interface{}, len(keys))
	for _i := range keys {
		_va[_i] = keys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, destination)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SDiffStore")
	}

	var r0 *IntCmd
	if rf, ok := ret.Get(0).(func(context.Context, string, ...string) *IntCmd); ok {
		r0 = rf(ctx, destination, keys...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*IntCmd)
		}
	}

	return r0
}

// MockSetCmdable_SDiffStore_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SDiffStore'
type MockSetCmdable_SDiffStore_Call struct {
	*mock.Call
}

// SDiffStore is a helper method to define mock.On call
//   - ctx context.Context
//   - destination string
//   - keys ...string
func (_e *MockSetCmdable_Expecter) SDiffStore(ctx interface{}, destination interface{}, keys ...interface{}) *MockSetCmdable_SDiffStore_Call {
	return &MockSetCmdable_SDiffStore_Call{Call: _e.mock.On("SDiffStore",
		append([]interface{}{ctx, destination}, keys...)...)}
}

func (_c *MockSetCmdable_SDiffStore_Call) Run(run func(ctx context.Context, destination string, keys ...string)) *MockSetCmdable_SDiffStore_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockSetCmdable_SDiffStore_Call) Return(_a0 *IntCmd) *MockSetCmdable_SDiffStore_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSetCmdable_SDiffStore_Call) RunAndReturn(run func(context.Context, string, ...string) *IntCmd) *MockSetCmdable_SDiffStore_Call {
	_c.Call.Return(run)
	return _c
}

// SInter provides a mock function with given fields: ctx, keys
func (_m *MockSetCmdable) SInter(ctx context.Context, keys ...string) *StringSliceCmd {
	_va := make([]interface{}, len(keys))
	for _i := range keys {
		_va[_i] = keys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SInter")
	}

	var r0 *StringSliceCmd
	if rf, ok := ret.Get(0).(func(context.Context, ...string) *StringSliceCmd); ok {
		r0 = rf(ctx, keys...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*StringSliceCmd)
		}
	}

	return r0
}

// MockSetCmdable_SInter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SInter'
type MockSetCmdable_SInter_Call struct {
	*mock.Call
}

// SInter is a helper method to define mock.On call
//   - ctx context.Context
//   - keys ...string
func (_e *MockSetCmdable_Expecter) SInter(ctx interface{}, keys ...interface{}) *MockSetCmdable_SInter_Call {
	return &MockSetCmdable_SInter_Call{Call: _e.mock.On("SInter",
		append([]interface{}{ctx}, keys...)...)}
}

func (_c *MockSetCmdable_SInter_Call) Run(run func(ctx context.Context, keys ...string)) *MockSetCmdable_SInter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *MockSetCmdable_SInter_Call) Return(_a0 *StringSliceCmd) *MockSetCmdable_SInter_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSetCmdable_SInter_Call) RunAndReturn(run func(context.Context, ...string) *StringSliceCmd) *MockSetCmdable_SInter_Call {
	_c.Call.Return(run)
	return _c
}

// SInterCard provides a mock function with given fields: ctx, limit, keys
func (_m *MockSetCmdable) SInterCard(ctx context.Context, limit int64, keys ...string) *IntCmd {
	_va := make([]interface{}, len(keys))
	for _i := range keys {
		_va[_i] = keys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, limit)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SInterCard")
	}

	var r0 *IntCmd
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...string) *IntCmd); ok {
		r0 = rf(ctx, limit, keys...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*IntCmd)
		}
	}

	return r0
}

// MockSetCmdable_SInterCard_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SInterCard'
type MockSetCmdable_SInterCard_Call struct {
	*mock.Call
}

// SInterCard is a helper method to define mock.On call
//   - ctx context.Context
//   - limit int64
//   - keys ...string
func (_e *MockSetCmdable_Expecter) SInterCard(ctx interface{}, limit interface{}, keys ...interface{}) *MockSetCmdable_SInterCard_Call {
	return &MockSetCmdable_SInterCard_Call{Call: _e.mock.On("SInterCard",
		append([]interface{}{ctx, limit}, keys...)...)}
}

func (_c *MockSetCmdable_SInterCard_Call) Run(run func(ctx context.Context, limit int64, keys ...string)) *MockSetCmdable_SInterCard_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(context.Context), args[1].(int64), variadicArgs...)
	})
	return _c
}

func (_c *MockSetCmdable_SInterCard_Call) Return(_a0 *IntCmd) *MockSetCmdable_SInterCard_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSetCmdable_SInterCard_Call) RunAndReturn(run func(context.Context, int64, ...string) *IntCmd) *MockSetCmdable_SInterCard_Call {
	_c.Call.Return(run)
	return _c
}

// SInterStore provides a mock function with given fields: ctx, destination, keys
func (_m *MockSetCmdable) SInterStore(ctx context.Context, destination string, keys ...string) *IntCmd {
	_va := make([]interface{}, len(keys))
	for _i := range keys {
		_va[_i] = keys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, destination)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SInterStore")
	}

	var r0 *IntCmd
	if rf, ok := ret.Get(0).(func(context.Context, string, ...string) *IntCmd); ok {
		r0 = rf(ctx, destination, keys...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*IntCmd)
		}
	}

	return r0
}

// MockSetCmdable_SInterStore_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SInterStore'
type MockSetCmdable_SInterStore_Call struct {
	*mock.Call
}

// SInterStore is a helper method to define mock.On call
//   - ctx context.Context
//   - destination string
//   - keys ...string
func (_e *MockSetCmdable_Expecter) SInterStore(ctx interface{}, destination interface{}, keys ...interface{}) *MockSetCmdable_SInterStore_Call {
	return &MockSetCmdable_SInterStore_Call{Call: _e.mock.On("SInterStore",
		append([]interface{}{ctx, destination}, keys...)...)}
}

func (_c *MockSetCmdable_SInterStore_Call) Run(run func(ctx context.Context, destination string, keys ...string)) *MockSetCmdable_SInterStore_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockSetCmdable_SInterStore_Call) Return(_a0 *IntCmd) *MockSetCmdable_SInterStore_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSetCmdable_SInterStore_Call) RunAndReturn(run func(context.Context, string, ...string) *IntCmd) *MockSetCmdable_SInterStore_Call {
	_c.Call.Return(run)
	return _c
}

// SIsMember provides a mock function with given fields: ctx, key, member
func (_m *MockSetCmdable) SIsMember(ctx context.Context, key string, member interface{}) *BoolCmd {
	ret := _m.Called(ctx, key, member)

	if len(ret) == 0 {
		panic("no return value specified for SIsMember")
	}

	var r0 *BoolCmd
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) *BoolCmd); ok {
		r0 = rf(ctx, key, member)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*BoolCmd)
		}
	}

	return r0
}

// MockSetCmdable_SIsMember_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SIsMember'
type MockSetCmdable_SIsMember_Call struct {
	*mock.Call
}

// SIsMember is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - member interface{}
func (_e *MockSetCmdable_Expecter) SIsMember(ctx interface{}, key interface{}, member interface{}) *MockSetCmdable_SIsMember_Call {
	return &MockSetCmdable_SIsMember_Call{Call: _e.mock.On("SIsMember", ctx, key, member)}
}

func (_c *MockSetCmdable_SIsMember_Call) Run(run func(ctx context.Context, key string, member interface{})) *MockSetCmdable_SIsMember_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(interface{}))
	})
	return _c
}

func (_c *MockSetCmdable_SIsMember_Call) Return(_a0 *BoolCmd) *MockSetCmdable_SIsMember_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSetCmdable_SIsMember_Call) RunAndReturn(run func(context.Context, string, interface{}) *BoolCmd) *MockSetCmdable_SIsMember_Call {
	_c.Call.Return(run)
	return _c
}

// SMIsMember provides a mock function with given fields: ctx, key, members
func (_m *MockSetCmdable) SMIsMember(ctx context.Context, key string, members ...interface{}) *BoolSliceCmd {
	var _ca []interface{}
	_ca = append(_ca, ctx, key)
	_ca = append(_ca, members...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SMIsMember")
	}

	var r0 *BoolSliceCmd
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) *BoolSliceCmd); ok {
		r0 = rf(ctx, key, members...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*BoolSliceCmd)
		}
	}

	return r0
}

// MockSetCmdable_SMIsMember_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SMIsMember'
type MockSetCmdable_SMIsMember_Call struct {
	*mock.Call
}

// SMIsMember is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - members ...interface{}
func (_e *MockSetCmdable_Expecter) SMIsMember(ctx interface{}, key interface{}, members ...interface{}) *MockSetCmdable_SMIsMember_Call {
	return &MockSetCmdable_SMIsMember_Call{Call: _e.mock.On("SMIsMember",
		append([]interface{}{ctx, key}, members...)...)}
}

func (_c *MockSetCmdable_SMIsMember_Call) Run(run func(ctx context.Context, key string, members ...interface{})) *MockSetCmdable_SMIsMember_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockSetCmdable_SMIsMember_Call) Return(_a0 *BoolSliceCmd) *MockSetCmdable_SMIsMember_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSetCmdable_SMIsMember_Call) RunAndReturn(run func(context.Context, string, ...interface{}) *BoolSliceCmd) *MockSetCmdable_SMIsMember_Call {
	_c.Call.Return(run)
	return _c
}

// SMembers provides a mock function with given fields: ctx, key
func (_m *MockSetCmdable) SMembers(ctx context.Context, key string) *StringSliceCmd {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for SMembers")
	}

	var r0 *StringSliceCmd
	if rf, ok := ret.Get(0).(func(context.Context, string) *StringSliceCmd); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*StringSliceCmd)
		}
	}

	return r0
}

// MockSetCmdable_SMembers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SMembers'
type MockSetCmdable_SMembers_Call struct {
	*mock.Call
}

// SMembers is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
func (_e *MockSetCmdable_Expecter) SMembers(ctx interface{}, key interface{}) *MockSetCmdable_SMembers_Call {
	return &MockSetCmdable_SMembers_Call{Call: _e.mock.On("SMembers", ctx, key)}
}

func (_c *MockSetCmdable_SMembers_Call) Run(run func(ctx context.Context, key string)) *MockSetCmdable_SMembers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockSetCmdable_SMembers_Call) Return(_a0 *StringSliceCmd) *MockSetCmdable_SMembers_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSetCmdable_SMembers_Call) RunAndReturn(run func(context.Context, string) *StringSliceCmd) *MockSetCmdable_SMembers_Call {
	_c.Call.Return(run)
	return _c
}

// SMembersMap provides a mock function with given fields: ctx, key
func (_m *MockSetCmdable) SMembersMap(ctx context.Context, key string) *StringStructMapCmd {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for SMembersMap")
	}

	var r0 *StringStructMapCmd
	if rf, ok := ret.Get(0).(func(context.Context, string) *StringStructMapCmd); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*StringStructMapCmd)
		}
	}

	return r0
}

// MockSetCmdable_SMembersMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SMembersMap'
type MockSetCmdable_SMembersMap_Call struct {
	*mock.Call
}

// SMembersMap is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
func (_e *MockSetCmdable_Expecter) SMembersMap(ctx interface{}, key interface{}) *MockSetCmdable_SMembersMap_Call {
	return &MockSetCmdable_SMembersMap_Call{Call: _e.mock.On("SMembersMap", ctx, key)}
}

func (_c *MockSetCmdable_SMembersMap_Call) Run(run func(ctx context.Context, key string)) *MockSetCmdable_SMembersMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockSetCmdable_SMembersMap_Call) Return(_a0 *StringStructMapCmd) *MockSetCmdable_SMembersMap_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSetCmdable_SMembersMap_Call) RunAndReturn(run func(context.Context, string) *StringStructMapCmd) *MockSetCmdable_SMembersMap_Call {
	_c.Call.Return(run)
	return _c
}

// SMove provides a mock function with given fields: ctx, source, destination, member
func (_m *MockSetCmdable) SMove(ctx context.Context, source string, destination string, member interface{}) *BoolCmd {
	ret := _m.Called(ctx, source, destination, member)

	if len(ret) == 0 {
		panic("no return value specified for SMove")
	}

	var r0 *BoolCmd
	if rf, ok := ret.Get(0).(func(context.Context, string, string, interface{}) *BoolCmd); ok {
		r0 = rf(ctx, source, destination, member)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*BoolCmd)
		}
	}

	return r0
}

// MockSetCmdable_SMove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SMove'
type MockSetCmdable_SMove_Call struct {
	*mock.Call
}

// SMove is a helper method to define mock.On call
//   - ctx context.Context
//   - source string
//   - destination string
//   - member interface{}
func (_e *MockSetCmdable_Expecter) SMove(ctx interface{}, source interface{}, destination interface{}, member interface{}) *MockSetCmdable_SMove_Call {
	return &MockSetCmdable_SMove_Call{Call: _e.mock.On("SMove", ctx, source, destination, member)}
}

func (_c *MockSetCmdable_SMove_Call) Run(run func(ctx context.Context, source string, destination string, member interface{})) *MockSetCmdable_SMove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(interface{}))
	})
	return _c
}

func (_c *MockSetCmdable_SMove_Call) Return(_a0 *BoolCmd) *MockSetCmdable_SMove_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSetCmdable_SMove_Call) RunAndReturn(run func(context.Context, string, string, interface{}) *BoolCmd) *MockSetCmdable_SMove_Call {
	_c.Call.Return(run)
	return _c
}

// SPop provides a mock function with given fields: ctx, key
func (_m *MockSetCmdable) SPop(ctx context.Context, key string) *StringCmd {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for SPop")
	}

	var r0 *StringCmd
	if rf, ok := ret.Get(0).(func(context.Context, string) *StringCmd); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*StringCmd)
		}
	}

	return r0
}

// MockSetCmdable_SPop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SPop'
type MockSetCmdable_SPop_Call struct {
	*mock.Call
}

// SPop is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
func (_e *MockSetCmdable_Expecter) SPop(ctx interface{}, key interface{}) *MockSetCmdable_SPop_Call {
	return &MockSetCmdable_SPop_Call{Call: _e.mock.On("SPop", ctx, key)}
}

func (_c *MockSetCmdable_SPop_Call) Run(run func(ctx context.Context, key string)) *MockSetCmdable_SPop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockSetCmdable_SPop_Call) Return(_a0 *StringCmd) *MockSetCmdable_SPop_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSetCmdable_SPop_Call) RunAndReturn(run func(context.Context, string) *StringCmd) *MockSetCmdable_SPop_Call {
	_c.Call.Return(run)
	return _c
}

// SPopN provides a mock function with given fields: ctx, key, count
func (_m *MockSetCmdable) SPopN(ctx context.Context, key string, count int64) *StringSliceCmd {
	ret := _m.Called(ctx, key, count)

	if len(ret) == 0 {
		panic("no return value specified for SPopN")
	}

	var r0 *StringSliceCmd
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) *StringSliceCmd); ok {
		r0 = rf(ctx, key, count)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*StringSliceCmd)
		}
	}

	return r0
}

// MockSetCmdable_SPopN_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SPopN'
type MockSetCmdable_SPopN_Call struct {
	*mock.Call
}

// SPopN is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - count int64
func (_e *MockSetCmdable_Expecter) SPopN(ctx interface{}, key interface{}, count interface{}) *MockSetCmdable_SPopN_Call {
	return &MockSetCmdable_SPopN_Call{Call: _e.mock.On("SPopN", ctx, key, count)}
}

func (_c *MockSetCmdable_SPopN_Call) Run(run func(ctx context.Context, key string, count int64)) *MockSetCmdable_SPopN_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(int64))
	})
	return _c
}

func (_c *MockSetCmdable_SPopN_Call) Return(_a0 *StringSliceCmd) *MockSetCmdable_SPopN_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSetCmdable_SPopN_Call) RunAndReturn(run func(context.Context, string, int64) *StringSliceCmd) *MockSetCmdable_SPopN_Call {
	_c.Call.Return(run)
	return _c
}

// SRandMember provides a mock function with given fields: ctx, key
func (_m *MockSetCmdable) SRandMember(ctx context.Context, key string) *StringCmd {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for SRandMember")
	}

	var r0 *StringCmd
	if rf, ok := ret.Get(0).(func(context.Context, string) *StringCmd); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*StringCmd)
		}
	}

	return r0
}

// MockSetCmdable_SRandMember_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SRandMember'
type MockSetCmdable_SRandMember_Call struct {
	*mock.Call
}

// SRandMember is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
func (_e *MockSetCmdable_Expecter) SRandMember(ctx interface{}, key interface{}) *MockSetCmdable_SRandMember_Call {
	return &MockSetCmdable_SRandMember_Call{Call: _e.mock.On("SRandMember", ctx, key)}
}

func (_c *MockSetCmdable_SRandMember_Call) Run(run func(ctx context.Context, key string)) *MockSetCmdable_SRandMember_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockSetCmdable_SRandMember_Call) Return(_a0 *StringCmd) *MockSetCmdable_SRandMember_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSetCmdable_SRandMember_Call) RunAndReturn(run func(context.Context, string) *StringCmd) *MockSetCmdable_SRandMember_Call {
	_c.Call.Return(run)
	return _c
}

// SRandMemberN provides a mock function with given fields: ctx, key, count
func (_m *MockSetCmdable) SRandMemberN(ctx context.Context, key string, count int64) *StringSliceCmd {
	ret := _m.Called(ctx, key, count)

	if len(ret) == 0 {
		panic("no return value specified for SRandMemberN")
	}

	var r0 *StringSliceCmd
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) *StringSliceCmd); ok {
		r0 = rf(ctx, key, count)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*StringSliceCmd)
		}
	}

	return r0
}

// MockSetCmdable_SRandMemberN_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SRandMemberN'
type MockSetCmdable_SRandMemberN_Call struct {
	*mock.Call
}

// SRandMemberN is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - count int64
func (_e *MockSetCmdable_Expecter) SRandMemberN(ctx interface{}, key interface{}, count interface{}) *MockSetCmdable_SRandMemberN_Call {
	return &MockSetCmdable_SRandMemberN_Call{Call: _e.mock.On("SRandMemberN", ctx, key, count)}
}

func (_c *MockSetCmdable_SRandMemberN_Call) Run(run func(ctx context.Context, key string, count int64)) *MockSetCmdable_SRandMemberN_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(int64))
	})
	return _c
}

func (_c *MockSetCmdable_SRandMemberN_Call) Return(_a0 *StringSliceCmd) *MockSetCmdable_SRandMemberN_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSetCmdable_SRandMemberN_Call) RunAndReturn(run func(context.Context, string, int64) *StringSliceCmd) *MockSetCmdable_SRandMemberN_Call {
	_c.Call.Return(run)
	return _c
}

// SRem provides a mock function with given fields: ctx, key, members
func (_m *MockSetCmdable) SRem(ctx context.Context, key string, members ...interface{}) *IntCmd {
	var _ca []interface{}
	_ca = append(_ca, ctx, key)
	_ca = append(_ca, members...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SRem")
	}

	var r0 *IntCmd
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) *IntCmd); ok {
		r0 = rf(ctx, key, members...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*IntCmd)
		}
	}

	return r0
}

// MockSetCmdable_SRem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SRem'
type MockSetCmdable_SRem_Call struct {
	*mock.Call
}

// SRem is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - members ...interface{}
func (_e *MockSetCmdable_Expecter) SRem(ctx interface{}, key interface{}, members ...interface{}) *MockSetCmdable_SRem_Call {
	return &MockSetCmdable_SRem_Call{Call: _e.mock.On("SRem",
		append([]interface{}{ctx, key}, members...)...)}
}

func (_c *MockSetCmdable_SRem_Call) Run(run func(ctx context.Context, key string, members ...interface{})) *MockSetCmdable_SRem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockSetCmdable_SRem_Call) Return(_a0 *IntCmd) *MockSetCmdable_SRem_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSetCmdable_SRem_Call) RunAndReturn(run func(context.Context, string, ...interface{}) *IntCmd) *MockSetCmdable_SRem_Call {
	_c.Call.Return(run)
	return _c
}

// SScan provides a mock function with given fields: ctx, key, cursor, match, count
func (_m *MockSetCmdable) SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd {
	ret := _m.Called(ctx, key, cursor, match, count)

	if len(ret) == 0 {
		panic("no return value specified for SScan")
	}

	var r0 *ScanCmd
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64, string, int64) *ScanCmd); ok {
		r0 = rf(ctx, key, cursor, match, count)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ScanCmd)
		}
	}

	return r0
}

// MockSetCmdable_SScan_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SScan'
type MockSetCmdable_SScan_Call struct {
	*mock.Call
}

// SScan is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - cursor uint64
//   - match string
//   - count int64
func (_e *MockSetCmdable_Expecter) SScan(ctx interface{}, key interface{}, cursor interface{}, match interface{}, count interface{}) *MockSetCmdable_SScan_Call {
	return &MockSetCmdable_SScan_Call{Call: _e.mock.On("SScan", ctx, key, cursor, match, count)}
}

func (_c *MockSetCmdable_SScan_Call) Run(run func(ctx context.Context, key string, cursor uint64, match string, count int64)) *MockSetCmdable_SScan_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(uint64), args[3].(string), args[4].(int64))
	})
	return _c
}

func (_c *MockSetCmdable_SScan_Call) Return(_a0 *ScanCmd) *MockSetCmdable_SScan_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSetCmdable_SScan_Call) RunAndReturn(run func(context.Context, string, uint64, string, int64) *ScanCmd) *MockSetCmdable_SScan_Call {
	_c.Call.Return(run)
	return _c
}

// SUnion provides a mock function with given fields: ctx, keys
func (_m *MockSetCmdable) SUnion(ctx context.Context, keys ...string) *StringSliceCmd {
	_va := make([]interface{}, len(keys))
	for _i := range keys {
		_va[_i] = keys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SUnion")
	}

	var r0 *StringSliceCmd
	if rf, ok := ret.Get(0).(func(context.Context, ...string) *StringSliceCmd); ok {
		r0 = rf(ctx, keys...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*StringSliceCmd)
		}
	}

	return r0
}

// MockSetCmdable_SUnion_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SUnion'
type MockSetCmdable_SUnion_Call struct {
	*mock.Call
}

// SUnion is a helper method to define mock.On call
//   - ctx context.Context
//   - keys ...string
func (_e *MockSetCmdable_Expecter) SUnion(ctx interface{}, keys ...interface{}) *MockSetCmdable_SUnion_Call {
	return &MockSetCmdable_SUnion_Call{Call: _e.mock.On("SUnion",
		append([]interface{}{ctx}, keys...)...)}
}

func (_c *MockSetCmdable_SUnion_Call) Run(run func(ctx context.Context, keys ...string)) *MockSetCmdable_SUnion_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *MockSetCmdable_SUnion_Call) Return(_a0 *StringSliceCmd) *MockSetCmdable_SUnion_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSetCmdable_SUnion_Call) RunAndReturn(run func(context.Context, ...string) *StringSliceCmd) *MockSetCmdable_SUnion_Call {
	_c.Call.Return(run)
	return _c
}

// SUnionStore provides a mock function with given fields: ctx, destination, keys
func (_m *MockSetCmdable) SUnionStore(ctx context.Context, destination string, keys ...string) *IntCmd {
	_va := make([]interface{}, len(keys))
	for _i := range keys {
		_va[_i] = keys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, destination)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SUnionStore")
	}

	var r0 *IntCmd
	if rf, ok := ret.Get(0).(func(context.Context, string, ...string) *IntCmd); ok {
		r0 = rf(ctx, destination, keys...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*IntCmd)
		}
	}

	return r0
}

// MockSetCmdable_SUnionStore_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SUnionStore'
type MockSetCmdable_SUnionStore_Call struct {
	*mock.Call
}

// SUnionStore is a helper method to define mock.On call
//   - ctx context.Context
//   - destination string
//   - keys ...string
func (_e *MockSetCmdable_Expecter) SUnionStore(ctx interface{}, destination interface{}, keys ...interface{}) *MockSetCmdable_SUnionStore_Call {
	return &MockSetCmdable_SUnionStore_Call{Call: _e.mock.On("SUnionStore",
		append([]interface{}{ctx, destination}, keys...)...)}
}

func (_c *MockSetCmdable_SUnionStore_Call) Run(run func(ctx context.Context, destination string, keys ...string)) *MockSetCmdable_SUnionStore_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockSetCmdable_SUnionStore_Call) Return(_a0 *IntCmd) *MockSetCmdable_SUnionStore_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSetCmdable_SUnionStore_Call) RunAndReturn(run func(context.Context, string, ...string) *IntCmd) *MockSetCmdable_SUnionStore_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSetCmdable creates a new instance of MockSetCmdable. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSetCmdable(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSetCmdable {
	mock := &MockSetCmdable{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
