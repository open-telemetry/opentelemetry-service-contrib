module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/dockerstatsreceiver

go 1.16

require (
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/docker/docker v20.10.8+incompatible
	github.com/gobwas/glob v0.2.3
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.0.0-00010101000000-000000000000
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/interval v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.33.1-0.20210827152330-09258f969908
	go.opentelemetry.io/collector/model v0.33.1-0.20210827152330-09258f969908
	go.uber.org/zap v1.19.0
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.3 // indirect
)

replace (
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../../internal/common
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/interval => ../../internal/interval
)
