package usecase

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/request"
)

func (uc *Usecase) UpdateObject(ctx context.Context, obj *request.Object) (*request.Object, error) {
	updatedObj, err := uc.repo.UpdateObject(ctx, obj)

	if err != nil {
		return &request.Object{}, err
	}

	return updatedObj, nil
}
