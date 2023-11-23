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

func (e *Endpoint) CreateObject(ctx context.Context, objData request.Object) (interface{}, error) {
	endpointSegment := e.newrelic.StartSegment(ctx, "CreateObjectEndpoint")
	var validates = validator.Validate(objData)

	if validates != nil {
		e.Log(ctx, constant.LogCategoryUsecase).Error(usecase.MethodCreateObject, _errors.ErrPayloadValidation)
		return validates, _errors.ErrPayloadValidation
	}

	createdObj, err := e.usecase.CreateObject(ctx, objData)
	if err != nil {
		e.Log(ctx, constant.LogCategoryUsecase).Error(usecase.MethodCreateObject, err)
		return nil, err
	}

	responseObj := &response.Object{}

	if err := copier.Copy(responseObj, createdObj); err != nil {
		e.Log(ctx, constant.LogCategoryUsecase).Error(usecase.MethodCreateObject, err)
		return nil, err
	}

	defer endpointSegment.End()

	e.usecase.Log(ctx, constant.LogCategoryUsecase).Success(usecase.MethodCreateObject, "success")
	return responseObj, nil
}
