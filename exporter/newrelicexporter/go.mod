module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/newrelicexporter

go 1.14

require (
	github.com/census-instrumentation/opencensus-proto v0.3.0
	github.com/newrelic/newrelic-telemetry-sdk-go v0.4.0
	github.com/stretchr/testify v1.6.1
	go.opencensus.io v0.22.5
	go.opentelemetry.io/collector v0.12.1-0.20201016230751-46aada6e3c3a
	go.uber.org/zap v1.16.0
	google.golang.org/grpc/examples v0.0.0-20200728194956-1c32b02682df // indirect
	google.golang.org/protobuf v1.25.0
)
