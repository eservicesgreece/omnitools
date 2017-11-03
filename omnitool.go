package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/spf13/viper"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var version string
var buildstamp string
var hash string
var ver = "esgtools v" + version + "\n Build: " + buildstamp + "\n GIT: " + hash + "\n Copyright (c) 2016-8 eServices Greece - https://eservices-greece.com\n"

var (
	app = kingpin.New("omnitool", ver)

	account      = app.Command("account", "account")
	accountUname = account.Arg("accountuname", "Account ID").Required().String()
	accountPass  = account.Arg("password", "Account Password").String()
	accountRInfo = account.Flag("RI", "credit,tax").String()

	config = app.Command("configdump", "Dump all entries in configuration")
)

func checkIfIsMap(typeFC interface{}) bool {
	if reflect.ValueOf(typeFC).Kind() == reflect.Map {
		return true
	}
	return false
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

func setupConfig() {
	viper.SetConfigName("omnitool") // name of config file (without extension)
	viper.SetConfigType("json")     // Set type to json
	// Set Configuration File paths
	viper.AddConfigPath("/etc/omnitool/")
	viper.AddConfigPath("$HOME/.omnitool")
	viper.AddConfigPath(".")

	//Fetch Configuration
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		fmt.Println("Config file not found, please create it under /etc/omnitool/omnitool.json or your $HOME")
		os.Exit(1)
	}
}

func main() {

	kingpin.Version(ver)
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.CommandLine.VersionFlag.Short('v')

	setupConfig()

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {

	case account.FullCommand():
		if viper.IsSet("accounts."+*accountUname+".pass") == false {
			fmt.Println("Account doesnt exist in omnitool.json")
		} else {
			fmt.Println(accountinfo("https://www.omnivoice.eu/", *accountUname, viper.GetString("accounts."+*accountUname+".pass"), *accountRInfo))
		}

	case config.FullCommand():
		dumpConfig()
	}
}
