module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/honeycombexporter

go 1.14

require (
	github.com/census-instrumentation/opencensus-proto v0.3.0
	github.com/google/go-cmp v0.5.4
	github.com/honeycombio/libhoney-go v1.15.2
	github.com/klauspost/compress v1.11.7
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.18.1-0.20210121231022-15322cf01768
	go.uber.org/zap v1.16.0
	google.golang.org/grpc/examples v0.0.0-20200728194956-1c32b02682df // indirect
	google.golang.org/protobuf v1.25.0
)
