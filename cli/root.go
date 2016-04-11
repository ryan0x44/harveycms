// Copyright © 2016 Ryan D <ryan0x44.com>
package cli

import (
	"github.com/spf13/cobra"
)

var RootDirectory string

var RootCmd = &cobra.Command{
	Use:              "harvey-cli",
	Short:            "Harvey CLI is a task runner for Harvey CMS",
	Long:             `Task runner for Harvey CMS <http://www.harveycms.org>`,
	PersistentPreRun: RootPresistentPreRun,
}

func init() {
	RootCmd.PersistentFlags().StringVarP(
		&RootDirectory,
		"dir",
		"d",
		"demo/",
		"Path to instance directory",
	)
	RootCmd.AddCommand(DbInitCmd)
}

func RootPresistentPreRun(cmd *cobra.Command, args []string) {
	LoadConfig(RootDirectory)
}
