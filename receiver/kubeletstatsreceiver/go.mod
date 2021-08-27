module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver

go 1.16

require (
	github.com/census-instrumentation/opencensus-proto v0.3.0
	github.com/onsi/ginkgo v1.14.1 // indirect
	github.com/onsi/gomega v1.10.2 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/interval v0.0.0-00010101000000-000000000000
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig v0.0.0-00010101000000-000000000000
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/kubelet v0.0.0-00010101000000-000000000000
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/opencensus v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.33.1-0.20210827152330-09258f969908
	go.opentelemetry.io/collector/model v0.33.1-0.20210827152330-09258f969908
	go.uber.org/zap v1.19.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.22.1
	k8s.io/apimachinery v0.22.1
	k8s.io/client-go v0.22.1
	k8s.io/kubelet v0.22.1

)

replace (
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../../internal/common
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/interval => ../../internal/interval
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig => ../../internal/k8sconfig
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/kubelet => ../../internal/kubelet
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/opencensus => ../../pkg/translator/opencensus
)
