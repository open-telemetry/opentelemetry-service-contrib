// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"

	model "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/aerospikereceiver/internal/model"
)

// Aerospike is an autogenerated mock type for the Aerospike type
type Aerospike struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Aerospike) Close() {
	_m.Called()
}

// Info provides a mock function with given fields:
func (_m *Aerospike) Info() (*model.NodeInfo, error) {
	ret := _m.Called()

	var r0 *model.NodeInfo
	if rf, ok := ret.Get(0).(func() *model.NodeInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.NodeInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NamespaceInfo provides a mock function with given fields: namespace
func (_m *Aerospike) NamespaceInfo(namespace string) (*model.NamespaceInfo, error) {
	ret := _m.Called(namespace)

	var r0 *model.NamespaceInfo
	if rf, ok := ret.Get(0).(func(string) *model.NamespaceInfo); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.NamespaceInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAerospike creates a new instance of Aerospike. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewAerospike(t testing.TB) *Aerospike {
	mock := &Aerospike{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
