package endpoint

import (
	"context"

	"github.com/google/uuid"
	"github.com/jabardigitalservice/super-app-services/event/src/constant"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/transport/handler/http/response"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/usecase"
	"github.com/jinzhu/copier"
)

func (e *Endpoint) GetOrganizationByID(ctx context.Context, id *uuid.UUID) (interface{}, error) {
	organization, err := e.usecase.GetOrganizationByID(ctx, id)
	if err != nil {
		e.Log(ctx, constant.LogCategoryUsecase).Error(usecase.MethodGetOrganizationByID, err)
		return nil, err
	}

	responseObj := &response.Organization{}

	if err := copier.Copy(responseObj, organization); err != nil {
		e.Log(ctx, constant.LogCategoryUsecase).Error(usecase.MethodGetOrganizationByID, err)
		return nil, err
	}

	e.usecase.Log(ctx, constant.LogCategoryUsecase).Success(usecase.MethodGetOrganizationByID, "success")
	return responseObj, nil
}
