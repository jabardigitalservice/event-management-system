package app

import (
	"context"
	"database/sql"
	"log"

	"github.com/fazpass/goliath/v3/database"
	"github.com/spf13/viper"
)

func InitPgsqlMaster(ctx context.Context) *sql.DB {
	var pgsqlClient, err = database.Init(ctx, database.Config{
		Driver: viper.GetString("DB_DRIVER"),
		Source: viper.GetString("DB_SOURCE_MASTER"),
	})

	if err != nil {
		log.Panic(err)
	}

	return pgsqlClient.(*sql.DB)
}

func InitPgsqlSlave(ctx context.Context) *sql.DB {
	var pgsqlClient, err = database.Init(ctx, database.Config{
		Driver: viper.GetString("DB_DRIVER"),
		Source: viper.GetString("DB_SOURCE_SLAVE"),
	})

	if err != nil {
		log.Panic(err)
	}

	return pgsqlClient.(*sql.DB)
}
