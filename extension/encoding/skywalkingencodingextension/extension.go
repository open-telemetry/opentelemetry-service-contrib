// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package skywalkingencodingextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/skywalkingencodingextension"

import (
	"context"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding"
)

const (
	skywalkingProto = "skywalking_proto"
)

var _ encoding.TracesUnmarshalerExtension = &skywalkingExtension{}
var _ ptrace.Unmarshaler = &skywalkingExtension{}

type skywalkingExtension struct {
	config      *Config
	unmarshaler ptrace.Unmarshaler
	//TODO:
	//traceMarshaler    ptrace.Marshaler
	//traceUnmarshaler  ptrace.Unmarshaler
	//logMarshaler      plog.Marshaler
	//logUnmarshaler    plog.Unmarshaler
	//metricMarshaler   pmetric.Marshaler
	//metricUnmarshaler pmetric.Unmarshaler
}

func (e *skywalkingExtension) UnmarshalTraces(buf []byte) (ptrace.Traces, error) {
	return e.unmarshaler.UnmarshalTraces(buf)
}

func (e *skywalkingExtension) Start(_ context.Context, _ component.Host) error {
	return nil
}

func (e *skywalkingExtension) Shutdown(_ context.Context) error {
	return nil
}
