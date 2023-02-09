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

package servicegraphprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/servicegraphprocessor"

import (
	"context"
	"go.opentelemetry.io/collector/featuregate"
	"time"

	"go.opencensus.io/stats/view"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"
)

const (
	// The value of "type" key in configuration.
	typeStr = "servicegraph"
	// The stability level of the processor.
	stability                = component.StabilityLevelAlpha
	virtualNodeFeatureGateId = "processor.servicegraph.virtualNode"
)

var virtualNodeFeatureGate *featuregate.Gate

func init() {
	virtualNodeFeatureGate = featuregate.GlobalRegistry().MustRegister(
		virtualNodeFeatureGateId,
		featuregate.StageAlpha,
		featuregate.WithRegisterDescription("When enabled, when the edge expires, processor checks if it has peer attributes(`db.name, net.sock.peer.addr, net.peer.name, rpc.service, http.url, http.target`), and then aggregate the metrics with virtual node."),
		featuregate.WithRegisterReferenceURL("https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/17196"),
	)
}

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

func createDefaultConfig() component.Config {
	return &Config{
		Store: StoreConfig{
			TTL:      2 * time.Second,
			MaxItems: 1000,
		},
	}
}

func createTracesProcessor(_ context.Context, params processor.CreateSettings, cfg component.Config, nextConsumer consumer.Traces) (processor.Traces, error) {
	return newProcessor(params.Logger, cfg, nextConsumer), nil
}
