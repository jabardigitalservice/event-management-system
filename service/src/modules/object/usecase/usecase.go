package usecase

import (
	"context"
	"time"

	"github.com/jabardigitalservice/super-app-services/event/src/app"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/entity"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/repository"
)

type (
	UsecaseInterface interface {
		CreateObject(ctx context.Context, objData entity.Object) (uint64, entity.Object, time.Time, time.Time, error)
	}

	Usecase struct {
		app  *app.App
		repo repository.RepositoryInterface
	}
)

func Init(app *app.App, repo repository.RepositoryInterface) UsecaseInterface {
	return &Usecase{
		app:  app,
		repo: repo,
	}
}
