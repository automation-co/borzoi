package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const ConfigFileNotFound = `
Config file not found in the working directory.
You can create one by running:

$ borzoi generate

For more information, please visit
https://github.com/automation-co/borzoi`

func ReadConfig() map[string]interface{} {

	// Setup viper to read from the config file
	viper.SetConfigName("borzoi")

	// Config file type
	viper.SetConfigType("json")

	// Search for config in the following places:
	// 1. The current working directory
	viper.AddConfigPath(".")

	// Load the config file
	err := viper.ReadInConfig()
	if err != nil {

		// If the config file doesn't exist, create it
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {

			fmt.Println(ConfigFileNotFound)

			os.Exit(1)

		} else {
			panic(err)
		}
	}

	return viper.AllSettings()

}
