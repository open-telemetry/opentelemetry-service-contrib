module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/saphanareceiver

go 1.17

require (
	github.com/SAP/go-hdb v0.105.5
	go.opentelemetry.io/collector v0.49.1-0.20220428142054-b34df3a7e7e9
	go.opentelemetry.io/collector/model v0.49.1-0.20220428142054-b34df3a7e7e9
	go.uber.org/zap v1.21.0 // indirect
)

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/scrapertest v0.49.0
	go.opentelemetry.io/collector/pdata v0.50.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/knadh/koanf v1.4.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.6.1 // indirect
	github.com/stretchr/objx v0.2.0 // indirect
	go.opencensus.io v0.23.0 // indirect
	go.opentelemetry.io/otel v1.6.3 // indirect
	go.opentelemetry.io/otel/metric v0.29.0 // indirect
	go.opentelemetry.io/otel/trace v1.6.3 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sys v0.0.0-20220128215802-99c3d69c2c27 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa // indirect
	google.golang.org/grpc v1.46.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

require (
	github.com/stretchr/testify v1.7.1
	go.uber.org/multierr v1.8.0
	golang.org/x/crypto v0.0.0-20220307211146-efcb8507fb70 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/scrapertest => ../../internal/scrapertest

replace go.opentelemetry.io/collector/semconv => go.opentelemetry.io/collector/semconv v0.0.0-20220422001137-87ab5de64ce4
