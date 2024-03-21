// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package probabilisticsamplerprocessor

import (
	"fmt"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/sampling"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

const (
	// These four can happen at runtime and be returned by
	// randomnessFromXXX()

	ErrInconsistentArrivingTValue samplerError = "inconsistent arriving t-value: span should not have been sampled"
	ErrMissingRandomness          samplerError = "missing randomness"
	ErrRandomnessInUse            samplerError = "log record has sampling randomness, equalizing or proportional mode recommended"
	ErrThresholdInUse             samplerError = "log record has sampling threshold, equalizing or proportional mode recommended"
)

const (
	// Hashing method: The constants below help translate user friendly percentages
	// to numbers direct used in sampling.
	numHashBucketsLg2     = 14
	numHashBuckets        = 0x4000 // Using a power of 2 to avoid division.
	bitMaskHashBuckets    = numHashBuckets - 1
	percentageScaleFactor = numHashBuckets / 100.0
)

// samplerErrors are conditions reported by the sampler that are somewhat
// ordinary and should log as info-level.
type samplerError string

var _ error = samplerError("")

func (s samplerError) Error() string {
	return string(s)
}

type SamplerMode string

const (
	HashSeed     SamplerMode = "hash_seed"
	Equalizing   SamplerMode = "equalizing"
	Proportional SamplerMode = "proportional"
	DefaultMode  SamplerMode = Proportional
	modeUnset    SamplerMode = ""
)

type randomnessNamer interface {
	randomness() sampling.Randomness
	policyName() string
}

type randomnessMethod sampling.Randomness

func (rm randomnessMethod) randomness() sampling.Randomness {
	return sampling.Randomness(rm)
}

type traceIDHashingMethod struct{ randomnessMethod }
type traceIDW3CSpecMethod struct{ randomnessMethod }
type samplingRandomnessMethod struct{ randomnessMethod }

type missingRandomnessMethod struct{}

func (rm missingRandomnessMethod) randomness() sampling.Randomness {
	return sampling.AllProbabilitiesRandomness
}

func (missingRandomnessMethod) policyName() string {
	return "missing_randomness"
}

type attributeHashingMethod struct {
	randomnessMethod
	attribute string
}

func (am attributeHashingMethod) policyName() string {
	return am.attribute
}

func (traceIDHashingMethod) policyName() string {
	return "trace_id_hash"
}

func (samplingRandomnessMethod) policyName() string {
	return "sampling_randomness"
}

func (traceIDW3CSpecMethod) policyName() string {
	return "trace_id_w3c"
}

var _ randomnessNamer = missingRandomnessMethod{}
var _ randomnessNamer = traceIDHashingMethod{}
var _ randomnessNamer = traceIDW3CSpecMethod{}
var _ randomnessNamer = samplingRandomnessMethod{}

func newMissingRandomnessMethod() randomnessNamer {
	return missingRandomnessMethod{}
}

func newSamplingRandomnessMethod(rnd sampling.Randomness) randomnessNamer {
	return samplingRandomnessMethod{randomnessMethod(rnd)}
}

func newTraceIDW3CSpecMethod(rnd sampling.Randomness) randomnessNamer {
	return traceIDW3CSpecMethod{randomnessMethod(rnd)}
}

func newTraceIDHashingMethod(rnd sampling.Randomness) randomnessNamer {
	return traceIDHashingMethod{randomnessMethod(rnd)}
}

func newAttributeHashingMethod(attribute string, rnd sampling.Randomness) randomnessNamer {
	return attributeHashingMethod{
		randomnessMethod: randomnessMethod(rnd),
		attribute:        attribute,
	}
}

type samplingCarrier interface {
	explicitRandomness() (randomnessNamer, bool)
	setExplicitRandomness(randomnessNamer)

	clearThreshold()
	threshold() (sampling.Threshold, bool)
	updateThreshold(sampling.Threshold) error

	reserialize() error
}

type dataSampler interface {
	// decide reports the result based on a probabilistic decision.
	decide(carrier samplingCarrier) sampling.Threshold

	// randomnessFromSpan extracts randomness and returns a carrier specific to traces data.
	randomnessFromSpan(s ptrace.Span) (randomness randomnessNamer, carrier samplingCarrier, err error)

	// randomnessFromLogRecord extracts randomness and returns a carrier specific to logs data.
	randomnessFromLogRecord(s plog.LogRecord) (randomness randomnessNamer, carrier samplingCarrier, err error)
}

var AllModes = []SamplerMode{HashSeed, Equalizing, Proportional}

func (sm *SamplerMode) UnmarshalText(in []byte) error {
	switch mode := SamplerMode(in); mode {
	case HashSeed,
		Equalizing,
		Proportional,
		modeUnset:
		*sm = mode
		return nil
	default:
		return fmt.Errorf("unsupported sampler mode %q", mode)
	}
}

// commonFields includes fields used in all sampler modes.
type commonFields struct {
	logger *zap.Logger
}

// hashingSampler is the original hash-based calculation.  It is an
// equalizing sampler with randomness calculation that matches the
// original implementation.  This hash-based implementation is limited
// to 14 bits of precision.
type hashingSampler struct {
	hashSeed        uint32
	tvalueThreshold sampling.Threshold

	// Logs only: name of attribute to obtain randomness
	logsRandomnessSourceAttribute string

	// Logs only: name of attribute to obtain randomness
	logsTraceIDEnabled bool

	consistentCommon
}

// consistentCommon implements update() for all samplers, which clears
// the sampling threshold when probability sampling decides false.
type consistentCommon struct {
	commonFields
}

// consistentTracestateCommon includes all except the legacy hash-based
// method, which overrides randomnessFromX.
type consistentTracestateCommon struct {
	// logsRandomnessSourceAttribute is used in non-strict mode
	// for logs data when no trace ID is available.
	logsRandomnessSourceAttribute string
	logsRandomnessHashSeed        uint32

	consistentCommon
}

// neverSampler always decides false.
type neverSampler struct {
	consistentTracestateCommon
}

// equalizingSampler adjusts thresholds absolutely.  Cannot be used with zero.
type equalizingSampler struct {
	// TraceID-randomness-based calculation
	tvalueThreshold sampling.Threshold

	consistentTracestateCommon
}

// proportionalSampler adjusts thresholds relatively.  Cannot be used with zero.
type proportionalSampler struct {
	// ratio in the range [2**-56, 1]
	ratio float64

	// prec is the precision in number of hex digits
	prec int

	consistentTracestateCommon
}

func isMissing(rnd randomnessNamer) bool {
	_, ok := rnd.(missingRandomnessMethod)
	return ok
}

// randomnessFromLogRecord (hashingSampler) uses a hash function over
// the TraceID
func (th *hashingSampler) randomnessFromLogRecord(l plog.LogRecord) (randomnessNamer, samplingCarrier, error) {
	rnd := newMissingRandomnessMethod()
	lrc, err := newLogRecordCarrier(l)

	if th.logsTraceIDEnabled {
		value := l.TraceID()
		// Note: this admits empty TraceIDs.
		rnd = newTraceIDHashingMethod(randomnessFromBytes(value[:], th.hashSeed))
	}

	if isMissing(rnd) && th.logsRandomnessSourceAttribute != "" {
		if value, ok := l.Attributes().Get(th.logsRandomnessSourceAttribute); ok {
			// Note: this admits zero-byte values.
			rnd = newAttributeHashingMethod(
				th.logsRandomnessSourceAttribute,
				randomnessFromBytes(getBytesFromValue(value), th.hashSeed),
			)
		}
	}

	if err != nil {
		// The sampling.randomness or sampling.threshold attributes
		// had a parse error, in this case.
		lrc = nil
	} else if _, hasRnd := lrc.explicitRandomness(); hasRnd {
		// If the log record contains a randomness value, do not set.
		err = ErrRandomnessInUse
	} else if _, hasTh := lrc.threshold(); hasTh {
		// If the log record contains a threshold value, do not set.
		err = ErrThresholdInUse
	} else if !isMissing(rnd) {
		// When no sampling information is already present and we have
		// calculated new randomness, add it to the record.
		lrc.setExplicitRandomness(rnd)
	}

	return rnd, lrc, err
}

func (ctc *consistentTracestateCommon) randomnessFromLogRecord(l plog.LogRecord) (randomnessNamer, samplingCarrier, error) {
	lrc, err := newLogRecordCarrier(l)
	rnd := newMissingRandomnessMethod()

	if err != nil {
		// Parse error in sampling.randomness or sampling.thresholdnil
		lrc = nil
	} else if rv, hasRnd := lrc.explicitRandomness(); hasRnd {
		rnd = rv
	} else if tid := l.TraceID(); !tid.IsEmpty() {
		rnd = newTraceIDW3CSpecMethod(sampling.TraceIDToRandomness(tid))
	} else {
		// The case of no TraceID remains.  Use the configured attribute.

		if ctc.logsRandomnessSourceAttribute == "" {
			// rnd continues to be missing
		} else if value, ok := l.Attributes().Get(ctc.logsRandomnessSourceAttribute); ok {
			rnd = newAttributeHashingMethod(
				ctc.logsRandomnessSourceAttribute,
				randomnessFromBytes(getBytesFromValue(value), ctc.logsRandomnessHashSeed),
			)
		}
	}

	return rnd, lrc, err
}

func randomnessFromBytes(b []byte, hashSeed uint32) sampling.Randomness {
	hashed32 := computeHash(b, hashSeed)
	hashed := uint64(hashed32 & bitMaskHashBuckets)

	// Ordinarily, hashed is compared against an acceptance
	// threshold i.e., sampled when hashed < scaledSamplerate,
	// which has the form R < T with T in [1, 2^14] and
	// R in [0, 2^14-1].
	//
	// Here, modify R to R' and T to T', so that the sampling
	// equation has identical form to the specification, i.e., T'
	// <= R', using:
	//
	//   T' = numHashBuckets-T
	//   R' = numHashBuckets-1-R
	//
	// As a result, R' has the correct most-significant 14 bits to
	// use in an R-value.
	rprime14 := uint64(numHashBuckets - 1 - hashed)

	// There are 18 unused bits from the FNV hash function.
	unused18 := uint64(hashed32 >> (32 - numHashBucketsLg2))
	mixed28 := unused18 ^ (unused18 << 10)

	// The 56 bit quantity here consists of, most- to least-significant:
	// - 14 bits: R' = numHashBuckets - 1 - hashed
	// - 28 bits: mixture of unused 18 bits
	// - 14 bits: original `hashed`.
	rnd56 := (rprime14 << 42) | (mixed28 << 14) | hashed

	// Note: by construction:
	// - OTel samplers make the same probabilistic decision with this r-value,
	// - only 14 out of 56 bits are used in the sampling decision,
	// - there are only 32 actual random bits.
	rnd, _ := sampling.UnsignedToRandomness(rnd56)
	return rnd
}

func (th *hashingSampler) randomnessFromSpan(s ptrace.Span) (randomnessNamer, samplingCarrier, error) {
	tid := s.TraceID()
	// Note: this admits empty TraceIDs.
	rnd := newTraceIDHashingMethod(randomnessFromBytes(tid[:], th.hashSeed))
	tsc := &tracestateCarrier{
		span: s,
	}

	var err error
	tsc.W3CTraceState, err = sampling.NewW3CTraceState(s.TraceState().AsRaw())
	if err != nil {
		return rnd, nil, err
	}

	// If the tracestate contains a proper R-value or T-value, we
	// have to leave it alone.  The user should not be using this
	// sampler mode if they are using specified forms of consistent
	// sampling in OTel.
	if _, has := tsc.explicitRandomness(); has {
		err = ErrRandomnessInUse
	} else if _, has := tsc.threshold(); has {
		err = ErrThresholdInUse
	} else {
		// When no sampling information is present, add a
		// Randomness value.
		tsc.setExplicitRandomness(rnd)
	}
	return rnd, tsc, err
}

func (ctc *consistentTracestateCommon) randomnessFromSpan(s ptrace.Span) (randomnessNamer, samplingCarrier, error) {
	rawts := s.TraceState().AsRaw()
	rnd := newMissingRandomnessMethod()
	tsc := &tracestateCarrier{
		span: s,
	}

	// Parse the arriving TraceState.
	var err error
	tsc.W3CTraceState, err = sampling.NewW3CTraceState(rawts)
	if err != nil {
		tsc = nil
	} else if rv, has := tsc.W3CTraceState.OTelValue().RValueRandomness(); has {
		// When the tracestate is OK and has r-value, use it.
		rnd = newSamplingRandomnessMethod(rv)
	} else if s.TraceID().IsEmpty() {
		// If the TraceID() is all zeros, which W3C calls an invalid TraceID.
		// rnd continues to be missing.
	} else {
		rnd = newTraceIDW3CSpecMethod(sampling.TraceIDToRandomness(s.TraceID()))
	}

	return rnd, tsc, err
}

func consistencyCheck(rnd randomnessNamer, carrier samplingCarrier, common commonFields) error {
	// Without randomness, do not check the threshold.
	if isMissing(rnd) {
		return ErrMissingRandomness
	}
	// Consistency check: if the TraceID is out of range, the
	// TValue is a lie.  If inconsistent, clear it and return an error.
	if tv, has := carrier.threshold(); has {
		if !tv.ShouldSample(rnd.randomness()) {
			// In case we fail open, the threshold is cleared as
			// recommended in the OTel spec.
			carrier.clearThreshold()
			return ErrInconsistentArrivingTValue
		}
	}

	return nil
}

// makeSample constructs a sampler. There are no errors, as the only
// potential error, out-of-range probability, is corrected automatically
// according to the README, which allows percents >100 to equal 100%.
//
// Extending this logic, we round very small probabilities up to the
// minimum supported value(s) which varies according to sampler mode.
func makeSampler(cfg *Config, common commonFields, isLogs bool) dataSampler {
	// README allows percents >100 to equal 100%.
	pct := cfg.SamplingPercentage
	if pct > 100 {
		pct = 100
	}
	mode := cfg.SamplerMode
	if mode == modeUnset {
		// Reasons to choose the legacy behavior include:
		// (a) having set the hash seed
		// (b) logs signal w/o trace ID source
		if cfg.HashSeed != 0 || (isLogs && cfg.AttributeSource != traceIDAttributeSource) {
			mode = HashSeed
		} else {
			mode = DefaultMode
		}
	}

	ccom := consistentCommon{
		commonFields: common,
	}
	ctcom := consistentTracestateCommon{
		logsRandomnessSourceAttribute: cfg.FromAttribute,
		logsRandomnessHashSeed:        cfg.HashSeed,
		consistentCommon:              ccom,
	}
	never := &neverSampler{
		consistentTracestateCommon: ctcom,
	}

	if pct == 0 {
		return never
	}
	// Note: Convert to float64 before dividing by 100, otherwise loss of precision.
	// If the probability is too small, round it up to the minimum.
	ratio := float64(pct) / 100
	// Like the pct > 100 test above, but for values too small to
	// express in 14 bits of precision.
	if ratio < sampling.MinSamplingProbability {
		ratio = sampling.MinSamplingProbability
	}

	switch mode {
	case Equalizing:
		// The error case below is ignored, we have rounded the probability so
		// that it is in-range
		threshold, _ := sampling.ProbabilityToThresholdWithPrecision(ratio, cfg.SamplingPrecision)

		return &equalizingSampler{
			tvalueThreshold: threshold,

			consistentTracestateCommon: ctcom,
		}

	case Proportional:
		return &proportionalSampler{
			ratio: ratio,
			prec:  cfg.SamplingPrecision,

			consistentTracestateCommon: ctcom,
		}

	default: // i.e., HashSeed

		// Note: the original hash function used in this code
		// is preserved to ensure consistency across updates.
		//
		//   uint32(pct * percentageScaleFactor)
		//
		// (a) carried out the multiplication in 32-bit precision
		// (b) rounded to zero instead of nearest.
		scaledSamplerate := uint32(pct * percentageScaleFactor)

		if scaledSamplerate == 0 {
			ccom.logger.Warn("probability rounded to zero", zap.Float32("percent", pct))
			return never
		}

		// Convert the accept threshold to a reject threshold,
		// then shift it into 56-bit value.
		reject := numHashBuckets - scaledSamplerate
		reject56 := uint64(reject) << 42

		threshold, _ := sampling.ThresholdFromUnsigned(reject56)

		return &hashingSampler{
			consistentCommon: ccom,
			tvalueThreshold:  threshold,

			// Logs specific:
			logsTraceIDEnabled:            cfg.AttributeSource == traceIDAttributeSource,
			logsRandomnessSourceAttribute: cfg.FromAttribute,
		}
	}
}

func (*neverSampler) decide(_ samplingCarrier) sampling.Threshold {
	return sampling.NeverSampleThreshold
}

func (th *hashingSampler) decide(carrier samplingCarrier) sampling.Threshold {
	return th.tvalueThreshold
}

func (te *equalizingSampler) decide(carrier samplingCarrier) sampling.Threshold {
	return te.tvalueThreshold
}

func (tp *proportionalSampler) decide(carrier samplingCarrier) sampling.Threshold {
	incoming := 1.0
	if tv, has := carrier.threshold(); has {
		incoming = tv.Probability()
	}

	// There is a potential here for the product probability to
	// underflow, which is checked here.
	threshold, err := sampling.ProbabilityToThresholdWithPrecision(incoming*tp.ratio, tp.prec)

	// Check the only known error condition.
	if err == sampling.ErrProbabilityRange {
		// Considered valid, a case where the sampling probability
		// has fallen below the minimum supported value and simply
		// becomes unsampled.
		return sampling.NeverSampleThreshold
	}
	return threshold
}
