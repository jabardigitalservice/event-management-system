package app

import (
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/spf13/viper"
)

var nrApp *newrelic.Application

func InitNewRelic() error {
	var err error
	nrApp, err = newrelic.NewApplication(
		newrelic.ConfigAppName(viper.GetString("NEW_RELIC_APP_NAME")),
		newrelic.ConfigLicense(viper.GetString("NEW_RELIC_LICENSE")),
		newrelic.ConfigAppLogForwardingEnabled(viper.GetBool("NEW_RELIC_CONFIG_APP_LOG_FORWARDING_ENABLED")),
		newrelic.ConfigDistributedTracerEnabled(viper.GetBool("NEW_RELIC_CONFIG_DISTRIBUTED_TRACER_ENABLED")),
		func(cfg *newrelic.Config) {
			cfg.ErrorCollector.RecordPanics = viper.GetBool("NEW_RELIC_RECORD_PANICS")
		},
	)

	return err
}
