// Copyright © 2016 Ryan D <ryan0x44.com>
package core

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Route struct {
	gorm.Model
	Path   string `sql:"type:text;not null;unique_index"`
	Module string `sql:"size:255;not null"`
	Model  string `sql:"size:255"`
	Entity int
	Active bool `sql:"not_null"`
}

type Setting struct {
	gorm.Model
	Key   string `sql:"size:255;not null;unique_index"`
	Value string `sql:"type:text"`
}
