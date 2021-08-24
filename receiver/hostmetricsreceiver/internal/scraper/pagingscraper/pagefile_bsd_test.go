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

// +build freebsd openbsd

package pagingscraper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const validFreeBSD = `Device:       1kB-blocks      Used:
/dev/gpt/swapfs    1048576          1234
/dev/md0         1048576          666
`

const validOpenBSD = `Device       1K-blocks      Used	Avail	Capacity	Priority
/dev/wd0b    655025          1234	653791	1%	0
`

const invalid = `Device:       512-blocks      Used:
/dev/gpt/swapfs    1048576          1234
/dev/md0         1048576          666
`

func TestParseSwapctlOutput_FreeBSD(t *testing.T) {
	assert := assert.New(t)
	stats, err := parseSwapctlOutput(validFreeBSD)
	assert.NoError(err)

	assert.Equal(*stats[0], pageFileStats{
		deviceName: "/dev/gpt/swapfs",
		usedBytes:  1263616,
		freeBytes:  1072478208,
	})

	assert.Equal(*stats[1], pageFileStats{
		deviceName: "/dev/md0",
		usedBytes:  681984,
		freeBytes:  1073059840,
	})
}

func TestParseSwapctlOutput_OpenBSD(t *testing.T) {
	assert := assert.New(t)
	stats, err := parseSwapctlOutput(validOpenBSD)
	assert.NoError(err)

	assert.Equal(*stats[0], pageFileStats{
		deviceName: "/dev/wd0b",
		usedBytes:  1234 * 1024,
		freeBytes:  653791 * 1024,
	})
}

func TestParseSwapctlOutput_Invalid(t *testing.T) {
	_, err := parseSwapctlOutput(invalid)
	assert.Error(t, err)
}

func TestParseSwapctlOutput_Empty(t *testing.T) {
	_, err := parseSwapctlOutput("")
	assert.Error(t, err)
}
