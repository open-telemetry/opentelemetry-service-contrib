// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package attrs

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/filetest"
)

func TestResolver(t *testing.T) {
	t.Parallel()

	for i := 0; i < 32; i++ {

		// Create a 4 bit string where each bit represents the value of a config option
		bitString := fmt.Sprintf("%05b", i)

		// Create a resolver with a config that matches the bit pattern of i
		r := Resolver{
			IncludeFileName:         bitString[0] == '1',
			IncludeFilePath:         bitString[1] == '1',
			IncludeFileNameResolved: bitString[2] == '1',
			IncludeFilePathResolved: bitString[3] == '1',
			IncludeFileInfos:        bitString[4] == '1',
		}

		t.Run(bitString, func(t *testing.T) {
			// Create a file
			tempDir := t.TempDir()
			temp := filetest.OpenTemp(t, tempDir)

			attributes, err := r.Resolve(temp.Name())
			assert.NoError(t, err)

			var expectLen int
			if r.IncludeFileName {
				expectLen++
				assert.Equal(t, filepath.Base(temp.Name()), attributes[LogFileName])
			} else {
				assert.Empty(t, attributes[LogFileName])
			}
			if r.IncludeFilePath {
				expectLen++
				assert.Equal(t, temp.Name(), attributes[LogFilePath])
			} else {
				assert.Empty(t, attributes[LogFilePath])
			}

			// We don't have an independent way to resolve the path, so the only meangingful validate
			// is to ensure that the resolver returns nothing vs something based on the config.
			if r.IncludeFileNameResolved {
				expectLen++
				assert.NotNil(t, attributes[LogFileNameResolved])
				assert.IsType(t, "", attributes[LogFileNameResolved])
			} else {
				assert.Empty(t, attributes[LogFileNameResolved])
			}
			if r.IncludeFilePathResolved {
				expectLen++
				assert.NotNil(t, attributes[LogFilePathResolved])
				assert.IsType(t, "", attributes[LogFilePathResolved])
			} else {
				assert.Empty(t, attributes[LogFilePathResolved])
			}
			if r.IncludeFileInfos {
				expectLen++
				assert.NotNil(t, attributes[LogFileOwner])
				assert.IsType(t, "", attributes[LogFileOwner])
				expectLen++
				assert.NotNil(t, attributes[LogFileGroup])
				assert.IsType(t, "", attributes[LogFileGroup])
			} else {
				assert.Empty(t, attributes[LogFileOwner])
				assert.Empty(t, attributes[LogFileOwner])
			}
			assert.Equal(t, expectLen, len(attributes))
		})
	}
}
