module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/dockerstatsreceiver

go 1.15

require (
	github.com/census-instrumentation/opencensus-proto v0.3.0
	github.com/docker/docker v20.10.6+incompatible
	github.com/gobwas/glob v0.2.3
	github.com/golang/protobuf v1.5.2
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.0.0-00010101000000-000000000000
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redisreceiver v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.26.1-0.20210513162346-453d1d0dd603
	go.uber.org/zap v1.16.0
	google.golang.org/protobuf v1.26.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../../internal/common

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redisreceiver => ../redisreceiver
