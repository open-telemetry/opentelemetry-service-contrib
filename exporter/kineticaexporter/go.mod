module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kineticaexporter

go 1.20

require (
	bitbucket.org/gisfederal/gpudb-api-go v0.0.9
	github.com/google/uuid v1.3.0
	github.com/stretchr/testify v1.8.4
	github.com/wk8/go-ordered-map v1.0.0
	go.opentelemetry.io/collector/component v0.85.1-0.20230919160920-512b9c3c8fec
	go.opentelemetry.io/collector/config/configopaque v0.85.1-0.20230919160920-512b9c3c8fec
	go.opentelemetry.io/collector/confmap v0.85.1-0.20230919160920-512b9c3c8fec
	go.opentelemetry.io/collector/exporter v0.85.1-0.20230919160920-512b9c3c8fec
	go.opentelemetry.io/collector/pdata v1.0.0-rcv0014.0.20230919160920-512b9c3c8fec
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.26.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-resty/resty/v2 v2.7.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/hamba/avro/v2 v2.13.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf v1.5.0 // indirect
	github.com/knadh/koanf/v2 v2.0.1 // indirect
	github.com/logrusorgru/aurora v2.0.3+incompatible // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.1-0.20220423185008-bf980b35cac4 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/samber/lo v1.38.1 // indirect
	github.com/ztrue/tracerr v0.3.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/collector v0.85.1-0.20230919160920-512b9c3c8fec // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.85.1-0.20230919160920-512b9c3c8fec // indirect
	go.opentelemetry.io/collector/consumer v0.85.1-0.20230919160920-512b9c3c8fec // indirect
	go.opentelemetry.io/collector/extension v0.85.1-0.20230919160920-512b9c3c8fec // indirect
	go.opentelemetry.io/collector/featuregate v1.0.0-rcv0014.0.20230919160920-512b9c3c8fec // indirect
	go.opentelemetry.io/collector/processor v0.85.1-0.20230919160920-512b9c3c8fec // indirect
	go.opentelemetry.io/collector/receiver v0.85.1-0.20230919160920-512b9c3c8fec // indirect
	go.opentelemetry.io/otel v1.18.0 // indirect
	go.opentelemetry.io/otel/metric v1.18.0 // indirect
	go.opentelemetry.io/otel/sdk v1.18.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v0.41.0 // indirect
	go.opentelemetry.io/otel/trace v1.18.0 // indirect
	golang.org/x/exp v0.0.0-20220303212507-bbda1eaf7a17 // indirect
	golang.org/x/net v0.15.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230911183012-2d3300fd4832 // indirect
	google.golang.org/grpc v1.58.1 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)
