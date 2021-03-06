// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package services

import (
	stdlog "log"
	"os"

	"github.com/euskadi31/go-service"
	"github.com/hyperscale/hyperfs/cmd/hyperfs-index/app/config"
	"github.com/hyperscale/hyperfs/pkg/hyperfs/version"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Services keys
const (
	LoggerKey = "service.logger"
)

func init() {
	service.Set(LoggerKey, func(c service.Container) interface{} {
		cfg := c.Get(ConfigKey).(*config.Configuration)

		zerolog.SetGlobalLevel(cfg.Logger.Level())

		zerolog.CallerSkipFrameCount = 2

		logger := zerolog.New(os.Stdout).With().
			Timestamp().
			Str("role", cfg.Logger.Prefix).
			Str("version", version.Get().Version).
			//Str("env", cfg.Environment.String()).
			Caller().
			Logger()

		fi, err := os.Stdin.Stat()
		if err != nil {
			log.Fatal().Err(err).Msg("Stdin.Stat failed")
		}

		if (fi.Mode() & os.ModeCharDevice) != 0 {
			logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		}

		stdlog.SetFlags(0)
		stdlog.SetOutput(logger)

		log.Logger = logger

		return logger // zerolog.Logger
	})
}
