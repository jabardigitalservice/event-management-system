package usecase

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/request"
)

func (uc *Usecase) UpdateObjectStatus(ctx context.Context, obj *request.Object) error {
	err := uc.repo.UpdateObjectStatus(ctx, obj, MethodUpdateObjectStatus)
	if err != nil {
		return err
	}
	return nil
}
