// Copyright The OpenTelemetry Authors
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

package googlecloudpubsubreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudpubsubreceiver"

import (
	"fmt"
	"regexp"

	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

var subscriptionMatcher = regexp.MustCompile(`projects/[a-z][a-z0-9\-]*/subscriptions/`)

type Config struct {
	// The mode the receiver is running in, this is either `pull` or `push` (default is `pull`)
	Mode string `mapstructure:"mode"`
	// Google Cloud Project ID where the Pubsub client will connect to
	ProjectID string `mapstructure:"project"`
	// User agent that will be used by the Pubsub client to connect to the service
	UserAgent string `mapstructure:"user_agent"`
	// Override of the Pubsub Endpoint, leave empty for the default endpoint
	Endpoint string `mapstructure:"endpoint"`
	// Only has effect if Endpoint is not ""
	Insecure bool `mapstructure:"insecure"`
	// Timeout for all API calls. If not set, defaults to 12 seconds.
	exporterhelper.TimeoutSettings `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct.

	// The fully qualified resource name of the Pubsub subscription
	Subscription string `mapstructure:"subscription"`
	// Lock down the encoding of the payload, leave empty for attribute based detection
	Encoding string `mapstructure:"encoding"`
	// Lock down the compression of the payload, leave empty for attribute based detection
	Compression string `mapstructure:"compression"`

	// The client id that will be used by Pubsub to make load balancing decisions
	ClientID string `mapstructure:"client_id"`

	// The configuration block, when selected mode is `push`
	Push PushConfig `mapstructure:"push"`
}

type PushConfig struct {
	// The path on which to listen for Pub/Sub messages
	Path string `mapstructure:"path"`
	// The settings for the HTTP server listening for requests.
	confighttp.HTTPServerSettings `mapstructure:",squash"`
}

func (c *PushConfig) validate() error {
	if c.Path == "" {
		c.Path = "/"
	}
	if c.Endpoint == "" {
		return fmt.Errorf("push endpoint needs to be defined in push mode")
	}
	return nil
}

func (config *Config) validateForLog() error {
	err := config.validate()
	if err != nil {
		return err
	}
	switch config.Encoding {
	case "":
	case "otlp_proto_log":
	case "raw_text":
	case "raw_json":
	default:
		return fmt.Errorf("log encoding %v is not supported.  supported encoding formats include [otlp_proto_log,raw_text,raw_json]", config.Encoding)
	}
	return nil
}

func (config *Config) validateForTrace() error {
	err := config.validate()
	if err != nil {
		return err
	}
	switch config.Encoding {
	case "":
	case "otlp_proto_trace":
	default:
		return fmt.Errorf("trace encoding %v is not supported.  supported encoding formats include [otlp_proto_trace]", config.Encoding)
	}
	return nil
}

func (config *Config) validateForMetric() error {
	err := config.validate()
	if err != nil {
		return err
	}
	switch config.Encoding {
	case "":
	case "otlp_proto_metric":
	default:
		return fmt.Errorf("metric encoding %v is not supported.  supported encoding formats include [otlp_proto_metric]", config.Encoding)
	}
	return nil
}

func (config *Config) validate() error {
	if config.Mode == "" {
		config.Mode = "pull"
	}
	switch config.Mode {
	case "pull":
		if !subscriptionMatcher.MatchString(config.Subscription) {
			return fmt.Errorf("subscription '%s' is not a valid format, use 'projects/<project_id>/subscriptions/<name>'", config.Subscription)
		}
	case "push":
		if config.Subscription != "" {
			return fmt.Errorf("subscription should not be specified in push mode")
		}
		err := config.Push.validate()
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("mode %v is not supported. supported modes include [pull,push]", config.Mode)
	}

	switch config.Compression {
	case "":
	case "gzip":
	default:
		return fmt.Errorf("compression %v is not supported.  supported compression formats include [gzip]", config.Compression)
	}
	return nil
}
