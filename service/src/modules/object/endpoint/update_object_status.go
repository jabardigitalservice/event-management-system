package endpoint

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/constant"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/request"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/usecase"
)

func (e *Endpoint) UpdateObjectStatus(ctx context.Context, obj *request.Object) error {
	err := e.usecase.UpdateObjectStatus(ctx, obj)
	if err != nil {
		return err
	}

	e.usecase.Log(ctx, constant.LogCategoryUsecase).Success(usecase.MethodUpdateObjectStatus, "success")
	return nil
}
