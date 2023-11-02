package endpoint

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/request"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/response"
	"github.com/jinzhu/copier"
)

func (e *Endpoint) GetObjects(ctx context.Context, params request.QueryParam) ([]response.Object, *int, error) {
	objects, count, err := e.usecase.GetObjects(ctx, params)
	if err != nil {
		return nil, nil, err
	}

	responseObjects := make([]response.Object, len(objects))

	for i, obj := range objects {
		responseObj := &response.Object{}
		if err := copier.Copy(responseObj, obj); err != nil {
			return nil, nil, err
		}
		responseObjects[i] = *responseObj
	}

	return responseObjects, count, nil
}
