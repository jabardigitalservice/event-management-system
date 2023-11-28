package usecase

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/entity"
)

func (uc *Usecase) CreateOrganization(ctx context.Context, obj entity.Organization) (*entity.Organization, error) {
	createdObj, err := uc.repo.CreateOrganization(ctx, obj, MethodCreateOrganization)
	if err != nil {
		return nil, err
	}

	return createdObj, nil
}
