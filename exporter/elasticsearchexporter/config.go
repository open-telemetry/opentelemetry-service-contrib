// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package elasticsearchexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter"

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"go.opentelemetry.io/collector/config/configcompression"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/exporter/exporterbatcher"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.uber.org/zap"
)

// Config defines configuration for Elastic exporter.
type Config struct {
	QueueSettings exporterhelper.QueueConfig `mapstructure:"sending_queue"`
	// Endpoints holds the Elasticsearch URLs the exporter should send events to.
	//
	// This setting is required if CloudID is not set and if the
	// ELASTICSEARCH_URL environment variable is not set.
	Endpoints []string `mapstructure:"endpoints"`

	// CloudID holds the cloud ID to identify the Elastic Cloud cluster to send events to.
	// https://www.elastic.co/guide/en/cloud/current/ec-cloud-id.html
	//
	// This setting is required if no URL is configured.
	CloudID string `mapstructure:"cloudid"`

	// NumWorkers configures the number of workers publishing bulk requests.
	NumWorkers int `mapstructure:"num_workers"`

	// This setting is required when logging pipelines used.
	LogsIndex string `mapstructure:"logs_index"`
	// fall back to pure LogsIndex, if 'elasticsearch.index.prefix' or 'elasticsearch.index.suffix' are not found in resource or attribute (prio: resource > attribute)
	LogsDynamicIndex DynamicIndexSetting `mapstructure:"logs_dynamic_index"`

	// This setting is required when the exporter is used in a metrics pipeline.
	MetricsIndex string `mapstructure:"metrics_index"`
	// fall back to pure MetricsIndex, if 'elasticsearch.index.prefix' or 'elasticsearch.index.suffix' are not found in resource attributes
	MetricsDynamicIndex DynamicIndexSetting `mapstructure:"metrics_dynamic_index"`

	// This setting is required when traces pipelines used.
	TracesIndex string `mapstructure:"traces_index"`
	// fall back to pure TracesIndex, if 'elasticsearch.index.prefix' or 'elasticsearch.index.suffix' are not found in resource or attribute (prio: resource > attribute)
	TracesDynamicIndex DynamicIndexSetting `mapstructure:"traces_dynamic_index"`

	// LogsDynamicID configures whether log record attribute `elasticsearch.document_id` is set as the document ID in ES.
	LogsDynamicID DynamicIDSettings `mapstructure:"logs_dynamic_id"`

	// Pipeline configures the ingest node pipeline name that should be used to process the
	// events.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ingest.html
	Pipeline string `mapstructure:"pipeline"`

	confighttp.ClientConfig `mapstructure:",squash"`
	Authentication          AuthenticationSettings `mapstructure:",squash"`
	Discovery               DiscoverySettings      `mapstructure:"discover"`
	Retry                   RetrySettings          `mapstructure:"retry"`
	Flush                   FlushSettings          `mapstructure:"flush"`
	Mapping                 MappingsSettings       `mapstructure:"mapping"`
	LogstashFormat          LogstashFormatSettings `mapstructure:"logstash_format"`

	// TelemetrySettings contains settings useful for testing/debugging purposes
	// This is experimental and may change at any time.
	TelemetrySettings `mapstructure:"telemetry"`

	// Batcher holds configuration for batching requests based on timeout
	// and size-based thresholds.
	//
	// Batcher is unused by default, in which case Flush will be used.
	// If Batcher.Enabled is non-nil (i.e. batcher::enabled is specified),
	// then the Flush will be ignored even if Batcher.Enabled is false.
	Batcher BatcherConfig `mapstructure:"batcher"`
}

// BatcherConfig holds configuration for exporterbatcher.
//
// This is a slightly modified version of exporterbatcher.Config,
// to enable tri-state Enabled: unset, false, true.
type BatcherConfig struct {
	// Enabled indicates whether to enqueue batches before sending
	// to the exporter. If Enabled is specified (non-nil),
	// then the exporter will not perform any buffering itself.
	Enabled *bool `mapstructure:"enabled"`

	// FlushTimeout sets the time after which a batch will be sent regardless of its size.
	FlushTimeout time.Duration `mapstructure:"flush_timeout"`

	exporterbatcher.MinSizeConfig `mapstructure:",squash"`
	exporterbatcher.MaxSizeConfig `mapstructure:",squash"`
}

type TelemetrySettings struct {
	LogRequestBody  bool `mapstructure:"log_request_body"`
	LogResponseBody bool `mapstructure:"log_response_body"`
}

type LogstashFormatSettings struct {
	Enabled         bool   `mapstructure:"enabled"`
	PrefixSeparator string `mapstructure:"prefix_separator"`
	DateFormat      string `mapstructure:"date_format"`
}

type DynamicIndexSetting struct {
	Enabled bool `mapstructure:"enabled"`
}

type DynamicIDSettings struct {
	Enabled bool `mapstructure:"enabled"`
}

// AuthenticationSettings defines user authentication related settings.
type AuthenticationSettings struct {
	// User is used to configure HTTP Basic Authentication.
	User string `mapstructure:"user"`

	// Password is used to configure HTTP Basic Authentication.
	Password configopaque.String `mapstructure:"password"`

	// APIKey is used to configure ApiKey based Authentication.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-create-api-key.html
	APIKey configopaque.String `mapstructure:"api_key"`
}

// DiscoverySettings defines Elasticsearch node discovery related settings.
// The exporter will check Elasticsearch regularly for available nodes
// and updates the list of hosts if discovery is enabled. Newly discovered
// nodes will automatically be used for load balancing.
//
// DiscoverySettings should not be enabled when operating Elasticsearch behind a proxy
// or load balancer.
//
// https://www.elastic.co/blog/elasticsearch-sniffing-best-practices-what-when-why-how
type DiscoverySettings struct {
	// OnStart, if set, instructs the exporter to look for available Elasticsearch
	// nodes the first time the exporter connects to the cluster.
	OnStart bool `mapstructure:"on_start"`

	// Interval instructs the exporter to renew the list of Elasticsearch URLs
	// with the given interval. URLs will not be updated if Interval is <=0.
	Interval time.Duration `mapstructure:"interval"`
}

// FlushSettings defines settings for configuring the write buffer flushing
// policy in the Elasticsearch exporter. The exporter sends a bulk request with
// all events already serialized into the send-buffer.
type FlushSettings struct {
	// Bytes sets the send buffer flushing limit.
	Bytes int `mapstructure:"bytes"`

	// Interval configures the max age of a document in the send buffer.
	Interval time.Duration `mapstructure:"interval"`
}

// RetrySettings defines settings for the HTTP request retries in the Elasticsearch exporter.
// Failed sends are retried with exponential backoff.
type RetrySettings struct {
	// Enabled allows users to disable retry without having to comment out all settings.
	Enabled bool `mapstructure:"enabled"`

	// MaxRequests configures how often an HTTP request is attempted before it is assumed to be failed.
	// Deprecated: use MaxRetries instead.
	MaxRequests int `mapstructure:"max_requests"`

	// MaxRetries configures how many times an HTTP request is retried.
	MaxRetries int `mapstructure:"max_retries"`

	// InitialInterval configures the initial waiting time if a request failed.
	InitialInterval time.Duration `mapstructure:"initial_interval"`

	// MaxInterval configures the max waiting time if consecutive requests failed.
	MaxInterval time.Duration `mapstructure:"max_interval"`

	// RetryOnStatus configures the status codes that trigger request or document level retries.
	RetryOnStatus []int `mapstructure:"retry_on_status"`
}

type MappingsSettings struct {
	// Mode configures the field mappings.
	Mode string `mapstructure:"mode"`
}

type MappingMode int

// Enum values for MappingMode.
const (
	MappingNone MappingMode = iota
	MappingECS
	MappingOTel
	MappingRaw
	MappingBodyMap
)

var (
	errConfigEndpointRequired = errors.New("exactly one of [endpoint, endpoints, cloudid] must be specified")
	errConfigEmptyEndpoint    = errors.New("endpoint must not be empty")
)

func (m MappingMode) String() string {
	switch m {
	case MappingNone:
		return ""
	case MappingECS:
		return "ecs"
	case MappingOTel:
		return "otel"
	case MappingRaw:
		return "raw"
	case MappingBodyMap:
		return "bodymap"
	default:
		return ""
	}
}

var mappingModes = func() map[string]MappingMode {
	table := map[string]MappingMode{}
	for _, m := range []MappingMode{
		MappingNone,
		MappingECS,
		MappingOTel,
		MappingRaw,
		MappingBodyMap,
	} {
		table[strings.ToLower(m.String())] = m
	}

	// config aliases
	table["no"] = MappingNone
	table["none"] = MappingNone

	return table
}()

const defaultElasticsearchEnvName = "ELASTICSEARCH_URL"

// Validate validates the elasticsearch server configuration.
func (cfg *Config) Validate() error {
	endpoints, err := cfg.endpoints()
	if err != nil {
		return err
	}
	for _, endpoint := range endpoints {
		if err := validateEndpoint(endpoint); err != nil {
			return fmt.Errorf("invalid endpoint %q: %w", endpoint, err)
		}
	}

	if _, ok := mappingModes[cfg.Mapping.Mode]; !ok {
		return fmt.Errorf("unknown mapping mode %q", cfg.Mapping.Mode)
	}

	if cfg.Compression != "none" && cfg.Compression != configcompression.TypeGzip {
		return errors.New("compression must be one of [none, gzip]")
	}

	if cfg.Retry.MaxRequests != 0 && cfg.Retry.MaxRetries != 0 {
		return errors.New("must not specify both retry::max_requests and retry::max_retries")
	}
	if cfg.Retry.MaxRequests < 0 {
		return errors.New("retry::max_requests should be non-negative")
	}
	if cfg.Retry.MaxRetries < 0 {
		return errors.New("retry::max_retries should be non-negative")
	}

	return nil
}

func (cfg *Config) endpoints() ([]string, error) {
	// Exactly one of endpoint, endpoints, or cloudid must be configured.
	// If none are set, then $ELASTICSEARCH_URL may be specified instead.
	var endpoints []string
	var numEndpointConfigs int
	if cfg.Endpoint != "" {
		numEndpointConfigs++
		endpoints = []string{cfg.Endpoint}
	}
	if len(cfg.Endpoints) > 0 {
		numEndpointConfigs++
		endpoints = cfg.Endpoints
	}
	if cfg.CloudID != "" {
		numEndpointConfigs++
		u, err := parseCloudID(cfg.CloudID)
		if err != nil {
			return nil, err
		}
		endpoints = []string{u.String()}
	}
	if numEndpointConfigs == 0 {
		if v := os.Getenv(defaultElasticsearchEnvName); v != "" {
			numEndpointConfigs++
			endpoints = strings.Split(v, ",")
			for i, endpoint := range endpoints {
				endpoints[i] = strings.TrimSpace(endpoint)
			}
		}
	}
	if numEndpointConfigs != 1 {
		return nil, errConfigEndpointRequired
	}
	return endpoints, nil
}

func validateEndpoint(endpoint string) error {
	if endpoint == "" {
		return errConfigEmptyEndpoint
	}

	u, err := url.Parse(endpoint)
	if err != nil {
		return err
	}
	switch u.Scheme {
	case "http", "https":
	default:
		return fmt.Errorf(`invalid scheme %q, expected "http" or "https"`, u.Scheme)
	}
	return nil
}

// Based on "addrFromCloudID" in go-elasticsearch.
func parseCloudID(input string) (*url.URL, error) {
	_, after, ok := strings.Cut(input, ":")
	if !ok {
		return nil, fmt.Errorf("invalid CloudID %q", input)
	}

	decoded, err := base64.StdEncoding.DecodeString(after)
	if err != nil {
		return nil, err
	}

	before, after, ok := strings.Cut(string(decoded), "$")
	if !ok {
		return nil, fmt.Errorf("invalid decoded CloudID %q", string(decoded))
	}
	return url.Parse(fmt.Sprintf("https://%s.%s", after, before))
}

// MappingMode returns the mapping.mode defined in the given cfg
// object. This method must be called after cfg.Validate() has been
// called without returning an error.
func (cfg *Config) MappingMode() MappingMode {
	return mappingModes[cfg.Mapping.Mode]
}

func handleDeprecatedConfig(cfg *Config, logger *zap.Logger) {
	if cfg.Retry.MaxRequests != 0 {
		cfg.Retry.MaxRetries = cfg.Retry.MaxRequests - 1
		// Do not set cfg.Retry.Enabled = false if cfg.Retry.MaxRequest = 1 to avoid breaking change on behavior
		logger.Warn("retry::max_requests has been deprecated, and will be removed in a future version. Use retry::max_retries instead.")
	}
}
