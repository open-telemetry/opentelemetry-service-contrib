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
package cluster // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/aerospikereceiver/cluster"

import (
	as "github.com/aerospike/aerospike-client-go/v5"
)

type SubsetCluster struct {
	nodes []Node
}

type nodeFactoryFunc func(*as.ClientPolicy, *as.Host, bool) (Node, error)

func NewSubsetCluster(policy *as.ClientPolicy, hosts []*as.Host, authEnabled bool) (*SubsetCluster, error) {
	return newSubsetCluster(policy, hosts, authEnabled, newConnNode)
}

func newSubsetCluster(policy *as.ClientPolicy, hosts []*as.Host, authEnabled bool, nodeFact nodeFactoryFunc) (*SubsetCluster, error) {
	nodes := make([]Node, len(hosts))

	// this is only used with 1 node for now (when collect-cluster-metrics is false)
	for i := range hosts {
		n, err := nodeFact(policy, hosts[i], authEnabled)
		if err != nil {
			return nil, err
		}

		nodes[i] = n
	}

	res := SubsetCluster{
		nodes: nodes,
	}

	return &res, nil
}

func (c *SubsetCluster) Close() {
	for _, node := range c.nodes {
		node.Close()
	}
}

func (c *SubsetCluster) GetNodes() []Node {
	return c.nodes
}
