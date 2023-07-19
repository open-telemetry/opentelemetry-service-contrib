// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

type testConfigCollection int

const (
	testSetDefault testConfigCollection = iota
	testSetAll
	testSetNone
)

func TestMetricsBuilder(t *testing.T) {
	tests := []struct {
		name      string
		configSet testConfigCollection
	}{
		{
			name:      "default",
			configSet: testSetDefault,
		},
		{
			name:      "all_set",
			configSet: testSetAll,
		},
		{
			name:      "none_set",
			configSet: testSetNone,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			start := pcommon.Timestamp(1_000_000_000)
			ts := pcommon.Timestamp(1_000_001_000)
			observedZapCore, observedLogs := observer.New(zap.WarnLevel)
			settings := receivertest.NewNopCreateSettings()
			settings.Logger = zap.New(observedZapCore)
			mb := NewMetricsBuilder(loadMetricsBuilderConfig(t, test.name), settings, WithStartTime(start))

			expectedWarnings := 0
			assert.Equal(t, expectedWarnings, observedLogs.Len())

			defaultMetricsCount := 0
			allMetricsCount := 0

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordK8sPodPhaseDataPoint(ts, 1)

			metrics := mb.Emit(WithK8sNamespaceName("k8s.namespace.name-val"), WithK8sNodeName("k8s.node.name-val"), WithK8sPodName("k8s.pod.name-val"), WithK8sPodUID("k8s.pod.uid-val"), WithOpencensusResourcetype("opencensus.resourcetype-val"))

			if test.configSet == testSetNone {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			attrCount := 0
			enabledAttrCount := 0
			attrVal, ok := rm.Resource().Attributes().Get("k8s.namespace.name")
			attrCount++
			assert.Equal(t, mb.resourceAttributesConfig.K8sNamespaceName.Enabled, ok)
			if mb.resourceAttributesConfig.K8sNamespaceName.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "k8s.namespace.name-val", attrVal.Str())
			}
			attrVal, ok = rm.Resource().Attributes().Get("k8s.node.name")
			attrCount++
			assert.Equal(t, mb.resourceAttributesConfig.K8sNodeName.Enabled, ok)
			if mb.resourceAttributesConfig.K8sNodeName.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "k8s.node.name-val", attrVal.Str())
			}
			attrVal, ok = rm.Resource().Attributes().Get("k8s.pod.name")
			attrCount++
			assert.Equal(t, mb.resourceAttributesConfig.K8sPodName.Enabled, ok)
			if mb.resourceAttributesConfig.K8sPodName.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "k8s.pod.name-val", attrVal.Str())
			}
			attrVal, ok = rm.Resource().Attributes().Get("k8s.pod.uid")
			attrCount++
			assert.Equal(t, mb.resourceAttributesConfig.K8sPodUID.Enabled, ok)
			if mb.resourceAttributesConfig.K8sPodUID.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "k8s.pod.uid-val", attrVal.Str())
			}
			attrVal, ok = rm.Resource().Attributes().Get("opencensus.resourcetype")
			attrCount++
			assert.Equal(t, mb.resourceAttributesConfig.OpencensusResourcetype.Enabled, ok)
			if mb.resourceAttributesConfig.OpencensusResourcetype.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "opencensus.resourcetype-val", attrVal.Str())
			}
			assert.Equal(t, enabledAttrCount, rm.Resource().Attributes().Len())
			assert.Equal(t, attrCount, 5)

			assert.Equal(t, 1, rm.ScopeMetrics().Len())
			ms := rm.ScopeMetrics().At(0).Metrics()
			if test.configSet == testSetDefault {
				assert.Equal(t, defaultMetricsCount, ms.Len())
			}
			if test.configSet == testSetAll {
				assert.Equal(t, allMetricsCount, ms.Len())
			}
			validatedMetrics := make(map[string]bool)
			for i := 0; i < ms.Len(); i++ {
				switch ms.At(i).Name() {
				case "k8s.pod.phase":
					assert.False(t, validatedMetrics["k8s.pod.phase"], "Found a duplicate in the metrics slice: k8s.pod.phase")
					validatedMetrics["k8s.pod.phase"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Current phase of the pod (1 - Pending, 2 - Running, 3 - Succeeded, 4 - Failed, 5 - Unknown)", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				}
			}
		})
	}
}
