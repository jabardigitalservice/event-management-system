package endpoint

import (
	"context"
	"fmt"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/request"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/response"
	"github.com/jinzhu/copier"
)

func (e *Endpoint) CreateObject(ctx context.Context, objData request.Object) (*response.Object, error) {
	createdObj, err := e.usecase.CreateObject(ctx, objData)
	if err != nil {
		return nil, err
	}

	responseObj := &response.Object{}

	// Use copier.Copy to copy data from createdObj to responseObj
	if err := copier.Copy(responseObj, createdObj); err != nil {
		return nil, err
	}

	// Debug print
	fmt.Println(createdObj)

	return responseObj, nil
}
