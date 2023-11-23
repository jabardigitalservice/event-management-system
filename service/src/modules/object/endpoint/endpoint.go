package endpoint

import (
	"context"

	"github.com/google/uuid"
	"github.com/jabardigitalservice/super-app-services/event/src/app"
	"github.com/jabardigitalservice/super-app-services/event/src/constant"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/request"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/response"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/usecase"
)

type (
	EndpointInterface interface {
		CreateObject(ctx context.Context, objData request.Object) (interface{}, error)
		GetObjects(ctx context.Context, params request.QueryParam) ([]response.Object, int, error)
		GetObjectByID(ctx context.Context, id *uuid.UUID) (interface{}, error)
		UpdateObject(ctx context.Context, obj *request.Object) (interface{}, error)
		UpdateObjectStatus(ctx context.Context, obj *request.Object) error
		DeleteObject(ctx context.Context, id *uuid.UUID) error
	}

	Endpoint struct {
		app      *app.App
		usecase  usecase.UsecaseInterface
		newrelic *app.NewRelicManager
		logger   *app.AppLogger
	}
)

func Init(app *app.App, usecase usecase.UsecaseInterface) EndpointInterface {
	return &Endpoint{
		app:      app,
		usecase:  usecase,
		newrelic: app.GetNewRelic(),
		logger:   app.GetAppLogger(),
	}
}

func (e *Endpoint) Log(ctx context.Context, category string) *app.AppLogger {
	return e.logger.Log(ctx, category, constant.ModuleNameobject)
}
