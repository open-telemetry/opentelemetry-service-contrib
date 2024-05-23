module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/signalfxreceiver

go 1.21.0

require (
	github.com/gorilla/mux v1.8.1
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter v0.101.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.101.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk v0.101.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest v0.101.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/signalfx v0.101.0
	github.com/signalfx/com_signalfx_metrics_protobuf v0.0.3
	github.com/stretchr/testify v1.9.0
	go.opentelemetry.io/collector/component v0.101.1-0.20240523155058-812210ba3685
	go.opentelemetry.io/collector/config/confighttp v0.101.1-0.20240523155058-812210ba3685
	go.opentelemetry.io/collector/config/configtls v0.101.1-0.20240523155058-812210ba3685
	go.opentelemetry.io/collector/confmap v0.101.1-0.20240523155058-812210ba3685
	go.opentelemetry.io/collector/consumer v0.101.1-0.20240523155058-812210ba3685
	go.opentelemetry.io/collector/exporter v0.101.1-0.20240523155058-812210ba3685
	go.opentelemetry.io/collector/pdata v1.8.1-0.20240523143024-6f5d43f9e405
	go.opentelemetry.io/collector/receiver v0.101.1-0.20240523155058-812210ba3685
	go.opentelemetry.io/otel/metric v1.26.0
	go.opentelemetry.io/otel/trace v1.26.0
	go.uber.org/zap v1.27.0
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-viper/mapstructure/v2 v2.0.0-alpha.1 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.8 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.1 // indirect
	github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.101.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchperresourceattr v0.101.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata v0.101.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil v0.101.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/prometheus/client_golang v1.19.1 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.53.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/rs/cors v1.10.1 // indirect
	github.com/shirou/gopsutil/v3 v3.24.4 // indirect
	github.com/shoenig/go-m1cpu v0.1.6 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.opentelemetry.io/collector v0.101.1-0.20240523155058-812210ba3685 // indirect
	go.opentelemetry.io/collector/config/configauth v0.101.1-0.20240523155058-812210ba3685 // indirect
	go.opentelemetry.io/collector/config/configcompression v1.8.1-0.20240523143024-6f5d43f9e405 // indirect
	go.opentelemetry.io/collector/config/configopaque v1.8.1-0.20240523143024-6f5d43f9e405 // indirect
	go.opentelemetry.io/collector/config/configretry v0.101.1-0.20240523155058-812210ba3685 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.101.1-0.20240523155058-812210ba3685 // indirect
	go.opentelemetry.io/collector/config/internal v0.101.1-0.20240523155058-812210ba3685 // indirect
	go.opentelemetry.io/collector/extension v0.101.1-0.20240523155058-812210ba3685 // indirect
	go.opentelemetry.io/collector/extension/auth v0.101.1-0.20240523155058-812210ba3685 // indirect
	go.opentelemetry.io/collector/featuregate v1.8.1-0.20240523143024-6f5d43f9e405 // indirect
	go.opentelemetry.io/collector/semconv v0.101.1-0.20240523155058-812210ba3685 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.51.0 // indirect
	go.opentelemetry.io/otel v1.26.0 // indirect
	go.opentelemetry.io/otel/exporters/prometheus v0.48.0 // indirect
	go.opentelemetry.io/otel/sdk v1.26.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.26.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240401170217-c3f982113cda // indirect
	google.golang.org/grpc v1.63.2 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter => ../../exporter/signalfxexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../pkg/pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../pkg/pdatautil

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../../internal/common

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk => ../../internal/splunk

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata => ../../pkg/experimentalmetricmetadata

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchperresourceattr => ../../pkg/batchperresourceattr

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/signalfx => ../../pkg/translator/signalfx

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)

// https://github.com/go-openapi/spec/issues/156
replace github.com/go-openapi/spec v0.20.5 => github.com/go-openapi/spec v0.20.6

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden => ../../pkg/golden
