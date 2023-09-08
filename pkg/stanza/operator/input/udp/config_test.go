// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package udp

import (
	"path/filepath"
	"testing"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/operatortest"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/split"
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
					cfg.Multiline = split.NewMultilineConfig()
					cfg.Multiline.LineStartPattern = "ABC"
					return cfg
				}(),
			},
		},
	}.Run(t)
}
