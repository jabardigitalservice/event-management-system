package usecase

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/category/transport/handler/http/request"
)

func (uc *Usecase) Createcategory(ctx context.Context, cat request.Category) (*request.Category, error) {

	createdCat, err := uc.repo.Createcategory(ctx, cat)
	if err != nil {
		return nil, err
	}

	return createdCat, nil
}
