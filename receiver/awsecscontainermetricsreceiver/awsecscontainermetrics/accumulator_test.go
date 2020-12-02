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
	"go.uber.org/zap"
)

var (
	v         = uint64(1)
	f         = 1.0
	floatZero = float64(0)
	logger    = zap.NewNop()

	memStats = map[string]uint64{"cache": v}

	mem = MemoryStats{
		Usage:          &v,
		MaxUsage:       &v,
		Limit:          &v,
		MemoryReserved: &v,
		MemoryUtilized: &v,
		Stats:          memStats,
	}

	disk = DiskStats{
		IoServiceBytesRecursives: []IoServiceBytesRecursive{
			{Op: "Read", Value: &v},
			{Op: "Write", Value: &v},
			{Op: "Total", Value: &v},
		},
	}

	networkStat = NetworkStats{
		RxBytes:   &v,
		RxPackets: &v,
		RxErrors:  &v,
		RxDropped: &v,
		TxBytes:   &v,
		TxPackets: &v,
		TxErrors:  &v,
		TxDropped: &v,
	}
	net = map[string]NetworkStats{"eth0": networkStat}

	netRate = NetworkRateStats{
		RxBytesPerSecond: &f,
		TxBytesPerSecond: &f,
	}

	percpu   = []*uint64{&v, &v}
	cpuUsage = CPUUsage{
		TotalUsage:        &v,
		UsageInKernelmode: &v,
		UsageInUserMode:   &v,
		PerCPUUsage:       percpu,
	}

	cpuStats = CPUStats{
		CPUUsage:       &cpuUsage,
		OnlineCpus:     &v,
		SystemCPUUsage: &v,
		CPUUtilized:    &v,
		CPUReserved:    &v,
	}
	containerStats = ContainerStats{
		Name:        "test",
		ID:          "001",
		Memory:      &mem,
		Disk:        &disk,
		Network:     net,
		NetworkRate: &netRate,
		CPU:         &cpuStats,
	}

	tm = TaskMetadata{
		Cluster:  "cluster-1",
		TaskARN:  "arn:aws:some-value/001",
		Family:   "task-def-family-1",
		Revision: "task-def-version",
		Containers: []ContainerMetadata{
			{ContainerName: "container-1", DockerID: "001", DockerName: "docker-container-1", Limits: Limit{CPU: &f, Memory: &v}},
		},
		Limits: Limit{CPU: &f, Memory: &v},
	}

	cstats = map[string]*ContainerStats{"001": &containerStats}
	acc    = metricDataAccumulator{
		mds: nil,
	}
)

func TestGetMetricsDataAllValid(t *testing.T) {
	acc.getMetricsData(cstats, tm, logger)
	require.Less(t, 0, len(acc.mds))
}

func TestGetMetricsDataMissingContainerStats(t *testing.T) {
	tm.Containers = []ContainerMetadata{
		{ContainerName: "container-1", DockerID: "001-Missing", DockerName: "docker-container-1", Limits: Limit{CPU: &f, Memory: &v}},
	}
	acc.getMetricsData(cstats, tm, logger)
	require.Less(t, 0, len(acc.mds))
}

func TestGetMetricsDataNilContainerStats(t *testing.T) {
	cstats = map[string]*ContainerStats{"001": nil}

	acc.getMetricsData(cstats, tm, logger)
	require.Less(t, 0, len(acc.mds))
}

func TestGetMetricsDataMissingContainerLimit(t *testing.T) {
	tm.Containers = []ContainerMetadata{
		{ContainerName: "container-1", DockerID: "001", DockerName: "docker-container-1"},
	}

	acc.getMetricsData(cstats, tm, logger)
	require.Less(t, 0, len(acc.mds))
}

func TestGetMetricsDataContainerLimitCpuNil(t *testing.T) {
	tm.Containers = []ContainerMetadata{
		{ContainerName: "container-1", DockerID: "001", DockerName: "docker-container-1", Limits: Limit{CPU: nil, Memory: &v}},
	}

	acc.getMetricsData(cstats, tm, logger)
	require.Less(t, 0, len(acc.mds))
}

func TestGetMetricsDataContainerLimitMemoryNil(t *testing.T) {
	tm.Containers = []ContainerMetadata{
		{ContainerName: "container-1", DockerID: "001", DockerName: "docker-container-1", Limits: Limit{CPU: &f, Memory: nil}},
	}

	acc.getMetricsData(cstats, tm, logger)
	require.Less(t, 0, len(acc.mds))
}

func TestGetMetricsDataMissingTaskLimit(t *testing.T) {
	tm = TaskMetadata{
		Cluster:  "cluster-1",
		TaskARN:  "arn:aws:some-value/001",
		Family:   "task-def-family-1",
		Revision: "task-def-version",
		Containers: []ContainerMetadata{
			{ContainerName: "container-1", DockerID: "001", DockerName: "docker-container-1", Limits: Limit{CPU: &f, Memory: &v}},
		},
	}

	acc.getMetricsData(cstats, tm, logger)
	require.Less(t, 0, len(acc.mds))
}

func TestGetMetricsDataTaskLimitCpuNil(t *testing.T) {
	tm = TaskMetadata{
		Cluster:  "cluster-1",
		TaskARN:  "arn:aws:some-value/001",
		Family:   "task-def-family-1",
		Revision: "task-def-version",
		Containers: []ContainerMetadata{
			{ContainerName: "container-1", DockerID: "001", DockerName: "docker-container-1", Limits: Limit{CPU: &f, Memory: &v}},
		},
		Limits: Limit{CPU: nil, Memory: &v},
	}

	acc.getMetricsData(cstats, tm, logger)
	require.Less(t, 0, len(acc.mds))
}

func TestGetMetricsDataTaskLimitMemoryNil(t *testing.T) {
	tm = TaskMetadata{
		Cluster:  "cluster-1",
		TaskARN:  "arn:aws:some-value/001",
		Family:   "task-def-family-1",
		Revision: "task-def-version",
		Containers: []ContainerMetadata{
			{ContainerName: "container-1", DockerID: "001", DockerName: "docker-container-1", Limits: Limit{CPU: &f, Memory: &v}},
		},
		Limits: Limit{CPU: &f, Memory: nil},
	}

	acc.getMetricsData(cstats, tm, logger)
	require.Less(t, 0, len(acc.mds))
}

func TestGetMetricsDataCpuReservedZero(t *testing.T) {
	tm = TaskMetadata{
		Cluster:  "cluster-1",
		TaskARN:  "arn:aws:some-value/001",
		Family:   "task-def-family-1",
		Revision: "task-def-version",
		Containers: []ContainerMetadata{
			{ContainerName: "container-1", DockerID: "001", DockerName: "docker-container-1", Limits: Limit{CPU: &floatZero, Memory: nil}},
		},
		Limits: Limit{CPU: &floatZero, Memory: &v},
	}

	acc.getMetricsData(cstats, tm, logger)
	require.Less(t, 0, len(acc.mds))
}
