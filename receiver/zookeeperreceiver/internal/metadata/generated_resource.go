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

// SetServerState sets provided value as "server.state" attribute.
func (rb *ResourceBuilder) SetServerState(val string) {
	if rb.config.ServerState.Enabled {
		rb.res.Attributes().PutStr("server.state", val)
	}
}

// SetZkVersion sets provided value as "zk.version" attribute.
func (rb *ResourceBuilder) SetZkVersion(val string) {
	if rb.config.ZkVersion.Enabled {
		rb.res.Attributes().PutStr("zk.version", val)
	}
}

// Emit returns the built resource and resets the internal builder state.
func (rb *ResourceBuilder) Emit() pcommon.Resource {
	r := rb.res
	rb.res = pcommon.NewResource()
	return r
}
