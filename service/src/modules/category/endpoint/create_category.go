package endpoint

import (
	"context"

	"github.com/fazpass/goliath/v3/helper/validator"
	_errors "github.com/jabardigitalservice/super-app-services/event/src/error"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/category/transport/handler/http/request"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/category/transport/handler/http/response"
	"github.com/jinzhu/copier"
)

func (e *Endpoint) Createcategory(ctx context.Context, catData request.Category) (interface{}, error) {
	var validates = validator.Validate(catData)

	if validates != nil {
		return validates, _errors.ErrPayloadValidation
	}

	createdObj, err := e.usecase.Createcategory(ctx, catData)
	if err != nil {

		return nil, err
	}

	responseObj := &response.Category{}

	if err := copier.Copy(responseObj, createdObj); err != nil {

		return nil, err
	}

	return responseObj, nil
}
