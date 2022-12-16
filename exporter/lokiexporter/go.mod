module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/lokiexporter

go 1.18

require (
	github.com/gogo/protobuf v1.3.2
	github.com/golang/snappy v0.0.4
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.67.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.67.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/loki v0.67.0
	github.com/prometheus/common v0.38.0
	github.com/stretchr/testify v1.8.1
	go.opentelemetry.io/collector v0.67.1-0.20221216193913-557c077e3cc3
	go.opentelemetry.io/collector/component v0.67.1-0.20221216193913-557c077e3cc3
	go.opentelemetry.io/collector/confmap v0.67.1-0.20221216193913-557c077e3cc3
	go.opentelemetry.io/collector/consumer v0.67.1-0.20221216193913-557c077e3cc3
	go.opentelemetry.io/collector/pdata v1.0.0-rc1.0.20221216193913-557c077e3cc3
	go.opentelemetry.io/collector/semconv v0.67.1-0.20221216200611-892d07dffbb8
	go.uber.org/multierr v1.9.0
	go.uber.org/zap v1.24.0
)

require (
	github.com/cenkalti/backoff/v4 v4.2.0 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grafana/regexp v0.0.0-20221005093135-b4c2bcb0a4b6 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.15.13 // indirect
	github.com/knadh/koanf v1.4.4 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/prometheus v0.40.6 // indirect
	github.com/rs/cors v1.8.2 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/collector/featuregate v0.67.1-0.20221216193913-557c077e3cc3 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.37.0 // indirect
	go.opentelemetry.io/otel v1.11.2 // indirect
	go.opentelemetry.io/otel/metric v0.34.0 // indirect
	go.opentelemetry.io/otel/trace v1.11.2 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	golang.org/x/net v0.4.0 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	google.golang.org/genproto v0.0.0-20221027153422-115e99e71e1c // indirect
	google.golang.org/grpc v1.51.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../../internal/common

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/loki => ../../pkg/translator/loki

retract v0.65.0
