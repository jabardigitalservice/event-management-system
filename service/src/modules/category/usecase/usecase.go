package usecase

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/app"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/category/repository"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/category/transport/handler/http/request"
)

type (
	UsecaseInterface interface {
		Createcategory(ctx context.Context, cat request.Category) (*request.Category, error)
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
