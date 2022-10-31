module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/vcenterreceiver

go 1.18

require (
	github.com/google/go-cmp v0.5.9
	github.com/vmware/govmomi v0.29.0
	go.opentelemetry.io/collector v0.63.2-0.20221031183340-2ed8c0c6ff9c
	go.opentelemetry.io/collector/pdata v0.63.2-0.20221031183340-2ed8c0c6ff9c
	go.uber.org/multierr v1.8.0
	go.uber.org/zap v1.23.0
)

require (
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/rogpeppe/go-internal v1.8.0 // indirect
	github.com/stretchr/testify v1.8.1
)

require gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect

require (
	github.com/basgys/goxml2json v1.1.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/scrapertest v0.63.0
	golang.org/x/sys v0.0.0-20220919091848-fb04ddd9f9c8 // indirect
)

require (
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf v1.4.4 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.opencensus.io v0.23.0 // indirect
	go.opentelemetry.io/otel v1.11.1 // indirect
	go.opentelemetry.io/otel/metric v0.33.0 // indirect
	go.opentelemetry.io/otel/trace v1.11.1 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/text v0.4.0 // indirect
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa // indirect
	google.golang.org/grpc v1.50.1 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/scrapertest => ../../internal/scrapertest
