// Copyright  OpenTelemetry Authors
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

package skywalkingreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/skywalkingreceiver"

import (
	"context"
	"io"

	common "skywalking.apache.org/repo/goapi/collect/common/v3"
	agent "skywalking.apache.org/repo/goapi/collect/language/agent/v3"
)

const (
	httpEventName = "http-tracing-event"
	failing       = "failing"
)

type traceSegmentReportService struct {
	agent.UnimplementedTraceSegmentReportServiceServer
}

func (s *traceSegmentReportService) Collect(stream agent.TraceSegmentReportService_CollectServer) error {
	for {
		var recData interface{}
		err := stream.RecvMsg(recData)
		if err == io.EOF {
			return stream.SendAndClose(&common.Commands{})
		}
		if err != nil {
			return err
		}

		// TODO: convert skywalking tracing to otel trace
	}
}

func (s *traceSegmentReportService) CollectInSync(ctx context.Context, segments *agent.SegmentCollection) (*common.Commands, error) {
	// TODO: convert skywalking tracing to otel trace
	return &common.Commands{}, nil
}
