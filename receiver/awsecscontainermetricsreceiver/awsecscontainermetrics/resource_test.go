// Copyright 2020, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package awsecscontainermetrics

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/consumer/pdata"
	"go.opentelemetry.io/collector/translator/conventions"
)

func TestContainerResource(t *testing.T) {

	cm := ContainerMetadata{
		ContainerName: "container-1",
		DockerID:      "001",
		DockerName:    "docker-container-1",
		Image:         "nginx:v1.0",
		ImageID:       "sha256:8cf1bfb43ff5d9b05af9b6b63983440f137",
		CreatedAt:     "2020-07-30T22:12:29.837074927Z",
		StartedAt:     "2020-07-30T22:12:31.153459485Z",
		FinishedAt:    "2020-07-31T22:12:29.837074927Z",
		KnownStatus:   "RUNNING",
	}

	r := containerResource(cm)
	require.NotNil(t, r)

	attrMap := r.Attributes()
	//require.EqualValues(t, 11, attrMap.Len())
	expected := map[string]string{
		conventions.AttributeContainerName:  "container-1",
		conventions.AttributeContainerID:    "001",
		AttributeECSDockerName:              "docker-container-1",
		conventions.AttributeContainerImage: "nginx:v1.0",
		AttributeContainerImageID:           "sha256:8cf1bfb43ff5d9b05af9b6b63983440f137",
		conventions.AttributeContainerTag:    "v1.0",
		AttributeContainerCreatedAt:         "2020-07-30T22:12:29.837074927Z",
		AttributeContainerStartedAt:         "2020-07-30T22:12:31.153459485Z",
		AttributeContainerFinishedAt:        "2020-07-31T22:12:29.837074927Z",
		AttributeContainerKnownStatus:       "RUNNING",
	}


	verifyAttributeMap(t, expected, attrMap)
}

func TestTaskResource(t *testing.T) {
	tm := TaskMetadata{
		Cluster:          "cluster-1",
		TaskARN:          "arn:aws:ecs:us-west-2:111122223333:task/default/158d1c8083dd49d6b527399fd6414f5c",
		Family:           "task-def-family-1",
		Revision:         "v1.2",
		AvailabilityZone: "us-west-2d",
		PullStartedAt:    "2020-10-02T00:43:06.202617438Z",
		PullStoppedAt:    "2020-10-02T00:43:06.31288465Z",
		KnownStatus:      "RUNNING",
		LaunchType:       "EC2",
	}
	r := taskResource(tm)
	require.NotNil(t, r)

	attrMap := r.Attributes()
	require.EqualValues(t, 13, attrMap.Len())
	expected := map[string]string{
		AttributeECSCluster:               "cluster-1",
		AttributeECSTaskARN:               "arn:aws:ecs:us-west-2:111122223333:task/default/158d1c8083dd49d6b527399fd6414f5c",
		AttributeECSTaskID:                "158d1c8083dd49d6b527399fd6414f5c",
		AttributeECSTaskFamily:            "task-def-family-1",
		AttributeECSTaskRevision:          "v1.2",
		conventions.AttributeCloudZone:    "us-west-2d",
		AttributeECSTaskPullStartedAt:     "2020-10-02T00:43:06.202617438Z",
		AttributeECSTaskPullStoppedAt:     "2020-10-02T00:43:06.31288465Z",
		AttributeECSTaskKnownStatus:       "RUNNING",
		AttributeECSTaskLaunchType:        "EC2",
		conventions.AttributeCloudRegion:  "us-west-2",
		conventions.AttributeCloudAccount: "111122223333",
	}

	verifyAttributeMap(t, expected, attrMap)
}

func TestTaskResourceWithClusterARN(t *testing.T) {
	tm := TaskMetadata{
		Cluster:          "arn:aws:ecs:us-west-2:803860917211:cluster/main-cluster",
		TaskARN:          "arn:aws:ecs:us-west-2:803860917211:cluster/main-cluster/c8083dd49d6b527399fd6414",
		Family:           "task-def-family-1",
		Revision:         "v1.2",
		AvailabilityZone: "us-west-2d",
		PullStartedAt:    "2020-10-02T00:43:06.202617438Z",
		PullStoppedAt:    "2020-10-02T00:43:06.31288465Z",
		KnownStatus:      "RUNNING",
		LaunchType:       "EC2",
	}
	r := taskResource(tm)
	require.NotNil(t, r)

	attrMap := r.Attributes()
	require.EqualValues(t, 13, attrMap.Len())

	expected := map[string]string{
		AttributeECSCluster:               "main-cluster",
		AttributeECSTaskARN:               "arn:aws:ecs:us-west-2:803860917211:cluster/main-cluster/c8083dd49d6b527399fd6414",
		AttributeECSTaskID:                "c8083dd49d6b527399fd6414",
		AttributeECSTaskFamily:            "task-def-family-1",
		AttributeECSTaskRevision:          "v1.2",
		conventions.AttributeCloudZone:    "us-west-2d",
		AttributeECSTaskPullStartedAt:     "2020-10-02T00:43:06.202617438Z",
		AttributeECSTaskPullStoppedAt:     "2020-10-02T00:43:06.31288465Z",
		AttributeECSTaskKnownStatus:       "RUNNING",
		AttributeECSTaskLaunchType:        "EC2",
		conventions.AttributeCloudRegion:  "us-west-2",
		conventions.AttributeCloudAccount: "803860917211",
	}

	verifyAttributeMap(t, expected, attrMap)
}

func verifyAttributeMap(t *testing.T, expected map[string]string, found pdata.AttributeMap) {
	for key, val := range expected {
		attributeVal, found := found.Get(key)
		require.EqualValues(t, true, found)

		require.EqualValues(t, val, attributeVal.StringVal())
	}
}

func TestGetResourceFromARN(t *testing.T) {
	region, accountID, taskID := getResourceFromARN("arn:aws:ecs:us-west-2:803860917211:task/test200/d22aaa11bf0e4ab19c2c940a1cbabbee")
	require.EqualValues(t, "us-west-2", region)
	require.EqualValues(t, "803860917211", accountID)
	require.EqualValues(t, "d22aaa11bf0e4ab19c2c940a1cbabbee", taskID)
	region, accountID, taskID = getResourceFromARN("")
	require.LessOrEqual(t, 0, len(region))
	require.LessOrEqual(t, 0, len(accountID))
	require.LessOrEqual(t, 0, len(taskID))
	region, accountID, taskID = getResourceFromARN("notarn:aws:ecs:us-west-2:803860917211:task/test2")
	require.LessOrEqual(t, 0, len(region))
	require.LessOrEqual(t, 0, len(accountID))
	require.LessOrEqual(t, 0, len(taskID))
}

func TestGetVersionFromImage(t *testing.T) {
	version := getVersionFromIamge("docker-Im.age:v1.0")
	require.EqualValues(t, "v1.0", version)
	version = getVersionFromIamge("dockerIm-agev1.0")
	require.EqualValues(t, "latest", version)
	version = getVersionFromIamge("")
	require.LessOrEqual(t, 0, len(version))
}

func TestGetNameFromCluster(t *testing.T) {
	clusterName := getNameFromCluster("arn:aws:ecs:region:012345678910:cluster/test")
	require.EqualValues(t, "test", clusterName)

	clusterName = getNameFromCluster("not-arn:aws:something/001")
	require.EqualValues(t, "not-arn:aws:something/001", clusterName)

	clusterName = getNameFromCluster("")
	require.LessOrEqual(t, 0, len(clusterName))

}
