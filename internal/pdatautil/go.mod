module github.com/open-telemetry/opentelemetry-collector-contrib/internal/pdatautil

go 1.22.0

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest v0.116.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil v0.116.0
	github.com/stretchr/testify v1.10.0
	go.opentelemetry.io/collector/pdata v1.22.1-0.20250107062214-ced38e8af2ae
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241015192408-796eee8c2d53 // indirect
	google.golang.org/grpc v1.69.2 // indirect
	google.golang.org/protobuf v1.36.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../pkg/pdatautil

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../pkg/pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden => ../../pkg/golden

replace go.opentelemetry.io/collector/scraper/scraperhelper v0.116.0 => go.opentelemetry.io/collector/scraper/scraperhelper v0.0.0-20250107062214-ced38e8af2ae

replace go.opentelemetry.io/collector/extension/xextension v0.116.0 => go.opentelemetry.io/collector/extension/xextension v0.0.0-20250107062214-ced38e8af2ae
