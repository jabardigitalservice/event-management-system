package app

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// Config holds the application configuration.
type Config struct {
	DBSourceMaster string
	DBSourceSlave  string
	DBDriver       string
}

// LoadConfig initializes and loads the application configuration into the Config struct.
func LoadConfig() (Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return Config{}, err
	}

	viper.AutomaticEnv() // Load environment variables

	var config Config

	// Use Viper to retrieve values from environment variables
	config.DBSourceMaster = viper.GetString("DB_SOURCE_MASTER")
	config.DBSourceSlave = viper.GetString("DB_SOURCE_SLAVE")
	config.DBDriver = viper.GetString("DB_DRIVER")

	return config, nil
}
