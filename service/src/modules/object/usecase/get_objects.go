package usecase

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/entity"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/request"
)

func (uc *Usecase) GetObjects(ctx context.Context, params request.QueryParam) ([]entity.Object, error) {
	objects, err := uc.repo.GetObjects(ctx, params)
	if err != nil {
		return nil, err
	}

	return objects, nil
}