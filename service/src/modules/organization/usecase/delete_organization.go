package usecase

import (
	"context"

	"github.com/google/uuid"
)

func (uc *Usecase) DeleteOrganization(ctx context.Context, id *uuid.UUID) error {
	err := uc.repo.DeleteOrganization(ctx, id, MethodDeleteOrganization)

	if err != nil {
		return err
	}

	return nil
}
