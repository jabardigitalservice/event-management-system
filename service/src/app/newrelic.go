package app

import (
	"github.com/newrelic/go-agent/v3/newrelic"
)

var nrApp *newrelic.Application

func InitNewRelic(config Config) error {
	var err error
	nrApp, err = newrelic.NewApplication(
		newrelic.ConfigAppName(config.NewRelicAppName),
		newrelic.ConfigLicense(config.NewRelicLicense),
		newrelic.ConfigAppLogForwardingEnabled(config.NewRelicConfigAppLogForwardingEnabled),
		newrelic.ConfigDistributedTracerEnabled(config.NewRelicConfigDistributedTracerEnabled),
		func(cfg *newrelic.Config) {
			cfg.ErrorCollector.RecordPanics = config.NewRelicRecordPanics
		},
	)

	return err
}
