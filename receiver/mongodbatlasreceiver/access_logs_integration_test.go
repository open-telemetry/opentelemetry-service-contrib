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

//go:build integration

package mongodbatlasreceiver

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas/mongodbatlas"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/receiver/receivertest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/golden"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest/plogtest"
)

func TestAccessLogsIntegration(t *testing.T) {
	mockClient := mockAccessLogsClient{}

	payloadFile, err := os.ReadFile(filepath.Join("testdata", "accesslogs", "sample-payloads", "sample-access-logs.json"))
	require.NoError(t, err)

	var accessLogs []*mongodbatlas.AccessLogs
	err = json.Unmarshal(payloadFile, &accessLogs)
	require.NoError(t, err)

	mockClient.On("GetProject", mock.Anything, testProjectName).Return(&mongodbatlas.Project{
		ID:    testProjectID,
		Name:  testProjectName,
		OrgID: testOrgID,
	}, nil)
	mockClient.On("GetClusters", mock.Anything, testProjectID).Return(
		[]mongodbatlas.Cluster{
			{
				ID:      testClusterID,
				GroupID: testProjectID,
				Name:    testClusterName,
			},
		},
		nil)
	mockClient.On("GetAccessLogs", mock.Anything, testProjectID, testClusterID, mock.Anything).Return(accessLogs, nil)

	sink := &consumertest.LogsSink{}
	fact := NewFactory()

	recv, err := fact.CreateLogsReceiver(
		context.Background(),
		receivertest.NewNopCreateSettings(),
		&Config{
			AccessLogs: &AccessLogsConfig{
				Projects: []ProjectConfig{
					{
						Name: testProjectName,
					},
				},
				PollInterval: 1 * time.Second,
			},
		},
		sink,
	)
	require.NoError(t, err)

	rcvr, ok := recv.(*combinedLogsReceiver)
	require.True(t, ok)
	rcvr.accessLogs.client = &mockClient

	err = recv.Start(context.Background(), componenttest.NewNopHost())
	require.NoError(t, err)

	require.Eventually(t, func() bool {
		return sink.LogRecordCount() > 0
	}, 5*time.Second, 10*time.Millisecond)

	err = recv.Shutdown(context.Background())
	require.NoError(t, err)

	logs := sink.AllLogs()[0]
	expectedLogs, err := golden.ReadLogs(filepath.Join("testdata", "accesslogs", "golden", "retrieved-logs.yaml"))
	require.NoError(t, err)
	require.NoError(t, plogtest.CompareLogs(expectedLogs, logs, plogtest.IgnoreObservedTimestamp()))
}
