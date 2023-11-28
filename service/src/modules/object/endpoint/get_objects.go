package endpoint

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/constant"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/request"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/response"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/usecase"
	"github.com/jinzhu/copier"
)

func (e *Endpoint) GetObjects(ctx context.Context, params request.QueryParam) ([]response.Object, int, error) {
	objects, count, err := e.usecase.GetObjects(ctx, params)
	if err != nil {
		e.Log(ctx, constant.LogCategoryUsecase).Error(usecase.MethodGetObjects, err)
		return nil, 0, err
	}

	responseObjects := make([]response.Object, len(objects))

	for i, obj := range objects {
		responseObj := &response.Object{}
		if err := copier.Copy(responseObj, obj); err != nil {
			return nil, 0, err
		}
		responseObjects[i] = *responseObj
	}
	e.usecase.Log(ctx, constant.LogCategoryUsecase).Success(usecase.MethodGetObjects, "success")
	return responseObjects, count, nil
}
