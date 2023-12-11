module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/splunkhecreceiver

go 1.20

require (
	github.com/gorilla/mux v1.8.1
	github.com/json-iterator/go v1.1.12
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/splunkhecexporter v0.90.1
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.90.1
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/sharedcomponent v0.90.1
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk v0.90.1
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest v0.90.1
	github.com/stretchr/testify v1.8.4
	go.opentelemetry.io/collector/component v0.90.2-0.20231208183206-eed3b4e9c5ef
	go.opentelemetry.io/collector/config/confighttp v0.90.2-0.20231208183206-eed3b4e9c5ef
	go.opentelemetry.io/collector/config/configtls v0.90.2-0.20231208183206-eed3b4e9c5ef
	go.opentelemetry.io/collector/confmap v0.90.2-0.20231208183206-eed3b4e9c5ef
	go.opentelemetry.io/collector/consumer v0.90.2-0.20231208183206-eed3b4e9c5ef
	go.opentelemetry.io/collector/exporter v0.90.2-0.20231208183206-eed3b4e9c5ef
	go.opentelemetry.io/collector/pdata v1.0.1-0.20231211180437-7ec38e5c1992
	go.opentelemetry.io/collector/receiver v0.90.2-0.20231208183206-eed3b4e9c5ef
	go.opentelemetry.io/collector/semconv v0.90.2-0.20231208183206-eed3b4e9c5ef
	go.uber.org/zap v1.26.0
)

require (
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/klauspost/compress v1.17.4 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.0.1 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.1-0.20220423185008-bf980b35cac4 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.90.1 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchperresourceattr v0.90.1 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil v0.90.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rs/cors v1.10.1 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/collector v0.90.2-0.20231208183206-eed3b4e9c5ef // indirect
	go.opentelemetry.io/collector/config/configauth v0.90.2-0.20231208183206-eed3b4e9c5ef // indirect
	go.opentelemetry.io/collector/config/configcompression v0.90.2-0.20231208183206-eed3b4e9c5ef // indirect
	go.opentelemetry.io/collector/config/configopaque v0.90.2-0.20231208183206-eed3b4e9c5ef // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.90.2-0.20231208183206-eed3b4e9c5ef // indirect
	go.opentelemetry.io/collector/config/internal v0.90.2-0.20231208183206-eed3b4e9c5ef // indirect
	go.opentelemetry.io/collector/extension v0.90.2-0.20231208183206-eed3b4e9c5ef // indirect
	go.opentelemetry.io/collector/extension/auth v0.90.2-0.20231208183206-eed3b4e9c5ef // indirect
	go.opentelemetry.io/collector/featuregate v1.0.1-0.20231211180437-7ec38e5c1992 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.46.1 // indirect
	go.opentelemetry.io/otel v1.21.0 // indirect
	go.opentelemetry.io/otel/metric v1.21.0 // indirect
	go.opentelemetry.io/otel/trace v1.21.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d // indirect
	google.golang.org/grpc v1.59.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/splunkhecexporter => ../../exporter/splunkhecexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk => ../../internal/splunk

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../../internal/common

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../pkg/pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../pkg/pdatautil

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchperresourceattr => ../../pkg/batchperresourceattr

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/sharedcomponent => ../../internal/sharedcomponent

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden => ../../pkg/golden
