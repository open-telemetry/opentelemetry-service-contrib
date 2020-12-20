// Copyright 2020 OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package metricstransformprocessor

import (
	"go.opentelemetry.io/collector/config/configmodels"
)

const (
	// IncludeFieldName is the mapstructure field name for Include field
	IncludeFieldName = "include"

	// MatchTypeFieldName is the mapstructure field name for MatchType field
	MatchTypeFieldName = "match_type"

	// MetricNameFieldName is the mapstructure field name for MetricName field
	MetricNameFieldName = "metric_name"

	// ActionFieldName is the mapstructure field name for Action field
	ActionFieldName = "action"

	// NewNameFieldName is the mapstructure field name for NewName field
	NewNameFieldName = "new_name"

	//
	GroupName = "group_name"

	// AggregationTypeFieldName is the mapstructure field name for AggregationType field
	AggregationTypeFieldName = "aggregation_type"

	// LabelFieldName is the mapstructure field name for Label field
	LabelFieldName = "label"

	// NewLabelFieldName is the mapstructure field name for NewLabel field
	NewLabelFieldName = "new_label"

	// NewValueFieldName is the mapstructure field name for NewValue field
	NewValueFieldName = "new_value"

	// SubmatchCaseFieldName is the mapstructure field name for SubmatchCase field
	SubmatchCaseFieldName = "submatch_case"
)

// Config defines configuration for Resource processor.
type Config struct {
	configmodels.ProcessorSettings `mapstructure:",squash"`

	// Transform specifies a list of transforms on metrics with each transform focusing on one metric.
	Transforms []Transform `mapstructure:"transforms"`
}

// Transform defines the transformation applied to the specific metric
type Transform struct {
	// MetricIncludeFilter is used to select the metric(s) to operate on.
	// REQUIRED
	MetricIncludeFilter FilterConfig `mapstructure:",squash"`

	// MetricName is used to select the metric to operate on.
	// DEPRECATED. Use MetricIncludeFilter instead.
	MetricName string `mapstructure:"metric_name"`

	// Action specifies the action performed on the matched metric.
	// REQUIRED
	Action ConfigAction `mapstructure:"action"`

	// NewName specifies the name of the new metric when inserting or updating.
	// REQUIRED only if Action is INSERT.
	NewName string `mapstructure:"new_name"`

	GroupName string `mapstructure:"group_name"`

	// AggregationType specifies how to aggregate.
	// REQUIRED only if Action is COMBINE.
	AggregationType AggregationType `mapstructure:"aggregation_type"`

	// SubmatchCase specifies what case to use for label values created from regexp submatches.
	SubmatchCase SubmatchCase `mapstructure:"submatch_case"`

	// Operations contains a list of operations that will be performed on the selected metric.
	Operations []Operation `mapstructure:"operations"`
}

type FilterConfig struct {
	// Include specifies the metric(s) to operate on.
	Include string `mapstructure:"include"`

	// MatchType determines how the Include string is matched: <strict|regexp>.
	MatchType MatchType `mapstructure:"match_type"`
}

// Operation defines the specific operation performed on the selected metrics.
type Operation struct {
	// Action specifies the action performed for this operation.
	// REQUIRED
	Action OperationAction `mapstructure:"action"`

	// Label identifies the exact label to operate on.
	Label string `mapstructure:"label"`

	// NewLabel determines the name to rename the identified label to.
	NewLabel string `mapstructure:"new_label"`

	// LabelSet is a list of labels to keep. All other labels are aggregated based on the AggregationType.
	LabelSet []string `mapstructure:"label_set"`

	// AggregationType specifies how to aggregate.
	AggregationType AggregationType `mapstructure:"aggregation_type"`

	// AggregatedValues is a list of label values to aggregate away.
	AggregatedValues []string `mapstructure:"aggregated_values"`

	// NewValue is used to set a new label value either when the operation is `AggregatedValues` or `AddLabel`.
	NewValue string `mapstructure:"new_value"`

	// ValueActions is a list of renaming actions for label values.
	ValueActions []ValueAction `mapstructure:"value_actions"`

	// LabelValue identifies the exact label value to operate on
	LabelValue string `mapstructure:"label_value"`
}

// ValueAction renames label values.
type ValueAction struct {
	// Value specifies the current label value.
	Value string `mapstructure:"value"`

	// NewValue specifies the label value to rename to.
	NewValue string `mapstructure:"new_value"`
}

// ConfigAction is the enum to capture the type of action to perform on a metric.
type ConfigAction string

const (
	// Insert adds a new metric to the batch with a new name.
	Insert ConfigAction = "insert"

	// Update updates an existing metric.
	Update ConfigAction = "update"

	// Combine combines multiple metrics into a single metric.
	Combine ConfigAction = "combine"

	//
	Group ConfigAction = "group"
)

var Actions = []ConfigAction{Insert, Update, Combine, Group}

func (ca ConfigAction) isValid() bool {
	for _, configAction := range Actions {
		if ca == configAction {
			return true
		}
	}

	return false
}

// OperationAction is the enum to capture the thress types of actions to perform for an operation.
type OperationAction string

const (
	// AddLabel adds a new label to an existing metric.
	AddLabel OperationAction = "add_label"

	// UpdateLabel applies name changes to label and/or label values.
	UpdateLabel OperationAction = "update_label"

	// DeleteLabelValue deletes a label value by also removing all the points associated with this label value
	DeleteLabelValue OperationAction = "delete_label_value"

	// ToggleScalarDataType changes the data type from int64 to double, or vice-versa
	ToggleScalarDataType OperationAction = "toggle_scalar_data_type"

	// AggregateLabels aggregates away all labels other than the ones in Operation.LabelSet
	// by the method indicated by Operation.AggregationType.
	AggregateLabels OperationAction = "aggregate_labels"

	// AggregateLabelValues aggregates away the values in Operation.AggregatedValues
	// by the method indicated by Operation.AggregationType.
	AggregateLabelValues OperationAction = "aggregate_label_values"
)

var OperationActions = []OperationAction{AddLabel, UpdateLabel, DeleteLabelValue, ToggleScalarDataType, AggregateLabels, AggregateLabelValues}

func (oa OperationAction) isValid() bool {
	for _, operationAction := range OperationActions {
		if oa == operationAction {
			return true
		}
	}

	return false
}

// AggregationType is the enum to capture the three types of aggregation for the aggregation operation.
type AggregationType string

const (
	// Sum indicates taking the sum of the aggregated data.
	Sum AggregationType = "sum"

	// Mean indicates taking the mean of the aggregated data.
	Mean AggregationType = "mean"

	// Min indicates taking the minimum of the aggregated data.
	Min AggregationType = "min"

	// Max indicates taking the max of the aggregated data.
	Max AggregationType = "max"
)

var AggregationTypes = []AggregationType{Sum, Mean, Min, Max}

func (at AggregationType) isValid() bool {
	for _, aggregationType := range AggregationTypes {
		if at == aggregationType {
			return true
		}
	}

	return false
}

// MatchType is the enum to capture the two types of matching metric(s) that should have operations applied to them.
type MatchType string

const (
	// StrictMatchType is the FilterType for filtering by exact string matches.
	StrictMatchType MatchType = "strict"

	// RegexpMatchType is the FilterType for filtering by regexp string matches.
	RegexpMatchType MatchType = "regexp"
)

var MatchTypes = []MatchType{StrictMatchType, RegexpMatchType}

func (mt MatchType) isValid() bool {
	for _, matchType := range MatchTypes {
		if mt == matchType {
			return true
		}
	}

	return false
}

// SubmatchCase is the enum to capture the two types of case changes to apply to submatches.
type SubmatchCase string

const (
	// Lower is the SubmatchCase for lower casing the submatch.
	Lower SubmatchCase = "lower"

	// Upper is the SubmatchCase for upper casing the submatch.
	Upper SubmatchCase = "upper"
)

var SubmatchCases = []SubmatchCase{Lower, Upper}

func (sc SubmatchCase) isValid() bool {
	for _, submatchCase := range SubmatchCases {
		if sc == submatchCase {
			return true
		}
	}

	return false
}
