package usecase

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/eventManagement/entity"
)

func (uc *Usecase) UpdateObject(ctx context.Context, obj entity.Object) (entity.Object, error) {
	updatedObject, err := uc.repo.UpdateObject(ctx, obj)
	if err != nil {
		return entity.Object{}, err
	}

	return updatedObject, nil
}
