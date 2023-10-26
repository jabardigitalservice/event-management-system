package usecase

import (
	"context"
)

func (uc *Usecase) DeleteObject(ctx context.Context, objectID uint64) error {
	return uc.repo.DeleteObject(ctx, objectID)
}
