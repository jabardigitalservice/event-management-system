package endpoint

import (
	"context"

	"github.com/fazpass/goliath/v3/helper/validator"
	_errors "github.com/jabardigitalservice/super-app-services/event/src/error"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/entity"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/transport/handler/http/request"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/transport/handler/http/response"
	"github.com/jinzhu/copier"
)

func (e *Endpoint) CreateOrganization(ctx context.Context, objData request.Organization) (interface{}, error) {
	if err := validator.Validate(objData); err != nil {
		return nil, _errors.ErrPayloadValidation
	}

	entityObj := entity.Organization{}

	if err := copier.Copy(&entityObj, objData); err != nil {
		return nil, err
	}

	createdObj, err := e.usecase.CreateOrganization(ctx, entityObj)
	if err != nil {
		return nil, err
	}

	responseObj := &response.Organization{}
	if err := copier.Copy(responseObj, createdObj); err != nil {
		return nil, err
	}

	return responseObj, nil
}
