package main

import (
	"fmt"
	"github.com/ryan0x44/harveycms/core"
)

func main() {
	fmt.Println("cli app")
	fmt.Println("initialising database")
	dbInit()
}

func dbInit() {
	db := core.DbConnect("dev:password@/harveycms?charset=utf8&parseTime=True&loc=Local")
	db.CreateTable(&core.Route{})
}
