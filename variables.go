package main

import kingpin "gopkg.in/alecthomas/kingpin.v2"

var version string
var buildstamp string
var hash string
var ver = "esgtools v" + version + "\n Build: " + buildstamp + "\n GIT: " + hash + "\n Copyright (c) 2016-8 eServices Greece - https://eservices-greece.com\n"

var (
	app = kingpin.New("omnitool", ver)

	account      = kingpin.Command("account", "account")
	accountUname = account.Arg("accountuname", "Account ID").Required().String()
	accountPass  = account.Arg("password", "Account Password").String()
	accountRInfo = account.Flag("ri", "credit,tax").Short('r').String()

	config = kingpin.Command("configdump", "Dump all entries in configuration")
)
