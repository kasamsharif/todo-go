package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// DBConfig struct defininf server and db
type DBConfig struct {
	Server   string
	Database string
}

func (dbc *DBConfig) Read() {
	if _, err := toml.DecodeFile("config.toml", &dbc); err != nil {
		log.Fatal(err)
	}
}
