// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package demonset // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/demonset"

import (
	agentmetricspb "github.com/census-instrumentation/opencensus-proto/gen-go/agent/metrics/v1"
	metricspb "github.com/census-instrumentation/opencensus-proto/gen-go/metrics/v1"
	resourcepb "github.com/census-instrumentation/opencensus-proto/gen-go/resource/v1"
	conventions "go.opentelemetry.io/collector/semconv/v1.6.1"
	appsv1 "k8s.io/api/apps/v1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/constants"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/utils"
)

var daemonSetCurrentScheduledMetric = &metricspb.MetricDescriptor{
	Name:        "k8s.daemonset.current_scheduled_nodes",
	Description: "Number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod",
	Unit:        "1",
	Type:        metricspb.MetricDescriptor_GAUGE_INT64,
}

var daemonSetDesiredScheduledMetric = &metricspb.MetricDescriptor{
	Name:        "k8s.daemonset.desired_scheduled_nodes",
	Description: "Number of nodes that should be running the daemon pod (including nodes currently running the daemon pod)",
	Unit:        "1",
	Type:        metricspb.MetricDescriptor_GAUGE_INT64,
}

var daemonSetMisScheduledMetric = &metricspb.MetricDescriptor{
	Name:        "k8s.daemonset.misscheduled_nodes",
	Description: "Number of nodes that are running the daemon pod, but are not supposed to run the daemon pod",
	Unit:        "1",
	Type:        metricspb.MetricDescriptor_GAUGE_INT64,
}

var daemonSetReadyMetric = &metricspb.MetricDescriptor{
	Name:        "k8s.daemonset.ready_nodes",
	Description: "Number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready",
	Unit:        "1",
	Type:        metricspb.MetricDescriptor_GAUGE_INT64,
}

// Transform transforms the pod to remove the fields that we don't use to reduce RAM utilization.
// IMPORTANT: Make sure to update this function before using new daemonset fields.
func Transform(ds *appsv1.DaemonSet) *appsv1.DaemonSet {
	return &appsv1.DaemonSet{
		ObjectMeta: metadata.TransformObjectMeta(ds.ObjectMeta),
		Status: appsv1.DaemonSetStatus{
			CurrentNumberScheduled: ds.Status.CurrentNumberScheduled,
			DesiredNumberScheduled: ds.Status.DesiredNumberScheduled,
			NumberMisscheduled:     ds.Status.NumberMisscheduled,
			NumberReady:            ds.Status.NumberReady,
		},
	}
}

func GetMetrics(ds *appsv1.DaemonSet) []*agentmetricspb.ExportMetricsServiceRequest {
	metrics := []*metricspb.Metric{
		{
			MetricDescriptor: daemonSetCurrentScheduledMetric,
			Timeseries: []*metricspb.TimeSeries{
				utils.GetInt64TimeSeries(int64(ds.Status.CurrentNumberScheduled)),
			},
		},
		{
			MetricDescriptor: daemonSetDesiredScheduledMetric,
			Timeseries: []*metricspb.TimeSeries{
				utils.GetInt64TimeSeries(int64(ds.Status.DesiredNumberScheduled)),
			},
		},
		{
			MetricDescriptor: daemonSetMisScheduledMetric,
			Timeseries: []*metricspb.TimeSeries{
				utils.GetInt64TimeSeries(int64(ds.Status.NumberMisscheduled)),
			},
		},
		{
			MetricDescriptor: daemonSetReadyMetric,
			Timeseries: []*metricspb.TimeSeries{
				utils.GetInt64TimeSeries(int64(ds.Status.NumberReady)),
			},
		},
	}

	return []*agentmetricspb.ExportMetricsServiceRequest{
		{
			Resource: getResource(ds),
			Metrics:  metrics,
		},
	}
}

func getResource(ds *appsv1.DaemonSet) *resourcepb.Resource {
	return &resourcepb.Resource{
		Type: constants.K8sType,
		Labels: map[string]string{
			conventions.AttributeK8SDaemonSetUID:  string(ds.UID),
			conventions.AttributeK8SDaemonSetName: ds.Name,
			conventions.AttributeK8SNamespaceName: ds.Namespace,
		},
	}
}

func GetMetadata(ds *appsv1.DaemonSet) map[experimentalmetricmetadata.ResourceID]*metadata.KubernetesMetadata {
	return map[experimentalmetricmetadata.ResourceID]*metadata.KubernetesMetadata{
		experimentalmetricmetadata.ResourceID(ds.UID): metadata.GetGenericMetadata(&ds.ObjectMeta, constants.K8sKindDaemonSet),
	}
}
