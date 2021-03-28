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

package groupbytraceprocessor

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/pdata"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchpersignal"
)

var (
	logger, _ = zap.NewDevelopment()
)

func TestTraceIsDispatchedAfterDuration(t *testing.T) {
	// prepare
	traces := simpleTraces()

	wgReceived := &sync.WaitGroup{} // we wait for the next (mock) processor to receive the trace
	config := Config{
		WaitDuration: time.Nanosecond,
		NumTraces:    10,
		NumWorkers:   1,
	}
	mockProcessor := &mockProcessor{
		onTraces: func(ctx context.Context, received pdata.Traces) error {
			assert.Equal(t, traces, received)
			wgReceived.Done()
			return nil
		},
	}

	wgDeleted := &sync.WaitGroup{} // we wait for the next (mock) processor to receive the trace
	backing := newMemoryStorage()
	st := &mockStorage{
		onCreateOrAppend: backing.createOrAppend,
		onGet:            backing.get,
		onDelete: func(traceID pdata.TraceID) ([]pdata.ResourceSpans, error) {
			wgDeleted.Done()
			return backing.delete(traceID)
		},
	}

	p := newGroupByTraceProcessor(logger, st, mockProcessor, config)
	ctx := context.Background()
	p.Start(ctx, nil)
	defer p.Shutdown(ctx)

	// test
	wgReceived.Add(1) // one should be received
	wgDeleted.Add(1)  // one should be deleted
	p.ConsumeTraces(ctx, traces)

	// verify
	wgReceived.Wait()
	wgDeleted.Wait()
}

func TestInternalCacheLimit(t *testing.T) {
	// prepare
	wg := &sync.WaitGroup{} // we wait for the next (mock) processor to receive the trace

	config := Config{
		// should be long enough for the test to run without traces being finished, but short enough to not
		// badly influence the testing experience
		WaitDuration: 50 * time.Millisecond,

		// we create 6 traces, only 5 should be at the storage in the end
		NumTraces: 5,

		NumWorkers: 1,
	}

	wg.Add(5) // 5 traces are expected to be received

	var receivedTraceIDs []pdata.TraceID
	mockProcessor := &mockProcessor{}
	mockProcessor.onTraces = func(ctx context.Context, received pdata.Traces) error {
		traceID := received.ResourceSpans().At(0).InstrumentationLibrarySpans().At(0).Spans().At(0).TraceID()
		receivedTraceIDs = append(receivedTraceIDs, traceID)
		wg.Done()
		return nil
	}

	st := newMemoryStorage()

	p := newGroupByTraceProcessor(logger, st, mockProcessor, config)

	ctx := context.Background()
	p.Start(ctx, nil)
	defer p.Shutdown(ctx)

	// test
	traceIDs := [][16]byte{
		{1, 2, 3, 4},
		{2, 3, 4, 5},
		{3, 4, 5, 6},
		{4, 5, 6, 7},
		{5, 6, 7, 8},
		{6, 7, 8, 9},
	}

	// 6 iterations
	for _, traceID := range traceIDs {
		batch := simpleTracesWithID(pdata.NewTraceID(traceID))
		p.ConsumeTraces(ctx, batch)
	}

	wg.Wait()

	// verify
	assert.Equal(t, 5, len(receivedTraceIDs))

	for i := 5; i > 0; i-- { // last 5 traces
		traceID := pdata.NewTraceID(traceIDs[i])
		assert.Contains(t, receivedTraceIDs, traceID)
	}

	// the first trace should have been evicted
	assert.NotContains(t, receivedTraceIDs, traceIDs[0])
}

func TestProcessorCapabilities(t *testing.T) {
	// prepare
	config := Config{
		WaitDuration: time.Nanosecond,
		NumTraces:    10,
		NumWorkers:   1,
	}
	st := newMemoryStorage()
	next := &mockProcessor{}

	// test
	p := newGroupByTraceProcessor(logger, st, next, config)
	caps := p.GetCapabilities()

	// verify
	assert.NotNil(t, p)
	assert.Equal(t, true, caps.MutatesConsumedData)
}

func TestProcessBatchDoesntFail(t *testing.T) {
	// prepare
	config := Config{
		WaitDuration: time.Nanosecond,
		NumTraces:    10,
		NumWorkers:   1,
	}
	st := newMemoryStorage()
	next := &mockProcessor{}

	traceID := pdata.NewTraceID([16]byte{1, 2, 3, 4})

	trace := pdata.NewTraces()
	rs := trace.ResourceSpans().AppendEmpty()
	ils := rs.InstrumentationLibrarySpans().AppendEmpty()
	span := ils.Spans().AppendEmpty()
	span.SetTraceID(traceID)
	span.SetSpanID(pdata.NewSpanID([8]byte{1, 2, 3, 4}))

	p := newGroupByTraceProcessor(logger, st, next, config)
	assert.NotNil(t, p)

	// test
	p.onTraceReceived(tracesWithID{id: traceID, td: trace}, p.eventMachine.workers[0])
}

func TestTraceDisappearedFromStorageBeforeReleasing(t *testing.T) {
	// prepare
	config := Config{
		WaitDuration: time.Second, // we are not waiting for this whole time
		NumTraces:    5,
		NumWorkers:   1,
	}
	st := &mockStorage{
		onGet: func(pdata.TraceID) ([]pdata.ResourceSpans, error) {
			return nil, nil
		},
	}
	next := &mockProcessor{}

	p := newGroupByTraceProcessor(logger, st, next, config)
	require.NotNil(t, p)

	traceID := pdata.NewTraceID([16]byte{1, 2, 3, 4})
	batch := simpleTracesWithID(traceID)

	ctx := context.Background()
	p.Start(ctx, nil)
	defer p.Shutdown(ctx)

	err := p.ConsumeTraces(context.Background(), batch)
	require.NoError(t, err)

	// test
	// we trigger this manually, instead of waiting the whole duration
	err = p.markAsReleased(traceID, p.eventMachine.workers[0].fire)

	// verify
	assert.Error(t, err)
}

func TestTraceErrorFromStorageWhileReleasing(t *testing.T) {
	// prepare
	config := Config{
		WaitDuration: time.Second, // we are not waiting for this whole time
		NumTraces:    5,
		NumWorkers:   1,
	}
	expectedError := errors.New("some unexpected error")
	st := &mockStorage{
		onGet: func(pdata.TraceID) ([]pdata.ResourceSpans, error) {
			return nil, expectedError
		},
	}
	next := &mockProcessor{}

	p := newGroupByTraceProcessor(logger, st, next, config)
	require.NotNil(t, p)

	traceID := pdata.NewTraceID([16]byte{1, 2, 3, 4})
	batch := simpleTracesWithID(traceID)

	ctx := context.Background()
	p.Start(ctx, nil)
	defer p.Shutdown(ctx)

	err := p.ConsumeTraces(context.Background(), batch)
	require.NoError(t, err)

	// test
	// we trigger this manually, instead of waiting the whole duration
	err = p.markAsReleased(traceID, p.eventMachine.workers[0].fire)

	// verify
	assert.True(t, errors.Is(err, expectedError))
}

func TestTraceErrorFromStorageWhileProcessingTrace(t *testing.T) {
	// prepare
	config := Config{
		WaitDuration: time.Second, // we are not waiting for this whole time
		NumTraces:    5,
		NumWorkers:   1,
	}
	expectedError := errors.New("some unexpected error")
	st := &mockStorage{
		onCreateOrAppend: func(pdata.TraceID, pdata.Traces) error {
			return expectedError
		},
	}
	next := &mockProcessor{}

	p := newGroupByTraceProcessor(logger, st, next, config)
	require.NotNil(t, p)

	traceID := pdata.NewTraceID([16]byte{1, 2, 3, 4})

	trace := pdata.NewTraces()
	rss := trace.ResourceSpans()
	rs := rss.AppendEmpty()
	ils := rs.InstrumentationLibrarySpans().AppendEmpty()
	span := ils.Spans().AppendEmpty()
	span.SetTraceID(traceID)
	span.SetSpanID(pdata.NewSpanID([8]byte{1, 2, 3, 4}))

	batch := batchpersignal.SplitTraces(trace)

	// test
	err := p.onTraceReceived(tracesWithID{id: traceID, td: batch[0]}, p.eventMachine.workers[0])

	// verify
	assert.True(t, errors.Is(err, expectedError))
}

func TestAddSpansToExistingTrace(t *testing.T) {
	// prepare
	wg := &sync.WaitGroup{}
	config := Config{
		WaitDuration: 50 * time.Millisecond,
		NumTraces:    5,
		NumWorkers:   1,
	}
	st := newMemoryStorage()

	var receivedTraces []pdata.ResourceSpans
	next := &mockProcessor{
		onTraces: func(ctx context.Context, traces pdata.Traces) error {
			require.Equal(t, 2, traces.ResourceSpans().Len())
			receivedTraces = append(receivedTraces, traces.ResourceSpans().At(0))
			receivedTraces = append(receivedTraces, traces.ResourceSpans().At(1))
			wg.Done()
			return nil
		},
	}

	p := newGroupByTraceProcessor(logger, st, next, config)
	require.NotNil(t, p)

	ctx := context.Background()
	p.Start(ctx, nil)
	defer p.Shutdown(ctx)

	traceID := pdata.NewTraceID([16]byte{1, 2, 3, 4})

	// test
	first := simpleTracesWithID(traceID)
	first.ResourceSpans().At(0).InstrumentationLibrarySpans().At(0).Spans().At(0).SetName("first-span")

	second := simpleTracesWithID(traceID)
	second.ResourceSpans().At(0).InstrumentationLibrarySpans().At(0).Spans().At(0).SetName("second-span")

	wg.Add(1)

	p.ConsumeTraces(context.Background(), first)
	p.ConsumeTraces(context.Background(), second)

	wg.Wait()

	// verify
	assert.Len(t, receivedTraces, 2)
}

func TestTraceErrorFromStorageWhileProcessingSecondTrace(t *testing.T) {
	// prepare
	config := Config{
		WaitDuration: time.Second, // we are not waiting for this whole time
		NumTraces:    5,
		NumWorkers:   1,
	}
	st := &mockStorage{}
	next := &mockProcessor{}

	p := newGroupByTraceProcessor(logger, st, next, config)
	require.NotNil(t, p)

	traceID := pdata.NewTraceID([16]byte{1, 2, 3, 4})

	trace := pdata.NewTraces()
	rss := trace.ResourceSpans()
	rs := rss.AppendEmpty()
	ils := rs.InstrumentationLibrarySpans().AppendEmpty()
	span := ils.Spans().AppendEmpty()
	span.SetTraceID(traceID)
	span.SetSpanID(pdata.NewSpanID([8]byte{1, 2, 3, 4}))

	batch := batchpersignal.SplitTraces(trace)

	// test
	err := p.onTraceReceived(tracesWithID{id: traceID, td: batch[0]}, p.eventMachine.workers[0])
	assert.NoError(t, err)

	expectedError := errors.New("some unexpected error")
	st.onCreateOrAppend = func(pdata.TraceID, pdata.Traces) error {
		return expectedError
	}

	// processing another batch for the same trace takes a slightly different code path
	err = p.onTraceReceived(tracesWithID{id: traceID, td: batch[0]}, p.eventMachine.workers[0])

	// verify
	assert.True(t, errors.Is(err, expectedError))
}

func TestErrorFromStorageWhileRemovingTrace(t *testing.T) {
	// prepare
	config := Config{
		WaitDuration: time.Second, // we are not waiting for this whole time
		NumTraces:    5,
		NumWorkers:   1,
	}
	expectedError := errors.New("some unexpected error")
	st := &mockStorage{
		onDelete: func(pdata.TraceID) ([]pdata.ResourceSpans, error) {
			return nil, expectedError
		},
	}
	next := &mockProcessor{}

	p := newGroupByTraceProcessor(logger, st, next, config)
	require.NotNil(t, p)

	traceID := pdata.NewTraceID([16]byte{1, 2, 3, 4})

	// test
	err := p.onTraceRemoved(traceID)

	// verify
	assert.True(t, errors.Is(err, expectedError))
}

func TestTraceNotFoundWhileRemovingTrace(t *testing.T) {
	// prepare
	config := Config{
		WaitDuration: time.Second, // we are not waiting for this whole time
		NumTraces:    5,
		NumWorkers:   1,
	}
	st := &mockStorage{
		onDelete: func(pdata.TraceID) ([]pdata.ResourceSpans, error) {
			return nil, nil
		},
	}
	next := &mockProcessor{}

	p := newGroupByTraceProcessor(logger, st, next, config)
	require.NotNil(t, p)

	traceID := pdata.NewTraceID([16]byte{1, 2, 3, 4})

	// test
	err := p.onTraceRemoved(traceID)

	// verify
	assert.Error(t, err)
}

func TestTracesAreDispatchedInIndividualBatches(t *testing.T) {
	// prepare
	wg := &sync.WaitGroup{}

	config := Config{
		WaitDuration: time.Nanosecond, // we are not waiting for this whole time
		NumTraces:    5,
		NumWorkers:   1,
	}
	st := newMemoryStorage()
	next := &mockProcessor{
		onTraces: func(_ context.Context, traces pdata.Traces) error {
			// we should receive two batches, each one with one trace
			assert.Equal(t, 1, traces.ResourceSpans().Len())
			wg.Done()
			return nil
		},
	}

	p := newGroupByTraceProcessor(logger, st, next, config)
	require.NotNil(t, p)

	ctx := context.Background()
	p.Start(ctx, nil)
	defer p.Shutdown(ctx)

	traceID := pdata.NewTraceID([16]byte{1, 2, 3, 4})

	firstTrace := pdata.NewTraces()
	firstRss := firstTrace.ResourceSpans()
	firstResourceSpans := firstRss.AppendEmpty()
	ils := firstResourceSpans.InstrumentationLibrarySpans().AppendEmpty()
	span := ils.Spans().AppendEmpty()
	span.SetTraceID(traceID)

	secondTraceID := pdata.NewTraceID([16]byte{2, 3, 4, 5})
	secondTrace := pdata.NewTraces()
	secondRss := secondTrace.ResourceSpans()
	secondResourceSpans := secondRss.AppendEmpty()
	secondIls := secondResourceSpans.InstrumentationLibrarySpans().AppendEmpty()
	secondSpan := secondIls.Spans().AppendEmpty()
	secondSpan.SetTraceID(secondTraceID)

	// test
	wg.Add(2)

	p.onTraceReceived(tracesWithID{id: traceID, td: firstTrace}, p.eventMachine.workers[0])
	p.onTraceReceived(tracesWithID{id: secondTraceID, td: secondTrace}, p.eventMachine.workers[0])

	wg.Wait()

	// verify
	// verification is done at onTraces from the mockProcessor
}

func TestSplitSameTraceIntoDifferentBatches(t *testing.T) {
	// prepare

	// we have 1 ResourceSpans with 2 ILS, resulting in two batches
	input := pdata.NewResourceSpans()
	input.InstrumentationLibrarySpans().Resize(2)

	// the first ILS has two spans
	firstILS := input.InstrumentationLibrarySpans().At(0)
	firstLibrary := firstILS.InstrumentationLibrary()
	firstLibrary.SetName("first-library")
	firstILS.Spans().Resize(2)
	firstSpan := firstILS.Spans().At(0)
	firstSpan.SetName("first-batch-first-span")
	firstSpan.SetTraceID(pdata.NewTraceID([16]byte{1, 2, 3, 4}))
	secondSpan := firstILS.Spans().At(1)
	secondSpan.SetName("first-batch-second-span")
	secondSpan.SetTraceID(pdata.NewTraceID([16]byte{1, 2, 3, 4}))

	// the second ILS has one span
	secondILS := input.InstrumentationLibrarySpans().At(1)
	secondLibrary := secondILS.InstrumentationLibrary()
	secondLibrary.SetName("second-library")
	thirdSpan := secondILS.Spans().AppendEmpty()
	thirdSpan.SetName("second-batch-first-span")
	thirdSpan.SetTraceID(pdata.NewTraceID([16]byte{1, 2, 3, 4}))

	// test
	batches := splitByTrace(input)

	// verify
	assert.Len(t, batches, 2)

	// first batch
	assert.Equal(t, pdata.NewTraceID([16]byte{1, 2, 3, 4}), batches[0].traceID)
	assert.Equal(t, firstLibrary.Name(), batches[0].rs.InstrumentationLibrarySpans().At(0).InstrumentationLibrary().Name())
	assert.Equal(t, firstSpan.Name(), batches[0].rs.InstrumentationLibrarySpans().At(0).Spans().At(0).Name())
	assert.Equal(t, secondSpan.Name(), batches[0].rs.InstrumentationLibrarySpans().At(0).Spans().At(1).Name())

	// second batch
	assert.Equal(t, pdata.NewTraceID([16]byte{1, 2, 3, 4}), batches[1].traceID)
	assert.Equal(t, secondLibrary.Name(), batches[1].rs.InstrumentationLibrarySpans().At(0).InstrumentationLibrary().Name())
	assert.Equal(t, thirdSpan.Name(), batches[1].rs.InstrumentationLibrarySpans().At(0).Spans().At(0).Name())
}

func TestSplitDifferentTracesIntoDifferentBatches(t *testing.T) {
	// prepare

	// we have 1 ResourceSpans with 1 ILS and two traceIDs, resulting in two batches
	input := pdata.NewResourceSpans()

	// the first ILS has two spans
	ils := input.InstrumentationLibrarySpans().AppendEmpty()
	library := ils.InstrumentationLibrary()
	library.SetName("first-library")
	ils.Spans().Resize(2)
	firstSpan := ils.Spans().At(0)
	firstSpan.SetName("first-batch-first-span")
	firstSpan.SetTraceID(pdata.NewTraceID([16]byte{1, 2, 3, 4}))
	secondSpan := ils.Spans().At(1)
	secondSpan.SetName("first-batch-second-span")
	secondSpan.SetTraceID(pdata.NewTraceID([16]byte{2, 3, 4, 5}))

	// test
	batches := splitByTrace(input)

	// verify
	assert.Len(t, batches, 2)

	// first batch
	assert.Equal(t, pdata.NewTraceID([16]byte{1, 2, 3, 4}), batches[0].traceID)
	assert.Equal(t, library.Name(), batches[0].rs.InstrumentationLibrarySpans().At(0).InstrumentationLibrary().Name())
	assert.Equal(t, firstSpan.Name(), batches[0].rs.InstrumentationLibrarySpans().At(0).Spans().At(0).Name())

	// second batch
	assert.Equal(t, pdata.NewTraceID([16]byte{2, 3, 4, 5}), batches[1].traceID)
	assert.Equal(t, library.Name(), batches[1].rs.InstrumentationLibrarySpans().At(0).InstrumentationLibrary().Name())
	assert.Equal(t, secondSpan.Name(), batches[1].rs.InstrumentationLibrarySpans().At(0).Spans().At(0).Name())
}

func TestSplitByTraceWithNilTraceID(t *testing.T) {
	// prepare
	input := pdata.NewResourceSpans()
	ils := input.InstrumentationLibrarySpans().AppendEmpty()
	firstSpan := ils.Spans().AppendEmpty()
	firstSpan.SetTraceID(pdata.NewTraceID([16]byte{}))

	// test
	batches := splitByTrace(input)

	// verify
	assert.Len(t, batches, 0)
}

func TestErrorOnProcessResourceSpansContinuesProcessing(t *testing.T) {
	// prepare
	config := Config{
		WaitDuration: time.Second, // we are not waiting for this whole time
		NumTraces:    5,
		NumWorkers:   1,
	}
	st := &mockStorage{}
	next := &mockProcessor{}

	p := newGroupByTraceProcessor(logger, st, next, config)
	require.NotNil(t, p)

	traceID := pdata.NewTraceID([16]byte{1, 2, 3, 4})

	trace := pdata.NewTraces()
	rss := trace.ResourceSpans()
	rs := rss.AppendEmpty()
	ils := rs.InstrumentationLibrarySpans().AppendEmpty()
	span := ils.Spans().AppendEmpty()
	span.SetTraceID(traceID)
	span.SetSpanID(pdata.NewSpanID([8]byte{1, 2, 3, 4}))

	expectedError := errors.New("some unexpected error")
	returnedError := false
	st.onCreateOrAppend = func(pdata.TraceID, pdata.Traces) error {
		returnedError = true
		return expectedError
	}

	// test
	p.onTraceReceived(tracesWithID{id: traceID, td: trace}, p.eventMachine.workers[0])

	// verify
	assert.True(t, returnedError)
}

func TestAsyncOnRelease(t *testing.T) {
	blockCh := make(chan struct{})
	blocker := &blockingConsumer{
		blockCh: blockCh,
	}

	sp := &groupByTraceProcessor{
		logger:       zap.NewNop(),
		nextConsumer: blocker,
	}
	assert.NoError(t, sp.onTraceReleased(nil))
	close(blockCh)
}

func BenchmarkConsumeTracesCompleteOnFirstBatch(b *testing.B) {
	// prepare
	config := Config{
		WaitDuration: 50 * time.Millisecond,
		NumTraces:    defaultNumTraces,
		NumWorkers:   2 * defaultNumWorkers,
	}
	st := newMemoryStorage()
	next := &mockProcessor{}

	p := newGroupByTraceProcessor(zap.NewNop(), st, next, config)
	require.NotNil(b, p)

	ctx := context.Background()
	p.Start(ctx, nil)
	defer p.Shutdown(ctx)

	for n := 0; n < b.N; n++ {
		traceID := pdata.NewTraceID([16]byte{byte(1 + n), 2, 3, 4})
		trace := simpleTracesWithID(traceID)
		p.ConsumeTraces(context.Background(), trace)
	}
}

type mockProcessor struct {
	mutex    sync.Mutex
	onTraces func(context.Context, pdata.Traces) error
}

var _ component.TracesProcessor = (*mockProcessor)(nil)

func (m *mockProcessor) ConsumeTraces(ctx context.Context, td pdata.Traces) error {
	if m.onTraces != nil {
		m.mutex.Lock()
		defer m.mutex.Unlock()
		return m.onTraces(ctx, td)
	}
	return nil
}
func (m *mockProcessor) GetCapabilities() component.ProcessorCapabilities {
	return component.ProcessorCapabilities{MutatesConsumedData: true}
}
func (m *mockProcessor) Shutdown(context.Context) error {
	return nil
}
func (m *mockProcessor) Start(_ context.Context, _ component.Host) error {
	return nil
}

type mockStorage struct {
	onCreateOrAppend func(pdata.TraceID, pdata.Traces) error
	onGet            func(pdata.TraceID) ([]pdata.ResourceSpans, error)
	onDelete         func(pdata.TraceID) ([]pdata.ResourceSpans, error)
	onStart          func() error
	onShutdown       func() error
}

var _ storage = (*mockStorage)(nil)

func (st *mockStorage) createOrAppend(traceID pdata.TraceID, trace pdata.Traces) error {
	if st.onCreateOrAppend != nil {
		return st.onCreateOrAppend(traceID, trace)
	}
	return nil
}
func (st *mockStorage) get(traceID pdata.TraceID) ([]pdata.ResourceSpans, error) {
	if st.onGet != nil {
		return st.onGet(traceID)
	}
	return nil, nil
}
func (st *mockStorage) delete(traceID pdata.TraceID) ([]pdata.ResourceSpans, error) {
	if st.onDelete != nil {
		return st.onDelete(traceID)
	}
	return nil, nil
}
func (st *mockStorage) start() error {
	if st.onStart != nil {
		return st.onStart()
	}
	return nil
}
func (st *mockStorage) shutdown() error {
	if st.onShutdown != nil {
		return st.onShutdown()
	}
	return nil
}

type blockingConsumer struct {
	blockCh <-chan struct{}
}

var _ consumer.Traces = (*blockingConsumer)(nil)

func (b *blockingConsumer) ConsumeTraces(context.Context, pdata.Traces) error {
	<-b.blockCh
	return nil
}

func simpleTraces() pdata.Traces {
	return simpleTracesWithID(pdata.NewTraceID([16]byte{1, 2, 3, 4}))
}

func simpleTracesWithID(traceID pdata.TraceID) pdata.Traces {
	traces := pdata.NewTraces()
	rs := traces.ResourceSpans().AppendEmpty()
	ils := rs.InstrumentationLibrarySpans().AppendEmpty()
	ils.Spans().AppendEmpty().SetTraceID(traceID)
	return traces
}
