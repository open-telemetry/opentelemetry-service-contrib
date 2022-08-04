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

package postgresqlreceiver

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/scrapertest"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/scrapertest/golden"
)

func TestUnsuccessfulScrape(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig().(*Config)
	cfg.Endpoint = "fake:11111"

	scraper := newPostgreSQLScraper(componenttest.NewNopReceiverCreateSettings(), cfg, &defaultClientFactory{})
	actualMetrics, err := scraper.scrape(context.Background())
	require.Error(t, err)

	require.NoError(t, scrapertest.CompareMetrics(pmetric.NewMetrics(), actualMetrics))
}

func TestScraper(t *testing.T) {
	factory := new(mockClientFactory)
	factory.initMocks([]string{"otel"})

	cfg := createDefaultConfig().(*Config)
	cfg.Databases = []string{"otel"}
	scraper := newPostgreSQLScraper(componenttest.NewNopReceiverCreateSettings(), cfg, factory)

	actualMetrics, err := scraper.scrape(context.Background())
	require.NoError(t, err)

	expectedFile := filepath.Join("testdata", "scraper", "otel", "expected.json")
	expectedMetrics, err := golden.ReadMetrics(expectedFile)
	require.NoError(t, err)

	require.NoError(t, scrapertest.CompareMetrics(expectedMetrics, actualMetrics))
}

func TestScraperNoDatabaseSingle(t *testing.T) {
	factory := new(mockClientFactory)
	factory.initMocks([]string{"otel"})

	cfg := createDefaultConfig().(*Config)
	scraper := newPostgreSQLScraper(componenttest.NewNopReceiverCreateSettings(), cfg, factory)

	actualMetrics, err := scraper.scrape(context.Background())
	require.NoError(t, err)

	expectedFile := filepath.Join("testdata", "scraper", "otel", "expected.json")
	expectedMetrics, err := golden.ReadMetrics(expectedFile)
	require.NoError(t, err)

	require.NoError(t, scrapertest.CompareMetrics(expectedMetrics, actualMetrics))
}

func TestScraperNoDatabaseMultiple(t *testing.T) {
	factory := mockClientFactory{}
	factory.initMocks([]string{"otel", "open", "telemetry"})

	cfg := createDefaultConfig().(*Config)
	scraper := newPostgreSQLScraper(componenttest.NewNopReceiverCreateSettings(), cfg, &factory)

	actualMetrics, err := scraper.scrape(context.Background())
	require.NoError(t, err)

	expectedFile := filepath.Join("testdata", "scraper", "multiple", "expected.json")
	expectedMetrics, err := golden.ReadMetrics(expectedFile)
	require.NoError(t, err)

	require.NoError(t, scrapertest.CompareMetrics(expectedMetrics, actualMetrics))
}

type mockClientFactory struct{ mock.Mock }
type mockClient struct{ mock.Mock }

var _ client = &mockClient{}

func (m *mockClient) Close() error {
	args := m.Called()
	return args.Error(0)
}

func (m *mockClient) getDatabaseStats(_ context.Context, databases []string) (map[string]databaseStats, error) {
	args := m.Called(databases)
	return args.Get(0).(map[string]databaseStats), args.Error(1)
}

func (m *mockClient) getBackends(_ context.Context, databases []string) (map[string]int64, error) {
	args := m.Called(databases)
	return args.Get(0).(map[string]int64), args.Error(1)
}

func (m *mockClient) getDatabaseSize(_ context.Context, databases []string) (map[string]int64, error) {
	args := m.Called(databases)
	return args.Get(0).(map[string]int64), args.Error(1)
}

func (m *mockClient) getDatabaseTableMetrics(ctx context.Context, database string) (map[string]tableStats, error) {
	args := m.Called(ctx, database)
	return args.Get(0).(map[string]tableStats), args.Error(1)
}

func (m *mockClient) getBlocksReadByTable(ctx context.Context, database string) (map[string]tableIOStats, error) {
	args := m.Called(ctx, database)
	return args.Get(0).(map[string]tableIOStats), args.Error(1)
}

func (m *mockClient) listDatabases(_ context.Context) ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}

func (m *mockClientFactory) getClient(c *Config, database string) (client, error) {
	args := m.Called(database)
	return args.Get(0).(client), args.Error(1)
}

func (m *mockClientFactory) initMocks(databases []string) {
	listClient := new(mockClient)
	listClient.initMocks("", databases, 0)
	m.On("getClient", "").Return(listClient, nil)

	for index, db := range databases {
		client := new(mockClient)
		client.initMocks(db, databases, index)
		m.On("getClient", db).Return(client, nil)
	}
}

func (m *mockClient) initMocks(database string, databases []string, index int) {
	m.On("Close").Return(nil)

	if database == "" {
		m.On("listDatabases").Return(databases, nil)

		commitsAndRollbacks := map[string]databaseStats{}
		dbSize := map[string]int64{}
		backends := map[string]int64{}

		for idx, db := range databases {
			commitsAndRollbacks[db] = databaseStats{
				transactionCommitted: int64(idx + 1),
				transactionRollback:  int64(idx + 2),
			}
			dbSize[db] = int64(idx + 4)
			backends[db] = int64(idx + 3)
		}

		m.On("getDatabaseStats", databases).Return(commitsAndRollbacks, nil)
		m.On("getDatabaseSize", databases).Return(dbSize, nil)
		m.On("getBackends", databases).Return(backends, nil)
	} else {
		table1 := "public.table1"
		table2 := "public.table2"
		tableMetrics := map[string]tableStats{
			tableKey(database, table1): {
				database: database,
				table:    table1,
				live:     int64(index + 7),
				dead:     int64(index + 8),
				inserts:  int64(89),
				upd:      int64(index + 40),
				del:      int64(index + 41),
				hotUpd:   int64(1),
			},
			tableKey(database, table2): {
				database: database,
				table:    table2,
				live:     int64(index + 7),
				dead:     int64(index + 8),
				inserts:  int64(89),
				upd:      int64(index + 40),
				del:      int64(index + 41),
				hotUpd:   int64(1),
			},
		}

		blocksMetrics := map[string]tableIOStats{
			tableKey(database, table1): {
				heapRead:  int64(index + 19),
				heapHit:   int64(index + 20),
				idxRead:   int64(index + 21),
				idxHit:    int64(index + 22),
				toastRead: int64(index + 23),
				toastHit:  int64(index + 24),
				tidxRead:  int64(index + 25),
				tidxHit:   int64(index + 26),
			},
			tableKey(database, table2): {
				heapRead:  int64(index + 27),
				heapHit:   int64(index + 28),
				idxRead:   int64(index + 29),
				idxHit:    int64(index + 30),
				toastRead: int64(index + 31),
				toastHit:  int64(index + 32),
				tidxRead:  int64(index + 33),
				tidxHit:   int64(index + 34),
			},
		}
		for _, db := range databases {
			m.On("getDatabaseTableMetrics", mock.Anything, db).Return(tableMetrics, nil)
			m.On("getBlocksReadByTable", mock.Anything, db).Return(blocksMetrics, nil)
		}
	}
}
