package repository

import (
	"context"
	"errors"
)

var (
	ErrServiceUnavailable = errors.New("service unavailable")
)

func (r *Repository) GetHealthCheck(ctx context.Context) (map[string]interface{}, error) {
	result := map[string]interface{}{
		"status": "Service Available",
	}
	return result, nil
}
