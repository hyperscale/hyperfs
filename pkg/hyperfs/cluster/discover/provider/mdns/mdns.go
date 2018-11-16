// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package mdns

import (
	"fmt"
	"net"

	"github.com/hashicorp/mdns"
)

// Provider struct
type Provider struct {
	service  string
	entries  chan *mdns.ServiceEntry
	shutdown chan struct{}
	addrs    []net.IP
}

// New mDNS provider
func New(service string) *Provider {
	p := &Provider{
		service: service,
		//entries:  make(chan *mdns.ServiceEntry, 10),
		//shutdown: make(chan struct{}),
		//addrs:    []net.IP{},
	}

	//o p.loop()

	return p
}

/*
func (p *Provider) loop() {
	for {
		select {
		case entry := <-p.entries:
			p.addrs = append(p.addrs, entry.Addr)
		case <-p.shutdown:
			return
		}
	}
}
*/

// Name implements discover.Provider
func (p *Provider) Name() string {
	return "mdns"
}

// Addrs implements discover.Provider
func (p *Provider) Addrs() ([]string, error) {
	addrs := []string{}
	entries := make(chan *mdns.ServiceEntry, 100)

	go func(entries <-chan *mdns.ServiceEntry) {
		for entry := range entries {
			fmt.Printf("Got new entry: %v\n", entry)
			addrs = append(addrs, entry.Addr.String())
		}
	}(entries)

	if err := mdns.Lookup(fmt.Sprintf("_%s._tcp", p.service), entries); err != nil {
		return addrs, err
	}
	close(entries)

	return addrs, nil
	//return p.addrs
}

// Close implements discover.Provider
func (p *Provider) Close() error {
	//p.shutdown <- struct{}{}

	return nil
}
