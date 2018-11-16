// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package discover

import (
	"fmt"

	"github.com/hyperscale/hyperfs/pkg/hyperfs/net"
	"github.com/pkg/errors"
)

// Discover interface
type Discover interface {
	AddProvider(provider Provider) error
	Addrs() ([]string, error)
	Close() error
}

type discover struct {
	providers map[string]Provider
}

// New Discover
func New() Discover {
	return &discover{
		providers: make(map[string]Provider),
	}
}

func (d *discover) AddProvider(provider Provider) error {
	if _, ok := d.providers[provider.Name()]; ok {
		return fmt.Errorf("provider %s already exists", provider.Name())
	}

	d.providers[provider.Name()] = provider

	return nil
}

func (d *discover) Addrs() ([]string, error) {
	addrs := []string{}

	node, err := net.ExternalIP()
	if err != nil {
		return addrs, errors.Wrap(err, "net.ExternalIP")
	}

	for _, provider := range d.providers {
		ips, err := provider.Addrs()
		if err != nil {
			return addrs, errors.Wrap(err, provider.Name())
		}

		for _, ip := range ips {
			if ip == node.String() {
				continue
			}

			addrs = append(addrs, ip)
		}
	}

	return addrs, nil
}

func (d *discover) Close() error {
	for _, provider := range d.providers {
		if err := provider.Close(); err != nil {
			return errors.Wrap(err, provider.Name())
		}
	}

	return nil
}
