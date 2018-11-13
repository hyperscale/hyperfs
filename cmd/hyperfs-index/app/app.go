// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/euskadi31/go-service"
	"github.com/hyperscale/hyperfs/cmd/hyperfs-index/app/services"
	"github.com/rs/zerolog/log"
)

// Run HyperFS Index server
func Run() error {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	_ = service.Get(services.LoggerKey)

	<-sig

	log.Info().Msg("Shutdown")

	return nil
}
