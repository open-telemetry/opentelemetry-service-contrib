// Code generated by mdatagen. DO NOT EDIT.

package azuremonitorexporter

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/confmap/confmaptest"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exportertest"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

func TestComponentFactoryType(t *testing.T) {
	require.Equal(t, "azuremonitor", NewFactory().Type().String())
}

func TestComponentConfigStruct(t *testing.T) {
	require.NoError(t, componenttest.CheckConfigStruct(NewFactory().CreateDefaultConfig()))
}

func TestComponentLifecycle(t *testing.T) {
	factory := NewFactory()

	tests := []struct {
		name     string
		createFn func(ctx context.Context, set exporter.CreateSettings, cfg component.Config) (component.Component, error)
	}{

		{
			name: "logs",
			createFn: func(ctx context.Context, set exporter.CreateSettings, cfg component.Config) (component.Component, error) {
				return factory.CreateLogsExporter(ctx, set, cfg)
			},
		},

		{
			name: "metrics",
			createFn: func(ctx context.Context, set exporter.CreateSettings, cfg component.Config) (component.Component, error) {
				return factory.CreateMetricsExporter(ctx, set, cfg)
			},
		},

		{
			name: "traces",
			createFn: func(ctx context.Context, set exporter.CreateSettings, cfg component.Config) (component.Component, error) {
				return factory.CreateTracesExporter(ctx, set, cfg)
			},
		},
	}

	cm, err := confmaptest.LoadConf("metadata.yaml")
	require.NoError(t, err)
	cfg := factory.CreateDefaultConfig()
	sub, err := cm.Sub("tests::config")
	require.NoError(t, err)
	require.NoError(t, sub.Unmarshal(&cfg))

	for _, test := range tests {
		t.Run(test.name+"-shutdown", func(t *testing.T) {
			c, err := test.createFn(context.Background(), exportertest.NewNopCreateSettings(), cfg)
			require.NoError(t, err)
			err = c.Shutdown(context.Background())
			require.NoError(t, err)
		})
		t.Run(test.name+"-lifecycle", func(t *testing.T) {
			c, err := test.createFn(context.Background(), exportertest.NewNopCreateSettings(), cfg)
			require.NoError(t, err)
			host := componenttest.NewNopHost()
			err = c.Start(context.Background(), host)
			require.NoError(t, err)
			require.NotPanics(t, func() {
				switch test.name {
				case "logs":
					e, ok := c.(exporter.Logs)
					require.True(t, ok)
					logs := generateLifecycleTestLogs()
					if !e.Capabilities().MutatesData {
						logs.MarkReadOnly()
					}
					err = e.ConsumeLogs(context.Background(), logs)
				case "metrics":
					e, ok := c.(exporter.Metrics)
					require.True(t, ok)
					metrics := generateLifecycleTestMetrics()
					if !e.Capabilities().MutatesData {
						metrics.MarkReadOnly()
					}
					err = e.ConsumeMetrics(context.Background(), metrics)
				case "traces":
					e, ok := c.(exporter.Traces)
					require.True(t, ok)
					traces := generateLifecycleTestTraces()
					if !e.Capabilities().MutatesData {
						traces.MarkReadOnly()
					}
					err = e.ConsumeTraces(context.Background(), traces)
				}
			})

			err = c.Shutdown(context.Background())
			require.NoError(t, err)
		})
	}
}

func generateLifecycleTestLogs() plog.Logs {
	logs := plog.NewLogs()
	rl := logs.ResourceLogs().AppendEmpty()
	rl.Resource().Attributes().PutStr("resource", "R1")
	l := rl.ScopeLogs().AppendEmpty().LogRecords().AppendEmpty()
	l.Body().SetStr("test log message")
	l.SetTimestamp(pcommon.NewTimestampFromTime(time.Now()))
	return logs
}

func generateLifecycleTestMetrics() pmetric.Metrics {
	metrics := pmetric.NewMetrics()
	rm := metrics.ResourceMetrics().AppendEmpty()
	rm.Resource().Attributes().PutStr("resource", "R1")
	m := rm.ScopeMetrics().AppendEmpty().Metrics().AppendEmpty()
	m.SetName("test_metric")
	dp := m.SetEmptyGauge().DataPoints().AppendEmpty()
	dp.Attributes().PutStr("test_attr", "value_1")
	dp.SetIntValue(123)
	dp.SetTimestamp(pcommon.NewTimestampFromTime(time.Now()))
	return metrics
}

func generateLifecycleTestTraces() ptrace.Traces {
	traces := ptrace.NewTraces()
	rs := traces.ResourceSpans().AppendEmpty()
	rs.Resource().Attributes().PutStr("resource", "R1")
	span := rs.ScopeSpans().AppendEmpty().Spans().AppendEmpty()
	span.Attributes().PutStr("test_attr", "value_1")
	span.SetName("test_span")
	span.SetStartTimestamp(pcommon.NewTimestampFromTime(time.Now().Add(-1 * time.Second)))
	span.SetEndTimestamp(pcommon.NewTimestampFromTime(time.Now()))
	return traces
}
