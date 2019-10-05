package config

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
)

type ConfigList struct {
	ApiKey     string
	ApiSecret  string
	LogFile    string
	MockString string

	DbName    string
	SQLDriver string
	Port      int
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v")
		os.Exit(1)
	}

	Config = ConfigList{
		ApiKey:     cfg.Section("sofalib").Key("api_key").String(),
		ApiSecret:  cfg.Section("sofalib").Key("api_secret").String(),
		LogFile:    cfg.Section("go-3daudio").Key("log_file").String(),
		MockString: cfg.Section("mock").Key("mock_string").String(),
		DbName: cfg.Section("db").Key("name").String(),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		Port: cfg.Section("web").Key("port").MustInt(),
	}
}
