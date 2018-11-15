// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package services

import (
	stdlog "log"

	"github.com/euskadi31/go-service"
	"github.com/google/uuid"
	"github.com/hashicorp/memberlist"
	"github.com/hyperscale/hyperfs/cmd/hyperfs-index/app/cluster"
	"github.com/hyperscale/hyperfs/cmd/hyperfs-index/app/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Services keys
const (
	MemberlistKey              = "service.memberlist"
	MemberlistNodeDelegateKey  = "service.memberlist.node.delegate"
	MemberlistEventDelegateKey = "service.memberlist.event.delegate"
)

func init() {
	service.Set(MemberlistEventDelegateKey, func(c service.Container) interface{} {
		return cluster.NewEventDelegate(16)
	})

	service.Set(MemberlistNodeDelegateKey, func(c service.Container) interface{} {
		return cluster.NewNodeDelegate()
	})

	service.Set(MemberlistKey, func(c service.Container) interface{} {
		cfg := c.Get(ConfigKey).(*config.Configuration)
		logger := c.Get(LoggerKey).(zerolog.Logger)
		nodeDelegate := c.Get(MemberlistNodeDelegateKey).(memberlist.Delegate)
		eventDelegate := c.Get(MemberlistEventDelegateKey).(cluster.EventDelegate)

		conf := memberlist.DefaultLANConfig()
		conf.Name = uuid.New().String()
		conf.Delegate = nodeDelegate
		//conf.BindPort = 0
		conf.Events = eventDelegate.Delegate()
		conf.Logger = stdlog.New(logger, "", 0)
		conf.SecretKey = []byte(cfg.Cluster.Key)

		m, err := memberlist.Create(conf)
		if err != nil {
			log.Fatal().Err(err).Msg(MemberlistKey)
		}

		node := m.LocalNode()
		log.Debug().Msgf("Local member %s:%d", node.Addr, node.Port)

		go func(events <-chan memberlist.NodeEvent) {
			for {
				select {
				case e := <-events:
					switch e.Event {
					case memberlist.NodeJoin:
						log.Debug().
							Str("name", e.Node.Name).
							Str("host", e.Node.Addr.String()).
							Uint16("port", e.Node.Port).
							Str("meta", string(e.Node.Meta)).
							Msg("Node join cluster")
					case memberlist.NodeLeave:
						log.Debug().
							Str("name", e.Node.Name).
							Str("host", e.Node.Addr.String()).
							Uint16("port", e.Node.Port).
							Str("meta", string(e.Node.Meta)).
							Msg("Node leave cluster")
					case memberlist.NodeUpdate:
						log.Debug().
							Str("name", e.Node.Name).
							Str("host", e.Node.Addr.String()).
							Uint16("port", e.Node.Port).
							Str("meta", string(e.Node.Meta)).
							Msg("Node update cluster")
					}
				}
			}
		}(eventDelegate.Events())

		return m // *memberlist.Memberlist
	})
}
