// Copyright 2020, OpenTelemetry Authors
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

package awscloudwatchlogsexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awscloudwatchlogsexporter"

import (
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/awsutil"
)

// Config represent a configuration for the CloudWatch logs exporter.
type Config struct {
	exporterhelper.RetrySettings `mapstructure:"retry_on_failure"`

	// LogGroupName is the name of CloudWatch log group which defines group of log streams
	// that share the same retention, monitoring, and access control settings.
	LogGroupName string `mapstructure:"log_group_name"`

	// LogStreamName is the name of CloudWatch log stream which is a sequence of log events
	// that share the same source.
	LogStreamName string `mapstructure:"log_stream_name"`

	// Endpoint is the CloudWatch Logs service endpoint which the requests
	// are forwarded to. https://docs.aws.amazon.com/general/latest/gr/cwl_region.html
	// e.g. logs.us-east-1.amazonaws.com
	// Optional.
	Endpoint string `mapstructure:"endpoint"`

	// LogRetention is the option to set the log retention policy for the CloudWatch Log Group. Defaults to Never Expire if not specified or set to 0
	// Possible values are 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, 2192, 2557, 2922, 3288, or 3653
	LogRetention int64 `mapstructure:"log_retention"`

	// QueueSettings is a subset of exporterhelper.QueueSettings,
	// because only QueueSize is user-settable due to how AWS CloudWatch API works
	QueueSettings QueueSettings `mapstructure:"sending_queue"`

	logger *zap.Logger

	awsutil.AWSSessionSettings `mapstructure:",squash"`
}

type QueueSettings struct {
	// QueueSize set the length of the sending queue
	QueueSize int `mapstructure:"queue_size"`
}

var _ component.Config = (*Config)(nil)

// Validate config
func (config *Config) Validate() error {
	if config.LogGroupName == "" {
		return errors.New("'log_group_name' must be set")
	}
	if config.LogStreamName == "" {
		return errors.New("'log_stream_name' must be set")
	}
	if config.QueueSettings.QueueSize < 1 {
		return errors.New("'sending_queue.queue_size' must be 1 or greater")
	}
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

func (config *Config) enforcedQueueSettings() exporterhelper.QueueSettings {
	return exporterhelper.QueueSettings{
		Enabled: true,
		// due to the sequence token, there can be only one request in flight
		NumConsumers: 1,
		QueueSize:    config.QueueSettings.QueueSize,
	}
}

// TODO(jbd): Add ARN role to config.
