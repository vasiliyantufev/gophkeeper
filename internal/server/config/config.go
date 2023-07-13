package config

import (
	"flag"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/sirupsen/logrus"
)

type Config struct {
	AddressGRPC string `env:"AddressGRPC"`
	AddressREST string `env:"AddressREST"`
	DSN         string `env:"DATABASE_DSN"`
	//MigrationsPath      string        `env:"ROOT_PATH"`
	DebugLevel          logrus.Level  `env:"DEBUG_LEVEL"`
	AccessTokenLifetime time.Duration `env:"ACCESS_TOKEN_LIFETIME"`
	FileFolder          string        `env:"DATA_FOLDER"`
}

// NewConfig - creates a new instance with the configuration for the server
func NewConfig(log *logrus.Logger) *Config {
	// Set default values
	configServer := Config{
		AddressGRPC: "localhost:8080",
		AddressREST: "localhost:8088",
		DSN:         "host=localhost port=5432 user=user password=password dbname=gophkeeper sslmode=disable",
		//MigrationsPath:      "file://./migrations",
		AccessTokenLifetime: 300 * time.Second,
		FileFolder:          "./data/server_keeper",
		DebugLevel:          logrus.DebugLevel,
	}

	flag.StringVar(&configServer.AddressGRPC, "g", configServer.AddressGRPC, "Server address")
	flag.StringVar(&configServer.DSN, "d", configServer.DSN, "Database configuration")
	flag.StringVar(&configServer.FileFolder, "f", configServer.FileFolder, "File Folder")
	flag.Parse()
	err := env.Parse(&configServer)
	if err != nil {
		log.Fatal(err)
	}
	log.Debug(configServer)

	return &configServer
}
