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

package kubernetesclusterreceiver

import (
	resourcepb "github.com/census-instrumentation/opencensus-proto/gen-go/resource/v1"
	"go.opencensus.io/resource/resourcekeys"
	corev1 "k8s.io/api/core/v1"
)

func getMetricsForReplicationController(rc *corev1.ReplicationController) []*resourceMetrics {
	if rc.Spec.Replicas == nil {
		return []*resourceMetrics{}
	}

	return []*resourceMetrics{
		{
			resource: getResourceForReplicationController(rc),
			metrics: getReplicaMetrics(
				"replication_controller",
				*rc.Spec.Replicas,
				rc.Status.AvailableReplicas,
			),
		},
	}

}

func getResourceForReplicationController(rc *corev1.ReplicationController) *resourcepb.Resource {
	return &resourcepb.Resource{
		Type: resourcekeys.K8SType,
		Labels: map[string]string{
			k8sKeyReplicationControllerUID:   string(rc.UID),
			k8sKeyReplicationControllerName:  rc.Name,
			resourcekeys.K8SKeyNamespaceName: rc.Namespace,
			resourcekeys.K8SKeyClusterName:   rc.ClusterName,
		},
	}
}

func getMetadataForReplicationController(rc *corev1.ReplicationController) []*KubernetesMetadata {
	return []*KubernetesMetadata{getGenericMetadata(&rc.ObjectMeta, "replicationcontroller")}
}
