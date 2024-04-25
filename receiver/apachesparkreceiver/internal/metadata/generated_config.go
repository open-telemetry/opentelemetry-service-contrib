// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/filter"
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

// MetricsConfig provides config for apachespark metrics.
type MetricsConfig struct {
	SparkDriverBlockManagerDiskUsage                   MetricConfig `mapstructure:"spark.driver.block_manager.disk.usage"`
	SparkDriverBlockManagerMemoryUsage                 MetricConfig `mapstructure:"spark.driver.block_manager.memory.usage"`
	SparkDriverCodeGeneratorCompilationAverageTime     MetricConfig `mapstructure:"spark.driver.code_generator.compilation.average_time"`
	SparkDriverCodeGeneratorCompilationCount           MetricConfig `mapstructure:"spark.driver.code_generator.compilation.count"`
	SparkDriverCodeGeneratorGeneratedClassAverageSize  MetricConfig `mapstructure:"spark.driver.code_generator.generated_class.average_size"`
	SparkDriverCodeGeneratorGeneratedClassCount        MetricConfig `mapstructure:"spark.driver.code_generator.generated_class.count"`
	SparkDriverCodeGeneratorGeneratedMethodAverageSize MetricConfig `mapstructure:"spark.driver.code_generator.generated_method.average_size"`
	SparkDriverCodeGeneratorGeneratedMethodCount       MetricConfig `mapstructure:"spark.driver.code_generator.generated_method.count"`
	SparkDriverCodeGeneratorSourceCodeAverageSize      MetricConfig `mapstructure:"spark.driver.code_generator.source_code.average_size"`
	SparkDriverCodeGeneratorSourceCodeOperations       MetricConfig `mapstructure:"spark.driver.code_generator.source_code.operations"`
	SparkDriverDagSchedulerJobActive                   MetricConfig `mapstructure:"spark.driver.dag_scheduler.job.active"`
	SparkDriverDagSchedulerJobCount                    MetricConfig `mapstructure:"spark.driver.dag_scheduler.job.count"`
	SparkDriverDagSchedulerStageCount                  MetricConfig `mapstructure:"spark.driver.dag_scheduler.stage.count"`
	SparkDriverDagSchedulerStageFailed                 MetricConfig `mapstructure:"spark.driver.dag_scheduler.stage.failed"`
	SparkDriverExecutorGcOperations                    MetricConfig `mapstructure:"spark.driver.executor.gc.operations"`
	SparkDriverExecutorGcTime                          MetricConfig `mapstructure:"spark.driver.executor.gc.time"`
	SparkDriverExecutorMemoryExecution                 MetricConfig `mapstructure:"spark.driver.executor.memory.execution"`
	SparkDriverExecutorMemoryJvm                       MetricConfig `mapstructure:"spark.driver.executor.memory.jvm"`
	SparkDriverExecutorMemoryPool                      MetricConfig `mapstructure:"spark.driver.executor.memory.pool"`
	SparkDriverExecutorMemoryStorage                   MetricConfig `mapstructure:"spark.driver.executor.memory.storage"`
	SparkDriverHiveExternalCatalogFileCacheHits        MetricConfig `mapstructure:"spark.driver.hive_external_catalog.file_cache_hits"`
	SparkDriverHiveExternalCatalogFilesDiscovered      MetricConfig `mapstructure:"spark.driver.hive_external_catalog.files_discovered"`
	SparkDriverHiveExternalCatalogHiveClientCalls      MetricConfig `mapstructure:"spark.driver.hive_external_catalog.hive_client_calls"`
	SparkDriverHiveExternalCatalogParallelListingJobs  MetricConfig `mapstructure:"spark.driver.hive_external_catalog.parallel_listing_jobs"`
	SparkDriverHiveExternalCatalogPartitionsFetched    MetricConfig `mapstructure:"spark.driver.hive_external_catalog.partitions_fetched"`
	SparkDriverJvmCPUTime                              MetricConfig `mapstructure:"spark.driver.jvm_cpu_time"`
	SparkDriverLiveListenerBusDropped                  MetricConfig `mapstructure:"spark.driver.live_listener_bus.dropped"`
	SparkDriverLiveListenerBusPosted                   MetricConfig `mapstructure:"spark.driver.live_listener_bus.posted"`
	SparkDriverLiveListenerBusProcessingTimeAverage    MetricConfig `mapstructure:"spark.driver.live_listener_bus.processing_time.average"`
	SparkDriverLiveListenerBusQueueSize                MetricConfig `mapstructure:"spark.driver.live_listener_bus.queue_size"`
	SparkExecutorDiskUsage                             MetricConfig `mapstructure:"spark.executor.disk.usage"`
	SparkExecutorGcTime                                MetricConfig `mapstructure:"spark.executor.gc_time"`
	SparkExecutorInputSize                             MetricConfig `mapstructure:"spark.executor.input_size"`
	SparkExecutorMemoryUsage                           MetricConfig `mapstructure:"spark.executor.memory.usage"`
	SparkExecutorShuffleIoSize                         MetricConfig `mapstructure:"spark.executor.shuffle.io.size"`
	SparkExecutorStorageMemoryUsage                    MetricConfig `mapstructure:"spark.executor.storage_memory.usage"`
	SparkExecutorTaskActive                            MetricConfig `mapstructure:"spark.executor.task.active"`
	SparkExecutorTaskLimit                             MetricConfig `mapstructure:"spark.executor.task.limit"`
	SparkExecutorTaskResult                            MetricConfig `mapstructure:"spark.executor.task.result"`
	SparkExecutorTime                                  MetricConfig `mapstructure:"spark.executor.time"`
	SparkJobStageActive                                MetricConfig `mapstructure:"spark.job.stage.active"`
	SparkJobStageResult                                MetricConfig `mapstructure:"spark.job.stage.result"`
	SparkJobTaskActive                                 MetricConfig `mapstructure:"spark.job.task.active"`
	SparkJobTaskResult                                 MetricConfig `mapstructure:"spark.job.task.result"`
	SparkStageDiskSpilled                              MetricConfig `mapstructure:"spark.stage.disk.spilled"`
	SparkStageExecutorCPUTime                          MetricConfig `mapstructure:"spark.stage.executor.cpu_time"`
	SparkStageExecutorRunTime                          MetricConfig `mapstructure:"spark.stage.executor.run_time"`
	SparkStageIoRecords                                MetricConfig `mapstructure:"spark.stage.io.records"`
	SparkStageIoSize                                   MetricConfig `mapstructure:"spark.stage.io.size"`
	SparkStageJvmGcTime                                MetricConfig `mapstructure:"spark.stage.jvm_gc_time"`
	SparkStageMemoryPeak                               MetricConfig `mapstructure:"spark.stage.memory.peak"`
	SparkStageMemorySpilled                            MetricConfig `mapstructure:"spark.stage.memory.spilled"`
	SparkStageShuffleBlocksFetched                     MetricConfig `mapstructure:"spark.stage.shuffle.blocks_fetched"`
	SparkStageShuffleFetchWaitTime                     MetricConfig `mapstructure:"spark.stage.shuffle.fetch_wait_time"`
	SparkStageShuffleIoDisk                            MetricConfig `mapstructure:"spark.stage.shuffle.io.disk"`
	SparkStageShuffleIoReadSize                        MetricConfig `mapstructure:"spark.stage.shuffle.io.read.size"`
	SparkStageShuffleIoRecords                         MetricConfig `mapstructure:"spark.stage.shuffle.io.records"`
	SparkStageShuffleIoWriteSize                       MetricConfig `mapstructure:"spark.stage.shuffle.io.write.size"`
	SparkStageShuffleWriteTime                         MetricConfig `mapstructure:"spark.stage.shuffle.write_time"`
	SparkStageStatus                                   MetricConfig `mapstructure:"spark.stage.status"`
	SparkStageTaskActive                               MetricConfig `mapstructure:"spark.stage.task.active"`
	SparkStageTaskResult                               MetricConfig `mapstructure:"spark.stage.task.result"`
	SparkStageTaskResultSize                           MetricConfig `mapstructure:"spark.stage.task.result_size"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		SparkDriverBlockManagerDiskUsage: MetricConfig{
			Enabled: true,
		},
		SparkDriverBlockManagerMemoryUsage: MetricConfig{
			Enabled: true,
		},
		SparkDriverCodeGeneratorCompilationAverageTime: MetricConfig{
			Enabled: true,
		},
		SparkDriverCodeGeneratorCompilationCount: MetricConfig{
			Enabled: true,
		},
		SparkDriverCodeGeneratorGeneratedClassAverageSize: MetricConfig{
			Enabled: true,
		},
		SparkDriverCodeGeneratorGeneratedClassCount: MetricConfig{
			Enabled: true,
		},
		SparkDriverCodeGeneratorGeneratedMethodAverageSize: MetricConfig{
			Enabled: true,
		},
		SparkDriverCodeGeneratorGeneratedMethodCount: MetricConfig{
			Enabled: true,
		},
		SparkDriverCodeGeneratorSourceCodeAverageSize: MetricConfig{
			Enabled: true,
		},
		SparkDriverCodeGeneratorSourceCodeOperations: MetricConfig{
			Enabled: true,
		},
		SparkDriverDagSchedulerJobActive: MetricConfig{
			Enabled: true,
		},
		SparkDriverDagSchedulerJobCount: MetricConfig{
			Enabled: true,
		},
		SparkDriverDagSchedulerStageCount: MetricConfig{
			Enabled: true,
		},
		SparkDriverDagSchedulerStageFailed: MetricConfig{
			Enabled: true,
		},
		SparkDriverExecutorGcOperations: MetricConfig{
			Enabled: true,
		},
		SparkDriverExecutorGcTime: MetricConfig{
			Enabled: true,
		},
		SparkDriverExecutorMemoryExecution: MetricConfig{
			Enabled: true,
		},
		SparkDriverExecutorMemoryJvm: MetricConfig{
			Enabled: true,
		},
		SparkDriverExecutorMemoryPool: MetricConfig{
			Enabled: true,
		},
		SparkDriverExecutorMemoryStorage: MetricConfig{
			Enabled: true,
		},
		SparkDriverHiveExternalCatalogFileCacheHits: MetricConfig{
			Enabled: true,
		},
		SparkDriverHiveExternalCatalogFilesDiscovered: MetricConfig{
			Enabled: true,
		},
		SparkDriverHiveExternalCatalogHiveClientCalls: MetricConfig{
			Enabled: true,
		},
		SparkDriverHiveExternalCatalogParallelListingJobs: MetricConfig{
			Enabled: true,
		},
		SparkDriverHiveExternalCatalogPartitionsFetched: MetricConfig{
			Enabled: true,
		},
		SparkDriverJvmCPUTime: MetricConfig{
			Enabled: true,
		},
		SparkDriverLiveListenerBusDropped: MetricConfig{
			Enabled: true,
		},
		SparkDriverLiveListenerBusPosted: MetricConfig{
			Enabled: true,
		},
		SparkDriverLiveListenerBusProcessingTimeAverage: MetricConfig{
			Enabled: true,
		},
		SparkDriverLiveListenerBusQueueSize: MetricConfig{
			Enabled: true,
		},
		SparkExecutorDiskUsage: MetricConfig{
			Enabled: true,
		},
		SparkExecutorGcTime: MetricConfig{
			Enabled: true,
		},
		SparkExecutorInputSize: MetricConfig{
			Enabled: true,
		},
		SparkExecutorMemoryUsage: MetricConfig{
			Enabled: true,
		},
		SparkExecutorShuffleIoSize: MetricConfig{
			Enabled: true,
		},
		SparkExecutorStorageMemoryUsage: MetricConfig{
			Enabled: true,
		},
		SparkExecutorTaskActive: MetricConfig{
			Enabled: true,
		},
		SparkExecutorTaskLimit: MetricConfig{
			Enabled: true,
		},
		SparkExecutorTaskResult: MetricConfig{
			Enabled: true,
		},
		SparkExecutorTime: MetricConfig{
			Enabled: true,
		},
		SparkJobStageActive: MetricConfig{
			Enabled: true,
		},
		SparkJobStageResult: MetricConfig{
			Enabled: true,
		},
		SparkJobTaskActive: MetricConfig{
			Enabled: true,
		},
		SparkJobTaskResult: MetricConfig{
			Enabled: true,
		},
		SparkStageDiskSpilled: MetricConfig{
			Enabled: true,
		},
		SparkStageExecutorCPUTime: MetricConfig{
			Enabled: true,
		},
		SparkStageExecutorRunTime: MetricConfig{
			Enabled: true,
		},
		SparkStageIoRecords: MetricConfig{
			Enabled: true,
		},
		SparkStageIoSize: MetricConfig{
			Enabled: true,
		},
		SparkStageJvmGcTime: MetricConfig{
			Enabled: true,
		},
		SparkStageMemoryPeak: MetricConfig{
			Enabled: true,
		},
		SparkStageMemorySpilled: MetricConfig{
			Enabled: true,
		},
		SparkStageShuffleBlocksFetched: MetricConfig{
			Enabled: true,
		},
		SparkStageShuffleFetchWaitTime: MetricConfig{
			Enabled: true,
		},
		SparkStageShuffleIoDisk: MetricConfig{
			Enabled: true,
		},
		SparkStageShuffleIoReadSize: MetricConfig{
			Enabled: true,
		},
		SparkStageShuffleIoRecords: MetricConfig{
			Enabled: true,
		},
		SparkStageShuffleIoWriteSize: MetricConfig{
			Enabled: true,
		},
		SparkStageShuffleWriteTime: MetricConfig{
			Enabled: true,
		},
		SparkStageStatus: MetricConfig{
			Enabled: true,
		},
		SparkStageTaskActive: MetricConfig{
			Enabled: true,
		},
		SparkStageTaskResult: MetricConfig{
			Enabled: true,
		},
		SparkStageTaskResultSize: MetricConfig{
			Enabled: true,
		},
	}
}

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool `mapstructure:"enabled"`
	// Experimental: MetricsInclude defines a list of filters for attribute values.
	// If the list is not empty, only metrics with matching resource attribute values will be emitted.
	MetricsInclude []filter.Config `mapstructure:"metrics_include"`
	// Experimental: MetricsExclude defines a list of filters for attribute values.
	// If the list is not empty, metrics with matching resource attribute values will not be emitted.
	// MetricsInclude has higher priority than MetricsExclude.
	MetricsExclude []filter.Config `mapstructure:"metrics_exclude"`

	enabledSetByUser bool
}

func (rac *ResourceAttributeConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(rac)
	if err != nil {
		return err
	}
	rac.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// ResourceAttributesConfig provides config for apachespark resource attributes.
type ResourceAttributesConfig struct {
	SparkApplicationID   ResourceAttributeConfig `mapstructure:"spark.application.id"`
	SparkApplicationName ResourceAttributeConfig `mapstructure:"spark.application.name"`
	SparkExecutorID      ResourceAttributeConfig `mapstructure:"spark.executor.id"`
	SparkJobID           ResourceAttributeConfig `mapstructure:"spark.job.id"`
	SparkStageAttemptID  ResourceAttributeConfig `mapstructure:"spark.stage.attempt.id"`
	SparkStageID         ResourceAttributeConfig `mapstructure:"spark.stage.id"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		SparkApplicationID: ResourceAttributeConfig{
			Enabled: true,
		},
		SparkApplicationName: ResourceAttributeConfig{
			Enabled: true,
		},
		SparkExecutorID: ResourceAttributeConfig{
			Enabled: true,
		},
		SparkJobID: ResourceAttributeConfig{
			Enabled: true,
		},
		SparkStageAttemptID: ResourceAttributeConfig{
			Enabled: false,
		},
		SparkStageID: ResourceAttributeConfig{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for apachespark metrics builder.
type MetricsBuilderConfig struct {
	Metrics            MetricsConfig            `mapstructure:"metrics"`
	ResourceAttributes ResourceAttributesConfig `mapstructure:"resource_attributes"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics:            DefaultMetricsConfig(),
		ResourceAttributes: DefaultResourceAttributesConfig(),
	}
}
