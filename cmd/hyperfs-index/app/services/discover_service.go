// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package services

import (
	"github.com/euskadi31/go-service"
	"github.com/hyperscale/hyperfs/pkg/hyperfs/cluster/discover"
	"github.com/hyperscale/hyperfs/pkg/hyperfs/cluster/discover/provider/mdns"
	"github.com/rs/zerolog/log"
)

// Services keys
const (
	DiscoverKey             = "service.discover"
	DiscoverMDNSProviderKey = "service.discover.provider.mdns"
)

func init() {

	service.Set(DiscoverMDNSProviderKey, func(c service.Container) interface{} {
		return mdns.New("hyperfs") // discover.Provider
	})

	service.Set(DiscoverKey, func(c service.Container) interface{} {
		mdnsProvider := c.Get(DiscoverMDNSProviderKey).(discover.Provider)

		d := discover.New()

		if err := d.AddProvider(mdnsProvider); err != nil {
			log.Fatal().Err(err).Msg("discover.AddProvider")
		}

		return d // discover.Discover
	})
}
