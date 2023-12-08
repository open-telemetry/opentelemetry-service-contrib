// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filter // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filter"

// Filter is an interface for matching values against a set of filters.
type Filter interface {
	// Matches returns true if the given value matches at least one
	// of the filters encapsulated by the Filter.
	Matches(any) bool
}
