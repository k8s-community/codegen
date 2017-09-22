// Copyright 2017 Kubernetes Community Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/kelseyhightower/envconfig"
	"{[( .ProjectPath )]}/pkg/logger"
)

const (
	// SERVICENAME contains a service name prefix which used in ENV variables
	SERVICENAME = "{[( .EnvPrefix )]}"
)

// Config contains ENV variables
type Config struct {
	// Local service host
	LocalHost string `split_words:"true"`
	// Local service port
	LocalPort int `split_words:"true"`
	// Logging level in logger.Level notation
	LogLevel logger.Level `split_words:"true"`
}

// Load settles ENV variables into Config structure
func (c *Config) Load(serviceName string) error {
	return envconfig.Process(serviceName, c)
}
