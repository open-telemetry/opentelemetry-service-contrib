module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlecloudexporter

go 1.21

require (
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/collector v0.45.0
	github.com/stretchr/testify v1.8.4
	go.opentelemetry.io/collector/component v0.94.2-0.20240216233521-0ab8f4443937
	go.opentelemetry.io/collector/confmap v0.94.2-0.20240216233521-0ab8f4443937
	go.opentelemetry.io/collector/exporter v0.94.2-0.20240216233521-0ab8f4443937
	go.opentelemetry.io/collector/featuregate v1.2.0
	go.opentelemetry.io/otel/metric v1.23.1
	go.opentelemetry.io/otel/trace v1.23.1
	go.uber.org/goleak v1.3.0
)

require (
	cloud.google.com/go v0.111.0 // indirect
	cloud.google.com/go/compute v1.23.3 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	cloud.google.com/go/logging v1.8.1 // indirect
	cloud.google.com/go/longrunning v0.5.4 // indirect
	cloud.google.com/go/monitoring v1.16.3 // indirect
	cloud.google.com/go/trace v1.10.4 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace v1.21.0 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/internal/resourcemapping v0.45.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.0.0-alpha.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/s2a-go v0.1.7 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.2 // indirect
	github.com/googleapis/gax-go/v2 v2.12.0 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.1-0.20231216201459-8508981c8b6c // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.18.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.46.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/tidwall/gjson v1.10.2 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/tidwall/tinylru v1.1.0 // indirect
	github.com/tidwall/wal v1.1.7 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/collector v0.94.2-0.20240216233521-0ab8f4443937 // indirect
	go.opentelemetry.io/collector/config/configretry v0.94.2-0.20240216233521-0ab8f4443937 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.94.2-0.20240216233521-0ab8f4443937 // indirect
	go.opentelemetry.io/collector/consumer v0.94.2-0.20240216233521-0ab8f4443937 // indirect
	go.opentelemetry.io/collector/extension v0.94.2-0.20240216233521-0ab8f4443937 // indirect
	go.opentelemetry.io/collector/pdata v1.2.0 // indirect
	go.opentelemetry.io/collector/receiver v0.94.2-0.20240216233521-0ab8f4443937 // indirect
	go.opentelemetry.io/collector/semconv v0.94.2-0.20240216233521-0ab8f4443937 // indirect
	go.opentelemetry.io/otel v1.23.1 // indirect
	go.opentelemetry.io/otel/exporters/prometheus v0.45.2 // indirect
	go.opentelemetry.io/otel/sdk v1.23.1 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.23.1 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	golang.org/x/crypto v0.18.0 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/oauth2 v0.16.0 // indirect
	golang.org/x/sync v0.5.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	google.golang.org/api v0.150.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto v0.0.0-20231212172506-995d672761c0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240102182953-50ed04b92917 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240102182953-50ed04b92917 // indirect
	google.golang.org/grpc v1.61.0 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
	v0.88.0
	v0.87.0
	v0.86.0
	v0.85.0
	v0.84.0
	v0.76.2
	v0.76.1
	v0.65.0
)
