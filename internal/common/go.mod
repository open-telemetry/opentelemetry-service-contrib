module github.com/open-telemetry/opentelemetry-collector-contrib/internal/common

go 1.14

require (
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/containerd/containerd v1.3.6 // indirect
	github.com/docker/docker v17.12.0-ce-rc1.0.20200706150819-a40b877fbb9e+incompatible
	github.com/docker/go-connections v0.4.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor v0.0.0
	github.com/stretchr/testify v1.6.1
	go.opentelemetry.io/collector v0.10.1-0.20200922190504-eb2127131b29
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor => ../../processor/resourcedetectionprocessor
