module github.com/open-telemetry/opentelemetry-collector-contrib/pkg/datadog

go 1.23.0

require (
	github.com/DataDog/datadog-agent/pkg/util/hostname/validate v0.62.2
	github.com/DataDog/opentelemetry-mapping-go/pkg/otlp/attributes v0.25.0
	github.com/DataDog/opentelemetry-mapping-go/pkg/otlp/metrics v0.25.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog v0.119.0
	github.com/stretchr/testify v1.10.0
	go.opentelemetry.io/collector/component v0.119.1-0.20250210123122-44b3eeda354c
	go.opentelemetry.io/collector/component/componenttest v0.119.1-0.20250210123122-44b3eeda354c
	go.opentelemetry.io/collector/config/configauth v0.119.1-0.20250210123122-44b3eeda354c
	go.opentelemetry.io/collector/config/confighttp v0.119.1-0.20250210123122-44b3eeda354c
	go.opentelemetry.io/collector/config/confignet v1.25.1-0.20250210123122-44b3eeda354c
	go.opentelemetry.io/collector/config/configopaque v1.25.1-0.20250210123122-44b3eeda354c
	go.opentelemetry.io/collector/config/configretry v1.25.1-0.20250210123122-44b3eeda354c
	go.opentelemetry.io/collector/config/configtls v1.25.1-0.20250210123122-44b3eeda354c
	go.opentelemetry.io/collector/confmap v1.25.1-0.20250210123122-44b3eeda354c
	go.opentelemetry.io/collector/confmap/xconfmap v0.0.0-20250210155359-76f44e1e21d1
	go.opentelemetry.io/collector/exporter v0.119.1-0.20250210123122-44b3eeda354c
	go.opentelemetry.io/collector/featuregate v1.25.1-0.20250210123122-44b3eeda354c
	go.uber.org/zap v1.27.0
)

require (
	cloud.google.com/go/compute/metadata v0.6.0 // indirect
	github.com/DataDog/datadog-agent/pkg/proto v0.64.0-devel // indirect
	github.com/DataDog/datadog-agent/pkg/util/log v0.62.2 // indirect
	github.com/DataDog/datadog-agent/pkg/util/scrubber v0.62.2 // indirect
	github.com/DataDog/datadog-agent/pkg/version v0.62.2 // indirect
	github.com/DataDog/datadog-api-client-go/v2 v2.35.0 // indirect
	github.com/DataDog/gohai v0.0.0-20230524154621-4316413895ee // indirect
	github.com/DataDog/opentelemetry-mapping-go/pkg/inframetadata v0.25.0 // indirect
	github.com/DataDog/opentelemetry-mapping-go/pkg/quantile v0.25.0 // indirect
	github.com/DataDog/sketches-go v1.4.6 // indirect
	github.com/DataDog/zstd v1.5.2 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp v1.26.0 // indirect
	github.com/aws/aws-sdk-go-v2 v1.36.1 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.29.6 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.17.59 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.28 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.32 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.32 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.203.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.12.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.12.13 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.24.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.28.14 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.33.14 // indirect
	github.com/aws/smithy-go v1.22.2 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/emicklei/go-restful/v3 v3.11.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/fxamacker/cbor/v2 v2.7.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/go-openapi/jsonpointer v0.21.0 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/swag v0.23.0 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/gnostic-models v0.6.8 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.2 // indirect
	github.com/lufia/plan9stats v0.0.0-20240226150601-1dcf7310316a // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/ecsutil v0.119.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.119.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig v0.119.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders v0.119.0 // indirect
	github.com/openshift/api v3.9.0+incompatible // indirect
	github.com/openshift/client-go v0.0.0-20210521082421-73d9475a9142 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/philhofer/fwd v1.1.3-0.20240916144458-20a13a1f6b7c // indirect
	github.com/pierrec/lz4/v4 v4.1.22 // indirect
	github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/power-devops/perfstat v0.0.0-20240221224432-82ca36839d55 // indirect
	github.com/rs/cors v1.11.1 // indirect
	github.com/shirou/gopsutil/v3 v3.22.12 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/tinylib/msgp v1.2.5 // indirect
	github.com/tklauser/go-sysconf v0.3.14 // indirect
	github.com/tklauser/numcpus v0.8.0 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/collector/client v1.25.1-0.20250210123122-44b3eeda354c // indirect
	go.opentelemetry.io/collector/config/configcompression v1.25.1-0.20250210123122-44b3eeda354c // indirect
	go.opentelemetry.io/collector/consumer v1.25.1-0.20250210123122-44b3eeda354c // indirect
	go.opentelemetry.io/collector/consumer/consumererror v0.119.1-0.20250210123122-44b3eeda354c // indirect
	go.opentelemetry.io/collector/extension v0.119.1-0.20250210123122-44b3eeda354c // indirect
	go.opentelemetry.io/collector/extension/auth v0.119.1-0.20250210123122-44b3eeda354c // indirect
	go.opentelemetry.io/collector/extension/xextension v0.119.1-0.20250210123122-44b3eeda354c // indirect
	go.opentelemetry.io/collector/pdata v1.25.1-0.20250210123122-44b3eeda354c // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.119.1-0.20250210123122-44b3eeda354c // indirect
	go.opentelemetry.io/collector/pipeline v0.119.1-0.20250210123122-44b3eeda354c // indirect
	go.opentelemetry.io/collector/semconv v0.119.1-0.20250210123122-44b3eeda354c // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.59.0 // indirect
	go.opentelemetry.io/otel v1.34.0 // indirect
	go.opentelemetry.io/otel/metric v1.34.0 // indirect
	go.opentelemetry.io/otel/sdk v1.34.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.34.0 // indirect
	go.opentelemetry.io/otel/trace v1.34.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/exp v0.0.0-20241217172543-b2144cdd0a67 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/oauth2 v0.24.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/term v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	golang.org/x/time v0.7.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241230172942-26aa7a208def // indirect
	google.golang.org/grpc v1.70.0 // indirect
	google.golang.org/protobuf v1.36.5 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gopkg.in/zorkian/go-datadog-api.v2 v2.30.0 // indirect
	k8s.io/api v0.31.3 // indirect
	k8s.io/apimachinery v0.31.3 // indirect
	k8s.io/client-go v0.31.3 // indirect
	k8s.io/klog/v2 v2.130.1 // indirect
	k8s.io/kube-openapi v0.0.0-20241105132330-32ad38e42d3f // indirect
	k8s.io/utils v0.0.0-20241104100929-3ea5e8cea738 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.5.0 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog => ../../internal/datadog

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders => ../../internal/metadataproviders

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig => ../../internal/k8sconfig

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../../internal/common

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/ecsutil => ../../internal/aws/ecsutil
