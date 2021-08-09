module github.com/open-telemetry/opentelemetry-collector-contrib/internal/stanza

go 1.16

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage v0.0.0-00010101000000-000000000000
	github.com/open-telemetry/opentelemetry-log-collection v0.20.0
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.31.1-0.20210809153342-28acc7d8b7f2
	go.opentelemetry.io/collector/model v0.31.1-0.20210809153342-28acc7d8b7f2
	go.uber.org/multierr v1.6.0
	go.uber.org/zap v1.18.1
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage => ../../extension/storage
