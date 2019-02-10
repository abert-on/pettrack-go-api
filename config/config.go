package config

import (
	"github.com/BurntSushi/toml"
)

// Config represents database server
type Config struct {
	Server   string
	Database string
}

// Read and parse a configuration file "config.toml"
func (c *Config) Read(filename string) error {
	_, err := toml.DecodeFile(filename, &c)
	return err
}
