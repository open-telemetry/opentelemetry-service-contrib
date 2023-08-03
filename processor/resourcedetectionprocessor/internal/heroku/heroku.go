// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package heroku // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/heroku"

import (
	"context"
	"os"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	conventions "go.opentelemetry.io/collector/semconv/v1.6.1"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/heroku/internal/metadata"
)

const (
	// TypeStr is type of detector.
	TypeStr = "heroku"
)

// NewDetector returns a detector which can detect resource attributes on Heroku
func NewDetector(set processor.CreateSettings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	cfg := dcfg.(Config)
	return &detector{
		logger: set.Logger,
		rb:     metadata.NewResourceBuilder(cfg.ResourceAttributes),
	}, nil
}

type detector struct {
	logger *zap.Logger
	rb     *metadata.ResourceBuilder
}

// Detect detects heroku metadata and returns a resource with the available ones
func (d *detector) Detect(_ context.Context) (resource pcommon.Resource, schemaURL string, err error) {
	dynoID, ok := os.LookupEnv("HEROKU_DYNO_ID")
	if !ok {
		d.logger.Debug("heroku metadata unavailable", zap.Error(err))
		return pcommon.NewResource(), "", nil
	}

	d.rb.SetCloudProvider("heroku")
	d.rb.SetServiceInstanceID(dynoID)
	if v, ok := os.LookupEnv("HEROKU_APP_ID"); ok {
		d.rb.SetHerokuAppID(v)
	}
	if v, ok := os.LookupEnv("HEROKU_APP_NAME"); ok {
		d.rb.SetServiceName(v)
	}
	if v, ok := os.LookupEnv("HEROKU_RELEASE_CREATED_AT"); ok {
		d.rb.SetHerokuReleaseCreationTimestamp(v)
	}
	if v, ok := os.LookupEnv("HEROKU_RELEASE_VERSION"); ok {
		d.rb.SetServiceVersion(v)
	}
	if v, ok := os.LookupEnv("HEROKU_SLUG_COMMIT"); ok {
		d.rb.SetHerokuReleaseCommit(v)
	}

	return d.rb.Emit(), conventions.SchemaURL, nil
}
