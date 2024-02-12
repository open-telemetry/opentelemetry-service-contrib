package metadata

import (
	"go.opentelemetry.io/collector/component"
)

var (
	Type = component.MustNewType("ack")
)

const (
	ExtensionStability = component.StabilityLevelAlpha
)
