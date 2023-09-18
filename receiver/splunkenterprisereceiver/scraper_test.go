// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package splunkenterprisereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/splunkenterprisereceiver"

import (
	"context"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.opentelemetry.io/collector/receiver/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/golden"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest/pmetrictest"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/splunkenterprisereceiver/internal/metadata"
)

// handler function for mock server
func mockIndexerThroughput(w http.ResponseWriter, _ *http.Request) {
	status := http.StatusOK
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write([]byte(`{"links":{},"origin":"https://somehost:8089/services/server/introspection/indexer","updated":"2023-07-31T21:41:07+00:00","generator":{"build":"82c987350fde","version":"9.0.1"},"entry":[{"name":"indexer","id":"https://34.213.134.166:8089/services/server/introspection/indexer/indexer","updated":"1970-01-01T00:00:00+00:00","links":{"alternate":"/services/server/introspection/indexer/indexer","list":"/services/server/introspection/indexer/indexer","edit":"/services/server/introspection/indexer/indexer"},"author":"system","acl":{"app":"","can_list":true,"can_write":true,"modifiable":false,"owner":"system","perms":{"read":["admin","splunk-system-role"],"write":["admin","splunk-system-role"]},"removable":false,"sharing":"system"},"content":{"average_KBps":25.579690815904478,"eai:acl":null,"reason":"","status":"normal"}}],"paging":{"total":1,"perPage":30,"offset":0},"messages":[]}`))
}

func mockIndexesExtended(w http.ResponseWriter, _ *http.Request) {
	status := http.StatusOK
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write([]byte(`{"links":{},"origin":"https://somehost:8089/services/data/indexes-extended","updated":"2023-09-18T13:17:38+00:00","generator":{"build":"82c987350fde","version":"9.0.1"},"entry":[{"name":"_audit","id":"https://somehost:8089/servicesNS/nobody/system/data/indexes-extended/_audit","updated":"2023-09-18T13:17:38+00:00","links":{"alternate":"/servicesNS/nobody/system/data/indexes-extended/_audit","list":"/servicesNS/nobody/system/data/indexes-extended/_audit"},"author":"nobody","acl":{"app":"system","can_list":true,"can_write":true,"modifiable":false,"owner":"nobody","perms":{"read":["*"],"write":[]},"removable":false,"sharing":"system"},"content":{"archiver.enableDataArchive":false,"archiver.maxDataArchiveRetentionPeriod":0,"assureUTF8":false,"bucketMerge.maxMergeSizeMB":1000,"bucketMerge.maxMergeTimeSpanSecs":7776000,"bucketMerge.minMergeSizeMB":750,"bucketMerging":false,"bucketRebuildMemoryHint":0,"bucket_dirs":{"cold":{"capacity":"0.000"},"home":{"capacity":"0.000","event_count":"107267027","event_max_time":"1695042546","event_min_time":"1663795123","hot_bucket_count":"1","warm_bucket_count":"50","warm_bucket_size":"19641.027"},"thawed":null},"coldPath":"$SPLUNK_DB/audit/colddb","coldPath.maxDataSizeMB":0,"coldPath_expanded":"/opt/splunk/var/lib/splunk/audit/colddb","coldToFrozenDir":"","coldToFrozenScript":"","compressRawdata":true,"currentDBSizeMB":"19855","datamodel_summary_size":"1342.055","datatype":"event","defaultDatabase":"main","disabled":false,"eai:acl":null,"enableDataIntegrityControl":false,"enableOnlineBucketRepair":true,"enableRealtimeSearch":true,"enableTsidxReduction":false,"federated.dataset":"","federated.provider":"","fileSystemExecutorWorkers":5,"frozenTimePeriodInSecs":188697600,"homePath":"$SPLUNK_DB/audit/db","homePath.maxDataSizeMB":0,"homePath_expanded":"/opt/splunk/var/lib/splunk/audit/db","hotBucketStreaming.deleteHotsAfterRestart":false,"hotBucketStreaming.extraBucketBuildingCmdlineArgs":null,"hotBucketStreaming.removeRemoteSlicesOnRoll":false,"hotBucketStreaming.reportStatus":false,"hotBucketStreaming.sendSlices":false,"hotBucketTimeRefreshInterval":10,"indexThreads":"auto","isInternal":true,"isReady":true,"isVirtual":false,"journalCompression":"zstd","lastChanceIndex":null,"lastInitSequenceNumber":1,"lastInitTime":1694724553,"maxBloomBackfillBucketAge":"30d","maxBucketSizeCacheEntries":0,"maxConcurrentOptimizes":6,"maxDataSize":"auto","maxGlobalDataSizeMB":0,"maxGlobalRawDataSizeMB":0,"maxHotBuckets":"auto","maxHotIdleSecs":0,"maxHotSpanSecs":7776000,"maxMemMB":5,"maxMetaEntries":1000000,"maxRunningProcessGroups":8,"maxRunningProcessGroupsLowPriority":1,"maxTime":"2023-09-18T13:17:35+0000","maxTimeUnreplicatedNoAcks":300,"maxTimeUnreplicatedWithAcks":60,"maxTotalDataSizeMB":500000,"maxWarmDBCount":300,"memPoolMB":"auto","metric.compressionBlockSize":1024,"metric.enableFloatingPointCompression":true,"metric.maxHotBuckets":"auto","metric.splitByIndexKeys":"","metric.stubOutRawdataJournal":true,"metric.timestampResolution":"s","metric.tsidxTargetSizeMB":1500,"minHotIdleSecsBeforeForceRoll":0,"minRawFileSyncSecs":"disable","minStreamGroupQueueSize":2000,"minTime":"2022-09-21T21:18:43+0000","name":"_audit","partialServiceMetaPeriod":0,"processTrackerServiceInterval":1,"quarantineFutureSecs":2592000,"quarantinePastSecs":77760000,"rawChunkSizeBytes":131072,"repFactor":0,"rotatePeriodInSecs":60,"rtRouterQueueSize":null,"rtRouterThreads":null,"selfStorageThreads":null,"serviceInactiveIndexesPeriod":60,"serviceMetaPeriod":25,"serviceOnlyAsNeeded":true,"serviceSubtaskTimingPeriod":30,"splitByIndexKeys":"","streamingTargetTsidxSyncPeriodMsec":5000,"summaryHomePath_expanded":"/opt/splunk/var/lib/splunk/audit/summary","suppressBannerList":"","suspendHotRollByDeleteQuery":false,"sync":0,"syncMeta":true,"thawedPath":"$SPLUNK_DB/audit/thaweddb","thawedPath_expanded":"/opt/splunk/var/lib/splunk/audit/thaweddb","throttleCheckPeriod":15,"timePeriodInSecBeforeTsidxReduction":604800,"totalEventCount":108411855,"total_bucket_count":"51","total_capacity":"500000.000","total_event_count":"107267027","total_raw_size":"67544.059","total_size":"19854.039","tsidxDedupPostingsListMaxTermsLimit":8388608,"tsidxReductionCheckPeriodInSec":600,"tsidxTargetSizeMB":1500,"tsidxWritingLevel":null,"tstatsHomePath":"volume:_splunk_summaries/audit/datamodel_summary","tstatsHomePath_expanded":"/opt/splunk/var/lib/splunk/audit/datamodel_summary","waitPeriodInSecsForManifestWrite":60,"warmToColdScript":""}}],"paging":{"total":40,"perPage":1,"offset":0},"messages":[]}`))
}

func mockIntrospectionQueues(w http.ResponseWriter, _ *http.Request) {
	status := http.StatusOK
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write([]byte(`{"links":{},"origin":"https://somehost:8089/services/server/introspection/queues","updated":"2023-09-18T13:37:45+00:00","generator":{"build":"82c987350fde","version":"9.0.1"},"entry":[{"name":"AEQ","id":"https://somehost:8089/services/server/introspection/queues/AEQ","updated":"1970-01-01T00:00:00+00:00","links":{"alternate":"/services/server/introspection/queues/AEQ","list":"/services/server/introspection/queues/AEQ","edit":"/services/server/introspection/queues/AEQ"},"author":"system","acl":{"app":"","can_list":true,"can_write":true,"modifiable":false,"owner":"system","perms":{"read":["admin","splunk-system-role"],"write":["admin","splunk-system-role"]},"removable":false,"sharing":"system"},"content":{"cntr_1_lookback_time":60,"cntr_2_lookback_time":600,"cntr_3_lookback_time":900,"current_size":1,"current_size_bytes":100,"eai:acl":null,"largest_size":3,"max_size_bytes":512000,"sampling_interval":1,"smallest_size":0,"value_cntr1_size_bytes_lookback":0,"value_cntr1_size_lookback":0,"value_cntr2_size_bytes_lookback":0,"value_cntr2_size_lookback":0,"value_cntr3_size_bytes_lookback":0,"value_cntr3_size_lookback":0}}],"paging":{"total":13,"perPage":1,"offset":0},"messages":[]}`))
}

// mock server create
func createMockServer() *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch strings.TrimSpace(r.URL.Path) {
		case "/services/server/introspection/indexer":
			mockIndexerThroughput(w, r)
		case "/services/data/indexes-extended":
			mockIndexesExtended(w, r)
		case "/services/server/introspection/queues":
			mockIntrospectionQueues(w, r)
		default:
			http.NotFoundHandler().ServeHTTP(w, r)
		}
	}))

	return ts
}

func TestScraper(t *testing.T) {
	ts := createMockServer()
	defer ts.Close()

	// in the future add more metrics
	metricsettings := metadata.MetricsBuilderConfig{}
	metricsettings.Metrics.SplunkIndexerThroughput.Enabled = true
	metricsettings.Metrics.SplunkDataIndexesExtendedTotalSize.Enabled = true
	metricsettings.Metrics.SplunkDataIndexesExtendedEventCount.Enabled = true
	metricsettings.Metrics.SplunkDataIndexesExtendedBucketCount.Enabled = true
	metricsettings.Metrics.SplunkDataIndexesExtendedRawSize.Enabled = true
	metricsettings.Metrics.SplunkDataIndexesExtendedBucketEventCount.Enabled = true
	metricsettings.Metrics.SplunkDataIndexesExtendedBucketHotCount.Enabled = true
	metricsettings.Metrics.SplunkDataIndexesExtendedBucketWarmCount.Enabled = true
	metricsettings.Metrics.SplunkServerIntrospectionQueuesCurrent.Enabled = true
	metricsettings.Metrics.SplunkServerIntrospectionQueuesCurrentBytes.Enabled = true	

	cfg := &Config{
		Username:          "admin",
		Password:          "securityFirst",
		MaxSearchWaitTime: 11 * time.Second,
		HTTPClientSettings: confighttp.HTTPClientSettings{
			Endpoint: ts.URL,
		},
		ScraperControllerSettings: scraperhelper.ScraperControllerSettings{
			CollectionInterval: 10 * time.Second,
			InitialDelay:       1 * time.Second,
		},
		MetricsBuilderConfig: metricsettings,
	}

	require.NoError(t, cfg.Validate())

	scraper := newSplunkMetricsScraper(receivertest.NewNopCreateSettings(), cfg)
	client := newSplunkEntClient(cfg)
	scraper.splunkClient = &client

	actualMetrics, err := scraper.scrape(context.Background())
	require.NoError(t, err)

	expectedFile := filepath.Join("testdata", "scraper", "expected.yaml")
  // golden.WriteMetrics(t, expectedFile, actualMetrics) // run tests with this line whenever metrics are modified

	expectedMetrics, err := golden.ReadMetrics(expectedFile)
	require.NoError(t, err)

	require.NoError(t, pmetrictest.CompareMetrics(expectedMetrics, actualMetrics, pmetrictest.IgnoreStartTimestamp(), pmetrictest.IgnoreTimestamp()))
}
