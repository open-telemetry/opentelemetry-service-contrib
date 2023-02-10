// Copyright The OpenTelemetry Authors
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

package logs

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.uber.org/zap/zaptest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/internal/testutil"
)

func TestSubmitLogs(t *testing.T) {
	logger := zaptest.NewLogger(t)

	tests := []struct {
		name        string
		payload     []datadogV2.HTTPLogItem
		testFn      func(jsonLogs testutil.JSONLogs, counter int)
		numRequests int
	}{
		{
			name: "same-tags",
			payload: []datadogV2.HTTPLogItem{{
				Ddsource: datadog.PtrString("golang"),
				Ddtags:   datadog.PtrString("tag1:true"),
				Hostname: datadog.PtrString("hostname"),
				Message:  "log 1",
				Service:  datadog.PtrString("server"),
				UnparsedObject: map[string]interface{}{
					"ddsource": "golang",
					"ddtags":   "tag1:true",
					"hostname": "hostname",
					"message":  "log 1",
					"service":  "server",
				},
			}, {
				Ddsource: datadog.PtrString("golang"),
				Ddtags:   datadog.PtrString("tag1:true"),
				Hostname: datadog.PtrString("hostname"),
				Message:  "log 2",
				Service:  datadog.PtrString("server"),
				UnparsedObject: map[string]interface{}{
					"ddsource": "golang",
					"ddtags":   "tag1:true",
					"hostname": "hostname",
					"message":  "log 2",
					"service":  "server",
				},
			}},
			testFn: func(jsonLogs testutil.JSONLogs, counter int) {
				switch counter {
				case 0:
					assert.True(t, jsonLogs.HasDDTag("tag1:true"))
					assert.Len(t, jsonLogs, 2)
				default:
					t.Fail()
				}
			},
			numRequests: 1,
		},
		{
			name: "different-tags",
			payload: []datadogV2.HTTPLogItem{{
				Ddsource: datadog.PtrString("golang"),
				Ddtags:   datadog.PtrString("tag1:true"),
				Hostname: datadog.PtrString("hostname"),
				Message:  "log 1",
				Service:  datadog.PtrString("server"),
				UnparsedObject: map[string]interface{}{
					"ddsource": "golang",
					"ddtags":   "tag1:true",
					"hostname": "hostname",
					"message":  "log 1",
					"service":  "server",
				},
			}, {
				Ddsource: datadog.PtrString("golang"),
				Ddtags:   datadog.PtrString("tag2:true"),
				Hostname: datadog.PtrString("hostname"),
				Message:  "log 2",
				Service:  datadog.PtrString("server"),
				UnparsedObject: map[string]interface{}{
					"ddsource": "golang",
					"ddtags":   "tag2:true",
					"hostname": "hostname",
					"message":  "log 2",
					"service":  "server",
				},
			}},
			testFn: func(jsonLogs testutil.JSONLogs, counter int) {
				switch counter {
				case 0:
					assert.True(t, jsonLogs.HasDDTag("tag1:true"))
					assert.Len(t, jsonLogs, 1)
				case 1:
					assert.True(t, jsonLogs.HasDDTag("tag2:true"))
					assert.Len(t, jsonLogs, 1)
				default:
					t.Fail()
				}
			},
			numRequests: 2,
		},
		{
			name: "two-batches",
			payload: []datadogV2.HTTPLogItem{{
				Ddsource: datadog.PtrString("golang"),
				Ddtags:   datadog.PtrString("tag1:true"),
				Hostname: datadog.PtrString("hostname"),
				Message:  "log 1",
				Service:  datadog.PtrString("server"),
				UnparsedObject: map[string]interface{}{
					"ddsource": "golang",
					"ddtags":   "tag1:true",
					"hostname": "hostname",
					"message":  "log 1",
					"service":  "server",
				},
			}, {
				Ddsource: datadog.PtrString("golang"),
				Ddtags:   datadog.PtrString("tag1:true"),
				Hostname: datadog.PtrString("hostname"),
				Message:  "log 2",
				Service:  datadog.PtrString("server"),
				UnparsedObject: map[string]interface{}{
					"ddsource": "golang",
					"ddtags":   "tag1:true",
					"hostname": "hostname",
					"message":  "log 2",
					"service":  "server",
				},
			}, {
				Ddsource: datadog.PtrString("golang"),
				Ddtags:   datadog.PtrString("tag2:true"),
				Hostname: datadog.PtrString("hostname"),
				Message:  "log 3",
				Service:  datadog.PtrString("server"),
				UnparsedObject: map[string]interface{}{
					"ddsource": "golang",
					"ddtags":   "tag2:true",
					"hostname": "hostname",
					"message":  "log 3",
					"service":  "server",
				},
			}, {
				Ddsource: datadog.PtrString("golang"),
				Ddtags:   datadog.PtrString("tag2:true"),
				Hostname: datadog.PtrString("hostname"),
				Message:  "log 4",
				Service:  datadog.PtrString("server"),
				UnparsedObject: map[string]interface{}{
					"ddsource": "golang",
					"ddtags":   "tag2:true",
					"hostname": "hostname",
					"message":  "log 4",
					"service":  "server",
				},
			}},
			testFn: func(jsonLogs testutil.JSONLogs, counter int) {
				switch counter {
				case 0:
					assert.True(t, jsonLogs.HasDDTag("tag1:true"))
					assert.Len(t, jsonLogs, 2)
				case 1:
					assert.True(t, jsonLogs.HasDDTag("tag2:true"))
					assert.Len(t, jsonLogs, 2)
				default:
					t.Fail()
				}
			},
			numRequests: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var numCalls int
			server := testutil.DatadogLogServerMock(func() (string, http.HandlerFunc) {
				return "/api/v2/logs", func(writer http.ResponseWriter, request *http.Request) {
					jsonLogs := testutil.MockLogsEndpoint(writer, request)
					tt.testFn(jsonLogs, numCalls)
					numCalls++
				}
			})
			defer server.Close()
			s := NewSender(server.URL, logger, exporterhelper.TimeoutSettings{Timeout: time.Second * 10}, true, true, "")
			if err := s.SubmitLogs(context.Background(), tt.payload); err != nil {
				t.Fatal(err)
			}
			assert.True(t, numCalls > 0 && numCalls == tt.numRequests)
		})
	}
}
