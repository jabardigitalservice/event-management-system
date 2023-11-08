package app

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBSourceMaster string
	DBSourceSlave  string
	DBDriver       string
	AppVersion     string
}

func LoadConfig() (Config, error) {
	viper.AutomaticEnv()

	var config Config

	config.DBSourceMaster = viper.GetString("DB_SOURCE_MASTER")
	config.DBSourceSlave = viper.GetString("DB_SOURCE_SLAVE")
	config.DBDriver = viper.GetString("DB_DRIVER")
	config.AppVersion = viper.GetString("APP_VERSION")

	return config, nil
}
