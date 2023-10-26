package endpoint

import (
	"context"
)

func (e *Endpoint) DeleteObject(ctx context.Context, id uint64) error {
	err := e.usecase.DeleteObject(ctx, id)
	return err
}
