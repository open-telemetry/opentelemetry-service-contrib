// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sampling // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/sampling"

import (
	"errors"
	"strconv"
	"strings"
)

const (
	// MaxAdjustedCount is 2^56 i.e. 0x100000000000000 i.e., 1<<56.
	MaxAdjustedCount = 1 << 56

	// NumHexDigits is the number of hex digits equalling 56 bits.
	// This is the limit of sampling precision.
	NumHexDigits = 56 / hexBits

	hexBits = 4
	hexBase = 16
)

// Threshold represents an exact sampling probability using 56 bits of
// precision.  A Threshold expresses the number of spans, out of 2**56,
// that are rejected.
//
// These 56 bits are compared against 56 bits of randomness, either
// extracted from an R-value or a TraceID having the W3C-specified
// randomness bit set.
//
// Because Thresholds store 56 bits of information and floating point
// values store 52 bits of significand, some conversions between
// Threshold and probability values are lossy.  The kinds of loss that
// occur depend on where in the probability scale it happens, as the
// step between adjacent floating point values adjusts with the exponent.
type Threshold struct {
	// unsigned is in the range [0, MaxAdjustedCount]
	// - 0 represents always sampling (0 Random values are less-than)
	// - 1 represents sampling 1-in-(MaxAdjustedCount-1)
	// - MaxAdjustedCount represents always sampling 1-in-
	unsigned uint64
}

var (
	// ErrTValueSize is returned for t-values longer than NumHexDigits hex digits.
	ErrTValueSize = errors.New("t-value exceeds 14 hex digits")

	// ErrEmptyTValue indicates no t-value was found, i.e., no threshold available.
	ErrTValueEmpty = errors.New("t-value is empty")

	// AlwaysSampleThreshold represents 100% sampling.
	AlwaysSampleThreshold = Threshold{unsigned: 0}
)

// TValueToThreshold returns a Threshold.  Because TValue strings
// have trailing zeros omitted, this function performs the reverse.
func TValueToThreshold(s string) (Threshold, error) {
	if len(s) > NumHexDigits {
		return AlwaysSampleThreshold, ErrTValueSize
	}
	if len(s) == 0 {
		return AlwaysSampleThreshold, ErrTValueEmpty
	}

	// Having checked length above, there are no range errors
	// possible.  Parse the hex string to an unsigned value.
	unsigned, err := strconv.ParseUint(s, hexBase, 64)
	if err != nil {
		return AlwaysSampleThreshold, err // e.g. parse error
	}

	// The unsigned value requires shifting to account for the
	// trailing zeros that were omitted by the encoding (see
	// TValue for the reverse).  Compute the number to shift by:
	extendByHexZeros := NumHexDigits - len(s)
	return Threshold{
		unsigned: unsigned << (hexBits * extendByHexZeros),
	}, nil
}

// TValue encodes a threshold, which is a variable-length hex string
// up to 14 characters.  The empty string is returned for 100%
// sampling.
func (th Threshold) TValue() string {
	// Always-sample is a special case because TrimRight() below
	// will trim it to the empty string, which represents no t-value.
	if th == AlwaysSampleThreshold {
		return "0"
	}
	// For thresholds other than the extremes, format a full-width
	// (14 digit) unsigned value with leading zeros, then, remove
	// the trailing zeros.  Use the logic for (Randomness).RValue().
	digits := Randomness(th).RValue()

	// Remove trailing zeros.
	return strings.TrimRight(digits, "0")
}

// ShouldSample returns true when the span passes this sampler's
// consistent sampling decision.
func (th Threshold) ShouldSample(rnd Randomness) bool {
	return rnd.unsigned >= th.unsigned
}

// ThresholdGreater allows direct comparison of Threshold values.
// Greater thresholds equate with smaller sampling probabilities.
func ThresholdGreater(a, b Threshold) bool {
	return a.unsigned > b.unsigned
}

// ThresholdLessThan allows direct comparison of Threshold values.
// Smaller thresholds equate with greater sampling probabilities.
func ThresholdLessThan(a, b Threshold) bool {
	return a.unsigned < b.unsigned
}
