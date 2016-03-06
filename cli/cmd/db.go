// Copyright © 2016 Ryan D <ryan0x44.com>
package cmd

import (
	"github.com/ryan0x44/harveycms/core"
	"github.com/spf13/cobra"
)

var DbInitCmd = &cobra.Command{
	Use:   "db:init",
	Short: "Initialise database",
	Long:  `Creates database and all required tables.`,
	Run: func(cmd *cobra.Command, args []string) {
		db := core.DbConnect("dev:password@/harveycms?charset=utf8&parseTime=True&loc=Local")
		db.CreateTable(&core.Route{})
	},
}
