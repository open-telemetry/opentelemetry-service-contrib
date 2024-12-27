// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package skywalkingencodingextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/skywalkingencodingextension"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

const (
	skywalkingProto = "skywalking_proto"
)

type skywalkingExtension struct {
	config      *Config
	unmarshaler ptrace.Unmarshaler
}

func (e *skywalkingExtension) UnmarshalTraces(buf []byte) (ptrace.Traces, error) {
	return e.unmarshaler.UnmarshalTraces(buf)
}

func (e *skywalkingExtension) Start(_ context.Context, _ component.Host) error {
	e.unmarshaler = skywalkingProtobufTrace{}
	return nil
}

func (e *skywalkingExtension) Shutdown(_ context.Context) error {
	return nil
}
