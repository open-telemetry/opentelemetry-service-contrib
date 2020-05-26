// Copyright 2020, OpenTelemetry Authors
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

package goldendataset

import (
	"io"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateTraces(t *testing.T) {
	random := io.Reader(rand.New(rand.NewSource(42)))
	rscSpans, err := GenerateResourceSpans("testdata/generated_pict_pairs_traces.txt",
		"testdata/generated_pict_pairs_spans.txt", random)
	assert.Nil(t, err)
	assert.Equal(t, 28, len(rscSpans))
}
