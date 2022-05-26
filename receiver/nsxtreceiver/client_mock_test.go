// Copyright  The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nsxtreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nsxtreceiver"
import (
	context "context"
	testing "testing"

	mock "github.com/stretchr/testify/mock"

	model "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nsxtreceiver/internal/model"
)

const (
	transportNode1    = "0e7bd3f2-bd49-4fa2-b650-9e9bfcdad827"
	transportNode2    = "f5045ed2-43ab-4a35-a2c5-d20c30a32292"
	transportNodeNic1 = "vmk10"
	transportNodeNic2 = "vmnic0"
	managerNode1      = "b7a79908-9808-4c9e-bb49-b70008993fcb"
	managerNodeNic1   = "eth0"
	managerNodeNic2   = "lo"
)

// MockClient is an autogenerated mock type for the MockClient type
type MockClient struct {
	mock.Mock
}

// ClusterNodes provides a mock function with given fields: ctx
func (m *MockClient) ClusterNodes(ctx context.Context) ([]model.ClusterNode, error) {
	ret := m.Called(ctx)

	var r0 []model.ClusterNode
	if rf, ok := ret.Get(0).(func(context.Context) []model.ClusterNode); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.ClusterNode)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InterfaceStatus provides a mock function with given fields: ctx, nodeID, interfaceID, class
func (m *MockClient) InterfaceStatus(ctx context.Context, nodeID string, interfaceID string, class nodeClass) (*model.NetworkInterfaceStats, error) {
	ret := m.Called(ctx, nodeID, interfaceID, class)

	var r0 *model.NetworkInterfaceStats
	if rf, ok := ret.Get(0).(func(context.Context, string, string, nodeClass) *model.NetworkInterfaceStats); ok {
		r0 = rf(ctx, nodeID, interfaceID, class)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.NetworkInterfaceStats)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, nodeClass) error); ok {
		r1 = rf(ctx, nodeID, interfaceID, class)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Interfaces provides a mock function with given fields: ctx, nodeID, class
func (m *MockClient) Interfaces(ctx context.Context, nodeID string, class nodeClass) ([]model.NetworkInterface, error) {
	ret := m.Called(ctx, nodeID, class)

	var r0 []model.NetworkInterface
	if rf, ok := ret.Get(0).(func(context.Context, string, nodeClass) []model.NetworkInterface); ok {
		r0 = rf(ctx, nodeID, class)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.NetworkInterface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, nodeClass) error); ok {
		r1 = rf(ctx, nodeID, class)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NodeStatus provides a mock function with given fields: ctx, nodeID, class
func (m *MockClient) NodeStatus(ctx context.Context, nodeID string, class nodeClass) (*model.NodeStatus, error) {
	ret := m.Called(ctx, nodeID, class)

	var r0 *model.NodeStatus
	if rf, ok := ret.Get(0).(func(context.Context, string, nodeClass) *model.NodeStatus); ok {
		r0 = rf(ctx, nodeID, class)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.NodeStatus)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, nodeClass) error); ok {
		r1 = rf(ctx, nodeID, class)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransportNodes provides a mock function with given fields: ctx
func (m *MockClient) TransportNodes(ctx context.Context) ([]model.TransportNode, error) {
	ret := m.Called(ctx)

	var r0 []model.TransportNode
	if rf, ok := ret.Get(0).(func(context.Context) []model.TransportNode); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.TransportNode)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockClient creates a new instance of MockClient. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockClient(t testing.TB) *MockClient {
	mock := &MockClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
