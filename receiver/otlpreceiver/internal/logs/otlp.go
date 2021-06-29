// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logs

import (
	"context"

	"go.opentelemetry.io/collector/client"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/model/otlpgrpc"
	"go.opentelemetry.io/collector/model/pdata"
	"go.opentelemetry.io/collector/obsreport"
)

const (
	dataFormatProtobuf = "protobuf"
	receiverTransport  = "grpc"
)

// Receiver is the type used to handle spans from OpenTelemetry exporters.
type Receiver struct {
	id           config.ComponentID
	nextConsumer consumer.Logs
	obsrecv      *obsreport.Receiver
}

// New creates a new Receiver reference.
func New(id config.ComponentID, nextConsumer consumer.Logs) *Receiver {
	return &Receiver{
		id:           id,
		nextConsumer: nextConsumer,
		obsrecv:      obsreport.NewReceiver(obsreport.ReceiverSettings{ReceiverID: id, Transport: receiverTransport}),
	}
}

// Export implements the service Export logs func.
func (r *Receiver) Export(ctx context.Context, ld pdata.Logs) (otlpgrpc.LogsResponse, error) {
	numSpans := ld.LogRecordCount()
	if numSpans == 0 {
		return otlpgrpc.NewLogsResponse(), nil
	}

	if c, ok := client.FromGRPC(ctx); ok {
		ctx = client.NewContext(ctx, c)
	}

	ctx = r.obsrecv.StartLogsOp(obsreport.ReceiverContext(ctx, r.id, receiverTransport))
	err := r.nextConsumer.ConsumeLogs(ctx, ld)
	r.obsrecv.EndLogsOp(ctx, dataFormatProtobuf, numSpans, err)

	return otlpgrpc.NewLogsResponse(), err
}
