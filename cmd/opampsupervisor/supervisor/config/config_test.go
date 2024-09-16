// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/open-telemetry/opamp-go/protobufs"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/config/configtls"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
		name          string
		config        Supervisor
		expectedError string
	}{
		{
			name: "Valid filled out config",
			config: Supervisor{
				Server: OpAMPServer{
					Endpoint: "wss://localhost:9090/opamp",
					Headers: http.Header{
						"Header1": []string{"HeaderValue"},
					},
					TLSSetting: configtls.ClientConfig{
						Insecure: true,
					},
				},
				Agent: Agent{
					Executable:              "${file_path}",
					OrphanDetectionInterval: 5 * time.Second,
				},
				Capabilities: Capabilities{
					AcceptsRemoteConfig: true,
				},
				Storage: Storage{
					Directory: "/etc/opamp-supervisor/storage",
				},
			},
		},
		{
			name: "Endpoint unspecified",
			config: Supervisor{
				Server: OpAMPServer{
					Headers: http.Header{
						"Header1": []string{"HeaderValue"},
					},
					TLSSetting: configtls.ClientConfig{
						Insecure: true,
					},
				},
				Agent: Agent{
					Executable:              "${file_path}",
					OrphanDetectionInterval: 5 * time.Second,
				},
				Capabilities: Capabilities{
					AcceptsRemoteConfig: true,
				},
				Storage: Storage{
					Directory: "/etc/opamp-supervisor/storage",
				},
			},
			expectedError: "server::endpoint must be specified",
		},
		{
			name: "Invalid URL",
			config: Supervisor{
				Server: OpAMPServer{
					Endpoint: "\000",
					Headers: http.Header{
						"Header1": []string{"HeaderValue"},
					},
					TLSSetting: configtls.ClientConfig{
						Insecure: true,
					},
				},
				Agent: Agent{
					Executable:              "${file_path}",
					OrphanDetectionInterval: 5 * time.Second,
				},
				Capabilities: Capabilities{
					AcceptsRemoteConfig: true,
				},
				Storage: Storage{
					Directory: "/etc/opamp-supervisor/storage",
				},
			},
			expectedError: "invalid URL for server::endpoint:",
		},
		{
			name: "Invalid endpoint scheme",
			config: Supervisor{
				Server: OpAMPServer{
					Endpoint: "tcp://localhost:9090/opamp",
					Headers: http.Header{
						"Header1": []string{"HeaderValue"},
					},
					TLSSetting: configtls.ClientConfig{
						Insecure: true,
					},
				},
				Agent: Agent{
					Executable:              "${file_path}",
					OrphanDetectionInterval: 5 * time.Second,
				},
				Capabilities: Capabilities{
					AcceptsRemoteConfig: true,
				},
				Storage: Storage{
					Directory: "/etc/opamp-supervisor/storage",
				},
			},
			expectedError: `invalid scheme "tcp" for server::endpoint, must be one of "http", "https", "ws", or "wss"`,
		},
		{
			name: "Invalid tls settings",
			config: Supervisor{
				Server: OpAMPServer{
					Endpoint: "wss://localhost:9090/opamp",
					Headers: http.Header{
						"Header1": []string{"HeaderValue"},
					},
					TLSSetting: configtls.ClientConfig{
						Insecure: true,
						Config: configtls.Config{
							MaxVersion: "1.2",
							MinVersion: "1.3",
						},
					},
				},
				Agent: Agent{
					Executable:              "${file_path}",
					OrphanDetectionInterval: 5 * time.Second,
				},
				Capabilities: Capabilities{
					AcceptsRemoteConfig: true,
				},
				Storage: Storage{
					Directory: "/etc/opamp-supervisor/storage",
				},
			},
			expectedError: "invalid server::tls settings:",
		},
		{
			name: "Empty agent executable path",
			config: Supervisor{
				Server: OpAMPServer{
					Endpoint: "wss://localhost:9090/opamp",
					Headers: http.Header{
						"Header1": []string{"HeaderValue"},
					},
					TLSSetting: configtls.ClientConfig{
						Insecure: true,
					},
				},
				Agent: Agent{
					Executable:              "",
					OrphanDetectionInterval: 5 * time.Second,
				},
				Capabilities: Capabilities{
					AcceptsRemoteConfig: true,
				},
				Storage: Storage{
					Directory: "/etc/opamp-supervisor/storage",
				},
			},
			expectedError: "agent::executable must be specified",
		},
		{
			name: "agent executable does not exist",
			config: Supervisor{
				Server: OpAMPServer{
					Endpoint: "wss://localhost:9090/opamp",
					Headers: http.Header{
						"Header1": []string{"HeaderValue"},
					},
					TLSSetting: configtls.ClientConfig{
						Insecure: true,
					},
				},
				Agent: Agent{
					Executable:              "./path/does/not/exist",
					OrphanDetectionInterval: 5 * time.Second,
				},
				Capabilities: Capabilities{
					AcceptsRemoteConfig: true,
				},
				Storage: Storage{
					Directory: "/etc/opamp-supervisor/storage",
				},
			},
			expectedError: "could not stat agent::executable path:",
		},
		{
			name: "Invalid orphan detection interval",
			config: Supervisor{
				Server: OpAMPServer{
					Endpoint: "wss://localhost:9090/opamp",
					Headers: http.Header{
						"Header1": []string{"HeaderValue"},
					},
					TLSSetting: configtls.ClientConfig{
						Insecure: true,
					},
				},
				Agent: Agent{
					Executable:              "${file_path}",
					OrphanDetectionInterval: -1,
				},
				Capabilities: Capabilities{
					AcceptsRemoteConfig: true,
				},
				Storage: Storage{
					Directory: "/etc/opamp-supervisor/storage",
				},
			},
			expectedError: "agent::orphan_detection_interval must be positive",
		},
		{
			name: "Invalid port number",
			config: Supervisor{
				Server: OpAMPServer{
					Endpoint: "wss://localhost:9090/opamp",
					Headers: http.Header{
						"Header1": []string{"HeaderValue"},
					},
					TLSSetting: configtls.ClientConfig{
						Insecure: true,
					},
				},
				Agent: Agent{
					Executable:              "${file_path}",
					OrphanDetectionInterval: 5 * time.Second,
					HealthCheckPort:         65536,
				},
				Capabilities: Capabilities{
					AcceptsRemoteConfig: true,
				},
				Storage: Storage{
					Directory: "/etc/opamp-supervisor/storage",
				},
			},
			expectedError: "agent::health_check_port must be a valid port number",
		},
		{
			name: "Zero value port number",
			config: Supervisor{
				Server: OpAMPServer{
					Endpoint: "wss://localhost:9090/opamp",
					Headers: http.Header{
						"Header1": []string{"HeaderValue"},
					},
					TLSSetting: configtls.ClientConfig{
						Insecure: true,
					},
				},
				Agent: Agent{
					Executable:              "${file_path}",
					OrphanDetectionInterval: 5 * time.Second,
					HealthCheckPort:         0,
				},
				Capabilities: Capabilities{
					AcceptsRemoteConfig: true,
				},
				Storage: Storage{
					Directory: "/etc/opamp-supervisor/storage",
				},
			},
		},
		{
			name: "Normal port number",
			config: Supervisor{
				Server: OpAMPServer{
					Endpoint: "wss://localhost:9090/opamp",
					Headers: http.Header{
						"Header1": []string{"HeaderValue"},
					},
					TLSSetting: configtls.ClientConfig{
						Insecure: true,
					},
				},
				Agent: Agent{
					Executable:              "${file_path}",
					OrphanDetectionInterval: 5 * time.Second,
					HealthCheckPort:         29848,
				},
				Capabilities: Capabilities{
					AcceptsRemoteConfig: true,
				},
				Storage: Storage{
					Directory: "/etc/opamp-supervisor/storage",
				},
			},
		},
	}

	// create some fake files for validating agent config
	tmpDir := t.TempDir()

	filePath := filepath.Join(tmpDir, "file")
	require.NoError(t, os.WriteFile(filePath, []byte{}, 0600))

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Fill in path to agent executable
			tc.config.Agent.Executable = os.Expand(tc.config.Agent.Executable,
				func(s string) string {
					if s == "file_path" {
						return filePath
					}
					return ""
				})

			err := tc.config.Validate()

			if tc.expectedError == "" {
				require.NoError(t, err)
			} else {
				require.ErrorContains(t, err, tc.expectedError)
			}
		})
	}
}

func TestCapabilities_SupportedCapabilities(t *testing.T) {
	testCases := []struct {
		name                      string
		capabilities              Capabilities
		expectedAgentCapabilities protobufs.AgentCapabilities
	}{
		{
			name:         "Default capabilities",
			capabilities: DefaultSupervisor().Capabilities,
			expectedAgentCapabilities: protobufs.AgentCapabilities_AgentCapabilities_ReportsStatus |
				protobufs.AgentCapabilities_AgentCapabilities_ReportsOwnMetrics |
				protobufs.AgentCapabilities_AgentCapabilities_ReportsEffectiveConfig |
				protobufs.AgentCapabilities_AgentCapabilities_ReportsHealth,
		},
		{
			name:                      "Empty capabilities",
			capabilities:              Capabilities{},
			expectedAgentCapabilities: protobufs.AgentCapabilities_AgentCapabilities_ReportsStatus,
		},
		{
			name: "Many capabilities",
			capabilities: Capabilities{
				AcceptsRemoteConfig:            true,
				AcceptsRestartCommand:          true,
				AcceptsOpAMPConnectionSettings: true,
				ReportsEffectiveConfig:         true,
				ReportsOwnMetrics:              true,
				ReportsHealth:                  true,
				ReportsRemoteConfig:            true,
			},
			expectedAgentCapabilities: protobufs.AgentCapabilities_AgentCapabilities_ReportsStatus |
				protobufs.AgentCapabilities_AgentCapabilities_ReportsEffectiveConfig |
				protobufs.AgentCapabilities_AgentCapabilities_ReportsHealth |
				protobufs.AgentCapabilities_AgentCapabilities_ReportsOwnMetrics |
				protobufs.AgentCapabilities_AgentCapabilities_AcceptsRemoteConfig |
				protobufs.AgentCapabilities_AgentCapabilities_ReportsRemoteConfig |
				protobufs.AgentCapabilities_AgentCapabilities_AcceptsRestartCommand |
				protobufs.AgentCapabilities_AgentCapabilities_AcceptsOpAMPConnectionSettings,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expectedAgentCapabilities, tc.capabilities.SupportedCapabilities())
		})
	}
}
