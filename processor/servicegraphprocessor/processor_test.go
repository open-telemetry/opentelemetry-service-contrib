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

package servicegraphprocessor

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/connector/connectortest"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exportertest"
	"go.opentelemetry.io/collector/exporter/otlpexporter"
	"go.opentelemetry.io/collector/featuregate"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/processor/processortest"
	semconv "go.opentelemetry.io/collector/semconv/v1.13.0"
	"go.uber.org/zap/zaptest"
)

func TestProcessorStart(t *testing.T) {
	// Create otlp exporters.
	otlpID, mexp, texp := newOTLPExporters(t)

	for _, tc := range []struct {
		name            string
		exporter        component.Component
		metricsExporter string
		wantErrorMsg    string
	}{
		{"export to active otlp metrics exporter", mexp, "otlp", ""},
		{"unable to find configured exporter in active exporter list", mexp, "prometheus", "failed to find metrics exporter: 'prometheus'; please configure metrics_exporter from one of: [otlp]"},
		{"export to active otlp traces exporter should error", texp, "otlp", "the exporter \"otlp\" isn't a metrics exporter"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Prepare
			exporters := map[component.DataType]map[component.ID]component.Component{
				component.DataTypeMetrics: {
					otlpID: tc.exporter,
				},
			}

			// Create servicegraph processor
			factory := NewFactory()
			cfg := factory.CreateDefaultConfig().(*Config)
			cfg.MetricsExporter = tc.metricsExporter

			procCreationParams := processortest.NewNopCreateSettings()
			traceProcessor, err := factory.CreateTracesProcessor(context.Background(), procCreationParams, cfg, consumertest.NewNop())
			require.NoError(t, err)

			// Test
			smp := traceProcessor.(*serviceGraphProcessor)
			err = smp.Start(context.Background(), newMockHost(exporters))

			// Verify
			if tc.wantErrorMsg != "" {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestConnectorStart(t *testing.T) {
	// Create servicegraph processor
	factory := NewConnectorFactory()
	cfg := factory.CreateDefaultConfig().(*Config)

	procCreationParams := connectortest.NewNopCreateSettings()
	traceProcessor, err := factory.CreateTracesToMetrics(context.Background(), procCreationParams, cfg, consumertest.NewNop())
	require.NoError(t, err)

	// Test
	smp := traceProcessor.(*serviceGraphProcessor)
	err = smp.Start(context.Background(), componenttest.NewNopHost())

	// Verify
	assert.NoError(t, err)
}

func TestProcessorShutdown(t *testing.T) {
	// Prepare
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig().(*Config)

	// Test
	next := new(consumertest.TracesSink)
	p := newProcessor(zaptest.NewLogger(t), cfg)
	p.tracesConsumer = next
	err := p.Shutdown(context.Background())

	// Verify
	assert.NoError(t, err)
}

func TestConnectorShutdown(t *testing.T) {
	// Prepare
	factory := NewConnectorFactory()
	cfg := factory.CreateDefaultConfig().(*Config)

	// Test
	next := new(consumertest.MetricsSink)
	p := newProcessor(zaptest.NewLogger(t), cfg)
	p.metricsConsumer = next
	err := p.Shutdown(context.Background())

	// Verify
	assert.NoError(t, err)
}

func TestProcessorConsume(t *testing.T) {
	metricsExporter := newMockMetricsExporter(func(md pmetric.Metrics) error {
		return verifyMetrics(t, md)
	})
	// set virtual node feature
	_ = featuregate.GlobalRegistry().Set(virtualNodeFeatureGate.ID(), true)

	for _, tc := range []struct {
		name         string
		cfg          Config
		sampleTraces ptrace.Traces
	}{
		{
			name: "traces with client and server span",
			cfg: Config{
				MetricsExporter: "mock",
				Dimensions:      []string{"some-attribute", "non-existing-attribute"},
			}, sampleTraces: buildSampleTrace("val"),
		},
		{
			name: "incomplete traces with server span lost",
			cfg: Config{
				MetricsExporter: "mock",
				Dimensions:      []string{"some-attribute", "non-existing-attribute"},
			},
			sampleTraces: incompleteClientTraces(),
		},
		{
			name: "incomplete traces with client span lost",
			cfg: Config{
				MetricsExporter: "mock",
				Dimensions:      []string{"some-attribute", "non-existing-attribute"},
			},
			sampleTraces: incompleteServerTraces(),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Prepare
			processor := newProcessor(zaptest.NewLogger(t), &tc.cfg)
			processor.tracesConsumer = consumertest.NewNop()
			mHost := newMockHost(map[component.DataType]map[component.ID]component.Component{
				component.DataTypeMetrics: {
					component.NewID("mock"): metricsExporter,
				},
			})

			assert.NoError(t, processor.Start(context.Background(), mHost))

			// Test & verify
			// The assertion is part of verifyMetrics func.
			assert.NoError(t, processor.ConsumeTraces(context.Background(), tc.sampleTraces))
			time.Sleep(time.Second * 2)
			// Shutdown the processor
			assert.NoError(t, processor.Shutdown(context.Background()))
		})
	}

	// unset virtual node feature
	_ = featuregate.GlobalRegistry().Set(virtualNodeFeatureGate.ID(), false)
}

func TestConnectorConsume(t *testing.T) {
	// Prepare
	cfg := &Config{
		Dimensions: []string{"some-attribute", "non-existing-attribute"},
	}

	conn := newProcessor(zaptest.NewLogger(t), cfg)
	conn.metricsConsumer = newMockMetricsExporter(func(md pmetric.Metrics) error {
		return verifyMetrics(t, md)
	})

	assert.NoError(t, conn.Start(context.Background(), componenttest.NewNopHost()))

	// Test & verify
	td := buildSampleTrace("val")
	// The assertion is part of verifyMetrics func.
	assert.NoError(t, conn.ConsumeTraces(context.Background(), td))

	// Shutdown the conn
	assert.NoError(t, conn.Shutdown(context.Background()))
}

func verifyMetrics(t *testing.T, md pmetric.Metrics) error {
	assert.Equal(t, 2, md.MetricCount())

	rms := md.ResourceMetrics()
	assert.Equal(t, 1, rms.Len())

	sms := rms.At(0).ScopeMetrics()
	assert.Equal(t, 1, sms.Len())

	ms := sms.At(0).Metrics()
	assert.Equal(t, 2, ms.Len())

	mCount := ms.At(0)
	verifyCount(t, mCount)

	mDuration := ms.At(1)
	verifyDuration(t, mDuration)

	return nil
}

func verifyCount(t *testing.T, m pmetric.Metric) {
	assert.Equal(t, "traces_service_graph_request_total", m.Name())

	assert.Equal(t, pmetric.MetricTypeSum, m.Type())
	dps := m.Sum().DataPoints()
	assert.Equal(t, 1, dps.Len())

	dp := dps.At(0)
	assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
	assert.Equal(t, int64(1), dp.IntValue())

	attributes := dp.Attributes()
	assert.Equal(t, 6, attributes.Len())
	verifyAttr(t, attributes, "client", "some-service")
	verifyAttr(t, attributes, "server", "some-service")
	verifyAttr(t, attributes, "failed", "false")
	verifyAttr(t, attributes, "some-attribute", "val")
}

func verifyDuration(t *testing.T, m pmetric.Metric) {
	assert.Equal(t, "traces_service_graph_request_duration_seconds", m.Name())

	assert.Equal(t, pmetric.MetricTypeHistogram, m.Type())
	dps := m.Histogram().DataPoints()
	assert.Equal(t, 1, dps.Len())

	dp := dps.At(0)
	assert.Equal(t, float64(1000), dp.Sum()) // Duration: 1sec
	assert.Equal(t, uint64(1), dp.Count())
	assert.Equal(t, []uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0}, dp.BucketCounts())

	attributes := dp.Attributes()
	assert.Equal(t, 4, attributes.Len())
	verifyAttr(t, attributes, "client", "some-service")
	verifyAttr(t, attributes, "server", "some-service")
	verifyAttr(t, attributes, "failed", "false")
	verifyAttr(t, attributes, "some-attribute", "val")
}

func verifyAttr(t *testing.T, attrs pcommon.Map, k, expected string) {
	v, ok := attrs.Get(k)
	assert.True(t, ok)
	assert.Equal(t, expected, v.AsString())
}

func buildSampleTrace(attrValue string) ptrace.Traces {
	tStart := time.Date(2022, 1, 2, 3, 4, 5, 6, time.UTC)
	tEnd := time.Date(2022, 1, 2, 3, 4, 6, 6, time.UTC)

	traces := ptrace.NewTraces()

	resourceSpans := traces.ResourceSpans().AppendEmpty()
	resourceSpans.Resource().Attributes().PutStr(semconv.AttributeServiceName, "some-service")

	scopeSpans := resourceSpans.ScopeSpans().AppendEmpty()

	var traceID pcommon.TraceID
	rand.Read(traceID[:])

	var clientSpanID, serverSpanID pcommon.SpanID
	rand.Read(clientSpanID[:])
	rand.Read(serverSpanID[:])

	clientSpan := scopeSpans.Spans().AppendEmpty()
	clientSpan.SetName("client span")
	clientSpan.SetSpanID(clientSpanID)
	clientSpan.SetTraceID(traceID)
	clientSpan.SetKind(ptrace.SpanKindClient)
	clientSpan.SetStartTimestamp(pcommon.NewTimestampFromTime(tStart))
	clientSpan.SetEndTimestamp(pcommon.NewTimestampFromTime(tEnd))
	clientSpan.Attributes().PutStr("some-attribute", attrValue) // Attribute selected as dimension for metrics

	serverSpan := scopeSpans.Spans().AppendEmpty()
	serverSpan.SetName("server span")
	serverSpan.SetSpanID(serverSpanID)
	serverSpan.SetTraceID(traceID)
	serverSpan.SetParentSpanID(clientSpanID)
	serverSpan.SetKind(ptrace.SpanKindServer)
	serverSpan.SetStartTimestamp(pcommon.NewTimestampFromTime(tStart))
	serverSpan.SetEndTimestamp(pcommon.NewTimestampFromTime(tEnd))

	return traces
}

func incompleteClientTraces() ptrace.Traces {
	tStart := time.Date(2022, 1, 2, 3, 4, 5, 6, time.UTC)
	tEnd := time.Date(2022, 1, 2, 3, 4, 6, 6, time.UTC)

	traces := ptrace.NewTraces()

	resourceSpans := traces.ResourceSpans().AppendEmpty()
	resourceSpans.Resource().Attributes().PutStr(semconv.AttributeServiceName, "some-client-service")

	scopeSpans := resourceSpans.ScopeSpans().AppendEmpty()
	anotherTraceID := pcommon.TraceID([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
	anotherClientSpanID := pcommon.SpanID([8]byte{1, 2, 3, 4, 4, 3, 2, 1})
	clientSpanNoServerSpan := scopeSpans.Spans().AppendEmpty()
	clientSpanNoServerSpan.SetName("client span")
	clientSpanNoServerSpan.SetSpanID(anotherClientSpanID)
	clientSpanNoServerSpan.SetTraceID(anotherTraceID)
	clientSpanNoServerSpan.SetKind(ptrace.SpanKindClient)
	clientSpanNoServerSpan.SetStartTimestamp(pcommon.NewTimestampFromTime(tStart))
	clientSpanNoServerSpan.SetEndTimestamp(pcommon.NewTimestampFromTime(tEnd))
	clientSpanNoServerSpan.Attributes().PutStr(semconv.AttributeNetSockPeerAddr, "127.10.10.1") // Attribute selected as dimension for metrics

	return traces
}

func incompleteServerTraces() ptrace.Traces {
	tStart := time.Date(2022, 1, 2, 3, 4, 5, 6, time.UTC)
	tEnd := time.Date(2022, 1, 2, 3, 4, 6, 6, time.UTC)

	traces := ptrace.NewTraces()

	resourceSpans := traces.ResourceSpans().AppendEmpty()
	resourceSpans.Resource().Attributes().PutStr(semconv.AttributeServiceName, "some-server-service")
	scopeSpans := resourceSpans.ScopeSpans().AppendEmpty()
	anotherTraceID := pcommon.TraceID([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 8, 7, 6, 5, 4, 3, 2, 1})
	serverSpanNoClientSpan := scopeSpans.Spans().AppendEmpty()
	serverSpanNoClientSpan.SetName("server span")
	serverSpanNoClientSpan.SetSpanID([8]byte{0x19, 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26})
	serverSpanNoClientSpan.SetTraceID(anotherTraceID)
	serverSpanNoClientSpan.SetKind(ptrace.SpanKindServer)
	serverSpanNoClientSpan.SetStartTimestamp(pcommon.NewTimestampFromTime(tStart))
	serverSpanNoClientSpan.SetEndTimestamp(pcommon.NewTimestampFromTime(tEnd))
	return traces
}

func newOTLPExporters(t *testing.T) (component.ID, exporter.Metrics, exporter.Traces) {
	otlpExpFactory := otlpexporter.NewFactory()
	otlpID := component.NewID("otlp")
	otlpConfig := &otlpexporter.Config{
		GRPCClientSettings: configgrpc.GRPCClientSettings{
			Endpoint: "example.com:1234",
		},
	}
	expCreationParams := exportertest.NewNopCreateSettings()
	mexp, err := otlpExpFactory.CreateMetricsExporter(context.Background(), expCreationParams, otlpConfig)
	require.NoError(t, err)
	texp, err := otlpExpFactory.CreateTracesExporter(context.Background(), expCreationParams, otlpConfig)
	require.NoError(t, err)
	return otlpID, mexp, texp
}

type mockHost struct {
	component.Host
	exps map[component.DataType]map[component.ID]component.Component
}

func newMockHost(exps map[component.DataType]map[component.ID]component.Component) component.Host {
	return &mockHost{
		Host: componenttest.NewNopHost(),
		exps: exps,
	}
}

func (m *mockHost) GetExporters() map[component.DataType]map[component.ID]component.Component {
	return m.exps
}

var _ exporter.Metrics = (*mockMetricsExporter)(nil)

func newMockMetricsExporter(verifyFunc func(md pmetric.Metrics) error) exporter.Metrics {
	return &mockMetricsExporter{verify: verifyFunc}
}

type mockMetricsExporter struct {
	verify func(md pmetric.Metrics) error
}

func (m *mockMetricsExporter) Start(context.Context, component.Host) error { return nil }

func (m *mockMetricsExporter) Shutdown(context.Context) error { return nil }

func (m *mockMetricsExporter) Capabilities() consumer.Capabilities { return consumer.Capabilities{} }

func (m *mockMetricsExporter) ConsumeMetrics(_ context.Context, md pmetric.Metrics) error {
	return m.verify(md)
}

func TestUpdateDurationMetrics(t *testing.T) {
	p := serviceGraphProcessor{
		reqTotal:                             make(map[string]int64),
		reqFailedTotal:                       make(map[string]int64),
		reqServerDurationSecondsSum:          make(map[string]float64),
		reqServerDurationSecondsCount:        make(map[string]uint64),
		reqServerDurationSecondsBucketCounts: make(map[string][]uint64),
		reqClientDurationSecondsSum:          make(map[string]float64),
		reqClientDurationSecondsCount:        make(map[string]uint64),
		reqClientDurationSecondsBucketCounts: make(map[string][]uint64),
		reqDurationBounds:                    defaultLatencyHistogramBucketsMs,
		keyToMetric:                          make(map[string]metricSeries),
		config: &Config{
			Dimensions: []string{},
		},
	}
	metricKey := p.buildMetricKey("foo", "bar", "", map[string]string{})

	testCases := []struct {
		caseStr  string
		duration float64
	}{

		{
			caseStr:  "index 0 latency",
			duration: 0,
		},
		{
			caseStr:  "out-of-range latency 1",
			duration: 25_000,
		},
		{
			caseStr:  "out-of-range latency 2",
			duration: 125_000,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.caseStr, func(t *testing.T) {
			p.updateDurationMetrics(metricKey, tc.duration, tc.duration)
		})
	}
}

func TestStaleSeriesCleanup(t *testing.T) {
	// Prepare
	cfg := &Config{
		MetricsExporter: "mock",
		Dimensions:      []string{"some-attribute", "non-existing-attribute"},
		Store: StoreConfig{
			MaxItems: 10,
			TTL:      time.Second,
		},
	}

	mockMetricsExporter := newMockMetricsExporter(func(md pmetric.Metrics) error { return nil })

	p := newProcessor(zaptest.NewLogger(t), cfg)
	p.tracesConsumer = consumertest.NewNop()

	mHost := newMockHost(map[component.DataType]map[component.ID]component.Component{
		component.DataTypeMetrics: {
			component.NewID("mock"): mockMetricsExporter,
		},
	})

	assert.NoError(t, p.Start(context.Background(), mHost))

	// ConsumeTraces
	td := buildSampleTrace("first")
	assert.NoError(t, p.ConsumeTraces(context.Background(), td))

	// Make series stale and force a cache cleanup
	for key, metric := range p.keyToMetric {
		metric.lastUpdated = 0
		p.keyToMetric[key] = metric
	}
	p.cleanCache()
	assert.Equal(t, 0, len(p.keyToMetric))

	// ConsumeTraces with a trace with different attribute value
	td = buildSampleTrace("second")
	assert.NoError(t, p.ConsumeTraces(context.Background(), td))

	// Shutdown the processor
	assert.NoError(t, p.Shutdown(context.Background()))
}
