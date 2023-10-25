package usecase

import (
	"context"
	"time"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/eventManagement/entity"
)

func (uc *Usecase) CreateObject(ctx context.Context, obj entity.Object) (uint64, entity.Object, time.Time, time.Time, error) {
	id, createdAt, updatedAt, err := uc.repo.CreateObject(ctx, obj)
	if err != nil {
		return 0, entity.Object{}, time.Time{}, time.Time{}, err
	}

	return id, obj, createdAt, updatedAt, nil
}
