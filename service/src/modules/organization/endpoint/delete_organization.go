package endpoint

import (
	"context"

	"github.com/google/uuid"
	"github.com/jabardigitalservice/super-app-services/event/src/constant"
	_errors "github.com/jabardigitalservice/super-app-services/event/src/error"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/usecase"
)

func (e *Endpoint) DeleteOrganization(ctx context.Context, id *uuid.UUID) error {
	err := e.usecase.DeleteOrganization(ctx, id)

	if err != nil {
		e.usecase.Log(ctx, constant.LogCategoryUsecase).Error(usecase.MethodDeleteOrganization, err)
		return _errors.ErrInvalidStatus
	}

	e.usecase.Log(ctx, constant.LogCategoryUsecase).Success(usecase.MethodDeleteOrganization, "success")
	return nil
}
