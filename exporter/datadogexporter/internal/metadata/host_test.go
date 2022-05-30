// Copyright The OpenTelemetry Authors
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

package metadata

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/internal/utils/cache"
)

func TestHost(t *testing.T) {

	logger := zap.NewNop()

	// Start with a fresh cache, the following test would fail
	// if the cache key is already set.
	cache.Cache.Delete(cache.CanonicalHostnameKey)

	p, err := buildCurrentProvider(logger, "test-host")
	require.NoError(t, err)
	host, err := p.Hostname(context.Background())
	require.NoError(t, err)
	assert.Equal(t, host, "test-host")

	// config.Config.Hostname does not get stored in the cache
	p, err = buildCurrentProvider(logger, "test-host-2")
	require.NoError(t, err)
	host, err = p.Hostname(context.Background())
	require.NoError(t, err)
	assert.Equal(t, host, "test-host-2")

	// Disable EC2 Metadata service to prevent fetching hostname from there,
	// in case the test is running on an EC2 instance
	const awsEc2MetadataDisabled = "AWS_EC2_METADATA_DISABLED"
	curr := os.Getenv(awsEc2MetadataDisabled)
	err = os.Setenv(awsEc2MetadataDisabled, "true")
	require.NoError(t, err)
	defer os.Setenv(awsEc2MetadataDisabled, curr)

	p, err = buildCurrentProvider(logger, "")
	require.NoError(t, err)
	host, err = p.Hostname(context.Background())
	require.NoError(t, err)
	osHostname, err := os.Hostname()
	require.NoError(t, err)
	assert.Contains(t, host, osHostname)
}
