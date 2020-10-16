// Copyright The OpenTelemetry Authors
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

package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddToCache(t *testing.T) {
	const testKey = "test_key"
	const testVal = "test_value"

	_, found := Get(testKey)
	assert.False(t, found)

	SetNoExpire(testKey, testVal)
	val, found := Get(testKey)
	assert.True(t, found)
	assert.Equal(t, testVal, val.(string))

	Delete(testKey)
	_, found = Get(testKey)
	assert.False(t, found)
}
