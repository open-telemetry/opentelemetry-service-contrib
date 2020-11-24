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

package signalfxreceiver

import (
	"strconv"
	"testing"
	"time"

	sfxpb "github.com/signalfx/com_signalfx_metrics_protobuf/model"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/consumer/pdata"
	"go.uber.org/zap"
)

func Test_signalFxV2ToMetricsData(t *testing.T) {
	now := time.Now()

	buildDefaulstSFxDataPt := func() *sfxpb.DataPoint {
		return &sfxpb.DataPoint{
			Metric:    "single",
			Timestamp: now.UnixNano() / 1e6,
			Value: sfxpb.Datum{
				IntValue: int64Ptr(13),
			},
			MetricType: sfxTypePtr(sfxpb.MetricType_GAUGE),
			Dimensions: buildNDimensions(3),
		}
	}

	buildDefaultMetricsData := func(typ pdata.MetricDataType, val interface{}) pdata.Metrics {
		out := pdata.NewMetrics()
		out.ResourceMetrics().Resize(1)
		rm := out.ResourceMetrics().At(0)
		rm.InitEmpty()
		rm.InstrumentationLibraryMetrics().Resize(1)
		ilm := rm.InstrumentationLibraryMetrics().At(0)
		ms := ilm.Metrics()

		ms.Resize(1)
		m := ms.At(0)
		m.InitEmpty()

		m.SetDataType(typ)
		m.SetName("single")

		var dps interface{}

		switch typ {
		case pdata.MetricDataTypeIntGauge:
			dps = m.IntGauge().DataPoints()
		case pdata.MetricDataTypeIntSum:
			m.IntSum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
			dps = m.IntSum().DataPoints()
		case pdata.MetricDataTypeDoubleGauge:
			dps = m.DoubleGauge().DataPoints()
		case pdata.MetricDataTypeDoubleSum:
			m.DoubleSum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
			dps = m.DoubleSum().DataPoints()
		}

		var labels pdata.StringMap

		switch typ {
		case pdata.MetricDataTypeIntGauge, pdata.MetricDataTypeIntSum:
			dps.(pdata.IntDataPointSlice).Resize(1)
			dp := dps.(pdata.IntDataPointSlice).At(0)
			labels = dp.LabelsMap()
			dp.SetTimestamp(pdata.TimestampUnixNano(now.Truncate(time.Millisecond).UnixNano()))
			dp.SetValue(int64(val.(int)))
		case pdata.MetricDataTypeDoubleGauge, pdata.MetricDataTypeDoubleSum:
			dps.(pdata.DoubleDataPointSlice).Resize(1)
			dp := dps.(pdata.DoubleDataPointSlice).At(0)
			labels = dp.LabelsMap()
			dp.SetTimestamp(pdata.TimestampUnixNano(now.Truncate(time.Millisecond).UnixNano()))
			dp.SetValue(val.(float64))
		}

		labels.InitFromMap(map[string]string{
			"k0": "v0",
			"k1": "v1",
			"k2": "v2",
		})
		labels.Sort()

		return out
	}

	tests := []struct {
		name                  string
		sfxDataPoints         []*sfxpb.DataPoint
		wantMetricsData       pdata.Metrics
		wantDroppedTimeseries int
	}{
		{
			name:            "int_gauge",
			sfxDataPoints:   []*sfxpb.DataPoint{buildDefaulstSFxDataPt()},
			wantMetricsData: buildDefaultMetricsData(pdata.MetricDataTypeIntGauge, 13),
		},
		{
			name: "double_gauge",
			sfxDataPoints: func() []*sfxpb.DataPoint {
				pt := buildDefaulstSFxDataPt()
				pt.MetricType = sfxTypePtr(sfxpb.MetricType_GAUGE)
				pt.Value = sfxpb.Datum{
					DoubleValue: float64Ptr(13.13),
				}
				return []*sfxpb.DataPoint{pt}
			}(),
			wantMetricsData: buildDefaultMetricsData(pdata.MetricDataTypeDoubleGauge, 13.13),
		},
		{
			name: "int_counter",
			sfxDataPoints: func() []*sfxpb.DataPoint {
				pt := buildDefaulstSFxDataPt()
				pt.MetricType = sfxTypePtr(sfxpb.MetricType_COUNTER)
				return []*sfxpb.DataPoint{pt}
			}(),
			wantMetricsData: func() pdata.Metrics {
				m := buildDefaultMetricsData(pdata.MetricDataTypeIntSum, 13)
				d := m.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics().At(0).IntSum()
				d.SetAggregationTemporality(pdata.AggregationTemporalityDelta)
				d.SetIsMonotonic(true)
				return m
			}(),
		},
		{
			name: "double_counter",
			sfxDataPoints: func() []*sfxpb.DataPoint {
				pt := buildDefaulstSFxDataPt()
				pt.MetricType = sfxTypePtr(sfxpb.MetricType_COUNTER)
				pt.Value = sfxpb.Datum{
					DoubleValue: float64Ptr(13.13),
				}
				return []*sfxpb.DataPoint{pt}
			}(),
			wantMetricsData: func() pdata.Metrics {
				m := buildDefaultMetricsData(pdata.MetricDataTypeDoubleSum, 13.13)
				d := m.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics().At(0).DoubleSum()
				d.SetAggregationTemporality(pdata.AggregationTemporalityDelta)
				d.SetIsMonotonic(true)
				return m
			}(),
		},
		{
			name: "nil_timestamp",
			sfxDataPoints: func() []*sfxpb.DataPoint {
				pt := buildDefaulstSFxDataPt()
				pt.Timestamp = 0
				return []*sfxpb.DataPoint{pt}
			}(),
			wantMetricsData: func() pdata.Metrics {
				md := buildDefaultMetricsData(pdata.MetricDataTypeIntGauge, 13)
				md.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics().At(0).IntGauge().DataPoints().At(0).SetTimestamp(0)
				return md
			}(),
		},
		{
			name: "empty_dimension_value",
			sfxDataPoints: func() []*sfxpb.DataPoint {
				pt := buildDefaulstSFxDataPt()
				pt.Dimensions[0].Value = ""
				return []*sfxpb.DataPoint{pt}
			}(),
			wantMetricsData: func() pdata.Metrics {
				md := buildDefaultMetricsData(pdata.MetricDataTypeIntGauge, 13)
				md.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics().At(0).IntGauge().DataPoints().At(0).LabelsMap().Update("k0", "")
				return md
			}(),
		},
		{
			name: "nil_dimension_ignored",
			sfxDataPoints: func() []*sfxpb.DataPoint {
				pt := buildDefaulstSFxDataPt()
				targetLen := 2*len(pt.Dimensions) + 1
				dimensions := make([]*sfxpb.Dimension, targetLen)
				copy(dimensions[1:], pt.Dimensions)
				assert.Equal(t, targetLen, len(dimensions))
				assert.Nil(t, dimensions[0])
				pt.Dimensions = dimensions
				return []*sfxpb.DataPoint{pt}
			}(),
			wantMetricsData: buildDefaultMetricsData(pdata.MetricDataTypeIntGauge, 13),
		},
		{
			name:            "nil_datapoint_ignored",
			sfxDataPoints:   []*sfxpb.DataPoint{nil, buildDefaulstSFxDataPt(), nil},
			wantMetricsData: buildDefaultMetricsData(pdata.MetricDataTypeIntGauge, 13),
		},
		{
			name: "drop_inconsistent_datapoints",
			sfxDataPoints: func() []*sfxpb.DataPoint {
				// nil Datum
				pt0 := buildDefaulstSFxDataPt()
				pt0.Value = sfxpb.Datum{}

				// nil expected Datum value
				pt1 := buildDefaulstSFxDataPt()
				pt1.Value.IntValue = nil

				// Non-supported type
				pt2 := buildDefaulstSFxDataPt()
				pt2.MetricType = sfxTypePtr(sfxpb.MetricType_ENUM)

				// Unknown type
				pt3 := buildDefaulstSFxDataPt()
				pt3.MetricType = sfxTypePtr(sfxpb.MetricType_CUMULATIVE_COUNTER + 1)

				return []*sfxpb.DataPoint{
					pt0, buildDefaulstSFxDataPt(), pt1, pt2, pt3}
			}(),
			wantMetricsData:       buildDefaultMetricsData(pdata.MetricDataTypeIntGauge, 13),
			wantDroppedTimeseries: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md, numDroppedTimeseries := signalFxV2ToMetrics(zap.NewNop(), tt.sfxDataPoints)
			assert.Equal(t, tt.wantMetricsData, md)
			assert.Equal(t, tt.wantDroppedTimeseries, numDroppedTimeseries)
		})
	}
}

func strPtr(s string) *string {
	return &s
}

func int64Ptr(i int64) *int64 {
	return &i
}

func float64Ptr(f float64) *float64 {
	return &f
}

func sfxTypePtr(t sfxpb.MetricType) *sfxpb.MetricType {
	return &t
}

func sfxCategoryPtr(t sfxpb.EventCategory) *sfxpb.EventCategory {
	return &t
}

func buildNDimensions(n uint) []*sfxpb.Dimension {
	d := make([]*sfxpb.Dimension, 0, n)
	for i := uint(0); i < n; i++ {
		idx := int(i)
		suffix := strconv.Itoa(idx)
		d = append(d, &sfxpb.Dimension{
			Key:   "k" + suffix,
			Value: "v" + suffix,
		})
	}
	return d
}
