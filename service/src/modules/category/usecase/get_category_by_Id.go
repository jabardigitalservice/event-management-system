package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/category/entity"
)

func (uc *Usecase) GetcategoryByID(ctx context.Context, id *uuid.UUID) (*entity.Category, error) {
	return uc.repo.GetCategoryByID(ctx, id)
}
