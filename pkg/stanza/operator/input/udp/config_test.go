// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package udp

import (
	"path/filepath"
	"testing"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/operatortest"
)

func TestUnmarshal(t *testing.T) {
	operatortest.ConfigUnmarshalTests{
		DefaultConfig: NewConfig(),
		TestsFile:     filepath.Join(".", "testdata", "config.yaml"),
		Tests: []operatortest.ConfigUnmarshalTest{
			{
				Name:      "default",
				ExpectErr: false,
				Expect:    NewConfig(),
			},
			{
				Name:      "all",
				ExpectErr: false,
				Expect: func() *Config {
					cfg := NewConfig()
					cfg.ListenAddress = "10.0.0.1:9000"
					cfg.AddAttributes = true
					cfg.Encoding = "utf-8"
					cfg.SplitConfig.LineStartPattern = "ABC"
					cfg.SplitConfig.LineEndPattern = ""
					return cfg
				}(),
			},
		},
	}.Run(t)
}

func TestAsyncUnmarshal(t *testing.T) {
	operatortest.ConfigUnmarshalTests{
		DefaultConfig: NewConfig(),
		TestsFile:     filepath.Join(".", "testdata", "config.yaml"),
		Tests: []operatortest.ConfigUnmarshalTest{
			{
				Name:      "async-enabled",
				ExpectErr: false,
				Expect: func() *Config {
					cfg := NewConfig()
					cfg.ListenAddress = "10.0.0.1:9000"
					cfg.AddAttributes = true
					cfg.Encoding = "utf-8"
					cfg.SplitConfig.LineStartPattern = "ABC"
					cfg.SplitConfig.LineEndPattern = ""
					cfg.AsyncConcurrentMode = true
					cfg.FixedAsyncReaderRoutineCount = 2
					cfg.FixedAsyncProcessorRoutineCount = 3
					cfg.MaxAsyncQueueLength = 500
					cfg.MaxGracefulShutdownTimeInMS = 1000
					return cfg
				}(),
			},
			{
				Name:      "async-disabled",
				ExpectErr: false,
				Expect: func() *Config {
					cfg := NewConfig()
					cfg.ListenAddress = "10.0.0.1:9000"
					cfg.AddAttributes = true
					cfg.Encoding = "utf-8"
					cfg.SplitConfig.LineStartPattern = "ABC"
					cfg.SplitConfig.LineEndPattern = ""
					cfg.AsyncConcurrentMode = false
					return cfg
				}(),
			},
		},
	}.Run(t)
}