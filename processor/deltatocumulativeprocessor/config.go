// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package deltatocumulativeprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor"

import (
	"fmt"
	"time"

	"go.opentelemetry.io/collector/component"
)

var _ component.ConfigValidator = (*Config)(nil)

type Config struct {
	MaxStale   time.Duration `mapstructure:"max_stale"`
	MaxStreams int           `mapstructure:"max_streams"`
}

func (c *Config) Validate() error {
	if c.MaxStale <= 0 {
		return fmt.Errorf("max_stale must be a positive duration (got %s)", c.MaxStale)
	}
	if c.MaxStreams < 0 {
		return fmt.Errorf("max_streams must be a positive number (got %d)", c.MaxStreams)
	}
	return nil
}

func createDefaultConfig() component.Config {
	return &Config{
		MaxStale: 5 * time.Minute,

		// disable. TODO: find good default
		// https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/31603
		MaxStreams: 0,
	}
}
