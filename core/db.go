// Copyright © 2016 Ryan D <ryan0x44.com>
package core

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

func DbConnect(DbUrl string) gorm.DB {
	db, err := gorm.Open("mysql", DbUrl)
	if err != nil {
		log.Panicf("Error connecting to database:\n%s", err)
	}
	return db
}
