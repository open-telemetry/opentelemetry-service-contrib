// Copyright 2018, OpenCensus Authors
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

package main

import (
	"testing"

	yaml "gopkg.in/yaml.v2"

	"github.com/census-instrumentation/opencensus-service/exporter/exporterparser"
)

// Issue #233: Zipkin receiver and exporter loopback detection
// would mistakenly report that "localhost:9410" and "localhost:9411"
// were equal, due to a mistake in parsing out their addresses,
// but also after IP resolution, the equivalence of ports was not being
// checked.
func TestZipkinReceiverExporterLogicalConflictChecks(t *testing.T) {
	regressionYAML := []byte(`
receivers:
    zipkin:
        address: "localhost:9410"

exporters:
    zipkin:
        endpoint: "http://localhost:9411/api/v2/spans"
`)

	cfg, err := parseOCAgentConfig(regressionYAML)
	if err != nil {
		t.Fatalf("Unexpected YAML parse error: %v", err)
	}
	if err := cfg.checkLogicalConflicts(regressionYAML); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if g, w := cfg.Receivers.Zipkin.Address, "localhost:9410"; g != w {
		t.Errorf("Receivers.Zipkin.EndpointURL mismatch\nGot: %s\nWant:%s", g, w)
	}

	var ecfg struct {
		Exporters *struct {
			Zipkin *exporterparser.ZipkinConfig `yaml:"zipkin"`
		} `yaml:"exporters"`
	}
	_ = yaml.Unmarshal(regressionYAML, &ecfg)
	if g, w := ecfg.Exporters.Zipkin.EndpointURL(), "http://localhost:9411/api/v2/spans"; g != w {
		t.Errorf("Exporters.Zipkin.EndpointURL mismatch\nGot: %s\nWant:%s", g, w)
	}
}
