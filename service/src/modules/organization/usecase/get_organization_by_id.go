package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/entity"
)

func (uc *Usecase) GetOrganizationByID(ctx context.Context, id *uuid.UUID) (*entity.Organization, error) {
	return uc.repo.GetOrganizationByID(ctx, id)
}
