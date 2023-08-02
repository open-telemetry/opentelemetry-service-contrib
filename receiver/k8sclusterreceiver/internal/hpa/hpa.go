// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package hpa // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/hpa"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	autoscalingv2 "k8s.io/api/autoscaling/v2"
	autoscalingv2beta2 "k8s.io/api/autoscaling/v2beta2"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/metadata"
)

func RecordMetricsBeta(mb *metadata.MetricsBuilder, hpa *autoscalingv2beta2.HorizontalPodAutoscaler, ts pcommon.Timestamp) {
	mb.RecordK8sHpaMaxReplicasDataPoint(ts, int64(hpa.Spec.MaxReplicas))
	mb.RecordK8sHpaMinReplicasDataPoint(ts, int64(*hpa.Spec.MinReplicas))
	mb.RecordK8sHpaCurrentReplicasDataPoint(ts, int64(hpa.Status.CurrentReplicas))
	mb.RecordK8sHpaDesiredReplicasDataPoint(ts, int64(hpa.Status.DesiredReplicas))
	rb := mb.NewResourceBuilder()
	rb.SetK8sHpaUID(string(hpa.UID))
	rb.SetK8sHpaName(hpa.Name)
	rb.SetK8sNamespaceName(hpa.Namespace)
	mb.EmitForResource(metadata.WithResource(rb.Emit()))
}

func RecordMetrics(mb *metadata.MetricsBuilder, hpa *autoscalingv2.HorizontalPodAutoscaler, ts pcommon.Timestamp) {
	mb.RecordK8sHpaMaxReplicasDataPoint(ts, int64(hpa.Spec.MaxReplicas))
	mb.RecordK8sHpaMinReplicasDataPoint(ts, int64(*hpa.Spec.MinReplicas))
	mb.RecordK8sHpaCurrentReplicasDataPoint(ts, int64(hpa.Status.CurrentReplicas))
	mb.RecordK8sHpaDesiredReplicasDataPoint(ts, int64(hpa.Status.DesiredReplicas))
	rb := mb.NewResourceBuilder()
	rb.SetK8sHpaUID(string(hpa.UID))
	rb.SetK8sHpaName(hpa.Name)
	rb.SetK8sNamespaceName(hpa.Namespace)
	mb.EmitForResource(metadata.WithResource(rb.Emit()))
}

func GetMetadata(hpa *autoscalingv2.HorizontalPodAutoscaler) map[experimentalmetricmetadata.ResourceID]*metadata.KubernetesMetadata {
	return map[experimentalmetricmetadata.ResourceID]*metadata.KubernetesMetadata{
		experimentalmetricmetadata.ResourceID(hpa.UID): metadata.GetGenericMetadata(&hpa.ObjectMeta, "HPA"),
	}
}

func GetMetadataBeta(hpa *autoscalingv2beta2.HorizontalPodAutoscaler) map[experimentalmetricmetadata.ResourceID]*metadata.KubernetesMetadata {
	return map[experimentalmetricmetadata.ResourceID]*metadata.KubernetesMetadata{
		experimentalmetricmetadata.ResourceID(hpa.UID): metadata.GetGenericMetadata(&hpa.ObjectMeta, "HPA"),
	}
}
