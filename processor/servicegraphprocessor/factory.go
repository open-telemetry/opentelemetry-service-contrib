// Copyright The OpenTelemetry Authors
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

package servicegraphprocessor // import "github.com/asserts/opentelemetry-collector-contrib/processor/servicegraphprocessor"

import (
	"context"
	"time"

	"go.opencensus.io/stats/view"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"
)

const (
	// The value of "type" key in configuration.
	typeStr = "servicegraph"
	// The stability level of the processor.
	stability          = component.StabilityLevelAlpha
	connectorStability = component.StabilityLevelDevelopment
)

// NewFactory creates a factory for the servicegraph processor.
func NewFactory() processor.Factory {
	// TODO: Handle this err
	_ = view.Register(serviceGraphProcessorViews()...)

	return processor.NewFactory(
		typeStr,
		createDefaultConfig,
		processor.WithTraces(createTracesProcessor, stability),
	)
}

// NewConnectorFactory creates a factory for the servicegraph connector.
func NewConnectorFactory() connector.Factory {
	// TODO: Handle this err
	_ = view.Register(serviceGraphProcessorViews()...)

	return connector.NewFactory(
		typeStr,
		createDefaultConfig,
		connector.WithTracesToMetrics(createTracesToMetricsConnector, connectorStability),
	)
}

func createDefaultConfig() component.Config {
	return &Config{
		Store: StoreConfig{
			TTL:      2 * time.Second,
			MaxItems: 1000,
		},
	}
}

func createTracesProcessor(_ context.Context, params processor.CreateSettings, cfg component.Config, nextConsumer consumer.Traces) (processor.Traces, error) {
	p := newProcessor(params.Logger, cfg)
	p.tracesConsumer = nextConsumer
	return p, nil
}

func createTracesToMetricsConnector(_ context.Context, params connector.CreateSettings, cfg component.Config, nextConsumer consumer.Metrics) (connector.Traces, error) {
	c := newProcessor(params.Logger, cfg)
	c.metricsConsumer = nextConsumer
	return c, nil
}
