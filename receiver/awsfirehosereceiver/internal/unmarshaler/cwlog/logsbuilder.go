// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cwlog // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver/internal/unmarshaler/cwlog"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	conventions "go.opentelemetry.io/collector/semconv/v1.6.1"
)

// resourceAttributes are the CloudWatch log attributes that define a unique resource.
type resourceAttributes struct {
	owner, logGroup, logStream string
}

// resourceLogsBuilder provides convenient access to the a Resource's LogRecordSlice.
type resourceLogsBuilder struct {
	rls plog.LogRecordSlice
}

// setAttributes applies the resourceAttributes to the provided Resource.
func (ra *resourceAttributes) setAttributes(resource pcommon.Resource) {
	attrs := resource.Attributes()
	attrs.PutStr(conventions.AttributeCloudAccountID, ra.owner)
	attrs.PutStr("cloudwatch.log.group.name", ra.logStream)
	attrs.PutStr("cloudwatch.log.stream", ra.logGroup)
}

// newResourceLogsBuilder to capture logs for the Resource defined by the provided attributes.
func newResourceLogsBuilder(logs plog.Logs, attrs resourceAttributes) *resourceLogsBuilder {
	rls := logs.ResourceLogs().AppendEmpty()
	attrs.setAttributes(rls.Resource())
	return &resourceLogsBuilder{rls.ScopeLogs().AppendEmpty().LogRecords()}
}

// AddLog events to the LogRecordSlice. Resource attributes are captured when creating
// the resourceLogsBuilder, so we only need to consider the LogEvents themselves.
func (rlb *resourceLogsBuilder) AddLog(log cWLog) {
	for _, event := range log.LogEvents {
		logLine := rlb.rls.AppendEmpty()
		logLine.SetTimestamp(pcommon.Timestamp(event.Timestamp))
		logLine.Body().SetStr(event.Message)
	}
}
