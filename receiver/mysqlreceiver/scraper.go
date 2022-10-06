// Copyright  OpenTelemetry Authors
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

package mysqlreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mysqlreceiver"

import (
	"context"
	"errors"
	"strconv"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/scrapererror"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mysqlreceiver/internal/metadata"
)

const (
	picosecondsInNanoseconds int64 = 1000
)

type mySQLScraper struct {
	sqlclient client
	logger    *zap.Logger
	config    *Config
	mb        *metadata.MetricsBuilder
}

func newMySQLScraper(
	settings component.ReceiverCreateSettings,
	config *Config,
) *mySQLScraper {
	return &mySQLScraper{
		logger: settings.Logger,
		config: config,
		mb:     metadata.NewMetricsBuilder(config.Metrics, settings.BuildInfo),
	}
}

// start starts the scraper by initializing the db client connection.
func (m *mySQLScraper) start(_ context.Context, host component.Host) error {
	sqlclient := newMySQLClient(m.config)

	err := sqlclient.Connect()
	if err != nil {
		return err
	}
	m.sqlclient = sqlclient

	return nil
}

// shutdown closes the db connection
func (m *mySQLScraper) shutdown(context.Context) error {
	if m.sqlclient == nil {
		return nil
	}
	return m.sqlclient.Close()
}

// scrape scrapes the mysql db metric stats, transforms them and labels them into a metric slices.
func (m *mySQLScraper) scrape(context.Context) (pmetric.Metrics, error) {
	if m.sqlclient == nil {
		return pmetric.Metrics{}, errors.New("failed to connect to http client")
	}

	now := pcommon.NewTimestampFromTime(time.Now())

	// collect innodb metrics.
	innodbStats, innoErr := m.sqlclient.getInnodbStats()
	if innoErr != nil {
		m.logger.Error("Failed to fetch InnoDB stats", zap.Error(innoErr))
	}

	errs := &scrapererror.ScrapeErrors{}
	for k, v := range innodbStats {
		if k != "buffer_pool_size" {
			continue
		}
		addPartialIfError(errs, m.mb.RecordMysqlBufferPoolLimitDataPoint(now, v))
	}

	// collect io_waits metrics.
	m.scrapeTableIoWaitsStats(now, errs)
	m.scrapeIndexIoWaitsStats(now, errs)

	// collect performance event statements metrics.
	m.scrapeStatementEventsStats(now, errs)

	// collect global status metrics.
	m.scrapeGlobalStats(now, errs)

	m.mb.EmitForResource(metadata.WithMysqlInstanceEndpoint(m.config.Endpoint))

	return m.mb.Emit(), errs.Combine()
}

func (m *mySQLScraper) scrapeGlobalStats(now pcommon.Timestamp, errs *scrapererror.ScrapeErrors) {
	globalStats, err := m.sqlclient.getGlobalStats()
	if err != nil {
		m.logger.Error("Failed to fetch global stats", zap.Error(err))
		errs.AddPartial(66, err)
		return
	}

	m.recordDataPages(now, globalStats, errs)
	m.recordDataUsage(now, globalStats, errs)

	for k, v := range globalStats {
		switch k {

		// buffer_pool.pages
		case "Innodb_buffer_pool_pages_data":
			addPartialIfError(errs, m.mb.RecordMysqlBufferPoolPagesDataPoint(now, v,
				metadata.AttributeBufferPoolPagesData))
		case "Innodb_buffer_pool_pages_free":
			addPartialIfError(errs, m.mb.RecordMysqlBufferPoolPagesDataPoint(now, v,
				metadata.AttributeBufferPoolPagesFree))
		case "Innodb_buffer_pool_pages_misc":
			addPartialIfError(errs, m.mb.RecordMysqlBufferPoolPagesDataPoint(now, v,
				metadata.AttributeBufferPoolPagesMisc))

		// buffer_pool.page_flushes
		case "Innodb_buffer_pool_pages_flushed":
			addPartialIfError(errs, m.mb.RecordMysqlBufferPoolPageFlushesDataPoint(now, v))

		// buffer_pool.operations
		case "Innodb_buffer_pool_read_ahead_rnd":
			addPartialIfError(errs, m.mb.RecordMysqlBufferPoolOperationsDataPoint(now, v,
				metadata.AttributeBufferPoolOperationsReadAheadRnd))
		case "Innodb_buffer_pool_read_ahead":
			addPartialIfError(errs, m.mb.RecordMysqlBufferPoolOperationsDataPoint(now, v,
				metadata.AttributeBufferPoolOperationsReadAhead))
		case "Innodb_buffer_pool_read_ahead_evicted":
			addPartialIfError(errs, m.mb.RecordMysqlBufferPoolOperationsDataPoint(now, v,
				metadata.AttributeBufferPoolOperationsReadAheadEvicted))
		case "Innodb_buffer_pool_read_requests":
			addPartialIfError(errs, m.mb.RecordMysqlBufferPoolOperationsDataPoint(now, v,
				metadata.AttributeBufferPoolOperationsReadRequests))
		case "Innodb_buffer_pool_reads":
			addPartialIfError(errs, m.mb.RecordMysqlBufferPoolOperationsDataPoint(now, v,
				metadata.AttributeBufferPoolOperationsReads))
		case "Innodb_buffer_pool_wait_free":
			addPartialIfError(errs, m.mb.RecordMysqlBufferPoolOperationsDataPoint(now, v,
				metadata.AttributeBufferPoolOperationsWaitFree))
		case "Innodb_buffer_pool_write_requests":
			addPartialIfError(errs, m.mb.RecordMysqlBufferPoolOperationsDataPoint(now, v,
				metadata.AttributeBufferPoolOperationsWriteRequests))

		// connection.errors
		case "Connection_errors_accept":
			addPartialIfError(errs, m.mb.RecordMysqlConnectionErrorsDataPoint(now, v,
				metadata.AttributeConnectionErrorAccept))
		case "Connection_errors_internal":
			addPartialIfError(errs, m.mb.RecordMysqlConnectionErrorsDataPoint(now, v,
				metadata.AttributeConnectionErrorInternal))
		case "Connection_errors_max_connections":
			addPartialIfError(errs, m.mb.RecordMysqlConnectionErrorsDataPoint(now, v,
				metadata.AttributeConnectionErrorMaxConnections))
		case "Connection_errors_peer_address":
			addPartialIfError(errs, m.mb.RecordMysqlConnectionErrorsDataPoint(now, v,
				metadata.AttributeConnectionErrorPeerAddress))
		case "Connection_errors_select":
			addPartialIfError(errs, m.mb.RecordMysqlConnectionErrorsDataPoint(now, v,
				metadata.AttributeConnectionErrorSelect))
		case "Connection_errors_tcpwrap":
			addPartialIfError(errs, m.mb.RecordMysqlConnectionErrorsDataPoint(now, v,
				metadata.AttributeConnectionErrorTcpwrap))

		// commands
		case "Com_stmt_execute":
			addPartialIfError(errs, m.mb.RecordMysqlCommandsDataPoint(now, v, metadata.AttributeCommandExecute))
		case "Com_stmt_close":
			addPartialIfError(errs, m.mb.RecordMysqlCommandsDataPoint(now, v, metadata.AttributeCommandClose))
		case "Com_stmt_fetch":
			addPartialIfError(errs, m.mb.RecordMysqlCommandsDataPoint(now, v, metadata.AttributeCommandFetch))
		case "Com_stmt_prepare":
			addPartialIfError(errs, m.mb.RecordMysqlCommandsDataPoint(now, v, metadata.AttributeCommandPrepare))
		case "Com_stmt_reset":
			addPartialIfError(errs, m.mb.RecordMysqlCommandsDataPoint(now, v, metadata.AttributeCommandReset))
		case "Com_stmt_send_long_data":
			addPartialIfError(errs, m.mb.RecordMysqlCommandsDataPoint(now, v, metadata.AttributeCommandSendLongData))

		// created tmps
		case "Created_tmp_disk_tables":
			addPartialIfError(errs, m.mb.RecordMysqlTmpResourcesDataPoint(now, v, metadata.AttributeTmpResourceDiskTables))
		case "Created_tmp_files":
			addPartialIfError(errs, m.mb.RecordMysqlTmpResourcesDataPoint(now, v, metadata.AttributeTmpResourceFiles))
		case "Created_tmp_tables":
			addPartialIfError(errs, m.mb.RecordMysqlTmpResourcesDataPoint(now, v, metadata.AttributeTmpResourceTables))

		// handlers
		case "Handler_commit":
			addPartialIfError(errs, m.mb.RecordMysqlHandlersDataPoint(now, v, metadata.AttributeHandlerCommit))
		case "Handler_delete":
			addPartialIfError(errs, m.mb.RecordMysqlHandlersDataPoint(now, v, metadata.AttributeHandlerDelete))
		case "Handler_discover":
			addPartialIfError(errs, m.mb.RecordMysqlHandlersDataPoint(now, v, metadata.AttributeHandlerDiscover))
		case "Handler_external_lock":
			addPartialIfError(errs, m.mb.RecordMysqlHandlersDataPoint(now, v, metadata.AttributeHandlerExternalLock))
		case "Handler_mrr_init":
			addPartialIfError(errs, m.mb.RecordMysqlHandlersDataPoint(now, v, metadata.AttributeHandlerMrrInit))
		case "Handler_prepare":
			addPartialIfError(errs, m.mb.RecordMysqlHandlersDataPoint(now, v, metadata.AttributeHandlerPrepare))
		case "Handler_read_first":
			addPartialIfError(errs, m.mb.RecordMysqlHandlersDataPoint(now, v, metadata.AttributeHandlerReadFirst))
		case "Handler_read_key":
			addPartialIfError(errs, m.mb.RecordMysqlHandlersDataPoint(now, v, metadata.AttributeHandlerReadKey))
		case "Handler_read_last":
			addPartialIfError(errs, m.mb.RecordMysqlHandlersDataPoint(now, v, metadata.AttributeHandlerReadLast))
		case "Handler_read_next":
			addPartialIfError(errs, m.mb.RecordMysqlHandlersDataPoint(now, v, metadata.AttributeHandlerReadNext))
		case "Handler_read_prev":
			addPartialIfError(errs, m.mb.RecordMysqlHandlersDataPoint(now, v, metadata.AttributeHandlerReadPrev))
		case "Handler_read_rnd":
			addPartialIfError(errs, m.mb.RecordMysqlHandlersDataPoint(now, v, metadata.AttributeHandlerReadRnd))
		case "Handler_read_rnd_next":
			addPartialIfError(errs, m.mb.RecordMysqlHandlersDataPoint(now, v, metadata.AttributeHandlerReadRndNext))
		case "Handler_rollback":
			addPartialIfError(errs, m.mb.RecordMysqlHandlersDataPoint(now, v, metadata.AttributeHandlerRollback))
		case "Handler_savepoint":
			addPartialIfError(errs, m.mb.RecordMysqlHandlersDataPoint(now, v, metadata.AttributeHandlerSavepoint))
		case "Handler_savepoint_rollback":
			addPartialIfError(errs, m.mb.RecordMysqlHandlersDataPoint(now, v, metadata.AttributeHandlerSavepointRollback))
		case "Handler_update":
			addPartialIfError(errs, m.mb.RecordMysqlHandlersDataPoint(now, v, metadata.AttributeHandlerUpdate))
		case "Handler_write":
			addPartialIfError(errs, m.mb.RecordMysqlHandlersDataPoint(now, v, metadata.AttributeHandlerWrite))

		// double_writes
		case "Innodb_dblwr_pages_written":
			addPartialIfError(errs, m.mb.RecordMysqlDoubleWritesDataPoint(now, v, metadata.AttributeDoubleWritesPagesWritten))
		case "Innodb_dblwr_writes":
			addPartialIfError(errs, m.mb.RecordMysqlDoubleWritesDataPoint(now, v, metadata.AttributeDoubleWritesWrites))

		// log_operations
		case "Innodb_log_waits":
			addPartialIfError(errs, m.mb.RecordMysqlLogOperationsDataPoint(now, v, metadata.AttributeLogOperationsWaits))
		case "Innodb_log_write_requests":
			addPartialIfError(errs, m.mb.RecordMysqlLogOperationsDataPoint(now, v, metadata.AttributeLogOperationsWriteRequests))
		case "Innodb_log_writes":
			addPartialIfError(errs, m.mb.RecordMysqlLogOperationsDataPoint(now, v, metadata.AttributeLogOperationsWrites))

		// operations
		case "Innodb_data_fsyncs":
			addPartialIfError(errs, m.mb.RecordMysqlOperationsDataPoint(now, v, metadata.AttributeOperationsFsyncs))
		case "Innodb_data_reads":
			addPartialIfError(errs, m.mb.RecordMysqlOperationsDataPoint(now, v, metadata.AttributeOperationsReads))
		case "Innodb_data_writes":
			addPartialIfError(errs, m.mb.RecordMysqlOperationsDataPoint(now, v, metadata.AttributeOperationsWrites))

		// page_operations
		case "Innodb_pages_created":
			addPartialIfError(errs, m.mb.RecordMysqlPageOperationsDataPoint(now, v, metadata.AttributePageOperationsCreated))
		case "Innodb_pages_read":
			addPartialIfError(errs, m.mb.RecordMysqlPageOperationsDataPoint(now, v,
				metadata.AttributePageOperationsRead))
		case "Innodb_pages_written":
			addPartialIfError(errs, m.mb.RecordMysqlPageOperationsDataPoint(now, v,
				metadata.AttributePageOperationsWritten))

		// row_locks
		case "Innodb_row_lock_waits":
			addPartialIfError(errs, m.mb.RecordMysqlRowLocksDataPoint(now, v, metadata.AttributeRowLocksWaits))
		case "Innodb_row_lock_time":
			addPartialIfError(errs, m.mb.RecordMysqlRowLocksDataPoint(now, v, metadata.AttributeRowLocksTime))

		// row_operations
		case "Innodb_rows_deleted":
			addPartialIfError(errs, m.mb.RecordMysqlRowOperationsDataPoint(now, v, metadata.AttributeRowOperationsDeleted))
		case "Innodb_rows_inserted":
			addPartialIfError(errs, m.mb.RecordMysqlRowOperationsDataPoint(now, v, metadata.AttributeRowOperationsInserted))
		case "Innodb_rows_read":
			addPartialIfError(errs, m.mb.RecordMysqlRowOperationsDataPoint(now, v,
				metadata.AttributeRowOperationsRead))
		case "Innodb_rows_updated":
			addPartialIfError(errs, m.mb.RecordMysqlRowOperationsDataPoint(now, v,
				metadata.AttributeRowOperationsUpdated))

		// locks
		case "Table_locks_immediate":
			addPartialIfError(errs, m.mb.RecordMysqlLocksDataPoint(now, v, metadata.AttributeLocksImmediate))
		case "Table_locks_waited":
			addPartialIfError(errs, m.mb.RecordMysqlLocksDataPoint(now, v, metadata.AttributeLocksWaited))

		// locked_connects
		case "Locked_connects":
			addPartialIfError(errs, m.mb.RecordMysqlLockedConnectsDataPoint(now, v))

		// joins
		case "Select_full_join":
			addPartialIfError(errs, m.mb.RecordMysqlJoinsDataPoint(now, v, metadata.AttributeJoinKindFull))
		case "Select_full_range_join":
			addPartialIfError(errs, m.mb.RecordMysqlJoinsDataPoint(now, v, metadata.AttributeJoinKindFullRange))
		case "Select_range":
			addPartialIfError(errs, m.mb.RecordMysqlJoinsDataPoint(now, v, metadata.AttributeJoinKindRange))
		case "Select_range_check":
			addPartialIfError(errs, m.mb.RecordMysqlJoinsDataPoint(now, v, metadata.AttributeJoinKindRangeCheck))
		case "Select_scan":
			addPartialIfError(errs, m.mb.RecordMysqlJoinsDataPoint(now, v, metadata.AttributeJoinKindScan))

		// open cache
		case "Table_open_cache_hits":
			addPartialIfError(errs, m.mb.RecordMysqlTableOpenCacheDataPoint(now, v, metadata.AttributeCacheStatusHit))
		case "Table_open_cache_misses":
			addPartialIfError(errs, m.mb.RecordMysqlTableOpenCacheDataPoint(now, v, metadata.AttributeCacheStatusMiss))
		case "Table_open_cache_overflows":
			addPartialIfError(errs, m.mb.RecordMysqlTableOpenCacheDataPoint(now, v, metadata.AttributeCacheStatusOverflow))

		// queries
		case "Queries":
			addPartialIfError(errs, m.mb.RecordMysqlQueryCountDataPoint(now, v))
		case "Questions":
			addPartialIfError(errs, m.mb.RecordMysqlQueryClientCountDataPoint(now, v))
		case "Slow_queries":
			addPartialIfError(errs, m.mb.RecordMysqlQuerySlowCountDataPoint(now, v))

		// sorts
		case "Sort_merge_passes":
			addPartialIfError(errs, m.mb.RecordMysqlSortsDataPoint(now, v, metadata.AttributeSortsMergePasses))
		case "Sort_range":
			addPartialIfError(errs, m.mb.RecordMysqlSortsDataPoint(now, v, metadata.AttributeSortsRange))
		case "Sort_rows":
			addPartialIfError(errs, m.mb.RecordMysqlSortsDataPoint(now, v, metadata.AttributeSortsRows))
		case "Sort_scan":
			addPartialIfError(errs, m.mb.RecordMysqlSortsDataPoint(now, v, metadata.AttributeSortsScan))

		// threads
		case "Threads_cached":
			addPartialIfError(errs, m.mb.RecordMysqlThreadsDataPoint(now, v, metadata.AttributeThreadsCached))
		case "Threads_connected":
			addPartialIfError(errs, m.mb.RecordMysqlThreadsDataPoint(now, v, metadata.AttributeThreadsConnected))
		case "Threads_created":
			addPartialIfError(errs, m.mb.RecordMysqlThreadsDataPoint(now, v, metadata.AttributeThreadsCreated))
		case "Threads_running":
			addPartialIfError(errs, m.mb.RecordMysqlThreadsDataPoint(now, v, metadata.AttributeThreadsRunning))

		// opened resources
		case "Opened_files":
			addPartialIfError(errs, m.mb.RecordMysqlOpenedResourcesDataPoint(now, v, metadata.AttributeOpenedResourcesFile))
		case "Opened_tables":
			addPartialIfError(errs, m.mb.RecordMysqlOpenedResourcesDataPoint(now, v, metadata.AttributeOpenedResourcesTable))
		case "Opened_table_definitions":
			addPartialIfError(errs, m.mb.RecordMysqlOpenedResourcesDataPoint(now, v, metadata.AttributeOpenedResourcesTableDefinition))

		// mysqlx_worker_threads
		case "Mysqlx_worker_threads":
			addPartialIfError(errs, m.mb.RecordMysqlMysqlxWorkerThreadsDataPoint(now, v, metadata.AttributeMysqlxThreadsAvailable))
		case "Mysqlx_worker_threads_active":
			addPartialIfError(errs, m.mb.RecordMysqlMysqlxWorkerThreadsDataPoint(now, v, metadata.AttributeMysqlxThreadsActive))

		// mysqlx_connections
		case "Mysqlx_connections_accepted":
			addPartialIfError(errs, m.mb.RecordMysqlMysqlxConnectionsDataPoint(now, v, metadata.AttributeConnectionStatusAccepted))
		case "Mysqlx_connections_closed":
			addPartialIfError(errs, m.mb.RecordMysqlMysqlxConnectionsDataPoint(now, v, metadata.AttributeConnectionStatusClosed))
		case "Mysqlx_connections_rejected":
			addPartialIfError(errs, m.mb.RecordMysqlMysqlxConnectionsDataPoint(now, v, metadata.AttributeConnectionStatusRejected))
		}
	}
}

func (m *mySQLScraper) scrapeTableIoWaitsStats(now pcommon.Timestamp, errs *scrapererror.ScrapeErrors) {
	tableIoWaitsStats, err := m.sqlclient.getTableIoWaitsStats()
	if err != nil {
		m.logger.Error("Failed to fetch table io_waits stats", zap.Error(err))
		errs.AddPartial(8, err)
		return
	}

	for i := 0; i < len(tableIoWaitsStats); i++ {
		s := tableIoWaitsStats[i]
		// counts
		m.mb.RecordMysqlTableIoWaitCountDataPoint(now, s.countDelete, metadata.AttributeIoWaitsOperationsDelete, s.name, s.schema)
		m.mb.RecordMysqlTableIoWaitCountDataPoint(now, s.countFetch, metadata.AttributeIoWaitsOperationsFetch, s.name, s.schema)
		m.mb.RecordMysqlTableIoWaitCountDataPoint(now, s.countInsert, metadata.AttributeIoWaitsOperationsInsert, s.name, s.schema)
		m.mb.RecordMysqlTableIoWaitCountDataPoint(now, s.countUpdate, metadata.AttributeIoWaitsOperationsUpdate, s.name, s.schema)

		// times
		m.mb.RecordMysqlTableIoWaitTimeDataPoint(
			now, s.timeDelete/picosecondsInNanoseconds, metadata.AttributeIoWaitsOperationsDelete, s.name, s.schema,
		)
		m.mb.RecordMysqlTableIoWaitTimeDataPoint(
			now, s.timeFetch/picosecondsInNanoseconds, metadata.AttributeIoWaitsOperationsFetch, s.name, s.schema,
		)
		m.mb.RecordMysqlTableIoWaitTimeDataPoint(
			now, s.timeInsert/picosecondsInNanoseconds, metadata.AttributeIoWaitsOperationsInsert, s.name, s.schema,
		)
		m.mb.RecordMysqlTableIoWaitTimeDataPoint(
			now, s.timeUpdate/picosecondsInNanoseconds, metadata.AttributeIoWaitsOperationsUpdate, s.name, s.schema,
		)
	}
}

func (m *mySQLScraper) scrapeIndexIoWaitsStats(now pcommon.Timestamp, errs *scrapererror.ScrapeErrors) {
	indexIoWaitsStats, err := m.sqlclient.getIndexIoWaitsStats()
	if err != nil {
		m.logger.Error("Failed to fetch index io_waits stats", zap.Error(err))
		errs.AddPartial(8, err)
		return
	}

	for i := 0; i < len(indexIoWaitsStats); i++ {
		s := indexIoWaitsStats[i]
		// counts
		m.mb.RecordMysqlIndexIoWaitCountDataPoint(now, s.countDelete, metadata.AttributeIoWaitsOperationsDelete, s.name, s.schema, s.index)
		m.mb.RecordMysqlIndexIoWaitCountDataPoint(now, s.countFetch, metadata.AttributeIoWaitsOperationsFetch, s.name, s.schema, s.index)
		m.mb.RecordMysqlIndexIoWaitCountDataPoint(now, s.countInsert, metadata.AttributeIoWaitsOperationsInsert, s.name, s.schema, s.index)
		m.mb.RecordMysqlIndexIoWaitCountDataPoint(now, s.countUpdate, metadata.AttributeIoWaitsOperationsUpdate, s.name, s.schema, s.index)

		// times
		m.mb.RecordMysqlIndexIoWaitTimeDataPoint(
			now, s.timeDelete/picosecondsInNanoseconds, metadata.AttributeIoWaitsOperationsDelete, s.name, s.schema, s.index,
		)
		m.mb.RecordMysqlIndexIoWaitTimeDataPoint(
			now, s.timeFetch/picosecondsInNanoseconds, metadata.AttributeIoWaitsOperationsFetch, s.name, s.schema, s.index,
		)
		m.mb.RecordMysqlIndexIoWaitTimeDataPoint(
			now, s.timeInsert/picosecondsInNanoseconds, metadata.AttributeIoWaitsOperationsInsert, s.name, s.schema, s.index,
		)
		m.mb.RecordMysqlIndexIoWaitTimeDataPoint(
			now, s.timeUpdate/picosecondsInNanoseconds, metadata.AttributeIoWaitsOperationsUpdate, s.name, s.schema, s.index,
		)
	}
}

func (m *mySQLScraper) scrapeStatementEventsStats(now pcommon.Timestamp, errs *scrapererror.ScrapeErrors) {
	statementEventsStats, err := m.sqlclient.getStatementEventsStats()
	if err != nil {
		m.logger.Error("Failed to fetch index io_waits stats", zap.Error(err))
		errs.AddPartial(8, err)
		return
	}

	for i := 0; i < len(statementEventsStats); i++ {
		s := statementEventsStats[i]
		m.mb.RecordMysqlStatementEventCountDataPoint(now, s.countCreatedTmpDiskTables, s.schema, s.digest, s.digestText, metadata.AttributeEventStateCreatedTmpDiskTables)
		m.mb.RecordMysqlStatementEventCountDataPoint(now, s.countCreatedTmpTables, s.schema, s.digest, s.digestText, metadata.AttributeEventStateCreatedTmpTables)
		m.mb.RecordMysqlStatementEventCountDataPoint(now, s.countErrors, s.schema, s.digest, s.digestText, metadata.AttributeEventStateErrors)
		m.mb.RecordMysqlStatementEventCountDataPoint(now, s.countNoIndexUsed, s.schema, s.digest, s.digestText, metadata.AttributeEventStateNoIndexUsed)
		m.mb.RecordMysqlStatementEventCountDataPoint(now, s.countRowsAffected, s.schema, s.digest, s.digestText, metadata.AttributeEventStateRowsAffected)
		m.mb.RecordMysqlStatementEventCountDataPoint(now, s.countRowsExamined, s.schema, s.digest, s.digestText, metadata.AttributeEventStateRowsExamined)
		m.mb.RecordMysqlStatementEventCountDataPoint(now, s.countRowsSent, s.schema, s.digest, s.digestText, metadata.AttributeEventStateRowsSent)
		m.mb.RecordMysqlStatementEventCountDataPoint(now, s.countSortMergePasses, s.schema, s.digest, s.digestText, metadata.AttributeEventStateSortMergePasses)
		m.mb.RecordMysqlStatementEventCountDataPoint(now, s.countSortRows, s.schema, s.digest, s.digestText, metadata.AttributeEventStateSortRows)
		m.mb.RecordMysqlStatementEventCountDataPoint(now, s.countWarnings, s.schema, s.digest, s.digestText, metadata.AttributeEventStateWarnings)

		m.mb.RecordMysqlStatementEventWaitTimeDataPoint(now, s.sumTimerWait/picosecondsInNanoseconds, s.schema, s.digest, s.digestText)
	}
}

func addPartialIfError(errors *scrapererror.ScrapeErrors, err error) {
	if err != nil {
		errors.AddPartial(1, err)
	}
}

func (m *mySQLScraper) recordDataPages(now pcommon.Timestamp, globalStats map[string]string, errors *scrapererror.ScrapeErrors) {
	dirty, err := parseInt(globalStats["Innodb_buffer_pool_pages_dirty"])
	if err != nil {
		errors.AddPartial(2, err) // we need dirty to calculate free, so 2 data points lost here
		return
	}
	m.mb.RecordMysqlBufferPoolDataPagesDataPoint(now, dirty, metadata.AttributeBufferPoolDataDirty)

	data, err := parseInt(globalStats["Innodb_buffer_pool_pages_data"])
	if err != nil {
		errors.AddPartial(1, err)
		return
	}
	m.mb.RecordMysqlBufferPoolDataPagesDataPoint(now, data-dirty, metadata.AttributeBufferPoolDataClean)
}

func (m *mySQLScraper) recordDataUsage(now pcommon.Timestamp, globalStats map[string]string, errors *scrapererror.ScrapeErrors) {
	dirty, err := parseInt(globalStats["Innodb_buffer_pool_bytes_dirty"])
	if err != nil {
		errors.AddPartial(2, err) // we need dirty to calculate free, so 2 data points lost here
		return
	}
	m.mb.RecordMysqlBufferPoolUsageDataPoint(now, dirty, metadata.AttributeBufferPoolDataDirty)

	data, err := parseInt(globalStats["Innodb_buffer_pool_bytes_data"])
	if err != nil {
		errors.AddPartial(1, err)
		return
	}
	m.mb.RecordMysqlBufferPoolUsageDataPoint(now, data-dirty, metadata.AttributeBufferPoolDataClean)
}

// parseInt converts string to int64.
func parseInt(value string) (int64, error) {
	return strconv.ParseInt(value, 10, 64)
}
