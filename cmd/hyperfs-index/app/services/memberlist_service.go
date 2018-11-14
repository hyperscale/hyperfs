// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package services

import (
	stdlog "log"

	"github.com/euskadi31/go-service"
	"github.com/google/uuid"
	"github.com/hashicorp/memberlist"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// MemberlistKey Hashicorp Memberlist service
const MemberlistKey = "service.memberlist"

func init() {
	service.Set(MemberlistKey, func(c service.Container) interface{} {
		logger := c.Get(LoggerKey).(zerolog.Logger)

		events := make(chan memberlist.NodeEvent, 16)

		conf := memberlist.DefaultLANConfig()
		conf.Name = uuid.New().String()
		conf.Delegate = &memberlistDelegate{}
		//conf.BindPort = 0
		conf.Events = &memberlist.ChannelEventDelegate{
			Ch: events,
		}
		conf.Logger = stdlog.New(logger, "", 0)

		m, err := memberlist.Create(conf)
		if err != nil {
			log.Fatal().Err(err).Msg(MemberlistKey)
		}

		node := m.LocalNode()
		log.Debug().Msgf("Local member %s:%d", node.Addr, node.Port)

		go func(events chan memberlist.NodeEvent) {
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
		}(events)

		return m // *memberlist.Memberlist
	})
}

type memberlistDelegate struct{}

func (d *memberlistDelegate) NodeMeta(limit int) []byte {
	log.Debug().Int("limit", limit).Msg("NodeMeta")

	return []byte("region=fr")
}

func (d *memberlistDelegate) NotifyMsg(b []byte) {
	log.Debug().Msgf("NotifyMsg: %s", string(b))

	if len(b) == 0 {
		return
	}

}

func (d *memberlistDelegate) GetBroadcasts(overhead int, limit int) [][]byte {
	log.Debug().Int("overhead", overhead).Int("limit", limit).Msg("GetBroadcasts")

	return [][]byte{}
}

func (d *memberlistDelegate) LocalState(join bool) []byte {
	log.Debug().Bool("join", join).Msg("LocalState")

	return []byte{}
}

func (d *memberlistDelegate) MergeRemoteState(buf []byte, join bool) {
	log.Debug().Bool("join", join).Msgf("MergeRemoteState: %s", string(buf))
}
