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

package datadogexporter

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHost(t *testing.T) {

	host := GetHost(&Config{TagsConfig: TagsConfig{Hostname: "test_host"}})
	assert.Equal(t, *host, "test_host")

	host = GetHost(&Config{})
	osHostname, err := os.Hostname()
	if err == nil {
		assert.Equal(t, *host, osHostname)
	} else {
		assert.Equal(t, *host, "unknown")
	}
}
