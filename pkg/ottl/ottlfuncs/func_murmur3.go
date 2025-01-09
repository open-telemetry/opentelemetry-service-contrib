// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"github.com/twmb/murmur3"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type murmur3Variant int

const (
	Murmur3Hash murmur3Variant = iota
	Murmur3Hash128
	Murmur3Hex
	Murmur3Hex128
)

func (v murmur3Variant) String() string {
	switch v {
	case Murmur3Hash:
		return "Murmur3Hash"
	case Murmur3Hash128:
		return "Murmur3Hash128"
	case Murmur3Hex:
		return "Murmur3Hex"
	case Murmur3Hex128:
		return "Murmur3Hex128"
	default:
		return fmt.Sprintf("Unknown(%d) Murmur3 variant", v)
	}
}

type Murmur3Arguments[K any] struct {
	Target ottl.StringGetter[K]
}

func NewMurmur3HashFactory[K any]() ottl.Factory[K] {
	return newMurmur3Factory[K](Murmur3Hash)
}

func NewMurmur3Hash128Factory[K any]() ottl.Factory[K] {
	return newMurmur3Factory[K](Murmur3Hash128)
}

func NewMurmur3HexFactory[K any]() ottl.Factory[K] {
	return newMurmur3Factory[K](Murmur3Hex)
}

func NewMurmur3Hex128Factory[K any]() ottl.Factory[K] {
	return newMurmur3Factory[K](Murmur3Hex128)
}

func newMurmur3Factory[K any](variant murmur3Variant) ottl.Factory[K] {
	return ottl.NewFactory(variant.String(), &Murmur3Arguments[K]{}, func(_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
		return createMurmur3Function[K](oArgs, variant)
	})
}

func createMurmur3Function[K any](oArgs ottl.Arguments, variant murmur3Variant) (ottl.ExprFunc[K], error) {
	args, ok := oArgs.(*Murmur3Arguments[K])
	if !ok {
		return nil, fmt.Errorf("%sFactory args must be of type *Murmur3Arguments[K]", variant.String())
	}

	return func(ctx context.Context, tCtx K) (any, error) {
		val, err := args.Target.Get(ctx, tCtx)
		if err != nil {
			return nil, err
		}

		switch variant {
		case Murmur3Hash:
			h := murmur3.Sum32([]byte(val))
			return int64(h), nil
		case Murmur3Hash128:
			h1, h2 := murmur3.Sum128([]byte(val))
			return []int64{int64(h1), int64(h2)}, nil
		// MurmurHash3 is sensitive to endianness
		// Hex returns the hexadecimal representation of the hash in little-endian
		case Murmur3Hex:
			h := murmur3.Sum32([]byte(val))
			b := make([]byte, 4)
			binary.LittleEndian.PutUint32(b, h)
			return hex.EncodeToString(b), nil
		case Murmur3Hex128:
			h1, h2 := murmur3.Sum128([]byte(val))
			b := make([]byte, 16)
			binary.LittleEndian.PutUint64(b[:8], h1)
			binary.LittleEndian.PutUint64(b[8:], h2)
			return hex.EncodeToString(b), nil
		default:
			return nil, fmt.Errorf("unknown murmur3 variant: %d", variant)
		}
	}, nil
}
