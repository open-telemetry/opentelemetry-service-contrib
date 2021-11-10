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

package prometheustracemetricsexporter

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

const (
	typeStr = config.Type("prometheustracemetrics")
)

// NewFactory creates a new Prometheus Trace Metrics exporter.
func NewFactory() component.ExporterFactory {
	return exporterhelper.NewFactory(
		typeStr,
		createDefaultConfig,
		exporterhelper.WithTraces(createTraceExporter))
}

func createTraceExporter(_ context.Context, set component.ExporterCreateSettings, cfg config.Exporter) (component.TracesExporter, error) {
	ptmCfg, ok := cfg.(*Config)
	if !ok {
		return nil, errors.New("invalid configuration")
	}

	return newExporter(set.Logger, ptmCfg), nil
}
