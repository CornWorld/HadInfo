package main

import (
	"flag"
	"github.com/gookit/ini/v2"
	"hadInfo/api"
	"hadInfo/db"
)

var configFile string

const defaultConfigPath = "./hadInfo.ini"
const defaultConfigOptions = `
	[db]
	host = localhost
	port = 5432
	user = postgres
	password = hadInfo
	sslMode = disable
	name = hadInfo
	`

func main() {
	flag.StringVar(&configFile, "config", defaultConfigPath, "config file path")
	flag.Parse()

	if err := ini.LoadStrings(defaultConfigOptions); err != nil {
		panic(err)
	}

	if err := ini.LoadExists(configFile); err != nil {
		panic(err)
	}

	db.Bootstrap()
	api.Bootstrap()

	defer db.Exit()
}
