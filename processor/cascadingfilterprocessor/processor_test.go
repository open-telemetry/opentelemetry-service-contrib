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

package cascadingfilterprocessor

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/consumer/pdata"
	tracetranslator "go.opentelemetry.io/collector/translator/trace"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/cascadingfilterprocessor/config"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/cascadingfilterprocessor/idbatcher"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/cascadingfilterprocessor/sampling"
)

const (
	defaultTestDecisionWait = 30 * time.Second
)

var testPolicy = []config.PolicyCfg{{
	Name:           "test-policy",
	SpansPerSecond: 1000,
}}

func TestSequentialTraceArrival(t *testing.T) {
	traceIds, batches := generateIdsAndBatches(128)
	cfg := config.Config{
		DecisionWait:            defaultTestDecisionWait,
		NumTraces:               uint64(2 * len(traceIds)),
		ExpectedNewTracesPerSec: 64,
		PolicyCfgs:              testPolicy,
	}
	sp, _ := newTraceProcessor(zap.NewNop(), consumertest.NewTracesNop(), cfg)
	tsp := sp.(*cascadingFilterSpanProcessor)
	for _, batch := range batches {
		tsp.ConsumeTraces(context.Background(), batch)
	}

	for i := range traceIds {
		d, ok := tsp.idToTrace.Load(traceKey(traceIds[i].Bytes()))
		require.True(t, ok, "Missing expected traceId")
		v := d.(*sampling.TraceData)
		require.Equal(t, int64(i+1), v.SpanCount, "Incorrect number of spans for entry %d", i)
	}
}

func TestConcurrentTraceArrival(t *testing.T) {
	traceIds, batches := generateIdsAndBatches(128)

	var wg sync.WaitGroup
	cfg := config.Config{
		DecisionWait:            defaultTestDecisionWait,
		NumTraces:               uint64(2 * len(traceIds)),
		ExpectedNewTracesPerSec: 64,
		PolicyCfgs:              testPolicy,
	}
	sp, _ := newTraceProcessor(zap.NewNop(), consumertest.NewTracesNop(), cfg)
	tsp := sp.(*cascadingFilterSpanProcessor)
	for _, batch := range batches {
		// Add the same traceId twice.
		wg.Add(2)
		go func(td pdata.Traces) {
			tsp.ConsumeTraces(context.Background(), td)
			wg.Done()
		}(batch)
		go func(td pdata.Traces) {
			tsp.ConsumeTraces(context.Background(), td)
			wg.Done()
		}(batch)
	}

	wg.Wait()

	for i := range traceIds {
		d, ok := tsp.idToTrace.Load(traceKey(traceIds[i].Bytes()))
		require.True(t, ok, "Missing expected traceId")
		v := d.(*sampling.TraceData)
		require.Equal(t, int64(i+1)*2, v.SpanCount, "Incorrect number of spans for entry %d", i)
	}
}

func TestSequentialTraceMapSize(t *testing.T) {
	traceIds, batches := generateIdsAndBatches(210)
	const maxSize = 100
	cfg := config.Config{
		DecisionWait:            defaultTestDecisionWait,
		NumTraces:               uint64(maxSize),
		ExpectedNewTracesPerSec: 64,
		PolicyCfgs:              testPolicy,
	}
	sp, _ := newTraceProcessor(zap.NewNop(), consumertest.NewTracesNop(), cfg)
	tsp := sp.(*cascadingFilterSpanProcessor)
	for _, batch := range batches {
		tsp.ConsumeTraces(context.Background(), batch)
	}

	// On sequential insertion it is possible to know exactly which traces should be still on the map.
	for i := 0; i < len(traceIds)-maxSize; i++ {
		_, ok := tsp.idToTrace.Load(traceKey(traceIds[i].Bytes()))
		require.False(t, ok, "Found unexpected traceId[%d] still on map (id: %v)", i, traceIds[i])
	}
}

func TestConcurrentTraceMapSize(t *testing.T) {
	_, batches := generateIdsAndBatches(210)
	const maxSize = 100
	var wg sync.WaitGroup
	cfg := config.Config{
		DecisionWait:            defaultTestDecisionWait,
		NumTraces:               uint64(maxSize),
		ExpectedNewTracesPerSec: 64,
		PolicyCfgs:              testPolicy,
	}
	sp, _ := newTraceProcessor(zap.NewNop(), consumertest.NewTracesNop(), cfg)
	tsp := sp.(*cascadingFilterSpanProcessor)
	for _, batch := range batches {
		wg.Add(1)
		go func(td pdata.Traces) {
			tsp.ConsumeTraces(context.Background(), td)
			wg.Done()
		}(batch)
	}

	wg.Wait()

	// Since we can't guarantee the order of insertion the only thing that can be checked is
	// if the number of traces on the map matches the expected value.
	cnt := 0
	tsp.idToTrace.Range(func(_ interface{}, _ interface{}) bool {
		cnt++
		return true
	})
	require.Equal(t, maxSize, cnt, "Incorrect traces count on idToTrace")
}

func TestSamplingPolicyTypicalPath(t *testing.T) {
	const maxSize = 100
	const decisionWaitSeconds = 5
	// For this test explicitly control the timer calls and batcher, and set a mock
	// sampling policy evaluator.
	msp := new(consumertest.TracesSink)
	mpe := &mockPolicyEvaluator{}
	mtt := &manualTTicker{}
	tsp := &cascadingFilterSpanProcessor{
		ctx:               context.Background(),
		nextConsumer:      msp,
		maxNumTraces:      maxSize,
		logger:            zap.NewNop(),
		decisionBatcher:   newSyncIDBatcher(decisionWaitSeconds),
		policies:          []*Policy{{Name: "mock-policy", Evaluator: mpe, ctx: context.TODO()}},
		deleteChan:        make(chan traceKey, maxSize),
		policyTicker:      mtt,
		maxSpansPerSecond: 10000,
	}

	_, batches := generateIdsAndBatches(210)
	currItem := 0
	numSpansPerBatchWindow := 10
	// First evaluations shouldn't have anything to evaluate, until decision wait time passed.
	for evalNum := 0; evalNum < decisionWaitSeconds; evalNum++ {
		for ; currItem < numSpansPerBatchWindow*(evalNum+1); currItem++ {
			tsp.ConsumeTraces(context.Background(), batches[currItem])
			require.True(t, mtt.Started, "Time ticker was expected to have started")
		}
		tsp.samplingPolicyOnTick()
		require.False(
			t,
			msp.SpansCount() != 0 || mpe.EvaluationCount != 0,
			"policy for initial items was evaluated before decision wait period",
		)
	}

	// Now the first batch that waited the decision period.
	mpe.NextDecision = sampling.Sampled
	tsp.samplingPolicyOnTick()
	require.False(
		t,
		msp.SpansCount() == 0 || mpe.EvaluationCount == 0,
		"policy should have been evaluated totalspans == %d and evaluationcount == %d",
		msp.SpansCount(),
		mpe.EvaluationCount,
	)

	require.Equal(t, numSpansPerBatchWindow, msp.SpansCount(), "not all spans of first window were accounted for")

	// Late span of a sampled trace should be sent directly down the pipeline exporter
	tsp.ConsumeTraces(context.Background(), batches[0])
	expectedNumWithLateSpan := numSpansPerBatchWindow + 1
	require.Equal(t, expectedNumWithLateSpan, msp.SpansCount(), "late span was not accounted for")
	require.Equal(t, 1, mpe.LateArrivingSpansCount, "policy was not notified of the late span")
}

func TestSamplingMultiplePolicies(t *testing.T) {
	const maxSize = 100
	const decisionWaitSeconds = 5
	// For this test explicitly control the timer calls and batcher, and set a mock
	// sampling policy evaluator.
	msp := new(consumertest.TracesSink)
	mpe1 := &mockPolicyEvaluator{}
	mpe2 := &mockPolicyEvaluator{}
	mtt := &manualTTicker{}
	tsp := &cascadingFilterSpanProcessor{
		ctx:             context.Background(),
		nextConsumer:    msp,
		maxNumTraces:    maxSize,
		logger:          zap.NewNop(),
		decisionBatcher: newSyncIDBatcher(decisionWaitSeconds),
		policies: []*Policy{
			{
				Name: "policy-1", Evaluator: mpe1, ctx: context.TODO(),
			},
			{
				Name: "policy-2", Evaluator: mpe2, ctx: context.TODO(),
			}},
		deleteChan:        make(chan traceKey, maxSize),
		policyTicker:      mtt,
		maxSpansPerSecond: 10000,
	}

	_, batches := generateIdsAndBatches(210)
	currItem := 0
	numSpansPerBatchWindow := 10
	// First evaluations shouldn't have anything to evaluate, until decision wait time passed.
	for evalNum := 0; evalNum < decisionWaitSeconds; evalNum++ {
		for ; currItem < numSpansPerBatchWindow*(evalNum+1); currItem++ {
			tsp.ConsumeTraces(context.Background(), batches[currItem])
			require.True(t, mtt.Started, "Time ticker was expected to have started")
		}
		tsp.samplingPolicyOnTick()
		require.False(
			t,
			msp.SpansCount() != 0 || mpe1.EvaluationCount != 0 || mpe2.EvaluationCount != 0,
			"policy for initial items was evaluated before decision wait period",
		)
	}

	// Both policies will decide to sample
	mpe1.NextDecision = sampling.Sampled
	mpe2.NextDecision = sampling.Sampled
	tsp.samplingPolicyOnTick()
	require.False(
		t,
		msp.SpansCount() == 0 || mpe1.EvaluationCount == 0 || mpe2.EvaluationCount == 0,
		"policy should have been evaluated totalspans == %d and evaluationcount(1) == %d and evaluationcount(2) == %d",
		msp.SpansCount(),
		mpe1.EvaluationCount,
		mpe2.EvaluationCount,
	)

	require.Equal(t, numSpansPerBatchWindow, msp.SpansCount(), "nextConsumer should've been called with exactly 1 batch of spans")

	// Late span of a sampled trace should be sent directly down the pipeline exporter
	tsp.ConsumeTraces(context.Background(), batches[0])
	expectedNumWithLateSpan := numSpansPerBatchWindow + 1
	require.Equal(t, expectedNumWithLateSpan, msp.SpansCount(), "late span was not accounted for")
	require.Equal(t, 1, mpe1.LateArrivingSpansCount, "1st policy was not notified of the late span")
	require.Equal(t, 0, mpe2.LateArrivingSpansCount, "2nd policy should not have been notified of the late span")
}

func TestSamplingPolicyDecisionNotSampled(t *testing.T) {
	const maxSize = 100
	const decisionWaitSeconds = 5
	// For this test explicitly control the timer calls and batcher, and set a mock
	// sampling policy evaluator.
	msp := new(consumertest.TracesSink)
	mpe := &mockPolicyEvaluator{}
	mtt := &manualTTicker{}
	tsp := &cascadingFilterSpanProcessor{
		ctx:               context.Background(),
		nextConsumer:      msp,
		maxNumTraces:      maxSize,
		logger:            zap.NewNop(),
		decisionBatcher:   newSyncIDBatcher(decisionWaitSeconds),
		policies:          []*Policy{{Name: "mock-policy", Evaluator: mpe, ctx: context.TODO()}},
		deleteChan:        make(chan traceKey, maxSize),
		policyTicker:      mtt,
		maxSpansPerSecond: 10000,
	}

	_, batches := generateIdsAndBatches(210)
	currItem := 0
	numSpansPerBatchWindow := 10
	// First evaluations shouldn't have anything to evaluate, until decision wait time passed.
	for evalNum := 0; evalNum < decisionWaitSeconds; evalNum++ {
		for ; currItem < numSpansPerBatchWindow*(evalNum+1); currItem++ {
			tsp.ConsumeTraces(context.Background(), batches[currItem])
			require.True(t, mtt.Started, "Time ticker was expected to have started")
		}
		tsp.samplingPolicyOnTick()
		require.False(
			t,
			msp.SpansCount() != 0 || mpe.EvaluationCount != 0,
			"policy for initial items was evaluated before decision wait period",
		)
	}

	// Now the first batch that waited the decision period.
	mpe.NextDecision = sampling.NotSampled
	tsp.samplingPolicyOnTick()
	require.EqualValues(t, 0, msp.SpansCount(), "exporter should have received zero spans")
	require.EqualValues(t, 4, mpe.EvaluationCount, "policy should have been evaluated 4 times")

	// Late span of a non-sampled trace should be ignored
	tsp.ConsumeTraces(context.Background(), batches[0])
	require.Equal(t, 0, msp.SpansCount())
	require.Equal(t, 1, mpe.LateArrivingSpansCount, "policy was not notified of the late span")

	mpe.NextDecision = sampling.Unspecified
	mpe.NextError = errors.New("mock policy error")
	tsp.samplingPolicyOnTick()
	require.EqualValues(t, 0, msp.SpansCount(), "exporter should have received zero spans")
	require.EqualValues(t, 6, mpe.EvaluationCount, "policy should have been evaluated 6 times")

	// Late span of a non-sampled trace should be ignored
	tsp.ConsumeTraces(context.Background(), batches[0])
	require.Equal(t, 0, msp.SpansCount())
	require.Equal(t, 2, mpe.LateArrivingSpansCount, "policy was not notified of the late span")
}

func TestMultipleBatchesAreCombinedIntoOne(t *testing.T) {
	const maxSize = 100
	const decisionWaitSeconds = 1
	// For this test explicitly control the timer calls and batcher, and set a mock
	// sampling policy evaluator.
	msp := new(consumertest.TracesSink)
	mpe := &mockPolicyEvaluator{}
	mtt := &manualTTicker{}
	tsp := &cascadingFilterSpanProcessor{
		ctx:               context.Background(),
		nextConsumer:      msp,
		maxNumTraces:      maxSize,
		logger:            zap.NewNop(),
		decisionBatcher:   newSyncIDBatcher(decisionWaitSeconds),
		policies:          []*Policy{{Name: "mock-policy", Evaluator: mpe, ctx: context.TODO()}},
		deleteChan:        make(chan traceKey, maxSize),
		policyTicker:      mtt,
		maxSpansPerSecond: 10000,
	}

	mpe.NextDecision = sampling.Sampled

	traceIds, batches := generateIdsAndBatches(3)
	for _, batch := range batches {
		require.NoError(t, tsp.ConsumeTraces(context.Background(), batch))
	}

	tsp.samplingPolicyOnTick()
	tsp.samplingPolicyOnTick()

	require.EqualValues(t, 3, len(msp.AllTraces()), "There should be three batches, one for each trace")

	expectedSpanIds := make(map[int][]pdata.SpanID)
	expectedSpanIds[0] = []pdata.SpanID{
		pdata.NewSpanID(tracetranslator.UInt64ToByteSpanID(uint64(1))),
	}
	expectedSpanIds[1] = []pdata.SpanID{
		pdata.NewSpanID(tracetranslator.UInt64ToByteSpanID(uint64(2))),
		pdata.NewSpanID(tracetranslator.UInt64ToByteSpanID(uint64(3))),
	}
	expectedSpanIds[2] = []pdata.SpanID{
		pdata.NewSpanID(tracetranslator.UInt64ToByteSpanID(uint64(4))),
		pdata.NewSpanID(tracetranslator.UInt64ToByteSpanID(uint64(5))),
		pdata.NewSpanID(tracetranslator.UInt64ToByteSpanID(uint64(6))),
	}

	receivedTraces := msp.AllTraces()
	for i, traceID := range traceIds {
		trace := findTrace(receivedTraces, traceID)
		require.NotNil(t, trace, "Trace was not received. TraceId %s", traceID.HexString())
		require.EqualValues(t, i+1, trace.SpanCount(), "The trace should have all of its spans in a single batch")

		expected := expectedSpanIds[i]
		got := collectSpanIds(trace)

		// might have received out of order, sort for comparison
		sort.Slice(got, func(i, j int) bool {
			a := tracetranslator.BytesToInt64SpanID(got[i].Bytes())
			b := tracetranslator.BytesToInt64SpanID(got[j].Bytes())
			return a < b
		})

		require.EqualValues(t, expected, got)
	}
}

func TestExceedingTheMaximumNumberOfTraces(t *testing.T) {
	testMaximumNumberOfTraces(t, 2, 410)
}

func TestNotExceedingTheMaximumNumberOfTraces(t *testing.T) {
	testMaximumNumberOfTraces(t, 5, 400)
}

func testMaximumNumberOfTraces(t *testing.T, maxSize uint64, expectedNumberOfSpans int) {
	const decisionWaitSeconds = 1
	const numOfBatches = 10
	const numOfSpansInBatch = 10
	msp := new(consumertest.TracesSink)
	mpe := &mockPolicyEvaluator{
		NextDecision: sampling.Sampled,
	}
	mtt := &manualTTicker{}
	tsp := &cascadingFilterSpanProcessor{
		ctx:               context.Background(),
		nextConsumer:      msp,
		maxNumTraces:      maxSize,
		logger:            zap.NewNop(),
		decisionBatcher:   newSyncIDBatcher(decisionWaitSeconds),
		policies:          []*Policy{{Name: "mock-policy", Evaluator: mpe, ctx: context.TODO()}},
		deleteChan:        make(chan traceKey, maxSize),
		policyTicker:      mtt,
		maxSpansPerSecond: 10000,
	}

	batches1 := generateBatchesForGivenId(1, numOfBatches, numOfSpansInBatch)
	batches2 := generateBatchesForGivenId(2, numOfBatches, numOfSpansInBatch)
	batches3 := generateBatchesForGivenId(3, numOfBatches, numOfSpansInBatch)
	batches4 := generateBatchesForGivenId(4, numOfBatches, numOfSpansInBatch)
	batches5 := generateBatchesForGivenId(5, numOfBatches, numOfSpansInBatch)
	for i := 0; i < numOfBatches; i++ {
		tsp.ConsumeTraces(context.Background(), batches1[i])
		tsp.samplingPolicyOnTick()
	}
	for i := 0; i < numOfBatches; i++ {
		tsp.ConsumeTraces(context.Background(), batches2[i])
		tsp.samplingPolicyOnTick()
	}
	require.EqualValues(t, 200, msp.SpansCount(), "exporter should have received 200 spans")
	mpe.NextDecision = sampling.NotSampled
	for i := 0; i < numOfBatches; i++ {
		tsp.ConsumeTraces(context.Background(), batches3[i])
		tsp.samplingPolicyOnTick()
	}
	mpe.NextDecision = sampling.Sampled
	require.EqualValues(t, 200, msp.SpansCount(), "exporter should have received 200 spans")
	for i := 0; i < numOfBatches; i++ {
		tsp.ConsumeTraces(context.Background(), batches4[i])
		tsp.samplingPolicyOnTick()
	}
	for i := 0; i < numOfBatches; i++ {
		tsp.ConsumeTraces(context.Background(), batches5[i])
		tsp.samplingPolicyOnTick()
	}
	currentExpectedNumberOfSpans := 4*numOfBatches*numOfSpansInBatch
	require.EqualValues(t, currentExpectedNumberOfSpans, msp.SpansCount(), fmt.Sprintf("exporter should have received %d spans", currentExpectedNumberOfSpans))
	time.Sleep(2 * time.Second)
	tsp.ConsumeTraces(context.Background(), batches3[0])
	tsp.samplingPolicyOnTick()
	tsp.samplingPolicyOnTick()
	require.EqualValues(t, expectedNumberOfSpans, msp.SpansCount(), fmt.Sprintf("exporter should have received %d spans", expectedNumberOfSpans))
}

func collectSpanIds(trace *pdata.Traces) []pdata.SpanID {
	spanIDs := make([]pdata.SpanID, 0)

	for i := 0; i < trace.ResourceSpans().Len(); i++ {
		ilss := trace.ResourceSpans().At(i).InstrumentationLibrarySpans()

		for j := 0; j < ilss.Len(); j++ {
			ils := ilss.At(j)

			for k := 0; k < ils.Spans().Len(); k++ {
				span := ils.Spans().At(k)
				spanIDs = append(spanIDs, span.SpanID())
			}
		}
	}

	return spanIDs
}

func findTrace(a []pdata.Traces, traceID pdata.TraceID) *pdata.Traces {
	for _, batch := range a {
		id := batch.ResourceSpans().At(0).InstrumentationLibrarySpans().At(0).Spans().At(0).TraceID()
		if traceID.Bytes() == id.Bytes() {
			return &batch
		}
	}
	return nil
}

func generateIdsAndBatches(numIds int) ([]pdata.TraceID, []pdata.Traces) {
	traceIds := make([]pdata.TraceID, numIds)
	spanID := 0
	var tds []pdata.Traces
	for i := 0; i < numIds; i++ {
		traceIds[i] = tracetranslator.UInt64ToTraceID(1, uint64(i+1))
		// Send each span in a separate batch
		for j := 0; j <= i; j++ {
			td := simpleTraces()
			span := td.ResourceSpans().At(0).InstrumentationLibrarySpans().At(0).Spans().At(0)
			span.SetTraceID(traceIds[i])

			spanID++
			span.SetSpanID(tracetranslator.UInt64ToSpanID(uint64(spanID)))
			tds = append(tds, td)
		}
	}

	return traceIds, tds
}

func generateBatchesForGivenId(id int, numOfBatches int, numOfSpansInBatch int) []pdata.Traces {
	spanID := 0
	var tds []pdata.Traces
	for i:=0; i< numOfBatches; i++ {
		trace := pdata.NewTraces()
		ils := pdata.NewInstrumentationLibrarySpans()
		rs := pdata.NewResourceSpans()
		rs.InstrumentationLibrarySpans().Append(ils)
		trace.ResourceSpans().Append(rs)
		for j:=0; j<numOfSpansInBatch; j++ {
			span := pdata.NewSpan()
			span.SetTraceID(tracetranslator.Int64ToTraceID(0, int64(id)))
			span.SetSpanID(tracetranslator.UInt64ToSpanID(uint64(spanID)))
			spanID++
			ils.Spans().Append(span)
		}

		tds = append(tds, trace)
	}
	return tds
}

func simpleTraces() pdata.Traces {
	return simpleTracesWithID(pdata.NewTraceID([16]byte{1, 2, 3, 4}))
}

func simpleTracesWithID(traceID pdata.TraceID) pdata.Traces {
	span := pdata.NewSpan()
	span.SetTraceID(traceID)

	ils := pdata.NewInstrumentationLibrarySpans()
	ils.Spans().Append(span)

	rs := pdata.NewResourceSpans()
	rs.InstrumentationLibrarySpans().Append(ils)

	traces := pdata.NewTraces()
	traces.ResourceSpans().Append(rs)

	return traces
}

type mockPolicyEvaluator struct {
	NextDecision           sampling.Decision
	NextError              error
	EvaluationCount        int
	LateArrivingSpansCount int
	OnDroppedSpansCount    int
}

var _ sampling.PolicyEvaluator = (*mockPolicyEvaluator)(nil)

func (m *mockPolicyEvaluator) OnLateArrivingSpans(sampling.Decision, []*pdata.Span) error {
	m.LateArrivingSpansCount++
	return m.NextError
}
func (m *mockPolicyEvaluator) Evaluate(_ pdata.TraceID, _ *sampling.TraceData) (sampling.Decision, error) {
	m.EvaluationCount++
	return m.NextDecision, m.NextError
}

type manualTTicker struct {
	Started bool
}

var _ tTicker = (*manualTTicker)(nil)

func (t *manualTTicker) Start(time.Duration) {
	t.Started = true
}

func (t *manualTTicker) OnTick() {
}

func (t *manualTTicker) Stop() {
}

type syncIDBatcher struct {
	sync.Mutex
	openBatch idbatcher.Batch
	batchPipe chan idbatcher.Batch
}

var _ idbatcher.Batcher = (*syncIDBatcher)(nil)

func newSyncIDBatcher(numBatches uint64) idbatcher.Batcher {
	batches := make(chan idbatcher.Batch, numBatches)
	for i := uint64(0); i < numBatches; i++ {
		batches <- nil
	}
	return &syncIDBatcher{
		batchPipe: batches,
	}
}

func (s *syncIDBatcher) AddToCurrentBatch(id pdata.TraceID) {
	s.Lock()
	s.openBatch = append(s.openBatch, id)
	s.Unlock()
}

func (s *syncIDBatcher) CloseCurrentAndTakeFirstBatch() (idbatcher.Batch, bool) {
	s.Lock()
	defer s.Unlock()
	firstBatch := <-s.batchPipe
	s.batchPipe <- s.openBatch
	s.openBatch = nil
	return firstBatch, true
}

func (s *syncIDBatcher) Stop() {
}
