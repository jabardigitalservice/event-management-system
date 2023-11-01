package endpoint

import (
	"context"

	"github.com/google/uuid"
)

func (e *Endpoint) DeleteObject(ctx context.Context, id *uuid.UUID) error {
	err := e.usecase.DeleteObject(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
