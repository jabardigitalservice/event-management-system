package endpoint

import (
	"context"

	"github.com/google/uuid"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/category/transport/handler/http/response"
	"github.com/jinzhu/copier"
)

func (e *Endpoint) GetcategoryByID(ctx context.Context, id *uuid.UUID) (interface{}, error) {
	cat, err := e.usecase.GetcategoryByID(ctx, id)
	if err != nil {
		return nil, err
	}

	responseObj := &response.Category{}

	if err := copier.Copy(responseObj, cat); err != nil {
		return nil, err
	}

	return responseObj, nil
}
