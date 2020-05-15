// Copyright 2019 OpenTelemetry Authors
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

package configtest

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	"go.opentelemetry.io/collector/config"
)

// NewViperFromYamlFile creates a viper instance that reads the given fileName as yaml config
// and can then be used to unmarshal the file contents to objects.
// Example usage for testing can be found in configtest_test.go
func NewViperFromYamlFile(t *testing.T, fileName string) *viper.Viper {
	// Open the file for reading.
	file, err := os.Open(filepath.Clean(fileName))
	require.NoErrorf(t, err, "unable to open the file %v", fileName)
	require.NotNil(t, file)

	defer func() {
		require.NoErrorf(t, file.Close(), "unable to close the file %v", fileName)
	}()

	// Read yaml config from file
	v := config.NewViper()
	v.SetConfigType("yaml")
	require.NoErrorf(t, v.ReadConfig(file), "unable to read the file %v", fileName)

	return v
}
