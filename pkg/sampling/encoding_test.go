// Copyright The OpenTelemetry Authors
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

package sampling

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func mustNot[T any](t T, err error) error {
	if err == nil {
		return fmt.Errorf("expected an error, got nil")
	}
	return err
}

func TestValidProbabilityToTValue(t *testing.T) {
	require.Equal(t, "8", must(ProbabilityToTValue(0.5)))
	require.Equal(t, "00000000000001", must(ProbabilityToTValue(0x1p-56)))
	require.Equal(t, "55555555555554", must(ProbabilityToTValue(1/3.)))
	require.Equal(t, "54", must(ProbabilityToTValue(0x54p-8))) // 0x54p-8 is approximately 1/3
	require.Equal(t, "01", must(ProbabilityToTValue(0x1p-8)))
	require.Equal(t, "0", must(ProbabilityToTValue(0)))
}

func TestInvalidProbabilityToTValue(t *testing.T) {
	// Too small
	require.Error(t, mustNot(ProbabilityToTValue(0x1p-57)))
	require.Error(t, mustNot(ProbabilityToTValue(0x1p-57)))

	// Too big
	require.Error(t, mustNot(ProbabilityToTValue(1.1)))
	require.Error(t, mustNot(ProbabilityToTValue(1.1)))
}

func TestTValueToProbability(t *testing.T) {
	require.Equal(t, 0.5, must(TValueToProbability("8")))
	require.Equal(t, 0x444p-12, must(TValueToProbability("444")))
	require.Equal(t, 0.0, must(TValueToProbability("0")))

	// 0x55555554p-32 is very close to 1/3
	require.InEpsilon(t, 1/3., must(TValueToProbability("55555554")), 1e-9)
}

func TestProbabilityToThreshold(t *testing.T) {
	require.Equal(t,
		Threshold{0x1p+55},
		must(ProbabilityToThreshold(0.5)))
	require.Equal(t,
		Threshold{1},
		must(ProbabilityToThreshold(0x1p-56)))
	require.Equal(t,
		Threshold{0x100},
		must(ProbabilityToThreshold(0x100p-56)))
	require.Equal(t,
		Threshold{2},
		must(ProbabilityToThreshold(0x1p-55)))
	require.Equal(t,
		Threshold{1 / MinSamplingProb},
		must(ProbabilityToThreshold(1.0)))

	require.Equal(t,
		Threshold{0x555p-12 / MinSamplingProb},
		must(TValueToThreshold("555")))
	require.Equal(t,
		Threshold{0x123p-20 / MinSamplingProb},
		must(TValueToThreshold("00123")))
}

func TestShouldSample(t *testing.T) {
	// Test four boundary conditions for 50% sampling,
	thresh := must(ProbabilityToThreshold(0.5))
	// Smallest TraceID that should sample.
	require.True(t, thresh.ShouldSample(RandomnessFromTraceID(pcommon.TraceID{
		// 9 meaningless bytes
		0xee, 0xee, 0xee, 0xee, 0xee, 0xee, 0xee, 0xee, 0xee,
		0, // randomness starts here
		0, 0, 0, 0, 0, 0,
	})))
	// Largest TraceID that should sample.
	require.True(t, thresh.ShouldSample(RandomnessFromTraceID(pcommon.TraceID{
		// 9 meaningless bytes
		0xee, 0xee, 0xee, 0xee, 0xee, 0xee, 0xee, 0xee, 0xee,
		0x7f, // randomness starts here
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	})))
	// Smallest TraceID that should NOT sample.
	require.False(t, thresh.ShouldSample(RandomnessFromTraceID(pcommon.TraceID{
		// 9 meaningless bytes
		0xee, 0xee, 0xee, 0xee, 0xee, 0xee, 0xee, 0xee, 0xee,
		0x80, // randomness starts here
		0, 0, 0, 0, 0, 0,
	})))
	// Largest TraceID that should NOT sample.
	require.False(t, thresh.ShouldSample(RandomnessFromTraceID(pcommon.TraceID{
		// 9 meaningless bytes
		0xee, 0xee, 0xee, 0xee, 0xee, 0xee, 0xee, 0xee, 0xee,
		0xff, // randomness starts here
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	})))
}

// The two benchmarks below were used to choose the implementation for
// the Threshold type in this package.  The results indicate that it
// is faster to compare a 56-bit number than to compare as 7 element []byte.

type benchTIDs [1024]pcommon.TraceID

func (tids *benchTIDs) init() {
	for i := range tids {
		binary.BigEndian.PutUint64(tids[i][:8], rand.Uint64())
		binary.BigEndian.PutUint64(tids[i][8:], rand.Uint64())
	}
}

// BenchmarkThresholdCompareAsUint64-10    	1000000000	         0.4515 ns/op	       0 B/op	       0 allocs/op
func BenchmarkThresholdCompareAsUint64(b *testing.B) {
	var tids benchTIDs
	var comps [1024]uint64
	tids.init()
	for i := range comps {
		comps[i] = (rand.Uint64() % 0x1p+56) + 1
	}

	b.ReportAllocs()
	b.ResetTimer()
	yes := 0
	no := 0
	for i := 0; i < b.N; i++ {
		tid := tids[i%len(tids)]
		comp := comps[i%len(comps)]
		// Read 8 bytes, mask to 7 bytes
		val := binary.BigEndian.Uint64(tid[8:]) & (0x1p+56 - 1)

		if val < comp {
			yes++
		} else {
			no++
		}
	}
}

// BenchmarkThresholdCompareAsBytes-10     	528679580	         2.288 ns/op	       0 B/op	       0 allocs/op
func BenchmarkThresholdCompareAsBytes(b *testing.B) {
	var tids benchTIDs
	var comps [1024][7]byte
	tids.init()
	for i := range comps {
		var e8 [8]byte
		binary.BigEndian.PutUint64(e8[:], rand.Uint64())
		copy(comps[i][:], e8[1:])
	}

	b.ReportAllocs()
	b.ResetTimer()
	yes := 0
	no := 0
	for i := 0; i < b.N; i++ {
		if bytes.Compare(tids[i%len(tids)][9:], comps[i%len(comps)][:]) <= 0 {
			yes++
		} else {
			no++
		}
	}
}
