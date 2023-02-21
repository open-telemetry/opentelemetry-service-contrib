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

//go:build windows
// +build windows

package system // import "github.com/asserts/opentelemetry-collector-contrib/exporter/datadogexporter/internal/metadata/internal/system"

func getSystemFQDN() (string, error) {
	// The Datadog Agent uses CGo to get the FQDN of the host
	// OpenTelemetry does not allow the use of CGo so this feature
	// is disabled on Windows
	return "", nil
}
