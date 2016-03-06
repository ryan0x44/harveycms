// Copyright © 2016 Ryan D <ryan0x44.com>
package main

import (
	"fmt"
	"os"
)

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
