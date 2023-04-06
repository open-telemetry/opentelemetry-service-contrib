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

package lokireceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/lokireceiver"

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"context"
	"fmt"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/snappy"
	"github.com/grafana/loki/pkg/push"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/confignet"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/common/testutil"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest/plogtest"
)

func sendToCollector(endpoint string, contentType string, contentEncoding string, body []byte) error {
	var buf bytes.Buffer

	switch contentEncoding {
	case "":
		buf = *bytes.NewBuffer(body)
	case "snappy":
		if contentType == jsonContentType {
			buf = *bytes.NewBuffer(body)
		} else {
			data := snappy.Encode(nil, body)
			buf = *bytes.NewBuffer(data)
		}
	case "gzip":
		zw := gzip.NewWriter(&buf)
		if _, err := zw.Write(body); err != nil {
			return err
		}
		if err := zw.Close(); err != nil {
			return err
		}
	case "deflate":
		fw := zlib.NewWriter(&buf)
		if _, err := fw.Write(body); err != nil {
			return nil
		}
		if err := fw.Close(); err != nil {
			return err
		}
	}

	req, err := http.NewRequest("POST", endpoint, &buf)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Content-Encoding", contentEncoding)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("failed to upload logs; HTTP status code: %d", resp.StatusCode)
	}
	return nil
}

func startGRPCServer(t *testing.T) (*grpc.ClientConn, *consumertest.LogsSink) {
	config := &Config{
		Protocols: Protocols{
			GRPC: &configgrpc.GRPCServerSettings{
				NetAddr: confignet.NetAddr{
					Endpoint:  testutil.GetAvailableLocalAddress(t),
					Transport: "tcp",
				},
			},
		},
		KeepTimestamp: true,
	}
	sink := new(consumertest.LogsSink)

	set := receivertest.NewNopCreateSettings()
	lr, err := newLokiReceiver(config, sink, set)
	require.NoError(t, err)

	require.NoError(t, lr.Start(context.Background(), componenttest.NewNopHost()))
	t.Cleanup(func() { require.NoError(t, lr.Shutdown(context.Background())) })

	conn, err := grpc.Dial(config.GRPC.NetAddr.Endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	require.NoError(t, err)
	return conn, sink
}

func startHTTPServer(t *testing.T) (string, *consumertest.LogsSink) {
	addr := testutil.GetAvailableLocalAddress(t)
	config := &Config{
		Protocols: Protocols{
			HTTP: &confighttp.HTTPServerSettings{
				Endpoint: addr,
			},
		},
		KeepTimestamp: true,
	}
	sink := new(consumertest.LogsSink)

	set := receivertest.NewNopCreateSettings()
	lr, err := newLokiReceiver(config, sink, set)
	require.NoError(t, err)

	require.NoError(t, lr.Start(context.Background(), componenttest.NewNopHost()))
	t.Cleanup(func() { require.NoError(t, lr.Shutdown(context.Background())) })

	return addr, sink
}

func TestSendingProtobufPushRequestToHTTPEndpoint(t *testing.T) {
	tests := []struct {
		name            string
		contentEncoding string
		contentType     string
		body            *push.PushRequest
		expected        plog.Logs
		err             error
	}{
		{
			name:            "Sending contentEncoding=\"\" contentType=application/x-protobuf to http endpoint",
			contentEncoding: "snappy",
			contentType:     pbContentType,
			body: &push.PushRequest{
				Streams: []push.Stream{
					{
						Labels: "{foo=\"bar\"}",
						Entries: []push.Entry{
							{
								Timestamp: time.Unix(0, 1676888496000000000),
								Line:      "logline 1",
							},
						},
					},
				},
			},
			expected: generateLogs([]Log{
				{
					Timestamp: 1676888496000000000,
					Attributes: map[string]interface{}{
						"foo": "bar",
					},
					Body: pcommon.NewValueStr("logline 1"),
				},
			}),
			err: nil,
		},
	}

	// Start http server
	addr, sink := startHTTPServer(t)

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Send push request to the Loki receiver.
			_, port, _ := net.SplitHostPort(addr)
			collectorAddr := fmt.Sprintf("http://localhost:%s/loki/api/v1/push", port)

			buf, err := proto.Marshal(tt.body)
			require.NoError(t, err)

			require.NoError(t, sendToCollector(collectorAddr, tt.contentType, tt.contentEncoding, buf))
			gotLogs := sink.AllLogs()
			require.NoError(t, plogtest.CompareLogs(tt.expected, gotLogs[i], plogtest.IgnoreObservedTimestamp()))
		})
	}
}

func TestSendingPushRequestToHTTPEndpoint(t *testing.T) {
	tests := []struct {
		name            string
		contentEncoding string
		contentType     string
		body            []byte
		expected        plog.Logs
		err             error
	}{
		{
			name:            "Sending contentEncoding=\"\" contentType=application/json to http endpoint",
			contentEncoding: "",
			contentType:     jsonContentType,
			body:            []byte(`{"streams": [{"stream": {"foo": "bar"},"values": [[ "1676888496000000000", "logline 1" ], [ "1676888497000000000", "logline 2" ]]}]}`),
			expected: generateLogs([]Log{
				{
					Timestamp: 1676888496000000000,
					Attributes: map[string]interface{}{
						"foo": "bar",
					},
					Body: pcommon.NewValueStr("logline 1"),
				},
				{
					Timestamp: 1676888497000000000,
					Attributes: map[string]interface{}{
						"foo": "bar",
					},
					Body: pcommon.NewValueStr("logline 2"),
				},
			}),
			err: nil,
		},
		{
			name:            "Sending contentEncoding=\"snappy\" contentType=application/json to http endpoint",
			contentEncoding: "snappy",
			contentType:     jsonContentType,
			body:            []byte(`{"streams": [{"stream": {"foo": "bar"},"values": [[ "1676888496000000000", "logline 1" ], [ "1676888497000000000", "logline 2" ]]}]}`),
			expected: generateLogs([]Log{
				{
					Timestamp: 1676888496000000000,
					Attributes: map[string]interface{}{
						"foo": "bar",
					},
					Body: pcommon.NewValueStr("logline 1"),
				},
				{
					Timestamp: 1676888497000000000,
					Attributes: map[string]interface{}{
						"foo": "bar",
					},
					Body: pcommon.NewValueStr("logline 2"),
				},
			}),
			err: nil,
		},
		{
			name:            "Sending contentEncoding=\"gzip\" contentType=application/json to http endpoint",
			contentEncoding: "gzip",
			contentType:     jsonContentType,
			body:            []byte(`{"streams": [{"stream": {"foo": "bar"},"values": [[ "1676888496000000000", "logline 1" ], [ "1676888497000000000", "logline 2" ]]}]}`),
			expected: generateLogs([]Log{
				{
					Timestamp: 1676888496000000000,
					Attributes: map[string]interface{}{
						"foo": "bar",
					},
					Body: pcommon.NewValueStr("logline 1"),
				},
				{
					Timestamp: 1676888497000000000,
					Attributes: map[string]interface{}{
						"foo": "bar",
					},
					Body: pcommon.NewValueStr("logline 2"),
				},
			}),
			err: nil,
		},
		{
			name:            "Sending contentEncoding=\"deflate\" contentType=application/json to http endpoint",
			contentEncoding: "deflate",
			contentType:     jsonContentType,
			body:            []byte(`{"streams": [{"stream": {"foo": "bar"},"values": [[ "1676888496000000000", "logline 1" ], [ "1676888497000000000", "logline 2" ]]}]}`),
			expected: generateLogs([]Log{
				{
					Timestamp: 1676888496000000000,
					Attributes: map[string]interface{}{
						"foo": "bar",
					},
					Body: pcommon.NewValueStr("logline 1"),
				},
				{
					Timestamp: 1676888497000000000,
					Attributes: map[string]interface{}{
						"foo": "bar",
					},
					Body: pcommon.NewValueStr("logline 2"),
				},
			}),
			err: nil,
		},
	}

	// Start http server
	addr, sink := startHTTPServer(t)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Send push request to the Loki receiver.
			_, port, _ := net.SplitHostPort(addr)
			collectorAddr := fmt.Sprintf("http://localhost:%s/loki/api/v1/push", port)

			require.NoError(t, sendToCollector(collectorAddr, tt.contentType, tt.contentEncoding, tt.body), "sending logs to http endpoint shouldn't have been failed")
			gotLogs := sink.AllLogs()
			require.NoError(t, plogtest.CompareLogs(tt.expected, gotLogs[0], plogtest.IgnoreObservedTimestamp()))
			sink.Reset()
		})
	}
}

func TestSendingPushRequestToGRPCEndpoint(t *testing.T) {
	// Start grpc server
	conn, sink := startGRPCServer(t)
	defer conn.Close()
	client := push.NewPusherClient(conn)

	tests := []struct {
		name     string
		body     *push.PushRequest
		expected plog.Logs
		err      error
	}{
		{
			name: "Sending logs to grpc endpoint",
			body: &push.PushRequest{
				Streams: []push.Stream{
					{
						Labels: "{foo=\"bar\"}",
						Entries: []push.Entry{
							{
								Timestamp: time.Unix(0, 1676888496000000000),
								Line:      "logline 1",
							},
						},
					},
				},
			},
			expected: generateLogs([]Log{
				{
					Timestamp: 1676888496000000000,
					Attributes: map[string]interface{}{
						"foo": "bar",
					},
					Body: pcommon.NewValueStr("logline 1"),
				},
			}),
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := client.Push(context.Background(), tt.body)
			assert.NoError(t, err, "should not have failed to post logs")
			assert.NotNil(t, resp, "response should not have been nil")

			gotLogs := sink.AllLogs()
			require.NoError(t, plogtest.CompareLogs(tt.expected, gotLogs[i], plogtest.IgnoreObservedTimestamp()))
		})
	}
}

type Log struct {
	Timestamp  int64
	Body       pcommon.Value
	Attributes map[string]interface{}
}

func generateLogs(logs []Log) plog.Logs {
	ld := plog.NewLogs()
	logSlice := ld.ResourceLogs().AppendEmpty().ScopeLogs().AppendEmpty().LogRecords()

	for _, log := range logs {
		lr := logSlice.AppendEmpty()
		_ = lr.Attributes().FromRaw(log.Attributes)
		lr.SetTimestamp(pcommon.Timestamp(log.Timestamp))
		lr.Body().SetStr(log.Body.AsString())
	}
	return ld
}
