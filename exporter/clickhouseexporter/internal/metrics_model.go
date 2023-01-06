// Copyright  The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/clickhouseexporter/internal"

import (
	"context"
	"database/sql"
	"encoding/hex"
	"fmt"
	"reflect"
	"sync"

	"github.com/ClickHouse/clickhouse-go/v2"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

var supportMetricsType = [...]string{createGaugeTableSQL, createSumTableSQL, createHistogramTableSQL, createExpHistogramTableSQL, createSummaryTableSQL}

type metricsModel interface {
	Add(metrics any, metaData *MetricsMetaData, name string, description string, unit string) error
	insert(ctx context.Context, tx *sql.Tx, logger *zap.Logger) error
}

type MetricsMetaData struct {
	ResAttr    map[string]string
	ResURL     string
	ScopeURL   string
	ScopeInstr pcommon.InstrumentationScope
}

func CreateMetricsTable(tableName string, ttlDays uint, db *sql.DB) error {
	var ttlExpr string
	if ttlDays > 0 {
		ttlExpr = fmt.Sprintf(`TTL toDateTime(TimeUnix) + toIntervalDay(%d)`, ttlDays)
	}
	for _, table := range supportMetricsType {
		query := fmt.Sprintf(table, tableName, ttlExpr)
		if _, err := db.Exec(query); err != nil {
			return fmt.Errorf("exec create metrics table sql: %w", err)
		}
	}
	return nil
}

func CreateMetricsModel(tableName string) map[pmetric.MetricType]metricsModel {
	metricsMap := make(map[pmetric.MetricType]metricsModel)
	metricsMap[pmetric.MetricTypeGauge] = &gaugeMetrics{
		insertSQL: fmt.Sprintf(insertGaugeTableSQL, tableName),
	}
	metricsMap[pmetric.MetricTypeSum] = &sumMetrics{
		insertSQL: fmt.Sprintf(insertSumTableSQL, tableName),
	}
	metricsMap[pmetric.MetricTypeHistogram] = &histogramMetrics{
		insertSQL: fmt.Sprintf(insertHistogramTableSQL, tableName),
	}
	metricsMap[pmetric.MetricTypeExponentialHistogram] = &expHistogramMetrics{
		insertSQL: fmt.Sprintf(insertExpHistogramTableSQL, tableName),
	}
	metricsMap[pmetric.MetricTypeSummary] = &summaryMetrics{
		insertSQL: fmt.Sprintf(insertSummaryTableSQL, tableName),
	}
	return metricsMap
}

func InsertMetrics(ctx context.Context, tx *sql.Tx, metricsMap map[pmetric.MetricType]metricsModel, logger *zap.Logger) error {
	errsChan := make(chan error, 5)
	wg := &sync.WaitGroup{}
	for _, m := range metricsMap {
		wg.Add(1)
		go func(m metricsModel, wg *sync.WaitGroup) {
			errsChan <- m.insert(ctx, tx, logger)
			wg.Done()
		}(m, wg)
	}
	wg.Wait()
	close(errsChan)
	var err []error
	for e := range errsChan {
		err = append(err, e)
	}
	return multierr.Combine(err...)
}

func convertExemplars(exemplars pmetric.ExemplarSlice) (clickhouse.ArraySet, clickhouse.ArraySet, clickhouse.ArraySet, clickhouse.ArraySet, clickhouse.ArraySet) {
	var (
		attrs    clickhouse.ArraySet
		times    clickhouse.ArraySet
		values   clickhouse.ArraySet
		traceIDs clickhouse.ArraySet
		spanIDs  clickhouse.ArraySet
	)
	for i := 0; i < exemplars.Len(); i++ {
		exemplar := exemplars.At(i)
		attrs = append(attrs, attributesToMap(exemplar.FilteredAttributes()))
		times = append(times, exemplar.Timestamp().AsTime())
		values = append(values, getValue(exemplar.IntValue(), exemplar.DoubleValue()))

		traceID, spanID := exemplar.TraceID(), exemplar.SpanID()
		traceIDs = append(traceIDs, hex.EncodeToString(traceID[:]))
		spanIDs = append(spanIDs, hex.EncodeToString(spanID[:]))
	}
	return attrs, times, values, traceIDs, spanIDs
}

// https://github.com/open-telemetry/opentelemetry-proto/blob/main/opentelemetry/proto/metrics/v1/metrics.proto#L358
// define two types for one datapoint value, clickhouse only use one value of float64 to store them
func getValue(intValue int64, floatValue float64) float64 {
	if intValue > 0 {
		return float64(intValue)
	}
	return floatValue
}

// Only support converting string type in pcommon.Value, other type will convert into an empty string
func attributesToMap(attributes pcommon.Map) map[string]string {
	m := make(map[string]string, attributes.Len())
	attributes.Range(func(k string, v pcommon.Value) bool {
		m[k] = v.Str()
		return true
	})
	return m
}

func convertSliceToArraySet(slice interface{}, logger *zap.Logger) clickhouse.ArraySet {
	var set clickhouse.ArraySet
	switch slice := slice.(type) {
	case []uint64:
		for _, item := range slice {
			set = append(set, item)
		}
	case []float64:
		for _, item := range slice {
			set = append(set, item)
		}
	default:
		logger.Warn("unsupported slice type", zap.String("type", reflect.TypeOf(slice).String()))
	}
	return set
}

func convertValueAtQuantile(valueAtQuantile pmetric.SummaryDataPointValueAtQuantileSlice) (clickhouse.ArraySet, clickhouse.ArraySet) {
	var (
		quantiles clickhouse.ArraySet
		values    clickhouse.ArraySet
	)
	for i := 0; i < valueAtQuantile.Len(); i++ {
		value := valueAtQuantile.At(i)
		quantiles = append(quantiles, value.Quantile())
		values = append(values, value.Value())
	}
	return quantiles, values
}
