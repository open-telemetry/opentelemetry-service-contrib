// Copyright 2020, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package awsemfexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter"

import (
	"errors"

	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/awsutil"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/resourcetotelemetry"
	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
)

// emfSupportedUnits contains the unit collection supported by CloudWatch backend service.
// https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricDatum.html
var emfSupportedUnits = cloudwatchlogs.StandardUnit_Values()

// Config defines configuration for AWS EMF exporter.
type Config struct {
	// AWSSessionSettings contains the common configuration options
	// for creating AWS session to communicate with backend
	awsutil.AWSSessionSettings `mapstructure:",squash"`
	// LogGroupName is the name of CloudWatch log group which defines group of log streams
	// that share the same retention, monitoring, and access control settings.
	LogGroupName string `mapstructure:"log_group_name"`
	// LogStreamName is the name of CloudWatch log stream which is a sequence of log events
	// that share the same source.
	LogStreamName string `mapstructure:"log_stream_name"`
	// Namespace is a container for CloudWatch metrics.
	// Metrics in different namespaces are isolated from each other.
	Namespace string `mapstructure:"namespace"`
	// DimensionRollupOption is the option for metrics dimension rollup. Three options are available, default option is "ZeroAndSingleDimensionRollup".
	// "ZeroAndSingleDimensionRollup" - Enable both zero dimension rollup and single dimension rollup
	// "SingleDimensionRollupOnly" - Enable single dimension rollup
	// "NoDimensionRollup" - No dimension rollup (only keep original metrics which contain all dimensions)
	DimensionRollupOption string `mapstructure:"dimension_rollup_option"`

	// LogRetention is the option to set the log retention policy for the CloudWatch Log Group. Defaults to Never Expire if not specified or set to 0
	// Possible values are 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, 2192, 2557, 2922, 3288, or 3653
	LogRetention int64 `mapstructure:"log_retention"`

	// ParseJSONEncodedAttributeValues is an array of attribute keys whose corresponding values are JSON-encoded as strings.
	// Those strings will be decoded to its original json structure.
	ParseJSONEncodedAttributeValues []string `mapstructure:"parse_json_encoded_attr_values"`

	// MetricDeclarations is the list of rules to be used to set dimensions for exported metrics.
	MetricDeclarations []*MetricDeclaration `mapstructure:"metric_declarations"`

	// MetricDescriptors is the list of override metric descriptors that are sent to the CloudWatch
	MetricDescriptors []MetricDescriptor `mapstructure:"metric_descriptors"`

	// OutputDestination is an option to specify the EMFExporter output. Default option is "cloudwatch"
	// "cloudwatch" - direct the exporter output to CloudWatch backend
	// "stdout" - direct the exporter output to stdout
	// TODO: we can support directing output to a file (in the future) while customer specifies a file path here.
	OutputDestination string `mapstructure:"output_destination"`

	// EKSFargateContainerInsightsEnabled is an option to reformat certin metric labels so that they take the form of a high level object
	// The end result will make the labels look like those coming out of ECS and be more easily injected into cloudwatch
	// Note that at the moment in order to use this feature the value "kubernetes" must also be added to the ParseJSONEncodedAttributeValues array in order to be used
	EKSFargateContainerInsightsEnabled bool `mapstructure:"eks_fargate_container_insights_enabled"`

	// ResourceToTelemetrySettings is the option for converting resource attrihutes to telemetry attributes.
	// "Enabled" - A boolean field to enable/disable this option. Default is `false`.
	// If enabled, all the resource attributes will be converted to metric labels by default.
	ResourceToTelemetrySettings resourcetotelemetry.Settings `mapstructure:"resource_to_telemetry_conversion"`

	// logger is the Logger used for writing error/warning logs
	logger *zap.Logger
}

var _ component.Config = (*Config)(nil)

// Validate filters out invalid metricDeclarations and metricDescriptors
func (config *Config) Validate() error {
	var validDeclarations []*MetricDeclaration
	for _, declaration := range config.MetricDeclarations {
		err := declaration.init(config.logger)
		if err != nil {
			config.logger.Warn("Dropped metric declaration.", zap.Error(err))
		} else {
			validDeclarations = append(validDeclarations, declaration)
		}
	}
	config.MetricDeclarations = validDeclarations

	var validDescriptors []MetricDescriptor
	for _, descriptor := range config.MetricDescriptors {
		if descriptor.MetricName == "" {
			continue
		}
		if ok := slices.Contains(emfSupportedUnits, descriptor.Unit); ok {
			validDescriptors = append(validDescriptors, descriptor)
		} else {
			config.logger.Warn("Dropped unsupported metric desctriptor.", zap.String("unit", descriptor.Unit))
		}
	}
	config.MetricDescriptors = validDescriptors

	if !isValidRetentionValue(config.LogRetention) {
		return errors.New("invalid value for retention policy.  Please make sure to use the following values: 0 (Never Expire), 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, 2192, 2557, 2922, 3288, or 3653")
	}

	return nil
}

// Added function to check if value is an accepted number of log retention days
func isValidRetentionValue(input int64) bool {
	switch input {
	case
		0,
		1,
		3,
		5,
		7,
		14,
		30,
		60,
		90,
		120,
		150,
		180,
		365,
		400,
		545,
		731,
		1827,
		2192,
		2557,
		2922,
		3288,
		3653:
		return true
	}
	return false
}
