package endpoint

import (
	"context"

	"github.com/go-playground/validator/v10"
	_errors "github.com/jabardigitalservice/super-app-services/event/src/error"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/request"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/response"
	"github.com/jinzhu/copier"
)

func (e *Endpoint) CreateObject(ctx context.Context, objData request.Object) (interface{}, error) {
	// Create a new validator instance
	validate := validator.New()

	if err := validate.Struct(objData); err != nil {

		return nil, _errors.ErrPayloadValidation
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
