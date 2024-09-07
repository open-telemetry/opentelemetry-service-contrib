// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datadog // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/datadog"

import "go.opentelemetry.io/collector/config/confignet"

// TracesConfig defines the traces exporter specific configuration options
type TracesConfig struct {
	// TCPAddr.Endpoint is the host of the Datadog intake server to send traces to.
	// If unset, the value is obtained from the Site.
	confignet.TCPAddrConfig `mapstructure:",squash"`

	// ignored resources
	// A blacklist of regular expressions can be provided to disable certain traces based on their resource name
	// all entries must be surrounded by double quotes and separated by commas.
	// ignore_resources: ["(GET|POST) /healthcheck"]
	IgnoreResources []string `mapstructure:"ignore_resources"`

	// SpanNameRemappings is the map of datadog span names and preferred name to map to. This can be used to
	// automatically map Datadog Span Operation Names to an updated value. All entries should be key/value pairs.
	// span_name_remappings:
	//   io.opentelemetry.javaagent.spring.client: spring.client
	//   instrumentation:express.server: express
	//   go.opentelemetry.io_contrib_instrumentation_net_http_otelhttp.client: http.client
	SpanNameRemappings map[string]string `mapstructure:"span_name_remappings"`

	// If set to true the OpenTelemetry span name will used in the Datadog resource name.
	// If set to false the resource name will be filled with the instrumentation library name + span kind.
	// The default value is `false`.
	SpanNameAsResourceName bool `mapstructure:"span_name_as_resource_name"`

	// If set to true, enables an additional stats computation check on spans to see they have an eligible `span.kind` (server, consumer, client, producer).
	// If enabled, a span with an eligible `span.kind` will have stats computed. If disabled, only top-level and measured spans will have stats computed.
	// NOTE: For stats computed from OTel traces, only top-level spans are considered when this option is off.
	// If you are sending OTel traces and want stats on non-top-level spans, this flag will need to be enabled.
	// If you are sending OTel traces and do not want stats computed by span kind, you need to disable this flag and disable `compute_top_level_by_span_kind`.
	ComputeStatsBySpanKind bool `mapstructure:"compute_stats_by_span_kind"`

	// If set to true, root spans and spans with a server or consumer `span.kind` will be marked as top-level.
	// Additionally, spans with a client or producer `span.kind` will have stats computed.
	// Enabling this config option may increase the number of spans that generate trace metrics, and may change which spans appear as top-level in Datadog.
	// ComputeTopLevelBySpanKind needs to be enabled in both the Datadog connector and Datadog exporter configs if both components are being used.
	// The default value is `false`.
	ComputeTopLevelBySpanKind bool `mapstructure:"compute_top_level_by_span_kind"`

	// If set to true, enables `peer.service` aggregation in the exporter. If disabled, aggregated trace stats will not include `peer.service` as a dimension.
	// For the best experience with `peer.service`, it is recommended to also enable `compute_stats_by_span_kind`.
	// If enabling both causes the datadog exporter to consume too many resources, try disabling `compute_stats_by_span_kind` first.
	// If the overhead remains high, it will be due to a high cardinality of `peer.service` values from the traces. You may need to check your instrumentation.
	// Deprecated: Please use PeerTagsAggregation instead
	PeerServiceAggregation bool `mapstructure:"peer_service_aggregation"`

	// If set to true, enables aggregation of peer related tags (e.g., `peer.service`, `db.instance`, etc.) in the datadog exporter.
	// If disabled, aggregated trace stats will not include these tags as dimensions on trace metrics.
	// For the best experience with peer tags, Datadog also recommends enabling `compute_stats_by_span_kind`.
	// If you are using an OTel tracer, it's best to have both enabled because client/producer spans with relevant peer tags
	// may not be marked by the datadog exporter as top-level spans.
	// If enabling both causes the datadog exporter to consume too many resources, try disabling `compute_stats_by_span_kind` first.
	// A high cardinality of peer tags or APM resources can also contribute to higher CPU and memory consumption.
	// You can check for the cardinality of these fields by making trace search queries in the Datadog UI.
	// The default list of peer tags can be found in https://github.com/DataDog/datadog-agent/blob/main/pkg/trace/stats/concentrator.go.
	PeerTagsAggregation bool `mapstructure:"peer_tags_aggregation"`

	// [BETA] Optional list of supplementary peer tags that go beyond the defaults. The Datadog backend validates all tags
	// and will drop ones that are unapproved. The default set of peer tags can be found at
	// https://github.com/DataDog/datadog-agent/blob/505170c4ac8c3cbff1a61cf5f84b28d835c91058/pkg/trace/stats/concentrator.go#L55.
	PeerTags []string `mapstructure:"peer_tags"`

	// TraceBuffer specifies the number of Datadog Agent TracerPayloads to buffer before dropping.
	// The default value is 0, meaning the Datadog Agent TracerPayloads are unbuffered.
	TraceBuffer int `mapstructure:"trace_buffer"`

	// FlushInterval defines the interval in seconds at which the writer flushes traces
	// to the intake; used in tests.
	FlushInterval float64 `mapstructure:"flush_interval"`
}
