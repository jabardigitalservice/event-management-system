package endpoint

import (
	"context"

	"github.com/fazpass/goliath/v3/helper/validator"
	_errors "github.com/jabardigitalservice/super-app-services/event/src/error"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/category/transport/handler/http/request"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/category/transport/handler/http/response"
	"github.com/jinzhu/copier"
)

func (e *Endpoint) UpdateCategory(ctx context.Context, obj *request.Category) (interface{}, error) {
	var validates = validator.Validate(obj)

	if validates != nil {
		return validates, _errors.ErrPayloadValidation
	}

	updatedObj, err := e.usecase.UpdateCategory(ctx, obj)

	if err != nil {
		return nil, err
	}

	responseObj := &response.Category{}

	if err := copier.Copy(responseObj, updatedObj); err != nil {
		return nil, err
	}

	return responseObj, nil
}
