// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/confmap"
)

// MetricConfig provides common config for a particular metric.
type MetricConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms)
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsConfig provides config for splunkenterprise metrics.
type MetricsConfig struct {
	SplunkAggregationQueueRatio                 MetricConfig `mapstructure:"splunk.aggregation.queue.ratio"`
	SplunkBucketsSearchableStatus               MetricConfig `mapstructure:"splunk.buckets.searchable.status"`
	SplunkDataIndexesExtendedBucketCount        MetricConfig `mapstructure:"splunk.data.indexes.extended.bucket.count"`
	SplunkDataIndexesExtendedBucketEventCount   MetricConfig `mapstructure:"splunk.data.indexes.extended.bucket.event.count"`
	SplunkDataIndexesExtendedBucketHotCount     MetricConfig `mapstructure:"splunk.data.indexes.extended.bucket.hot.count"`
	SplunkDataIndexesExtendedBucketWarmCount    MetricConfig `mapstructure:"splunk.data.indexes.extended.bucket.warm.count"`
	SplunkDataIndexesExtendedEventCount         MetricConfig `mapstructure:"splunk.data.indexes.extended.event.count"`
	SplunkDataIndexesExtendedRawSize            MetricConfig `mapstructure:"splunk.data.indexes.extended.raw.size"`
	SplunkDataIndexesExtendedTotalSize          MetricConfig `mapstructure:"splunk.data.indexes.extended.total.size"`
	SplunkIndexerAvgRate                        MetricConfig `mapstructure:"splunk.indexer.avg.rate"`
	SplunkIndexerCPUTime                        MetricConfig `mapstructure:"splunk.indexer.cpu.time"`
	SplunkIndexerQueueRatio                     MetricConfig `mapstructure:"splunk.indexer.queue.ratio"`
	SplunkIndexerRawWriteTime                   MetricConfig `mapstructure:"splunk.indexer.raw.write.time"`
	SplunkIndexerThroughput                     MetricConfig `mapstructure:"splunk.indexer.throughput"`
	SplunkIndexesAvgSize                        MetricConfig `mapstructure:"splunk.indexes.avg.size"`
	SplunkIndexesAvgUsage                       MetricConfig `mapstructure:"splunk.indexes.avg.usage"`
	SplunkIndexesBucketCount                    MetricConfig `mapstructure:"splunk.indexes.bucket.count"`
	SplunkIndexesMedianDataAge                  MetricConfig `mapstructure:"splunk.indexes.median.data.age"`
	SplunkIndexesSize                           MetricConfig `mapstructure:"splunk.indexes.size"`
	SplunkIoAvgIops                             MetricConfig `mapstructure:"splunk.io.avg.iops"`
	SplunkLicenseIndexUsage                     MetricConfig `mapstructure:"splunk.license.index.usage"`
	SplunkParseQueueRatio                       MetricConfig `mapstructure:"splunk.parse.queue.ratio"`
	SplunkPipelineSetCount                      MetricConfig `mapstructure:"splunk.pipeline.set.count"`
	SplunkSchedulerAvgExecutionLatency          MetricConfig `mapstructure:"splunk.scheduler.avg.execution.latency"`
	SplunkSchedulerAvgRunTime                   MetricConfig `mapstructure:"splunk.scheduler.avg.run.time"`
	SplunkSchedulerCompletionRatio              MetricConfig `mapstructure:"splunk.scheduler.completion.ratio"`
	SplunkServerIntrospectionQueuesCurrent      MetricConfig `mapstructure:"splunk.server.introspection.queues.current"`
	SplunkServerIntrospectionQueuesCurrentBytes MetricConfig `mapstructure:"splunk.server.introspection.queues.current.bytes"`
	SplunkTypingQueueRatio                      MetricConfig `mapstructure:"splunk.typing.queue.ratio"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		SplunkAggregationQueueRatio: MetricConfig{
			Enabled: true,
		},
		SplunkBucketsSearchableStatus: MetricConfig{
			Enabled: true,
		},
		SplunkDataIndexesExtendedBucketCount: MetricConfig{
			Enabled: false,
		},
		SplunkDataIndexesExtendedBucketEventCount: MetricConfig{
			Enabled: false,
		},
		SplunkDataIndexesExtendedBucketHotCount: MetricConfig{
			Enabled: false,
		},
		SplunkDataIndexesExtendedBucketWarmCount: MetricConfig{
			Enabled: false,
		},
		SplunkDataIndexesExtendedEventCount: MetricConfig{
			Enabled: false,
		},
		SplunkDataIndexesExtendedRawSize: MetricConfig{
			Enabled: false,
		},
		SplunkDataIndexesExtendedTotalSize: MetricConfig{
			Enabled: false,
		},
		SplunkIndexerAvgRate: MetricConfig{
			Enabled: true,
		},
		SplunkIndexerCPUTime: MetricConfig{
			Enabled: true,
		},
		SplunkIndexerQueueRatio: MetricConfig{
			Enabled: true,
		},
		SplunkIndexerRawWriteTime: MetricConfig{
			Enabled: true,
		},
		SplunkIndexerThroughput: MetricConfig{
			Enabled: false,
		},
		SplunkIndexesAvgSize: MetricConfig{
			Enabled: true,
		},
		SplunkIndexesAvgUsage: MetricConfig{
			Enabled: true,
		},
		SplunkIndexesBucketCount: MetricConfig{
			Enabled: true,
		},
		SplunkIndexesMedianDataAge: MetricConfig{
			Enabled: true,
		},
		SplunkIndexesSize: MetricConfig{
			Enabled: true,
		},
		SplunkIoAvgIops: MetricConfig{
			Enabled: true,
		},
		SplunkLicenseIndexUsage: MetricConfig{
			Enabled: true,
		},
		SplunkParseQueueRatio: MetricConfig{
			Enabled: true,
		},
		SplunkPipelineSetCount: MetricConfig{
			Enabled: true,
		},
		SplunkSchedulerAvgExecutionLatency: MetricConfig{
			Enabled: true,
		},
		SplunkSchedulerAvgRunTime: MetricConfig{
			Enabled: true,
		},
		SplunkSchedulerCompletionRatio: MetricConfig{
			Enabled: true,
		},
		SplunkServerIntrospectionQueuesCurrent: MetricConfig{
			Enabled: false,
		},
		SplunkServerIntrospectionQueuesCurrentBytes: MetricConfig{
			Enabled: false,
		},
		SplunkTypingQueueRatio: MetricConfig{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for splunkenterprise metrics builder.
type MetricsBuilderConfig struct {
	Metrics MetricsConfig `mapstructure:"metrics"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics: DefaultMetricsConfig(),
	}
}
