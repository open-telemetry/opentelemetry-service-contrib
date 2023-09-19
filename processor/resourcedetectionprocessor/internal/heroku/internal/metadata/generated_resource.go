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

// SetCloudProvider sets provided value as "cloud.provider" attribute.
func (rb *ResourceBuilder) SetCloudProvider(val string) {
	if rb.config.CloudProvider.Enabled {
		rb.res.Attributes().PutStr("cloud.provider", val)
	}
}

// SetHerokuAppID sets provided value as "heroku.app.id" attribute.
func (rb *ResourceBuilder) SetHerokuAppID(val string) {
	if rb.config.HerokuAppID.Enabled {
		rb.res.Attributes().PutStr("heroku.app.id", val)
	}
}

// SetHerokuDynoID sets provided value as "heroku.dyno.id" attribute.
func (rb *ResourceBuilder) SetHerokuDynoID(val string) {
	if rb.config.HerokuDynoID.Enabled {
		rb.res.Attributes().PutStr("heroku.dyno.id", val)
	}
}

// SetHerokuReleaseCommit sets provided value as "heroku.release.commit" attribute.
func (rb *ResourceBuilder) SetHerokuReleaseCommit(val string) {
	if rb.config.HerokuReleaseCommit.Enabled {
		rb.res.Attributes().PutStr("heroku.release.commit", val)
	}
}

// SetHerokuReleaseCreationTimestamp sets provided value as "heroku.release.creation_timestamp" attribute.
func (rb *ResourceBuilder) SetHerokuReleaseCreationTimestamp(val string) {
	if rb.config.HerokuReleaseCreationTimestamp.Enabled {
		rb.res.Attributes().PutStr("heroku.release.creation_timestamp", val)
	}
}

// SetServiceInstanceID sets provided value as "service.instance.id" attribute.
func (rb *ResourceBuilder) SetServiceInstanceID(val string) {
	if rb.config.ServiceInstanceID.Enabled {
		rb.res.Attributes().PutStr("service.instance.id", val)
	}
}

// SetServiceName sets provided value as "service.name" attribute.
func (rb *ResourceBuilder) SetServiceName(val string) {
	if rb.config.ServiceName.Enabled {
		rb.res.Attributes().PutStr("service.name", val)
	}
}

// SetServiceVersion sets provided value as "service.version" attribute.
func (rb *ResourceBuilder) SetServiceVersion(val string) {
	if rb.config.ServiceVersion.Enabled {
		rb.res.Attributes().PutStr("service.version", val)
	}
}

// Emit returns the built resource and resets the internal builder state.
func (rb *ResourceBuilder) Emit() pcommon.Resource {
	r := rb.res
	rb.res = pcommon.NewResource()
	return r
}
