// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"io"

	"github.com/spf13/cobra"
)

// NewCmdCluster builder
func NewCmdCluster(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cluster",
		Short: "Manage cluster",
	}

	cmd.AddCommand(NewCmdClusterInfo(out))

	return cmd
}
