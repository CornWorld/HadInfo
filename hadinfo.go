package main

import (
	"flag"
	"github.com/gookit/ini/v2"
	"hadInfo/api"
	"hadInfo/db"
)

var configFile string

const defaultConfigPath = "./hadInfo.ini"

func main() {
	flag.StringVar(&configFile, "config", defaultConfigPath, "config file path")

	flag.Parse()

	err := ini.LoadStrings(`
	[db]
	host = localhost
	port = 5432
	user = postgres
	password = hadInfo
	sslMode = disable
	name = hadInfo
	`) // Load default config
	if err != nil {
		panic(err)
	}

	if err := ini.LoadExists(configFile); err != nil {
		panic(err)
	}

	db.Bootstrap()
	api.Bootstrap()

	defer db.Exit()
}
