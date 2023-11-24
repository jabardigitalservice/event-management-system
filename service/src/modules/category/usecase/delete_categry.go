package usecase

import (
	"context"

	"github.com/google/uuid"
)

func (uc *Usecase) DeleteCategory(ctx context.Context, id *uuid.UUID) error {
	err := uc.repo.DeleteCategory(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
