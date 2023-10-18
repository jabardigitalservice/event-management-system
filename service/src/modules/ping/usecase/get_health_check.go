package usecase

import (
	"context"
)

func (uc *UseCase) GetHealthCheck(ctx context.Context) (map[string]interface{}, error) {
	healthCheck, err := uc.repo.GetHealthCheck(ctx)
	if err != nil {
		return nil, err
	}

	return healthCheck, nil
}
