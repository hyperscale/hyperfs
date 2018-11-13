// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package services

import (
	"flag"
	"os"

	"github.com/euskadi31/go-service"
)

// FlagsKey flags service
const FlagsKey = "service.flags"

func init() {
	service.Set(FlagsKey, func(c service.Container) interface{} {
		cmd := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

		return cmd // *flag.FlagSet
	})
}
