// Copyright © 2016 Ryan D <ryan0x44.com>
package core

import (
	"time"
)

type Route struct {
	Path      string `sql:"type:text;not null;unique_index"`
	Module    string `sql:"size:255;not null"`
	Model     string `sql:"size:255"`
	Entity    int
	Active    bool `sql:"not_null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Setting struct {
	Key       string `sql:"size:255;not null;unique_index"`
	Value     string `sql:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
