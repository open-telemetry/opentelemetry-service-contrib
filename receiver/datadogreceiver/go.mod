module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/datadogreceiver

go 1.19

require (
	github.com/DataDog/datadog-agent/pkg/trace v0.45.0-rc.3
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/sharedcomponent v0.75.0
	github.com/stretchr/testify v1.8.2
	github.com/vmihailenco/msgpack/v4 v4.3.12
	go.opentelemetry.io/collector v0.76.1-0.20230426191218-56daa378f504
	go.opentelemetry.io/collector/component v0.76.1-0.20230426191218-56daa378f504
	go.opentelemetry.io/collector/consumer v0.76.1-0.20230426191218-56daa378f504
	go.opentelemetry.io/collector/pdata v1.0.0-rc9.0.20230426191218-56daa378f504
	go.opentelemetry.io/collector/receiver v0.76.1-0.20230426191218-56daa378f504
	go.opentelemetry.io/collector/semconv v0.76.1-0.20230426191218-56daa378f504
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.16.5 // indirect
	github.com/knadh/koanf v1.5.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/philhofer/fwd v1.1.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rs/cors v1.9.0 // indirect
	github.com/tinylib/msgp v1.1.8 // indirect
	github.com/vmihailenco/tagparser v0.1.2 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/collector/confmap v0.76.1-0.20230426191218-56daa378f504 // indirect
	go.opentelemetry.io/collector/exporter v0.76.1-0.20230426191218-56daa378f504 // indirect
	go.opentelemetry.io/collector/featuregate v0.76.1-0.20230426191218-56daa378f504 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.40.0 // indirect
	go.opentelemetry.io/otel v1.14.0 // indirect
	go.opentelemetry.io/otel/metric v0.37.0 // indirect
	go.opentelemetry.io/otel/trace v1.14.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20230306155012-7f2fa6fef1f4 // indirect
	google.golang.org/grpc v1.54.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/sharedcomponent => ../../internal/sharedcomponent
