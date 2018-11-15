// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"github.com/hashicorp/go-msgpack/codec"
	"github.com/hashicorp/memberlist"
	"github.com/hyperscale/hyperfs/pkg/hyperfs/cluster"
	"github.com/rs/zerolog/log"
)

// NodeDelegate struct
type NodeDelegate struct{}

// NewNodeDelegate constructor
func NewNodeDelegate() memberlist.Delegate {
	return &NodeDelegate{}
}

// NodeMeta implements memberlist.Delegate
func (d *NodeDelegate) NodeMeta(limit int) (meta []byte) {
	log.Debug().Int("limit", limit).Msg("NodeMeta")

	if err := codec.NewEncoderBytes(&meta, &codec.MsgpackHandle{}).Encode(&cluster.NodeMeta{
		Cluster: "test",
		Region:  "fr",
		Type:    cluster.NodeTypeIndex,
	}); err != nil {
		log.Error().Err(err).Msg("codec.NewEncoderBytes")
	}

	return
}

// NotifyMsg implements memberlist.Delegate
func (d *NodeDelegate) NotifyMsg(b []byte) {
	log.Debug().Msgf("NotifyMsg: %s", string(b))

	if len(b) == 0 {
		return
	}

}

// GetBroadcasts implements memberlist.Delegate
func (d *NodeDelegate) GetBroadcasts(overhead int, limit int) [][]byte {
	log.Debug().Int("overhead", overhead).Int("limit", limit).Msg("GetBroadcasts")

	return [][]byte{}
}

// LocalState implements memberlist.Delegate
func (d *NodeDelegate) LocalState(join bool) []byte {
	log.Debug().Bool("join", join).Msg("LocalState")

	return []byte{}
}

// MergeRemoteState implements memberlist.Delegate
func (d *NodeDelegate) MergeRemoteState(buf []byte, join bool) {
	log.Debug().Bool("join", join).Msgf("MergeRemoteState: %s", string(buf))
}
