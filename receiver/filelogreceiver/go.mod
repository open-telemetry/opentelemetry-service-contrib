module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/filelogreceiver

go 1.22.0

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza v0.108.0
	github.com/stretchr/testify v1.9.0
	go.opentelemetry.io/collector/component v0.109.0
	go.opentelemetry.io/collector/confmap v1.15.0
	go.opentelemetry.io/collector/consumer v0.109.0 // indirect
	go.opentelemetry.io/collector/pdata v1.15.0
	go.opentelemetry.io/collector/receiver v0.109.0
	go.opentelemetry.io/collector/semconv v0.109.0 // indirect
	go.opentelemetry.io/otel/metric v1.29.0 // indirect
	go.opentelemetry.io/otel/trace v1.29.0 // indirect
	go.uber.org/goleak v1.3.0
)

require go.opentelemetry.io/collector/consumer/consumertest v0.109.0

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bmatcuk/doublestar/v4 v4.6.1 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/elastic/lunes v0.1.0 // indirect
	github.com/expr-lang/expr v1.16.9 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.1.0 // indirect
	github.com/goccy/go-json v0.10.3 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/jonboulle/clockwork v0.4.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.1 // indirect
	github.com/leodido/go-syslog/v4 v4.1.0 // indirect
	github.com/leodido/ragel-machinery v0.0.0-20190525184631-5f46317e436b // indirect
	github.com/magefile/mage v1.15.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.20.2 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.57.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/valyala/fastjson v1.6.4 // indirect
	go.opentelemetry.io/collector v0.109.0 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.109.0 // indirect
	go.opentelemetry.io/collector/consumer/consumerprofiles v0.109.0 // indirect
	go.opentelemetry.io/collector/extension v0.109.0 // indirect
	go.opentelemetry.io/collector/extension/experimental/storage v0.109.0 // indirect
	go.opentelemetry.io/collector/featuregate v1.15.0 // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.109.0 // indirect
	go.opentelemetry.io/collector/receiver/receiverprofiles v0.109.0 // indirect
	go.opentelemetry.io/otel v1.29.0 // indirect
	go.opentelemetry.io/otel/exporters/prometheus v0.51.0 // indirect
	go.opentelemetry.io/otel/sdk v1.29.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.29.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	gonum.org/v1/gonum v0.15.1 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240822170219-fc7c04adadcd // indirect
	google.golang.org/grpc v1.66.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage => ../../extension/storage

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza => ../../pkg/stanza

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../pkg/pdatautil

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../pkg/pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden => ../../pkg/golden

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../../internal/common
