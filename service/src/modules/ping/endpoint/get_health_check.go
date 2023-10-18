package endpoint

import (
	"context"
)

func (ep *Endpoint) GetHealthCheck(ctx context.Context) (map[string]interface{}, error) {
	result, err := ep.usecase.GetHealthCheck(ctx)
	return result, err
}
