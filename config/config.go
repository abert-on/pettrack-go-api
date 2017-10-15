package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Represents database server
type Config struct {
	Server   string
	Database string
}

// Read and parse a configuration file "config.toml"
func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
