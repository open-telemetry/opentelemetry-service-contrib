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

package data

// Resource information.
//
// Must use NewResource functions to create new instances.
// Important: zero-initialized instance is not valid for use.
type Resource struct {
	// Set of attributes that describe the resource.
	attributes AttributesMap
}

func NewResource() *Resource {
	return &Resource{}
}

func (m *Resource) Attributes() AttributesMap {
	return m.attributes
}

func (m *Resource) SetAttributes(v AttributesMap) {
	m.attributes = v
}
