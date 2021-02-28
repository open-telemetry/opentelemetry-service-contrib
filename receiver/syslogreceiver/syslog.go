package syslogreceiver

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/stanza"
	"github.com/open-telemetry/opentelemetry-log-collection/operator"
	"github.com/open-telemetry/opentelemetry-log-collection/operator/builtin/input/syslog"
	"github.com/open-telemetry/opentelemetry-log-collection/operator/builtin/input/tcp"
	"github.com/open-telemetry/opentelemetry-log-collection/operator/builtin/input/udp"
	syslogparser "github.com/open-telemetry/opentelemetry-log-collection/operator/builtin/parser/syslog"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configmodels"
	"gopkg.in/yaml.v2"
)


const typeStr = "syslog"

// NewFactory creates a factory for filelog receiver
func NewFactory() component.ReceiverFactory {
	return stanza.NewFactory(ReceiverType{})
}

// ReceiverType implements stanza.LogReceiverType
// to create a syslog receiver
type ReceiverType struct{}

// Type is the receiver type
func (f ReceiverType) Type() configmodels.Type {
	return configmodels.Type(typeStr)
}

// CreateDefaultConfig creates a config with type and version
func (f ReceiverType) CreateDefaultConfig() configmodels.Receiver {
	 return &SysLogConfig{
		BaseConfig: stanza.BaseConfig{
			ReceiverSettings: configmodels.ReceiverSettings{
				TypeVal: configmodels.Type(typeStr),
				NameVal: typeStr,
			},
			Operators: stanza.OperatorConfigs{},
		},
		Input: stanza.InputConfig{},
	}
}

// BaseConfig gets the base config from config, for now
func (f ReceiverType) BaseConfig(cfg configmodels.Receiver) stanza.BaseConfig {
	return cfg.(*SysLogConfig).BaseConfig
}

// SysLogConfig defines configuration for the syslog receiver
type SysLogConfig struct {
	stanza.BaseConfig `mapstructure:",squash"`
	Input             stanza.InputConfig `mapstructure:",remain"`
}


// DecodeInputConfig unmarshals the input operator
func (f ReceiverType) DecodeInputConfig(cfg configmodels.Receiver) (*operator.Config, error) {
	logConfig := cfg.(*SysLogConfig)
	yamlBytes, _ := yaml.Marshal(logConfig.Input)
	inputCfg := syslog.NewSyslogInputConfig("syslog_input")
	if err := yaml.Unmarshal(yamlBytes, &inputCfg); err != nil {
		return nil, err
	}
	if inputCfg.Syslog != nil {
		inputCfg.Syslog = syslogparser.NewSyslogParserConfig("syslog_parser")
	}
	if inputCfg.Tcp != nil {
		inputCfg.Tcp = tcp.NewTCPInputConfig("tcp_input")
	}
	if inputCfg.Udp != nil {
		inputCfg.Udp = udp.NewUDPInputConfig("udp_input")
	}

	if err := yaml.Unmarshal(yamlBytes, &inputCfg); err != nil {
		return nil, err
	}
	return &operator.Config{Builder: inputCfg}, nil
}
