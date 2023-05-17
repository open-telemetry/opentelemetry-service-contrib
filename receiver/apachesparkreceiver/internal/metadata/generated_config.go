// Code generated by mdatagen. DO NOT EDIT.

package metadata

import "go.opentelemetry.io/collector/confmap"

// MetricConfig provides common config for a particular metric.
type MetricConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms, confmap.WithErrorUnused())
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsConfig provides config for apachespark metrics.
type MetricsConfig struct {
	SparkDriverBlockManagerDiskSpaceUsed                    MetricConfig `mapstructure:"spark.driver.block_manager.disk.space_used"`
	SparkDriverBlockManagerMemoryRemaining                  MetricConfig `mapstructure:"spark.driver.block_manager.memory.remaining"`
	SparkDriverBlockManagerMemoryUsed                       MetricConfig `mapstructure:"spark.driver.block_manager.memory.used"`
	SparkDriverCodeGeneratorCompilationAverageTime          MetricConfig `mapstructure:"spark.driver.code_generator.compilation.average_time"`
	SparkDriverCodeGeneratorCompilationCount                MetricConfig `mapstructure:"spark.driver.code_generator.compilation.count"`
	SparkDriverCodeGeneratorGeneratedClassAverageSize       MetricConfig `mapstructure:"spark.driver.code_generator.generated_class.average_size"`
	SparkDriverCodeGeneratorGeneratedClassCount             MetricConfig `mapstructure:"spark.driver.code_generator.generated_class.count"`
	SparkDriverCodeGeneratorGeneratedMethodAverageSize      MetricConfig `mapstructure:"spark.driver.code_generator.generated_method.average_size"`
	SparkDriverCodeGeneratorGeneratedMethodCount            MetricConfig `mapstructure:"spark.driver.code_generator.generated_method.count"`
	SparkDriverCodeGeneratorSourceCodeAverageSize           MetricConfig `mapstructure:"spark.driver.code_generator.source_code.average_size"`
	SparkDriverCodeGeneratorSourceCodeCount                 MetricConfig `mapstructure:"spark.driver.code_generator.source_code.count"`
	SparkDriverDagSchedulerJobsActive                       MetricConfig `mapstructure:"spark.driver.dag_scheduler.jobs.active"`
	SparkDriverDagSchedulerJobsAll                          MetricConfig `mapstructure:"spark.driver.dag_scheduler.jobs.all"`
	SparkDriverDagSchedulerStagesFailed                     MetricConfig `mapstructure:"spark.driver.dag_scheduler.stages.failed"`
	SparkDriverDagSchedulerStagesRunning                    MetricConfig `mapstructure:"spark.driver.dag_scheduler.stages.running"`
	SparkDriverDagSchedulerStagesWaiting                    MetricConfig `mapstructure:"spark.driver.dag_scheduler.stages.waiting"`
	SparkDriverExecutorMetricsGcCount                       MetricConfig `mapstructure:"spark.driver.executor_metrics.gc.count"`
	SparkDriverExecutorMetricsGcTime                        MetricConfig `mapstructure:"spark.driver.executor_metrics.gc.time"`
	SparkDriverExecutorMetricsJvmMemory                     MetricConfig `mapstructure:"spark.driver.executor_metrics.jvm_memory"`
	SparkDriverExecutorMetricsMemoryExecution               MetricConfig `mapstructure:"spark.driver.executor_metrics.memory.execution"`
	SparkDriverExecutorMetricsMemoryPool                    MetricConfig `mapstructure:"spark.driver.executor_metrics.memory.pool"`
	SparkDriverExecutorMetricsMemoryStorage                 MetricConfig `mapstructure:"spark.driver.executor_metrics.memory.storage"`
	SparkDriverHiveExternalCatalogFileCacheHits             MetricConfig `mapstructure:"spark.driver.hive_external_catalog.file_cache_hits"`
	SparkDriverHiveExternalCatalogFilesDiscovered           MetricConfig `mapstructure:"spark.driver.hive_external_catalog.files_discovered"`
	SparkDriverHiveExternalCatalogHiveClientCalls           MetricConfig `mapstructure:"spark.driver.hive_external_catalog.hive_client_calls"`
	SparkDriverHiveExternalCatalogParallelListingJobs       MetricConfig `mapstructure:"spark.driver.hive_external_catalog.parallel_listing_jobs"`
	SparkDriverHiveExternalCatalogPartitionsFetched         MetricConfig `mapstructure:"spark.driver.hive_external_catalog.partitions_fetched"`
	SparkDriverJvmCPUTime                                   MetricConfig `mapstructure:"spark.driver.jvm_cpu_time"`
	SparkDriverLiveListenerBusEventsDropped                 MetricConfig `mapstructure:"spark.driver.live_listener_bus.events_dropped"`
	SparkDriverLiveListenerBusEventsPosted                  MetricConfig `mapstructure:"spark.driver.live_listener_bus.events_posted"`
	SparkDriverLiveListenerBusListenerProcessingTimeAverage MetricConfig `mapstructure:"spark.driver.live_listener_bus.listener_processing_time.average"`
	SparkDriverLiveListenerBusQueueSize                     MetricConfig `mapstructure:"spark.driver.live_listener_bus.queue_size"`
	SparkExecutorDiskUsed                                   MetricConfig `mapstructure:"spark.executor.disk_used"`
	SparkExecutorDuration                                   MetricConfig `mapstructure:"spark.executor.duration"`
	SparkExecutorGcTime                                     MetricConfig `mapstructure:"spark.executor.gc_time"`
	SparkExecutorInputSize                                  MetricConfig `mapstructure:"spark.executor.input_size"`
	SparkExecutorMemoryUsed                                 MetricConfig `mapstructure:"spark.executor.memory_used"`
	SparkExecutorShuffleIoSize                              MetricConfig `mapstructure:"spark.executor.shuffle.io.size"`
	SparkExecutorStorageMemoryTotal                         MetricConfig `mapstructure:"spark.executor.storage_memory.total"`
	SparkExecutorStorageMemoryUsed                          MetricConfig `mapstructure:"spark.executor.storage_memory.used"`
	SparkExecutorTasksActive                                MetricConfig `mapstructure:"spark.executor.tasks.active"`
	SparkExecutorTasksMax                                   MetricConfig `mapstructure:"spark.executor.tasks.max"`
	SparkExecutorTasksResults                               MetricConfig `mapstructure:"spark.executor.tasks.results"`
	SparkJobStagesActive                                    MetricConfig `mapstructure:"spark.job.stages.active"`
	SparkJobStagesResults                                   MetricConfig `mapstructure:"spark.job.stages.results"`
	SparkJobTasksActive                                     MetricConfig `mapstructure:"spark.job.tasks.active"`
	SparkJobTasksResults                                    MetricConfig `mapstructure:"spark.job.tasks.results"`
	SparkStageDiskSpaceSpilled                              MetricConfig `mapstructure:"spark.stage.disk_space_spilled"`
	SparkStageExecutorCPUTime                               MetricConfig `mapstructure:"spark.stage.executor.cpu_time"`
	SparkStageExecutorRunTime                               MetricConfig `mapstructure:"spark.stage.executor.run_time"`
	SparkStageIoRecords                                     MetricConfig `mapstructure:"spark.stage.io.records"`
	SparkStageIoSize                                        MetricConfig `mapstructure:"spark.stage.io.size"`
	SparkStageJvmGcTime                                     MetricConfig `mapstructure:"spark.stage.jvm_gc_time"`
	SparkStageMemorySpilled                                 MetricConfig `mapstructure:"spark.stage.memory_spilled"`
	SparkStagePeakExecutionMemory                           MetricConfig `mapstructure:"spark.stage.peak_execution_memory"`
	SparkStageShuffleBlocksFetched                          MetricConfig `mapstructure:"spark.stage.shuffle.blocks_fetched"`
	SparkStageShuffleFetchWaitTime                          MetricConfig `mapstructure:"spark.stage.shuffle.fetch_wait_time"`
	SparkStageShuffleIoRecords                              MetricConfig `mapstructure:"spark.stage.shuffle.io.records"`
	SparkStageShuffleIoSize                                 MetricConfig `mapstructure:"spark.stage.shuffle.io.size"`
	SparkStageShuffleRemoteDataReadToDisk                   MetricConfig `mapstructure:"spark.stage.shuffle.remote_data_read_to_disk"`
	SparkStageShuffleWriteTime                              MetricConfig `mapstructure:"spark.stage.shuffle.write_time"`
	SparkStageTaskActive                                    MetricConfig `mapstructure:"spark.stage.task.active"`
	SparkStageTaskResultSize                                MetricConfig `mapstructure:"spark.stage.task.result_size"`
	SparkStageTaskResults                                   MetricConfig `mapstructure:"spark.stage.task.results"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		SparkDriverBlockManagerDiskSpaceUsed: MetricConfig{
			Enabled: true,
		},
		SparkDriverBlockManagerMemoryRemaining: MetricConfig{
			Enabled: true,
		},
		SparkDriverBlockManagerMemoryUsed: MetricConfig{
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
		SparkDriverCodeGeneratorSourceCodeCount: MetricConfig{
			Enabled: true,
		},
		SparkDriverDagSchedulerJobsActive: MetricConfig{
			Enabled: true,
		},
		SparkDriverDagSchedulerJobsAll: MetricConfig{
			Enabled: true,
		},
		SparkDriverDagSchedulerStagesFailed: MetricConfig{
			Enabled: true,
		},
		SparkDriverDagSchedulerStagesRunning: MetricConfig{
			Enabled: true,
		},
		SparkDriverDagSchedulerStagesWaiting: MetricConfig{
			Enabled: true,
		},
		SparkDriverExecutorMetricsGcCount: MetricConfig{
			Enabled: true,
		},
		SparkDriverExecutorMetricsGcTime: MetricConfig{
			Enabled: true,
		},
		SparkDriverExecutorMetricsJvmMemory: MetricConfig{
			Enabled: true,
		},
		SparkDriverExecutorMetricsMemoryExecution: MetricConfig{
			Enabled: true,
		},
		SparkDriverExecutorMetricsMemoryPool: MetricConfig{
			Enabled: true,
		},
		SparkDriverExecutorMetricsMemoryStorage: MetricConfig{
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
		SparkDriverLiveListenerBusEventsDropped: MetricConfig{
			Enabled: true,
		},
		SparkDriverLiveListenerBusEventsPosted: MetricConfig{
			Enabled: true,
		},
		SparkDriverLiveListenerBusListenerProcessingTimeAverage: MetricConfig{
			Enabled: true,
		},
		SparkDriverLiveListenerBusQueueSize: MetricConfig{
			Enabled: true,
		},
		SparkExecutorDiskUsed: MetricConfig{
			Enabled: true,
		},
		SparkExecutorDuration: MetricConfig{
			Enabled: true,
		},
		SparkExecutorGcTime: MetricConfig{
			Enabled: true,
		},
		SparkExecutorInputSize: MetricConfig{
			Enabled: true,
		},
		SparkExecutorMemoryUsed: MetricConfig{
			Enabled: true,
		},
		SparkExecutorShuffleIoSize: MetricConfig{
			Enabled: true,
		},
		SparkExecutorStorageMemoryTotal: MetricConfig{
			Enabled: true,
		},
		SparkExecutorStorageMemoryUsed: MetricConfig{
			Enabled: true,
		},
		SparkExecutorTasksActive: MetricConfig{
			Enabled: true,
		},
		SparkExecutorTasksMax: MetricConfig{
			Enabled: true,
		},
		SparkExecutorTasksResults: MetricConfig{
			Enabled: true,
		},
		SparkJobStagesActive: MetricConfig{
			Enabled: true,
		},
		SparkJobStagesResults: MetricConfig{
			Enabled: true,
		},
		SparkJobTasksActive: MetricConfig{
			Enabled: true,
		},
		SparkJobTasksResults: MetricConfig{
			Enabled: true,
		},
		SparkStageDiskSpaceSpilled: MetricConfig{
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
		SparkStageMemorySpilled: MetricConfig{
			Enabled: true,
		},
		SparkStagePeakExecutionMemory: MetricConfig{
			Enabled: true,
		},
		SparkStageShuffleBlocksFetched: MetricConfig{
			Enabled: true,
		},
		SparkStageShuffleFetchWaitTime: MetricConfig{
			Enabled: true,
		},
		SparkStageShuffleIoRecords: MetricConfig{
			Enabled: true,
		},
		SparkStageShuffleIoSize: MetricConfig{
			Enabled: true,
		},
		SparkStageShuffleRemoteDataReadToDisk: MetricConfig{
			Enabled: true,
		},
		SparkStageShuffleWriteTime: MetricConfig{
			Enabled: true,
		},
		SparkStageTaskActive: MetricConfig{
			Enabled: true,
		},
		SparkStageTaskResultSize: MetricConfig{
			Enabled: true,
		},
		SparkStageTaskResults: MetricConfig{
			Enabled: true,
		},
	}
}

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool `mapstructure:"enabled"`
}

// ResourceAttributesConfig provides config for apachespark resource attributes.
type ResourceAttributesConfig struct {
	SparkApplicationID   ResourceAttributeConfig `mapstructure:"spark.application.id"`
	SparkApplicationName ResourceAttributeConfig `mapstructure:"spark.application.name"`
	SparkExecutorID      ResourceAttributeConfig `mapstructure:"spark.executor.id"`
	SparkJobID           ResourceAttributeConfig `mapstructure:"spark.job.id"`
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
