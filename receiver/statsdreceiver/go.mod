module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver

go 1.18

require (
	// Note that this Lightstep metrics SDK dependency is temporary.  The same
	// code is under review upstream, see
	// https://github.com/open-telemetry/opentelemetry-go/pull/3022
	github.com/lightstep/otel-launcher-go/lightstep/sdk/metric v1.11.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.59.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.58.0
	github.com/stretchr/testify v1.8.0
	go.opencensus.io v0.23.0
	go.opentelemetry.io/collector v0.59.1-0.20220909192754-8d66f408a79a
	go.opentelemetry.io/collector/pdata v0.59.1-0.20220909192754-8d66f408a79a
	go.opentelemetry.io/otel v1.10.0

	// Note! This unreleased otel-go dependency (would-be v0.32)
	// changes histogram inclusivity.
	// https://github.com/open-telemetry/opentelemetry-go/pull/2982.
	go.opentelemetry.io/otel/sdk/metric v0.31.1-0.20220826135333-55b49c407e07
	go.uber.org/multierr v1.8.0
	go.uber.org/zap v1.23.0
	gonum.org/v1/gonum v0.12.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-farm v0.0.0-20200201041132-a6ae2369ad13 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf v1.4.3 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.opentelemetry.io/otel/metric v0.31.0 // indirect
	go.opentelemetry.io/otel/sdk v1.10.0 // indirect
	go.opentelemetry.io/otel/trace v1.10.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	golang.org/x/exp v0.0.0-20200224162631-6cc2880d07d6 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sys v0.0.0-20220808155132-1c4a2a72c664 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220112215332-a9c7c0acf9f2 // indirect
	google.golang.org/grpc v1.49.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../../internal/common

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal
