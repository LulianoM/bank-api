package database

import (
	"context"
	"fmt"

	"github.com/LulianoM/bank-api/internal/config"
	"github.com/sethvargo/go-envconfig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect(config config.Config) error {

	ctx := context.Background()

	if err := envconfig.Process(ctx, &config); err != nil {
		return err
	}

	DB, err = gorm.Open(postgres.Open(getDSN(config)))
	if err != nil {
		return err
	}

	if err = AutoMigrate(); err != nil {
		return err
	}

	return nil
}

func AutoMigrate() error {
	DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	if err = DB.AutoMigrate(Tables...); err != nil {
		return err
	}
	return nil
}

func getDSN(config config.Config) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DatabaseEnvs.DatabaseHost,
		config.DatabaseEnvs.DatabaseUser,
		config.DatabaseEnvs.DatabasePassword,
		config.DatabaseEnvs.DatabaseName,
		config.DatabaseEnvs.DatabasePort,
	)
}
