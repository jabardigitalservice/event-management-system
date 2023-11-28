package endpoint

import (
	"context"

	"github.com/google/uuid"
	"github.com/jabardigitalservice/super-app-services/event/src/constant"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/response"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/usecase"
	"github.com/jinzhu/copier"
)

func (e *Endpoint) GetObjectByID(ctx context.Context, id *uuid.UUID) (interface{}, error) {
	object, err := e.usecase.GetObjectByID(ctx, id)
	if err != nil {
		e.Log(ctx, constant.LogCategoryUsecase).Error(usecase.MethodGetObjectByID, err)
		return nil, err
	}

	responseObj := &response.Object{}

	if err := copier.Copy(responseObj, object); err != nil {
		e.Log(ctx, constant.LogCategoryUsecase).Error(usecase.MethodGetObjectByID, err)
		return nil, err
	}

	e.usecase.Log(ctx, constant.LogCategoryUsecase).Success(usecase.MethodGetObjectByID, "success")
	return responseObj, nil
}
