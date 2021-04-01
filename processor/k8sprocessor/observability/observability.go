// Copyright 2020 OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package observability

import (
	"context"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
)

// TODO: re-think if processor should register it's own telemetry views or if some other
// mechanism should be used by the collector to discover views from all components

func init() {
	view.Register(
		viewPodsUpdated,
		viewPodsAdded,
		viewPodsDeleted,
		viewIPLookupMiss,
		viewPodTableSize,
	)
}

var (
	mPodsUpdated  = stats.Int64("otelcol_k8s_pod_updated", "Number of pod update events received", "1")
	mPodsAdded    = stats.Int64("otelcol_k8s_pod_added", "Number of pod add events received", "1")
	mPodsDeleted  = stats.Int64("otelcol_k8s_pod_deleted", "Number of pod delete events received", "1")
	mPodTableSize = stats.Int64("otelcol_k8s_pod_table_size", "Size of table containing pod info", "1")

	mIDLookupMiss = stats.Int64("otelcol_k8s_ip_lookup_miss", "Number of times pod by identifier (IP, UID) lookup failed.", "1")
)

var viewPodsUpdated = &view.View{
	Name:        mPodsUpdated.Name(),
	Description: mPodsUpdated.Description(),
	Measure:     mPodsUpdated,
	Aggregation: view.Sum(),
}

var viewPodsAdded = &view.View{
	Name:        mPodsAdded.Name(),
	Description: mPodsAdded.Description(),
	Measure:     mPodsAdded,
	Aggregation: view.Sum(),
}

var viewPodsDeleted = &view.View{
	Name:        mPodsDeleted.Name(),
	Description: mPodsDeleted.Description(),
	Measure:     mPodsDeleted,
	Aggregation: view.Sum(),
}

var viewIPLookupMiss = &view.View{
	Name:        mIDLookupMiss.Name(),
	Description: mIDLookupMiss.Description(),
	Measure:     mIDLookupMiss,
	Aggregation: view.Sum(),
}
var viewPodTableSize = &view.View{
	Name:        mPodTableSize.Name(),
	Description: mPodTableSize.Description(),
	Measure:     mPodTableSize,
	Aggregation: view.LastValue(),
}

// RecordPodUpdated increments the metric that records pod update events received.
func RecordPodUpdated() {
	stats.Record(context.Background(), mPodsUpdated.M(int64(1)))
}

// RecordPodAdded increments the metric that records pod add events receiver.
func RecordPodAdded() {
	stats.Record(context.Background(), mPodsAdded.M(int64(1)))
}

// RecordPodDeleted increments the metric that records pod events deleted.
func RecordPodDeleted() {
	stats.Record(context.Background(), mPodsDeleted.M(int64(1)))
}

// RecordIDLookupMiss increments the metric that records Pod lookup by ID (IP, UID) misses.
func RecordIDLookupMiss() {
	stats.Record(context.Background(), mIDLookupMiss.M(int64(1)))
}

// RecordPodTableSize store size of pod table field in WatchClient
func RecordPodTableSize(podTableSize int64) {
	stats.Record(context.Background(), mPodTableSize.M(podTableSize))
}
