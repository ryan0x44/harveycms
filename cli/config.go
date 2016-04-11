// Copyright © 2016 Ryan D <ryan0x44.com>
package cli

import (
	"github.com/spf13/viper"
	"log"
)

func LoadConfig(rootDirectory string) {
	viper.SetConfigName("config")
	viper.AddConfigPath(rootDirectory)
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("Error reading config file '%s':\n%s\n", rootDirectory, err)
	}
}
