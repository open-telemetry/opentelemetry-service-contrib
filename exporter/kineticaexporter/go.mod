module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kineticaexporter

go 1.21

require (
	github.com/google/uuid v1.6.0
	github.com/kineticadb/kinetica-api-go v0.0.3
	github.com/samber/lo v1.39.0
	github.com/stretchr/testify v1.9.0
	github.com/wk8/go-ordered-map/v2 v2.1.8
	go.opentelemetry.io/collector/component v0.98.1-0.20240416135553-49cc9e05e3a9
	go.opentelemetry.io/collector/config/configopaque v1.5.1-0.20240416135553-49cc9e05e3a9
	go.opentelemetry.io/collector/confmap v0.98.1-0.20240416135553-49cc9e05e3a9
	go.opentelemetry.io/collector/exporter v0.98.1-0.20240416135553-49cc9e05e3a9
	go.opentelemetry.io/collector/pdata v1.5.1-0.20240416135553-49cc9e05e3a9
	go.opentelemetry.io/otel/metric v1.25.0
	go.opentelemetry.io/otel/trace v1.25.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.27.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

require (
	github.com/bahlo/generic-list-go v0.2.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-resty/resty/v2 v2.7.0 // indirect
	github.com/go-viper/mapstructure/v2 v2.0.0-alpha.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/hamba/avro/v2 v2.13.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.1 // indirect
	github.com/logrusorgru/aurora v2.0.3+incompatible // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.1-0.20231216201459-8508981c8b6c // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_golang v1.19.0 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.52.3 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/ztrue/tracerr v0.3.0 // indirect
	go.opentelemetry.io/collector v0.98.1-0.20240416135553-49cc9e05e3a9 // indirect
	go.opentelemetry.io/collector/config/configretry v0.98.1-0.20240416135553-49cc9e05e3a9 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.98.1-0.20240416135553-49cc9e05e3a9 // indirect
	go.opentelemetry.io/collector/consumer v0.98.1-0.20240416135553-49cc9e05e3a9 // indirect
	go.opentelemetry.io/collector/extension v0.98.1-0.20240416135553-49cc9e05e3a9 // indirect
	go.opentelemetry.io/collector/receiver v0.98.1-0.20240416135553-49cc9e05e3a9 // indirect
	go.opentelemetry.io/otel v1.25.0 // indirect
	go.opentelemetry.io/otel/exporters/prometheus v0.47.0 // indirect
	go.opentelemetry.io/otel/sdk v1.25.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.25.0 // indirect
	golang.org/x/exp v0.0.0-20231214170342-aacd6d4b4611 // indirect
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240227224415-6ceb2ff114de // indirect
	google.golang.org/grpc v1.63.2 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)
