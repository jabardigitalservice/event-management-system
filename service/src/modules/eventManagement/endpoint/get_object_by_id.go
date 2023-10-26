package endpoint

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/eventManagement/entity"
)

func (e *Endpoint) GetObjectByID(ctx context.Context, objectID uint64) (entity.Object, error) {
	obj, err := e.usecase.GetObjectByID(ctx, objectID)
	if err != nil {
		return entity.Object{}, err
	}

	return obj, nil
}
