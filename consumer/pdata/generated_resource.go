// Copyright The OpenTelemetry Authors
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

// Code generated by "cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "go run cmd/pdatagen/main.go".

package pdata

import (
	otlpresource "go.opentelemetry.io/collector/internal/data/protogen/resource/v1"
)

// Resource information.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewResource function to create new instances.
// Important: zero-initialized instance is not valid for use.
type Resource struct {
	orig *otlpresource.Resource
}

func newResource(orig *otlpresource.Resource) Resource {
	return Resource{orig: orig}
}

// NewResource creates a new empty Resource.
//
// This must be used only in testing code since no "Set" method available.
func NewResource() Resource {
	return newResource(&otlpresource.Resource{})
}

// Attributes returns the Attributes associated with this Resource.
func (ms Resource) Attributes() AttributeMap {
	return newAttributeMap(&(*ms.orig).Attributes)
}

// CopyTo copies all properties from the current struct to the dest.
func (ms Resource) CopyTo(dest Resource) {
	ms.Attributes().CopyTo(dest.Attributes())
}
