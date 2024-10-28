// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package admission // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/otelarrow/admission"

import (
	"container/list"
	"context"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	grpccodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrTooMuchWaiting = status.Error(grpccodes.ResourceExhausted, "rejecting request, too much pending data")
var ErrRequestTooLarge = status.Errorf(grpccodes.InvalidArgument, "rejecting request, request is too large")

// BoundedQueue is a LIFO-oriented admission-controlled Queue.
type BoundedQueue struct {
	maxLimitAdmit int64
	maxLimitWait  int64
	tracer        trace.Tracer

	// lock protects currentAdmitted, currentWaiting, and waiters

	lock            sync.Mutex
	currentAdmitted int64
	currentWaiting  int64
	waiters         *list.List // of *waiter
}

var _ Queue = &BoundedQueue{}

// waiter is an item in the BoundedQueue waiters list.
type waiter struct {
	notify  N
	pending int64
}

// NewBoundedQueue returns a LIFO-oriented Queue implementation which
// admits `maxLimitAdmit` bytes concurrently and allows up to
// `maxLimitWait` bytes to wait for admission.
func NewBoundedQueue(ts component.TelemetrySettings, maxLimitAdmit, maxLimitWait int64) Queue {
	return &BoundedQueue{
		maxLimitAdmit: maxLimitAdmit,
		maxLimitWait:  maxLimitWait,
		waiters:       list.New(),
		tracer:        ts.TracerProvider.Tracer("github.com/open-telemetry/opentelemetry-collector-contrib/internal/otelarrow"),
	}
}

// acquireOrGetWaiter returns with three distinct conditions depending
// on whether it was accepted, rejected, or asked to wait.
//
// - element=nil, error=nil: the fast success path
// - element=nil, error=non-nil: the fast failure path
// - element=non-nil, error=non-nil: the slow success path
func (bq *BoundedQueue) acquireOrGetWaiter(pending int64) (*list.Element, error) {
	bq.lock.Lock()
	defer bq.lock.Unlock()

	if pending > bq.maxLimitAdmit {
		// when the request will never succeed because it is
		// individually over the total limit, fail fast.
		return nil, ErrRequestTooLarge
	}

	if bq.currentAdmitted+pending <= bq.maxLimitAdmit {
		// the fast success path.
		bq.currentAdmitted += pending
		return nil, nil
	}

	// since we were unable to admit, check if we can wait.
	if bq.currentWaiting+pending > bq.maxLimitWait {
		return nil, ErrTooMuchWaiting
	}

	// otherwise we need to wait
	return bq.addWaiterLocked(pending), nil
}

// Acquire implements Queue.
func (bq *BoundedQueue) Acquire(ctx context.Context, pending int64) (ReleaseFunc, error) {
	element, err := bq.acquireOrGetWaiter(pending)
	parentSpan := trace.SpanFromContext(ctx)
	pendingAttr := trace.WithAttributes(attribute.Int64("pending", pending))

	if err != nil {
		parentSpan.AddEvent("admission rejected (fast path)", pendingAttr)
		return noopRelease, err
	} else if element == nil {
		parentSpan.AddEvent("admission accepted (fast path)", pendingAttr)
		return bq.releaseFunc(pending), nil
	}

	parentSpan.AddEvent("enter admission queue")

	ctx, span := bq.tracer.Start(ctx, "admission_blocked", pendingAttr)
	defer span.End()

	waiter := element.Value.(*waiter)

	select {
	case <-waiter.notify.Chan():
		parentSpan.AddEvent("admission accepted (slow path)", pendingAttr)
		return bq.releaseFunc(pending), nil

	case <-ctx.Done():
		bq.lock.Lock()
		defer bq.lock.Unlock()

		if waiter.notify.HasBeenNotified() {
			// We were also admitted, which can happen
			// concurrently with cancellation. Make sure
			// to release since no one else will do it.
			bq.releaseLocked(pending)
		} else {
			// Remove ourselves from the list of waiters
			// so that we can't be admitted in the future.
			bq.removeWaiterLocked(pending, element)
			bq.admitWaitersLocked()
		}

		parentSpan.AddEvent("admission rejected (canceled)", pendingAttr)
		return noopRelease, status.Error(grpccodes.Canceled, context.Cause(ctx).Error())
	}
}

func (bq *BoundedQueue) admitWaitersLocked() {
	for bq.waiters.Len() != 0 {
		// Ensure there is enough room to admit the next waiter.
		element := bq.waiters.Back()
		waiter := element.Value.(*waiter)
		if bq.currentAdmitted+waiter.pending > bq.maxLimitAdmit {
			// Returning means continuing to wait for the
			// most recent arrival to get service by another release.
			return
		}

		// Release the next waiter and tell it that it has been admitted.
		bq.removeWaiterLocked(waiter.pending, element)
		bq.currentAdmitted += waiter.pending

		waiter.notify.Notify()
	}
}

func (bq *BoundedQueue) addWaiterLocked(pending int64) *list.Element {
	bq.currentWaiting += pending
	return bq.waiters.PushBack(&waiter{
		pending: pending,
		notify:  newNotification(),
	})
}

func (bq *BoundedQueue) removeWaiterLocked(pending int64, element *list.Element) {
	bq.currentWaiting -= pending
	bq.waiters.Remove(element)
}

func (bq *BoundedQueue) releaseLocked(pending int64) {
	bq.currentAdmitted -= pending
	bq.admitWaitersLocked()
}

func (bq *BoundedQueue) releaseFunc(pending int64) ReleaseFunc {
	return func() {
		bq.lock.Lock()
		defer bq.lock.Unlock()

		bq.releaseLocked(pending)
	}
}
