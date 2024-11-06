// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logzioexporter

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/confmap/confmaptest"
	"go.opentelemetry.io/collector/exporter/exportertest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/logzioexporter/internal/metadata"
)

func TestCreateDefaultConfig(t *testing.T) {
	cfg := createDefaultConfig()
	assert.NotNil(t, cfg, "failed to create default config")
	assert.NoError(t, componenttest.CheckConfigStruct(cfg))
}

func TestCreateTraces(t *testing.T) {

	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()

	sub, err := cm.Sub(component.NewIDWithName(metadata.Type, "2").String())
	require.NoError(t, err)
	require.NoError(t, sub.Unmarshal(cfg))

	params := exportertest.NewNopSettings()
	exporter, err := factory.CreateTraces(context.Background(), params, cfg)
	assert.NoError(t, err)
	assert.NotNil(t, exporter)
}

func TestGenerateUrl(t *testing.T) {
	type generateURLTest struct {
		endpoint string
		region   string
		expected string
	}
	var generateURLTests = []generateURLTest{
		{"", "us", "https://listener.logz.io:8071/?token=token"},
		{"", "", "https://listener.logz.io:8071/?token=token"},
		{"https://doesnotexist.com", "", "https://doesnotexist.com"},
		{"https://doesnotexist.com", "us", "https://doesnotexist.com"},
		{"https://doesnotexist.com", "not-valid", "https://doesnotexist.com"},
		{"", "not-valid", "https://listener.logz.io:8071/?token=token"},
		{"", "US", "https://listener.logz.io:8071/?token=token"},
		{"", "Us", "https://listener.logz.io:8071/?token=token"},
		{"", "EU", "https://listener-eu.logz.io:8071/?token=token"},
	}
	for _, test := range generateURLTests {
		clientConfig := confighttp.NewDefaultClientConfig()
		clientConfig.Endpoint = test.endpoint
		cfg := &Config{
			Region:       test.region,
			Token:        "token",
			ClientConfig: clientConfig,
		}
		output, _ := generateTracesEndpoint(cfg)
		require.Equal(t, test.expected, output)
	}
}

func TestGetTracesListenerURL(t *testing.T) {
	type getListenerURLTest struct {
		arg1     string
		expected string
	}
	var getListenerURLTests = []getListenerURLTest{
		{"us", "https://listener.logz.io:8071"},
		{"eu", "https://listener-eu.logz.io:8071"},
		{"au", "https://listener-au.logz.io:8071"},
		{"ca", "https://listener-ca.logz.io:8071"},
		{"nl", "https://listener-nl.logz.io:8071"},
		{"uk", "https://listener-uk.logz.io:8071"},
		{"wa", "https://listener-wa.logz.io:8071"},
		{"not-valid", "https://listener.logz.io:8071"},
		{"", "https://listener.logz.io:8071"},
		{"US", "https://listener.logz.io:8071"},
		{"Us", "https://listener.logz.io:8071"},
	}
	for _, test := range getListenerURLTests {
		output := getTracesListenerURL(test.arg1)
		require.Equal(t, test.expected, output)
	}
}

func TestGetLogsListenerURL(t *testing.T) {
	type getListenerURLTest struct {
		arg1     string
		expected string
	}
	var getListenerURLTests = []getListenerURLTest{
		{"us", "https://otlp-listener.logz.io/v1/logs"},
		{"eu", "https://otlp-listener-eu.logz.io/v1/logs"},
		{"au", "https://otlp-listener-au.logz.io/v1/logs"},
		{"ca", "https://otlp-listener-ca.logz.io/v1/logs"},
		{"nl", "https://otlp-listener-nl.logz.io/v1/logs"},
		{"uk", "https://otlp-listener-uk.logz.io/v1/logs"},
		{"wa", "https://otlp-listener-wa.logz.io/v1/logs"},
		{"not-valid", "https://otlp-listener.logz.io/v1/logs"},
		{"", "https://otlp-listener.logz.io/v1/logs"},
		{"US", "https://otlp-listener.logz.io/v1/logs"},
		{"Us", "https://otlp-listener.logz.io/v1/logs"},
	}
	for _, test := range getListenerURLTests {
		output := getLogsListenerURL(test.arg1)
		require.Equal(t, test.expected, output)
	}
}
