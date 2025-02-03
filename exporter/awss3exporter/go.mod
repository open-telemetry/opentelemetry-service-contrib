module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awss3exporter

go 1.22.0

require (
	github.com/aws/aws-sdk-go-v2 v1.34.0
	github.com/aws/aws-sdk-go-v2/config v1.29.2
	github.com/aws/aws-sdk-go-v2/credentials v1.17.55
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.17.54
	github.com/aws/aws-sdk-go-v2/service/s3 v1.74.1
	github.com/aws/aws-sdk-go-v2/service/sts v1.33.10
	github.com/stretchr/testify v1.10.0
	github.com/tilinna/clock v1.1.0
	go.opentelemetry.io/collector/component v0.118.1-0.20250203014413-643a35ffbcea
	go.opentelemetry.io/collector/component/componenttest v0.118.1-0.20250203014413-643a35ffbcea
	go.opentelemetry.io/collector/config/configcompression v1.24.1-0.20250203014413-643a35ffbcea
	go.opentelemetry.io/collector/confmap v1.24.1-0.20250203014413-643a35ffbcea
	go.opentelemetry.io/collector/consumer v1.24.1-0.20250203014413-643a35ffbcea
	go.opentelemetry.io/collector/exporter v0.118.1-0.20250203014413-643a35ffbcea
	go.opentelemetry.io/collector/exporter/exportertest v0.118.1-0.20250203014413-643a35ffbcea
	go.opentelemetry.io/collector/otelcol/otelcoltest v0.118.1-0.20250203014413-643a35ffbcea
	go.opentelemetry.io/collector/pdata v1.24.1-0.20250203014413-643a35ffbcea
	go.uber.org/goleak v1.3.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.27.0
)

require (
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.6.8 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.25 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.29 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.29 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.2 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.3.29 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.12.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.5.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.12.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.18.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.24.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.28.11 // indirect
	github.com/aws/smithy-go v1.22.2 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/ebitengine/purego v0.8.1 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.25.1 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.2 // indirect
	github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/prometheus/client_golang v1.20.5 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.62.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/shirou/gopsutil/v4 v4.24.12 // indirect
	github.com/spf13/cobra v1.8.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/collector/component/componentstatus v0.118.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/config/configretry v1.24.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.118.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/confmap/provider/envprovider v1.24.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/confmap/provider/fileprovider v1.24.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/confmap/provider/httpprovider v1.24.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/confmap/provider/yamlprovider v1.24.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/connector v0.118.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/connector/connectortest v0.118.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/connector/xconnector v0.118.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/consumer/consumererror v0.118.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/consumer/consumertest v0.118.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/consumer/xconsumer v0.118.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/exporter/xexporter v0.118.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/extension v0.118.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/extension/extensioncapabilities v0.118.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/extension/extensiontest v0.118.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/extension/xextension v0.118.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/featuregate v1.24.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/internal/fanoutconsumer v0.118.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/otelcol v0.118.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.118.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/pdata/testdata v0.118.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/pipeline v0.118.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/pipeline/xpipeline v0.118.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/processor v0.118.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/processor/processortest v0.118.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/processor/xprocessor v0.118.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/receiver v0.118.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/receiver/receivertest v0.118.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/receiver/xreceiver v0.118.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/semconv v0.118.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/collector/service v0.118.1-0.20250203014413-643a35ffbcea // indirect
	go.opentelemetry.io/contrib/bridges/otelzap v0.9.0 // indirect
	go.opentelemetry.io/contrib/config v0.14.0 // indirect
	go.opentelemetry.io/contrib/propagators/b3 v1.34.0 // indirect
	go.opentelemetry.io/otel v1.34.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc v0.10.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp v0.10.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v1.34.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp v1.34.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.34.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.34.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.34.0 // indirect
	go.opentelemetry.io/otel/exporters/prometheus v0.56.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdoutlog v0.10.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.34.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.34.0 // indirect
	go.opentelemetry.io/otel/log v0.10.0 // indirect
	go.opentelemetry.io/otel/metric v1.34.0 // indirect
	go.opentelemetry.io/otel/sdk v1.34.0 // indirect
	go.opentelemetry.io/otel/sdk/log v0.10.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.34.0 // indirect
	go.opentelemetry.io/otel/trace v1.34.0 // indirect
	go.opentelemetry.io/proto/otlp v1.5.0 // indirect
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	gonum.org/v1/gonum v0.15.1 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250115164207-1a7da9e5054f // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f // indirect
	google.golang.org/grpc v1.70.0 // indirect
	google.golang.org/protobuf v1.36.4 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
	v0.76.2
	v0.76.1
)
