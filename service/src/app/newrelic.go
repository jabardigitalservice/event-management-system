package app

import (
	"context"
	"log"

	"github.com/newrelic/go-agent/v3/newrelic"
)

type NewRelicManager struct {
	*newrelic.Application
}

func NewNewRelicManager(config Config) (*NewRelicManager, error) {
	var err error
	nrApp, err := newrelic.NewApplication(
		newrelic.ConfigAppName(config.NewRelicAppName),
		newrelic.ConfigLicense(config.NewRelicLicense),
		newrelic.ConfigAppLogForwardingEnabled(config.NewRelicConfigAppLogForwardingEnabled),
		newrelic.ConfigDistributedTracerEnabled(config.NewRelicConfigDistributedTracerEnabled),
		func(cfg *newrelic.Config) {
			cfg.ErrorCollector.RecordPanics = config.NewRelicRecordPanics
		},
	)

	if err != nil {
		return nil, err
	}

	return &NewRelicManager{Application: nrApp}, nil
}

func (nr *NewRelicManager) StartSegment(ctx context.Context, segmentName string) *newrelic.Segment {
	txn := newrelic.FromContext(ctx)
	if txn == nil {
		log.Println("No New Relic transaction found in the context. Creating a default transaction.")
		defaultTxn := nr.Application.StartTransaction("defaultTransaction")
		if defaultTxn == nil {
			log.Println("Failed to start a default New Relic transaction.")
			return nil
		}
		defer defaultTxn.End()
		txn = defaultTxn
	}

	segment := txn.StartSegment(segmentName)
	if segment == nil {
		log.Println("Failed to start New Relic segment.")
		return nil
	}

	return segment
}
