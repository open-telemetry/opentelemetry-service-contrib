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

package conventions

// OpenTelemetry Semantic Convention values for general Span attribute names.
// See: https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/semantic_conventions/span-general.md
const (
	AttributeComponent    = "component"
	AttributeEnduserID    = "enduser.id"
	AttributeEnduserRole  = "enduser.role"
	AttributeEnduserScope = "enduser.scope"
	AttributeNetHostIP    = "net.host.ip"
	AttributeNetHostName  = "net.host.name"
	AttributeNetHostPort  = "net.host.port"
	AttributeNetPeerIP    = "net.peer.ip"
	AttributeNetPeerName  = "net.peer.name"
	AttributeNetPeerPort  = "net.peer.port"
	AttributeNetTransport = "net.transport"
	AttributePeerService  = "peer.service"
)

// OpenTelemetry Semantic Convention values for component attribute values.
// Possibly being removed due to issue #336
const (
	ComponentTypeHTTP = "http"
	ComponentTypeGRPC = "grpc"
)
