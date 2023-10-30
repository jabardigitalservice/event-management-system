package endpoint

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/response"
	"github.com/jinzhu/copier"
)

func (e *Endpoint) GetObjects(ctx context.Context, page int, perPage int) ([]response.Object, error) {
	objects, err := e.usecase.GetObjects(ctx, page, perPage)
	if err != nil {
		return nil, err
	}

	responseObjects := make([]response.Object, len(objects))

	for i, obj := range objects {
		responseObj := &response.Object{}
		if err := copier.Copy(responseObj, obj); err != nil {
			return nil, err
		}
		responseObjects[i] = *responseObj
	}

	return responseObjects, nil
}
