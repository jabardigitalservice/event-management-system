package usecase

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/category/entity"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/category/transport/handler/http/request"
)

func (uc *Usecase) Getcategories(ctx context.Context, params request.QueryParam) ([]entity.Category, int, error) {
	categories, err := uc.repo.Getcategories(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	count, err := uc.repo.CountFilteredCategories(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	return categories, count, nil
}
