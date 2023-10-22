package usecase

import (
	"context"

	_error "github.com/jabardigitalservice/super-app-services/event/src/error"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/eventManagement/entity"
)

func (uc *Usecase) CreateOrganizations(ctx context.Context, orgData entity.Organization) (uint64, entity.Organization, error) {
	if orgData.Name == "" || orgData.Email == "" {
		return 0, entity.Organization{}, _error.ErrInvalidInput
	}

	id, err := uc.repo.CreateOrganization(ctx, orgData)
	if err != nil {
		return 0, entity.Organization{}, err
	}

	return id, orgData, nil
}
