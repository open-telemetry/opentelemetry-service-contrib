// Copyright 2019, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//lint:file-ignore U1000 t.Skip() flaky test causes unused function warning.

package opencensusreceiver

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	commonpb "github.com/census-instrumentation/opencensus-proto/gen-go/agent/common/v1"
	agentmetricspb "github.com/census-instrumentation/opencensus-proto/gen-go/agent/metrics/v1"
	agenttracepb "github.com/census-instrumentation/opencensus-proto/gen-go/agent/trace/v1"
	metricspb "github.com/census-instrumentation/opencensus-proto/gen-go/metrics/v1"
	tracepb "github.com/census-instrumentation/opencensus-proto/gen-go/trace/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/open-telemetry/opentelemetry-collector/component/componenttest"
	"github.com/open-telemetry/opentelemetry-collector/consumer"
	"github.com/open-telemetry/opentelemetry-collector/consumer/consumerdata"
	"github.com/open-telemetry/opentelemetry-collector/exporter/exportertest"
	"github.com/open-telemetry/opentelemetry-collector/internal"
	"github.com/open-telemetry/opentelemetry-collector/observability/observabilitytest"
	"github.com/open-telemetry/opentelemetry-collector/testutils"
)

const ocReceiver = "oc_receiver_test"

// TODO(ccaraman): Migrate tests to use assert for validating functionality.
func TestGrpcGateway_endToEnd(t *testing.T) {
	addr := testutils.GetAvailableLocalAddress(t)

	// Set the buffer count to 1 to make it flush the test span immediately.
	sink := new(exportertest.SinkTraceExporterOld)
	ocr, err := New(ocReceiver, "tcp", addr, sink, nil)
	require.NoError(t, err, "Failed to create trace receiver: %v", err)

	err = ocr.Start(context.Background(), componenttest.NewNopHost())
	require.NoError(t, err, "Failed to start trace receiver: %v", err)
	defer ocr.Shutdown(context.Background())

	// TODO(songy23): make starting server deterministic
	// Wait for the servers to start
	<-time.After(10 * time.Millisecond)

	url := fmt.Sprintf("http://%s/v1/trace", addr)

	// Verify that CORS is not enabled by default, but that it gives an 405
	// method not allowed error.
	verifyCorsResp(t, url, "origin.com", 405, false)

	traceJSON := []byte(`
    {
       "node":{"identifier":{"hostName":"testHost"}},
       "spans":[
          {
              "traceId":"W47/95gDgQPSabYzgT/GDA==",
              "spanId":"7uGbfsPBsXM=",
              "name":{"value":"testSpan"},
              "startTime":"2018-12-13T14:51:00Z",
              "endTime":"2018-12-13T14:51:01Z",
              "attributes": {
                "attributeMap": {
                  "attr1": {"intValue": "55"}
                }
              }
          }
       ]
    }`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(traceJSON))
	require.NoError(t, err, "Error creating trace POST request: %v", err)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	require.NoError(t, err, "Error posting trace to grpc-gateway server: %v", err)

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading response from trace grpc-gateway, %v", err)
	}
	respStr := string(respBytes)

	err = resp.Body.Close()
	if err != nil {
		t.Errorf("Error closing response body, %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Unexpected status from trace grpc-gateway: %v", resp.StatusCode)
	}

	if respStr != "" {
		t.Errorf("Got unexpected response from trace grpc-gateway: %v", respStr)
	}

	got := sink.AllTraces()

	want := []consumerdata.TraceData{
		{
			Node: &commonpb.Node{
				Identifier: &commonpb.ProcessIdentifier{HostName: "testHost"},
			},

			Spans: []*tracepb.Span{
				{
					TraceId:   []byte{0x5B, 0x8E, 0xFF, 0xF7, 0x98, 0x3, 0x81, 0x3, 0xD2, 0x69, 0xB6, 0x33, 0x81, 0x3F, 0xC6, 0xC},
					SpanId:    []byte{0xEE, 0xE1, 0x9B, 0x7E, 0xC3, 0xC1, 0xB1, 0x73},
					Name:      &tracepb.TruncatableString{Value: "testSpan"},
					StartTime: internal.TimeToTimestamp(time.Unix(1544712660, 0).UTC()),
					EndTime:   internal.TimeToTimestamp(time.Unix(1544712661, 0).UTC()),
					Attributes: &tracepb.Span_Attributes{
						AttributeMap: map[string]*tracepb.AttributeValue{
							"attr1": {
								Value: &tracepb.AttributeValue_IntValue{IntValue: 55},
							},
						},
					},
				},
			},
			SourceFormat: "oc_trace",
		},
	}

	assert.EqualValues(t, want, got)
}

func TestTraceGrpcGatewayCors_endToEnd(t *testing.T) {
	addr := testutils.GetAvailableLocalAddress(t)
	corsOrigins := []string{"allowed-*.com"}

	sink := new(exportertest.SinkTraceExporterOld)
	ocr, err := New(ocReceiver, "tcp", addr, sink, nil, WithCorsOrigins(corsOrigins))
	require.NoError(t, err, "Failed to create trace receiver: %v", err)
	defer ocr.Shutdown(context.Background())

	err = ocr.Start(context.Background(), componenttest.NewNopHost())
	require.NoError(t, err, "Failed to start trace receiver: %v", err)

	// TODO(songy23): make starting server deterministic
	// Wait for the servers to start
	<-time.After(10 * time.Millisecond)

	url := fmt.Sprintf("http://%s/v1/trace", addr)

	// Verify allowed domain gets responses that allow CORS.
	verifyCorsResp(t, url, "allowed-origin.com", 200, true)

	// Verify disallowed domain gets responses that disallow CORS.
	verifyCorsResp(t, url, "disallowed-origin.com", 200, false)
}

func TestMetricsGrpcGatewayCors_endToEnd(t *testing.T) {
	addr := testutils.GetAvailableLocalAddress(t)
	corsOrigins := []string{"allowed-*.com"}

	sink := new(exportertest.SinkMetricsExporterOld)
	ocr, err := New(ocReceiver, "tcp", addr, nil, sink, WithCorsOrigins(corsOrigins))
	require.NoError(t, err, "Failed to create metrics receiver: %v", err)
	defer ocr.Shutdown(context.Background())

	err = ocr.Start(context.Background(), componenttest.NewNopHost())
	require.NoError(t, err, "Failed to start metrics receiver: %v", err)

	// TODO(songy23): make starting server deterministic
	// Wait for the servers to start
	<-time.After(10 * time.Millisecond)

	url := fmt.Sprintf("http://%s/v1/metrics", addr)

	// Verify allowed domain gets responses that allow CORS.
	verifyCorsResp(t, url, "allowed-origin.com", 200, true)

	// Verify disallowed domain gets responses that disallow CORS.
	verifyCorsResp(t, url, "disallowed-origin.com", 200, false)
}

// As per Issue https://github.com/census-instrumentation/opencensus-service/issues/366
// the agent's mux should be able to accept all Proto affiliated content-types and not
// redirect them to the web-grpc-gateway endpoint.
func TestAcceptAllGRPCProtoAffiliatedContentTypes(t *testing.T) {
	t.Skip("Currently a flaky test as we need a way to flush all written traces")

	addr := testutils.GetAvailableLocalAddress(t)
	cbts := new(exportertest.SinkTraceExporterOld)
	ocr, err := New(ocReceiver, "tcp", addr, cbts, nil)
	require.NoError(t, err, "Failed to create trace receiver: %v", err)

	err = ocr.Start(context.Background(), componenttest.NewNopHost())
	require.NoError(t, err, "Failed to start the trace receiver: %v", err)
	defer ocr.Shutdown(context.Background())

	// Now start the client with the various Proto affiliated gRPC Content-SubTypes as per:
	//      https://godoc.org/google.golang.org/grpc#CallContentSubtype
	protoAffiliatedContentSubTypes := []string{"", "proto"}
	for _, subContentType := range protoAffiliatedContentSubTypes {
		if err := runContentTypeTests(addr, asSubContentType, subContentType); err != nil {
			t.Errorf("%q subContentType failed to send proto: %v", subContentType, err)
		}
	}

	// Now start the client with the various Proto affiliated gRPC Content-Types,
	// as we encountered in https://github.com/census-instrumentation/opencensus-service/issues/366
	protoAffiliatedContentTypes := []string{"application/grpc", "application/grpc+proto"}
	for _, contentType := range protoAffiliatedContentTypes {
		if err := runContentTypeTests(addr, asContentType, contentType); err != nil {
			t.Errorf("%q Content-type failed to send proto: %v", contentType, err)
		}
	}

	// Before we exit we have to verify that we got exactly 4 TraceService requests.
	wantLen := len(protoAffiliatedContentSubTypes) + len(protoAffiliatedContentTypes)
	gotReqs := cbts.AllTraces()
	if len(gotReqs) != wantLen {
		t.Errorf("Receiver ExportTraceServiceRequest length mismatch:: Got %d Want %d", len(gotReqs), wantLen)
	}
}

const (
	asSubContentType = true
	asContentType    = false
)

func runContentTypeTests(addr string, contentTypeDesignation bool, contentType string) error {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDisableRetry(),
	}

	if contentTypeDesignation == asContentType {
		opts = append(opts, grpc.WithDefaultCallOptions(
			grpc.Header(&metadata.MD{"Content-Type": []string{contentType}})))
	} else {
		opts = append(opts, grpc.WithDefaultCallOptions(grpc.CallContentSubtype(contentType)))
	}

	cc, err := grpc.Dial(addr, opts...)
	if err != nil {
		return fmt.Errorf("Creating grpc.ClientConn: %v", err)
	}
	defer cc.Close()

	// First step is to send the Node.
	acc := agenttracepb.NewTraceServiceClient(cc)

	stream, err := acc.Export(context.Background())
	if err != nil {
		return fmt.Errorf("Initializing the export stream: %v", err)
	}

	msg := &agenttracepb.ExportTraceServiceRequest{
		Node: &commonpb.Node{
			Attributes: map[string]string{
				"sub-type": contentType,
			},
		},
		Spans: []*tracepb.Span{
			{
				TraceId: []byte{
					0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
					0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10,
				},
			},
		},
	}
	return stream.Send(msg)
}

func verifyCorsResp(t *testing.T, url string, origin string, wantStatus int, wantAllowed bool) {
	req, err := http.NewRequest("OPTIONS", url, nil)
	require.NoError(t, err, "Error creating trace OPTIONS request: %v", err)
	req.Header.Set("Origin", origin)
	req.Header.Set("Access-Control-Request-Method", "POST")

	client := &http.Client{}
	resp, err := client.Do(req)
	require.NoError(t, err, "Error sending OPTIONS to grpc-gateway server: %v", err)

	err = resp.Body.Close()
	if err != nil {
		t.Errorf("Error closing OPTIONS response body, %v", err)
	}

	if resp.StatusCode != wantStatus {
		t.Errorf("Unexpected status from OPTIONS: %v", resp.StatusCode)
	}

	gotAllowOrigin := resp.Header.Get("Access-Control-Allow-Origin")
	gotAllowMethods := resp.Header.Get("Access-Control-Allow-Methods")

	wantAllowOrigin := ""
	wantAllowMethods := ""
	if wantAllowed {
		wantAllowOrigin = origin
		wantAllowMethods = "POST"
	}

	if gotAllowOrigin != wantAllowOrigin {
		t.Errorf("Unexpected Access-Control-Allow-Origin: %v", gotAllowOrigin)
	}
	if gotAllowMethods != wantAllowMethods {
		t.Errorf("Unexpected Access-Control-Allow-Methods: %v", gotAllowMethods)
	}
}

func TestStopWithoutStartNeverCrashes(t *testing.T) {
	addr := testutils.GetAvailableLocalAddress(t)
	ocr, err := New(ocReceiver, "tcp", addr, nil, nil)
	require.NoError(t, err, "Failed to create an OpenCensus receiver: %v", err)
	// Stop it before ever invoking Start*.
	ocr.stop()
}

func TestNewPortAlreadyUsed(t *testing.T) {
	addr := testutils.GetAvailableLocalAddress(t)
	ln, err := net.Listen("tcp", addr)
	require.NoError(t, err, "failed to listen on %q: %v", addr, err)
	defer ln.Close()

	r, err := New(ocReceiver, "tcp", addr, nil, nil)
	require.Error(t, err)
	require.Nil(t, r)
}

func TestMultipleStopReceptionShouldNotError(t *testing.T) {
	addr := testutils.GetAvailableLocalAddress(t)
	r, err := New(ocReceiver, "tcp", addr, new(exportertest.SinkTraceExporterOld), new(exportertest.SinkMetricsExporterOld))
	require.NoError(t, err)
	require.NotNil(t, r)

	require.NoError(t, r.Start(context.Background(), componenttest.NewNopHost()))
	require.NoError(t, r.Shutdown(context.Background()))
}

func TestStartWithoutConsumersShouldFail(t *testing.T) {
	addr := testutils.GetAvailableLocalAddress(t)
	r, err := New(ocReceiver, "tcp", addr, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, r)

	require.Error(t, r.Start(context.Background(), componenttest.NewNopHost()))
}

func tempSocketName(t *testing.T) string {
	tmpfile, err := ioutil.TempFile("", "sock")
	require.NoError(t, err)
	socket := tmpfile.Name()
	err = os.Remove(socket)
	require.NoError(t, err)

	return socket
}

func TestReceiveOnUnixDomainSocket_endToEnd(t *testing.T) {
	socketName := tempSocketName(t)
	cbts := new(exportertest.SinkTraceExporterOld)
	r, err := New(ocReceiver, "unix", socketName, cbts, nil)
	require.NoError(t, err)
	require.NotNil(t, r)
	require.NoError(t, r.Start(context.Background(), componenttest.NewNopHost()))
	defer r.Shutdown(context.Background())

	// Wait for the servers to start
	<-time.After(10 * time.Millisecond)

	span := `
{
 "node": {
 },
 "spans": [
   {
     "trace_id": "YpsR8/le4OgjwSSxhjlrEg==",
     "span_id": "2CogcbJh7Ko=",
     "socket": {
       "value": "/abc",
       "truncated_byte_count": 0
     },
     "kind": "SPAN_KIND_UNSPECIFIED",
     "start_time": "2020-01-09T11:13:53.187Z",
     "end_time": "2020-01-09T11:13:53.187Z"
	}
 ]
}
`
	c := http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (conn net.Conn, err error) {
				return net.Dial("unix", socketName)
			},
		},
	}

	response, err := c.Post("http://unix/v1/trace", "application/json", strings.NewReader(span))
	require.NoError(t, err)
	defer response.Body.Close()

	require.Equal(t, 200, response.StatusCode)
}

// TestOCReceiverTrace_HandleNextConsumerResponse checks if the trace receiver
// is returning the proper response (return and metrics) when the next consumer
// in the pipeline reports error. The test changes the responses returned by the
// next trace consumer, checks if data was passed down the pipeline and if
// proper metrics were recorded. It also uses all endpoints supported by the
// trace receiver.
func TestOCReceiverTrace_HandleNextConsumerResponse(t *testing.T) {
	type ingestionStateTest struct {
		okToIngest   bool
		expectedCode codes.Code
	}
	tests := []struct {
		name                         string
		expectedReceivedBatches      int
		expectedIngestionBlockedRPCs int
		ingestionStates              []ingestionStateTest
	}{
		{
			name:                         "IngestTest",
			expectedReceivedBatches:      2,
			expectedIngestionBlockedRPCs: 1,
			ingestionStates: []ingestionStateTest{
				{
					okToIngest:   true,
					expectedCode: codes.OK,
				},
				{
					okToIngest:   false,
					expectedCode: codes.Unknown,
				},
				{
					okToIngest:   true,
					expectedCode: codes.OK,
				},
			},
		},
	}

	addr := testutils.GetAvailableLocalAddress(t)
	msg := &agenttracepb.ExportTraceServiceRequest{
		Node: &commonpb.Node{
			ServiceInfo: &commonpb.ServiceInfo{Name: "test-svc"},
		},
		Spans: []*tracepb.Span{
			{
				TraceId: []byte{
					0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
					0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10,
				},
			},
		},
	}

	exportBidiFn := func(
		t *testing.T,
		cc *grpc.ClientConn,
		msg *agenttracepb.ExportTraceServiceRequest) error {

		acc := agenttracepb.NewTraceServiceClient(cc)
		stream, err := acc.Export(context.Background())
		require.NoError(t, err)
		require.NotNil(t, stream)

		err = stream.Send(msg)
		stream.CloseSend()
		if err == nil {
			for {
				if _, err = stream.Recv(); err != nil {
					if err == io.EOF {
						err = nil
					}
					break
				}
			}
		}

		return err
	}

	exporters := []struct {
		receiverTag string
		exportFn    func(
			t *testing.T,
			cc *grpc.ClientConn,
			msg *agenttracepb.ExportTraceServiceRequest) error
	}{
		{
			receiverTag: "oc_trace",
			exportFn:    exportBidiFn,
		},
	}
	for _, exporter := range exporters {
		for _, tt := range tests {
			t.Run(tt.name+"/"+exporter.receiverTag, func(t *testing.T) {
				doneFn := observabilitytest.SetupRecordedMetricsTest()
				defer doneFn()

				sink := new(sinkTraceConsumer)

				var opts []Option
				ocr, err := New(ocReceiver, "tcp", addr, nil, nil, opts...)
				require.Nil(t, err)
				require.NotNil(t, ocr)

				ocr.traceConsumer = sink
				require.NoError(t, ocr.Start(context.Background(), componenttest.NewNopHost()))
				defer ocr.Shutdown(context.Background())

				cc, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
				if err != nil {
					t.Errorf("grpc.Dial: %v", err)
				}
				defer cc.Close()

				for _, ingestionState := range tt.ingestionStates {
					if ingestionState.okToIngest {
						sink.SetConsumeTraceError(nil)
					} else {
						sink.SetConsumeTraceError(fmt.Errorf("%q: consumer error", tt.name))
					}

					err = exporter.exportFn(t, cc, msg)

					status, ok := status.FromError(err)
					require.True(t, ok)
					assert.Equal(t, ingestionState.expectedCode, status.Code())
				}

				require.Equal(t, tt.expectedReceivedBatches, len(sink.AllTraces()))
				require.Nil(
					t,
					observabilitytest.CheckValueViewReceiverReceivedSpans(
						exporter.receiverTag,
						tt.expectedReceivedBatches),
				)
			})
		}
	}
}

// TestOCReceiverMetrics_HandleNextConsumerResponse checks if the metrics receiver
// is returning the proper response (return and metrics) when the next consumer
// in the pipeline reports error. The test changes the responses returned by the
// next trace consumer, checks if data was passed down the pipeline and if
// proper metrics were recorded. It also uses all endpoints supported by the
// metrics receiver.
func TestOCReceiverMetrics_HandleNextConsumerResponse(t *testing.T) {
	type ingestionStateTest struct {
		okToIngest   bool
		expectedCode codes.Code
	}
	tests := []struct {
		name                         string
		expectedReceivedBatches      int
		expectedIngestionBlockedRPCs int
		ingestionStates              []ingestionStateTest
	}{
		{
			name:                         "IngestTest",
			expectedReceivedBatches:      2,
			expectedIngestionBlockedRPCs: 1,
			ingestionStates: []ingestionStateTest{
				{
					okToIngest:   true,
					expectedCode: codes.OK,
				},
				{
					okToIngest:   false,
					expectedCode: codes.Unknown,
				},
				{
					okToIngest:   true,
					expectedCode: codes.OK,
				},
			},
		},
	}

	descriptor := &metricspb.MetricDescriptor{
		Name:        "testMetric",
		Description: "metric descriptor",
		Unit:        "1",
		Type:        metricspb.MetricDescriptor_GAUGE_INT64,
	}
	point := &metricspb.Point{
		Timestamp: internal.TimeToTimestamp(time.Now().UTC()),
		Value: &metricspb.Point_Int64Value{
			Int64Value: int64(1),
		},
	}
	ts := &metricspb.TimeSeries{
		Points: []*metricspb.Point{point},
	}
	metric := &metricspb.Metric{
		MetricDescriptor: descriptor,
		Timeseries:       []*metricspb.TimeSeries{ts},
	}

	addr := testutils.GetAvailableLocalAddress(t)
	msg := &agentmetricspb.ExportMetricsServiceRequest{
		Node: &commonpb.Node{
			ServiceInfo: &commonpb.ServiceInfo{Name: "test-svc"},
		},
		Metrics: []*metricspb.Metric{metric},
	}

	exportBidiFn := func(
		t *testing.T,
		cc *grpc.ClientConn,
		msg *agentmetricspb.ExportMetricsServiceRequest) error {

		acc := agentmetricspb.NewMetricsServiceClient(cc)
		stream, err := acc.Export(context.Background())
		require.NoError(t, err)
		require.NotNil(t, stream)

		err = stream.Send(msg)
		stream.CloseSend()
		if err == nil {
			for {
				if _, err = stream.Recv(); err != nil {
					if err == io.EOF {
						err = nil
					}
					break
				}
			}
		}

		return err
	}

	exporters := []struct {
		receiverTag string
		exportFn    func(
			t *testing.T,
			cc *grpc.ClientConn,
			msg *agentmetricspb.ExportMetricsServiceRequest) error
	}{
		{
			receiverTag: "oc_metrics",
			exportFn:    exportBidiFn,
		},
	}
	for _, exporter := range exporters {
		for _, tt := range tests {
			t.Run(tt.name+"/"+exporter.receiverTag, func(t *testing.T) {
				doneFn := observabilitytest.SetupRecordedMetricsTest()
				defer doneFn()

				sink := new(sinkMetricsConsumer)

				var opts []Option
				ocr, err := New(ocReceiver, "tcp", addr, nil, nil, opts...)
				require.Nil(t, err)
				require.NotNil(t, ocr)

				ocr.metricsConsumer = sink
				require.Nil(t, ocr.Start(context.Background(), componenttest.NewNopHost()))
				defer ocr.Shutdown(context.Background())

				cc, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
				if err != nil {
					t.Errorf("grpc.Dial: %v", err)
				}
				defer cc.Close()

				for _, ingestionState := range tt.ingestionStates {
					if ingestionState.okToIngest {
						sink.SetConsumeMetricsError(nil)
					} else {
						sink.SetConsumeMetricsError(fmt.Errorf("%q: consumer error", tt.name))
					}

					err = exporter.exportFn(t, cc, msg)

					status, ok := status.FromError(err)
					require.True(t, ok)
					assert.Equal(t, ingestionState.expectedCode, status.Code())
				}

				require.Equal(t, tt.expectedReceivedBatches, len(sink.AllMetrics()))
				require.Nil(
					t,
					observabilitytest.CheckValueViewReceiverReceivedTimeSeries(
						exporter.receiverTag,
						tt.expectedReceivedBatches),
				)
			})
		}
	}
}

type sinkTraceConsumer struct {
	// consumeTraceError is the error to be returned when ConsumeTraceData is
	// called
	consumeTraceError error

	traces []consumerdata.TraceData
}

var _ consumer.TraceConsumerOld = (*sinkTraceConsumer)(nil)

func (stc *sinkTraceConsumer) ConsumeTraceData(ctx context.Context, td consumerdata.TraceData) error {
	if stc.consumeTraceError == nil {
		stc.traces = append(stc.traces, td)
	}
	return stc.consumeTraceError
}

func (stc *sinkTraceConsumer) SetConsumeTraceError(err error) {
	stc.consumeTraceError = err
}

func (stc *sinkTraceConsumer) AllTraces() []consumerdata.TraceData {
	return stc.traces[:]
}

type sinkMetricsConsumer struct {
	// consumeMetricsError is the error to be returned when ConsumeMetricsData is
	// called
	consumeMetricsError error

	metrics []consumerdata.MetricsData
}

var _ consumer.MetricsConsumerOld = (*sinkMetricsConsumer)(nil)

func (smc *sinkMetricsConsumer) ConsumeMetricsData(ctx context.Context, md consumerdata.MetricsData) error {
	if smc.consumeMetricsError == nil {
		smc.metrics = append(smc.metrics, md)
	}
	return smc.consumeMetricsError
}

func (smc *sinkMetricsConsumer) SetConsumeMetricsError(err error) {
	smc.consumeMetricsError = err
}

func (smc *sinkMetricsConsumer) AllMetrics() []consumerdata.MetricsData {
	return smc.metrics[:]
}
