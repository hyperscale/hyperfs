// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"context"

	"github.com/hyperscale/hyperfs/cmd/hyperfs-index/app"
	"github.com/hyperscale/hyperfs/pkg/hyperfs/net"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func main() {
	extIP, _ := net.ExternalIP()
	locIP, _ := net.LocalIP()
	outIP, _ := net.GetOutboundIP()

	log.Fatal().Str("ext_ip", extIP.String()).Str("loc_ip", locIP.String()).Str("out_ip", outIP.String()).Msg("Net")

	if err := app.Run(); err != nil {
		if errors.Cause(err) == context.Canceled {
			log.Debug().Err(err).Msg("ignore error since context is cancelled")
		} else {
			log.Fatal().Err(err).Msg("hyperfs run failed")
		}
	}
}
