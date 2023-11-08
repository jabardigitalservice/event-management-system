package endpoint

import (
	"context"

	"github.com/google/uuid"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/transport/handler/http/response"
	"github.com/jinzhu/copier"
)

func (e *Endpoint) GetOrganizationByID(ctx context.Context, id *uuid.UUID) (interface{}, error) {
	organization, err := e.usecase.GetOrganizationByID(ctx, id)
	if err != nil {
		return nil, err
	}

	responseObj := &response.Organization{}

	if err := copier.Copy(responseObj, organization); err != nil {
		return nil, err
	}

	return responseObj, nil
}
