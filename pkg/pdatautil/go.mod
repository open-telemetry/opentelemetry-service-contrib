module github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil

go 1.22.0

require (
	github.com/cespare/xxhash/v2 v2.3.0
	github.com/stretchr/testify v1.9.0
	go.opentelemetry.io/collector/pdata v1.14.2-0.20240829190554-7da6b618a7ee
	go.uber.org/goleak v1.3.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
	google.golang.org/grpc v1.65.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
	v0.76.2
	v0.76.1
)
