// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"io"

	"github.com/hyperscale/hyperfs/pkg/hyperfs/version"
	"github.com/spf13/cobra"
)

// NewCmdVersion builder
func NewCmdVersion(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version information",
		RunE: func(cmd *cobra.Command, args []string) error {
			if _, err := fmt.Fprintf(out, "%s\n", version.Get().Version); err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}
