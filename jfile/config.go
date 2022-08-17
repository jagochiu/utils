package jfile

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

func GetConfig(configPath string) *ini.File {
	cfg, err := ini.LoadSources(ini.LoadOptions{IgnoreInlineComment: true}, configPath)
	if err != nil {
		log.Fatalf("[ERROR] Fail to parse config file %v", err)
	}
	return cfg
}

func GetGinConfig(path string) *ini.File {
	configPath := path + `debug.ini`
	if mode, okay := os.LookupEnv(`GIN_MODE`); okay && mode == `release` {
		configPath = path + `config.ini`
	}
	return GetConfig(configPath)
}
