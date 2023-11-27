package usecase

import (
	"context"

	"github.com/google/uuid"
)

func (uc *Usecase) DeleteObject(ctx context.Context, id *uuid.UUID) error {
	err := uc.repo.DeleteObject(ctx, id, MethodDeleteObject)

	if err != nil {
		return err
	}

	return nil
}
