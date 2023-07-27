// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package chronyreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/chronyreceiver"

import (
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/receiver/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/chronyreceiver/internal/chrony"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/chronyreceiver/internal/metadata"
)

type Config struct {
	scraperhelper.ScraperControllerSettings `mapstructure:",squash"`
	metadata.MetricsBuilderConfig           `mapstructure:",squash"`
	// Endpoint is the published address or unix socket
	// that allows clients to connect to:
	// The allowed format is:
	//   unix:///path/to/chronyd/unix.sock
	//   udp://localhost:323
	//
	// The default value is unix:///var/run/chrony/chronyd.sock
	Endpoint string `mapstructure:"endpoint"`
}

var (
	_ component.Config = (*Config)(nil)

	errInvalidValue = errors.New("invalid value")
)

func newDefaultCongfig() component.Config {
	return &Config{
		ScraperControllerSettings: scraperhelper.NewDefaultScraperControllerSettings(metadata.Type),
		MetricsBuilderConfig:      metadata.DefaultMetricsBuilderConfig(),
		Endpoint:                  "unix:///var/run/chrony/chronyd.sock",
	}
}

func (c *Config) Validate() error {
	_, _, err := chrony.SplitNetworkEndpoint(c.Endpoint)
	return err
}
