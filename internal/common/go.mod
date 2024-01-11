module github.com/open-telemetry/opentelemetry-collector-contrib/internal/common

go 1.20

require (
	github.com/stretchr/testify v1.8.4
	go.opentelemetry.io/collector/featuregate v1.0.2-0.20240110091511-bf804d6c4ecc
	go.uber.org/zap v1.26.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)
