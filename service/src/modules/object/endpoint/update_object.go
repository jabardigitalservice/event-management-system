package endpoint

import (
	"context"

	"github.com/fazpass/goliath/v3/helper/validator"
	"github.com/jabardigitalservice/super-app-services/event/src/constant"
	_errors "github.com/jabardigitalservice/super-app-services/event/src/error"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/request"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/response"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/usecase"
	"github.com/jinzhu/copier"
)

func (e *Endpoint) UpdateObject(ctx context.Context, obj *request.Object) (interface{}, error) {
	var validates = validator.Validate(obj)

	if validates != nil {
		return validates, _errors.ErrPayloadValidation
	}
	updatedObj, err := e.usecase.UpdateObject(ctx, obj)

	if err != nil {
		return nil, err
	}

	responseObj := &response.Object{}

	if err := copier.Copy(responseObj, updatedObj); err != nil {
		return nil, err
	}

	e.usecase.Log(ctx, constant.LogCategoryUsecase).Success(usecase.MethodUpdateObject, "success")
	return responseObj, nil
}
