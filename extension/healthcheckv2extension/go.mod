module github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckv2extension

go 1.22.0

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.114.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/status v0.114.0
	github.com/stretchr/testify v1.10.0
	go.opentelemetry.io/collector/component v0.114.1-0.20241130171227-c52d625647be
	go.opentelemetry.io/collector/component/componentstatus v0.114.1-0.20241130171227-c52d625647be
	go.opentelemetry.io/collector/component/componenttest v0.114.1-0.20241130171227-c52d625647be
	go.opentelemetry.io/collector/config/configgrpc v0.114.1-0.20241130171227-c52d625647be
	go.opentelemetry.io/collector/config/confighttp v0.114.1-0.20241130171227-c52d625647be
	go.opentelemetry.io/collector/config/confignet v1.20.1-0.20241130171227-c52d625647be
	go.opentelemetry.io/collector/config/configtls v1.20.1-0.20241130171227-c52d625647be
	go.opentelemetry.io/collector/confmap v1.20.1-0.20241130171227-c52d625647be
	go.opentelemetry.io/collector/extension v0.114.1-0.20241130171227-c52d625647be
	go.opentelemetry.io/collector/extension/extensioncapabilities v0.114.1-0.20241130171227-c52d625647be
	go.opentelemetry.io/collector/extension/extensiontest v0.114.1-0.20241130171227-c52d625647be
	go.uber.org/goleak v1.3.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.27.0
	google.golang.org/grpc v1.67.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.2 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/mostynb/go-grpc-compression v1.2.3 // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rs/cors v1.11.1 // indirect
	go.opentelemetry.io/collector/client v1.20.1-0.20241130171227-c52d625647be // indirect
	go.opentelemetry.io/collector/config/configauth v0.114.1-0.20241130171227-c52d625647be // indirect
	go.opentelemetry.io/collector/config/configcompression v1.20.1-0.20241130171227-c52d625647be // indirect
	go.opentelemetry.io/collector/config/configopaque v1.20.1-0.20241130171227-c52d625647be // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.114.1-0.20241130171227-c52d625647be // indirect
	go.opentelemetry.io/collector/config/internal v0.114.1-0.20241130171227-c52d625647be // indirect
	go.opentelemetry.io/collector/extension/auth v0.114.1-0.20241130171227-c52d625647be // indirect
	go.opentelemetry.io/collector/featuregate v1.20.1-0.20241130171227-c52d625647be // indirect
	go.opentelemetry.io/collector/pdata v1.20.1-0.20241130171227-c52d625647be // indirect
	go.opentelemetry.io/collector/pipeline v0.114.1-0.20241130171227-c52d625647be // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.56.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.56.0 // indirect
	go.opentelemetry.io/otel v1.32.0 // indirect
	go.opentelemetry.io/otel/metric v1.32.0 // indirect
	go.opentelemetry.io/otel/sdk v1.32.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.32.0 // indirect
	go.opentelemetry.io/otel/trace v1.32.0 // indirect
	golang.org/x/net v0.31.0 // indirect
	golang.org/x/sys v0.27.0 // indirect
	golang.org/x/text v0.20.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241007155032-5fefd90f89a9 // indirect
	google.golang.org/protobuf v1.35.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../../internal/common

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/status => ../../pkg/status
