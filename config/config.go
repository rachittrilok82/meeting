package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// These show the database connection
type Config struct {
	Server   string
	Database string
}

// These are use to read the credentials of the database
func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
