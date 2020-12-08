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

package ecs

import (
	"encoding/json"
	"log"
	"net/http"

	"go.uber.org/zap"
)

type metadataProvider interface {
	fetchTask(tmde string) (*TaskMetadata, error)
	fetchContainer(tmde string) (*Container, error)
}

type TaskMetadata struct {
	Cluster          string
	LaunchType       string // TODO: Change to enum when defined in otel collector convent
	TaskARN          string
	Family           string
	AvailabilityZone string
	Containers       []Container
}

type Container struct {
	DockerID     string `json:"DockerId"`
	ContainerARN string
	Type         string
	KnownStatus  string
	LogDriver    string
	LogOptions   LogData
}

type LogData struct {
	LogGroup string `json:"awslogs-group"`
	Region   string `json:"awslogs-region"`
	Stream   string `json:"awslogs-stream"`
}

type metadataClient struct {
	client *http.Client
	logger *zap.Logger
}

var _ metadataProvider = &metadataClient{}

// Retrieves the metadata for a task running on Amazon ECS
func (md *metadataClient) fetchTask(tmde string) (*TaskMetadata, error) {
	ret, err := fetch(tmde+"/task", md, true)
	if ret == nil {
		return nil, err
	}

	return ret.(*TaskMetadata), err
}

// Retrieves the metadata for the Amazon ECS Container the collector is running on
func (md *metadataClient) fetchContainer(tmde string) (*Container, error) {
	ret, err := fetch(tmde, md, false)
	if ret == nil {
		return nil, err
	}

	return ret.(*Container), err
}

func fetch(tmde string, md *metadataClient, task bool) (tmdeResp interface{}, err error) {
	// TODO(jbd): Use the zap logger instead of log.Print*.
	req, err := http.NewRequest(http.MethodGet, tmde, nil)

	if err != nil {
		log.Printf("Received error constructing request to ECS Task Metadata Endpoint: %v", err)
		return nil, err
	}

	resp, err := md.client.Do(req)
	if err != nil {
		log.Printf("Received error from ECS Task Metadata Endpoint: %v", err)
		return nil, err
	}

	if task {
		tmdeResp = &TaskMetadata{}
	} else {
		tmdeResp = &Container{}
	}

	err = json.NewDecoder(resp.Body).Decode(tmdeResp)
	defer resp.Body.Close()

	if err != nil {
		log.Printf("Encountered unexpected error reading response from ECS Task Metadata Endpoint: %v", err)
		return nil, err
	}

	return tmdeResp, nil
}
