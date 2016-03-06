// Copyright © 2016 Ryan D <ryan0x44.com>
package main

import (
	"fmt"
	"github.com/ryan0x44/harveycms/cli/cmd"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
