// Copyright The OpenTelemetry Authors
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

package cwmetricstream // import "github.com/asserts/opentelemetry-collector-contrib/receiver/awsfirehosereceiver/internal/unmarshaler/cwmetricstream"

// The cWMetric is the format for the CloudWatch metric stream records.
//
// More details can be found at:
// https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-formats-json.html
type cWMetric struct {
	// MetricStreamName is the name of the CloudWatch metric stream.
	MetricStreamName string `json:"metric_stream_name"`
	// AccountID is the AWS account ID associated with the metric.
	AccountID string `json:"account_id"`
	// Region is the AWS region for the metric.
	Region string `json:"region"`
	// Namespace is the CloudWatch namespace the metric is in.
	Namespace string `json:"namespace"`
	// MetricName is the name of the metric.
	MetricName string `json:"metric_name"`
	// Dimensions is a map of name/value pairs that help to
	// differentiate a metric.
	Dimensions map[string]string `json:"dimensions"`
	// Timestamp is the milliseconds since epoch for
	// the metric.
	Timestamp int64 `json:"timestamp"`
	// Value is the cWMetricValue, which has the min, max,
	// sum, and count.
	Value *cWMetricValue `json:"value"`
	// Unit is the unit for the metric.
	//
	// More details can be found at:
	// https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricDatum.html
	Unit string `json:"unit"`
}

// The cWMetricValue is the actual values of the CloudWatch metric.
type cWMetricValue struct {
	// Max is the highest value observed.
	Max float64 `json:"max"`
	// Min is the lowest value observed.
	Min float64 `json:"min"`
	// Sum is the sum of data points collected.
	Sum float64 `json:"sum"`
	// Count is the number of data points.
	Count float64 `json:"count"`
}
