// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// NewCmdClusterInfo builder
func NewCmdClusterInfo(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "Print the cluster information",
		RunE: func(cmd *cobra.Command, args []string) error {

			const padding = 3
			w := tabwriter.NewWriter(out, 0, 0, padding, ' ', 0)

			fmt.Fprintln(w, "IP\tROLE\tVERSION\tSTATUS\t")
			fmt.Fprintln(w, "10.12.5.18\tStorage\tv1.2.4\tReady\t")
			fmt.Fprintln(w, "10.12.5.19\tStorage\tv1.2.4\tReady\t")
			fmt.Fprintln(w, "10.12.5.20\tIndex\tv1.2.4\tReady\t")
			fmt.Fprintln(w, "10.12.5.21\tIndex\tv1.2.4\tReady\t")
			fmt.Fprintln(w, "10.12.5.22\tApi\tv1.2.4\tReady\t")

			return w.Flush()
		},
	}

	return cmd
}
