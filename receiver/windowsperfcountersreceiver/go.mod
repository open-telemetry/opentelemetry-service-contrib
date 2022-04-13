module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowsperfcountersreceiver

go 1.17

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/winperfcounters v0.48.0
	github.com/stretchr/testify v1.7.1
	go.opentelemetry.io/collector v0.48.1-0.20220412005140-8eb68f40028d
	go.uber.org/multierr v1.8.0
	go.uber.org/zap v1.21.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/knadh/koanf v1.4.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/scrapertest v0.48.0
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.opencensus.io v0.23.0 // indirect
	go.opentelemetry.io/collector/pdata v0.0.0-00010101000000-000000000000
	go.opentelemetry.io/otel v1.6.3 // indirect
	go.opentelemetry.io/otel/metric v0.29.0 // indirect
	go.opentelemetry.io/otel/trace v1.6.3 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

require golang.org/x/sys v0.0.0-20220128215802-99c3d69c2c27 // indirect

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/scrapertest => ../../internal/scrapertest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/winperfcounters => ../../pkg/winperfcounters

replace go.opentelemetry.io/collector/pdata => go.opentelemetry.io/collector/pdata v0.0.0-20220412005140-8eb68f40028d
