// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package retain // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/retain"

import (
	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

// Deprecated [v0.97.0] Use Factory.NewDefaultConfig instead.
func NewConfig() *Config {
	return NewFactory().NewDefaultConfig(operatorType.String()).(*Config)
}

// Deprecated [v0.97.0] Use Factory.NewDefaultConfig instead.
func NewConfigWithID(operatorID string) *Config {
	return NewFactory().NewDefaultConfig(operatorID).(*Config)
}

// Config is the configuration of a retain operator
type Config struct {
	helper.TransformerConfig `mapstructure:",squash"`
	Fields                   []entry.Field `mapstructure:"fields"`
}

// Deprecated [v0.97.0] Use Factory.CreateOperator instead.
func (c Config) Build(logger *zap.SugaredLogger) (operator.Operator, error) {
	set := component.TelemetrySettings{}
	if logger != nil {
		set.Logger = logger.Desugar()
	}
	return NewFactory().CreateOperator(&c, set)
}
