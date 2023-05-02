module github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextension

go 1.19

require (
	github.com/jaegertracing/jaeger v1.41.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.76.3
	github.com/stretchr/testify v1.8.2
	go.opencensus.io v0.24.0
	go.opentelemetry.io/collector v0.76.2-0.20230502195822-4df44379e094
	go.opentelemetry.io/collector/component v0.76.2-0.20230502195822-4df44379e094
	go.opentelemetry.io/collector/confmap v0.76.2-0.20230502195822-4df44379e094
	go.uber.org/zap v1.24.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/klauspost/compress v1.16.5 // indirect
	github.com/knadh/koanf v1.5.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rs/cors v1.9.0 // indirect
	go.opentelemetry.io/collector/featuregate v0.76.2-0.20230502195822-4df44379e094 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.41.0 // indirect
	go.opentelemetry.io/otel v1.15.0 // indirect
	go.opentelemetry.io/otel/metric v0.38.0 // indirect
	go.opentelemetry.io/otel/trace v1.15.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/grpc v1.54.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../../internal/common

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)
