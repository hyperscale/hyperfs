// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package config

import (
	"github.com/hyperscale/hyperfs/pkg/hyperfs/cluster"
	"github.com/hyperscale/hyperfs/pkg/hyperfs/logger"
)

// Configuration struct
type Configuration struct {
	Logger  *logger.Configuration
	Cluster *cluster.Configuration
}

// NewConfiguration constructor
func NewConfiguration() *Configuration {
	return &Configuration{
		Logger:  &logger.Configuration{},
		Cluster: &cluster.Configuration{},
	}
}
