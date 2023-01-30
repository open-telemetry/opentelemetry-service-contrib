module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver

go 1.18

require (
	github.com/cenkalti/backoff/v4 v4.2.0
	github.com/mongodb-forks/digest v1.0.4
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.70.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest v0.70.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza v0.70.0
	github.com/stretchr/testify v1.8.1
	go.mongodb.org/atlas v0.21.0
	go.opentelemetry.io/collector v0.70.1-0.20230130215412-26bd7b2bf678
	go.opentelemetry.io/collector/component v0.70.1-0.20230130215412-26bd7b2bf678
	go.opentelemetry.io/collector/confmap v0.70.1-0.20230130215412-26bd7b2bf678
	go.opentelemetry.io/collector/consumer v0.70.1-0.20230130215412-26bd7b2bf678
	go.opentelemetry.io/collector/pdata v1.0.0-rc4.0.20230130215412-26bd7b2bf678
	go.uber.org/multierr v1.9.0
	go.uber.org/zap v1.24.0
)

require (
	github.com/antonmedv/expr v1.10.5 // indirect
	github.com/benbjohnson/clock v1.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf v1.5.0 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/observiq/ctimefmt v1.0.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil v0.70.0 // indirect
	github.com/openlyinc/pointy v1.2.0 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/collector/featuregate v0.70.1-0.20230130215412-26bd7b2bf678 // indirect
	go.opentelemetry.io/otel v1.12.0 // indirect
	go.opentelemetry.io/otel/metric v0.35.0 // indirect
	go.opentelemetry.io/otel/trace v1.12.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	gonum.org/v1/gonum v0.12.0 // indirect
	google.golang.org/genproto v0.0.0-20221118155620-16455021b5e6 // indirect
	google.golang.org/grpc v1.52.3 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../../internal/common

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza => ../../pkg/stanza

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage => ../../extension/storage

retract v0.65.0

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../pkg/pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../pkg/pdatautil
