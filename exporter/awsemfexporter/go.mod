module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter

go 1.17

require (
	github.com/aws/aws-sdk-go v1.42.27
	github.com/census-instrumentation/opencensus-proto v0.3.0
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.3.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/awsutil v0.41.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/cloudwatch v0.41.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/metrics v0.41.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/resourcetotelemetry v0.41.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/opencensus v0.41.0
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.41.1-0.20220105191026-60823e25df1c
	go.opentelemetry.io/collector/model v0.41.1-0.20220105191026-60823e25df1c
	go.uber.org/zap v1.20.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/cenkalti/backoff/v4 v4.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/knadh/koanf v1.4.0 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.41.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/stretchr/objx v0.1.1 // indirect
	go.opencensus.io v0.23.0 // indirect
	go.opentelemetry.io/otel v1.3.0 // indirect
	go.opentelemetry.io/otel/metric v0.26.0 // indirect
	go.opentelemetry.io/otel/trace v1.3.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	golang.org/x/net v0.0.0-20211216030914-fe4d6282115f // indirect
	golang.org/x/sys v0.0.0-20211205182925-97ca703d548d // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa // indirect
	google.golang.org/grpc v1.43.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/metrics => ./../../internal/aws/metrics

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/awsutil => ./../../internal/aws/awsutil

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/cloudwatch => ./../../internal/aws/cloudwatch

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/resourcetotelemetry => ../../pkg/resourcetotelemetry

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/opencensus => ../../pkg/translator/opencensus
