// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package app

import (
	"os"
	"os/signal"
	"net/http"
	"syscall"
	"encoding/json"
	"time"

	"github.com/euskadi31/go-service"
	"github.com/hashicorp/memberlist"
	"github.com/hyperscale/hyperfs/cmd/hyperfs-index/app/services"
	"github.com/hyperscale/hyperfs/pkg/hyperfs/cluster"
	"github.com/hyperscale/hyperfs/pkg/hyperfs/cluster/discover"
	"github.com/rs/zerolog/log"
	"github.com/pkg/errors"
)

// Run HyperFS Index server
func Run() error {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	_ = service.Get(services.LoggerKey)
	m := service.Get(services.MemberlistKey).(*memberlist.Memberlist)
	cltr := service.Get(services.ClusterKey).(cluster.Cluster)
	dscvr := service.Get(services.DiscoverKey).(discover.Discover)

	if err := cltr.Run(); err != nil {
		return errors.Wrap(err, "cluster.Run")
	}

	addrs, err := dscvr.Addrs()
	if err != nil {
		return errors.Wrap(err, "discover.Addrs")
	}

	if len(addrs) > 0 {
		if _, err := m.Join(addrs); err != nil {
			return errors.Wrap(err, "memberlist.Join")
		}
	}

	log.Info().Msg("Rinning")


	http.HandleFunc("/health", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/members", func(w http.ResponseWriter, req *http.Request) {
		json.NewEncoder(w).Encode(m.Members())
	})

	go func() {
		if err := http.ListenAndServe(":8000", nil); err != nil {
			log.Panic().Err(err).Msg("ListenAndServe")
		}
	}()



	<-sig

	if err := m.Leave(2 * time.Second); err != nil {
		return errors.Wrap(err, "memberlist.Leave")
	}

	if err := m.Shutdown(); err != nil {
		return errors.Wrap(err, "memberlist.Shutdown")
	}

	if err := cltr.Close(); err != nil {
		return errors.Wrap(err, "cluster.Close")
	}

	log.Info().Msg("Shutdown")

	return nil
}
