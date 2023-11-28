package usecase

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/transport/handler/http/request"
)

func (uc *Usecase) UpdateOrganization(ctx context.Context, obj *request.Organization) (*request.Organization, error) {
	updatedObj, err := uc.repo.UpdateOrganization(ctx, obj, MethodUpdateOrganization)

	if err != nil {
		return nil, err
	}

	return updatedObj, nil
}
