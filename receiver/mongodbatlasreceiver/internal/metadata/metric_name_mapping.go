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

package metadata

import (
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/atlas/mongodbatlas"
	"go.opentelemetry.io/collector/model/pdata"
)

type metricMappingData struct {
	metricName string
	attributes map[string]pdata.AttributeValue
}

var metricNameMapping = map[string]metricMappingData{
	// MongoDB CPU usage. For hosts with more than one CPU core, these values can exceed 100%.

	"PROCESS_CPU_USER": {"mongodbatlas.process.cpu.usage", map[string]pdata.AttributeValue{
		"state":       pdata.NewAttributeValueString("user"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},

	"MAX_PROCESS_CPU_USER": {"mongodbatlas.process.cpu.usage", map[string]pdata.AttributeValue{
		"state":       pdata.NewAttributeValueString("user"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},

	"PROCESS_CPU_KERNEL": {"mongodbatlas.process.cpu.usage", map[string]pdata.AttributeValue{
		"state":       pdata.NewAttributeValueString("kernel"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},

	"MAX_PROCESS_CPU_KERNEL": {"mongodbatlas.process.cpu.usage", map[string]pdata.AttributeValue{
		"state":       pdata.NewAttributeValueString("kernel"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},

	"PROCESS_CPU_CHILDREN_USER": {"mongodbatlas.process.cpu.children.usage", map[string]pdata.AttributeValue{
		"state":       pdata.NewAttributeValueString("user"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},

	"MAX_PROCESS_CPU_CHILDREN_USER": {"mongodbatlas.process.cpu.children.usage", map[string]pdata.AttributeValue{
		"state":       pdata.NewAttributeValueString("user"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},

	"PROCESS_CPU_CHILDREN_KERNEL": {"mongodbatlas.process.cpu.children.usage", map[string]pdata.AttributeValue{
		"state":       pdata.NewAttributeValueString("kernel"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},

	"MAX_PROCESS_CPU_CHILDREN_KERNEL": {"mongodbatlas.process.cpu.children.usage", map[string]pdata.AttributeValue{
		"state":       pdata.NewAttributeValueString("kernel"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},

	// MongoDB CPU usage scaled to a range of 0% to 100%. Atlas computes this value by dividing by the number of CPU cores.

	"PROCESS_NORMALIZED_CPU_USER": {"mongodbatlas.process.cpu.normalized.usage", map[string]pdata.AttributeValue{
		"state":       pdata.NewAttributeValueString("user"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},

	"MAX_PROCESS_NORMALIZED_CPU_USER": {"mongodbatlas.process.cpu.normalized.usage", map[string]pdata.AttributeValue{
		"state":       pdata.NewAttributeValueString("user"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},

	"PROCESS_NORMALIZED_CPU_KERNEL": {"mongodbatlas.process.cpu.normalized.usage", map[string]pdata.AttributeValue{
		"state":       pdata.NewAttributeValueString("kernel"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},

	"MAX_PROCESS_NORMALIZED_CPU_KERNEL": {"mongodbatlas.process.cpu.normalized.usage", map[string]pdata.AttributeValue{
		"state":       pdata.NewAttributeValueString("kernel"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},

	"PROCESS_NORMALIZED_CPU_CHILDREN_USER": {"mongodbatlas.process.cpu.children.normalized.usage", map[string]pdata.AttributeValue{
		"state":       pdata.NewAttributeValueString("user"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},

	// Context: Process
	"MAX_PROCESS_NORMALIZED_CPU_CHILDREN_USER": {"mongodbatlas.process.cpu.children.normalized.usage", map[string]pdata.AttributeValue{
		"state":       pdata.NewAttributeValueString("user"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},

	"PROCESS_NORMALIZED_CPU_CHILDREN_KERNEL": {"mongodbatlas.process.cpu.children.normalized.usage", map[string]pdata.AttributeValue{
		"state":       pdata.NewAttributeValueString("kernel"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},

	"MAX_PROCESS_NORMALIZED_CPU_CHILDREN_KERNEL": {"mongodbatlas.process.cpu.children.normalized.usage", map[string]pdata.AttributeValue{
		"state":       pdata.NewAttributeValueString("kernel"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},

	// Rate of asserts for a MongoDB process found in the asserts document that the serverStatus command generates.

	"ASSERT_REGULAR": {"mongodbatlas.process.asserts", map[string]pdata.AttributeValue{
		"type": pdata.NewAttributeValueString("regular"),
	}},

	"ASSERT_WARNING": {"mongodbatlas.process.asserts", map[string]pdata.AttributeValue{
		"type": pdata.NewAttributeValueString("warning"),
	}},

	"ASSERT_MSG": {"mongodbatlas.process.asserts", map[string]pdata.AttributeValue{
		"type": pdata.NewAttributeValueString("msg"),
	}},

	"ASSERT_USER": {"mongodbatlas.process.asserts", map[string]pdata.AttributeValue{
		"type": pdata.NewAttributeValueString("user"),
	}},

	// Amount of data flushed in the background.

	"BACKGROUND_FLUSH_AVG": {"mongodbatlas.process.background_flush", map[string]pdata.AttributeValue{}},

	// Amount of bytes in the WiredTiger storage engine cache and tickets found in the wiredTiger.cache and wiredTiger.concurrentTransactions documents that the serverStatus command generates.

	"CACHE_BYTES_READ_INTO": {"mongodbatlas.process.cache.io", map[string]pdata.AttributeValue{
		"direction": pdata.NewAttributeValueString("read_into"),
	}},

	"CACHE_BYTES_WRITTEN_FROM": {"mongodbatlas.process.cache.io", map[string]pdata.AttributeValue{
		"direction": pdata.NewAttributeValueString("written_from"),
	}},

	"CACHE_DIRTY_BYTES": {"mongodbatlas.process.cache.size", map[string]pdata.AttributeValue{
		"status": pdata.NewAttributeValueString("dirty"),
	}},

	"CACHE_USED_BYTES": {"mongodbatlas.process.cache.size", map[string]pdata.AttributeValue{
		"status": pdata.NewAttributeValueString("used"),
	}},

	"TICKETS_AVAILABLE_READS": {"mongodbatlas.process.tickets", map[string]pdata.AttributeValue{
		"type": pdata.NewAttributeValueString("available_reads"),
	}},

	"TICKETS_AVAILABLE_WRITE": {"mongodbatlas.process.tickets", map[string]pdata.AttributeValue{
		"type": pdata.NewAttributeValueString("available_writes"),
	}},

	// Number of connections to a MongoDB process found in the connections document that the serverStatus command generates.
	"CONNECTIONS": {"mongodbatlas.process.connections", map[string]pdata.AttributeValue{}},

	// Number of cursors for a MongoDB process found in the metrics.cursor document that the serverStatus command generates.
	"CURSORS_TOTAL_OPEN": {"mongodbatlas.process.cursors", map[string]pdata.AttributeValue{
		"state": pdata.NewAttributeValueString("open"),
	}},

	"CURSORS_TOTAL_TIMED_OUT": {"mongodbatlas.process.cursors", map[string]pdata.AttributeValue{
		"state": pdata.NewAttributeValueString("timed_out"),
	}},

	// Numbers of Memory Issues and Page Faults for a MongoDB process.
	"EXTRA_INFO_PAGE_FAULTS": {"mongodbatlas.process.page_faults", map[string]pdata.AttributeValue{
		"type": pdata.NewAttributeValueString("extra_info"),
	}},

	"GLOBAL_ACCESSES_NOT_IN_MEMORY": {"mongodbatlas.process.page_faults", map[string]pdata.AttributeValue{
		"type": pdata.NewAttributeValueString("global_accesses_not_in_memory"),
	}},
	"GLOBAL_PAGE_FAULT_EXCEPTIONS_THROWN": {"mongodbatlas.process.page_faults", map[string]pdata.AttributeValue{
		"type": pdata.NewAttributeValueString("exceptions_thrown"),
	}},

	// Number of operations waiting on locks for the MongoDB process that the serverStatus command generates. Cloud Manager computes these values based on the type of storage engine.
	"GLOBAL_LOCK_CURRENT_QUEUE_TOTAL": {"mongodbatlas.process.global_lock", map[string]pdata.AttributeValue{
		"state": pdata.NewAttributeValueString("current_queue_total"),
	}},
	"GLOBAL_LOCK_CURRENT_QUEUE_READERS": {"mongodbatlas.process.global_lock", map[string]pdata.AttributeValue{
		"state": pdata.NewAttributeValueString("current_queue_readers"),
	}},
	"GLOBAL_LOCK_CURRENT_QUEUE_WRITERS": {"mongodbatlas.process.global_lock", map[string]pdata.AttributeValue{
		"state": pdata.NewAttributeValueString("current_queue_writers"),
	}},

	// Number of index btree operations.
	"INDEX_COUNTERS_BTREE_ACCESSES": {"mongodbatlas.process.index.counters", map[string]pdata.AttributeValue{
		"type": pdata.NewAttributeValueString("btree_accesses"),
	}},
	"INDEX_COUNTERS_BTREE_HITS": {"mongodbatlas.process.index.counters", map[string]pdata.AttributeValue{
		"type": pdata.NewAttributeValueString("btree_hits"),
	}},
	"INDEX_COUNTERS_BTREE_MISSES": {"mongodbatlas.process.index.counters", map[string]pdata.AttributeValue{
		"type": pdata.NewAttributeValueString("btree_misses"),
	}},
	"INDEX_COUNTERS_BTREE_MISS_RATIO": {"mongodbatlas.process.index.btree_miss_ratio", map[string]pdata.AttributeValue{}},

	// Number of journaling operations.
	"JOURNALING_COMMITS_IN_WRITE_LOCK": {"mongodbatlas.process.journaling.commits", map[string]pdata.AttributeValue{
		"status": pdata.NewAttributeValueString("in_write_lock"),
	}},
	"JOURNALING_MB":                  {"mongodbatlas.process.journaling.written", map[string]pdata.AttributeValue{}},
	"JOURNALING_WRITE_DATA_FILES_MB": {"mongodbatlas.process.journaling.data_files", map[string]pdata.AttributeValue{}},

	// Amount of memory for a MongoDB process found in the mem document that the serverStatus command collects.
	"MEMORY_RESIDENT": {"mongodbatlas.process.memory.usage", map[string]pdata.AttributeValue{
		"state": pdata.NewAttributeValueString("resident"),
	}},
	"MEMORY_VIRTUAL": {"mongodbatlas.process.memory.usage", map[string]pdata.AttributeValue{
		"state": pdata.NewAttributeValueString("virtual"),
	}},

	"MEMORY_MAPPED": {"mongodbatlas.process.memory.usage", map[string]pdata.AttributeValue{
		"state": pdata.NewAttributeValueString("mapped"),
	}},
	"COMPUTED_MEMORY": {"mongodbatlas.process.memory.usage", map[string]pdata.AttributeValue{
		"state": pdata.NewAttributeValueString("computed"),
	}},

	// Amount of throughput for MongoDB process found in the network document that the serverStatus command collects.

	"NETWORK_BYTES_IN": {"mongodbatlas.process.network.io", map[string]pdata.AttributeValue{
		"direction": pdata.NewAttributeValueString("receive"),
	}},
	"NETWORK_BYTES_OUT": {"mongodbatlas.process.network.io", map[string]pdata.AttributeValue{
		"direction": pdata.NewAttributeValueString("transmit"),
	}},
	"NETWORK_NUM_REQUESTS": {"mongodbatlas.process.network.requests", map[string]pdata.AttributeValue{}},

	// Durations and throughput of the MongoDB process' oplog.
	"OPLOG_SLAVE_LAG_MASTER_TIME": {"mongodbatlas.process.oplog.time", map[string]pdata.AttributeValue{
		"type": pdata.NewAttributeValueString("slave_lag_master_time"),
	}},
	"OPLOG_MASTER_TIME": {"mongodbatlas.process.oplog.time", map[string]pdata.AttributeValue{
		"type": pdata.NewAttributeValueString("master_time"),
	}},
	"OPLOG_MASTER_LAG_TIME_DIFF": {"mongodbatlas.process.oplog.time", map[string]pdata.AttributeValue{
		"type": pdata.NewAttributeValueString("master_lag_time_diff"),
	}},
	"OPLOG_RATE_GB_PER_HOUR": {"mongodbatlas.process.oplog.rate", map[string]pdata.AttributeValue{
		"aggregation": pdata.NewAttributeValueString("hour"),
	}},

	// Number of database operations on a MongoDB process since the process last started.

	"DB_STORAGE_TOTAL": {"mongodbatlas.process.db.storage", map[string]pdata.AttributeValue{
		"status": pdata.NewAttributeValueString("total"),
	}},

	"DB_DATA_SIZE_TOTAL": {"mongodbatlas.process.db.storage", map[string]pdata.AttributeValue{
		"status": pdata.NewAttributeValueString("data_size_total"),
	}},
	"DB_INDEX_SIZE_TOTAL": {"mongodbatlas.process.db.storage", map[string]pdata.AttributeValue{
		"status": pdata.NewAttributeValueString("index_size_total"),
	}},
	"DB_DATA_SIZE_TOTAL_WO_SYSTEM": {"mongodbatlas.process.db.storage", map[string]pdata.AttributeValue{
		"status": pdata.NewAttributeValueString("data_size_total_wo_system"),
	}},

	// Rate of database operations on a MongoDB process since the process last started found in the opcounters document that the serverStatus command collects.
	"OPCOUNTER_CMD": {"mongodbatlas.process.db.operations.rate", map[string]pdata.AttributeValue{
		"operation": pdata.NewAttributeValueString("cmd"),
		"role":      pdata.NewAttributeValueString("primary"),
	}},
	"OPCOUNTER_QUERY": {"mongodbatlas.process.db.operations.rate", map[string]pdata.AttributeValue{
		"operation": pdata.NewAttributeValueString("query"),
		"role":      pdata.NewAttributeValueString("primary"),
	}},
	"OPCOUNTER_UPDATE": {"mongodbatlas.process.db.operations.rate", map[string]pdata.AttributeValue{
		"operation": pdata.NewAttributeValueString("update"),
		"role":      pdata.NewAttributeValueString("primary"),
	}},
	"OPCOUNTER_DELETE": {"mongodbatlas.process.db.operations.rate", map[string]pdata.AttributeValue{
		"operation": pdata.NewAttributeValueString("delete"),
		"role":      pdata.NewAttributeValueString("primary"),
	}},
	"OPCOUNTER_GETMORE": {"mongodbatlas.process.db.operations.rate", map[string]pdata.AttributeValue{
		"operation": pdata.NewAttributeValueString("getmore"),
		"role":      pdata.NewAttributeValueString("primary"),
	}},
	"OPCOUNTER_INSERT": {"mongodbatlas.process.db.operations.rate", map[string]pdata.AttributeValue{
		"operation": pdata.NewAttributeValueString("insert"),
		"role":      pdata.NewAttributeValueString("primary"),
	}},

	// Rate of database operations on MongoDB secondaries found in the opcountersRepl document that the serverStatus command collects.
	"OPCOUNTER_REPL_CMD": {"mongodbatlas.process.db.operations.rate", map[string]pdata.AttributeValue{
		"operation": pdata.NewAttributeValueString("cmd"),
		"role":      pdata.NewAttributeValueString("replica"),
	}},
	"OPCOUNTER_REPL_UPDATE": {"mongodbatlas.process.db.operations.rate", map[string]pdata.AttributeValue{
		"operation": pdata.NewAttributeValueString("update"),
		"role":      pdata.NewAttributeValueString("replica"),
	}},
	"OPCOUNTER_REPL_DELETE": {"mongodbatlas.process.db.operations.rate", map[string]pdata.AttributeValue{
		"operation": pdata.NewAttributeValueString("delete"),
		"role":      pdata.NewAttributeValueString("replica"),
	}},
	"OPCOUNTER_REPL_INSERT": {"mongodbatlas.process.db.operations.rate", map[string]pdata.AttributeValue{
		"operation": pdata.NewAttributeValueString("insert"),
		"role":      pdata.NewAttributeValueString("replica"),
	}},

	// Average rate of documents returned, inserted, updated, or deleted per second during a selected time period.
	"DOCUMENT_METRICS_RETURNED": {"mongodbatlas.process.db.document.rate", map[string]pdata.AttributeValue{
		"status": pdata.NewAttributeValueString("returned"),
	}},
	"DOCUMENT_METRICS_INSERTED": {"mongodbatlas.process.db.document.rate", map[string]pdata.AttributeValue{
		"status": pdata.NewAttributeValueString("inserted"),
	}},
	"DOCUMENT_METRICS_UPDATED": {"mongodbatlas.process.db.document.rate", map[string]pdata.AttributeValue{
		"status": pdata.NewAttributeValueString("updated"),
	}},
	"DOCUMENT_METRICS_DELETED": {"mongodbatlas.process.db.document.rate", map[string]pdata.AttributeValue{
		"status": pdata.NewAttributeValueString("deleted"),
	}},

	// Average rate for operations per second during a selected time period that perform a sort but cannot perform the sort using an index.
	"OPERATIONS_SCAN_AND_ORDER": {"mongodbatlas.process.db.operations.rate", map[string]pdata.AttributeValue{
		"operation": pdata.NewAttributeValueString("scan_and_order"),
	}},

	// Average execution time in milliseconds per read, write, or command operation during a selected time period.
	"OP_EXECUTION_TIME_READS": {"mongodbatlas.process.db.operations.time", map[string]pdata.AttributeValue{
		"status": pdata.NewAttributeValueString("reads"),
	}},
	"OP_EXECUTION_TIME_WRITES": {"mongodbatlas.process.db.operations.time", map[string]pdata.AttributeValue{
		"status": pdata.NewAttributeValueString("writes"),
	}},
	"OP_EXECUTION_TIME_COMMANDS": {"mongodbatlas.process.db.operations.time", map[string]pdata.AttributeValue{
		"status": pdata.NewAttributeValueString("commands"),
	}},

	// Number of times the host restarted within the previous hour.
	"RESTARTS_IN_LAST_HOUR": {"mongodbatlas.process.restarts", map[string]pdata.AttributeValue{
		"span": pdata.NewAttributeValueString("hour"),
	}},

	// Average rate per second to scan index items during queries and query-plan evaluations found in the value of totalKeysExamined from the explain command.
	"QUERY_EXECUTOR_SCANNED": {"mongodbatlas.process.db.query_executor.scanned", map[string]pdata.AttributeValue{
		"type": pdata.NewAttributeValueString("index items"),
	}},

	// Average rate of documents scanned per second during queries and query-plan evaluations found in the value of totalDocsExamined from the explain command.
	"QUERY_EXECUTOR_SCANNED_OBJECTS": {"mongodbatlas.process.db.query_executor.scanned", map[string]pdata.AttributeValue{
		"type": pdata.NewAttributeValueString("objects"),
	}},

	// Ratio of the number of index items scanned to the number of documents returned.
	"QUERY_TARGETING_SCANNED_PER_RETURNED": {"mongodbatlas.process.db.query_targeting.scanned_per_returned", map[string]pdata.AttributeValue{
		"type": pdata.NewAttributeValueString("index items"),
	}},

	// Ratio of the number of documents scanned to the number of documents returned.
	"QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED": {"mongodbatlas.process.db.query_targeting.scanned_per_returned", map[string]pdata.AttributeValue{
		"type": pdata.NewAttributeValueString("objects"),
	}},

	// CPU usage of processes on the host. For hosts with more than one CPU core, this value can exceed 100%.
	"SYSTEM_CPU_USER": {"mongodbatlas.system.cpu.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("user"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},
	"MAX_SYSTEM_CPU_USER": {"mongodbatlas.system.cpu.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("user"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},
	"SYSTEM_CPU_KERNEL": {"mongodbatlas.system.cpu.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("kernel"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},
	"MAX_SYSTEM_CPU_KERNEL": {"mongodbatlas.system.cpu.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("kernel"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},
	"SYSTEM_CPU_NICE": {"mongodbatlas.system.cpu.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("nice"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},
	"MAX_SYSTEM_CPU_NICE": {"mongodbatlas.system.cpu.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("nice"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},
	"SYSTEM_CPU_IOWAIT": {"mongodbatlas.system.cpu.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("iowait"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},
	"MAX_SYSTEM_CPU_IOWAIT": {"mongodbatlas.system.cpu.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("iowait"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},
	"SYSTEM_CPU_IRQ": {"mongodbatlas.system.cpu.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("irq"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},
	"MAX_SYSTEM_CPU_IRQ": {"mongodbatlas.system.cpu.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("irq"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},
	"SYSTEM_CPU_SOFTIRQ": {"mongodbatlas.system.cpu.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("softirq"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},
	"MAX_SYSTEM_CPU_SOFTIRQ": {"mongodbatlas.system.cpu.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("softirq"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},
	"SYSTEM_CPU_GUEST": {"mongodbatlas.system.cpu.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("guest"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},
	"MAX_SYSTEM_CPU_GUEST": {"mongodbatlas.system.cpu.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("guest"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},
	"SYSTEM_CPU_STEAL": {"mongodbatlas.system.cpu.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("steal"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},
	"MAX_SYSTEM_CPU_STEAL": {"mongodbatlas.system.cpu.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("steal"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},

	// CPU usage of processes on the host scaled to a range of 0 to 100% by dividing by the number of CPU cores.
	"SYSTEM_NORMALIZED_CPU_USER": {"mongodbatlas.system.cpu.normalized.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("user"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},
	"MAX_SYSTEM_NORMALIZED_CPU_USER": {"mongodbatlas.system.cpu.normalized.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("user"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},
	"MAX_SYSTEM_NORMALIZED_CPU_NICE": {"mongodbatlas.system.cpu.normalized.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("nice"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},
	"SYSTEM_NORMALIZED_CPU_KERNEL": {"mongodbatlas.system.cpu.normalized.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("kernel"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},
	"MAX_SYSTEM_NORMALIZED_CPU_KERNEL": {"mongodbatlas.system.cpu.normalized.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("kernel"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},
	"SYSTEM_NORMALIZED_CPU_NICE": {"mongodbatlas.system.cpu.normalized.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("nice"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},
	"SYSTEM_NORMALIZED_CPU_IOWAIT": {"mongodbatlas.system.cpu.normalized.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("iowait"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},
	"MAX_SYSTEM_NORMALIZED_CPU_IOWAIT": {"mongodbatlas.system.cpu.normalized.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("iowait"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},
	"SYSTEM_NORMALIZED_CPU_IRQ": {"mongodbatlas.system.cpu.normalized.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("irq"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},
	"MAX_SYSTEM_NORMALIZED_CPU_IRQ": {"mongodbatlas.system.cpu.normalized.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("irq"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},
	"SYSTEM_NORMALIZED_CPU_SOFTIRQ": {"mongodbatlas.system.cpu.normalized.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("softirq"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},
	"MAX_SYSTEM_NORMALIZED_CPU_SOFTIRQ": {"mongodbatlas.system.cpu.normalized.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("softirq"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},
	"SYSTEM_NORMALIZED_CPU_GUEST": {"mongodbatlas.system.cpu.normalized.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("guest"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},
	"MAX_SYSTEM_NORMALIZED_CPU_GUEST": {"mongodbatlas.system.cpu.normalized.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("guest"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},
	"SYSTEM_NORMALIZED_CPU_STEAL": {"mongodbatlas.system.cpu.normalized.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("steal"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},
	"MAX_SYSTEM_NORMALIZED_CPU_STEAL": {"mongodbatlas.system.cpu.normalized.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("steal"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},

	// Physical memory usage, in bytes, that the host uses.
	"SYSTEM_MEMORY_AVAILABLE": {"mongodbatlas.system.memory.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("available"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},
	"MAX_SYSTEM_MEMORY_AVAILABLE": {"mongodbatlas.system.memory.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("available"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},
	"SYSTEM_MEMORY_BUFFERS": {"mongodbatlas.system.memory.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("buffers"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},
	"MAX_SYSTEM_MEMORY_BUFFERS": {"mongodbatlas.system.memory.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("buffers"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},
	"SYSTEM_MEMORY_CACHED": {"mongodbatlas.system.memory.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("cached"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},
	"MAX_SYSTEM_MEMORY_CACHED": {"mongodbatlas.system.memory.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("cached"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},
	"SYSTEM_MEMORY_FREE": {"mongodbatlas.system.memory.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("free"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},
	"MAX_SYSTEM_MEMORY_FREE": {"mongodbatlas.system.memory.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("free"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},
	"SYSTEM_MEMORY_SHARED": {"mongodbatlas.system.memory.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("shared"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},
	"MAX_SYSTEM_MEMORY_SHARED": {"mongodbatlas.system.memory.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("shared"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},
	"SYSTEM_MEMORY_USED": {"mongodbatlas.system.memory.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("used"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},
	"MAX_SYSTEM_MEMORY_USED": {"mongodbatlas.system.memory.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("used"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},

	// Average rate of physical bytes per second that the eth0 network interface received and transmitted.
	"SYSTEM_NETWORK_IN": {"mongodbatlas.system.network.io", map[string]pdata.AttributeValue{
		"direction":   pdata.NewAttributeValueString("receive"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},
	"MAX_SYSTEM_NETWORK_IN": {"mongodbatlas.system.network.io", map[string]pdata.AttributeValue{
		"direction":   pdata.NewAttributeValueString("receive"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},
	"SYSTEM_NETWORK_OUT": {"mongodbatlas.system.network.io", map[string]pdata.AttributeValue{
		"direction":   pdata.NewAttributeValueString("transmit"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},
	"MAX_SYSTEM_NETWORK_OUT": {"mongodbatlas.system.network.io", map[string]pdata.AttributeValue{
		"direction":   pdata.NewAttributeValueString("transmit"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},

	// Total amount of memory that swap uses.
	"SWAP_USAGE_USED": {"mongodbatlas.system.paging.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("used"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},
	"MAX_SWAP_USAGE_USED": {"mongodbatlas.system.paging.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("used"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},
	"SWAP_USAGE_FREE": {"mongodbatlas.system.paging.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("free"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},
	"MAX_SWAP_USAGE_FREE": {"mongodbatlas.system.paging.usage", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("free"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},

	// Total amount of memory written and read from swap.
	"SWAP_IO_IN": {"mongodbatlas.system.paging.io", map[string]pdata.AttributeValue{
		"direction":   pdata.NewAttributeValueString("in"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},
	"MAX_SWAP_IO_IN": {"mongodbatlas.system.paging.io", map[string]pdata.AttributeValue{
		"direction":   pdata.NewAttributeValueString("in"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},
	"SWAP_IO_OUT": {"mongodbatlas.system.paging.io", map[string]pdata.AttributeValue{
		"direction":   pdata.NewAttributeValueString("out"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},
	"MAX_SWAP_IO_OUT": {"mongodbatlas.system.paging.io", map[string]pdata.AttributeValue{
		"direction":   pdata.NewAttributeValueString("out"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},

	// Memory usage, in bytes, that Atlas Search processes use.
	"FTS_PROCESS_RESIDENT_MEMORY": {"mongodbatlas.system.fts.memory.usage", map[string]pdata.AttributeValue{
		"state": pdata.NewAttributeValueString("resident"),
	}},
	"FTS_PROCESS_VIRTUAL_MEMORY": {"mongodbatlas.system.fts.memory.usage", map[string]pdata.AttributeValue{
		"state": pdata.NewAttributeValueString("virtual"),
	}},
	"FTS_PROCESS_SHARED_MEMORY": {"mongodbatlas.system.fts.memory.usage", map[string]pdata.AttributeValue{
		"state": pdata.NewAttributeValueString("shared"),
	}},
	"FTS_MEMORY_MAPPED": {"mongodbatlas.system.fts.memory.usage", map[string]pdata.AttributeValue{
		"state": pdata.NewAttributeValueString("mapped"),
	}},

	// Disk space, in bytes, that Atlas Search indexes use.
	"FTS_DISK_USAGE": {"mongodbatlas.system.fts.disk.utilization", map[string]pdata.AttributeValue{
		"status": pdata.NewAttributeValueString("used"),
	}},

	// Percentage of CPU that Atlas Search processes use.
	"FTS_PROCESS_CPU_USER": {"mongodbatlas.system.fts.cpu.usage", map[string]pdata.AttributeValue{
		"status": pdata.NewAttributeValueString("user"),
	}},
	"FTS_PROCESS_CPU_KERNEL": {"mongodbatlas.system.fts.cpu.usage", map[string]pdata.AttributeValue{
		"status": pdata.NewAttributeValueString("kernel"),
	}},
	"FTS_PROCESS_NORMALIZED_CPU_USER": {"mongodbatlas.system.fts.cpu.normalized.usage", map[string]pdata.AttributeValue{
		"status": pdata.NewAttributeValueString("user"),
	}},
	"FTS_PROCESS_NORMALIZED_CPU_KERNEL": {"mongodbatlas.system.fts.cpu.normalized.usage", map[string]pdata.AttributeValue{
		"status": pdata.NewAttributeValueString("kernel"),
	}},

	// Process Disk Measurements (https://docs.atlas.mongodb.com/reference/api/process-disks-measurements/)

	// Measures throughput of I/O operations for the disk partition used for MongoDB.
	"DISK_PARTITION_IOPS_READ": {"mongodbatlas.disk.partition.iops", map[string]pdata.AttributeValue{
		"direction":   pdata.NewAttributeValueString("read"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},

	"MAX_DISK_PARTITION_IOPS_READ": {"mongodbatlas.disk.partition.iops", map[string]pdata.AttributeValue{
		"direction":   pdata.NewAttributeValueString("read"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},

	"DISK_PARTITION_IOPS_WRITE": {"mongodbatlas.disk.partition.iops", map[string]pdata.AttributeValue{
		"direction":   pdata.NewAttributeValueString("write"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},

	"MAX_DISK_PARTITION_IOPS_WRITE": {"mongodbatlas.disk.partition.iops", map[string]pdata.AttributeValue{
		"direction":   pdata.NewAttributeValueString("write"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},

	"DISK_PARTITION_IOPS_TOTAL": {"mongodbatlas.disk.partition.iops", map[string]pdata.AttributeValue{
		"direction":   pdata.NewAttributeValueString("total"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},

	"MAX_DISK_PARTITION_IOPS_TOTAL": {"mongodbatlas.disk.partition.iops", map[string]pdata.AttributeValue{
		"direction":   pdata.NewAttributeValueString("total"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},

	"DISK_PARTITION_UTILIZATION": {"mongodbatlas.disk.partition.utilization", map[string]pdata.AttributeValue{
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},

	"MAX_DISK_PARTITION_UTILIZATION": {"mongodbatlas.disk.partition.utilization", map[string]pdata.AttributeValue{
		"aggregation": pdata.NewAttributeValueString("max"),
	}},

	// The percentage of time during which requests are being issued to and serviced by the partition.
	// This includes requests from any process, not just MongoDB processes.
	"DISK_PARTITION_LATENCY_READ": {"mongodbatlas.disk.partition.latency", map[string]pdata.AttributeValue{
		"direction":   pdata.NewAttributeValueString("read"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},

	"MAX_DISK_PARTITION_LATENCY_READ": {"mongodbatlas.disk.partition.latency", map[string]pdata.AttributeValue{
		"direction":   pdata.NewAttributeValueString("read"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},

	"DISK_PARTITION_LATENCY_WRITE": {"mongodbatlas.disk.partition.latency", map[string]pdata.AttributeValue{
		"direction":   pdata.NewAttributeValueString("write"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},

	"MAX_DISK_PARTITION_LATENCY_WRITE": {"mongodbatlas.disk.partition.latency", map[string]pdata.AttributeValue{
		"direction":   pdata.NewAttributeValueString("write"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},

	// Measures latency per operation type of the disk partition used by MongoDB.
	"DISK_PARTITION_SPACE_FREE": {"mongodbatlas.disk.partition.space", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("free"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},

	"MAX_DISK_PARTITION_SPACE_FREE": {"mongodbatlas.disk.partition.space", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("free"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},

	"DISK_PARTITION_SPACE_USED": {"mongodbatlas.disk.partition.space", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("used"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},

	"MAX_DISK_PARTITION_SPACE_USED": {"mongodbatlas.disk.partition.space", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("used"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},

	"DISK_PARTITION_SPACE_PERCENT_FREE": {"mongodbatlas.disk.partition.utilization", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("free"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},
	"MAX_DISK_PARTITION_SPACE_PERCENT_FREE": {"mongodbatlas.disk.partition.utilization", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("free"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},
	"DISK_PARTITION_SPACE_PERCENT_USED": {"mongodbatlas.disk.partition.utilization", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("used"),
		"aggregation": pdata.NewAttributeValueString("avg"),
	}},
	"MAX_DISK_PARTITION_SPACE_PERCENT_USED": {"mongodbatlas.disk.partition.utilization", map[string]pdata.AttributeValue{
		"status":      pdata.NewAttributeValueString("used"),
		"aggregation": pdata.NewAttributeValueString("max"),
	}},

	// Process Database Measurements (https://docs.atlas.mongodb.com/reference/api/process-disks-measurements/)
	"DATABASE_COLLECTION_COUNT": {"mongodbatlas.db.counts", map[string]pdata.AttributeValue{
		"type": pdata.NewAttributeValueString("collection"),
	}},
	"DATABASE_INDEX_COUNT": {"mongodbatlas.db.counts", map[string]pdata.AttributeValue{
		"type": pdata.NewAttributeValueString("index"),
	}},
	"DATABASE_EXTENT_COUNT": {"mongodbatlas.db.counts", map[string]pdata.AttributeValue{
		"type": pdata.NewAttributeValueString("extent"),
	}},
	"DATABASE_OBJECT_COUNT": {"mongodbatlas.db.counts", map[string]pdata.AttributeValue{
		"type": pdata.NewAttributeValueString("object"),
	}},
	"DATABASE_VIEW_COUNT": {"mongodbatlas.db.counts", map[string]pdata.AttributeValue{
		"type": pdata.NewAttributeValueString("view"),
	}},
	"DATABASE_AVERAGE_OBJECT_SIZE": {"mongodbatlas.db.size", map[string]pdata.AttributeValue{
		"type": pdata.NewAttributeValueString("object"),
	}},
	"DATABASE_STORAGE_SIZE": {"mongodbatlas.db.size", map[string]pdata.AttributeValue{
		"type": pdata.NewAttributeValueString("storage"),
	}},
	"DATABASE_INDEX_SIZE": {"mongodbatlas.db.size", map[string]pdata.AttributeValue{
		"type": pdata.NewAttributeValueString("index"),
	}},
	"DATABASE_DATA_SIZE": {"mongodbatlas.db.size", map[string]pdata.AttributeValue{
		"type": pdata.NewAttributeValueString("data"),
	}},
}

func mappedMetricByName(name string) (MetricIntf, map[string]pdata.AttributeValue) {
	info, found := metricNameMapping[name]
	if !found {
		return nil, nil
	}

	metricinf := Metrics.ByName(info.metricName)
	return metricinf, info.attributes
}

func inferMetricType(mongoType string) (pdata.MetricDataType, string) {
	var dataType pdata.MetricDataType
	switch mongoType {
	case "BYTES", "MEGABYTES", "SECONDS", "KILOBYTES":
		dataType = pdata.MetricDataTypeSum
	default:
		dataType = pdata.MetricDataTypeGauge
	}

	return dataType, strings.ToLower(mongoType)
}

func buildMetricIntf(meas *mongodbatlas.Measurements) MetricIntf {
	metricType, metricUnits := inferMetricType(meas.Units)
	metricName := fmt.Sprintf("mongodb.atlas.%s", strings.ToLower(meas.Name))
	return &metricImpl{
		metricName,
		func(m pdata.Metric) {
			m.SetDataType(metricType)
			m.SetUnit(metricUnits)
		},
	}
}

func MeasurementsToMetric(meas *mongodbatlas.Measurements, buildUnrecognized bool) (*pdata.Metric, error) {
	intf, attrs := mappedMetricByName(meas.Name)
	if intf == nil {
		if buildUnrecognized {
			intf = buildMetricIntf(meas)
		} else {
			return nil, nil // Not an error- simply skipping undocumented metrics
		}
	}
	m := pdata.NewMetric()
	intf.Init(m)
	switch m.DataType() {
	case pdata.MetricDataTypeGauge:
		datapoints := m.Gauge().DataPoints()
		err := addDataPoints(datapoints, meas, attrs)
		if err != nil {
			return nil, err
		}
	case pdata.MetricDataTypeSum:
		datapoints := m.Sum().DataPoints()
		err := addDataPoints(datapoints, meas, attrs)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unrecognized data type for metric '%s'", meas.Name)
	}

	return &m, nil
}

func addDataPoints(datapoints pdata.NumberDataPointSlice, meas *mongodbatlas.Measurements, attrs map[string]pdata.AttributeValue) error {
	for _, point := range meas.DataPoints {
		if point.Value != nil {
			dp := datapoints.AppendEmpty()
			curTime, err := time.Parse(time.RFC3339, point.Timestamp)
			if err != nil {
				return err
			}
			for k, v := range attrs {
				dp.Attributes().Upsert(k, v)
			}
			dp.SetTimestamp(pdata.NewTimestampFromTime(curTime))
			dp.SetDoubleVal(float64(*point.Value))
		}
	}
	return nil
}
