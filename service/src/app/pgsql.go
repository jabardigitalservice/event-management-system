package app

import (
	"context"
	"database/sql"
	"log"

	"github.com/fazpass/goliath/v3/database"
)

func InitPgsqlMaster(ctx context.Context, config Config) *sql.DB {
	source := config.DBSourceMaster
	driver := config.DBDriver

	pgsqlClient, err := database.Init(ctx, database.Config{
		Driver: driver,
		Source: source,
	})

	if err != nil {
		log.Panic(err)
	}

	return pgsqlClient.(*sql.DB)
}

func InitPgsqlSlave(ctx context.Context, config Config) *sql.DB {
	source := config.DBSourceSlave
	driver := config.DBDriver

	pgsqlClient, err := database.Init(ctx, database.Config{
		Driver: driver,
		Source: source,
	})

	if err != nil {
		log.Panic(err)
	}

	return pgsqlClient.(*sql.DB)
}
