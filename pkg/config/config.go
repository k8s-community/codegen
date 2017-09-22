package config

import (
	"github.com/kelseyhightower/envconfig"
)

const (
	// SERVICENAME contains a service name prefix which used in ENV variables
	SERVICENAME = "CODEGEN"
)

// Config contains ENV variables
type Config struct {
	// Local service host
	LocalHost string `split_words:"true"`
	// Local service port
	LocalPort int `split_words:"true"`
	// Logging level in logger.Level notation
	LogLevel int `split_words:"true"`
}

// Load settles ENV variables into Config structure
func (c *Config) Load(serviceName string) error {
	return envconfig.Process(serviceName, c)
}
