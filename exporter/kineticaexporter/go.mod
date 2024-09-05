module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kineticaexporter

go 1.22.0

require (
	github.com/google/uuid v1.6.0
	github.com/kineticadb/kinetica-api-go v0.0.5
	github.com/samber/lo v1.47.0
	github.com/stretchr/testify v1.9.0
	github.com/wk8/go-ordered-map/v2 v2.1.8
	go.opentelemetry.io/collector/component v0.108.2-0.20240904075637-48b11ba1c5f8
	go.opentelemetry.io/collector/config/configopaque v1.14.2-0.20240904075637-48b11ba1c5f8
	go.opentelemetry.io/collector/confmap v1.14.2-0.20240904075637-48b11ba1c5f8
	go.opentelemetry.io/collector/exporter v0.108.2-0.20240904075637-48b11ba1c5f8
	go.opentelemetry.io/collector/pdata v1.14.2-0.20240904075637-48b11ba1c5f8
	go.uber.org/goleak v1.3.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.27.0
)

require (
	github.com/bahlo/generic-list-go v0.2.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-resty/resty/v2 v2.12.0 // indirect
	github.com/go-viper/mapstructure/v2 v2.1.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/hamba/avro/v2 v2.25.1 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.1 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.1-0.20231216201459-8508981c8b6c // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_golang v1.20.2 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.57.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/ztrue/tracerr v0.4.0 // indirect
	go.opentelemetry.io/collector v0.108.2-0.20240904075637-48b11ba1c5f8 // indirect
	go.opentelemetry.io/collector/config/configretry v1.14.2-0.20240904075637-48b11ba1c5f8 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.108.2-0.20240904075637-48b11ba1c5f8 // indirect
	go.opentelemetry.io/collector/consumer v0.108.2-0.20240904075637-48b11ba1c5f8 // indirect
	go.opentelemetry.io/collector/consumer/consumerprofiles v0.108.2-0.20240904075637-48b11ba1c5f8 // indirect
	go.opentelemetry.io/collector/consumer/consumertest v0.108.2-0.20240904075637-48b11ba1c5f8 // indirect
	go.opentelemetry.io/collector/exporter/exporterprofiles v0.108.1 // indirect
	go.opentelemetry.io/collector/extension v0.108.2-0.20240904075637-48b11ba1c5f8 // indirect
	go.opentelemetry.io/collector/extension/experimental/storage v0.0.0-20240902085631-cd1f7623fe67 // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.108.2-0.20240904075637-48b11ba1c5f8 // indirect
	go.opentelemetry.io/collector/receiver v0.108.2-0.20240904075637-48b11ba1c5f8 // indirect
	go.opentelemetry.io/collector/receiver/receiverprofiles v0.108.2-0.20240904075637-48b11ba1c5f8 // indirect
	go.opentelemetry.io/otel v1.29.0 // indirect
	go.opentelemetry.io/otel/exporters/prometheus v0.51.0 // indirect
	go.opentelemetry.io/otel/metric v1.29.0 // indirect
	go.opentelemetry.io/otel/sdk v1.29.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.29.0 // indirect
	go.opentelemetry.io/otel/trace v1.29.0 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240822170219-fc7c04adadcd // indirect
	google.golang.org/grpc v1.66.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)
