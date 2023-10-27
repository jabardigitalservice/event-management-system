package endpoint

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/entity"
)

func (e *Endpoint) CreateObject(ctx context.Context, objData entity.Object) (*entity.Object, error) {
	createdObj, err := e.usecase.CreateObject(ctx, objData)
	if err != nil {
		return nil, err
	}

	return createdObj, nil
}
