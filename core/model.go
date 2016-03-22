// Copyright © 2016 Ryan D <ryan0x44.com>
package core

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Route struct {
	ID          uint `gorm:"primary_key"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Path        string `sql:"type:text;not null;unique_index"`
	Module      string `sql:"size:255;not null"`
	EntityModel string `sql:"size:255"`
	EntityId    int
	Active      bool `sql:"not_null"`
}

type Setting struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Key       string `sql:"size:255;not null;unique_index"`
	Value     string `sql:"type:text"`
}

func DbInit(db *gorm.DB) {
	if !db.HasTable(&Route{}) {
		_ = db.CreateTable(&Route{})
	}
	if !db.HasTable(&Setting{}) {
		_ = db.CreateTable(&Setting{})
	}
}
