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

package mongodbatlasreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver"

import (
	"context"
	"fmt"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/receiver/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver/internal/metadata"
)

const (
	typeStr              = "mongodbatlas"
	stability            = component.StabilityLevelBeta
	defaultGranularity   = "PT1M" // 1-minute, as per https://docs.atlas.mongodb.com/reference/api/process-measurements/
	defaultAlertsEnabled = false
)

// receivers is a mapping of receiver ID to a receiver that is shared between metrics and log scraping
var receivers = make(map[string]*receiver)

// NewFactory creates a factory for MongoDB Atlas receiver
func NewFactory() component.ReceiverFactory {
	return component.NewReceiverFactory(
		typeStr,
		createDefaultConfig,
		component.WithMetricsReceiverAndStabilityLevel(createMetricsReceiver, stability),
		component.WithLogsReceiverAndStabilityLevel(createCombinedLogReceiver, stability))

}

func createMetricsReceiver(
	_ context.Context,
	params component.ReceiverCreateSettings,
	rConf config.Receiver,
	consumer consumer.Metrics,
) (component.MetricsReceiver, error) {
	cfg := rConf.(*Config)
	recv, err := getReceiver(params, cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to create a MongoDB Atlas Receiver instance: %w", err)
	}

	ms, err := newMongoDBAtlasScraper(recv)
	if err != nil {
		return nil, fmt.Errorf("unable to create a MongoDB Atlas Scaper instance: %w", err)
	}

	return scraperhelper.NewScraperControllerReceiver(&cfg.ScraperControllerSettings, params, consumer, scraperhelper.AddScraper(ms))
}

func createAlertsLogReceiver(
	_ context.Context,
	params component.ReceiverCreateSettings,
	rConf config.Receiver,
	consumer consumer.Logs,
) (component.LogsReceiver, error) {
	cfg := rConf.(*Config)
	recv, err := newAlertsReceiver(params.Logger, cfg.Alerts, consumer)
	if err != nil {
		return nil, fmt.Errorf("unable to create a MongoDB Atlas Receiver instance: %w", err)
	}

	return recv, nil
}

func createCombinedLogReceiver(
	_ context.Context,
	params component.ReceiverCreateSettings,
	rConf config.Receiver,
	consumer consumer.Logs,
) (component.LogsReceiver, error) {
	cfg := rConf.(*Config)

	var err error
	recv := &combindedLogsReceiver{}

	// If alerts is enabled create alerts receiver
	if cfg.Alerts.Enabled {
		recv.alerts, err = newAlertsReceiver(params.Logger, cfg.Alerts, consumer)
		if err != nil {
			return nil, fmt.Errorf("unable to create a MongoDB Atlas Alerts Receiver instance: %w", err)
		}
	}

	// If logs is enabled create logs receiver
	if cfg.Logs.Enabled {
		recv.logs, err = getReceiver(params, cfg)
		if err != nil {
			return nil, fmt.Errorf("unable to create a MongoDB Atlas Logs Receiver instance: %w", err)
		}
		recv.logs.consumer = consumer
	}

	return recv, nil
}

func createDefaultConfig() config.Receiver {
	return &Config{
		ScraperControllerSettings: scraperhelper.NewDefaultScraperControllerSettings(typeStr),
		Granularity:               defaultGranularity,
		RetrySettings:             exporterhelper.NewDefaultRetrySettings(),
		Metrics:                   metadata.DefaultMetricsSettings(),
		Alerts: AlertConfig{
			Enabled: defaultAlertsEnabled,
		},
		Logs: LogConfig{
			Enabled:  true,
			Projects: []*Project{},
		},
	}
}

// getReceiver ensures we only create a single receiver per receiver ID because it is shared between logs and metrics
func getReceiver(params component.ReceiverCreateSettings, cfg *Config) (*receiver, error) {
	mongoDBReceiver, ok := receivers[cfg.ID().String()]
	var err error
	if !ok {
		mongoDBReceiver, err = newMongoDBAtlasReciever(params, cfg)
		receivers[cfg.ID().String()] = mongoDBReceiver
	}

	return mongoDBReceiver, err
}
