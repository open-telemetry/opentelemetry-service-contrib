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

package fileconsumer // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer"

import (
	"bufio"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

type splitterFactory interface {
	Build(maxLogSize int) (*helper.Splitter, error)
}

type multilineSplitterFactory struct {
	*helper.SplitterConfig
}

var _ splitterFactory = (*multilineSplitterFactory)(nil)

func newMultilineSplitterFactory(splitter helper.SplitterConfig) *multilineSplitterFactory {
	return &multilineSplitterFactory{
		SplitterConfig: &splitter,
	}

}

// Build builds Multiline Splitter struct
func (factory *multilineSplitterFactory) Build(maxLogSize int) (*helper.Splitter, error) {
	return factory.SplitterConfig.Build(false, maxLogSize)
}

type customizeSplitterFactory struct {
	*helper.SplitterConfig
	SplitFunc bufio.SplitFunc
}

var _ splitterFactory = (*customizeSplitterFactory)(nil)

func newCustomizeSplitterFactory(
	splitter helper.SplitterConfig,
	splitFunc bufio.SplitFunc) *customizeSplitterFactory {
	return &customizeSplitterFactory{
		SplitterConfig: &splitter,
		SplitFunc:      splitFunc,
	}
}

// Build builds Multiline Splitter struct
func (factory *customizeSplitterFactory) Build(maxLogSize int) (*helper.Splitter, error) {
	return &helper.Splitter{
		Encoding:  helper.Encoding{},
		SplitFunc: factory.SplitFunc,
	}, nil
}
