package endpoint

import (
	"context"

	"github.com/google/uuid"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/response"
	"github.com/jinzhu/copier"
)

func (e *Endpoint) GetObjectByID(ctx context.Context, id *uuid.UUID) (interface{}, error) {
	object, err := e.usecase.GetObjectByID(ctx, id)
	if err != nil {
		return nil, err
	}

	responseObj := &response.Object{}

	if err := copier.Copy(responseObj, object); err != nil {
		return nil, err
	}

	return responseObj, nil
}
