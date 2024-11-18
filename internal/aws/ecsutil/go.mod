module github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/ecsutil

go 1.22.0

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.113.0
	github.com/stretchr/testify v1.9.0
	go.opentelemetry.io/collector/component v0.113.1-0.20241115165626-8b99b8023ca3
	go.opentelemetry.io/collector/component/componenttest v0.0.0-20241112045949-9c4c007a12b4
	go.opentelemetry.io/collector/config/confighttp v0.113.1-0.20241115165626-8b99b8023ca3
	go.uber.org/goleak v1.3.0
	go.uber.org/zap v1.27.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rs/cors v1.11.1 // indirect
	go.opentelemetry.io/collector/client v1.19.1-0.20241115165626-8b99b8023ca3 // indirect
	go.opentelemetry.io/collector/config/configauth v0.113.1-0.20241115165626-8b99b8023ca3 // indirect
	go.opentelemetry.io/collector/config/configcompression v1.19.1-0.20241115165626-8b99b8023ca3 // indirect
	go.opentelemetry.io/collector/config/configopaque v1.19.1-0.20241115165626-8b99b8023ca3 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.113.1-0.20241115165626-8b99b8023ca3 // indirect
	go.opentelemetry.io/collector/config/configtls v1.19.1-0.20241115165626-8b99b8023ca3 // indirect
	go.opentelemetry.io/collector/config/internal v0.113.1-0.20241115165626-8b99b8023ca3 // indirect
	go.opentelemetry.io/collector/extension v0.113.1-0.20241115165626-8b99b8023ca3 // indirect
	go.opentelemetry.io/collector/extension/auth v0.113.1-0.20241115165626-8b99b8023ca3 // indirect
	go.opentelemetry.io/collector/pdata v1.19.1-0.20241115165626-8b99b8023ca3 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.56.0 // indirect
	go.opentelemetry.io/otel v1.32.0 // indirect
	go.opentelemetry.io/otel/metric v1.32.0 // indirect
	go.opentelemetry.io/otel/sdk v1.32.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.32.0 // indirect
	go.opentelemetry.io/otel/trace v1.32.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.31.0 // indirect
	golang.org/x/sys v0.27.0 // indirect
	golang.org/x/text v0.20.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240822170219-fc7c04adadcd // indirect
	google.golang.org/grpc v1.67.1 // indirect
	google.golang.org/protobuf v1.35.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../../../internal/common

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)
