package endpoint

import (
	"context"

	"github.com/fazpass/goliath/v3/helper/validator"
	"github.com/jabardigitalservice/super-app-services/event/src/constant"
	_errors "github.com/jabardigitalservice/super-app-services/event/src/error"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/entity"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/transport/handler/http/request"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/transport/handler/http/response"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/usecase"
	"github.com/jinzhu/copier"
)

func (e *Endpoint) CreateOrganization(ctx context.Context, objData request.Organization) (interface{}, error) {
	if err := validator.Validate(objData); err != nil {
		e.usecase.Log(ctx, constant.LogCategoryUsecase).Error(usecase.MethodCreateOrganization, _errors.ErrPayloadValidation)
		return nil, _errors.ErrPayloadValidation
	}

	entityObj := entity.Organization{}

	if err := copier.Copy(&entityObj, objData); err != nil {
		e.usecase.Log(ctx, constant.LogCategoryUsecase).Error(usecase.MethodCreateOrganization, err)
		return nil, err
	}

	createdObj, err := e.usecase.CreateOrganization(ctx, entityObj)
	if err != nil {
		e.usecase.Log(ctx, constant.LogCategoryUsecase).Error(usecase.MethodCreateOrganization, err)
		return nil, err
	}

	responseObj := &response.Organization{}
	if err := copier.Copy(responseObj, createdObj); err != nil {
		e.usecase.Log(ctx, constant.LogCategoryUsecase).Error(usecase.MethodCreateOrganization, err)
		return nil, err
	}

	e.usecase.Log(ctx, constant.LogCategoryUsecase).Success(usecase.MethodCreateOrganization, "success")
	return responseObj, nil
}
