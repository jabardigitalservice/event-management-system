package endpoint

import (
	"context"

	"github.com/fazpass/goliath/v3/helper/validator"
	"github.com/jabardigitalservice/super-app-services/event/src/constant"
	_errors "github.com/jabardigitalservice/super-app-services/event/src/error"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/transport/handler/http/request"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/transport/handler/http/response"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/usecase"
	"github.com/jinzhu/copier"
)

func (e *Endpoint) UpdateOrganization(ctx context.Context, obj *request.Organization) (interface{}, error) {
	var validates = validator.Validate(obj)

	if validates != nil {
		e.usecase.Log(ctx, constant.LogCategoryUsecase).Error(usecase.MethodUpdateOrganization, _errors.ErrPayloadValidation)
		return validates, _errors.ErrPayloadValidation
	}
	updatedObj, err := e.usecase.UpdateOrganization(ctx, obj)

	if err != nil {
		e.usecase.Log(ctx, constant.LogCategoryUsecase).Error(usecase.MethodUpdateOrganization, err)
		return nil, err
	}

	responseObj := &response.Organization{}

	if err := copier.Copy(responseObj, updatedObj); err != nil {
		e.usecase.Log(ctx, constant.LogCategoryUsecase).Error(usecase.MethodUpdateOrganization, err)
		return nil, err
	}
	e.usecase.Log(ctx, constant.LogCategoryUsecase).Success(usecase.MethodUpdateOrganization, "success")
	return responseObj, nil
}
