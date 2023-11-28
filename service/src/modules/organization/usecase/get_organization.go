package usecase

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/entity"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/transport/handler/http/request"
)

func (uc *Usecase) GetOrganizations(ctx context.Context, params request.QueryParam) ([]entity.Organization, int, error) {
	organizations, err := uc.repo.GetOrganizations(ctx, params, MethodGetOrganizations)
	if err != nil {
		return nil, 0, err
	}

	count, err := uc.repo.CountFilteredOrganizations(ctx, params, MethodGetOrganizations)
	if err != nil {
		return nil, 0, err
	}

	return organizations, count, nil
}
