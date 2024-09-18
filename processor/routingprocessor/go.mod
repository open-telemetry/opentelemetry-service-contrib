module github.com/open-telemetry/opentelemetry-collector-contrib/processor/routingprocessor

go 1.22.0

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl v0.109.0
	github.com/stretchr/testify v1.9.0
	go.opentelemetry.io/collector/client v1.15.1-0.20240918193345-a3c0565031b0
	go.opentelemetry.io/collector/component v0.109.1-0.20240918193345-a3c0565031b0
	go.opentelemetry.io/collector/config/configgrpc v0.109.1-0.20240918193345-a3c0565031b0
	go.opentelemetry.io/collector/config/configtelemetry v0.109.1-0.20240918193345-a3c0565031b0
	go.opentelemetry.io/collector/confmap v1.15.1-0.20240918193345-a3c0565031b0
	go.opentelemetry.io/collector/consumer v0.109.1-0.20240918193345-a3c0565031b0
	go.opentelemetry.io/collector/consumer/consumertest v0.109.1-0.20240918193345-a3c0565031b0
	go.opentelemetry.io/collector/exporter v0.109.1-0.20240918193345-a3c0565031b0
	go.opentelemetry.io/collector/exporter/otlpexporter v0.109.1-0.20240918193345-a3c0565031b0
	go.opentelemetry.io/collector/pdata v1.15.1-0.20240918193345-a3c0565031b0
	go.opentelemetry.io/collector/processor v0.109.1-0.20240918193345-a3c0565031b0
	go.opentelemetry.io/otel v1.30.0
	go.opentelemetry.io/otel/metric v1.30.0
	go.opentelemetry.io/otel/sdk/metric v1.30.0
	go.opentelemetry.io/otel/trace v1.30.0
	go.uber.org/goleak v1.3.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.27.0
	google.golang.org/grpc v1.66.2
)

require (
	github.com/alecthomas/participle/v2 v2.1.1 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/elastic/go-grok v0.3.1 // indirect
	github.com/elastic/lunes v0.1.0 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.1.0 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.1 // indirect
	github.com/magefile/mage v1.15.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mostynb/go-grpc-compression v1.2.3 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.109.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/ua-parser/uap-go v0.0.0-20240611065828-3a4781585db6 // indirect
	go.opentelemetry.io/collector v0.109.1-0.20240918193345-a3c0565031b0 // indirect
	go.opentelemetry.io/collector/component/componentstatus v0.109.1-0.20240918193345-a3c0565031b0 // indirect
	go.opentelemetry.io/collector/config/configauth v0.109.1-0.20240918193345-a3c0565031b0 // indirect
	go.opentelemetry.io/collector/config/configcompression v1.15.1-0.20240918193345-a3c0565031b0 // indirect
	go.opentelemetry.io/collector/config/confignet v0.109.1-0.20240918193345-a3c0565031b0 // indirect
	go.opentelemetry.io/collector/config/configopaque v1.15.1-0.20240918193345-a3c0565031b0 // indirect
	go.opentelemetry.io/collector/config/configretry v1.15.1-0.20240918193345-a3c0565031b0 // indirect
	go.opentelemetry.io/collector/config/configtls v1.15.1-0.20240918193345-a3c0565031b0 // indirect
	go.opentelemetry.io/collector/config/internal v0.109.1-0.20240918193345-a3c0565031b0 // indirect
	go.opentelemetry.io/collector/consumer/consumerprofiles v0.109.1-0.20240918193345-a3c0565031b0 // indirect
	go.opentelemetry.io/collector/exporter/exporterprofiles v0.109.1-0.20240918193345-a3c0565031b0 // indirect
	go.opentelemetry.io/collector/extension v0.109.1-0.20240918193345-a3c0565031b0 // indirect
	go.opentelemetry.io/collector/extension/auth v0.109.1-0.20240918193345-a3c0565031b0 // indirect
	go.opentelemetry.io/collector/extension/experimental/storage v0.109.1-0.20240918193345-a3c0565031b0 // indirect
	go.opentelemetry.io/collector/featuregate v1.15.1-0.20240918193345-a3c0565031b0 // indirect
	go.opentelemetry.io/collector/internal/globalgates v0.109.1-0.20240918193345-a3c0565031b0 // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.109.1-0.20240918193345-a3c0565031b0 // indirect
	go.opentelemetry.io/collector/pdata/testdata v0.109.1-0.20240918193345-a3c0565031b0 // indirect
	go.opentelemetry.io/collector/processor/processorprofiles v0.109.1-0.20240918193345-a3c0565031b0 // indirect
	go.opentelemetry.io/collector/receiver v0.109.1-0.20240918193345-a3c0565031b0 // indirect
	go.opentelemetry.io/collector/receiver/receiverprofiles v0.109.1-0.20240918193345-a3c0565031b0 // indirect
	go.opentelemetry.io/collector/semconv v0.109.1-0.20240918193345-a3c0565031b0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.55.0 // indirect
	go.opentelemetry.io/otel/sdk v1.30.0 // indirect
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842 // indirect
	golang.org/x/net v0.29.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/text v0.18.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240903143218-8af14fe29dc1 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl => ../../pkg/ottl

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../pkg/pdatautil

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../pkg/pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden => ../../pkg/golden
