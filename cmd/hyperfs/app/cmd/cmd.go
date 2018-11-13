// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"io"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hyperfs",
	Short: "HyperFS Command Line",
}

// NewHyperFSCommand create commands
func NewHyperFSCommand(out io.Writer, err io.Writer) *cobra.Command {
	rootCmd.AddCommand(NewCmdCluster(out))
	rootCmd.AddCommand(NewCmdVersion(out))

	return rootCmd
}
