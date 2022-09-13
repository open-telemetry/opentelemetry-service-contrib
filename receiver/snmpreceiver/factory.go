// Copyright 2020 OpenTelemetry Authors
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

package snmpreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/snmpreceiver"

import (
	"context"
	"errors"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

const (
	typeStr   = "snmp"
	stability = component.StabilityLevelBeta

	defaultCollectionInterval = 10 // In seconds
	defaultEndpoint           = "localhost:161"
	defaultVersion            = "v2c"
	defaultCommunity          = "public"
)

var errConfigNotSNMP = errors.New("config was not a SNMP receiver config")

// NewFactory creates a new receiver factory for SNMP
func NewFactory() component.ReceiverFactory {
	return component.NewReceiverFactory(
		typeStr,
		createDefaultConfig,
		component.WithMetricsReceiver(createMetricsReceiver, stability))
}

// createDefaultConfig creates a config for Big-IP with as many default values as possible
func createDefaultConfig() config.Receiver {
	return &Config{
		ScraperControllerSettings: scraperhelper.ScraperControllerSettings{
			ReceiverSettings:   config.NewReceiverSettings(config.NewComponentID(typeStr)),
			CollectionInterval: defaultCollectionInterval * time.Second,
		},
		Endpoint:  defaultEndpoint,
		Version:   defaultVersion,
		Community: defaultCommunity,
	}
}

// createMetricsReceiver creates the metric receiver for SNMP
func createMetricsReceiver(
	_ context.Context,
	params component.ReceiverCreateSettings,
	config config.Receiver,
	consumer consumer.Metrics,
) (component.MetricsReceiver, error) {
	snmpConfig, ok := config.(*Config)
	if !ok {
		return nil, errConfigNotSNMP
	}

	snmpScraper := newScraper(params.Logger, snmpConfig, params)
	scraper, err := scraperhelper.NewScraper(typeStr, snmpScraper.scrape, scraperhelper.WithStart(snmpScraper.start))
	if err != nil {
		return nil, err
	}

	return scraperhelper.NewScraperControllerReceiver(&snmpConfig.ScraperControllerSettings, params, consumer, scraperhelper.AddScraper(scraper))
}
