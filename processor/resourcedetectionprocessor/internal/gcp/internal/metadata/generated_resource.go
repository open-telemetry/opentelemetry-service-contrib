// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// ResourceBuilder is a helper struct to build resources predefined in metadata.yaml.
// The ResourceBuilder is not thread-safe and must not to be used in multiple goroutines.
type ResourceBuilder struct {
	config ResourceAttributesConfig
	res    pcommon.Resource
}

// NewResourceBuilder creates a new ResourceBuilder. This method should be called on the start of the application.
func NewResourceBuilder(rac ResourceAttributesConfig) *ResourceBuilder {
	return &ResourceBuilder{
		config: rac,
		res:    pcommon.NewResource(),
	}
}

// SetCloudAccountID sets provided value as "cloud.account.id" attribute.
func (rb *ResourceBuilder) SetCloudAccountID(val string) {
	if rb.config.CloudAccountID.Enabled {
		rb.res.Attributes().PutStr("cloud.account.id", val)
	}
}

// SetCloudAvailabilityZone sets provided value as "cloud.availability_zone" attribute.
func (rb *ResourceBuilder) SetCloudAvailabilityZone(val string) {
	if rb.config.CloudAvailabilityZone.Enabled {
		rb.res.Attributes().PutStr("cloud.availability_zone", val)
	}
}

// SetCloudPlatform sets provided value as "cloud.platform" attribute.
func (rb *ResourceBuilder) SetCloudPlatform(val string) {
	if rb.config.CloudPlatform.Enabled {
		rb.res.Attributes().PutStr("cloud.platform", val)
	}
}

// SetCloudProvider sets provided value as "cloud.provider" attribute.
func (rb *ResourceBuilder) SetCloudProvider(val string) {
	if rb.config.CloudProvider.Enabled {
		rb.res.Attributes().PutStr("cloud.provider", val)
	}
}

// SetCloudRegion sets provided value as "cloud.region" attribute.
func (rb *ResourceBuilder) SetCloudRegion(val string) {
	if rb.config.CloudRegion.Enabled {
		rb.res.Attributes().PutStr("cloud.region", val)
	}
}

// SetFaasID sets provided value as "faas.id" attribute.
func (rb *ResourceBuilder) SetFaasID(val string) {
	if rb.config.FaasID.Enabled {
		rb.res.Attributes().PutStr("faas.id", val)
	}
}

// SetFaasName sets provided value as "faas.name" attribute.
func (rb *ResourceBuilder) SetFaasName(val string) {
	if rb.config.FaasName.Enabled {
		rb.res.Attributes().PutStr("faas.name", val)
	}
}

// SetFaasVersion sets provided value as "faas.version" attribute.
func (rb *ResourceBuilder) SetFaasVersion(val string) {
	if rb.config.FaasVersion.Enabled {
		rb.res.Attributes().PutStr("faas.version", val)
	}
}

// SetGcpCloudRunJobExecution sets provided value as "gcp.cloud_run.job.execution" attribute.
func (rb *ResourceBuilder) SetGcpCloudRunJobExecution(val string) {
	if rb.config.GcpCloudRunJobExecution.Enabled {
		rb.res.Attributes().PutStr("gcp.cloud_run.job.execution", val)
	}
}

// SetGcpCloudRunJobTaskIndex sets provided value as "gcp.cloud_run.job.task_index" attribute.
func (rb *ResourceBuilder) SetGcpCloudRunJobTaskIndex(val string) {
	if rb.config.GcpCloudRunJobTaskIndex.Enabled {
		rb.res.Attributes().PutStr("gcp.cloud_run.job.task_index", val)
	}
}

// SetGcpGceInstanceHostname sets provided value as "gcp.gce.instance.hostname" attribute.
func (rb *ResourceBuilder) SetGcpGceInstanceHostname(val string) {
	if rb.config.GcpGceInstanceHostname.Enabled {
		rb.res.Attributes().PutStr("gcp.gce.instance.hostname", val)
	}
}

// SetGcpGceInstanceName sets provided value as "gcp.gce.instance.name" attribute.
func (rb *ResourceBuilder) SetGcpGceInstanceName(val string) {
	if rb.config.GcpGceInstanceName.Enabled {
		rb.res.Attributes().PutStr("gcp.gce.instance.name", val)
	}
}

// SetHostID sets provided value as "host.id" attribute.
func (rb *ResourceBuilder) SetHostID(val string) {
	if rb.config.HostID.Enabled {
		rb.res.Attributes().PutStr("host.id", val)
	}
}

// SetHostName sets provided value as "host.name" attribute.
func (rb *ResourceBuilder) SetHostName(val string) {
	if rb.config.HostName.Enabled {
		rb.res.Attributes().PutStr("host.name", val)
	}
}

// SetHostType sets provided value as "host.type" attribute.
func (rb *ResourceBuilder) SetHostType(val string) {
	if rb.config.HostType.Enabled {
		rb.res.Attributes().PutStr("host.type", val)
	}
}

// SetK8sClusterName sets provided value as "k8s.cluster.name" attribute.
func (rb *ResourceBuilder) SetK8sClusterName(val string) {
	if rb.config.K8sClusterName.Enabled {
		rb.res.Attributes().PutStr("k8s.cluster.name", val)
	}
}

// Emit returns the built resource and resets the internal builder state.
func (rb *ResourceBuilder) Emit() pcommon.Resource {
	r := rb.res
	rb.res = pcommon.NewResource()
	return r
}
