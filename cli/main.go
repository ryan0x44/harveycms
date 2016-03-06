// Copyright © 2016 Ryan D <ryan0x44.com>
package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var MainCmd = &cobra.Command{
	Use:   "harvey-cli",
	Short: "Harvey CLI is a task runner for Harvey CMS",
	Long:  `Task runner for Harvey CMS <http://www.harveycms.org>`,
}

func main() {
	MainCmd.AddCommand(dbInitCmd)
	if err := MainCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
