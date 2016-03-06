// Copyright © 2016 Ryan D <ryan0x44.com>
package main

import (
	"github.com/ryan0x44/harveycms/core"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var DbInitCmd = &cobra.Command{
	Use:   "db:init",
	Short: "Initialise database",
	Long:  `Creates database and all required tables.`,
	Run: func(cmd *cobra.Command, args []string) {
		dbUrl := fmt.Sprintf(
			"%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local",
			viper.Get("db.user"),
			viper.Get("db.pass"),
			viper.Get("db.host"),
			viper.Get("db.name"),
		)
		db := core.DbConnect(dbUrl)
		db.CreateTable(&core.Route{})
	},
}
