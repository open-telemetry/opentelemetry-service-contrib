// Copyright 2020, OpenTelemetry Authors
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

// Package awscloudwatchlogsexporter provides a logging exporter for the OpenTelemetry collector.
// This package is subject to change and may break configuration settings and behavior.
package awscloudwatchlogsexporter

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

const typeStr = "awscloudwatchlogs"

func NewFactory() component.ExporterFactory {
	return exporterhelper.NewFactory(
		typeStr,
		createDefaultConfig,
		exporterhelper.WithLogs(createLogsExporter))
}

func createDefaultConfig() config.Exporter {
	return &Config{
		ExporterSettings: config.NewExporterSettings(config.NewID(typeStr)),
	}
}

func createLogsExporter(_ context.Context, params component.ExporterCreateSettings, cfg config.Exporter) (component.LogsExporter, error) {
	oCfg, ok := cfg.(*Config)
	if !ok {
		return nil, errors.New("invalid configuration type; can't cast to awscloudwatchlogsexporter.Config")
	}

	exporter := &exporter{config: oCfg, logger: params.Logger}
	return exporterhelper.NewLogsExporter(
		oCfg,
		params.Logger,
		exporter.PushLogs,
		exporterhelper.WithQueue(exporterhelper.QueueSettings{
			Enabled:      true,
			NumConsumers: 1, // due to the sequence token, there can be only one request in flight
		}),
		exporterhelper.WithRetry(oCfg.RetrySettings),
	)
}
