module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusexecreceiver

go 1.14

require (
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51
	github.com/prometheus/common v0.14.0
	github.com/prometheus/prometheus v1.8.2-0.20200827201422-1195cc24e3c8
	github.com/shirou/gopsutil v2.20.9+incompatible // indirect
	github.com/stretchr/testify v1.6.1
	go.opentelemetry.io/collector v0.11.1-0.20201006165100-07236c11fb27
	go.uber.org/zap v1.16.0
)
