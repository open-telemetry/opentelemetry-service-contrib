module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nsxtreceiver

go 1.20

require (
	github.com/google/go-cmp v0.6.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest v0.92.0
	github.com/stretchr/testify v1.8.4
	github.com/vmware/go-vmware-nsxt v0.0.0-20230223012718-d31b8a1ca05e
	go.opentelemetry.io/collector/component v0.92.1-0.20240110091511-bf804d6c4ecc
	go.opentelemetry.io/collector/config/confighttp v0.92.1-0.20240110091511-bf804d6c4ecc
	go.opentelemetry.io/collector/config/configopaque v0.92.1-0.20240110091511-bf804d6c4ecc
	go.opentelemetry.io/collector/confmap v0.92.1-0.20240110091511-bf804d6c4ecc
	go.opentelemetry.io/collector/consumer v0.92.1-0.20240110091511-bf804d6c4ecc
	go.opentelemetry.io/collector/pdata v1.0.2-0.20240110091511-bf804d6c4ecc
	go.opentelemetry.io/collector/receiver v0.92.1-0.20240110091511-bf804d6c4ecc
	go.opentelemetry.io/otel/metric v1.21.0
	go.opentelemetry.io/otel/trace v1.21.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.26.0
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.4 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.0.1 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.1-0.20220423185008-bf980b35cac4 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil v0.92.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rs/cors v1.10.1 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/collector v0.92.1-0.20240110091511-bf804d6c4ecc // indirect
	go.opentelemetry.io/collector/config/configauth v0.92.1-0.20240110091511-bf804d6c4ecc // indirect
	go.opentelemetry.io/collector/config/configcompression v0.92.1-0.20240110091511-bf804d6c4ecc // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.92.1-0.20240110091511-bf804d6c4ecc // indirect
	go.opentelemetry.io/collector/config/configtls v0.92.1-0.20240110091511-bf804d6c4ecc // indirect
	go.opentelemetry.io/collector/config/internal v0.92.1-0.20240110091511-bf804d6c4ecc // indirect
	go.opentelemetry.io/collector/extension v0.92.1-0.20240110091511-bf804d6c4ecc // indirect
	go.opentelemetry.io/collector/extension/auth v0.92.1-0.20240110091511-bf804d6c4ecc // indirect
	go.opentelemetry.io/collector/featuregate v1.0.2-0.20240110091511-bf804d6c4ecc // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.46.1 // indirect
	go.opentelemetry.io/otel v1.21.0 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231002182017-d307bd883b97 // indirect
	google.golang.org/grpc v1.60.1 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../pkg/pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../pkg/pdatautil

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden => ../../pkg/golden
