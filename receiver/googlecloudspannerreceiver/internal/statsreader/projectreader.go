// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package statsreader // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/statsreader"

import (
	"context"
	"strings"

	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/metadata"
)

type ProjectReader struct {
	databaseReaders []CompositeReader
	logger          *zap.Logger
}

func NewProjectReader(databaseReaders []CompositeReader, logger *zap.Logger) *ProjectReader {
	return &ProjectReader{
		databaseReaders: databaseReaders,
		logger:          logger,
	}
}

func (projectReader *ProjectReader) Shutdown() {
	for _, databaseReader := range projectReader.databaseReaders {
		projectReader.logger.Info("Shutting down projectReader for database",
			zap.String("database", databaseReader.Name()))
		databaseReader.Shutdown()
	}
}

func (projectReader *ProjectReader) Read(ctx context.Context) ([]*metadata.MetricsDataPoint, error) {
	var result []*metadata.MetricsDataPoint

	for _, databaseReader := range projectReader.databaseReaders {
		dataPoints, err := databaseReader.Read(ctx)
		if err != nil {
			return nil, err
		}

		result = append(result, dataPoints...)
	}

	return result, nil
}

func (projectReader *ProjectReader) Name() string {
	databaseReaderNames := make([]string, len(projectReader.databaseReaders))

	for i, databaseReader := range projectReader.databaseReaders {
		databaseReaderNames[i] = databaseReader.Name()
	}

	return "Project reader for: " + strings.Join(databaseReaderNames, ",")
}
