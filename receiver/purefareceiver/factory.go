// Copyright 2022 The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package purefareceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/purefareceiver"

// This file implements Factory for Array scraper.
import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/consumer"
)

// NewFactory creates a factory for Pure Storage FlashArray receiver.
const (
	typeStr   = "purefa"
	stability = component.StabilityLevelDevelopment
)

func NewFactory() component.ReceiverFactory {
	return component.NewReceiverFactory(
		typeStr,
		createDefaultConfig,
		component.WithMetricsReceiver(createMetricsReceiver, stability))
}

func createDefaultConfig() component.Config {
	return &Config{
		ReceiverSettings:   config.NewReceiverSettings(component.NewID(typeStr)),
		HTTPClientSettings: confighttp.HTTPClientSettings{},
		Settings: &Settings{
			ReloadIntervals: &ReloadIntervals{
				Array: 15 * time.Second,
			},
		},
	}
}

func createMetricsReceiver(
	_ context.Context,
	set component.ReceiverCreateSettings,
	rCfg component.Config,
	next consumer.Metrics,
) (component.MetricsReceiver, error) {
	cfg, ok := rCfg.(*Config)
	if !ok {
		return nil, fmt.Errorf("a purefa receiver config was expected by the receiver factory, but got %T", rCfg)
	}
	return newReceiver(cfg, set, next), nil
}
