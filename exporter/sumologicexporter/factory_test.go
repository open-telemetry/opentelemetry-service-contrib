// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sumologicexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sumologicexporter"

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configauth"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

func TestType(t *testing.T) {
	factory := NewFactory()
	pType := factory.Type()
	assert.Equal(t, pType, component.Type("sumologic"))
}

func TestCreateDefaultConfig(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()
	qs := exporterhelper.NewDefaultQueueSettings()
	qs.Enabled = false

	assert.Equal(t, cfg, &Config{
		CompressEncoding:   "gzip",
		MaxRequestBodySize: 1_048_576,
		LogFormat:          "otlp",
		MetricFormat:       "otlp",
		Client:             "otelcol",
		ClearLogsTimestamp: true,
		JSONLogs: JSONLogs{
			LogKey:       "log",
			AddTimestamp: true,
			TimestampKey: "timestamp",
		},
		TraceFormat: "otlp",

		HTTPClientSettings: confighttp.HTTPClientSettings{
			Timeout: 5 * time.Second,
			Auth: &configauth.Authentication{
				AuthenticatorID: component.NewID("sumologic"),
			},
		},
		RetrySettings: exporterhelper.NewDefaultRetrySettings(),
		QueueSettings: qs,
	})

	assert.NoError(t, component.ValidateConfig(cfg))
}
