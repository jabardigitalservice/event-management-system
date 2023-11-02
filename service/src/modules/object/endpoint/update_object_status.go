package endpoint

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/request"
)

func (e *Endpoint) UpdateObjectStatus(ctx context.Context, obj *request.Object) error {
	err := e.usecase.UpdateObjectStatus(ctx, obj)
	if err != nil {
		return err
	}

	return nil
}
