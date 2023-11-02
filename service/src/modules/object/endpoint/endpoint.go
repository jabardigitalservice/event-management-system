package endpoint

import (
	"context"

	"github.com/google/uuid"
	"github.com/jabardigitalservice/super-app-services/event/src/app"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/entity"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/request"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/response"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/usecase"
)

type (
	EndpointInterface interface {
		CreateObject(ctx context.Context, objData request.Object) (interface{}, error)
		GetObjects(ctx context.Context, params request.QueryParam) ([]response.Object, *int, error)
		GetObjectByID(ctx context.Context, id *uuid.UUID) (*entity.Object, error)
		UpdateObject(ctx context.Context, obj *request.Object) (interface{}, error)
		DeleteObject(ctx context.Context, id *uuid.UUID) error
	}

	Endpoint struct {
		app     *app.App
		usecase usecase.UsecaseInterface
	}
)

func Init(app *app.App, usecase usecase.UsecaseInterface) EndpointInterface {
	return &Endpoint{
		app:     app,
		usecase: usecase,
	}
}
