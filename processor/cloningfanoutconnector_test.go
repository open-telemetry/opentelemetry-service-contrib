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

package processor

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/pdatautil"
	"go.opentelemetry.io/collector/exporter/exportertest"
	"go.opentelemetry.io/collector/internal/data/testdata"
)

func TestTraceProcessorCloningNotMultiplexing(t *testing.T) {
	processors := []consumer.TraceConsumer{
		new(exportertest.SinkTraceExporter),
	}

	tfc := NewTracesCloningFanOutConnector(processors)

	assert.Same(t, processors[0], tfc)
}

func TestTraceProcessorCloningMultiplexing(t *testing.T) {
	processors := make([]consumer.TraceConsumer, 3)
	for i := range processors {
		processors[i] = new(exportertest.SinkTraceExporter)
	}

	tfc := NewTracesCloningFanOutConnector(processors)
	td := testdata.GenerateTraceDataTwoSpansSameResource()

	var wantSpansCount = 0
	for i := 0; i < 2; i++ {
		wantSpansCount += td.SpanCount()
		err := tfc.ConsumeTraces(context.Background(), td)
		if err != nil {
			t.Errorf("Wanted nil got error")
			return
		}
	}

	for i, p := range processors {
		m := p.(*exportertest.SinkTraceExporter)
		assert.Equal(t, wantSpansCount, m.SpansCount())
		spanOrig := td.ResourceSpans().At(0).InstrumentationLibrarySpans().At(0).Spans().At(0)
		allTraces := m.AllTraces()
		spanClone := allTraces[0].ResourceSpans().At(0).InstrumentationLibrarySpans().At(0).Spans().At(0)
		if i < len(processors)-1 {
			assert.True(t, td.ResourceSpans().At(0).Resource() != allTraces[0].ResourceSpans().At(0).Resource())
			assert.True(t, spanOrig != spanClone)
		} else {
			assert.True(t, td.ResourceSpans().At(0).Resource() == allTraces[0].ResourceSpans().At(0).Resource())
			assert.True(t, spanOrig == spanClone)
		}
		assert.EqualValues(t, td.ResourceSpans().At(0).Resource(), allTraces[0].ResourceSpans().At(0).Resource())
		assert.EqualValues(t, spanOrig, spanClone)
	}
}

func TestMetricsProcessorCloningNotMultiplexing(t *testing.T) {
	processors := []consumer.MetricsConsumer{
		new(exportertest.SinkMetricsExporter),
	}

	tfc := NewMetricsCloningFanOutConnector(processors)

	assert.Same(t, processors[0], tfc)
}

func TestMetricsProcessorCloningMultiplexing(t *testing.T) {
	processors := make([]consumer.MetricsConsumer, 3)
	for i := range processors {
		processors[i] = new(exportertest.SinkMetricsExporter)
	}

	mfc := NewMetricsCloningFanOutConnector(processors)
	md := testdata.GenerateMetricDataWithCountersHistogramAndSummary()

	var wantMetricsCount = 0
	for i := 0; i < 2; i++ {
		wantMetricsCount += md.MetricCount()
		err := mfc.ConsumeMetrics(context.Background(), pdatautil.MetricsFromInternalMetrics(md))
		if err != nil {
			t.Errorf("Wanted nil got error")
			return
		}
	}

	for i, p := range processors {
		m := p.(*exportertest.SinkMetricsExporter)
		assert.Equal(t, wantMetricsCount, m.MetricsCount())
		metricOrig := md.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics().At(0)
		allMetrics := m.AllMetrics()
		metricClone := pdatautil.MetricsToInternalMetrics(allMetrics[0]).ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics().At(0)
		if i < len(processors)-1 {
			assert.True(t, md.ResourceMetrics().At(0).Resource() != pdatautil.MetricsToInternalMetrics(allMetrics[0]).ResourceMetrics().At(0).Resource())
			assert.True(t, metricOrig != metricClone)
		} else {
			assert.True(t, md.ResourceMetrics().At(0).Resource() == pdatautil.MetricsToInternalMetrics(allMetrics[0]).ResourceMetrics().At(0).Resource())
			assert.True(t, metricOrig == metricClone)
		}
		assert.EqualValues(t, md.ResourceMetrics().At(0).Resource(), pdatautil.MetricsToInternalMetrics(allMetrics[0]).ResourceMetrics().At(0).Resource())
		assert.EqualValues(t, metricOrig, metricClone)
	}
}
