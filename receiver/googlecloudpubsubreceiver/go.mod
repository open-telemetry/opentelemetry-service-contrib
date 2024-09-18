module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudpubsubreceiver

go 1.22.0

require (
	cloud.google.com/go/logging v1.11.0
	cloud.google.com/go/pubsub v1.43.0
	github.com/google/go-cmp v0.6.0
	github.com/iancoleman/strcase v0.3.0
	github.com/json-iterator/go v1.1.12
	github.com/stretchr/testify v1.9.0
	go.opentelemetry.io/collector/component v0.109.1-0.20240916171416-d3a171889c53
	go.opentelemetry.io/collector/confmap v1.15.1-0.20240918135221-59c083f830f3
	go.opentelemetry.io/collector/consumer v0.109.1-0.20240916171416-d3a171889c53
	go.opentelemetry.io/collector/consumer/consumertest v0.109.1-0.20240916171416-d3a171889c53
	go.opentelemetry.io/collector/exporter v0.109.1-0.20240916171416-d3a171889c53
	go.opentelemetry.io/collector/pdata v1.15.1-0.20240918193345-a3c0565031b0
	go.opentelemetry.io/collector/receiver v0.109.1-0.20240916171416-d3a171889c53
	go.uber.org/goleak v1.3.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.27.0
	google.golang.org/api v0.197.0
	google.golang.org/genproto v0.0.0-20240903143218-8af14fe29dc1
	google.golang.org/genproto/googleapis/api v0.0.0-20240903143218-8af14fe29dc1
	google.golang.org/grpc v1.66.2
	google.golang.org/protobuf v1.34.2
)

require (
	cloud.google.com/go v0.115.1 // indirect
	cloud.google.com/go/auth v0.9.3 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.4 // indirect
	cloud.google.com/go/compute/metadata v0.5.0 // indirect
	cloud.google.com/go/iam v1.2.0 // indirect
	cloud.google.com/go/longrunning v0.6.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.1.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/s2a-go v0.1.8 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.4 // indirect
	github.com/googleapis/gax-go/v2 v2.13.0 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.1 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.20.3 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.59.1 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	go.einride.tech/aip v0.68.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/collector/config/configretry v1.15.1-0.20240918193345-a3c0565031b0 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.109.1-0.20240916171416-d3a171889c53 // indirect
	go.opentelemetry.io/collector/consumer/consumerprofiles v0.109.1-0.20240916171416-d3a171889c53 // indirect
	go.opentelemetry.io/collector/extension v0.109.1-0.20240916171416-d3a171889c53 // indirect
	go.opentelemetry.io/collector/extension/experimental/storage v0.109.1-0.20240916171416-d3a171889c53 // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.109.1-0.20240916171416-d3a171889c53 // indirect
	go.opentelemetry.io/collector/receiver/receiverprofiles v0.109.1-0.20240916171416-d3a171889c53 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.54.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.54.0 // indirect
	go.opentelemetry.io/otel v1.30.0 // indirect
	go.opentelemetry.io/otel/exporters/prometheus v0.52.0 // indirect
	go.opentelemetry.io/otel/metric v1.30.0 // indirect
	go.opentelemetry.io/otel/sdk v1.30.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.30.0 // indirect
	go.opentelemetry.io/otel/trace v1.30.0 // indirect
	golang.org/x/crypto v0.27.0 // indirect
	golang.org/x/net v0.29.0 // indirect
	golang.org/x/oauth2 v0.23.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/text v0.18.0 // indirect
	golang.org/x/time v0.6.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240903143218-8af14fe29dc1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)
