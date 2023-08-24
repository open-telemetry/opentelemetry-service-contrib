// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"context"
	"fmt"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type MicroArguments[K any] struct {
	Duration ottl.DurationGetter[K] `ottlarg:"0"`
}

func NewMicroFactory[K any]() ottl.Factory[K] {
	return ottl.NewFactory("Micro", &MicroArguments[K]{}, createMicroFunction[K])
}
func createMicroFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	args, ok := oArgs.(*MicroArguments[K])

	if !ok {
		return nil, fmt.Errorf("MicroFactory args must be of type *MicroArguments[K]")
	}

	return Micro(args.Duration)
}

func Micro[K any](duration ottl.DurationGetter[K]) (ottl.ExprFunc[K], error) {
	return func(ctx context.Context, tCtx K) (interface{}, error) {
		d, err := duration.Get(ctx, tCtx)
		if err != nil {
			return nil, err
		}
		return d.Microseconds(), nil
	}, nil
}
