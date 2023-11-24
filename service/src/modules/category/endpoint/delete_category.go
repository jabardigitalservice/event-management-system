package endpoint

import (
	"context"

	"github.com/google/uuid"
	_errors "github.com/jabardigitalservice/super-app-services/event/src/error"
)

func (e *Endpoint) DeleteCategory(ctx context.Context, id *uuid.UUID) error {
	err := e.usecase.DeleteCategory(ctx, id)

	if err != nil {
		return _errors.ErrInvalidStatus
	}

	return nil
}
