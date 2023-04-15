// Copyright OpenTelemetry Authors
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

package awss3exporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awss3exporter"

import (
	"errors"

	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type Marshaler interface {
	MarshalTraces(td ptrace.Traces) ([]byte, error)
	MarshalLogs(ld plog.Logs) ([]byte, error)
	MarshalMetrics(md pmetric.Metrics) ([]byte, error)
	Format() string
}

var (
	ErrUnknownMarshaler = errors.New("unknown marshaler")
)

func NewMarshaler(name string, logger *zap.Logger) (Marshaler, error) {
	marshaler := &S3Marshaler{logger: logger}
	switch name {
	case "otlp", "otlp_proto":
		marshaler.logsMarshaler = &plog.ProtoMarshaler{}
		marshaler.tracesMarshaler = &ptrace.ProtoMarshaler{}
		marshaler.metricsMarshaler = &pmetric.ProtoMarshaler{}
		marshaler.format = "proto"
	case "otlp_json":
		marshaler.logsMarshaler = &plog.JSONMarshaler{}
		marshaler.tracesMarshaler = &ptrace.JSONMarshaler{}
		marshaler.metricsMarshaler = &pmetric.JSONMarshaler{}
		marshaler.format = "json"
	default:
		return nil, ErrUnknownMarshaler
	}
	return marshaler, nil
}
