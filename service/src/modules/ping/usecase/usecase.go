package usecase

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/app"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/ping/repository"
)

type (
	UseCaseInterface interface {
		GetHealthCheck(context.Context) (map[string]interface{}, error)
	}

	UseCase struct {
		app  *app.App
		repo repository.RepositoryInterface
	}
)

func Init(app *app.App, repo repository.RepositoryInterface) UseCaseInterface {
	return &UseCase{
		app:  app,
		repo: repo,
	}
}
