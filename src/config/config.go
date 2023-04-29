package config

//unused

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
)

type Configuration struct {
	Port          string
	SQLDriver     string
	SQLsrc        string
	WindowsDBName string
	LinuxDBName   string
	MacDBName     string
	LogFile       string
	ErrFile       string
	AccessLogFile string
	Static        string
}

//go:embed config.json
var configFile []byte

func LoadConfig() *Configuration {
	var Config Configuration

	var err error

	err = json.Unmarshal(configFile, &Config)
	if err != nil {
		log.Fatalln("LoadConfig Unmarshal = ", err)
	}

	fmt.Println(Config)

	return &Config
}
