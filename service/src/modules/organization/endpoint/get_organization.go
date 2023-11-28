package endpoint

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/constant"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/transport/handler/http/request"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/transport/handler/http/response"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/usecase"
	"github.com/jinzhu/copier"
)

func (e *Endpoint) GetOrganizations(ctx context.Context, params request.QueryParam) ([]response.Organization, int, error) {
	organizations, count, err := e.usecase.GetOrganizations(ctx, params)
	if err != nil {
		e.Log(ctx, constant.LogCategoryUsecase).Error(usecase.MethodGetOrganizations, err)
		return nil, 0, err
	}

	responseOrganizations := make([]response.Organization, len(organizations))

	for i, org := range organizations {
		responseOrg := &response.Organization{}
		if err := copier.Copy(responseOrg, org); err != nil {
			e.Log(ctx, constant.LogCategoryUsecase).Error(usecase.MethodGetOrganizations, err)
			return nil, 0, err
		}
		responseOrganizations[i] = *responseOrg
	}

	e.usecase.Log(ctx, constant.LogCategoryUsecase).Success(usecase.MethodGetOrganizations, "success")
	return responseOrganizations, count, nil
}
