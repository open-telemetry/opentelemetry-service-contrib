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

package migrate

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func TestConditionalAttributeSetApply(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name   string
		cond   *ConditionalAttributeSet
		check  string
		attr   pcommon.Map
		expect pcommon.Map
	}{
		{
			name:  "No changes defined",
			cond:  NewConditionalAttributeSet[string, string, string](map[string]string{}),
			check: "database operation",
			attr: testHelperBuildMap(func(m pcommon.Map) {
				m.PutStr("service.version", "v0.0.0")
			}),
			expect: testHelperBuildMap(func(m pcommon.Map) {
				m.PutStr("service.version", "v0.0.0")
			}),
		},
		{
			name: "Not matched check value",
			cond: NewConditionalAttributeSet(
				map[string]string{
					"service.version": "application.version",
				},
				"application start",
			),
			check: "datatbase operation",
			attr: testHelperBuildMap(func(m pcommon.Map) {
				m.PutStr("service.version", "v0.0.0")
			}),
			expect: testHelperBuildMap(func(m pcommon.Map) {
				m.PutStr("service.version", "v0.0.0")
			}),
		},
		{
			name: "No condition set, applys to all",
			cond: NewConditionalAttributeSet[string, string, string](
				map[string]string{
					"service.version": "application.version",
				},
			),
			check: "datatbase operation",
			attr: testHelperBuildMap(func(m pcommon.Map) {
				m.PutStr("service.version", "v0.0.0")
			}),
			expect: testHelperBuildMap(func(m pcommon.Map) {
				m.PutStr("application.version", "v0.0.0")
			}),
		},
		{
			name: "Matched condition, setting value",
			cond: NewConditionalAttributeSet(
				map[string]string{
					"service.version": "application.version",
				},
				"application start",
				"application stop",
			),
			check: "application start",
			attr: testHelperBuildMap(func(m pcommon.Map) {
				m.PutStr("service.version", "v0.0.0")
			}),
			expect: testHelperBuildMap(func(m pcommon.Map) {
				m.PutStr("application.version", "v0.0.0")
			}),
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.NoError(t, tc.cond.Apply(tc.attr, tc.check))
			assert.Equal(t, tc.expect.AsRaw(), tc.attr.AsRaw(), "Must match the expected value")
		})
	}
}

func TestConditionalAttributeSetRollback(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name   string
		cond   *ConditionalAttributeSet
		check  string
		attr   pcommon.Map
		expect pcommon.Map
	}{
		{
			name:  "No changes defined",
			cond:  NewConditionalAttributeSet[string, string, string](map[string]string{}),
			check: "database operation",
			attr: testHelperBuildMap(func(m pcommon.Map) {
				m.PutStr("service.version", "v0.0.0")
			}),
			expect: testHelperBuildMap(func(m pcommon.Map) {
				m.PutStr("service.version", "v0.0.0")
			}),
		},
		{
			name: "Not matched check value",
			cond: NewConditionalAttributeSet(
				map[string]string{
					"service.version": "application.version",
				},
				"application start",
			),
			check: "datatbase operation",
			attr: testHelperBuildMap(func(m pcommon.Map) {
				m.PutStr("service.version", "v0.0.0")
			}),
			expect: testHelperBuildMap(func(m pcommon.Map) {
				m.PutStr("service.version", "v0.0.0")
			}),
		},
		{
			name: "No condition set, applys to all",
			cond: NewConditionalAttributeSet[string, string, string](
				map[string]string{
					"service.version": "application.version",
				},
			),
			check: "datatbase operation",
			attr: testHelperBuildMap(func(m pcommon.Map) {
				m.PutStr("application.version", "v0.0.0")
			}),
			expect: testHelperBuildMap(func(m pcommon.Map) {
				m.PutStr("service.version", "v0.0.0")
			}),
		},
		{
			name: "Matched condition, setting value",
			cond: NewConditionalAttributeSet(
				map[string]string{
					"service.version": "application.version",
				},
				"application start",
				"application stop",
			),
			check: "application start",
			attr: testHelperBuildMap(func(m pcommon.Map) {
				m.PutStr("application.version", "v0.0.0")
			}),
			expect: testHelperBuildMap(func(m pcommon.Map) {
				m.PutStr("service.version", "v0.0.0")
			}),
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.NoError(t, tc.cond.Rollback(tc.attr, tc.check))
			assert.Equal(t, tc.expect.AsRaw(), tc.attr.AsRaw(), "Must match the expected value")
		})
	}
}

func TestConditionalAttribueSetSliceApply(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name   string
		slice  *ConditionalAttributeSetSlice
		check  string
		attrs  pcommon.Map
		expect pcommon.Map
	}{
		{
			name:  "No changes",
			slice: NewConditionalAttributeSetSlice(),
			check: "application start",
			attrs: testHelperBuildMap(func(m pcommon.Map) {
				m.PutStr("service.version", "v0.0.0")
			}),
			expect: testHelperBuildMap(func(m pcommon.Map) {
				m.PutStr("service.version", "v0.0.0")
			}),
		},
		{
			name: "Not matched check value",
			slice: NewConditionalAttributeSetSlice(
				NewConditionalAttributeSet[string, string, string](
					map[string]string{
						"service_version": "service.version",
					},
				),
				// intentially silly to be make it clear
				// that this should not be applied
				NewConditionalAttributeSet(
					map[string]string{
						"service.version": "shark.attack",
					},
					"shark spotted",
				),
			),
			check: "application start",
			attrs: testHelperBuildMap(func(m pcommon.Map) {
				m.PutStr("service_version", "v0.0.0")
			}),
			expect: testHelperBuildMap(func(m pcommon.Map) {
				m.PutStr("service.version", "v0.0.0")
			}),
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.NoError(t, tc.slice.Apply(tc.attrs, tc.check))
			assert.Equal(t, tc.expect.AsRaw(), tc.attrs.AsRaw(), "Must match the expected values")
		})
	}
}

func TestConditionalAttribueSetSliceRollback(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name   string
		slice  *ConditionalAttributeSetSlice
		check  string
		attrs  pcommon.Map
		expect pcommon.Map
	}{
		{
			name:  "No changes",
			slice: NewConditionalAttributeSetSlice(),
			check: "application start",
			attrs: testHelperBuildMap(func(m pcommon.Map) {
				m.PutStr("service.version", "v0.0.0")
			}),
			expect: testHelperBuildMap(func(m pcommon.Map) {
				m.PutStr("service.version", "v0.0.0")
			}),
		},
		{
			name: "Not matched check value",
			slice: NewConditionalAttributeSetSlice(
				NewConditionalAttributeSet[string, string, string](
					map[string]string{
						"service_version": "service.version",
					},
				),
				// intentially silly to be make it clear
				// that this should not be applied
				NewConditionalAttributeSet(
					map[string]string{
						"service.version": "shark.attack",
					},
					"shark spotted",
				),
			),
			check: "application start",
			attrs: testHelperBuildMap(func(m pcommon.Map) {
				m.PutStr("service.version", "v0.0.0")
			}),
			expect: testHelperBuildMap(func(m pcommon.Map) {
				m.PutStr("service_version", "v0.0.0")
			}),
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.NoError(t, tc.slice.Rollback(tc.attrs, tc.check))
			assert.Equal(t, tc.expect.AsRaw(), tc.attrs.AsRaw(), "Must match the expected values")
		})
	}
}
