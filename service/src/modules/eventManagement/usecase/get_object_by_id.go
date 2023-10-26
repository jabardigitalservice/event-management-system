package usecase

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/eventManagement/entity"
)

func (uc *Usecase) GetObjectByID(ctx context.Context, objectID uint64) (entity.Object, error) {
	obj, err := uc.repo.GetObjectByID(ctx, objectID)
	if err != nil {
		return entity.Object{}, err
	}

	return obj, nil
}
