// Copyright 2021 OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build linux
// +build linux

package journaldreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/journaldreceiver"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/adapter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/journald"
)

const (
	typeStr   = "journald"
	stability = component.StabilityLevelAlpha
)

// NewFactory creates a factory for journald receiver
func NewFactory() component.ReceiverFactory {
	return adapter.NewFactory(ReceiverType{}, stability)
}

// ReceiverType implements adapter.LogReceiverType
// to create a journald receiver
type ReceiverType struct{}

// Type is the receiver type
func (f ReceiverType) Type() config.Type {
	return typeStr
}

// CreateDefaultConfig creates a config with type and version
func (f ReceiverType) CreateDefaultConfig() config.Receiver {
	return &JournaldConfig{
		BaseConfig: adapter.BaseConfig{
			ReceiverSettings: config.NewReceiverSettings(config.NewComponentID(typeStr)),
			Operators:        adapter.OperatorConfigs{},
		},
		InputConfig: *journald.NewConfig(),
	}
}

// BaseConfig gets the base config from config, for now
func (f ReceiverType) BaseConfig(cfg config.Receiver) adapter.BaseConfig {
	return cfg.(*JournaldConfig).BaseConfig
}

// JournaldConfig defines configuration for the journald receiver
type JournaldConfig struct {
	adapter.BaseConfig `mapstructure:",squash"`
	InputConfig        journald.Config `mapstructure:",squash"`
}

// DecodeInputConfig unmarshals the input operator
func (f ReceiverType) DecodeInputConfig(cfg config.Receiver) (*operator.Config, error) {
	logConfig := cfg.(*JournaldConfig)
	return &operator.Config{Builder: &logConfig.InputConfig}, nil
}
