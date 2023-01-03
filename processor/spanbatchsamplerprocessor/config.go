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

package spanbatchsamplerprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/spanbatchsamplerprocessor"

import (
	"fmt"

	"go.opentelemetry.io/collector/config"
)

// Config defines configuration for processor.
type Config struct {
	config.ProcessorSettings `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct
	// TokensPerBatch determines how many requests to send in each batch
	TokensPerBatch int  `mapstructure:"tokens"`
	MostPopular    bool `mapstructure:"most_popular"`
}

// Validate checks if the processor configuration is valid
func (cfg *Config) Validate() error {
	if cfg.TokensPerBatch <= 0 {
		return fmt.Errorf("tokens (%d) must be greater then zero", cfg.TokensPerBatch)
	}
	return nil
}
