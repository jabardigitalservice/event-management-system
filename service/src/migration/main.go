package main

import (
	"log"

	"github.com/fazpass/goliath/v3/config"
	"github.com/fazpass/goliath/v3/database"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Panic(err)
	}

	db, err := database.InitPostgresql(viper.GetString("DB_DRIVER"), viper.GetString("DB_SOURCE_MASTER"))
	if err != nil {
		log.Panic(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(viper.GetString("DB_MIGRATE_PATH"), viper.GetString("DB_DRIVER"), driver)
	if err != nil {
		log.Panic(err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Panic(err)
	}
}
