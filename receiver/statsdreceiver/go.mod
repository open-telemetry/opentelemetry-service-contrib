module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver

go 1.16

require (
	github.com/mattn/go-colorable v0.1.7 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe
	github.com/stretchr/testify v1.7.0
	go.opencensus.io v0.23.0
	go.opentelemetry.io/collector v0.29.1-0.20210712222151-aa60edff162c
	go.opentelemetry.io/collector/model v0.0.0-00010101000000-000000000000
	go.opentelemetry.io/otel v1.0.0-RC1
	go.uber.org/zap v1.18.1
	gopkg.in/square/go-jose.v2 v2.5.1 // indirect
)

replace go.opentelemetry.io/collector/model => go.opentelemetry.io/collector/model v0.0.0-20210712222151-aa60edff162c
