// Copyright 2017 Kubernetes Community Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"

	"{[( .ProjectPath )]}/pkg/config"
	"{[( .ProjectPath )]}/pkg/service"
	"{[( .ProjectPath )]}/pkg/system"
)

func main() {
	// Load ENV configuration
	cfg := new(config.Config)
	if err := cfg.Load(config.SERVICENAME); err != nil {
		log.Fatal(err)
	}

	// Configure service and get router
	router, logger, err := service.Setup(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Listen and serve handlers
	go router.Listen(fmt.Sprintf("%s:%d", cfg.LocalHost, cfg.LocalPort))

	// Wait signals
	signals := system.NewSignals()
	if err := signals.Wait(logger, new(system.Handling)); err != nil {
		logger.Fatal(err)
	}
}
