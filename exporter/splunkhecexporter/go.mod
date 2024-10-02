module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/splunkhecexporter

go 1.22.0

require (
	github.com/cenkalti/backoff/v4 v4.3.0
	github.com/goccy/go-json v0.10.3
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.110.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk v0.110.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchperresourceattr v0.110.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden v0.110.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest v0.110.0
	github.com/stretchr/testify v1.9.0
	github.com/testcontainers/testcontainers-go v0.31.0
	go.opentelemetry.io/collector/component v0.110.1-0.20240927195042-40396d5fc50c
	go.opentelemetry.io/collector/config/confighttp v0.110.1-0.20240927195042-40396d5fc50c
	go.opentelemetry.io/collector/config/configopaque v1.16.1-0.20240927195042-40396d5fc50c
	go.opentelemetry.io/collector/config/configretry v1.16.1-0.20240927195042-40396d5fc50c
	go.opentelemetry.io/collector/config/configtls v1.16.1-0.20240927195042-40396d5fc50c
	go.opentelemetry.io/collector/confmap v1.16.1-0.20240927195042-40396d5fc50c
	go.opentelemetry.io/collector/consumer v0.110.1-0.20240927195042-40396d5fc50c
	go.opentelemetry.io/collector/consumer/consumertest v0.110.1-0.20240927195042-40396d5fc50c
	go.opentelemetry.io/collector/exporter v0.110.1-0.20240927195042-40396d5fc50c
	go.opentelemetry.io/collector/pdata v1.16.1-0.20240927195042-40396d5fc50c
	go.opentelemetry.io/collector/semconv v0.110.1-0.20240927195042-40396d5fc50c
	go.opentelemetry.io/otel v1.30.0
	go.opentelemetry.io/otel/metric v1.30.0
	go.opentelemetry.io/otel/sdk/metric v1.30.0
	go.opentelemetry.io/otel/trace v1.30.0
	go.uber.org/goleak v1.3.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.27.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	dario.cat/mergo v1.0.0 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/Microsoft/hcsshim v0.11.4 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/containerd/containerd v1.7.15 // indirect
	github.com/containerd/log v0.1.0 // indirect
	github.com/cpuguy83/dockercfg v0.3.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/distribution/reference v0.5.0 // indirect
	github.com/docker/docker v26.1.5+incompatible // indirect
	github.com/docker/go-connections v0.5.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-viper/mapstructure/v2 v2.1.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.1 // indirect
	github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/moby/docker-image-spec v1.3.1 // indirect
	github.com/moby/patternmatcher v0.6.0 // indirect
	github.com/moby/sys/sequential v0.5.0 // indirect
	github.com/moby/sys/user v0.1.0 // indirect
	github.com/moby/term v0.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil v0.110.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/rs/cors v1.11.1 // indirect
	github.com/shirou/gopsutil/v3 v3.24.5 // indirect
	github.com/shoenig/go-m1cpu v0.1.6 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.opentelemetry.io/collector/client v1.16.1-0.20240927195042-40396d5fc50c // indirect
	go.opentelemetry.io/collector/component/componentprofiles v0.110.1-0.20240927195042-40396d5fc50c // indirect
	go.opentelemetry.io/collector/config/configauth v0.110.1-0.20240927195042-40396d5fc50c // indirect
	go.opentelemetry.io/collector/config/configcompression v1.16.1-0.20240927195042-40396d5fc50c // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.110.1-0.20240927195042-40396d5fc50c // indirect
	go.opentelemetry.io/collector/config/internal v0.110.1-0.20240927195042-40396d5fc50c // indirect
	go.opentelemetry.io/collector/consumer/consumerprofiles v0.110.1-0.20240927195042-40396d5fc50c // indirect
	go.opentelemetry.io/collector/exporter/exporterprofiles v0.110.1-0.20240927195042-40396d5fc50c // indirect
	go.opentelemetry.io/collector/extension v0.110.1-0.20240927195042-40396d5fc50c // indirect
	go.opentelemetry.io/collector/extension/auth v0.110.1-0.20240927195042-40396d5fc50c // indirect
	go.opentelemetry.io/collector/extension/experimental/storage v0.110.1-0.20240927195042-40396d5fc50c // indirect
	go.opentelemetry.io/collector/internal/globalsignal v0.110.1-0.20240927195042-40396d5fc50c // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.110.1-0.20240927195042-40396d5fc50c // indirect
	go.opentelemetry.io/collector/pipeline v0.110.1-0.20240927195042-40396d5fc50c // indirect
	go.opentelemetry.io/collector/receiver v0.110.1-0.20240927195042-40396d5fc50c // indirect
	go.opentelemetry.io/collector/receiver/receiverprofiles v0.110.1-0.20240927195042-40396d5fc50c // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.55.0 // indirect
	go.opentelemetry.io/otel/sdk v1.30.0 // indirect
	golang.org/x/crypto v0.27.0 // indirect
	golang.org/x/mod v0.19.0 // indirect
	golang.org/x/net v0.29.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/text v0.18.0 // indirect
	golang.org/x/tools v0.23.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240822170219-fc7c04adadcd // indirect
	google.golang.org/grpc v1.66.2 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk => ../../internal/splunk

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchperresourceattr => ../../pkg/batchperresourceattr

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../pkg/pdatautil

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../pkg/pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden => ../../pkg/golden
