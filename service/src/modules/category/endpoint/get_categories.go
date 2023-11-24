package endpoint

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/category/transport/handler/http/request"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/category/transport/handler/http/response"
	"github.com/jinzhu/copier"
)

func (e *Endpoint) Getcategories(ctx context.Context, params request.QueryParam) ([]response.Category, int, error) {
	categories, count, err := e.usecase.Getcategories(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	responseCategories := make([]response.Category, len(categories))

	for i, cat := range categories {
		responseCat := &response.Category{}
		if err := copier.Copy(responseCat, cat); err != nil {
			return nil, 0, err
		}
		responseCategories[i] = *responseCat
	}

	return responseCategories, count, nil
}
