package usecase

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/request"
)

func (uc *Usecase) CreateObject(ctx context.Context, obj request.Object) (*request.Object, error) {
	createdObj, err := uc.repo.CreateObject(ctx, obj)
	if err != nil {
		return nil, err
	}

	return &createdObj, nil
}
