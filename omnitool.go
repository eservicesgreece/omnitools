package main

import (
	"fmt"

	"github.com/spf13/viper"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {

	kingpin.Version(ver)
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.CommandLine.VersionFlag.Short('v')

	var appFlags = kingpin.Parse()

	setupConfig()

	switch appFlags {

	case account.FullCommand():
		if *accountPass != "" {
			fmt.Println(accountinfo(viper.GetString("general.baseURL"), *accountUname, *accountPass, *accountRInfo))
		} else {

			if viper.IsSet("accounts."+*accountUname+".pass") == false {
				fmt.Println("Account doesnt exist in omnitool.json")
			} else {
				fmt.Println(accountinfo(viper.GetString("general.baseURL"), *accountUname, viper.GetString("accounts."+*accountUname+".pass"), *accountRInfo))
			}
		}

	case config.FullCommand():
		dumpConfig()
	}
}
