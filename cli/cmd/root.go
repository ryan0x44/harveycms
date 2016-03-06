// Copyright © 2016 Ryan D <ryan0x44.com>
package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "harvey-cli",
	Short: "Harvey CLI is a task runner for Harvey CMS",
	Long:  `Task runner for Harvey CMS <http://www.harveycms.org>`,
}

func init() {
	var ConfigFile string
	RootCmd.PersistentFlags().StringVarP(
		&ConfigFile,
		"config",
		"c",
		"demo/config.toml",
		"Path to config file",
	)
	RootCmd.AddCommand(DbInitCmd)
}
