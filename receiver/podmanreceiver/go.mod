module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/podmanreceiver

go 1.16

require (
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.33.1-0.20210826200354-479f46434f9a
	go.opentelemetry.io/collector/model v0.33.1-0.20210826200354-479f46434f9a // indirect
	go.uber.org/zap v1.19.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../../internal/common
