package endpoint

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func (e *Endpoint) StartSegment(ctx context.Context, segmentName string) *newrelic.Segment {
	return e.app.GetNewRelic().StartSegment(ctx, segmentName)
}
