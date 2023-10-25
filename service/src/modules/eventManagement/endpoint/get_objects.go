package endpoint

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/eventManagement/entity"
)

func (e *Endpoint) GetObjects(ctx context.Context, page int, perPage int) ([]entity.Object, error) {
	objects, err := e.usecase.GetObjects(ctx, page, perPage)
	if err != nil {
		return nil, err
	}

	return objects, nil
}