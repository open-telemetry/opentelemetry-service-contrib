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

package datasetexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datasetexporter"

import (
	"fmt"
	"os"
	"time"

	"github.com/scalyr/dataset-go/pkg/buffer"
	datasetConfig "github.com/scalyr/dataset-go/pkg/config"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

const maxDelay = 15 * time.Millisecond
const tracesMaxWait = 5 * time.Second

type TracesSettings struct {
	MaxWait time.Duration `mapstructure:"max_wait"`
}

// NewDefaultTracesSettings returns the default settings for TracesSettings.
func NewDefaultTracesSettings() TracesSettings {
	return TracesSettings{
		MaxWait: tracesMaxWait,
	}
}

type Config struct {
	DatasetURL                     string        `mapstructure:"dataset_url"`
	APIKey                         string        `mapstructure:"api_key"`
	MaxDelay                       time.Duration `mapstructure:"max_delay"`
	GroupBy                        []string      `mapstructure:"group_by"`
	TracesSettings                 `mapstructure:"traces"`
	exporterhelper.RetrySettings   `mapstructure:"retry_on_failure"`
	exporterhelper.QueueSettings   `mapstructure:"sending_queue"`
	exporterhelper.TimeoutSettings `mapstructure:"timeout"`
}

func (c *Config) Unmarshal(conf *confmap.Conf) error {
	if err := conf.Unmarshal(c, confmap.WithErrorUnused()); err != nil {
		return fmt.Errorf("cannot unmarshal config: %w", err)
	}

	if len(c.DatasetURL) == 0 {
		c.DatasetURL = os.Getenv("DATASET_URL")
	}
	if len(c.APIKey) == 0 {
		c.APIKey = os.Getenv("DATASET_API_KEY")
	}

	if c.MaxDelay == 0 {
		c.MaxDelay = maxDelay
	}

	if c.TracesSettings.MaxWait == 0 {
		c.TracesSettings.MaxWait = tracesMaxWait
	}

	return nil
}

// Validate checks if all required fields in Config are set and have valid values.
// If any of the required fields are missing or have invalid values, it returns an error.
func (c *Config) Validate() error {
	if c.APIKey == "" {
		return fmt.Errorf("api_key is required")
	}
	if c.DatasetURL == "" {
		return fmt.Errorf("dataset_url is required")
	}

	return nil
}

// String returns a string representation of the Config object.
// It includes all the fields and their values in the format "field_name: field_value".
func (c *Config) String() string {
	s := ""
	s += fmt.Sprintf("%s: %s; ", "DatasetURL", c.DatasetURL)
	s += fmt.Sprintf("%s: %s; ", "MaxDelay", c.MaxDelay)
	s += fmt.Sprintf("%s: %s; ", "GroupBy", c.GroupBy)
	s += fmt.Sprintf("%s: %+v; ", "TracesSettings", c.TracesSettings)
	s += fmt.Sprintf("%s: %+v; ", "RetrySettings", c.RetrySettings)
	s += fmt.Sprintf("%s: %+v; ", "QueueSettings", c.QueueSettings)
	s += fmt.Sprintf("%s: %+v", "TimeoutSettings", c.TimeoutSettings)

	return s
}

func (c *Config) Convert() (*ExporterConfig, error) {
	err := c.Validate()
	if err != nil {
		return nil, fmt.Errorf("config is not valid: %w", err)
	}

	return &ExporterConfig{
			datasetConfig: &datasetConfig.DataSetConfig{
				Endpoint:       c.DatasetURL,
				Tokens:         datasetConfig.DataSetTokens{WriteLog: c.APIKey},
				MaxBufferDelay: c.MaxDelay,
				MaxPayloadB:    buffer.LimitBufferSize,
				GroupBy:        c.GroupBy,
				RetryBase:      5 * time.Second,
			},
			tracesSettings: c.TracesSettings,
		},
		nil
}

type ExporterConfig struct {
	datasetConfig  *datasetConfig.DataSetConfig
	tracesSettings TracesSettings
}
