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

package pagingscraper

import (
	"context"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/model/pdata"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/pagingscraper/internal/metadata"
)

func TestScrape(t *testing.T) {
	type testCase struct {
		name              string
		config            Config
		bootTimeFunc      func() (uint64, error)
		expectedStartTime pdata.Timestamp
		initializationErr string
	}

	testCases := []testCase{
		{
			name:   "Standard",
			config: Config{Metrics: metadata.DefaultMetricsSettings()},
		},
		//{
		//	name:              "Validate Start Time",
		//	config:            Config{Metrics: metadata.DefaultMetricsSettings()},
		//	bootTimeFunc:      func() (uint64, error) { return 100, nil },
		//	expectedStartTime: 100 * 1e9,
		//},
		//{
		//	name:              "Boot Time Error",
		//	config:            Config{Metrics: metadata.DefaultMetricsSettings()},
		//	bootTimeFunc:      func() (uint64, error) { return 0, errors.New("err1") },
		//	initializationErr: "err1",
		//},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			scraper := newPagingScraper(context.Background(), &test.config)
			if test.bootTimeFunc != nil {
				scraper.bootTime = test.bootTimeFunc
			}

			err := scraper.start(context.Background(), componenttest.NewNopHost())
			if test.initializationErr != "" {
				assert.EqualError(t, err, test.initializationErr)
				return
			}
			require.NoError(t, err, "Failed to initialize paging scraper: %v", err)

			md, err := scraper.scrape(context.Background())
			require.NoError(t, err)
			metrics := md.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics()

			// expect 3 metrics (windows does not currently support the faults metric)
			expectedMetrics := 3
			if runtime.GOOS == "windows" {
				expectedMetrics = 2
			}
			assert.Equal(t, expectedMetrics, md.MetricCount())

			startIndex := 0
			if runtime.GOOS != "windows" {
				//assertPageFaultsMetricValid(t, metrics.At(startIndex), test.expectedStartTime)
				//startIndex++
			}

			assertPagingOperationsMetricValid(t, metrics.At(startIndex), test.expectedStartTime)
			internal.AssertSameTimeStampForMetrics(t, metrics, 0, metrics.Len()-1)
			startIndex++

			assertPagingUsageMetricValid(t, metrics.At(startIndex))
			internal.AssertSameTimeStampForMetrics(t, metrics, startIndex, metrics.Len())
		})
	}
}

func assertPagingUsageMetricValid(t *testing.T, hostPagingUsageMetric pdata.Metric) {
	expected := pdata.NewMetric()
	expected.SetName("system.paging.usage")
	expected.SetDescription("Swap (unix) or pagefile (windows) usage.")
	expected.SetUnit("By")
	expected.SetDataType(pdata.MetricDataTypeSum)
	internal.AssertDescriptorEqual(t, expected, hostPagingUsageMetric)

	// it's valid for a system to have no swap space  / paging file, so if no data points were returned, do no validation
	if hostPagingUsageMetric.Sum().DataPoints().Len() == 0 {
		return
	}

	// expect at least used, free & cached datapoint
	expectedDataPoints := 3
	// windows does not return a cached datapoint
	if runtime.GOOS == "windows" || runtime.GOOS == "linux" {
		expectedDataPoints = 2
	}

	assert.GreaterOrEqual(t, hostPagingUsageMetric.Sum().DataPoints().Len(), expectedDataPoints)
	internal.AssertSumMetricHasAttributeValue(t, hostPagingUsageMetric, 0, "state", pdata.NewAttributeValueString(metadata.AttributeState.Used))
	internal.AssertSumMetricHasAttributeValue(t, hostPagingUsageMetric, 1, "state", pdata.NewAttributeValueString(metadata.AttributeState.Free))
	// Windows and Linux do not support cached state label
	if runtime.GOOS != "windows" && runtime.GOOS != "linux" {
		internal.AssertSumMetricHasAttributeValue(t, hostPagingUsageMetric, 2, "state", pdata.NewAttributeValueString(metadata.AttributeState.Cached))
	}

	// on Windows and Linux, also expect the page file device name label
	if runtime.GOOS == "windows" || runtime.GOOS == "linux" {
		internal.AssertSumMetricHasAttribute(t, hostPagingUsageMetric, 0, "device")
		internal.AssertSumMetricHasAttribute(t, hostPagingUsageMetric, 1, "device")
	}
}

func assertPagingOperationsMetricValid(t *testing.T, pagingMetric pdata.Metric, startTime pdata.Timestamp) {
	expected := pdata.NewMetric()
	expected.SetName("system.paging.operations")
	expected.SetDescription("The number of paging operations.")
	expected.SetUnit("{operations}")
	expected.SetDataType(pdata.MetricDataTypeSum)
	internal.AssertDescriptorEqual(t, expected, pagingMetric)

	if startTime != 0 {
		internal.AssertSumMetricStartTimeEquals(t, pagingMetric, startTime)
	}

	// expect an in & out datapoint, for both major and minor paging types (windows does not currently support minor paging data)
	expectedDataPoints := 4
	if runtime.GOOS == "windows" {
		expectedDataPoints = 2
	}
	assert.Equal(t, expectedDataPoints, pagingMetric.Sum().DataPoints().Len())

	internal.AssertSumMetricHasAttributeValue(t, pagingMetric, 0, "type", pdata.NewAttributeValueString(metadata.AttributeType.Major))
	internal.AssertSumMetricHasAttributeValue(t, pagingMetric, 0, "direction", pdata.NewAttributeValueString(metadata.AttributeDirection.PageIn))
	internal.AssertSumMetricHasAttributeValue(t, pagingMetric, 1, "type", pdata.NewAttributeValueString(metadata.AttributeType.Major))
	internal.AssertSumMetricHasAttributeValue(t, pagingMetric, 1, "direction", pdata.NewAttributeValueString(metadata.AttributeDirection.PageOut))
	if runtime.GOOS != "windows" {
		internal.AssertSumMetricHasAttributeValue(t, pagingMetric, 2, "type", pdata.NewAttributeValueString(metadata.AttributeType.Minor))
		internal.AssertSumMetricHasAttributeValue(t, pagingMetric, 2, "direction", pdata.NewAttributeValueString(metadata.AttributeDirection.PageIn))
		internal.AssertSumMetricHasAttributeValue(t, pagingMetric, 3, "type", pdata.NewAttributeValueString(metadata.AttributeType.Minor))
		internal.AssertSumMetricHasAttributeValue(t, pagingMetric, 3, "direction", pdata.NewAttributeValueString(metadata.AttributeDirection.PageOut))
	}
}

func assertPageFaultsMetricValid(t *testing.T, pageFaultsMetric pdata.Metric, startTime pdata.Timestamp) {
	expected := pdata.NewMetric()
	expected.SetName("system.paging.faults")
	expected.SetDescription("The number of page faults.")
	expected.SetUnit("{faults}")
	expected.SetDataType(pdata.MetricDataTypeSum)
	internal.AssertDescriptorEqual(t, expected, pageFaultsMetric)

	if startTime != 0 {
		internal.AssertSumMetricStartTimeEquals(t, pageFaultsMetric, startTime)
	}

	assert.Equal(t, 2, pageFaultsMetric.Sum().DataPoints().Len())
	internal.AssertSumMetricHasAttributeValue(t, pageFaultsMetric, 0, "type", pdata.NewAttributeValueString(metadata.AttributeType.Major))
	internal.AssertSumMetricHasAttributeValue(t, pageFaultsMetric, 1, "type", pdata.NewAttributeValueString(metadata.AttributeType.Minor))
}
