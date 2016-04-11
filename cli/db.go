// Copyright © 2016 Ryan D <ryan0x44.com>
package cli

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/ryan0x44/harveycms/core"
	"github.com/ryan0x44/harveycms/modules/pages"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
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
		db := DbConnect(dbUrl)
		db.Set("gorm:table_options", "ENGINE=InnoDB")
		core.DbInit(&db)
		pages.DbInit(&db)
	},
}

func DbConnect(DbUrl string) gorm.DB {
	db, err := gorm.Open("mysql", DbUrl)
	if err != nil {
		log.Panicf("Error connecting to database:\n%s", err)
	}
	return db
}
