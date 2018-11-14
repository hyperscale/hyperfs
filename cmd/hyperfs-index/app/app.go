// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package app

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/euskadi31/go-service"
	"github.com/hashicorp/memberlist"
	"github.com/hyperscale/hyperfs/cmd/hyperfs-index/app/services"
	"github.com/rs/zerolog/log"
)

// Run HyperFS Index server
func Run() error {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	_ = service.Get(services.LoggerKey)
	m := service.Get(services.MemberlistKey).(*memberlist.Memberlist)

	log.Info().Msg("Rinning")

	<-sig

	if err := m.Leave(2 * time.Second); err != nil {
		log.Error().Err(err).Msg("memberlist.Leave")
	}

	if err := m.Shutdown(); err != nil {
		log.Error().Err(err).Msg("memberlist.Shutdown")
	}

	log.Info().Msg("Shutdown")

	return nil
}
