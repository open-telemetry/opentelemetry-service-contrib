// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package clickhouseexporter

import (
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/cenkalti/backoff/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/confmap/confmaptest"
	"go.opentelemetry.io/collector/exporter/exporterhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/clickhouseexporter/internal/metadata"
)

const defaultEndpoint = "clickhouse://127.0.0.1:9000"

func TestLoadConfig(t *testing.T) {
	t.Parallel()

	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)

	defaultCfg := createDefaultConfig()
	defaultCfg.(*Config).Endpoint = defaultEndpoint

	storageID := component.MustNewIDWithName("file_storage", "clickhouse")

	tests := []struct {
		id       component.ID
		expected component.Config
	}{

		{
			id:       component.NewIDWithName(metadata.Type, ""),
			expected: defaultCfg,
		},
		{
			id: component.NewIDWithName(metadata.Type, "full"),
			expected: &Config{
				Endpoint:         defaultEndpoint,
				Database:         "otel",
				Username:         "foo",
				Password:         "bar",
				TTL:              72 * time.Hour,
				LogsTableName:    "otel_logs",
				TracesTableName:  "otel_traces",
				MetricsTableName: "otel_metrics",
				TimeoutSettings: exporterhelper.TimeoutSettings{
					Timeout: 5 * time.Second,
				},
				BackOffConfig: configretry.BackOffConfig{
					Enabled:             true,
					InitialInterval:     5 * time.Second,
					MaxInterval:         30 * time.Second,
					MaxElapsedTime:      300 * time.Second,
					RandomizationFactor: backoff.DefaultRandomizationFactor,
					Multiplier:          backoff.DefaultMultiplier,
				},
				ConnectionParams: map[string]string{},
				QueueSettings: exporterhelper.QueueSettings{
					Enabled:      true,
					NumConsumers: 1,
					QueueSize:    100,
					StorageID:    &storageID,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.id.String(), func(t *testing.T) {
			factory := NewFactory()
			cfg := factory.CreateDefaultConfig()

			sub, err := cm.Sub(tt.id.String())
			require.NoError(t, err)
			require.NoError(t, component.UnmarshalConfig(sub, cfg))

			assert.NoError(t, component.ValidateConfig(cfg))
			assert.Equal(t, tt.expected, cfg)
		})
	}
}

func withDefaultConfig(fns ...func(*Config)) *Config {
	cfg := createDefaultConfig().(*Config)
	for _, fn := range fns {
		fn(cfg)
	}
	return cfg
}

func TestConfig_buildDSN(t *testing.T) {
	type fields struct {
		Endpoint         string
		Username         string
		Password         string
		Database         string
		ConnectionParams map[string]string
		AsyncInsert      *bool
	}
	type args struct {
		database string
	}
	type ChOptions struct {
		Secure      bool
		DialTimeout time.Duration
		Compress    clickhouse.CompressionMethod
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		want          string
		wantChOptions ChOptions
		wantErr       error
	}{
		{
			name: "valid default config",
			fields: fields{
				Endpoint: defaultEndpoint,
			},
			args: args{},
			wantChOptions: ChOptions{
				Secure: false,
			},
			want: "clickhouse://127.0.0.1:9000/default?async_insert=true",
		},
		{
			name: "Support tcp scheme",
			fields: fields{
				Endpoint: "tcp://127.0.0.1:9000",
			},
			args: args{},
			wantChOptions: ChOptions{
				Secure: false,
			},
			want: "tcp://127.0.0.1:9000/default?async_insert=true",
		},
		{
			name: "prefers database name from config over from DSN",
			fields: fields{
				Endpoint: defaultEndpoint,
				Username: "foo",
				Password: "bar",
				Database: "otel",
			},
			args: args{
				database: "otel",
			},
			wantChOptions: ChOptions{
				Secure: false,
			},
			want: "clickhouse://foo:bar@127.0.0.1:9000/otel?async_insert=true",
		},
		{
			name: "use database name from DSN if not set in config",
			fields: fields{
				Endpoint: "clickhouse://foo:bar@127.0.0.1:9000/otel",
				Username: "foo",
				Password: "bar",
				Database: "",
			},
			args: args{
				database: "",
			},
			wantChOptions: ChOptions{
				Secure: false,
			},
			want: "clickhouse://foo:bar@127.0.0.1:9000/otel?async_insert=true",
		},
		{
			name: "invalid config",
			fields: fields{
				Endpoint: "127.0.0.1:9000",
			},
			wantChOptions: ChOptions{
				Secure: false,
			},
			wantErr: errConfigInvalidEndpoint,
		},
		{
			name: "Auto enable TLS connection based on scheme",
			fields: fields{
				Endpoint: "https://127.0.0.1:9000",
			},
			wantChOptions: ChOptions{
				Secure: true,
			},
			args: args{},
			want: "https://127.0.0.1:9000/default?async_insert=true&secure=true",
		},
		{
			name: "Preserve query parameters",
			fields: fields{
				Endpoint: "clickhouse://127.0.0.1:9000?secure=true&foo=bar",
			},
			wantChOptions: ChOptions{
				Secure: true,
			},
			args: args{},
			want: "clickhouse://127.0.0.1:9000/default?async_insert=true&foo=bar&secure=true",
		},
		{
			name: "Parse clickhouse settings",
			fields: fields{
				Endpoint: "https://127.0.0.1:9000?secure=true&dial_timeout=30s&compress=lz4",
			},
			wantChOptions: ChOptions{
				Secure:      true,
				DialTimeout: 30 * time.Second,
				Compress:    clickhouse.CompressionLZ4,
			},
			args: args{},
			want: "https://127.0.0.1:9000/default?async_insert=true&compress=lz4&dial_timeout=30s&secure=true",
		},
		{
			name: "Should respect connection parameters",
			fields: fields{
				Endpoint:         "clickhouse://127.0.0.1:9000?foo=bar",
				ConnectionParams: map[string]string{"secure": "true"},
			},
			wantChOptions: ChOptions{
				Secure: true,
			},
			args: args{},
			want: "clickhouse://127.0.0.1:9000/default?async_insert=true&foo=bar&secure=true",
		},
		{
			name: "support replace database in DSN to default database",
			fields: fields{
				Endpoint: "tcp://127.0.0.1:9000/otel",
			},
			args: args{
				database: defaultDatabase,
			},
			want: "tcp://127.0.0.1:9000/default?async_insert=true",
		},
		{
			name: "when config option is missing, preserve async_insert false in DSN",
			fields: fields{
				Endpoint: "tcp://127.0.0.1:9000?async_insert=false",
			},
			want: "tcp://127.0.0.1:9000/default?async_insert=false",
		},
		{
			name: "when config option is missing, preserve async_insert true in DSN",
			fields: fields{
				Endpoint: "tcp://127.0.0.1:9000?async_insert=true",
			},
			want: "tcp://127.0.0.1:9000/default?async_insert=true",
		},
		{
			name: "ignore config option when async_insert is present in connection params as false",
			fields: fields{
				Endpoint:         "tcp://127.0.0.1:9000?async_insert=false",
				ConnectionParams: map[string]string{"async_insert": "false"},
				AsyncInsert:      configBool(true),
			},

			want: "tcp://127.0.0.1:9000/default?async_insert=false",
		},
		{
			name: "ignore config option when async_insert is present in connection params as true",
			fields: fields{
				Endpoint:         "tcp://127.0.0.1:9000?async_insert=false",
				ConnectionParams: map[string]string{"async_insert": "true"},
				AsyncInsert:      configBool(false),
			},

			want: "tcp://127.0.0.1:9000/default?async_insert=true",
		},
		{
			name: "ignore config option when async_insert is present in DSN as false",
			fields: fields{
				Endpoint:    "tcp://127.0.0.1:9000?async_insert=false",
				AsyncInsert: configBool(true),
			},

			want: "tcp://127.0.0.1:9000/default?async_insert=false",
		},
		{
			name: "use async_insert true config option when it is not present in DSN",
			fields: fields{
				Endpoint:    "tcp://127.0.0.1:9000",
				AsyncInsert: configBool(true),
			},

			want: "tcp://127.0.0.1:9000/default?async_insert=true",
		},
		{
			name: "use async_insert false config option when it is not present in DSN",
			fields: fields{
				Endpoint:    "tcp://127.0.0.1:9000",
				AsyncInsert: configBool(false),
			},

			want: "tcp://127.0.0.1:9000/default?async_insert=false",
		},
		{
			name: "set async_insert to true when not present in config or DSN",
			fields: fields{
				Endpoint: "tcp://127.0.0.1:9000",
			},

			want: "tcp://127.0.0.1:9000/default?async_insert=true",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &Config{
				Endpoint:         tt.fields.Endpoint,
				Username:         tt.fields.Username,
				Password:         configopaque.String(tt.fields.Password),
				Database:         tt.fields.Database,
				ConnectionParams: tt.fields.ConnectionParams,
				AsyncInsert:      tt.fields.AsyncInsert,
			}
			got, err := cfg.buildDSN(tt.args.database)

			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr, "buildDSN(%v)", tt.args.database)
			} else {
				// Validate DSN
				opts, err := clickhouse.ParseDSN(got)
				assert.NoError(t, err)
				assert.Equalf(t, tt.wantChOptions.Secure, opts.TLS != nil, "TLSConfig is not nil")
				assert.Equalf(t, tt.wantChOptions.DialTimeout, opts.DialTimeout, "DialTimeout is not nil")
				if tt.wantChOptions.Compress != 0 {
					assert.Equalf(t, tt.wantChOptions.Compress, opts.Compression.Method, "Compress is not nil")
				}
				assert.Equalf(t, tt.want, got, "buildDSN(%v)", tt.args.database)
			}

		})
	}
}

func TestTableEngineConfigParsing(t *testing.T) {
	t.Parallel()
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)

	tests := []struct {
		id       component.ID
		expected string
	}{
		{
			id:       component.NewIDWithName(metadata.Type, "table-engine-empty"),
			expected: "MergeTree()",
		},
		{
			id:       component.NewIDWithName(metadata.Type, "table-engine-name-only"),
			expected: "ReplicatedReplacingMergeTree()",
		},
		{
			id:       component.NewIDWithName(metadata.Type, "table-engine-full"),
			expected: "ReplicatedReplacingMergeTree('/clickhouse/tables/{shard}/table_name', '{replica}', ver)",
		},
		{
			id:       component.NewIDWithName(metadata.Type, "table-engine-params-only"),
			expected: "MergeTree()",
		},
	}

	for _, tt := range tests {
		t.Run(tt.id.String(), func(t *testing.T) {
			factory := NewFactory()
			cfg := factory.CreateDefaultConfig()

			sub, err := cm.Sub(tt.id.String())
			require.NoError(t, err)
			require.NoError(t, component.UnmarshalConfig(sub, cfg))

			assert.NoError(t, component.ValidateConfig(cfg))
			assert.Equal(t, tt.expected, cfg.(*Config).TableEngineString())
		})
	}
}

func TestClusterString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "",
			expected: "",
		},
		{
			input:    "cluster_a_b",
			expected: "ON CLUSTER cluster_a_b",
		},
		{
			input:    "cluster a b",
			expected: "ON CLUSTER cluster a b",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("ClusterString case %d", i), func(t *testing.T) {
			cfg := createDefaultConfig()
			cfg.(*Config).Endpoint = defaultEndpoint
			cfg.(*Config).ClusterName = tt.input

			assert.NoError(t, component.ValidateConfig(cfg))
			assert.Equal(t, tt.expected, cfg.(*Config).ClusterString())
		})
	}
}
