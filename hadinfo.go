package main

import (
	"flag"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/gookit/ini/v2"
	"github.com/sirupsen/logrus"
	"hadInfo/api"
	"hadInfo/db"
	"time"
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

	logrus.SetLevel(logrus.TraceLevel) // TODO
	logrus.SetFormatter(&nested.Formatter{
		TimestampFormat: time.RFC3339,
	})

	db.Bootstrap()
	api.Bootstrap()

	defer db.Exit()
}
