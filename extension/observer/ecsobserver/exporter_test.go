// Copyright  OpenTelemetry Authors
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

package ecsobserver

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

func TestTaskExporter(t *testing.T) {
	exp := newTaskExporter(zap.NewExample(), "ecs-cluster-1")

	t.Run("invalid ip", func(t *testing.T) {
		_, err := exp.ExportTask(&Task{
			Task:       &ecs.Task{},
			Definition: &ecs.TaskDefinition{},
		})
		assert.Error(t, err)
		assert.IsType(t, &ErrPrivateIPNotFound{}, err)
	})

	awsVpcTask := &ecs.Task{
		TaskArn:           aws.String("arn:task:t2"),
		TaskDefinitionArn: aws.String("t2"),
		Attachments: []*ecs.Attachment{
			{
				Type: aws.String("ElasticNetworkInterface"),
				Details: []*ecs.KeyValuePair{
					{
						Name:  aws.String("privateIPv4Address"),
						Value: aws.String("172.168.1.1"),
					},
				},
			},
		},
	}
	awsVpcTaskDef := &ecs.TaskDefinition{
		ContainerDefinitions: []*ecs.ContainerDefinition{
			{
				PortMappings: []*ecs.PortMapping{
					{
						ContainerPort: aws.Int64(2112),
						HostPort:      aws.Int64(2113),
					},
				},
			},
		},
		NetworkMode: aws.String(ecs.NetworkModeAwsvpc),
	}
	t.Run("single target", func(t *testing.T) {
		task := &Task{
			Task:       awsVpcTask,
			Definition: awsVpcTaskDef,
			Matched: []MatchedContainer{
				{
					Targets: []MatchedTarget{
						{
							MatcherType: MatcherTypeDockerLabel,
							Port:        2112,
							Job:         "PROM_JOB_1",
						},
					},
				},
			},
		}

		targets, err := exp.ExportTask(task)
		require.NoError(t, err)
		assert.Len(t, targets, 1)
		assert.Equal(t, "172.168.1.1:2113", targets[0].Address)
	})

	t.Run("multiple target in one container", func(t *testing.T) {
		task := &Task{
			Task:       awsVpcTask,
			Definition: awsVpcTaskDef,
			Service:    &ecs.Service{ServiceName: aws.String("svc-1")},
			Matched: []MatchedContainer{
				{
					Targets: []MatchedTarget{
						{
							MatcherType: MatcherTypeDockerLabel,
							Port:        2112,
							Job:         "PROM_JOB_1",
						},
						{
							Port: 404, // invalid in the middle, but shouldn't stop the export
						},
						{
							MatcherType: MatcherTypeService,
							Port:        2112,
							Job:         "PROM_JOB_Service",
							MetricsPath: "/service_metrics",
						},
					},
				},
			},
		}

		targets, err := exp.ExportTask(task)
		require.Error(t, err)
		merr := multierr.Errors(err)
		require.Len(t, merr, 1)
		assert.IsType(t, &ErrMappedPortNotFound{}, merr[0])
		assert.Len(t, targets, 2)
	})

	t.Run("ec2", func(t *testing.T) {
		task := &Task{
			Task: &ecs.Task{
				TaskArn:           aws.String("arn:task:t2"),
				TaskDefinitionArn: aws.String("t2"),
				Containers: []*ecs.Container{
					{
						Name: aws.String("c1"),
						NetworkBindings: []*ecs.NetworkBinding{
							{
								ContainerPort: aws.Int64(2112),
								HostPort:      aws.Int64(2114),
							},
						},
					},
				},
			},
			Definition: &ecs.TaskDefinition{
				ContainerDefinitions: []*ecs.ContainerDefinition{
					{
						Name: aws.String("c1"),
						PortMappings: []*ecs.PortMapping{
							{
								ContainerPort: aws.Int64(2112),
							},
						},
					},
				},
				NetworkMode: aws.String(ecs.NetworkModeBridge),
			},
			EC2: &ec2.Instance{PrivateIpAddress: aws.String("172.168.1.2")},
			Matched: []MatchedContainer{
				{
					Targets: []MatchedTarget{
						{
							MatcherType: MatcherTypeDockerLabel,
							Port:        2112,
							Job:         "PROM_JOB_1",
						},
					},
				},
			},
		}

		targets, err := exp.ExportTask(task)
		require.NoError(t, err)
		assert.Len(t, targets, 1)
		assert.Equal(t, "172.168.1.2:2114", targets[0].Address)
	})

	validMatched := []MatchedContainer{
		{
			Targets: []MatchedTarget{
				{
					MatcherType: MatcherTypeDockerLabel,
					Port:        2112,
					Job:         "PROM_JOB_1",
				},
			},
		},
	}
	invalidMatched := []MatchedContainer{
		{
			Targets: []MatchedTarget{
				{
					MatcherType: MatcherTypeDockerLabel,
					Port:        404,
					Job:         "PROM_JOB_1",
				},
				{
					MatcherType: MatcherTypeDockerLabel,
					Port:        405,
					Job:         "PROM_JOB_1",
				},
			},
		},
	}
	validTask := &Task{
		Task:       awsVpcTask,
		Definition: awsVpcTaskDef,
		Matched:    validMatched,
	}
	invalidTask := &Task{
		Task:       awsVpcTask,
		Definition: awsVpcTaskDef,
		Matched:    invalidMatched,
	}
	t.Run("all valid tasks", func(t *testing.T) {
		// Just make sure the for loop is right, done care about the content, they are tested in previous cases
		targets, err := exp.ExportTasks([]*Task{validTask, validTask})
		assert.NoError(t, err)
		assert.Len(t, targets, 2)
	})

	t.Run("invalid tasks", func(t *testing.T) {
		targets, err := exp.ExportTasks([]*Task{validTask, invalidTask, validTask, invalidTask, invalidTask})
		require.Error(t, err)
		assert.Len(t, targets, 2)
		merr := multierr.Errors(err)
		// each invalid task has two invalid match, we have 3 invalid tasks
		// multierr flatten multierr when appending
		assert.Len(t, merr, 2*3)
	})
}
