// Copyright 2020, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package alibabacloudlogserviceexporter

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configmodels"
	"go.opentelemetry.io/collector/consumer/consumerdata"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.uber.org/zap"
)

// NewMetricsExporter return a new LogSerice metrics exporter.
func NewMetricsExporter(logger *zap.Logger, cfg configmodels.Exporter) (component.MetricsExporterOld, error) {

	l := &logServiceMetricsSender{
		logger: logger,
	}

	var err error
	if l.producer, err = NewProducer(cfg.(*Config), logger); err != nil {
		return nil, err
	}
	logger.Info("Create LogService metrics exporter success")

	return exporterhelper.NewMetricsExporterOld(
		cfg,
		l.pushMetricsData)
}

type logServiceMetricsSender struct {
	logger   *zap.Logger
	producer Producer
}

func (s *logServiceMetricsSender) pushMetricsData(
	ctx context.Context,
	td consumerdata.MetricsData,
) (droppedTimeSeries int, err error) {
	logs, droppedTimeSeries, err := metricsDataToLogServiceData(s.logger, td)
	if len(logs) > 0 {
		err = s.producer.SendLogs(logs)
	}
	return droppedTimeSeries, err
}
