module github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/otlpencodingextension

go 1.20

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding v0.92.0
	github.com/stretchr/testify v1.8.4
	go.opentelemetry.io/collector/component v0.92.1-0.20240110091511-bf804d6c4ecc
	go.opentelemetry.io/collector/confmap v0.92.1-0.20240110091511-bf804d6c4ecc
	go.opentelemetry.io/collector/extension v0.92.1-0.20240110091511-bf804d6c4ecc
	go.opentelemetry.io/collector/pdata v1.0.2-0.20240112172857-83d463ceba06
	go.opentelemetry.io/otel/metric v1.21.0
	go.opentelemetry.io/otel/trace v1.21.0
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
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
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.92.1-0.20240110091511-bf804d6c4ecc // indirect
	go.opentelemetry.io/collector/featuregate v1.0.2-0.20240112172857-83d463ceba06 // indirect
	go.opentelemetry.io/otel v1.21.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	golang.org/x/net v0.18.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231012201019-e917dd12ba7a // indirect
	google.golang.org/grpc v1.60.1 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding => ../
