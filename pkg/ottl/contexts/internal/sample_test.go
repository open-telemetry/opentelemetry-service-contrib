// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal"
import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pprofile"
)

// create Test_ProfilePathGetSetter
func Test_SamplePathGetSetter(t *testing.T) {
	// create tests
	tests := []struct {
		path string
		val  any
	}{
		{
			path: "locations_length",
			val:  int32(1),
		},
		{
			path: "locations_start_index",
			val:  int32(2),
		},
		{
			path: "attribute_indices",
			val:  createInt32Slice(3),
		},
		{
			path: "timestamps_unix_nano",
			val:  createUInt64Slice(4),
		},
		{
			path: "value",
			val:  createInt64Slice(5),
		},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			path := &TestPath[*sampleContext]{N: tt.path}

			sample := createSample()

			accessor, err := SamplePathGetSetter[*sampleContext](path)
			assert.NoError(t, err)

			err = accessor.Set(context.Background(), newSampleContext(sample), tt.val)
			assert.NoError(t, err)

			got, err := accessor.Get(context.Background(), newSampleContext(sample))
			assert.NoError(t, err)

			assert.Equal(t, tt.val, got)
		})
	}
}

func createSample() pprofile.Sample {
	sample := pprofile.NewSample()
	sample.AttributeIndices().Append(1)
	sample.SetLocationsLength(2)
	sample.SetLocationsStartIndex(3)
	sample.TimestampsUnixNano().Append(4)
	sample.Value().Append(5)
	return sample
}

type sampleContext struct {
	sample pprofile.Sample
}

func (p *sampleContext) GetSample() pprofile.Sample {
	return p.sample
}

func newSampleContext(sample pprofile.Sample) *sampleContext {
	return &sampleContext{sample: sample}
}

func createInt32Slice(n int32) pcommon.Int32Slice {
	sl := pcommon.NewInt32Slice()
	sl.Append(n)
	return sl
}

func createUInt64Slice(n uint64) pcommon.UInt64Slice {
	sl := pcommon.NewUInt64Slice()
	sl.Append(n)
	return sl
}

func createInt64Slice(n int64) pcommon.Int64Slice {
	sl := pcommon.NewInt64Slice()
	sl.Append(n)
	return sl
}
