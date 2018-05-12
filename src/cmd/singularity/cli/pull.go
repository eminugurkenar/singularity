/*
  Copyright (c) 2018, Sylabs, Inc. All rights reserved.

  This software is licensed under a 3-clause BSD license.  Please
  consult LICENSE file distributed with the sources of this project regarding
  your rights to use or distribute this software.
*/
package cli

import (
	"github.com/singularityware/singularity/src/pkg/libexec"
	"github.com/spf13/cobra"

	"github.com/singularityware/singularity/docs"
)

var (
	// PullLibraryURI holds the base URI to a Sylabs library API instance
	PullLibraryURI string
)

func init() {
	pullCmd.Flags().SetInterspersed(false)
	SingularityCmd.AddCommand(pullCmd)

	pullCmd.Flags().BoolVarP(&Force, "force", "F", false, "overwrite an image file if it exists")
	pullCmd.Flags().StringVar(&PullLibraryURI, "libraryuri", "http://localhost:5150", "")
}

var pullCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Args: cobra.RangeArgs(1, 2),

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			libexec.PullImage(args[0], args[1], PullLibraryURI, Force)
			return
		}
		libexec.PullImage("", args[0], PullLibraryURI, Force)
	},

	Use:     docs.PullUse,
	Short:   docs.PullShort,
	Long:    docs.PullLong,
	Example: docs.PullExample,
}
