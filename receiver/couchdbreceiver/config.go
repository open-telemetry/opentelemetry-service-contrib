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

package couchdbreceiver

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
	"go.uber.org/multierr"
)

var (
	defaultEndpoint = "http://localhost:5984"

	// Errors for missing required config fields.
	errMissingUsername = errors.New(`no "username" specified in config`)
	errMissingPassword = errors.New(`no "password" specified in config`)

	// Errors for invalid url components in the endpoint.
	errInvalidScheme   = errors.New(`"endpoint" requires schema of "http" or "https"`)
	errInvalidEndpoint = errors.New(`"endpoint" %q must be in the form of <scheme>://<hostname>:<port>`)
	errInvalidHost     = errors.New(`"endpoint" %q requires a hostname `)
)

// Config defines the configuration for the various elements of the receiver agent.
type Config struct {
	scraperhelper.ScraperControllerSettings `mapstructure:",squash"`
	confighttp.HTTPClientSettings           `mapstructure:",squash"`
	Username                                string `mapstructure:"username"`
	Password                                string `mapstructure:"password"`
	AllNodes                                bool   `mapstructure:"all_nodes"`
}

// Validate validates missing and invalid configuration fields.
func (cfg *Config) Validate() error {
	if missingFields := validateNotEmpty(cfg); missingFields != nil {
		return missingFields
	}
	return validateEndpoint(cfg)
}

// validateNotEmpty validates that the required fields are not empty.
func validateNotEmpty(cfg *Config) error {
	var err error
	if cfg.Username == "" {
		err = multierr.Append(err, errMissingUsername)
	}
	if cfg.Password == "" {
		err = multierr.Append(err, errMissingPassword)
	}

	return err
}

// validateEndpoint validates the endpoint by parsing the url.
func validateEndpoint(cfg *Config) error {

	// validSchema is used because url.Parse may not return an error without a schema even though it is an invalid url.
	if !validSchema(cfg.Endpoint) {
		return errInvalidScheme
	}

	u, err := url.Parse(cfg.Endpoint)
	if err != nil {
		return fmt.Errorf(errInvalidEndpoint.Error(), cfg.Endpoint)
	}

	if u.Hostname() == "" {
		return fmt.Errorf(errInvalidHost.Error(), cfg.Endpoint)
	}

	return nil
}

// validSchema returns true if any http protocol is found.
func validSchema(rawUrl string) bool {
	lowerUrl := strings.ToLower(rawUrl)
	return strings.HasPrefix(lowerUrl, "http://") || strings.HasPrefix(lowerUrl, "https://")
}
