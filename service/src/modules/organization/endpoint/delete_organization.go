package endpoint

import (
	"context"

	"github.com/google/uuid"
)

func (e *Endpoint) DeleteOrganization(ctx context.Context, id *uuid.UUID) error {
	err := e.usecase.DeleteOrganization(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
