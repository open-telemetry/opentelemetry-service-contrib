module github.com/open-telemetry/opentelemetry-collector-contrib/testbed

go 1.21.0

require (
	github.com/fluent/fluent-logger-golang v1.9.0
	github.com/jaegertracing/jaeger v1.60.0
	github.com/open-telemetry/opentelemetry-collector-contrib/connector/routingconnector v0.107.0
	github.com/open-telemetry/opentelemetry-collector-contrib/connector/spanmetricsconnector v0.107.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/carbonexporter v0.107.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/opencensusexporter v0.107.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusexporter v0.107.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sapmexporter v0.107.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter v0.107.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/syslogexporter v0.107.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/zipkinexporter v0.107.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.107.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.107.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk v0.107.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden v0.107.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza v0.107.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/jaeger v0.107.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver v0.107.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/datadogreceiver v0.107.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jaegerreceiver v0.107.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/opencensusreceiver v0.107.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver v0.107.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sapmreceiver v0.107.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/signalfxreceiver v0.107.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/splunkhecreceiver v0.107.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/syslogreceiver v0.107.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zipkinreceiver v0.107.0
	github.com/open-telemetry/opentelemetry-collector-contrib/testbed/mockdatasenders/mockdatadogagentexporter v0.107.0
	github.com/prometheus/common v0.55.0
	github.com/prometheus/prometheus v0.53.1
	github.com/shirou/gopsutil/v4 v4.24.7
	github.com/stretchr/testify v1.9.0
	go.opentelemetry.io/collector/component v0.107.0
	go.opentelemetry.io/collector/config/configcompression v1.13.0
	go.opentelemetry.io/collector/config/configgrpc v0.107.0
	go.opentelemetry.io/collector/config/confighttp v0.107.0
	go.opentelemetry.io/collector/config/confignet v0.107.0
	go.opentelemetry.io/collector/config/configretry v1.13.0
	go.opentelemetry.io/collector/config/configtls v1.13.0
	go.opentelemetry.io/collector/confmap v0.107.0
	go.opentelemetry.io/collector/confmap/provider/fileprovider v0.107.0
	go.opentelemetry.io/collector/connector v0.107.0
	go.opentelemetry.io/collector/consumer v0.107.0
	go.opentelemetry.io/collector/exporter v0.107.0
	go.opentelemetry.io/collector/exporter/debugexporter v0.107.0
	go.opentelemetry.io/collector/exporter/otlpexporter v0.107.0
	go.opentelemetry.io/collector/exporter/otlphttpexporter v0.107.0
	go.opentelemetry.io/collector/extension v0.107.0
	go.opentelemetry.io/collector/extension/ballastextension v0.107.0
	go.opentelemetry.io/collector/extension/zpagesextension v0.107.0
	go.opentelemetry.io/collector/otelcol v0.107.0
	go.opentelemetry.io/collector/pdata v1.13.0
	go.opentelemetry.io/collector/processor v0.107.0
	go.opentelemetry.io/collector/processor/batchprocessor v0.107.0
	go.opentelemetry.io/collector/processor/memorylimiterprocessor v0.107.0
	go.opentelemetry.io/collector/receiver v0.107.0
	go.opentelemetry.io/collector/receiver/otlpreceiver v0.107.0
	go.opentelemetry.io/collector/semconv v0.107.0
	go.uber.org/goleak v1.3.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.27.0
	golang.org/x/text v0.17.0
	google.golang.org/grpc v1.65.0
)

require (
	cloud.google.com/go/auth v0.5.1 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.2 // indirect
	cloud.google.com/go/compute/metadata v0.3.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.11.1 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.6.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.8.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5 v5.7.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4 v4.3.0 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v1.2.2 // indirect
	github.com/Code-Hex/go-generics-cache v1.5.1 // indirect
	github.com/DataDog/agent-payload/v5 v5.0.130 // indirect
	github.com/DataDog/datadog-agent/pkg/proto v0.56.0-rc.11 // indirect
	github.com/DataDog/datadog-agent/pkg/trace/exportable v0.0.0-20201016145401-4646cf596b02 // indirect
	github.com/DataDog/datadog-api-client-go/v2 v2.29.0 // indirect
	github.com/DataDog/zstd v1.5.2 // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/alecthomas/participle/v2 v2.1.1 // indirect
	github.com/alecthomas/units v0.0.0-20231202071711-9a357b53e9c9 // indirect
	github.com/apache/thrift v0.20.0 // indirect
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/aws/aws-sdk-go v1.53.16 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/census-instrumentation/opencensus-proto v0.4.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cncf/xds/go v0.0.0-20240423153145-555b57ec207b // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dennwc/varint v1.0.0 // indirect
	github.com/digitalocean/godo v1.117.0 // indirect
	github.com/distribution/reference v0.5.0 // indirect
	github.com/docker/docker v26.1.5+incompatible // indirect
	github.com/docker/go-connections v0.5.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/elastic/go-grok v0.3.1 // indirect
	github.com/emicklei/go-restful/v3 v3.11.0 // indirect
	github.com/envoyproxy/go-control-plane v0.12.0 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.0.4 // indirect
	github.com/expr-lang/expr v1.16.9 // indirect
	github.com/fatih/color v1.16.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-openapi/jsonpointer v0.20.2 // indirect
	github.com/go-openapi/jsonreference v0.20.4 // indirect
	github.com/go-openapi/swag v0.22.9 // indirect
	github.com/go-resty/resty/v2 v2.13.1 // indirect
	github.com/go-viper/mapstructure/v2 v2.0.0 // indirect
	github.com/go-zookeeper/zk v1.0.3 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/goccy/go-json v0.10.3 // indirect
	github.com/gogo/googleapis v1.4.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/gnostic-models v0.6.8 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/s2a-go v0.1.7 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.2 // indirect
	github.com/googleapis/gax-go/v2 v2.12.4 // indirect
	github.com/gophercloud/gophercloud v1.12.0 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/grafana/regexp v0.0.0-20240518133315-a468a5bfb3bc // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.21.0 // indirect
	github.com/hashicorp/consul/api v1.29.1 // indirect
	github.com/hashicorp/cronexpr v1.1.2 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.6.3 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.7 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/hashicorp/nomad/api v0.0.0-20240604134157-e73d8bb1140d // indirect
	github.com/hashicorp/serf v0.10.1 // indirect
	github.com/hetznercloud/hcloud-go/v2 v2.9.0 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/ionos-cloud/sdk-go/v6 v6.1.11 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/jonboulle/clockwork v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.1 // indirect
	github.com/kolo/xmlrpc v0.0.0-20220921171641-a4b6fa1dd06b // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/leodido/go-syslog/v4 v4.1.0 // indirect
	github.com/leodido/ragel-machinery v0.0.0-20190525184631-5f46317e436b // indirect
	github.com/lightstep/go-expohisto v1.0.0 // indirect
	github.com/linode/linodego v1.35.0 // indirect
	github.com/lufia/plan9stats v0.0.0-20220913051719-115f729f3c8c // indirect
	github.com/magefile/mage v1.15.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/miekg/dns v1.1.59 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.1-0.20231216201459-8508981c8b6c // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/moby/docker-image-spec v1.3.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mostynb/go-grpc-compression v1.2.3 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/mwitkow/go-conntrack v0.0.0-20190716064945-2f068394615f // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/ackextension v0.107.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/exp/metrics v0.107.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/sharedcomponent v0.107.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchperresourceattr v0.107.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata v0.107.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl v0.107.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil v0.107.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/resourcetotelemetry v0.107.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/opencensus v0.107.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheus v0.107.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/signalfx v0.107.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/zipkin v0.107.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.0 // indirect
	github.com/openzipkin/zipkin-go v0.4.3 // indirect
	github.com/ovh/go-ovh v1.5.1 // indirect
	github.com/philhofer/fwd v1.1.3-0.20240612014219-fbbf4953d986 // indirect
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/power-devops/perfstat v0.0.0-20220216144756-c35f1ee13d7c // indirect
	github.com/prometheus/client_golang v1.19.1 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common/sigv4 v0.1.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/rs/cors v1.11.0 // indirect
	github.com/scaleway/scaleway-sdk-go v1.0.0-beta.27 // indirect
	github.com/shoenig/go-m1cpu v0.1.6 // indirect
	github.com/signalfx/com_signalfx_metrics_protobuf v0.0.3 // indirect
	github.com/signalfx/sapm-proto v0.14.0 // indirect
	github.com/soheilhy/cmux v0.1.5 // indirect
	github.com/spf13/cobra v1.8.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tinylib/msgp v1.2.0 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/valyala/fastjson v1.6.4 // indirect
	github.com/vultr/govultr/v2 v2.17.2 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/collector v0.107.0 // indirect
	go.opentelemetry.io/collector/client v1.13.0 // indirect
	go.opentelemetry.io/collector/component/componentprofiles v0.107.0 // indirect
	go.opentelemetry.io/collector/config/configauth v0.107.0 // indirect
	go.opentelemetry.io/collector/config/configopaque v1.13.0 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.107.0 // indirect
	go.opentelemetry.io/collector/config/internal v0.107.0 // indirect
	go.opentelemetry.io/collector/consumer/consumerprofiles v0.107.0 // indirect
	go.opentelemetry.io/collector/consumer/consumertest v0.107.0 // indirect
	go.opentelemetry.io/collector/extension/auth v0.107.0 // indirect
	go.opentelemetry.io/collector/featuregate v1.13.0 // indirect
	go.opentelemetry.io/collector/internal/globalgates v0.107.0 // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.107.0 // indirect
	go.opentelemetry.io/collector/service v0.107.0 // indirect
	go.opentelemetry.io/contrib/config v0.8.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.53.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.53.0 // indirect
	go.opentelemetry.io/contrib/propagators/b3 v1.28.0 // indirect
	go.opentelemetry.io/contrib/zpages v0.53.0 // indirect
	go.opentelemetry.io/otel v1.28.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp v0.4.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v1.28.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp v1.28.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.28.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.28.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.28.0 // indirect
	go.opentelemetry.io/otel/exporters/prometheus v0.50.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.28.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.28.0 // indirect
	go.opentelemetry.io/otel/log v0.4.0 // indirect
	go.opentelemetry.io/otel/metric v1.28.0 // indirect
	go.opentelemetry.io/otel/sdk v1.28.0 // indirect
	go.opentelemetry.io/otel/sdk/log v0.4.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.28.0 // indirect
	go.opentelemetry.io/otel/trace v1.28.0 // indirect
	go.opentelemetry.io/proto/otlp v1.3.1 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/crypto v0.26.0 // indirect
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842 // indirect
	golang.org/x/mod v0.18.0 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/oauth2 v0.21.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/term v0.23.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.22.0 // indirect
	gonum.org/v1/gonum v0.15.0 // indirect
	google.golang.org/api v0.183.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240723171418-e6d459c13d2a // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240723171418-e6d459c13d2a // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gotest.tools/v3 v3.5.0 // indirect
	k8s.io/api v0.29.3 // indirect
	k8s.io/apimachinery v0.29.3 // indirect
	k8s.io/client-go v0.29.3 // indirect
	k8s.io/klog/v2 v2.120.1 // indirect
	k8s.io/kube-openapi v0.0.0-20240228011516-70dd3763d340 // indirect
	k8s.io/utils v0.0.0-20230726121419-3b25d923346b // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/connector/spanmetricsconnector => ../connector/spanmetricsconnector

replace github.com/open-telemetry/opentelemetry-collector-contrib/connector/routingconnector => ../connector/routingconnector

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/carbonexporter => ../exporter/carbonexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/opencensusexporter => ../exporter/opencensusexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusexporter => ../exporter/prometheusexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusremotewriteexporter => ../exporter/prometheusremotewriteexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sapmexporter => ../exporter/sapmexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter => ../exporter/signalfxexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/splunkhecexporter => ../exporter/splunkhecexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/syslogexporter => ../exporter/syslogexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/zipkinexporter => ../exporter/zipkinexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../internal/common

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/sharedcomponent => ../internal/sharedcomponent

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk => ../internal/splunk

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchperresourceattr => ../pkg/batchperresourceattr

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata => ../pkg/experimentalmetricmetadata

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/opencensus => ../pkg/translator/opencensus

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheus => ../pkg/translator/prometheus

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheusremotewrite => ../pkg/translator/prometheusremotewrite

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/signalfx => ../pkg/translator/signalfx

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/zipkin => ../pkg/translator/zipkin

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver => ../receiver/carbonreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/datadogreceiver => ../receiver/datadogreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jaegerreceiver => ../receiver/jaegerreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/opencensusreceiver => ../receiver/opencensusreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver => ../receiver/prometheusreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sapmreceiver => ../receiver/sapmreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/signalfxreceiver => ../receiver/signalfxreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/splunkhecreceiver => ../receiver/splunkhecreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/syslogreceiver => ../receiver/syslogreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zipkinreceiver => ../receiver/zipkinreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/testbed/mockdatasenders/mockdatadogagentexporter => ./mockdatasenders/mockdatadogagentexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/jaeger => ../pkg/translator/jaeger

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../internal/coreinternal

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/resourcetotelemetry => ../pkg/resourcetotelemetry

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza => ../pkg/stanza

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage => ../extension/storage

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../pkg/pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../pkg/pdatautil

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden => ../pkg/golden

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/ackextension => ../extension/ackextension

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl => ../pkg/ottl

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/exp/metrics => ../internal/exp/metrics
