// Copyright The OpenTelemetry Authors
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

package mongodbreceiver

import (
	"errors"
	"strings"
)

/*
This file defines the expected partial errors in the test cases
*/
var (
	errAllPartialMetrics = errors.New(
		strings.Join(
			[]string{
				"failed to collect metric mongodb.cache.operations with attribute(s) miss, hit: could not find key for metric",
				"failed to collect metric mongodb.cursor.count: could not find key for metric",
				"failed to collect metric mongodb.cursor.timeout.count: could not find key for metric",
				"failed to collect metric mongodb.global_lock.time: could not find key for metric",
				"failed to collect metric bytesIn: could not find key for metric",
				"failed to collect metric bytesOut: could not find key for metric",
				"failed to collect metric numRequests: could not find key for metric",
				"failed to collect metric mongodb.operation.count with attribute(s) delete: could not find key for metric",
				"failed to collect metric mongodb.operation.count with attribute(s) getmore: could not find key for metric",
				"failed to collect metric mongodb.operation.count with attribute(s) command: could not find key for metric",
				"failed to collect metric mongodb.operation.count with attribute(s) insert: could not find key for metric",
				"failed to collect metric mongodb.operation.count with attribute(s) query: could not find key for metric",
				"failed to collect metric mongodb.operation.count with attribute(s) update: could not find key for metric",
				"failed to collect metric mongodb.session.count: could not find key for metric",
				"failed to collect metric mongodb.operation.time: could not find key for metric",
				"failed to collect metric mongodb.collection.count with attribute(s) fakedatabase: could not find key for metric",
				"failed to collect metric mongodb.data.size with attribute(s) fakedatabase: could not find key for metric",
				"failed to collect metric mongodb.extent.count with attribute(s) fakedatabase: could not find key for metric",
				"failed to collect metric mongodb.index.size with attribute(s) fakedatabase: could not find key for metric",
				"failed to collect metric mongodb.index.count with attribute(s) fakedatabase: could not find key for metric",
				"failed to collect metric mongodb.object.count with attribute(s) fakedatabase: could not find key for metric",
				"failed to collect metric mongodb.storage.size with attribute(s) fakedatabase: could not find key for metric",
				"failed to collect metric mongodb.connection.count with attribute(s) available, fakedatabase: could not find key for metric",
				"failed to collect metric mongodb.connection.count with attribute(s) current, fakedatabase: could not find key for metric",
				"failed to collect metric mongodb.connection.count with attribute(s) active, fakedatabase: could not find key for metric",
				"failed to collect metric mongodb.document.operation.count with attribute(s) inserted, fakedatabase: could not find key for metric",
				"failed to collect metric mongodb.document.operation.count with attribute(s) updated, fakedatabase: could not find key for metric",
				"failed to collect metric mongodb.document.operation.count with attribute(s) deleted, fakedatabase: could not find key for metric",
				"failed to collect metric mongodb.memory.usage with attribute(s) resident, fakedatabase: could not find key for metric",
				"failed to collect metric mongodb.memory.usage with attribute(s) virtual, fakedatabase: could not find key for metric",
				"failed to collect metric mongodb.index.access.count with attribute(s) fakedatabase, orders: could not find key for index access metric",
				"failed to collect metric mongodb.index.access.count with attribute(s) fakedatabase, products: could not find key for index access metric",
			}, "; "))
	errAllClientFailedFetch = errors.New(
		strings.Join(
			[]string{
				"failed to fetch admin server status metrics: some admin server status error",
				"failed to fetch top stats metrics: some top stats error",
				"failed to fetch database stats metrics: some database stats error",
				"failed to fetch server status metrics: some server status error",
				"failed to fetch index stats metrics: some index stats error",
				"failed to fetch index stats metrics: some index stats error",
			}, "; "))

	errCollectionNames = errors.New(
		strings.Join(
			[]string{
				"failed to fetch admin server status metrics: some admin server status error",
				"failed to fetch top stats metrics: some top stats error",
				"failed to fetch database stats metrics: some database stats error",
				"failed to fetch server status metrics: some server status error",
				"failed to fetch collection names: some collection names error",
			}, "; "))
	)
