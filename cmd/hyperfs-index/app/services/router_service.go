// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/euskadi31/go-server"
	"github.com/euskadi31/go-service"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
)

// Services keys
const (
	RouterKey = "service.router"
)

func init() {
	service.Set(RouterKey, func(c service.Container) interface{} {
		// cfg := c.Get(ConfigKey).(*config.Configuration)
		logger := c.Get(LoggerKey).(zerolog.Logger)

		router := server.NewRouter()

		router.Use(hlog.NewHandler(logger))
		router.Use(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
			hlog.FromRequest(r).Info().
				Str("method", r.Method).
				Str("url", r.URL.String()).
				Int("status", status).
				Int("size", size).
				Dur("duration", duration).
				Msg(fmt.Sprintf("%s %s", r.Method, r.URL.String()))
		}))
		router.Use(hlog.RemoteAddrHandler("ip"))
		router.Use(hlog.UserAgentHandler("user_agent"))
		router.Use(hlog.RefererHandler("referer"))
		router.Use(hlog.RequestIDHandler("req_id", "Request-Id"))

		router.EnableHealthCheck()
		router.EnableMetrics()

		return router // *server.Router
	})
}
