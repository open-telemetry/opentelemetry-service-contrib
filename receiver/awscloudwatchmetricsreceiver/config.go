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

package awscloudwatchmetricsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscloudwatchmetricsreceiver"

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"go.uber.org/multierr"
)

var (
	defaultPollInterval = time.Minute
)

// Config is the overall config structure for the awscloudwatchmetricsreceiver
type Config struct {
	Region       string         `mapstructure:"region"`
	Profile      string         `mapstructure:"profile"`
	IMDSEndpoint string         `mapstructure:"imds_endpoint"`
	PollInterval time.Duration  `mapstructure:"poll_interval"`
	Metrics      *MetricsConfig `mapstructure:"metrics"`
}

// MetricsConfig is the configuration for the metrics part of the receiver
// added this so we could expand to other inputs such as autodiscover
type MetricsConfig struct {
	Names []*NamedConfig `mapstructure:"named"`
}

// NamesConfig is the configuration for the metric namespace and metric names
// https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html
type NamedConfig struct {
	Namespace      string                   `mapstructure:"namespace"`
	MetricName     string                   `mapstructure:"metric_name"`
	Period         time.Duration            `mapstructure:"period"`
	AwsAggregation string                   `mapstructure:"aws_aggregation"`
	Dimensions     []MetricDimensionsConfig `mapstructure:"dimensions"`
}

// MetricDimensionConfig is the configuration for the metric dimensions
type MetricDimensionsConfig struct {
	Name  string `mapstructure:"Name"`
	Value string `mapstructure:"Value"`
}

var (
	errNoMetricsConfigured = errors.New("no metrics configured")
	errNoRegion            = errors.New("no region was specified")
	errInvalidPollInterval = errors.New("poll interval is incorrect, it must be a duration greater than one second")

	// https://docs.aws.amazon.com/cli/latest/reference/cloudwatch/get-metric-data.html
	// GetMetricData supports up to 500 metrics per API call
	errTooManyMetrics = errors.New("too many metrics defined")

	// https://docs.aws.amazon.com/cli/latest/reference/cloudwatch/get-metric-data.html
	errEmptyDimensions      = errors.New("dimensions name and value is empty")
	errTooManyDimensions    = errors.New("you cannot define more than 30 dimensions for a metric")
	errDimensionColonPrefix = errors.New("dimension name cannot start with a colon")

	errInvalidAwsAggregation = errors.New("invalid AWS aggregation")
)

func (cfg *Config) Validate() error {
	if cfg.Region == "" {
		return errNoRegion
	}

	if cfg.IMDSEndpoint != "" {
		_, err := url.ParseRequestURI(cfg.IMDSEndpoint)
		if err != nil {
			return fmt.Errorf("unable to parse URI for imds_endpoint: %w", err)
		}
	}

	if cfg.PollInterval < time.Second {
		return errInvalidPollInterval
	}
	var errs error
	errs = multierr.Append(errs, cfg.validateMetricsConfig())
	return errs
}

func (cfg *Config) validateMetricsConfig() error {
	if cfg.Metrics == nil {
		return errNoMetricsConfigured
	}
	return cfg.validateNamedConfig()
}

func (cfg *Config) validateNamedConfig() error {
	if cfg.Metrics.Names == nil {
		return errNoMetricsConfigured
	}
	return cfg.validateDimensionsConfig()
}

func (cfg *Config) validateDimensionsConfig() error {
	var errs error

	metricsNames := cfg.Metrics.Names
	if len(metricsNames) > 500 {
		return errTooManyMetrics
	}
	for _, name := range metricsNames {
		if name.Namespace == "" {
			return errNoMetricsConfigured
		}
		err := validateAwsAggregation(name.AwsAggregation)
		if err != nil {
			return err
		}
		if name.MetricName == "" {
			return errNoMetricsConfigured
		}
		errs = multierr.Append(errs, validate(name.Dimensions))
	}
	return errs
}

// https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html
func validateAwsAggregation(agg string) error {
	switch {
	case agg == "SampleCount":
		return nil
	case agg == "Sum":
		return nil
	case agg == "Average":
		return nil
	case agg == "Minimum":
		return nil
	case agg == "Maximum":
		return nil
	case strings.HasPrefix(agg, "p"):
		return nil
	case strings.HasPrefix(agg, "TM"):
		return nil
	case agg == "IQM":
		return nil
	case strings.HasPrefix(agg, "PR"):
		return nil
	case strings.HasPrefix(agg, "TC"):
		return nil
	case strings.HasPrefix(agg, "TS"):
		return nil
	default:
		return errInvalidAwsAggregation
	}
}

func validate(nmd []MetricDimensionsConfig) error {
	for _, dimensionConfig := range nmd {
		if dimensionConfig.Name == "" || dimensionConfig.Value == "" {
			return errEmptyDimensions
		}
		if strings.HasPrefix(dimensionConfig.Name, ":") {
			return errDimensionColonPrefix
		}
	}
	if len(nmd) > 30 {
		return errTooManyDimensions
	}
	return nil
}
