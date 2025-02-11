// Copyright The OpenTelemetry Authors
// Copyright (c) 2018 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package filesource

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/jaegertracing/jaeger-idl/proto-gen/api_v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

const snapshotLocation = "./fixtures/"

// Snapshots can be regenerated via:
//
//	REGENERATE_SNAPSHOTS=true go test -v ./plugin/sampling/strategyprovider/static/provider_test.go
var regenerateSnapshots = os.Getenv("REGENERATE_SNAPSHOTS") == "true"

// strategiesJSON returns the strategy with
// a given probability.
func strategiesJSON(probability float32) string {
	strategy := fmt.Sprintf(`
		{
			"default_strategy": {
				"type": "probabilistic",
				"param": 0.5
			},
			"service_strategies": [
				{
					"service": "foo",
					"type": "probabilistic",
					"param": %.1f
				},
				{
					"service": "bar",
					"type": "ratelimiting",
					"param": 5
				}
			]
		}
		`,
		probability,
	)
	return strategy
}

func deepCopy(s *api_v2.SamplingStrategyResponse) (*api_v2.SamplingStrategyResponse, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)
	err := enc.Encode(*s)
	if err != nil {
		return nil, err
	}
	var copyValue api_v2.SamplingStrategyResponse
	err = dec.Decode(&copyValue)
	return &copyValue, err
}

// Returns strategies in JSON format. Used for testing
// URL option for sampling strategies.
func mockStrategyServer(t *testing.T) (*httptest.Server, *atomic.Pointer[string]) {
	var strategy atomic.Pointer[string]
	value := strategiesJSON(0.8)
	strategy.Store(&value)
	f := func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/bad-content":
			_, err := w.Write([]byte("bad-content"))
			assert.NoError(t, err)
			return

		case "/bad-status":
			w.WriteHeader(http.StatusNotFound)
			return

		case "/service-unavailable":
			w.WriteHeader(http.StatusServiceUnavailable)
			return

		default:
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			_, err := w.Write([]byte(*strategy.Load()))
			assert.NoError(t, err)
		}
	}
	mockserver := httptest.NewServer(http.HandlerFunc(f))
	t.Cleanup(func() {
		mockserver.Close()
	})
	return mockserver, &strategy
}

func TestStrategyStoreWithFile(t *testing.T) {
	_, err := NewFileSource(Options{StrategiesFile: "fileNotFound.json"}, zap.NewNop())
	require.ErrorContains(t, err, "failed to read strategies file fileNotFound.json")

	_, err = NewFileSource(Options{StrategiesFile: "fixtures/bad_strategies.json"}, zap.NewNop())
	require.EqualError(t, err,
		"failed to unmarshal strategies: json: cannot unmarshal string into Go value of type filesource.strategies")

	// Test default strategy
	zapCore, logs := observer.New(zap.InfoLevel)
	logger := zap.New(zapCore)
	provider, err := NewFileSource(Options{}, logger)
	require.NoError(t, err)
	message := logs.FilterMessage("No sampling strategies source provided, using defaults")
	assert.Equal(t, 1, message.Len(), "Expected No sampling strategies provided log message")
	s, err := provider.GetSamplingStrategy(context.Background(), "foo")
	require.NoError(t, err)
	assert.EqualValues(t, makeResponse(api_v2.SamplingStrategyType_PROBABILISTIC, 0.001), *s)

	// Test reading strategies from a file
	provider, err = NewFileSource(Options{StrategiesFile: "fixtures/strategies.json"}, logger)
	require.NoError(t, err)
	s, err = provider.GetSamplingStrategy(context.Background(), "foo")
	require.NoError(t, err)
	assert.EqualValues(t, makeResponse(api_v2.SamplingStrategyType_PROBABILISTIC, 0.8), *s)

	s, err = provider.GetSamplingStrategy(context.Background(), "bar")
	require.NoError(t, err)
	assert.EqualValues(t, makeResponse(api_v2.SamplingStrategyType_RATE_LIMITING, 5), *s)

	s, err = provider.GetSamplingStrategy(context.Background(), "default")
	require.NoError(t, err)
	assert.EqualValues(t, makeResponse(api_v2.SamplingStrategyType_PROBABILISTIC, 0.5), *s)
}

func TestStrategyStoreWithURL(t *testing.T) {
	// Test default strategy when URL is temporarily unavailable.
	zapCore, logs := observer.New(zap.InfoLevel)
	logger := zap.New(zapCore)
	mockServer, _ := mockStrategyServer(t)
	provider, err := NewFileSource(Options{StrategiesFile: mockServer.URL + "/service-unavailable"}, logger)
	require.NoError(t, err)
	message := logs.FilterMessage("No sampling strategies found or URL is unavailable, using defaults")
	assert.Equal(t, 1, message.Len(), "Expected No sampling strategies found log message.")
	s, err := provider.GetSamplingStrategy(context.Background(), "foo")
	require.NoError(t, err)
	assert.EqualValues(t, makeResponse(api_v2.SamplingStrategyType_PROBABILISTIC, 0.001), *s)

	// Test downloading strategies from a URL.
	provider, err = NewFileSource(Options{StrategiesFile: mockServer.URL}, logger)
	require.NoError(t, err)

	s, err = provider.GetSamplingStrategy(context.Background(), "foo")
	require.NoError(t, err)
	assert.EqualValues(t, makeResponse(api_v2.SamplingStrategyType_PROBABILISTIC, 0.8), *s)

	s, err = provider.GetSamplingStrategy(context.Background(), "bar")
	require.NoError(t, err)
	assert.EqualValues(t, makeResponse(api_v2.SamplingStrategyType_RATE_LIMITING, 5), *s)
}

func TestPerOperationSamplingStrategies(t *testing.T) {
	tests := []struct {
		options Options
	}{
		{Options{StrategiesFile: "fixtures/operation_strategies.json"}},
		{Options{
			StrategiesFile:             "fixtures/operation_strategies.json",
			IncludeDefaultOpStrategies: true,
		}},
	}

	for _, tc := range tests {
		zapCore, logs := observer.New(zap.InfoLevel)
		logger := zap.New(zapCore)
		provider, err := NewFileSource(tc.options, logger)
		message := logs.FilterMessage("Operation strategies only supports probabilistic sampling at the moment," +
			"'op2' defaulting to probabilistic sampling with probability 0.800000")
		assert.Equal(t, 1, message.Len(), "Expected Operation strategies only supports probabilistic sampling, op2 change to probability 0.8 log message")
		message = logs.FilterMessage("Operation strategies only supports probabilistic sampling at the moment," +
			"'op4' defaulting to probabilistic sampling with probability 0.001000")
		assert.Equal(t, 1, message.Len(), "Expected Operation strategies only supports probabilistic sampling, op2 change to probability 0.8 log message")
		require.NoError(t, err)

		expected := makeResponse(api_v2.SamplingStrategyType_PROBABILISTIC, 0.8)

		s, err := provider.GetSamplingStrategy(context.Background(), "foo")
		require.NoError(t, err)
		assert.Equal(t, api_v2.SamplingStrategyType_PROBABILISTIC, s.StrategyType)
		assert.Equal(t, *expected.ProbabilisticSampling, *s.ProbabilisticSampling)

		require.NotNil(t, s.OperationSampling)
		opSampling := s.OperationSampling
		assert.InDelta(t, 0.8, opSampling.DefaultSamplingProbability, 0.01)
		require.Len(t, opSampling.PerOperationStrategies, 4)

		assert.Equal(t, "op6", opSampling.PerOperationStrategies[0].Operation)
		assert.InDelta(t, 0.5, opSampling.PerOperationStrategies[0].ProbabilisticSampling.SamplingRate, 0.01)
		assert.Equal(t, "op1", opSampling.PerOperationStrategies[1].Operation)
		assert.InDelta(t, 0.2, opSampling.PerOperationStrategies[1].ProbabilisticSampling.SamplingRate, 0.01)
		assert.Equal(t, "op0", opSampling.PerOperationStrategies[2].Operation)
		assert.InDelta(t, 0.2, opSampling.PerOperationStrategies[2].ProbabilisticSampling.SamplingRate, 0.01)
		assert.Equal(t, "op7", opSampling.PerOperationStrategies[3].Operation)
		assert.InDelta(t, 1.0, opSampling.PerOperationStrategies[3].ProbabilisticSampling.SamplingRate, 0.01)

		expected = makeResponse(api_v2.SamplingStrategyType_RATE_LIMITING, 5)

		s, err = provider.GetSamplingStrategy(context.Background(), "bar")
		require.NoError(t, err)
		assert.Equal(t, api_v2.SamplingStrategyType_RATE_LIMITING, s.StrategyType)
		assert.Equal(t, *expected.RateLimitingSampling, *s.RateLimitingSampling)

		require.NotNil(t, s.OperationSampling)
		opSampling = s.OperationSampling
		assert.InDelta(t, 0.001, opSampling.DefaultSamplingProbability, 1e-4)
		require.Len(t, opSampling.PerOperationStrategies, 5)
		assert.Equal(t, "op3", opSampling.PerOperationStrategies[0].Operation)
		assert.InDelta(t, 0.3, opSampling.PerOperationStrategies[0].ProbabilisticSampling.SamplingRate, 0.01)
		assert.Equal(t, "op5", opSampling.PerOperationStrategies[1].Operation)
		assert.InDelta(t, 0.4, opSampling.PerOperationStrategies[1].ProbabilisticSampling.SamplingRate, 0.01)
		assert.Equal(t, "op0", opSampling.PerOperationStrategies[2].Operation)
		assert.InDelta(t, 0.2, opSampling.PerOperationStrategies[2].ProbabilisticSampling.SamplingRate, 0.01)
		assert.Equal(t, "op6", opSampling.PerOperationStrategies[3].Operation)
		assert.InDelta(t, 0.0, opSampling.PerOperationStrategies[3].ProbabilisticSampling.SamplingRate, 0.01)
		assert.Equal(t, "op7", opSampling.PerOperationStrategies[4].Operation)
		assert.InDelta(t, 1.0, opSampling.PerOperationStrategies[4].ProbabilisticSampling.SamplingRate, 0.01)

		s, err = provider.GetSamplingStrategy(context.Background(), "default")
		require.NoError(t, err)
		expectedRsp := makeResponse(api_v2.SamplingStrategyType_PROBABILISTIC, 0.5)
		expectedRsp.OperationSampling = &api_v2.PerOperationSamplingStrategies{
			DefaultSamplingProbability: 0.5,
			PerOperationStrategies: []*api_v2.OperationSamplingStrategy{
				{
					Operation: "op0",
					ProbabilisticSampling: &api_v2.ProbabilisticSamplingStrategy{
						SamplingRate: 0.2,
					},
				},
				{
					Operation: "op6",
					ProbabilisticSampling: &api_v2.ProbabilisticSamplingStrategy{
						SamplingRate: 0,
					},
				},
				{
					Operation: "op7",
					ProbabilisticSampling: &api_v2.ProbabilisticSamplingStrategy{
						SamplingRate: 1,
					},
				},
			},
		}
		assert.EqualValues(t, expectedRsp, *s)
	}
}

func TestMissingServiceSamplingStrategyTypes(t *testing.T) {
	zapCore, logs := observer.New(zap.InfoLevel)
	logger := zap.New(zapCore)
	provider, err := NewFileSource(Options{StrategiesFile: "fixtures/missing-service-types.json"}, logger)
	message := logs.FilterMessage("Failed to parse sampling strategy")
	assert.Equal(t, "Failed to parse sampling strategy", message.All()[0].Message)
	require.NoError(t, err)

	expected := makeResponse(api_v2.SamplingStrategyType_PROBABILISTIC, defaultSamplingProbability)

	s, err := provider.GetSamplingStrategy(context.Background(), "foo")
	require.NoError(t, err)
	assert.Equal(t, api_v2.SamplingStrategyType_PROBABILISTIC, s.StrategyType)
	assert.Equal(t, *expected.ProbabilisticSampling, *s.ProbabilisticSampling)

	require.NotNil(t, s.OperationSampling)
	opSampling := s.OperationSampling
	assert.InDelta(t, defaultSamplingProbability, opSampling.DefaultSamplingProbability, 1e-4)
	require.Len(t, opSampling.PerOperationStrategies, 1)
	assert.Equal(t, "op1", opSampling.PerOperationStrategies[0].Operation)
	assert.InDelta(t, 0.2, opSampling.PerOperationStrategies[0].ProbabilisticSampling.SamplingRate, 0.001)

	expected = makeResponse(api_v2.SamplingStrategyType_PROBABILISTIC, defaultSamplingProbability)

	s, err = provider.GetSamplingStrategy(context.Background(), "bar")
	require.NoError(t, err)
	assert.Equal(t, api_v2.SamplingStrategyType_PROBABILISTIC, s.StrategyType)
	assert.Equal(t, *expected.ProbabilisticSampling, *s.ProbabilisticSampling)

	require.NotNil(t, s.OperationSampling)
	opSampling = s.OperationSampling
	assert.InDelta(t, 0.001, opSampling.DefaultSamplingProbability, 1e-4)
	require.Len(t, opSampling.PerOperationStrategies, 2)
	assert.Equal(t, "op3", opSampling.PerOperationStrategies[0].Operation)
	assert.InDelta(t, 0.3, opSampling.PerOperationStrategies[0].ProbabilisticSampling.SamplingRate, 0.01)
	assert.Equal(t, "op5", opSampling.PerOperationStrategies[1].Operation)
	assert.InDelta(t, 0.4, opSampling.PerOperationStrategies[1].ProbabilisticSampling.SamplingRate, 0.01)

	s, err = provider.GetSamplingStrategy(context.Background(), "default")
	require.NoError(t, err)
	assert.EqualValues(t, makeResponse(api_v2.SamplingStrategyType_PROBABILISTIC, 0.5), *s)
}

func TestParseStrategy(t *testing.T) {
	tests := []struct {
		strategy serviceStrategy
		expected api_v2.SamplingStrategyResponse
	}{
		{
			strategy: serviceStrategy{
				Service:  "svc",
				strategy: strategy{Type: "probabilistic", Param: 0.2},
			},
			expected: makeResponse(api_v2.SamplingStrategyType_PROBABILISTIC, 0.2),
		},
		{
			strategy: serviceStrategy{
				Service:  "svc",
				strategy: strategy{Type: "ratelimiting", Param: 3.5},
			},
			expected: makeResponse(api_v2.SamplingStrategyType_RATE_LIMITING, 3),
		},
	}
	zapCore, logs := observer.New(zap.InfoLevel)
	logger := zap.New(zapCore)
	provider := &samplingProvider{logger: logger}
	for _, test := range tests {
		tt := test
		t.Run("", func(t *testing.T) {
			assert.EqualValues(t, tt.expected, *provider.parseStrategy(&tt.strategy.strategy))
		})
	}
	assert.Empty(t, logs.Len())

	// Test nonexistent strategy type
	actual := *provider.parseStrategy(&strategy{Type: "blah", Param: 3.5})
	expected := makeResponse(api_v2.SamplingStrategyType_PROBABILISTIC, defaultSamplingProbability)
	assert.EqualValues(t, expected, actual)
	message := logs.FilterMessage("Failed to parse sampling strategy")
	assert.Equal(t, 1, message.Len(), "Expected Failed to parse sampling strategy log message.")
}

func makeResponse(samplerType api_v2.SamplingStrategyType, param float64) (resp api_v2.SamplingStrategyResponse) {
	resp.StrategyType = samplerType
	if samplerType == api_v2.SamplingStrategyType_PROBABILISTIC {
		resp.ProbabilisticSampling = &api_v2.ProbabilisticSamplingStrategy{
			SamplingRate: param,
		}
	} else if samplerType == api_v2.SamplingStrategyType_RATE_LIMITING {
		resp.RateLimitingSampling = &api_v2.RateLimitingSamplingStrategy{
			MaxTracesPerSecond: int32(param),
		}
	}
	return resp
}

func TestDeepCopy(t *testing.T) {
	s := &api_v2.SamplingStrategyResponse{
		StrategyType: api_v2.SamplingStrategyType_PROBABILISTIC,
		ProbabilisticSampling: &api_v2.ProbabilisticSamplingStrategy{
			SamplingRate: 0.5,
		},
	}
	cp, err := deepCopy(s)
	require.NoError(t, err)
	assert.NotSame(t, cp, s)
	assert.EqualValues(t, cp, s)
}

func TestAutoUpdateStrategyWithFile(t *testing.T) {
	tempFile, _ := os.CreateTemp("", "for_go_test_*.json")
	require.NoError(t, tempFile.Close())
	defer func() {
		require.NoError(t, os.Remove(tempFile.Name()))
	}()

	// copy known fixture content into temp file which we can later overwrite
	srcFile, dstFile := "fixtures/strategies.json", tempFile.Name()
	srcBytes, err := os.ReadFile(srcFile)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(dstFile, srcBytes, 0o600))

	ss, err := NewFileSource(Options{
		StrategiesFile: dstFile,
		ReloadInterval: time.Millisecond * 10,
	}, zap.NewNop())
	require.NoError(t, err)
	provider := ss.(*samplingProvider)
	defer provider.Close()

	// confirm baseline value
	s, err := provider.GetSamplingStrategy(context.Background(), "foo")
	require.NoError(t, err)
	assert.EqualValues(t, makeResponse(api_v2.SamplingStrategyType_PROBABILISTIC, 0.8), *s)

	// verify that reloading is a no-op
	value := provider.reloadSamplingStrategy(provider.samplingStrategyLoader(dstFile), string(srcBytes))
	assert.Equal(t, string(srcBytes), value)

	// update file with new probability of 0.9
	newStr := strings.Replace(string(srcBytes), "0.8", "0.9", 1)
	require.NoError(t, os.WriteFile(dstFile, []byte(newStr), 0o600))

	// wait for reload timer
	for i := 0; i < 1000; i++ { // wait up to 1sec
		s, err = provider.GetSamplingStrategy(context.Background(), "foo")
		require.NoError(t, err)
		if s.ProbabilisticSampling != nil && s.ProbabilisticSampling.SamplingRate == 0.9 {
			break
		}
		time.Sleep(1 * time.Millisecond)
	}
	assert.EqualValues(t, makeResponse(api_v2.SamplingStrategyType_PROBABILISTIC, 0.9), *s)
}

func TestAutoUpdateStrategyWithURL(t *testing.T) {
	mockServer, mockStrategy := mockStrategyServer(t)
	ss, err := NewFileSource(Options{
		StrategiesFile: mockServer.URL,
		ReloadInterval: 10 * time.Millisecond,
	}, zap.NewNop())
	require.NoError(t, err)
	provider := ss.(*samplingProvider)
	defer provider.Close()

	// confirm baseline value
	s, err := provider.GetSamplingStrategy(context.Background(), "foo")
	require.NoError(t, err)
	assert.EqualValues(t, makeResponse(api_v2.SamplingStrategyType_PROBABILISTIC, 0.8), *s)

	// verify that reloading in no-op
	value := provider.reloadSamplingStrategy(
		provider.samplingStrategyLoader(mockServer.URL),
		*mockStrategy.Load(),
	)
	assert.Equal(t, *mockStrategy.Load(), value)

	// update original strategies with new probability of 0.9
	{
		v09 := strategiesJSON(0.9)
		mockStrategy.Store(&v09)
	}

	// wait for reload timer
	for i := 0; i < 1000; i++ { // wait up to 1sec
		s, err = provider.GetSamplingStrategy(context.Background(), "foo")
		require.NoError(t, err)
		if s.ProbabilisticSampling != nil && s.ProbabilisticSampling.SamplingRate == 0.9 {
			break
		}
		time.Sleep(1 * time.Millisecond)
	}
	assert.EqualValues(t, makeResponse(api_v2.SamplingStrategyType_PROBABILISTIC, 0.9), *s)
}

func TestAutoUpdateStrategyErrors(t *testing.T) {
	tempFile, _ := os.CreateTemp("", "for_go_test_*.json")
	require.NoError(t, tempFile.Close())
	defer func() {
		_ = os.Remove(tempFile.Name())
	}()

	zapCore, logs := observer.New(zap.InfoLevel)
	logger := zap.New(zapCore)

	s, err := NewFileSource(Options{
		StrategiesFile: "fixtures/strategies.json",
		ReloadInterval: time.Hour,
	}, logger)
	require.NoError(t, err)
	provider := s.(*samplingProvider)
	defer provider.Close()

	// check invalid file path or read failure
	assert.Equal(t, "blah", provider.reloadSamplingStrategy(provider.samplingStrategyLoader(tempFile.Name()+"bad-path"), "blah"))
	assert.Len(t, logs.FilterMessage("failed to re-load sampling strategies").All(), 1)

	// check bad file content
	require.NoError(t, os.WriteFile(tempFile.Name(), []byte("bad value"), 0o600))
	assert.Equal(t, "blah", provider.reloadSamplingStrategy(provider.samplingStrategyLoader(tempFile.Name()), "blah"))
	assert.Len(t, logs.FilterMessage("failed to update sampling strategies").All(), 1)

	// check invalid url
	assert.Equal(t, "duh", provider.reloadSamplingStrategy(provider.samplingStrategyLoader("bad-url"), "duh"))
	assert.Len(t, logs.FilterMessage("failed to re-load sampling strategies").All(), 2)

	// check status code other than 200
	mockServer, _ := mockStrategyServer(t)
	assert.Equal(t, "duh", provider.reloadSamplingStrategy(provider.samplingStrategyLoader(mockServer.URL+"/bad-status"), "duh"))
	assert.Len(t, logs.FilterMessage("failed to re-load sampling strategies").All(), 3)

	// check bad content from url
	assert.Equal(t, "duh", provider.reloadSamplingStrategy(provider.samplingStrategyLoader(mockServer.URL+"/bad-content"), "duh"))
	assert.Len(t, logs.FilterMessage("failed to update sampling strategies").All(), 2)
}

func TestServiceNoPerOperationStrategies(t *testing.T) {
	// given setup of strategy provider with no specific per operation sampling strategies
	// and option "sampling.strategies.bugfix-5270=true"
	provider, err := NewFileSource(Options{
		StrategiesFile:             "fixtures/service_no_per_operation.json",
		IncludeDefaultOpStrategies: true,
	}, zap.NewNop())
	require.NoError(t, err)

	for _, service := range []string{"ServiceA", "ServiceB"} {
		t.Run(service, func(t *testing.T) {
			strategy, err := provider.GetSamplingStrategy(context.Background(), service)
			require.NoError(t, err)
			strategyJSON, err := json.MarshalIndent(strategy, "", "  ")
			require.NoError(t, err)

			testName := strings.ReplaceAll(t.Name(), "/", "_")
			snapshotFile := filepath.Join(snapshotLocation, testName+".json")
			expectedServiceResponse, err := os.ReadFile(snapshotFile)
			require.NoError(t, err)

			assert.JSONEq(t, string(expectedServiceResponse), string(strategyJSON),
				"comparing against stored snapshot. Use REGENERATE_SNAPSHOTS=true to rebuild snapshots.")

			if regenerateSnapshots {
				err = os.WriteFile(snapshotFile, strategyJSON, 0o600)
				require.NoError(t, err)
			}
		})
	}
}

func TestServiceNoPerOperationStrategiesDeprecatedBehavior(t *testing.T) {
	// test case to be removed along with removal of strategy_store.parseStrategies_deprecated,
	// see https://github.com/jaegertracing/jaeger/issues/5270 for more details

	// given setup of strategy provider with no specific per operation sampling strategies
	provider, err := NewFileSource(Options{
		StrategiesFile: "fixtures/service_no_per_operation.json",
	}, zap.NewNop())
	require.NoError(t, err)

	for _, service := range []string{"ServiceA", "ServiceB"} {
		t.Run(service, func(t *testing.T) {
			strategy, err := provider.GetSamplingStrategy(context.Background(), service)
			require.NoError(t, err)
			strategyJSON, err := json.MarshalIndent(strategy, "", "  ")
			require.NoError(t, err)

			testName := strings.ReplaceAll(t.Name(), "/", "_")
			snapshotFile := filepath.Join(snapshotLocation, testName+".json")
			expectedServiceResponse, err := os.ReadFile(snapshotFile)
			require.NoError(t, err)

			assert.JSONEq(t, string(expectedServiceResponse), string(strategyJSON),
				"comparing against stored snapshot. Use REGENERATE_SNAPSHOTS=true to rebuild snapshots.")

			if regenerateSnapshots {
				err = os.WriteFile(snapshotFile, strategyJSON, 0o600)
				require.NoError(t, err)
			}
		})
	}
}

func TestSamplingStrategyLoader(t *testing.T) {
	provider := &samplingProvider{logger: zap.NewNop()}
	// invalid file path
	loader := provider.samplingStrategyLoader("not-exists")
	_, err := loader()
	require.ErrorContains(t, err, "failed to read strategies file not-exists")

	// status code other than 200
	mockServer, _ := mockStrategyServer(t)
	loader = provider.samplingStrategyLoader(mockServer.URL + "/bad-status")
	_, err = loader()
	require.ErrorContains(t, err, "receiving 404 Not Found while downloading strategies file")

	// should download content from URL
	loader = provider.samplingStrategyLoader(mockServer.URL + "/bad-content")
	content, err := loader()
	require.NoError(t, err)
	assert.Equal(t, "bad-content", string(content))
}
