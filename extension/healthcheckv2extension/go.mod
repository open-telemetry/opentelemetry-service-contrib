module github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckv2extension

go 1.21.0

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.101.0
	github.com/stretchr/testify v1.9.0
	go.opentelemetry.io/collector/component v0.101.1-0.20240523155058-812210ba3685
	go.opentelemetry.io/collector/config/configgrpc v0.101.1-0.20240523155058-812210ba3685
	go.opentelemetry.io/collector/config/confighttp v0.101.1-0.20240523155058-812210ba3685
	go.opentelemetry.io/collector/config/confignet v0.101.1-0.20240523155058-812210ba3685
	go.opentelemetry.io/collector/config/configtls v0.101.1-0.20240523155058-812210ba3685
	go.opentelemetry.io/collector/confmap v0.101.1-0.20240523155058-812210ba3685
	go.opentelemetry.io/collector/extension v0.101.1-0.20240523155058-812210ba3685
	go.opentelemetry.io/otel/metric v1.27.0
	go.opentelemetry.io/otel/trace v1.27.0
	go.uber.org/goleak v1.3.0
	go.uber.org/zap v1.27.0
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.0.0-alpha.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/klauspost/compress v1.17.8 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.1 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/mostynb/go-grpc-compression v1.2.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.19.1 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.53.0 // indirect
	github.com/prometheus/procfs v0.15.0 // indirect
	github.com/rs/cors v1.10.1 // indirect
	go.opentelemetry.io/collector v0.101.1-0.20240523155058-812210ba3685 // indirect
	go.opentelemetry.io/collector/config/configauth v0.101.1-0.20240523155058-812210ba3685 // indirect
	go.opentelemetry.io/collector/config/configcompression v1.8.1-0.20240523143024-6f5d43f9e405 // indirect
	go.opentelemetry.io/collector/config/configopaque v1.8.1-0.20240523143024-6f5d43f9e405 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.101.1-0.20240523155058-812210ba3685 // indirect
	go.opentelemetry.io/collector/config/internal v0.101.1-0.20240523155058-812210ba3685 // indirect
	go.opentelemetry.io/collector/extension/auth v0.101.1-0.20240523155058-812210ba3685 // indirect
	go.opentelemetry.io/collector/featuregate v1.8.1-0.20240523143024-6f5d43f9e405 // indirect
	go.opentelemetry.io/collector/pdata v1.8.1-0.20240523143024-6f5d43f9e405 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.51.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.51.0 // indirect
	go.opentelemetry.io/otel v1.27.0 // indirect
	go.opentelemetry.io/otel/exporters/prometheus v0.49.0 // indirect
	go.opentelemetry.io/otel/sdk v1.27.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.27.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240401170217-c3f982113cda // indirect
	google.golang.org/grpc v1.64.0 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../../internal/common
