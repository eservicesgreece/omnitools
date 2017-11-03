package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func setupConfig() {
	viper.SetConfigName("omnitool")
	viper.SetConfigType("json")
	// Set Configuration File paths
	viper.AddConfigPath("/etc/omnitool/")
	viper.AddConfigPath("$HOME/.omnitool")
	viper.AddConfigPath(".")

	//Fetch Configuration
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config file not found, please create it under /etc/omnitool/omnitool.json or your $HOME")
		os.Exit(1)
	}
}
func dumpConfig() {
	for setting, value := range viper.AllSettings() {
		if checkIfIsMap(value) {
			fmt.Println(setting)
			for option, oValue := range value.(map[string]interface{}) {
				fmt.Println(option, "=", oValue)
			}
		} else {
			fmt.Println("Setting:", setting, "=", value)
		}
	}
}
