package usecase

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/eventManagement/entity"
)

func (uc *Usecase) GetObjects(ctx context.Context, page int, perPage int) ([]entity.Object, error) {
	objects, err := uc.repo.GetObjects(ctx, page, perPage)
	if err != nil {
		return nil, err
	}

	return objects, nil
}
