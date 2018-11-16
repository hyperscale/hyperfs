// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package services

import (
	"github.com/euskadi31/go-service"
	"github.com/hyperscale/hyperfs/cmd/hyperfs-index/app/config"
	"github.com/hyperscale/hyperfs/pkg/hyperfs/cluster"
)

// Services keys
const (
	ClusterKey = "service.cluster"
)

func init() {
	service.Set(ClusterKey, func(c service.Container) interface{} {
		cfg := c.Get(ConfigKey).(*config.Configuration)

		return cluster.New(cfg.Cluster) // cluster.Cluster
	})
}
