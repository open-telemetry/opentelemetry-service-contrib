// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics

import (
	"context"
	"errors"
	"net"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/pmetric/pmetricotlp"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/otelarrow/admission"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/otelarrow/testdata"
)

const (
	maxBytes = 250
)

type testSink struct {
	consumertest.MetricsSink
	context.Context
	context.CancelFunc
}

func newTestSink() *testSink {
	ctx, cancel := context.WithCancel(context.Background())
	return &testSink{
		Context:    ctx,
		CancelFunc: cancel,
	}
}

func (ts *testSink) unblock() {
	time.Sleep(10 * time.Millisecond)
	ts.CancelFunc()
}

func (ts *testSink) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	<-ts.Context.Done()
	return ts.MetricsSink.ConsumeMetrics(ctx, md)
}

func TestExport_Success(t *testing.T) {
	md := testdata.GenerateMetrics(1)
	req := pmetricotlp.NewExportRequestFromMetrics(md)

	metricsSink := newTestSink()
	metricsClient, selfExp, selfProv := makeMetricsServiceClient(t, metricsSink)

	go metricsSink.unblock()
	resp, err := metricsClient.Export(context.Background(), req)
	require.NoError(t, err, "Failed to export trace: %v", err)
	require.NotNil(t, resp, "The response is missing")

	require.Len(t, metricsSink.AllMetrics(), 1)
	assert.EqualValues(t, md, metricsSink.AllMetrics()[0])

	// One self-tracing spans is issued.
	require.NoError(t, selfProv.ForceFlush(context.Background()))
	require.Len(t, selfExp.GetSpans(), 1)
}

func TestExport_EmptyRequest(t *testing.T) {
	metricsSink := newTestSink()
	metricsClient, selfExp, selfProv := makeMetricsServiceClient(t, metricsSink)
	empty := pmetricotlp.NewExportRequest()

	go metricsSink.unblock()
	resp, err := metricsClient.Export(context.Background(), empty)
	assert.NoError(t, err, "Failed to export trace: %v", err)
	assert.NotNil(t, resp, "The response is missing")

	require.Empty(t, metricsSink.AllMetrics())

	// No self-tracing spans are issued.
	require.NoError(t, selfProv.ForceFlush(context.Background()))
	require.Empty(t, selfExp.GetSpans())
}

func TestExport_ErrorConsumer(t *testing.T) {
	md := testdata.GenerateMetrics(1)
	req := pmetricotlp.NewExportRequestFromMetrics(md)

	metricsClient, selfExp, selfProv := makeMetricsServiceClient(t, consumertest.NewErr(errors.New("my error")))
	resp, err := metricsClient.Export(context.Background(), req)
	assert.EqualError(t, err, "rpc error: code = Unknown desc = my error")
	assert.Equal(t, pmetricotlp.ExportResponse{}, resp)

	// One self-tracing spans is issued.
	require.NoError(t, selfProv.ForceFlush(context.Background()))
	require.Len(t, selfExp.GetSpans(), 1)
}

func TestExport_AdmissionRequestTooLarge(t *testing.T) {
	md := testdata.GenerateMetrics(10)
	metricsSink := newTestSink()
	req := pmetricotlp.NewExportRequestFromMetrics(md)
	metricsClient, selfExp, selfProv := makeMetricsServiceClient(t, metricsSink)

	go metricsSink.unblock()
	resp, err := metricsClient.Export(context.Background(), req)
	assert.EqualError(t, err, "rpc error: code = InvalidArgument desc = rejecting request, request is too large")
	assert.Equal(t, pmetricotlp.ExportResponse{}, resp)

	// One self-tracing spans is issued.
	require.NoError(t, selfProv.ForceFlush(context.Background()))
	require.Len(t, selfExp.GetSpans(), 1)
}

func TestExport_AdmissionLimitExceeded(t *testing.T) {
	md := testdata.GenerateMetrics(1)
	metricsSink := newTestSink()
	req := pmetricotlp.NewExportRequestFromMetrics(md)

	metricsClient, selfExp, selfProv := makeMetricsServiceClient(t, metricsSink)

	var wait sync.WaitGroup
	wait.Add(10)

	var expectSuccess atomic.Int32

	for i := 0; i < 10; i++ {
		go func() {
			defer wait.Done()
			_, err := metricsClient.Export(context.Background(), req)
			if err == nil {
				// some succeed!
				expectSuccess.Add(1)
				return
			}
			assert.EqualError(t, err, "rpc error: code = ResourceExhausted desc = rejecting request, too much pending data")
		}()
	}

	metricsSink.unblock()
	wait.Wait()

	// 10 self-tracing spans are issued
	require.NoError(t, selfProv.ForceFlush(context.Background()))
	require.Len(t, selfExp.GetSpans(), 10)

	// Expect the correct number of success and failure.
	testSuccess := 0
	for _, span := range selfExp.GetSpans() {
		switch span.Status.Code {
		case codes.Ok, codes.Unset:
			testSuccess++
		}
	}
	require.Equal(t, int(expectSuccess.Load()), testSuccess)
}

func makeMetricsServiceClient(t *testing.T, mc consumer.Metrics) (pmetricotlp.GRPCClient, *tracetest.InMemoryExporter, *trace.TracerProvider) {
	addr, exp, tp := otlpReceiverOnGRPCServer(t, mc)
	cc, err := grpc.NewClient(addr.String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err, "Failed to create the TraceServiceClient: %v", err)
	t.Cleanup(func() {
		require.NoError(t, cc.Close())
	})

	return pmetricotlp.NewGRPCClient(cc), exp, tp
}

func otlpReceiverOnGRPCServer(t *testing.T, mc consumer.Metrics) (net.Addr, *tracetest.InMemoryExporter, *trace.TracerProvider) {
	ln, err := net.Listen("tcp", "localhost:")
	require.NoError(t, err, "Failed to find an available address to run the gRPC server: %v", err)

	t.Cleanup(func() {
		require.NoError(t, ln.Close())
	})

	exp := tracetest.NewInMemoryExporter()

	tp := trace.NewTracerProvider(trace.WithSyncer(exp))
	telset := componenttest.NewNopTelemetrySettings()
	telset.TracerProvider = tp

	set := receivertest.NewNopSettings()
	set.TelemetrySettings = telset

	set.ID = component.NewIDWithName(component.MustNewType("otlp"), "metrics")
	obsrecv, err := receiverhelper.NewObsReport(receiverhelper.ObsReportSettings{
		ReceiverID:             set.ID,
		Transport:              "grpc",
		ReceiverCreateSettings: set,
	})
	require.NoError(t, err)
	bq := admission.NewBoundedQueue(telset, maxBytes, 0)
	r, err := New(telset, mc, obsrecv, bq)
	require.NoError(t, err)

	// Now run it as a gRPC server
	srv := grpc.NewServer()
	pmetricotlp.RegisterGRPCServer(srv, r)
	go func() {
		_ = srv.Serve(ln)
	}()

	return ln.Addr(), exp, tp
}
