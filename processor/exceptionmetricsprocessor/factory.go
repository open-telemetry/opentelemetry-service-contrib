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

package exceptionmetricsprocessor

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/featuregate"
	"go.opentelemetry.io/collector/processor"
)

const (
	// The value of "type" key in configuration.
	typeStr = "exceptionmetrics"
	// The stability level of the processor.
	stability = component.StabilityLevelDevelopment
)

// NewFactory creates a factory for the exceptionmetrics processor.
func NewFactory() processor.Factory {
	return processor.NewFactory(
		typeStr,
		createDefaultConfig,
		processor.WithTraces(createTracesProcessor, stability),
	)
}

func createDefaultConfig() component.Config {
	return &Config{
		DimensionsCacheSize: defaultDimensionsCacheSize,
		Dimensions: []Dimension{
			{Name: "exception.type"},
			{Name: "exception.message"},
		},
		skipSanitizeLabel: featuregate.GetRegistry().IsEnabled(dropSanitizationGateID),
	}
}

func createTracesProcessor(_ context.Context, params processor.CreateSettings, cfg component.Config, nextConsumer consumer.Traces) (processor.Traces, error) {
	return newProcessor(params.Logger, cfg, nextConsumer)
}
