// Copyright © 2016 Ryan D <ryan0x44.com>
package pages

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Page struct {
	gorm.Model
}

type PageVersion struct {
	gorm.Model
	Page        Page
	PageID      int
	Name        string `sql:"size:255;not null"`
	IsPublished bool   `sql:"not_null"`
	PublishedAt time.Time
}

func DbModel(db gorm.DB) {
	_ = db.Model(&PageVersion{}).Related(&Page{})
}

func DbInit(db gorm.DB) {
	_ = db.CreateTable(&Page{})
	_ = db.CreateTable(&PageVersion{})
}
