// Copyright 2020, OpenTelemetry Authors
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

package splunkhecexporter

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"go.opentelemetry.io/collector/consumer/consumerdata"
	"go.uber.org/zap"
)

const (
	hecEventMetricType = "metric"
	unknownHostName    = "unknown"

	// Some standard dimension keys.
	// upper bound dimension key for histogram buckets.
	upperBoundDimensionKey = "upper_bound"
	// quantile dimension key for summary quantiles.
	quantileDimensionKey = "quantile"
)

type splunkMetric struct {
	Time       int64                  `json:"time"`                 // epoch time
	Host       string                 `json:"host"`                 // hostname
	Source     string                 `json:"source,omitempty"`     // optional description of the source of the event; typically the app's name
	SourceType string                 `json:"sourcetype,omitempty"` // optional name of a Splunk parsing configuration; this is usually inferred by Splunk
	Index      string                 `json:"index,omitempty"`      // optional name of the Splunk index to store the event in; not required if the token has a default index set in Splunk
	Event      string                 `json:"event"`                // type of event: this is a metric.
	Fields     map[string]interface{} `json:"fields"`               // metric data
}

func metricDataToSplunk(logger *zap.Logger, data consumerdata.MetricsData, config *Config) ([]*splunkMetric, int, error) {
	host := data.Resource.Labels["host.hostname"]
	if host == "" {
		host = unknownHostName
	}
	splunkMetrics := make([]*splunkMetric, 0)
	for _, metric := range data.Metrics {
		for _, timeSeries := range metric.Timeseries {
			for _, tsPoint := range timeSeries.Points {
				sm := &splunkMetric{
					Time:       timestampToEpochMilliseconds(tsPoint.GetTimestamp()),
					Host:       host,
					Source:     config.Source,
					SourceType: config.SourceType,
					Index:      config.Index,
					Event:      hecEventMetricType,
					Fields:     map[string]interface{}{}, // TODO fill fields
				}
				// TODO change metric_name computation.
				sm.Fields[fmt.Sprintf("metric_name:%s", data.Resource.Type)] = tsPoint.GetValue()
				splunkMetrics = append(splunkMetrics, sm)
			}
		}
	}

	return splunkMetrics, 0, nil
}

func timestampToEpochMilliseconds(ts *timestamp.Timestamp) int64 {
	if ts == nil {
		return 0
	}
	return ts.GetSeconds()*1e3 + int64(ts.GetNanos()/1e6)
}
