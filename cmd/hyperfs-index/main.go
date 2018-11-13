// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grandcat/zeroconf"
	"github.com/rs/zerolog/log"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	hostname, _ := os.Hostname()

	server, err := zeroconf.Register(hostname, "_workstation._tcp", "local.", 42424, []string{"txtv=0", "lo=1", "la=2"}, nil)
	if err != nil {
		log.Panic().Err(err).Msg("zeroconf.Register")
	}
	defer server.Shutdown()

	// Discover all services on the network (e.g. _workstation._tcp)
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		log.Error().Err(err).Msg("Failed to initialize resolver")
	}

	nodes := []*zeroconf.ServiceEntry{}

	entries := make(chan *zeroconf.ServiceEntry)
	go func(results <-chan *zeroconf.ServiceEntry) {
		for entry := range results {
			fmt.Printf("Entry: %+v\n", entry)

			nodes = append(nodes, entry)
		}

		fmt.Println("No more entries.")
	}(entries)

	ctx := context.Background()

	if err = resolver.Browse(ctx, "_workstation._tcp", "local.", entries); err != nil {
		log.Error().Err(err).Msg("Failed to browse")
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, req *http.Request) {
		json.NewEncoder(w).Encode(true)
	})

	http.HandleFunc("/members", func(w http.ResponseWriter, req *http.Request) {
		json.NewEncoder(w).Encode(nodes)
	})

	go func() {
		if err := http.ListenAndServe(":8000", nil); err != nil {
			log.Panic().Err(err).Msg("ListenAndServe")
		}
	}()

	<-sig

	log.Info().Msg("Shutdown")
}
