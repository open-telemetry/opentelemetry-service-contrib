module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jmxreceiver

go 1.16

require (
	github.com/mattn/go-colorable v0.1.7 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/testutil v0.0.0-00010101000000-000000000000
	github.com/shirou/gopsutil v3.21.7+incompatible
	github.com/stretchr/testify v1.7.0
	github.com/testcontainers/testcontainers-go v0.11.1
	go.opentelemetry.io/collector v0.33.1-0.20210820002854-d3000232f8f6
	go.opentelemetry.io/collector/model v0.33.1-0.20210820002854-d3000232f8f6
	go.uber.org/atomic v1.9.0
	go.uber.org/zap v1.19.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/testutil => ../../internal/coreinternal/testutil
