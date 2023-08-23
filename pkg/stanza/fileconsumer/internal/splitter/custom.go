// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package splitter // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/splitter"

import (
	"bufio"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/tokenize"
)

type customFactory struct {
	Flusher  tokenize.FlusherConfig
	Splitter bufio.SplitFunc
}

var _ Factory = (*customFactory)(nil)

func NewCustomFactory(
	flusher tokenize.FlusherConfig,
	splitter bufio.SplitFunc) Factory {
	return &customFactory{
		Flusher:  flusher,
		Splitter: splitter,
	}
}

// Build builds Multiline Splitter struct
func (factory *customFactory) Build(_ int) (bufio.SplitFunc, error) {
	return factory.Flusher.Build().SplitFunc(factory.Splitter), nil
}
