module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlserverreceiver

go 1.20

require (
	github.com/google/go-cmp v0.6.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden v0.91.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest v0.91.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/winperfcounters v0.91.0
	github.com/stretchr/testify v1.8.4
	go.opentelemetry.io/collector/component v0.91.1-0.20240109173641-c5a2c78d6143
	go.opentelemetry.io/collector/confmap v0.91.1-0.20240109173641-c5a2c78d6143
	go.opentelemetry.io/collector/consumer v0.91.1-0.20240109173641-c5a2c78d6143
	go.opentelemetry.io/collector/pdata v1.0.1-0.20240109173641-c5a2c78d6143
	go.opentelemetry.io/collector/receiver v0.91.1-0.20240109173641-c5a2c78d6143
	go.opentelemetry.io/otel/metric v1.21.0
	go.opentelemetry.io/otel/trace v1.21.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.26.0
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.0.1 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.1-0.20220423185008-bf980b35cac4 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil v0.91.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/collector v0.91.1-0.20240109173641-c5a2c78d6143 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.91.1-0.20240109173641-c5a2c78d6143 // indirect
	go.opentelemetry.io/collector/featuregate v1.0.1-0.20240109173641-c5a2c78d6143 // indirect
	go.opentelemetry.io/otel v1.21.0 // indirect
	golang.org/x/net v0.18.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d // indirect
	google.golang.org/grpc v1.59.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../pkg/pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../pkg/pdatautil

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/winperfcounters => ../../pkg/winperfcounters

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden => ../../pkg/golden
