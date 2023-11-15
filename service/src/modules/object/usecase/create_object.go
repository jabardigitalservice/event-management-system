package usecase

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/request"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func (uc *Usecase) CreateObject(ctx context.Context, obj request.Object) (*request.Object, error) {
	txn := newrelic.FromContext(ctx)
	usecaseSegment := txn.StartSegment("CreateObjectUsecase")

	createdObj, err := uc.repo.CreateObject(ctx, obj)
	if err != nil {
		return nil, err
	}
	usecaseSegment.End()
	return &createdObj, nil
}
