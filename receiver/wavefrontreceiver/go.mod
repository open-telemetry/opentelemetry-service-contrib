module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/wavefrontreceiver

go 1.21.0

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/collectd v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver v0.103.0
	github.com/stretchr/testify v1.9.0
	go.opentelemetry.io/collector/component v0.103.1-0.20240628190926-ee4eb85f58e9
	go.opentelemetry.io/collector/config/confignet v0.103.1-0.20240628190926-ee4eb85f58e9
	go.opentelemetry.io/collector/confmap v0.103.1-0.20240628190926-ee4eb85f58e9
	go.opentelemetry.io/collector/consumer v0.103.1-0.20240628190926-ee4eb85f58e9
	go.opentelemetry.io/collector/pdata v1.10.1-0.20240701160413-fd36d05133d0
	go.opentelemetry.io/collector/receiver v0.103.1-0.20240628190926-ee4eb85f58e9
	go.opentelemetry.io/otel/metric v1.27.0
	go.opentelemetry.io/otel/trace v1.27.0
	go.uber.org/goleak v1.3.0
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.0.0-alpha.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.1 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.19.1 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.54.0 // indirect
	github.com/prometheus/procfs v0.15.0 // indirect
	go.opentelemetry.io/collector v0.103.1-0.20240628190926-ee4eb85f58e9 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.103.1-0.20240628190926-ee4eb85f58e9 // indirect
	go.opentelemetry.io/collector/featuregate v1.10.1-0.20240701160413-fd36d05133d0 // indirect
	go.opentelemetry.io/otel v1.27.0 // indirect
	go.opentelemetry.io/otel/exporters/prometheus v0.49.0 // indirect
	go.opentelemetry.io/otel/sdk v1.27.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.27.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240520151616-dc85e6b867a5 // indirect
	google.golang.org/grpc v1.64.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver => ../carbonreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../../internal/common

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/collectd => ../../internal/collectd
