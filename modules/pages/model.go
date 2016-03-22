// Copyright © 2016 Ryan D <ryan0x44.com>
package pages

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Page struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Versions  []PageVersion
}

type PageVersion struct {
	ID          uint `gorm:"primary_key"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Page        Page
	PageId      uint
	Name        string `sql:"size:255;not null"`
	IsPublished bool   `sql:"not_null"`
	PublishedAt time.Time
}

func DbInit(db *gorm.DB) {
	if !db.HasTable(&Page{}) {
		db.CreateTable(&Page{})
	}
	if !db.HasTable(&PageVersion{}) {
		db.CreateTable(&PageVersion{})
		db.Model(&PageVersion{}).AddForeignKey("page_id", "pages(id)", "CASCADE", "RESTRICT")
		db.Model(&PageVersion{}).AddUniqueIndex("is_published_unique", "page_id", "is_published")
	}
}
