module github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/hostobserver

go 1.16

require (
	github.com/mattn/go-colorable v0.1.7 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer v0.0.0-00010101000000-000000000000
	github.com/shirou/gopsutil v3.21.7+incompatible
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.31.1-0.20210804191544-3cfe4f8c5d3e
	go.opentelemetry.io/collector/model v0.31.1-0.20210804191544-3cfe4f8c5d3e // indirect
	go.uber.org/zap v1.18.1
	gopkg.in/square/go-jose.v2 v2.5.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer => ../
