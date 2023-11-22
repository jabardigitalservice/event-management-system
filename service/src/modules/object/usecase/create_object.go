package usecase

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/constant"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/request"
)

func (uc *Usecase) CreateObject(ctx context.Context, obj request.Object) (*request.Object, error) {
	method := "CreateObjectUsecase"
	usecaseSegment := uc.newrelic.StartSegment(ctx, "CreateObjectUsecase")

	createdObj, err := uc.repo.CreateObject(ctx, obj)
	if err != nil {
		uc.Log(ctx, constant.LogCategoryApp).Error(method, err)
		return nil, err
	}
	defer usecaseSegment.End()
	return &createdObj, nil
}
