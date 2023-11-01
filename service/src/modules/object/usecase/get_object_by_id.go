package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/entity"
)

func (uc *Usecase) GetObjectByID(ctx context.Context, id *uuid.UUID) (*entity.Object, error) {
	return uc.repo.GetObjectByID(ctx, id)
}
