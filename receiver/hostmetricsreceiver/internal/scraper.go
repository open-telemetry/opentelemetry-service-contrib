// Copyright  OpenTelemetry Authors
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

package internal

import (
	"context"

	"go.uber.org/zap"

	"go.opentelemetry.io/collector/consumer/pdata"
)

// Scraper gathers metrics from the host machine and converts
// these into internal metrics format.
type Scraper interface {
	// Initialize performs any timely initialization tasks such as
	// setting up performance counters for initial collection.
	Initialize(ctx context.Context) error
	// Close should clean up any unmanaged resources such as
	// performance counter handles.
	Close(ctx context.Context) error
	// ScrapeMetrics returns relevant scraped metrics. If errors occur
	// scraping some metrics, an error should be returned, but any
	// metrics that were successfully scraped should still be returned.
	ScrapeMetrics(ctx context.Context) (pdata.MetricSlice, error)
}

// Factory can create a Scraper.
type Factory interface {
	// CreateDefaultConfig creates the default configuration for the Scraper.
	CreateDefaultConfig() Config

	// CreateMetricsScraper creates a scraper based on this config.
	// If the config is not valid, error will be returned instead.
	CreateMetricsScraper(
		ctx context.Context,
		logger *zap.Logger,
		cfg Config) (Scraper, error)
}

// Config is the configuration of a scraper.
type Config interface {
}

// ConfigSettings provides common settings for scraper configuration.
type ConfigSettings struct {
}
