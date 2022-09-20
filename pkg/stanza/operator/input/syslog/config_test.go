// Copyright The OpenTelemetry Authors
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

package syslog

import (
	"path/filepath"
	"testing"

	"go.opentelemetry.io/collector/config/configtls"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper/operatortest"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/tcp"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/udp"
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
				Name:      "tcp",
				ExpectErr: false,
				Expect: func() *Config {
					cfg := NewConfig()
					cfg.Protocol = "rfc3164"
					cfg.Location = "foo"
					cfg.TCP = &tcp.NewConfig().BaseConfig
					cfg.TCP.MaxLogSize = 1000000
					cfg.TCP.ListenAddress = "10.0.0.1:9000"
					cfg.TCP.AddAttributes = true
					cfg.TCP.Encoding = helper.NewEncodingConfig()
					cfg.TCP.Encoding.Encoding = "utf-16"
					cfg.TCP.Multiline = helper.NewMultilineConfig()
					cfg.TCP.Multiline.LineStartPattern = "ABC"
					cfg.TCP.TLS = &configtls.TLSServerSetting{
						TLSSetting: configtls.TLSSetting{
							CertFile: "foo",
							KeyFile:  "foo2",
							CAFile:   "foo3",
						},
						ClientCAFile: "foo4",
					}
					return cfg
				}(),
			},
			{
				Name:      "udp",
				ExpectErr: false,
				Expect: func() *Config {
					cfg := NewConfig()
					cfg.Protocol = "rfc5424"
					cfg.Location = "foo"
					cfg.UDP = &udp.NewConfig().BaseConfig
					cfg.UDP.ListenAddress = "10.0.0.1:9000"
					cfg.UDP.AddAttributes = true
					cfg.UDP.Encoding = helper.NewEncodingConfig()
					cfg.UDP.Encoding.Encoding = "utf-16"
					cfg.UDP.Multiline = helper.NewMultilineConfig()
					cfg.UDP.Multiline.LineStartPattern = "ABC"
					return cfg
				}(),
			},
		},
	}.Run(t)
}
