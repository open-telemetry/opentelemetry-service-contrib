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

package probabilisticsamplerprocessor

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/model/pdata"
)

func TestNewLogsProcessor(t *testing.T) {
	tests := []struct {
		name         string
		nextConsumer consumer.Logs
		cfg          *Config
		wantErr      bool
	}{
		{
			name: "nil_nextConsumer",
			cfg: &Config{
				ProcessorSettings:  config.NewProcessorSettings(config.NewComponentID(typeStr)),
				SamplingPercentage: 15.5,
			},
			wantErr: true,
		},
		{
			name:         "happy_path",
			nextConsumer: consumertest.NewNop(),
			cfg: &Config{
				ProcessorSettings:  config.NewProcessorSettings(config.NewComponentID(typeStr)),
				SamplingPercentage: 15.5,
			},
		},
		{
			name:         "happy_path_hash_seed",
			nextConsumer: consumertest.NewNop(),
			cfg: &Config{
				ProcessorSettings:  config.NewProcessorSettings(config.NewComponentID(typeStr)),
				SamplingPercentage: 13.33,
				HashSeed:           4321,
			},
		},
		{
			name:         "with_severity",
			nextConsumer: consumertest.NewNop(),
			cfg: &Config{
				ProcessorSettings:  config.NewProcessorSettings(config.NewComponentID(typeStr)),
				SamplingPercentage: 13.33,
				HashSeed:           4321,
				Severity: []severityPair{
					{"error", 90},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newLogsProcessor(tt.nextConsumer, tt.cfg)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
			}
		})
	}
}

func TestLogsSampling(t *testing.T) {
	tests := []struct {
		name     string
		cfg      *Config
		received int
	}{
		{
			name: "happy_path",
			cfg: &Config{
				ProcessorSettings:  config.NewProcessorSettings(config.NewComponentID(typeStr)),
				SamplingPercentage: 100,
			},
			received: 2,
		},
		{
			name: "nothing",
			cfg: &Config{
				ProcessorSettings:  config.NewProcessorSettings(config.NewComponentID(typeStr)),
				SamplingPercentage: 0,
			},
			received: 0,
		},
		{
			name: "half",
			cfg: &Config{
				ProcessorSettings:  config.NewProcessorSettings(config.NewComponentID(typeStr)),
				SamplingPercentage: 49,
			},
			received: 1,
		},
		{
			name: "nothing_except_errors",
			cfg: &Config{
				ProcessorSettings:  config.NewProcessorSettings(config.NewComponentID(typeStr)),
				SamplingPercentage: 0,
				Severity: []severityPair{
					{"error", 100},
				},
			},
			received: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sink := new(consumertest.LogsSink)
			processor, err := newLogsProcessor(sink, tt.cfg)
			require.NoError(t, err)
			logs := pdata.NewLogs()
			lr := logs.ResourceLogs().AppendEmpty().ScopeLogs().AppendEmpty().LogRecords()
			for i := 0; i < 5; i++ {
				first := lr.AppendEmpty()
				first.SetTimestamp(pdata.Timestamp(time.Unix(1649400860, 0).Unix()))
				first.SetSeverityNumber(pdata.SeverityNumberDEBUG)
				second := lr.AppendEmpty()
				second.SetTimestamp(pdata.Timestamp(time.Unix(12345555432, 0).Unix()))
				second.SetSeverityNumber(pdata.SeverityNumberERROR)
			}
			err = processor.ConsumeLogs(context.Background(), logs)
			require.NoError(t, err)
			sunk := sink.AllLogs()
			numReceived := 0
			if len(sunk) > 0 && sunk[0].ResourceLogs().Len() > 0 {
				numReceived = sunk[0].ResourceLogs().At(0).ScopeLogs().At(0).LogRecords().Len()
			}
			assert.Equal(t, tt.received*5, numReceived)
		})
	}
}
