// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package sqlserverreceiver

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest/pmetrictest"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/winperfcounters"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlserverreceiver/internal/metadata"
)

// mockPerfCounterWatcher is an autogenerated mock type for the PerfCounterWatcher type
type mockPerfCounterWatcher struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *mockPerfCounterWatcher) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Path provides a mock function with given fields:
func (_m *mockPerfCounterWatcher) Path() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ScrapeData provides a mock function with given fields:
func (_m *mockPerfCounterWatcher) ScrapeData() ([]winperfcounters.CounterValue, error) {
	ret := _m.Called()

	var r0 []winperfcounters.CounterValue
	if rf, ok := ret.Get(0).(func() []winperfcounters.CounterValue); ok {
		r0 = rf()
	} else if ret.Get(0) != nil {
		r0 = ret.Get(0).([]winperfcounters.CounterValue)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *mockPerfCounterWatcher) Reset() error {
	return nil
}

func TestSqlServerScraper(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig().(*Config)
	logger, obsLogs := observer.New(zap.WarnLevel)
	settings := receivertest.NewNopSettings()
	settings.Logger = zap.New(logger)
	s := newSQLServerPCScraper(settings, cfg)

	assert.NoError(t, s.start(context.Background(), nil))
	assert.Len(t, s.watcherRecorders, 0)
	assert.Equal(t, 21, obsLogs.Len())
	assert.Equal(t, 21, obsLogs.FilterMessageSnippet("failed to create perf counter with path \\SQLServer:").Len())
	assert.Equal(t, 21, obsLogs.FilterMessageSnippet("The specified object was not found on the computer.").Len())
	assert.Equal(t, 1, obsLogs.FilterMessageSnippet("\\SQLServer:General Statistics\\").Len())
	assert.Equal(t, 3, obsLogs.FilterMessageSnippet("\\SQLServer:SQL Statistics\\").Len())
	assert.Equal(t, 2, obsLogs.FilterMessageSnippet("\\SQLServer:Locks(_Total)\\").Len())
	assert.Equal(t, 6, obsLogs.FilterMessageSnippet("\\SQLServer:Buffer Manager\\").Len())
	assert.Equal(t, 1, obsLogs.FilterMessageSnippet("\\SQLServer:Access Methods(_Total)\\").Len())
	assert.Equal(t, 8, obsLogs.FilterMessageSnippet("\\SQLServer:Databases(*)\\").Len())

	metrics, err := s.scrape(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 0, metrics.ResourceMetrics().Len())

	err = s.shutdown(context.Background())
	require.NoError(t, err)
}

var goldenScrapePath = filepath.Join("testdata", "golden_scrape.yaml")
var goldenNamedInstanceScrapePath = filepath.Join("testdata", "golden_named_instance_scrape.yaml")
var dbInstance = "db-instance"

func TestScrape(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		factory := NewFactory()
		cfg := factory.CreateDefaultConfig().(*Config)
		settings := receivertest.NewNopSettings()
		scraper := newSQLServerPCScraper(settings, cfg)

		for i, rec := range perfCounterRecorders {
			perfCounterWatcher := &mockPerfCounterWatcher{}
			perfCounterWatcher.On("ScrapeData").Return([]winperfcounters.CounterValue{{InstanceName: dbInstance, Value: float64(i)}}, nil)
			for _, recorder := range rec.recorders {
				scraper.watcherRecorders = append(scraper.watcherRecorders, watcherRecorder{
					watcher:  perfCounterWatcher,
					recorder: recorder,
				})
			}
		}

		scrapeData, err := scraper.scrape(context.Background())
		require.NoError(t, err)

		expectedMetrics, err := golden.ReadMetrics(goldenScrapePath)
		require.NoError(t, err)

		require.NoError(t, pmetrictest.CompareMetrics(expectedMetrics, scrapeData,
			pmetrictest.IgnoreMetricDataPointsOrder(), pmetrictest.IgnoreStartTimestamp(), pmetrictest.IgnoreTimestamp()))
	})

	t.Run("named", func(t *testing.T) {
		factory := NewFactory()
		cfg := factory.CreateDefaultConfig().(*Config)
		cfg.MetricsBuilderConfig = metadata.MetricsBuilderConfig{
			Metrics: metadata.DefaultMetricsConfig(),
			ResourceAttributes: metadata.ResourceAttributesConfig{
				SqlserverDatabaseName: metadata.ResourceAttributeConfig{
					Enabled: true,
				},
				SqlserverInstanceName: metadata.ResourceAttributeConfig{
					Enabled: true,
				},
				SqlserverComputerName: metadata.ResourceAttributeConfig{
					Enabled: true,
				},
			},
		}
		cfg.ComputerName = "CustomServer"
		cfg.InstanceName = "CustomInstance"

		settings := receivertest.NewNopSettings()
		scraper := newSQLServerPCScraper(settings, cfg)

		for i, rec := range perfCounterRecorders {
			perfCounterWatcher := &mockPerfCounterWatcher{}
			perfCounterWatcher.On("ScrapeData").Return([]winperfcounters.CounterValue{{InstanceName: dbInstance, Value: float64(i)}}, nil)
			for _, recorder := range rec.recorders {
				scraper.watcherRecorders = append(scraper.watcherRecorders, watcherRecorder{
					watcher:  perfCounterWatcher,
					recorder: recorder,
				})
			}
		}

		scrapeData, err := scraper.scrape(context.Background())
		require.NoError(t, err)

		expectedMetrics, err := golden.ReadMetrics(goldenNamedInstanceScrapePath)
		require.NoError(t, err)

		require.NoError(t, pmetrictest.CompareMetrics(expectedMetrics, scrapeData,
			pmetrictest.IgnoreMetricDataPointsOrder(), pmetrictest.IgnoreStartTimestamp(), pmetrictest.IgnoreTimestamp()))
	})
}
