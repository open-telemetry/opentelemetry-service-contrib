// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Aerospike is an autogenerated mock type for the Aerospike type
type Aerospike struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Aerospike) Close() {
	_m.Called()
}

// Info provides a mock function with given fields:
func (_m *Aerospike) Info() map[string]map[string]string {
	ret := _m.Called()

	var r0 map[string]map[string]string
	if rf, ok := ret.Get(0).(func() map[string]map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]map[string]string)
		}
	}

	return r0
}

// NamespaceInfo provides a mock function with given fields:
func (_m *Aerospike) NamespaceInfo() map[string]map[string]map[string]string {
	ret := _m.Called()

	var r0 map[string]map[string]map[string]string
	if rf, ok := ret.Get(0).(func() map[string]map[string]map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]map[string]map[string]string)
		}
	}

	return r0
}

type mockConstructorTestingTNewAerospike interface {
	mock.TestingT
	Cleanup(func())
}

// NewAerospike creates a new instance of Aerospike. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAerospike(t mockConstructorTestingTNewAerospike) *Aerospike {
	mock := &Aerospike{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
