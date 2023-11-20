package usecase

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/request"
)

func (uc *Usecase) CreateObject(ctx context.Context, obj request.Object) (*request.Object, error) {
	usecaseSegment := uc.app.GetNewRelic().StartSegment(ctx, "CreateObjectUsecase")

	createdObj, err := uc.repo.CreateObject(ctx, obj)
	if err != nil {
		return nil, err
	}
	defer usecaseSegment.End()
	return &createdObj, nil
}
