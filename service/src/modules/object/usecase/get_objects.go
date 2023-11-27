package usecase

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/entity"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/request"
)

func (uc *Usecase) GetObjects(ctx context.Context, params request.QueryParam) ([]entity.Object, int, error) {
	objects, err := uc.repo.GetObjects(ctx, params, MethodGetObjects)
	if err != nil {
		return nil, 0, err
	}

	count, err := uc.repo.CountFilteredObjects(ctx, params, MethodGetObjects)
	if err != nil {
		return nil, 0, err
	}

	return objects, count, nil
}
