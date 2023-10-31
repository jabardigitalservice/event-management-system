package endpoint

import (
	"context"
	"log"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/request"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/response"
	"github.com/jinzhu/copier"
)

func (e *Endpoint) GetObjects(ctx context.Context, params request.QueryParam) ([]response.Object, error) {
	log.Printf("GetObjects - Received parameters: %#v", params)
	objects, err := e.usecase.GetObjects(ctx, params)
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

	log.Printf("GetObjects Response: %#v", responseObjects)
	return responseObjects, nil
}
