// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package cluster

import (
	stdnet "net"
	"os"

	"github.com/hashicorp/mdns"
	"github.com/hyperscale/hyperfs/pkg/hyperfs/net"
	"github.com/pkg/errors"
)

// Cluster interface
type Cluster interface {
	Run() error
	Close() error
}

type cluster struct {
	cfg  *Configuration
	mdns *mdns.Server
}

// New Cluster
func New(cfg *Configuration) Cluster {
	return &cluster{
		cfg: cfg,
	}
}

func (c *cluster) Run() error {
	//ip, err := net.ExternalIP()
	ip, err := net.GetOutboundIP()
	if err != nil {
		return errors.Wrap(err, "net.GetOutboundIP")
	}

	// Setup our service export
	host, _ := os.Hostname()
	info := []string{c.cfg.Name}
	service, err := mdns.NewMDNSService(host, "_hyperfs._tcp", "", "", 7946, []stdnet.IP{ip}, info)
	if err != nil {
		return errors.Wrap(err, "mdns.NewMDNSService")
	}

	mdns, err := mdns.NewServer(&mdns.Config{
		Zone: service,
	})
	if err != nil {
		return errors.Wrap(err, "mdns.NewServer")
	}

	c.mdns = mdns

	return nil
}

func (c *cluster) Close() error {
	if err := c.mdns.Shutdown(); err != nil {
		return errors.Wrap(err, "mdns.Shutdown")
	}

	return nil
}
