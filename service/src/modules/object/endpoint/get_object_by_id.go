package endpoint

import (
	"context"

	"github.com/google/uuid"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/entity"
)

func (e *Endpoint) GetObjectByID(ctx context.Context, id *uuid.UUID) (*entity.Object, error) {
	object, err := e.usecase.GetObjectByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return object, nil
}
