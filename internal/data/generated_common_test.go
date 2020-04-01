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

// Code generated by "internal/data_generator/main.go". DO NOT EDIT.
// To regenerate this file run "go run internal/data_generator/main.go".

package data

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestInstrumentationLibrary(t *testing.T) {
	ms := NewInstrumentationLibrary()
	assert.EqualValues(t, true,ms.IsNil())
	ms.InitEmpty()
	assert.EqualValues(t, false, ms.IsNil())

	assert.EqualValues(t, "", ms.Name())
	testValName := "test_name"
	ms.SetName(testValName)
	assert.EqualValues(t, testValName, ms.Name())

	assert.EqualValues(t, "", ms.Version())
	testValVersion := "test_version"
	ms.SetVersion(testValVersion)
	assert.EqualValues(t, testValVersion, ms.Version())

	assert.EqualValues(t, generateTestInstrumentationLibrary(), ms)
}

func generateTestInstrumentationLibrary() InstrumentationLibrary {
	tv := NewInstrumentationLibrary()
	tv.InitEmpty()
	fillTestInstrumentationLibrary(tv)
	return tv
}

func fillTestInstrumentationLibrary(tv InstrumentationLibrary) {
	tv.SetName("test_name")
	tv.SetVersion("test_version")
}
