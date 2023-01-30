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

package configschema

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadFields(t *testing.T) {
	fields, err := readFields(
		reflect.ValueOf(testConfig{Name: "foo"}),
		func(v reflect.Value) (map[string]string, error) { return nil, nil },
	)
	require.NoError(t, err)
	assert.Equal(t, "configschema.testConfig", fields.Type)
	nameField := fields.CfgFields[0]
	assert.Equal(t, "name", nameField.Name)
	assert.Equal(t, "string", nameField.Kind)
	assert.Equal(t, "foo", nameField.Default)
}
