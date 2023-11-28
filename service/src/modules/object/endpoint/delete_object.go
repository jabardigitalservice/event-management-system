package endpoint

import (
	"context"

	"github.com/google/uuid"
	"github.com/jabardigitalservice/super-app-services/event/src/constant"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/usecase"
)

func (e *Endpoint) DeleteObject(ctx context.Context, id *uuid.UUID) error {
	err := e.usecase.DeleteObject(ctx, id)

	if err != nil {
		e.usecase.Log(ctx, constant.LogCategoryUsecase).Error(usecase.MethodDeleteObject, err)
		return err
	}

	e.usecase.Log(ctx, constant.LogCategoryUsecase).Success(usecase.MethodDeleteObject, "success")
	return nil
}
