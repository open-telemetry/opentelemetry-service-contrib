// Copyright  The OpenTelemetry Authors
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

package flinkmetricsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/flinkmetricsreceiver"

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/flinkmetricsreceiver/internal/models"
)

// The API endpoints required to collect metrics.
const (
	// jobmanagerMetricEndpoint gets jobmanager metrics.
	jobmanagerMetricEndpoint = "/jobmanager/metrics"
	// taskmanagersEndpoint gets taskmanager IDs.
	taskmanagersEndpoint = "/taskmanagers"
	// taskmanagersMetricEndpoint gets taskmanager using a taskmanager ID.
	taskmanagersMetricEndpoint = "/taskmanagers/%s/metrics"
	// jobsEndpoint gets job IDs.
	jobsEndpoint = "/jobs"
	// jobsOverviewEndpoint gets job IDs with associated Job names.
	jobsOverviewEndpoint = "/jobs/overview"
	// jobsWithIDEndpoint gets vertex IDs using a job ID.
	jobsWithIDEndpoint = "/jobs/%s"
	// jobsMetricEndpoint gets job metrics using a job ID.
	jobsMetricEndpoint = "/jobs/%s/metrics"
	// verticesEndpoint gets subtask index's using a job and vertex ID.
	verticesEndpoint = "/jobs/%s/vertices/%s"
	// subtaskMetricEndpoint gets subtask metrics using a job ID, vertex ID and subtask index.
	subtaskMetricEndpoint = "/jobs/%s/vertices/%s/subtasks/%v/metrics"
)

type client interface {
	GetJobmanagerMetrics(ctx context.Context) (*models.JobmanagerMetrics, error)
	GetTaskmanagersMetrics(ctx context.Context) ([]*models.TaskmanagerMetrics, error)
	GetJobsMetrics(ctx context.Context) ([]*models.JobMetrics, error)
	GetSubtasksMetrics(ctx context.Context) ([]*models.SubtaskMetrics, error)
}

type flinkClient struct {
	client       *http.Client
	hostEndpoint string
	hostName     string
	logger       *zap.Logger
}

func newClient(cfg *Config, host component.Host, settings component.TelemetrySettings, logger *zap.Logger) (client, error) {
	httpClient, err := cfg.ToClient(host.GetExtensions(), settings)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP Client: %w", err)
	}

	hostName, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	return &flinkClient{
		client:       httpClient,
		hostName:     hostName,
		hostEndpoint: cfg.Endpoint,
		logger:       logger,
	}, nil
}

func (c *flinkClient) get(ctx context.Context, path string) ([]byte, error) {
	// Construct endpoint and create request
	url := c.hostEndpoint + path
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create get request for path %s: %w", path, err)
	}

	// Make request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make http request: %w", err)
	}

	// Defer body close
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			c.logger.Warn("failed to close response body", zap.Error(closeErr))
		}
	}()

	// Check for OK status code
	if resp.StatusCode != http.StatusOK {
		c.logger.Debug("flink API non-200", zap.Error(err), zap.Int("status_code", resp.StatusCode))

		// Attempt to extract the error payload
		payloadData, err := io.ReadAll(resp.Body)
		if err != nil {
			c.logger.Debug("failed to read payload error message", zap.Error(err))
		} else {
			c.logger.Debug("flink API Error", zap.ByteString("api_error", payloadData))
		}

		return nil, fmt.Errorf("non 200 code returned %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

// getMetrics makes a request to a metric endpoint to get the metric names, the another request building a query to get the metric values.
func (c *flinkClient) getMetrics(ctx context.Context, path string) (*models.MetricsResponse, error) {
	// Get the metric names
	var metrics *models.MetricsResponse
	body, err := c.get(ctx, path)
	if err != nil {
		c.logger.Debug("Failed to retrieve metric names", zap.Error(err))
		return nil, err
	}

	// Populates the metric names
	err = json.Unmarshal(body, &metrics)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	// Construct a get query parameter using comma-separated list of string values to select specific metrics
	query := []string{}
	for _, metricName := range *metrics {
		query = append(query, metricName.ID)
	}
	metricsPath := path + "?get=" + strings.Join(query, ",")

	// get the metric values using the query
	body, err = c.get(ctx, metricsPath)
	if err != nil {
		c.logger.Debug("Failed to retrieve metric values", zap.Error(err))
		return nil, err
	}

	// Populates metric values
	err = json.Unmarshal(body, &metrics)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return metrics, nil
}

// GetJobManagerMetrics gets the jobmanager metrics.
func (c *flinkClient) GetJobmanagerMetrics(ctx context.Context) (*models.JobmanagerMetrics, error) {
	// Get the metric names and values for jobmanager
	metrics, err := c.getMetrics(ctx, jobmanagerMetricEndpoint)
	if err != nil {
		return nil, err
	}

	// Add a hostname used to identify between multiple jobmanager instances
	return &models.JobmanagerMetrics{
		Host:    c.hostName,
		Metrics: *metrics,
	}, nil
}

type taskmanagerResult struct {
	taskmanagerInstance models.TaskmanagerMetrics
	err                 error
}

// GetTaskmanagersMetrics gets the Taskmanager metrics for each taskmanager.
func (c *flinkClient) GetTaskmanagersMetrics(ctx context.Context) ([]*models.TaskmanagerMetrics, error) {
	// Get the taskmanager id list
	var taskManagerIDs *models.TaskmanagerIDsResponse
	body, err := c.get(ctx, taskmanagersEndpoint)
	if err != nil {
		c.logger.Debug("Failed to retrieve taskmanager IDs", zap.Error(err))
		return nil, err
	}

	// Populates taskmanager id names
	err = json.Unmarshal(body, &taskManagerIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	// TaskManagerInstances stores all metric data for each taskmanager id
	var taskManagerInstances []*models.TaskmanagerMetrics

	// Get taskmanager metrics for each task manager id
	for _, taskmanager := range taskManagerIDs.Taskmanagers {
		metrics, err := c.getMetrics(ctx, fmt.Sprintf(taskmanagersMetricEndpoint, taskmanager.ID))
		if err != nil {
			return nil, err
		}
		// Taskmanger ID is in the form of host:id
		host := strings.Split(taskmanager.ID, ":")
		taskManagerInstance := models.TaskmanagerMetrics{
			TaskmanagerID: taskmanager.ID,
			Host:          host[0],
			Metrics:       *metrics,
		}
		taskManagerInstances = append(taskManagerInstances, &taskManagerInstance)
	}

	return taskManagerInstances, nil
}

// GetJobsMetrics gets the job metrics for each job.
func (c *flinkClient) GetJobsMetrics(ctx context.Context) ([]*models.JobMetrics, error) {
	// Get the job id and name list
	var jobIDs *models.JobOverviewResponse
	body, err := c.get(ctx, jobsOverviewEndpoint)
	if err != nil {
		c.logger.Debug("Failed to retrieve job IDs", zap.Error(err))
		return nil, err
	}

	// Populates job id and names
	err = json.Unmarshal(body, &jobIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	// JobInstances stores all metric data for each job id
	var jobInstances []*models.JobMetrics

	// Get the job metrics for each job id
	for _, job := range jobIDs.Jobs {
		metrics, err := c.getMetrics(ctx, fmt.Sprintf(jobsMetricEndpoint, job.Jid))
		if err != nil {
			return nil, err
		}
		jobInstance := models.JobMetrics{
			Host:    c.hostName,
			JobName: job.Name,
			Metrics: *metrics,
		}
		jobInstances = append(jobInstances, &jobInstance)
	}

	return jobInstances, nil
}

// GetSubtasksMetrics gets subtask metrics for each job id, vertex id and subtask index.
func (c *flinkClient) GetSubtasksMetrics(ctx context.Context) ([]*models.SubtaskMetrics, error) {
	var subtaskMetricsStorage []*models.SubtaskMetrics
	// Get the job id's
	var jobsResponse *models.JobsResponse
	body, err := c.get(ctx, jobsEndpoint)
	if err != nil {
		c.logger.Debug("Failed to retrieve job IDs", zap.Error(err))
		return nil, err
	}

	// Populates the job id
	err = json.Unmarshal(body, &jobsResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	// Get vertices for each job
	for _, job := range jobsResponse.Jobs {
		var jobsWithIDResponse *models.JobsWithIDResponse
		query := fmt.Sprintf(jobsWithIDEndpoint, job.ID)
		body, err = c.get(ctx, query)
		if err != nil {
			c.logger.Debug("Failed to retrieve job with ID", zap.Error(err))
			return nil, err
		}

		// Populates the job response with vertices info
		err = json.Unmarshal(body, &jobsWithIDResponse)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
		}
		// gets subtask info for each vertex id
		for _, vertex := range jobsWithIDResponse.Vertices {
			var vertexResponse *models.VerticesResponse
			query := fmt.Sprintf(verticesEndpoint, job.ID, vertex.ID)
			body, err = c.get(ctx, query)
			if err != nil {
				c.logger.Debug("Failed to retrieve vertex with ID", zap.Error(err))
				return nil, err
			}

			// Populates the vertex response with subtask info
			err = json.Unmarshal(body, &vertexResponse)
			if err != nil {
				return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
			}

			// gets subtask metrics for each vertex id
			for _, subtask := range vertexResponse.Subtasks {
				query := fmt.Sprintf(subtaskMetricEndpoint, job.ID, vertex.ID, subtask.Subtask)
				subtaskMetrics, err := c.getMetrics(ctx, query)
				if err != nil {
					c.logger.Debug("Failed to retrieve subtasks metrics", zap.Error(err))
					return nil, err
				}

				// stores subtask info with additional attribute values to uniquely identify metrics
				subtaskMetricsStorage = append(subtaskMetricsStorage,
					&models.SubtaskMetrics{
						Host:          subtask.Host,
						TaskmanagerID: subtask.TaskmanagerID,
						JobName:       jobsWithIDResponse.Name,
						TaskName:      vertex.Name,
						SubtaskIndex:  fmt.Sprintf("%v", subtask.Subtask),
						Metrics:       *subtaskMetrics,
					})
			}
		}
	}
	return subtaskMetricsStorage, nil
}
