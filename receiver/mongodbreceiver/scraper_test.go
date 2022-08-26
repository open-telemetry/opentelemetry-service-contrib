// Copyright The OpenTelemetry Authors
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

package mongodbreceiver

import (
	"context"
	"errors"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/go-version"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/scrapererror"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/scrapertest"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/scrapertest/golden"
)

func TestNewMongodbScraper(t *testing.T) {
	f := NewFactory()
	cfg := f.CreateDefaultConfig().(*Config)

	scraper := newMongodbScraper(componenttest.NewNopReceiverCreateSettings(), cfg)
	require.NotEmpty(t, scraper.config.hostlist())
}

func TestScraperLifecycle(t *testing.T) {
	now := time.Now()
	f := NewFactory()
	cfg := f.CreateDefaultConfig().(*Config)

	scraper := newMongodbScraper(componenttest.NewNopReceiverCreateSettings(), cfg)
	require.NoError(t, scraper.start(context.Background(), componenttest.NewNopHost()))
	require.NoError(t, scraper.shutdown(context.Background()))

	require.Less(t, time.Since(now), 100*time.Millisecond, "component start and stop should be very fast")
}

var (
	errAllPartialMetrics = errors.New(
		strings.Join(
			[]string{
				"failed to collect metric mongodb.cache.operations with attribute(s) miss, hit: could not find key for metric",
				"failed to collect metric mongodb.cursor.count: could not find key for metric",
				"failed to collect metric mongodb.cursor.timeout.count: could not find key for metric",
				"failed to collect metric mongodb.global_lock.time: could not find key for metric",
				"failed to collect metric bytesIn: could not find key for metric",
				"failed to collect metric bytesOut: could not find key for metric",
				"failed to collect metric numRequests: could not find key for metric",
				"failed to collect metric mongodb.operation.count with attribute(s) delete: could not find key for metric",
				"failed to collect metric mongodb.operation.count with attribute(s) getmore: could not find key for metric",
				"failed to collect metric mongodb.operation.count with attribute(s) command: could not find key for metric",
				"failed to collect metric mongodb.operation.count with attribute(s) insert: could not find key for metric",
				"failed to collect metric mongodb.operation.count with attribute(s) query: could not find key for metric",
				"failed to collect metric mongodb.operation.count with attribute(s) update: could not find key for metric",
				"failed to collect metric mongodb.session.count: could not find key for metric",
				"failed to collect metric mongodb.operation.time: could not find key for metric",
				"failed to collect metric mongodb.collection.count with attribute(s) fakedatabase: could not find key for metric",
				"failed to collect metric mongodb.data.size with attribute(s) fakedatabase: could not find key for metric",
				"failed to collect metric mongodb.extent.count with attribute(s) fakedatabase: could not find key for metric",
				"failed to collect metric mongodb.index.size with attribute(s) fakedatabase: could not find key for metric",
				"failed to collect metric mongodb.index.count with attribute(s) fakedatabase: could not find key for metric",
				"failed to collect metric mongodb.object.count with attribute(s) fakedatabase: could not find key for metric",
				"failed to collect metric mongodb.storage.size with attribute(s) fakedatabase: could not find key for metric",
				"failed to collect metric mongodb.connection.count with attribute(s) available, fakedatabase: could not find key for metric",
				"failed to collect metric mongodb.connection.count with attribute(s) current, fakedatabase: could not find key for metric",
				"failed to collect metric mongodb.connection.count with attribute(s) active, fakedatabase: could not find key for metric",
				"failed to collect metric mongodb.document.operation.count with attribute(s) inserted, fakedatabase: could not find key for metric",
				"failed to collect metric mongodb.document.operation.count with attribute(s) updated, fakedatabase: could not find key for metric",
				"failed to collect metric mongodb.document.operation.count with attribute(s) deleted, fakedatabase: could not find key for metric",
				"failed to collect metric mongodb.memory.usage with attribute(s) resident, fakedatabase: could not find key for metric",
				"failed to collect metric mongodb.memory.usage with attribute(s) virtual, fakedatabase: could not find key for metric",
				"failed to collect metric mongodb.index.access.count with attribute(s) fakedatabase, orders: could not find key for index access metric",
				"failed to collect metric mongodb.index.access.count with attribute(s) fakedatabase, products: could not find key for index access metric",
				"failed to collect metric mongodb.global_lock.active_clients with attribute(s) readers: could not find key for metric",
				"failed to collect metric mongodb.global_lock.active_clients with attribute(s) writers: could not find key for metric",
				"failed to collect metric mongodb.global_lock.active_clients with attribute(s) total: could not find key for metric",
				"failed to collect metric mongodb.global_lock.current_queue with attribute(s) readers: could not find key for metric",
				"failed to collect metric mongodb.global_lock.current_queue with attribute(s) writers: could not find key for metric",
				"failed to collect metric mongodb.global_lock.current_queue with attribute(s) total: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Collection, R: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Collection, W: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Collection, r: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Collection, w: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Database, R: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Database, W: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Database, r: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Database, w: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Global, R: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Global, W: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Global, r: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Global, w: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Metadata, R: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Metadata, W: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Metadata, r: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Metadata, w: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Mutex, R: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Mutex, W: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Mutex, r: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Mutex, w: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, ParallelBatchWriterMode, R: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, ParallelBatchWriterMode, W: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, ParallelBatchWriterMode, r: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, ParallelBatchWriterMode, w: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, ReplicationStateTransition, R: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, ReplicationStateTransition, W: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, ReplicationStateTransition, r: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, ReplicationStateTransition, w: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, oplog, R: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, oplog, W: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, oplog, r: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, oplog, w: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Collection, R: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Collection, W: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Collection, r: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Collection, w: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Database, R: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Database, W: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Database, r: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Database, w: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Global, R: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Global, W: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Global, r: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Global, w: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Metadata, R: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Metadata, W: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Metadata, r: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Metadata, w: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Mutex, R: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Mutex, W: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Mutex, r: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Mutex, w: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, ParallelBatchWriterMode, R: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, ParallelBatchWriterMode, W: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, ParallelBatchWriterMode, r: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, ParallelBatchWriterMode, w: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, ReplicationStateTransition, R: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, ReplicationStateTransition, W: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, ReplicationStateTransition, r: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, ReplicationStateTransition, w: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, oplog, R: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, oplog, W: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, oplog, r: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, oplog, w: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Collection, R: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Collection, W: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Collection, r: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Collection, w: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Database, R: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Database, W: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Database, r: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Database, w: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Global, R: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Global, W: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Global, r: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Global, w: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Metadata, R: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Metadata, W: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Metadata, r: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Metadata, w: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Mutex, R: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Mutex, W: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Mutex, r: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Mutex, w: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, ParallelBatchWriterMode, R: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, ParallelBatchWriterMode, W: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, ParallelBatchWriterMode, r: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, ParallelBatchWriterMode, w: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, ReplicationStateTransition, R: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, ReplicationStateTransition, W: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, ReplicationStateTransition, r: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, ReplicationStateTransition, w: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, oplog, R: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, oplog, W: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, oplog, r: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, oplog, w: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Collection, R: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Collection, W: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Collection, r: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Collection, w: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Database, R: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Database, W: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Database, r: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Database, w: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Global, R: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Global, W: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Global, r: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Global, w: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Metadata, R: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Metadata, W: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Metadata, r: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Metadata, w: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Mutex, R: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Mutex, W: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Mutex, r: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Mutex, w: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, ParallelBatchWriterMode, R: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, ParallelBatchWriterMode, W: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, ParallelBatchWriterMode, r: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, ParallelBatchWriterMode, w: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, ReplicationStateTransition, R: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, ReplicationStateTransition, W: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, ReplicationStateTransition, r: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, ReplicationStateTransition, w: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, oplog, R: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, oplog, W: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, oplog, r: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, oplog, w: could not find key for metric",
			}, "; "))
	errAllClientFailedFetch = errors.New(
		strings.Join(
			[]string{
				"failed to fetch admin server status metrics: some admin server status error",
				"failed to fetch top stats metrics: some top stats error",
				"failed to fetch database stats metrics: some database stats error",
				"failed to fetch server status metrics: some server status error",
				"failed to fetch index stats metrics: some index stats error",
				"failed to fetch index stats metrics: some index stats error",
			}, "; "))

	errCollectionNames = errors.New(
		strings.Join(
			[]string{
				"failed to fetch admin server status metrics: some admin server status error",
				"failed to fetch top stats metrics: some top stats error",
				"failed to fetch database stats metrics: some database stats error",
				"failed to fetch server status metrics: some server status error",
				"failed to fetch collection names: some collection names error",
			}, "; "))
	errLocksPartialMetrics = errors.New(
		strings.Join(
			[]string{
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Collection, R: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Collection, W: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Collection, w: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Database, w: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Global, R: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Metadata, R: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Metadata, W: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Metadata, r: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Metadata, w: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Mutex, R: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Mutex, W: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Mutex, r: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, Mutex, w: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, ParallelBatchWriterMode, R: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, ParallelBatchWriterMode, W: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, ParallelBatchWriterMode, r: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, ParallelBatchWriterMode, w: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, ReplicationStateTransition, R: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, ReplicationStateTransition, W: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, ReplicationStateTransition, r: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, ReplicationStateTransition, w: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, oplog, R: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, oplog, W: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_count with attribute(s) fakedatabase, oplog, w: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Collection, R: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Collection, W: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Collection, r: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Collection, w: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Database, R: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Database, w: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Global, R: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Global, W: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Global, r: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Global, w: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Metadata, R: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Metadata, W: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Metadata, r: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Metadata, w: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Mutex, R: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Mutex, W: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Mutex, r: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, Mutex, w: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, ParallelBatchWriterMode, R: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, ParallelBatchWriterMode, W: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, ParallelBatchWriterMode, r: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, ParallelBatchWriterMode, w: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, ReplicationStateTransition, R: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, ReplicationStateTransition, W: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, ReplicationStateTransition, r: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, ReplicationStateTransition, w: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, oplog, R: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, oplog, W: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, oplog, r: could not find key for metric",
				"failed to collect metric mongodb.locks.acquire_wait_count with attribute(s) fakedatabase, oplog, w: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Collection, R: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Collection, W: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Collection, r: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Collection, w: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Database, R: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Database, W: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Database, r: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Database, w: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Global, R: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Global, W: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Global, r: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Global, w: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Metadata, R: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Metadata, W: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Metadata, r: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Metadata, w: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Mutex, R: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Mutex, W: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Mutex, r: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, Mutex, w: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, ParallelBatchWriterMode, R: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, ParallelBatchWriterMode, W: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, ParallelBatchWriterMode, r: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, ParallelBatchWriterMode, w: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, ReplicationStateTransition, R: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, ReplicationStateTransition, W: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, ReplicationStateTransition, r: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, ReplicationStateTransition, w: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, oplog, R: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, oplog, W: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, oplog, r: could not find key for metric",
				"failed to collect metric mongodb.locks.deadlock_count with attribute(s) fakedatabase, oplog, w: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Collection, R: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Collection, W: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Collection, r: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Collection, w: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Database, R: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Database, w: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Global, R: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Global, W: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Global, r: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Global, w: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Metadata, R: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Metadata, W: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Metadata, r: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Metadata, w: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Mutex, R: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Mutex, W: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Mutex, r: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, Mutex, w: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, ParallelBatchWriterMode, R: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, ParallelBatchWriterMode, W: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, ParallelBatchWriterMode, r: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, ParallelBatchWriterMode, w: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, ReplicationStateTransition, R: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, ReplicationStateTransition, W: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, ReplicationStateTransition, r: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, ReplicationStateTransition, w: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, oplog, R: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, oplog, W: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, oplog, r: could not find key for metric",
				"failed to collect metric mongodb.locks.time_acquiring_micros with attribute(s) fakedatabase, oplog, w: could not find key for metric",
			}, "; "))
)

func TestScraperScrape(t *testing.T) {
	testCases := []struct {
		desc              string
		partialErr        bool
		setupMockClient   func(t *testing.T) client
		expectedMetricGen func(t *testing.T) pmetric.Metrics
		expectedErr       error
	}{
		{
			desc:       "Nil client",
			partialErr: false,
			setupMockClient: func(t *testing.T) client {
				return nil
			},
			expectedMetricGen: func(t *testing.T) pmetric.Metrics {
				return pmetric.NewMetrics()
			},
			expectedErr: errors.New("no client was initialized before calling scrape"),
		},
		{
			desc:       "Failed to get version",
			partialErr: false,
			setupMockClient: func(t *testing.T) client {
				fc := &fakeClient{}
				mongo40, err := version.NewVersion("4.0")
				require.NoError(t, err)
				fc.On("GetVersion", mock.Anything).Return(mongo40, errors.New("some version error"))
				return fc
			},
			expectedMetricGen: func(t *testing.T) pmetric.Metrics {
				return pmetric.NewMetrics()
			},
			expectedErr: errors.New("unable to determine version of mongo scraping against: some version error"),
		},
		{
			desc:       "Failed to fetch database names",
			partialErr: true,
			setupMockClient: func(t *testing.T) client {
				fc := &fakeClient{}
				mongo40, err := version.NewVersion("4.0")
				require.NoError(t, err)
				fc.On("GetVersion", mock.Anything).Return(mongo40, nil)
				fc.On("ListDatabaseNames", mock.Anything, mock.Anything, mock.Anything).Return([]string{}, errors.New("some database names error"))
				return fc
			},
			expectedMetricGen: func(t *testing.T) pmetric.Metrics {
				return pmetric.NewMetrics()
			},
			expectedErr: errors.New("failed to fetch database names: some database names error"),
		},
		{
			desc:       "Failed to fetch collection names",
			partialErr: true,
			setupMockClient: func(t *testing.T) client {
				fc := &fakeClient{}
				mongo40, err := version.NewVersion("4.0")
				require.NoError(t, err)
				require.NoError(t, err)
				fakeDatabaseName := "fakedatabase"
				fc.On("GetVersion", mock.Anything).Return(mongo40, nil)
				fc.On("ListDatabaseNames", mock.Anything, mock.Anything, mock.Anything).Return([]string{fakeDatabaseName}, nil)
				fc.On("ServerStatus", mock.Anything, fakeDatabaseName).Return(bson.M{}, errors.New("some server status error"))
				fc.On("ServerStatus", mock.Anything, "admin").Return(bson.M{}, errors.New("some admin server status error"))
				fc.On("DBStats", mock.Anything, fakeDatabaseName).Return(bson.M{}, errors.New("some database stats error"))
				fc.On("TopStats", mock.Anything).Return(bson.M{}, errors.New("some top stats error"))
				fc.On("ListCollectionNames", mock.Anything, fakeDatabaseName).Return([]string{}, errors.New("some collection names error"))
				return fc
			},
			expectedMetricGen: func(t *testing.T) pmetric.Metrics {
				goldenPath := filepath.Join("testdata", "scraper", "partial_scrape.json")
				expectedMetrics, err := golden.ReadMetrics(goldenPath)
				require.NoError(t, err)
				return expectedMetrics
			},
			expectedErr: errCollectionNames,
		},
		{
			desc:       "Failed to scrape client stats",
			partialErr: true,
			setupMockClient: func(t *testing.T) client {
				fc := &fakeClient{}
				mongo40, err := version.NewVersion("4.0")
				require.NoError(t, err)
				require.NoError(t, err)
				fakeDatabaseName := "fakedatabase"
				fc.On("GetVersion", mock.Anything).Return(mongo40, nil)
				fc.On("ListDatabaseNames", mock.Anything, mock.Anything, mock.Anything).Return([]string{fakeDatabaseName}, nil)
				fc.On("ServerStatus", mock.Anything, fakeDatabaseName).Return(bson.M{}, errors.New("some server status error"))
				fc.On("ServerStatus", mock.Anything, "admin").Return(bson.M{}, errors.New("some admin server status error"))
				fc.On("DBStats", mock.Anything, fakeDatabaseName).Return(bson.M{}, errors.New("some database stats error"))
				fc.On("TopStats", mock.Anything).Return(bson.M{}, errors.New("some top stats error"))
				fc.On("ListCollectionNames", mock.Anything, fakeDatabaseName).Return([]string{"products", "orders"}, nil)
				fc.On("IndexStats", mock.Anything, fakeDatabaseName, "products").Return([]bson.M{}, errors.New("some index stats error"))
				fc.On("IndexStats", mock.Anything, fakeDatabaseName, "orders").Return([]bson.M{}, errors.New("some index stats error"))
				return fc
			},
			expectedMetricGen: func(t *testing.T) pmetric.Metrics {
				goldenPath := filepath.Join("testdata", "scraper", "partial_scrape.json")
				expectedMetrics, err := golden.ReadMetrics(goldenPath)
				require.NoError(t, err)
				return expectedMetrics
			},
			expectedErr: errAllClientFailedFetch,
		},
		{
			desc:       "Failed to scrape with partial errors on metrics",
			partialErr: true,
			setupMockClient: func(t *testing.T) client {
				fc := &fakeClient{}
				mongo40, err := version.NewVersion("4.0")
				require.NoError(t, err)
				wiredTigerStorage, err := loadOnlyStorageEngineAsMap()
				require.NoError(t, err)
				fakeDatabaseName := "fakedatabase"
				indexStats, err := loadIndexStatsAsMap("error")
				require.NoError(t, err)
				fc.On("GetVersion", mock.Anything).Return(mongo40, nil)
				fc.On("ListDatabaseNames", mock.Anything, mock.Anything, mock.Anything).Return([]string{fakeDatabaseName}, nil)
				fc.On("ServerStatus", mock.Anything, fakeDatabaseName).Return(bson.M{}, nil)
				fc.On("ServerStatus", mock.Anything, "admin").Return(wiredTigerStorage, nil)
				fc.On("DBStats", mock.Anything, fakeDatabaseName).Return(bson.M{}, nil)
				fc.On("TopStats", mock.Anything).Return(bson.M{}, nil)
				fc.On("ListCollectionNames", mock.Anything, fakeDatabaseName).Return([]string{"products", "orders"}, nil)
				fc.On("IndexStats", mock.Anything, fakeDatabaseName, "products").Return(indexStats, nil)
				fc.On("IndexStats", mock.Anything, fakeDatabaseName, "orders").Return(indexStats, nil)
				return fc
			},
			expectedMetricGen: func(t *testing.T) pmetric.Metrics {
				goldenPath := filepath.Join("testdata", "scraper", "partial_scrape.json")
				expectedMetrics, err := golden.ReadMetrics(goldenPath)
				require.NoError(t, err)
				return expectedMetrics
			},
			expectedErr: errAllPartialMetrics,
		},
		{
			desc:       "Successful scrape with partial errors on missing Lock Metrics",
			partialErr: true,
			setupMockClient: func(t *testing.T) client {
				fc := &fakeClient{}
				adminStatus, err := loadAdminStatusAsMap()
				require.NoError(t, err)
				ss, err := loadServerStatusAsMap()
				require.NoError(t, err)
				dbStats, err := loadDBStatsAsMap()
				require.NoError(t, err)
				topStats, err := loadTopAsMap()
				require.NoError(t, err)
				productsIndexStats, err := loadIndexStatsAsMap("products")
				require.NoError(t, err)
				ordersIndexStats, err := loadIndexStatsAsMap("orders")
				require.NoError(t, err)
				mongo40, err := version.NewVersion("4.0")
				require.NoError(t, err)
				fakeDatabaseName := "fakedatabase"
				fc.On("GetVersion", mock.Anything).Return(mongo40, nil)
				fc.On("ListDatabaseNames", mock.Anything, mock.Anything, mock.Anything).Return([]string{fakeDatabaseName}, nil)
				fc.On("ServerStatus", mock.Anything, fakeDatabaseName).Return(ss, nil)
				fc.On("ServerStatus", mock.Anything, "admin").Return(adminStatus, nil)
				fc.On("DBStats", mock.Anything, fakeDatabaseName).Return(dbStats, nil)
				fc.On("TopStats", mock.Anything).Return(topStats, nil)
				fc.On("ListCollectionNames", mock.Anything, fakeDatabaseName).Return([]string{"products", "orders"}, nil)
				fc.On("IndexStats", mock.Anything, fakeDatabaseName, "products").Return(productsIndexStats, nil)
				fc.On("IndexStats", mock.Anything, fakeDatabaseName, "orders").Return(ordersIndexStats, nil)
				return fc
			},
			expectedMetricGen: func(t *testing.T) pmetric.Metrics {
				goldenPath := filepath.Join("testdata", "scraper", "expected.json")
				expectedMetrics, err := golden.ReadMetrics(goldenPath)
				require.NoError(t, err)
				return expectedMetrics
			},
			expectedErr: errLocksPartialMetrics,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			scraper := newMongodbScraper(componenttest.NewNopReceiverCreateSettings(), createDefaultConfig().(*Config))
			scraper.client = tc.setupMockClient(t)
			actualMetrics, err := scraper.scrape(context.Background())

			if tc.expectedErr == nil {
				require.NoError(t, err)
			} else {
				if strings.Contains(err.Error(), ";") {
					// metrics with attributes use a map and errors can be returned in random order so sorting is required.
					// The first error message would noe have a leading whitespace and hence split on "; "
					actualErrs := strings.Split(err.Error(), "; ")
					sort.Strings(actualErrs)
					// The first error message would noe have a leading whitespace and hence split on "; "
					expectedErrs := strings.Split(tc.expectedErr.Error(), "; ")
					sort.Strings(expectedErrs)
					require.Equal(t, actualErrs, expectedErrs)
				} else {
					require.EqualError(t, err, tc.expectedErr.Error())
				}
			}

			if tc.partialErr {
				require.True(t, scrapererror.IsPartialScrapeError(err))
			} else {
				require.False(t, scrapererror.IsPartialScrapeError(err))
			}
			expectedMetrics := tc.expectedMetricGen(t)

			err = scrapertest.CompareMetrics(expectedMetrics, actualMetrics)
			require.NoError(t, err)
		})
	}
}

func TestTopMetricsAggregation(t *testing.T) {
	mont := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mont.Close()

	loadedTop, err := loadTop()
	require.NoError(t, err)

	mont.Run("test top stats are aggregated correctly", func(mt *mtest.T) {
		mt.AddMockResponses(loadedTop)
		driver := mt.Client
		client := mongodbClient{
			Client: driver,
			logger: zap.NewNop(),
		}
		var doc bson.M
		doc, err = client.TopStats(context.Background())
		require.NoError(t, err)

		collectionPathNames, err := digForCollectionPathNames(doc)
		require.NoError(t, err)
		require.ElementsMatch(t, collectionPathNames,
			[]string{
				"config.transactions",
				"test.admin",
				"test.orders",
				"admin.system.roles",
				"local.system.replset",
				"test.products",
				"admin.system.users",
				"admin.system.version",
				"config.system.sessions",
				"local.oplog.rs",
				"local.startup_log",
			})

		actualOperationTimeValues, err := aggregateOperationTimeValues(doc, collectionPathNames, operationsMap)
		require.NoError(t, err)

		// values are taken from testdata/top.json
		expectedInsertValues := 0 + 0 + 0 + 0 + 0 + 11302 + 0 + 1163 + 0 + 0 + 0
		expectedQueryValues := 0 + 0 + 6072 + 0 + 0 + 0 + 44 + 0 + 0 + 0 + 2791
		expectedUpdateValues := 0 + 0 + 0 + 0 + 0 + 0 + 0 + 0 + 155 + 9962 + 0
		expectedRemoveValues := 0 + 0 + 0 + 0 + 0 + 0 + 0 + 0 + 0 + 3750 + 0
		expectedGetmoreValues := 0 + 0 + 0 + 0 + 0 + 0 + 0 + 0 + 0 + 0 + 0
		expectedCommandValues := 540 + 397 + 4009 + 0 + 0 + 23285 + 0 + 10993 + 0 + 10116 + 0
		require.EqualValues(t, expectedInsertValues, actualOperationTimeValues["insert"])
		require.EqualValues(t, expectedQueryValues, actualOperationTimeValues["queries"])
		require.EqualValues(t, expectedUpdateValues, actualOperationTimeValues["update"])
		require.EqualValues(t, expectedRemoveValues, actualOperationTimeValues["remove"])
		require.EqualValues(t, expectedGetmoreValues, actualOperationTimeValues["getmore"])
		require.EqualValues(t, expectedCommandValues, actualOperationTimeValues["commands"])
	})
}
