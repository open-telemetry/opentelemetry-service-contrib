// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package trace

import (
	"context"
	"errors"
	"net"
	"testing"

	"github.com/open-telemetry/otel-arrow/collector/testdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/pdata/ptrace/ptraceotlp"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestExport(t *testing.T) {
	td := testdata.GenerateTraces(1)
	req := ptraceotlp.NewExportRequestFromTraces(td)

	traceSink := new(consumertest.TracesSink)
	traceClient := makeTraceServiceClient(t, traceSink)
	resp, err := traceClient.Export(context.Background(), req)
	require.NoError(t, err, "Failed to export trace: %v", err)
	require.NotNil(t, resp, "The response is missing")

	require.Len(t, traceSink.AllTraces(), 1)
	assert.EqualValues(t, td, traceSink.AllTraces()[0])
}

func TestExport_EmptyRequest(t *testing.T) {
	traceSink := new(consumertest.TracesSink)
	traceClient := makeTraceServiceClient(t, traceSink)
	resp, err := traceClient.Export(context.Background(), ptraceotlp.NewExportRequest())
	assert.NoError(t, err, "Failed to export trace: %v", err)
	assert.NotNil(t, resp, "The response is missing")
}

func TestExport_ErrorConsumer(t *testing.T) {
	td := testdata.GenerateTraces(1)
	req := ptraceotlp.NewExportRequestFromTraces(td)

	traceClient := makeTraceServiceClient(t, consumertest.NewErr(errors.New("my error")))
	resp, err := traceClient.Export(context.Background(), req)
	assert.EqualError(t, err, "rpc error: code = Unknown desc = my error")
	assert.Equal(t, ptraceotlp.ExportResponse{}, resp)
}

func makeTraceServiceClient(t *testing.T, tc consumer.Traces) ptraceotlp.GRPCClient {
	addr := otlpReceiverOnGRPCServer(t, tc)
	cc, err := grpc.NewClient(addr.String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err, "Failed to create the TraceServiceClient: %v", err)
	t.Cleanup(func() {
		require.NoError(t, cc.Close())
	})

	return ptraceotlp.NewGRPCClient(cc)
}

func otlpReceiverOnGRPCServer(t *testing.T, tc consumer.Traces) net.Addr {
	ln, err := net.Listen("tcp", "localhost:")
	require.NoError(t, err, "Failed to find an available address to run the gRPC server: %v", err)

	t.Cleanup(func() {
		require.NoError(t, ln.Close())
	})

	set := receivertest.NewNopCreateSettings()
	set.ID = component.NewIDWithName(component.MustNewType("otlp"), "trace")
	obsrecv, err := receiverhelper.NewObsReport(receiverhelper.ObsReportSettings{
		ReceiverID:             set.ID,
		Transport:              "grpc",
		ReceiverCreateSettings: set,
	})
	require.NoError(t, err)
	r := New(tc, obsrecv)
	// Now run it as a gRPC server
	srv := grpc.NewServer()
	ptraceotlp.RegisterGRPCServer(srv, r)
	go func() {
		_ = srv.Serve(ln)
	}()

	return ln.Addr()
}
