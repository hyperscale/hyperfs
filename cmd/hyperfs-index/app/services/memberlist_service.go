// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package services

import (
	"os"

	"github.com/euskadi31/go-service"
	"github.com/hashicorp/memberlist"
	"github.com/rs/zerolog/log"
)

// MemberlistKey Hashicorp Memberlist service
const MemberlistKey = "service.memberlist"

func init() {
	service.Set(MemberlistKey, func(c service.Container) interface{} {
		hostname, err := os.Hostname()
		if err != nil {
			log.Fatal().Err(err).Msg("os.Hostname")
		}

		cfg := memberlist.DefaultLANConfig()
		cfg.Name = hostname

		s, err := memberlist.Create(cfg)
		if err != nil {
			log.Fatal().Err(err).Msg(MemberlistKey)
		}

		return s
	})
}
