// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sattributesprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/consumerprofiles"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/processorhelper"
	"go.opentelemetry.io/collector/processor/processorhelper/processorhelperprofiles"
	"go.opentelemetry.io/collector/processor/processorprofiles"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor/internal/kube"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor/internal/metadata"
)

var kubeClientProvider = kube.ClientProvider(nil)
var consumerCapabilities = consumer.Capabilities{MutatesData: true}
var defaultExcludes = ExcludeConfig{Pods: []ExcludePodConfig{{Name: "jaeger-agent"}, {Name: "jaeger-collector"}}}

// NewFactory returns a new factory for the k8s processor.
func NewFactory() processor.Factory {
	return processorprofiles.NewFactory(
		metadata.Type,
		createDefaultConfig,
		processorprofiles.WithTraces(createTracesProcessor, metadata.TracesStability),
		processorprofiles.WithMetrics(createMetricsProcessor, metadata.MetricsStability),
		processorprofiles.WithLogs(createLogsProcessor, metadata.LogsStability),
		processorprofiles.WithProfiles(createProfilesProcessor, metadata.ProfilesStability),
	)
}

func createDefaultConfig() component.Config {
	return &Config{
		APIConfig: k8sconfig.APIConfig{AuthType: k8sconfig.AuthTypeServiceAccount},
		Exclude:   defaultExcludes,
		Extract: ExtractConfig{
			Metadata: enabledAttributes(),
		},
	}
}

func createTracesProcessor(
	ctx context.Context,
	params processor.Settings,
	cfg component.Config,
	next consumer.Traces,
) (processor.Traces, error) {
	return createTracesProcessorWithOptions(ctx, params, cfg, next)
}

func createLogsProcessor(
	ctx context.Context,
	params processor.Settings,
	cfg component.Config,
	nextLogsConsumer consumer.Logs,
) (processor.Logs, error) {
	return createLogsProcessorWithOptions(ctx, params, cfg, nextLogsConsumer)
}

func createMetricsProcessor(
	ctx context.Context,
	params processor.Settings,
	cfg component.Config,
	nextMetricsConsumer consumer.Metrics,
) (processor.Metrics, error) {
	return createMetricsProcessorWithOptions(ctx, params, cfg, nextMetricsConsumer)
}

func createProfilesProcessor(
	ctx context.Context,
	params processor.Settings,
	cfg component.Config,
	nextProfilesConsumer consumerprofiles.Profiles,
) (processorprofiles.Profiles, error) {
	return createProfilesProcessorWithOptions(ctx, params, cfg, nextProfilesConsumer)
}

func createTracesProcessorWithOptions(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	next consumer.Traces,
	options ...option,
) (processor.Traces, error) {
	kp := createKubernetesProcessor(set, cfg, options...)

	return processorhelper.NewTraces(
		ctx,
		set,
		cfg,
		next,
		kp.processTraces,
		processorhelper.WithCapabilities(consumerCapabilities),
		processorhelper.WithStart(kp.Start),
		processorhelper.WithShutdown(kp.Shutdown))
}

func createMetricsProcessorWithOptions(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	nextMetricsConsumer consumer.Metrics,
	options ...option,
) (processor.Metrics, error) {
	kp := createKubernetesProcessor(set, cfg, options...)

	return processorhelper.NewMetrics(
		ctx,
		set,
		cfg,
		nextMetricsConsumer,
		kp.processMetrics,
		processorhelper.WithCapabilities(consumerCapabilities),
		processorhelper.WithStart(kp.Start),
		processorhelper.WithShutdown(kp.Shutdown))
}

func createLogsProcessorWithOptions(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	nextLogsConsumer consumer.Logs,
	options ...option,
) (processor.Logs, error) {
	kp := createKubernetesProcessor(set, cfg, options...)

	return processorhelper.NewLogs(
		ctx,
		set,
		cfg,
		nextLogsConsumer,
		kp.processLogs,
		processorhelper.WithCapabilities(consumerCapabilities),
		processorhelper.WithStart(kp.Start),
		processorhelper.WithShutdown(kp.Shutdown))
}

func createProfilesProcessorWithOptions(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	nextProfilesConsumer consumerprofiles.Profiles,
	options ...option,
) (processorprofiles.Profiles, error) {
	kp := createKubernetesProcessor(set, cfg, options...)

	return processorhelperprofiles.NewProfiles(
		ctx,
		set,
		cfg,
		nextProfilesConsumer,
		kp.processProfiles,
		processorhelperprofiles.WithCapabilities(consumerCapabilities),
		processorhelperprofiles.WithStart(kp.Start),
		processorhelperprofiles.WithShutdown(kp.Shutdown),
	)
}

func createKubernetesProcessor(
	params processor.Settings,
	cfg component.Config,
	options ...option,
) *kubernetesprocessor {
	kp := &kubernetesprocessor{logger: params.Logger,
		cfg:               cfg,
		options:           options,
		telemetrySettings: params.TelemetrySettings,
	}

	return kp
}

func createProcessorOpts(cfg component.Config) []option {
	oCfg := cfg.(*Config)
	var opts []option
	if oCfg.Passthrough {
		opts = append(opts, withPassthrough())
	}

	// extraction rules
	opts = append(opts, withExtractMetadata(oCfg.Extract.Metadata...))
	opts = append(opts, withExtractLabels(oCfg.Extract.Labels...))
	opts = append(opts, withExtractAnnotations(oCfg.Extract.Annotations...))

	// filters
	opts = append(opts, withFilterNode(oCfg.Filter.Node, oCfg.Filter.NodeFromEnvVar))
	opts = append(opts, withFilterNamespace(oCfg.Filter.Namespace))
	opts = append(opts, withFilterLabels(oCfg.Filter.Labels...))
	opts = append(opts, withFilterFields(oCfg.Filter.Fields...))
	opts = append(opts, withAPIConfig(oCfg.APIConfig))

	opts = append(opts, withExtractPodAssociations(oCfg.Association...))

	opts = append(opts, withExcludes(oCfg.Exclude))

	return opts
}
