// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package expr // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/expr"

import (
	"context"
)

// BoolExpr is an interface that allows matching a context K against a configuration of a match.
type BoolExpr[K any] interface {
	// Evaluate match context K against the configuration of ctx
	Eval(ctx context.Context, tCtx K) (bool, error)
}

// A BoolExpr wrapper that inverts the predicate
type notMatcher[K any] struct {
	matcher BoolExpr[K]
}

func (nm notMatcher[K]) Eval(ctx context.Context, tCtx K) (bool, error) {
	ret, err := nm.matcher.Eval(ctx, tCtx)
	return !ret, err
}

func Not[K any](matcher BoolExpr[K]) BoolExpr[K] {
	return notMatcher[K]{matcher: matcher}
}

// A BoolExpr wrapper that is the logical OR on a list of matches, `matchers`
type orMatcher[K any] struct {
	matchers []BoolExpr[K]
}

func (om orMatcher[K]) Eval(ctx context.Context, tCtx K) (bool, error) {
	for _, matcher := range om.matchers {
		ret, err := matcher.Eval(ctx, tCtx)
		if err != nil {
			return false, err
		}
		if ret {
			return true, nil
		}
	}
	return false, nil
}

// Create an Or BoolExpr from a list of matchers
func Or[K any](matchers ...BoolExpr[K]) BoolExpr[K] {
	switch len(matchers) {
	case 0:
		return nil
	case 1:
		return matchers[0]
	default:
		return orMatcher[K]{matchers: matchers}
	}
}

// A BoolExpr wrapper that is the logical AND on a list of matches, `matchers`
type andMatcher[K any] struct {
	matchers []BoolExpr[K]
}

func (am andMatcher[K]) Eval(ctx context.Context, tCtx K) (bool, error) {
	for _, matcher := range am.matchers {
		ret, err := matcher.Eval(ctx, tCtx)
		if err != nil {
			return false, err
		}
		if !ret {
			return false, nil
		}
	}
	return true, nil
}

// Create an And BoolExpr from a list of matchers
func And[K any](matchers ...BoolExpr[K]) BoolExpr[K] {
	switch len(matchers) {
	case 0:
		return nil
	case 1:
		return matchers[0]
	default:
		return andMatcher[K]{matchers: matchers}
	}
}
