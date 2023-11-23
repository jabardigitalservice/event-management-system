package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/jabardigitalservice/super-app-services/event/src/app"
	"github.com/jabardigitalservice/super-app-services/event/src/constant"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/entity"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/repository"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/request"
)

type (
	UsecaseInterface interface {
		CreateObject(ctx context.Context, obj request.Object) (*request.Object, error)
		GetObjects(ctx context.Context, params request.QueryParam) ([]entity.Object, int, error)
		GetObjectByID(ctx context.Context, id *uuid.UUID) (*entity.Object, error)
		UpdateObject(ctx context.Context, obj *request.Object) (*request.Object, error)
		UpdateObjectStatus(ctx context.Context, obj *request.Object) error
		DeleteObject(ctx context.Context, id *uuid.UUID) error
	}

	Usecase struct {
		app      *app.App
		repo     repository.RepositoryInterface
		newrelic *app.NewRelicManager
		logger   *app.AppLogger
	}
)

const (
	MethodCreateObject       = "object:create-object"
	MethodGetObjects         = "object:get-objects"
	MethodGetObjectByID      = "object:get-objects-by-id"
	MethodUpdateObject       = "object:update-object"
	MethodUpdateObjectStatus = "object:update-object-status"
	MethodDeleteObject       = "object:delete-object"
)

func Init(app *app.App, repo repository.RepositoryInterface) UsecaseInterface {
	return &Usecase{
		app:      app,
		repo:     repo,
		newrelic: app.GetNewRelic(),
		logger:   app.GetAppLogger(),
	}
}

func (uc *Usecase) Log(ctx context.Context, category string) *app.AppLogger {
	return uc.logger.Log(ctx, category, constant.ModuleNameobject)
}
