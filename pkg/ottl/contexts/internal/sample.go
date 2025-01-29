// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pprofile"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

const (
	SampleContextName = "sample"
)

type SampleContext interface {
	GetSample() pprofile.Sample
}

func SamplePathGetSetter[K SampleContext](path ottl.Path[K]) (ottl.GetSetter[K], error) {
	if path == nil {
		return nil, FormatDefaultErrorMessage(SampleContextName, SampleContextName, "Sample", ProfileRef)
	}
	switch path.Name() {
	case "locations_length":
		return accessLocationsLength[K](), nil
	case "locations_start_index":
		return accessLocationsStartIndex[K](), nil
	case "attribute_indices":
		return accessSampleAttributeIndices[K](), nil
	case "timestamps_unix_nano":
		return accessTimestampsUnixNano[K](), nil
	case "value":
		return accessSampleValue[K](), nil
	default:
		return nil, FormatDefaultErrorMessage(path.Name(), path.String(), "Profile", ProfileRef)
	}
}

func accessLocationsLength[K SampleContext]() ottl.StandardGetSetter[K] {
	return ottl.StandardGetSetter[K]{
		Getter: func(_ context.Context, tCtx K) (any, error) {
			return tCtx.GetSample().LocationsLength(), nil
		},
		Setter: func(_ context.Context, tCtx K, val any) error {
			if v, ok := val.(int32); ok {
				tCtx.GetSample().SetLocationsLength(v)
			}
			return nil
		},
	}
}

func accessLocationsStartIndex[K SampleContext]() ottl.StandardGetSetter[K] {
	return ottl.StandardGetSetter[K]{
		Getter: func(_ context.Context, tCtx K) (any, error) {
			return tCtx.GetSample().LocationsStartIndex(), nil
		},
		Setter: func(_ context.Context, tCtx K, val any) error {
			if v, ok := val.(int32); ok {
				tCtx.GetSample().SetLocationsStartIndex(v)
			}
			return nil
		},
	}
}

func accessSampleAttributeIndices[K SampleContext]() ottl.StandardGetSetter[K] {
	return ottl.StandardGetSetter[K]{
		Getter: func(_ context.Context, tCtx K) (any, error) {
			return tCtx.GetSample().AttributeIndices(), nil
		},
		Setter: func(_ context.Context, tCtx K, val any) error {
			if v, ok := val.(pcommon.Int32Slice); ok {
				v.CopyTo(tCtx.GetSample().AttributeIndices())
			}
			return nil
		},
	}
}

func accessTimestampsUnixNano[K SampleContext]() ottl.StandardGetSetter[K] {
	return ottl.StandardGetSetter[K]{
		Getter: func(_ context.Context, tCtx K) (any, error) {
			return tCtx.GetSample().TimestampsUnixNano(), nil
		},
		Setter: func(_ context.Context, tCtx K, val any) error {
			if v, ok := val.(pcommon.UInt64Slice); ok {
				v.CopyTo(tCtx.GetSample().TimestampsUnixNano())
			}
			return nil
		},
	}
}

func accessSampleValue[K SampleContext]() ottl.StandardGetSetter[K] {
	return ottl.StandardGetSetter[K]{
		Getter: func(_ context.Context, tCtx K) (any, error) {
			return tCtx.GetSample().Value(), nil
		},
		Setter: func(_ context.Context, tCtx K, val any) error {
			if v, ok := val.(pcommon.Int64Slice); ok {
				v.CopyTo(tCtx.GetSample().Value())
			}
			return nil
		},
	}
}
