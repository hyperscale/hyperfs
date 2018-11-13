// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package app

import (
	"os"

	"github.com/hyperscale/hyperfs/cmd/hyperfs/app/cmd"
)

// Run Command Line Application
func Run() error {
	return cmd.NewHyperFSCommand(os.Stdout, os.Stderr).Execute()
}
