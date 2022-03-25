// Code generated by mockery v2.10.0. DO NOT EDIT.

package azureblobreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureblobreceiver"

import (
	context "context"

	consumer "go.opentelemetry.io/collector/consumer"

	mock "github.com/stretchr/testify/mock"
)

type MockTracesDataConsumer struct {
	mock.Mock
}

// ConsumeTracesJSON provides a mock function with given fields: ctx, json
func (_m *MockTracesDataConsumer) ConsumeTracesJSON(ctx context.Context, json []byte) error {
	ret := _m.Called(ctx, json)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte) error); ok {
		r0 = rf(ctx, json)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetNextTracesConsumer provides a mock function with given fields: nextracesConsumer
func (_m *MockTracesDataConsumer) SetNextTracesConsumer(nextracesConsumer consumer.Traces) {
	_m.Called(nextracesConsumer)
}

func NewMockTracesDataConsumer() *MockTracesDataConsumer {
	return &MockTracesDataConsumer{}
}
