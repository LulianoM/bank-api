package config

import (
	"context"
	"errors"
	"log"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	ServerPort   string `env:"SERVER_PORT"`
	DatabaseEnvs DatabaseConfig
}

type DatabaseConfig struct {
	DatabaseName     string `env:"DATABASE_NAME"`
	DatabaseHost     string `env:"DATABASE_HOST"`
	DatabasePort     string `env:"DATABASE_PORT"`
	DatabaseUser     string `env:"DATABASE_USER"`
	DatabasePassword string `env:"DATABASE_PASSWORD"`
}

var globalConfig Config

func InitConfigs() {
	ctx := context.Background()
	var config Config

	if err := envconfig.Process(ctx, &config); err != nil {
		log.Fatal(err)
	}

	globalConfig = config
}

func init() {
	InitConfigs()
}

func GetConfig() Config {
	return globalConfig
}

func SetConfig(cfg Config) Config {
	globalConfig = cfg
	return globalConfig
}

func (this Config) Validate() error {
	if this.DatabaseEnvs.DatabaseName == "" || this.DatabaseEnvs.DatabaseHost == "" || this.DatabaseEnvs.DatabasePort == "" {
		return errors.New("cannot initilize with empty database environment variable")
	}

	if this.DatabaseEnvs.DatabaseUser == "" || this.DatabaseEnvs.DatabasePassword == "" {
		return errors.New("cannot initilize with empty database credentials environment variable")
	}

	return nil
}
