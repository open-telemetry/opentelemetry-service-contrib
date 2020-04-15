// Copyright 2020, OpenTelemetry Authors
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

package k8sclusterreceiver

import (
	resourcepb "github.com/census-instrumentation/opencensus-proto/gen-go/resource/v1"
	"go.opencensus.io/resource/resourcekeys"
	appsv1 "k8s.io/api/apps/v1"
)

func getMetricsForReplicaSet(rs *appsv1.ReplicaSet) []*resourceMetrics {
	if rs.Spec.Replicas == nil {
		return nil
	}

	return []*resourceMetrics{
		{
			resource: getResourceForReplicaSet(rs),
			metrics: getReplicaMetrics(
				"replica_set",
				*rs.Spec.Replicas,
				rs.Status.AvailableReplicas,
			),
		},
	}

}

func getResourceForReplicaSet(rs *appsv1.ReplicaSet) *resourcepb.Resource {
	return &resourcepb.Resource{
		Type: resourcekeys.K8SType,
		Labels: map[string]string{
			k8sKeyReplicaSetUID:              string(rs.UID),
			k8sKeyReplicaSetName:             rs.Name,
			resourcekeys.K8SKeyNamespaceName: rs.Namespace,
			resourcekeys.K8SKeyClusterName:   rs.ClusterName,
		},
	}
}

func getMetadataForReplicaSet(rs *appsv1.ReplicaSet) []*KubernetesMetadata {
	return []*KubernetesMetadata{getGenericMetadata(&rs.ObjectMeta, "replicaset")}
}
