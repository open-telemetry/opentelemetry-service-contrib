// Copyright  The OpenTelemetry Authors
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

package vcenterreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/vcenterreceiver"

import (
	"errors"
	"fmt"
	"net/url"

	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/configtls"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
	"go.uber.org/multierr"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/vcenterreceiver/internal/metadata"
)

// Config is the configuration of the receiver
type Config struct {
	config.ReceiverSettings `mapstructure:",squash"`
	MetricsConfig           *MetricsConfig `mapstructure:"metrics,omitempty"`
}

// MetricsConfig is the metrics configuration of the receiver
type MetricsConfig struct {
	scraperhelper.ScraperControllerSettings `mapstructure:",squash"`
	configtls.TLSClientSetting              `mapstructure:"tls,omitempty"`
	Settings                                metadata.MetricsSettings `mapstructure:"settings"`
	Endpoint                                string                   `mapstructure:"endpoint"`
	Username                                string                   `mapstructure:"username"`
	Password                                string                   `mapstructure:"password"`
}

// Validate checks to see if the supplied config will work for the receiver
func (c *Config) Validate() error {
	var err error
	metricsErr := c.validateMetricsConfig()
	if metricsErr != nil {
		err = multierr.Append(err, metricsErr)
	}

	return err
}

// ID returns the ID of the component.
func (c *Config) ID() config.ComponentID {
	// defaulting to use the MetricsConfig ID
	return c.MetricsConfig.ID()
}

func (c *Config) validateMetricsConfig() error {
	mc := c.MetricsConfig
	if mc == nil {
		return nil
	}

	if mc.Endpoint == "" {
		return errors.New("no endpoint was provided")
	}

	var err error
	res, err := url.Parse(mc.Endpoint)
	if err != nil {
		err = multierr.Append(err, fmt.Errorf("unable to parse url %s: %w", c.MetricsConfig.Endpoint, err))
		return err
	}

	if res.Scheme != "http" && res.Scheme != "https" {
		err = multierr.Append(err, errors.New("url scheme must be http or https"))
	}

	if mc.Username == "" {
		err = multierr.Append(err, errors.New("username not provided and is required"))
	}

	if mc.Password == "" {
		err = multierr.Append(err, errors.New("password not provided and is required"))
	}

	if _, tlsErr := mc.LoadTLSConfig(); err != nil {
		err = multierr.Append(err, fmt.Errorf("error loading tls configuration: %w", tlsErr))
	}

	return err
}

// SDKUrl returns the url for the vCenter SDK
func (c *Config) SDKUrl() (*url.URL, error) {
	res, err := url.Parse(c.MetricsConfig.Endpoint)
	if err != nil {
		return res, err
	}
	res.Path = "/sdk"
	return res, nil
}
