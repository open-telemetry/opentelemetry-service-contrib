// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tokenize // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/tokenize"

import (
	"bufio"

	"golang.org/x/text/encoding"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/trim"
)

// SplitterConfig consolidates MultilineConfig and FlusherConfig
type SplitterConfig struct {
	Flusher   FlusherConfig   `mapstructure:",squash,omitempty"`
	Multiline MultilineConfig `mapstructure:"multiline,omitempty"`
}

// NewSplitterConfig returns default SplitterConfig
func NewSplitterConfig() SplitterConfig {
	return SplitterConfig{
		Multiline: NewMultilineConfig(),
		Flusher:   FlusherConfig{Period: DefaultFlushPeriod},
	}
}

// Build builds bufio.SplitFunc based on the config
func (c *SplitterConfig) Build(enc encoding.Encoding, flushAtEOF bool, maxLogSize int, trimFunc trim.Func) (bufio.SplitFunc, error) {
	splitFunc, err := c.Multiline.Build(enc, flushAtEOF, maxLogSize, trimFunc)
	if err != nil {
		return nil, err
	}
	return c.Flusher.Wrap(splitFunc, trimFunc), nil
}
