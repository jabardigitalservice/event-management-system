package app

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBSourceMaster                         string
	DBSourceSlave                          string
	DBDriver                               string
	NewRelicAppName                        string
	NewRelicLicense                        string
	NewRelicConfigAppLogForwardingEnabled  bool
	NewRelicConfigDistributedTracerEnabled bool
	NewRelicRecordPanics                   bool
}

func LoadConfig() (Config, error) {
	viper.AutomaticEnv()

	config := Config{
		DBSourceMaster:                         viper.GetString("DB_SOURCE_MASTER"),
		DBSourceSlave:                          viper.GetString("DB_SOURCE_SLAVE"),
		DBDriver:                               viper.GetString("DB_DRIVER"),
		NewRelicAppName:                        viper.GetString("NEW_RELIC_APP_NAME"),
		NewRelicLicense:                        viper.GetString("NEW_RELIC_LICENSE"),
		NewRelicConfigAppLogForwardingEnabled:  viper.GetBool("NEW_RELIC_CONFIG_APP_LOG_FORWARDING_ENABLED"),
		NewRelicConfigDistributedTracerEnabled: viper.GetBool("NEW_RELIC_CONFIG_DISTRIBUTED_TRACER_ENABLED"),
		NewRelicRecordPanics:                   viper.GetBool("NEW_RELIC_RECORD_PANICS"),
	}

	return config, nil
}
