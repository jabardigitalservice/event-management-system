package usecase

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/category/transport/handler/http/request"
)

func (uc *Usecase) UpdateCategory(ctx context.Context, obj *request.Category) (*request.Category, error) {
	updatedObj, err := uc.repo.UpdateCategory(ctx, obj)

	if err != nil {
		return nil, err
	}

	return updatedObj, nil
}
