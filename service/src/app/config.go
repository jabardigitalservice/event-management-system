package app

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBSourceMaster string
	DBSourceSlave  string
	DBDriver       string
}

func LoadConfig() (Config, error) {
	viper.AutomaticEnv()

	var config Config

	config.DBSourceMaster = viper.GetString("DB_SOURCE_MASTER")
	config.DBSourceSlave = viper.GetString("DB_SOURCE_SLAVE")
	config.DBDriver = viper.GetString("DB_DRIVER")

	return config, nil
}
