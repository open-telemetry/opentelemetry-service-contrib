// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package testhelpers // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextensionv2/internal/testhelpers"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextensionv2/internal/status"
)

// PipelineMetadata groups together component and instance IDs for a hypothetical pipeline used
// for testing purposes.
type PipelineMetadata struct {
	PipelineID  component.ID
	ReceiverID  *component.InstanceID
	ProcessorID *component.InstanceID
	ExporterID  *component.InstanceID
}

// InstanceIDs returns a slice of instanceIDs for components within the hypothetical pipeline.
func (p *PipelineMetadata) InstanceIDs() []*component.InstanceID {
	return []*component.InstanceID{p.ReceiverID, p.ProcessorID, p.ExporterID}
}

// NewPipelineMetadata returns a metadata for a hypothetical pipeline.
func NewPipelineMetadata(typeVal component.Type) *PipelineMetadata {
	pipelineID := component.NewID(typeVal)
	return &PipelineMetadata{
		PipelineID: pipelineID,
		ReceiverID: &component.InstanceID{
			ID:   component.NewIDWithName(typeVal, "in"),
			Kind: component.KindReceiver,
			PipelineIDs: map[component.ID]struct{}{
				pipelineID: {},
			},
		},
		ProcessorID: &component.InstanceID{
			ID:   component.NewID("batch"),
			Kind: component.KindProcessor,
			PipelineIDs: map[component.ID]struct{}{
				pipelineID: {},
			},
		},
		ExporterID: &component.InstanceID{
			ID:   component.NewIDWithName(typeVal, "out"),
			Kind: component.KindExporter,
			PipelineIDs: map[component.ID]struct{}{
				pipelineID: {},
			},
		},
	}
}

// NewPipelines returns a map of hypothetical pipelines identified by their stringified typeVal.
func NewPipelines(typeVals ...component.Type) map[string]*PipelineMetadata {
	result := make(map[string]*PipelineMetadata, len(typeVals))
	for _, val := range typeVals {
		result[string(val)] = NewPipelineMetadata(val)
	}
	return result
}

// SeedAggregator records a status event for each instanceID.
func SeedAggregator(
	agg *status.Aggregator,
	instanceIDs []*component.InstanceID,
	statuses ...component.Status,
) {
	for _, st := range statuses {
		for _, id := range instanceIDs {
			agg.RecordStatus(id, component.NewStatusEvent(st))
		}
	}
}
