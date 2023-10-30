package endpoint

import (
	"context"

	"github.com/fazpass/goliath/v3/helper/validator"
	_errors "github.com/jabardigitalservice/super-app-services/event/src/error"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/request"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/response"
	"github.com/jinzhu/copier"
)

func (e *Endpoint) CreateObject(ctx context.Context, objData request.Object) (interface{}, error) {
	var validates = validator.Validate(objData)

	if validates != nil {
		return validates, _errors.ErrPayloadValidation
	}

	createdObj, err := e.usecase.CreateObject(ctx, objData)
	if err != nil {
		return nil, err
	}

	responseObj := &response.Object{}

	if err := copier.Copy(responseObj, createdObj); err != nil {
		return nil, err
	}

	return responseObj, nil
}
