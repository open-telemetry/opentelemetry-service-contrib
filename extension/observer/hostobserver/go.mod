module github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/hostobserver

go 1.14

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer v0.0.0-00010101000000-000000000000
	github.com/shirou/gopsutil v3.20.12+incompatible
	github.com/stretchr/testify v1.6.1
	go.opentelemetry.io/collector v0.17.1-0.20210111170652-28e3cc1d73ab
	go.uber.org/zap v1.16.0
	google.golang.org/grpc/examples v0.0.0-20200728194956-1c32b02682df // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer => ../
