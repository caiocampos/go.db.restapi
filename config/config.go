package config

import (
	"github.com/BurntSushi/toml"
)

type config struct {
	App app
	DB  database
}

type app struct {
	Port          int
	JsonProcessor string `toml:"json_processor"`
}

type database struct {
	Server   string
	Database string
}

// TOMLConfig represents the toml configuration file
var TOMLConfig *config

// ReadTOML method reads and parses the TOML configuration file
func ReadTOML() error {
	if TOMLConfig == nil {
		TOMLConfig = &config{}
		_, err := toml.DecodeFile("config.toml", &TOMLConfig)
		if err != nil {
			return err
		}
	}
	return nil
}
