// Copyright 2019, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package octrace

// Option interface defines for configuration settings to be applied to receivers.
//
// WithReceiver applies the configuration to the given receiver.
type Option func(*Receiver)

// WithWorkerCount sets the number of worker goroutines that will be started
// for the receiver
func WithWorkerCount(workerCount int) Option {
	return func(r *Receiver) {
		r.numWorkers = workerCount
	}
}
