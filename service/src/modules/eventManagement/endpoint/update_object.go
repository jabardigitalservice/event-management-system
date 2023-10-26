package endpoint

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/eventManagement/entity"
)

func (e *Endpoint) UpdateObject(ctx context.Context, obj entity.Object) (entity.Object, error) {
	updatedObject, err := e.usecase.UpdateObject(ctx, obj)
	if err != nil {
		return entity.Object{}, err
	}

	return updatedObject, nil
}
