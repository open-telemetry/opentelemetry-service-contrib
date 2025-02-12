// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusremotewritereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusremotewritereceiver"

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/snappy"
	promconfig "github.com/prometheus/prometheus/config"
	writev2 "github.com/prometheus/prometheus/prompb/io/prometheus/write/v2"
	"github.com/prometheus/prometheus/storage/remote"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest/pmetrictest"
)

var writeV2RequestFixture = &writev2.Request{
	Symbols: []string{"", "__name__", "test_metric1", "job", "service-x/test", "instance", "107cn001", "d", "e", "foo", "bar", "f", "g", "h", "i", "Test gauge for test purposes", "Maybe op/sec who knows (:", "Test counter for test purposes"},
	Timeseries: []writev2.TimeSeries{
		{
			Metadata:   writev2.Metadata{Type: writev2.Metadata_METRIC_TYPE_GAUGE},
			LabelsRefs: []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, // Symbolized writeRequestFixture.Timeseries[0].Labels
			Samples:    []writev2.Sample{{Value: 1, Timestamp: 1}},
		},
		{
			Metadata:   writev2.Metadata{Type: writev2.Metadata_METRIC_TYPE_GAUGE},
			LabelsRefs: []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, // Same series as first. Should use the same resource metrics.
			Samples:    []writev2.Sample{{Value: 2, Timestamp: 2}},
		},
		{
			Metadata:   writev2.Metadata{Type: writev2.Metadata_METRIC_TYPE_GAUGE},
			LabelsRefs: []uint32{1, 2, 3, 9, 5, 10, 7, 8, 9, 10}, // This series has different label values for job and instance.
			Samples:    []writev2.Sample{{Value: 2, Timestamp: 2}},
		},
	},
}

func setupMetricsReceiver(t *testing.T) *prometheusRemoteWriteReceiver {
	t.Helper()

	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()

	prwReceiver, err := factory.CreateMetrics(context.Background(), receivertest.NewNopSettings(), cfg, consumertest.NewNop())
	assert.NoError(t, err)
	assert.NotNil(t, prwReceiver, "metrics receiver creation failed")

	return prwReceiver.(*prometheusRemoteWriteReceiver)
}

func TestHandlePRWContentTypeNegotiation(t *testing.T) {
	for _, tc := range []struct {
		name         string
		contentType  string
		expectedCode int
	}{
		{
			name:         "no content type",
			contentType:  "",
			expectedCode: http.StatusUnsupportedMediaType,
		},
		{
			name:         "unsupported content type",
			contentType:  "application/json",
			expectedCode: http.StatusUnsupportedMediaType,
		},
		{
			name:         "x-protobuf/no proto parameter",
			contentType:  "application/x-protobuf",
			expectedCode: http.StatusUnsupportedMediaType,
		},
		{
			name:         "x-protobuf/v1 proto parameter",
			contentType:  fmt.Sprintf("application/x-protobuf;proto=%s", promconfig.RemoteWriteProtoMsgV1),
			expectedCode: http.StatusUnsupportedMediaType,
		},
		{
			name:         "x-protobuf/v2 proto parameter",
			contentType:  fmt.Sprintf("application/x-protobuf;proto=%s", promconfig.RemoteWriteProtoMsgV2),
			expectedCode: http.StatusNoContent,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			body := writev2.Request{}
			pBuf := proto.NewBuffer(nil)
			err := pBuf.Marshal(&body)
			assert.NoError(t, err)

			var compressedBody []byte
			snappy.Encode(compressedBody, pBuf.Bytes())

			req := httptest.NewRequest(http.MethodPost, "/api/v1/write", bytes.NewBuffer(compressedBody))

			req.Header.Set("Content-Type", tc.contentType)
			req.Header.Set("Content-Encoding", "snappy")
			w := httptest.NewRecorder()

			prwReceiver := setupMetricsReceiver(t)
			prwReceiver.handlePRW(w, req)
			resp := w.Result()

			assert.Equal(t, tc.expectedCode, resp.StatusCode)
			if tc.expectedCode == http.StatusNoContent { // We went until the end
				assert.NotEmpty(t, resp.Header.Get("X-Prometheus-Remote-Write-Samples-Written"))
				assert.NotEmpty(t, resp.Header.Get("X-Prometheus-Remote-Write-Histograms-Written"))
				assert.NotEmpty(t, resp.Header.Get("X-Prometheus-Remote-Write-Exemplars-Written"))
			}
		})
	}
}

func TestTranslateV2(t *testing.T) {
	prwReceiver := setupMetricsReceiver(t)
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	for _, tc := range []struct {
		name            string
		request         *writev2.Request
		expectError     string
		expectedMetrics pmetric.Metrics
		expectedStats   remote.WriteResponseStats
	}{
		{
			name: "duplicated scope name and version",
			request: &writev2.Request{
				Symbols: []string{
					"",
					"__name__", "test_metric",
					"job", "service-x/test",
					"instance", "107cn001",
					"otel_scope_name", "scope1",
					"otel_scope_version", "v1",
					"otel_scope_name", "scope2",
					"otel_scope_version", "v2",
					"d", "e",
					"foo", "bar",
				},
				Timeseries: []writev2.TimeSeries{
					{
						Metadata:   writev2.Metadata{Type: writev2.Metadata_METRIC_TYPE_GAUGE},
						LabelsRefs: []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 15, 16}, // Same scope: scope_name: scope1. scope_version v1
						Samples:    []writev2.Sample{{Value: 1, Timestamp: 1}},
					},
					{
						Metadata:   writev2.Metadata{Type: writev2.Metadata_METRIC_TYPE_GAUGE},
						LabelsRefs: []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 15, 16}, // Same scope: scope_name: scope1. scope_version v1
						Samples:    []writev2.Sample{{Value: 2, Timestamp: 2}},
					},
					{
						Metadata:   writev2.Metadata{Type: writev2.Metadata_METRIC_TYPE_GAUGE},
						LabelsRefs: []uint32{1, 2, 3, 4, 5, 6, 11, 12, 13, 14, 17, 18}, // Different scope: scope_name: scope2. scope_version v2
						Samples:    []writev2.Sample{{Value: 3, Timestamp: 3}},
					},
				},
			},
			expectedMetrics: func() pmetric.Metrics {
				expected := pmetric.NewMetrics()
				rm1 := expected.ResourceMetrics().AppendEmpty()
				rmAttributes1 := rm1.Resource().Attributes()
				rmAttributes1.PutStr("service.namespace", "service-x")
				rmAttributes1.PutStr("service.name", "test")
				rmAttributes1.PutStr("service.instance.id", "107cn001")
				sm1 := rm1.ScopeMetrics().AppendEmpty()
				sm1.Scope().SetName("scope1")
				sm1.Scope().SetVersion("v1")
				sm1Attributes := sm1.Metrics().AppendEmpty().SetEmptyGauge().DataPoints().AppendEmpty().Attributes()
				sm1Attributes.PutStr("d", "e")
				sm2Attributes := sm1.Metrics().AppendEmpty().SetEmptyGauge().DataPoints().AppendEmpty().Attributes()
				sm2Attributes.PutStr("d", "e")

				sm2 := rm1.ScopeMetrics().AppendEmpty()
				sm2.Scope().SetName("scope2")
				sm2.Scope().SetVersion("v2")
				sm3Attributes := sm2.Metrics().AppendEmpty().SetEmptyGauge().DataPoints().AppendEmpty().Attributes()
				sm3Attributes.PutStr("foo", "bar")
				return expected
			}(),
			expectedStats: remote.WriteResponseStats{},
		},
		{
			name: "missing metric name",
			request: &writev2.Request{
				Symbols: []string{"", "foo", "bar"},
				Timeseries: []writev2.TimeSeries{
					{
						LabelsRefs: []uint32{1, 2},
						Samples:    []writev2.Sample{{Value: 1, Timestamp: 1}},
					},
				},
			},
			expectError: "missing metric name in labels",
		},
		{
			name: "duplicate label",
			request: &writev2.Request{
				Symbols: []string{"", "__name__", "test"},
				Timeseries: []writev2.TimeSeries{
					{
						LabelsRefs: []uint32{1, 2, 1, 2},
						Samples:    []writev2.Sample{{Value: 1, Timestamp: 1}},
					},
				},
			},
			expectError: `duplicate label "__name__" in labels`,
		},
		{
			name:    "valid request",
			request: writeV2RequestFixture,
			expectedMetrics: func() pmetric.Metrics {
				expected := pmetric.NewMetrics()

				// Resource 1: for job "service-x/test" → namespace "service-x", name "test", instance "107cn001"
				rm1 := expected.ResourceMetrics().AppendEmpty()
				rm1.Resource().Attributes().PutStr("service.namespace", "service-x")
				rm1.Resource().Attributes().PutStr("service.name", "test")
				rm1.Resource().Attributes().PutStr("service.instance.id", "107cn001")
				sm1 := rm1.ScopeMetrics().AppendEmpty()
				sm1.Scope().SetName(buildName)
				sm1.Scope().SetVersion(buildVersion)
				// Timeseries 1 gauge metric.
				m1 := sm1.Metrics().AppendEmpty().SetEmptyGauge()
				dp1 := m1.DataPoints().AppendEmpty()
				dp1.Attributes().PutStr("d", "e")
				dp1.Attributes().PutStr("foo", "bar")
				// Timeseries 2 gauge metric.
				m2 := sm1.Metrics().AppendEmpty().SetEmptyGauge()
				dp2 := m2.DataPoints().AppendEmpty()
				dp2.Attributes().PutStr("d", "e")
				dp2.Attributes().PutStr("foo", "bar")

				// Resource 2: for job "foo" with instance "bar"
				rm2 := expected.ResourceMetrics().AppendEmpty()
				rm2.Resource().Attributes().PutStr("service.name", "foo")
				rm2.Resource().Attributes().PutStr("service.instance.id", "bar")
				sm2 := rm2.ScopeMetrics().AppendEmpty()
				sm2.Scope().SetName(buildName)
				sm2.Scope().SetVersion(buildVersion)
				m3 := sm2.Metrics().AppendEmpty().SetEmptyGauge()
				dp3 := m3.DataPoints().AppendEmpty()
				dp3.Attributes().PutStr("d", "e")
				dp3.Attributes().PutStr("foo", "bar")

				return expected
			}(),
			expectedStats: remote.WriteResponseStats{},
		},
		{
			name: "provided otel_scope_name and otel_scope_version",
			request: &writev2.Request{
				Symbols: []string{
					"", "__name__", "metric_with_scope",
					"otel_scope_name", "custom_scope",
					"otel_scope_version", "v1.0",
					"job", "service-y/custom",
					"instance", "instance-1",
				},
				Timeseries: []writev2.TimeSeries{
					{
						Metadata:   writev2.Metadata{Type: writev2.Metadata_METRIC_TYPE_GAUGE},
						LabelsRefs: []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
						Samples:    []writev2.Sample{{Value: 10, Timestamp: 100}},
					},
				},
			},
			expectedMetrics: func() pmetric.Metrics {
				expected := pmetric.NewMetrics()
				rm := expected.ResourceMetrics().AppendEmpty()
				// For job "service-y/custom", resource attributes are set as follows:
				rm.Resource().Attributes().PutStr("service.namespace", "service-y")
				rm.Resource().Attributes().PutStr("service.name", "custom")
				rm.Resource().Attributes().PutStr("service.instance.id", "instance-1")
				sm := rm.ScopeMetrics().AppendEmpty()
				// Expect the provided scope values.
				sm.Scope().SetName("custom_scope")
				sm.Scope().SetVersion("v1.0")
				m := sm.Metrics().AppendEmpty().SetEmptyGauge()
				// The datapoint will have no extra attributes because __name__, otel_scope_name/version,
				// job and instance are filtered out.
				_ = m.DataPoints().AppendEmpty()
				return expected
			}(),
			expectedStats: remote.WriteResponseStats{},
		},
		{
			name: "missing otel_scope_name/version falls back to buildName/buildVersion",
			request: &writev2.Request{
				Symbols: []string{
					"",                // index 0
					"__name__",        // index 1
					"metric_no_scope", // index 2
					"job",             // index 3
					"service-z/xyz",   // index 4
					"instance",        // index 5
					"inst-42",         // index 6
					"d",               // index 7
					"e",               // index 8
					"foo",             // index 9
					"bar",             // index 10
				},
				Timeseries: []writev2.TimeSeries{
					{
						Metadata:   writev2.Metadata{Type: writev2.Metadata_METRIC_TYPE_GAUGE},
						LabelsRefs: []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
						Samples:    []writev2.Sample{{Value: 5, Timestamp: 50}},
					},
				},
			},
			expectedMetrics: func() pmetric.Metrics {
				expected := pmetric.NewMetrics()
				rm := expected.ResourceMetrics().AppendEmpty()
				// parseJobAndInstance splits "service-z/xyz" into namespace "service-z" and name "xyz"
				rm.Resource().Attributes().PutStr("service.namespace", "service-z")
				rm.Resource().Attributes().PutStr("service.name", "xyz")
				rm.Resource().Attributes().PutStr("service.instance.id", "inst-42")
				sm := rm.ScopeMetrics().AppendEmpty()
				// Since otel_scope_name/version are missing, the code falls back to buildName/buildVersion.
				sm.Scope().SetName(buildName)
				sm.Scope().SetVersion(buildVersion)
				m := sm.Metrics().AppendEmpty().SetEmptyGauge()
				dp := m.DataPoints().AppendEmpty()
				// Expect the attributes from the labels that are added by addDatapoints.
				dp.Attributes().PutStr("d", "e")
				dp.Attributes().PutStr("foo", "bar")
				return expected
			}(),
			expectedStats: remote.WriteResponseStats{},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			metrics, stats, err := prwReceiver.translateV2(ctx, tc.request)
			if tc.expectError != "" {
				assert.ErrorContains(t, err, tc.expectError)
				return
			}

			assert.NoError(t, err)
			assert.NoError(t, pmetrictest.CompareMetrics(tc.expectedMetrics, metrics))
			assert.Equal(t, tc.expectedStats, stats)
		})
	}
}
