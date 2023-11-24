package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/jabardigitalservice/super-app-services/event/src/app"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/category/entity"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/category/repository"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/category/transport/handler/http/request"
)

type (
	UsecaseInterface interface {
		Createcategory(ctx context.Context, cat request.Category) (*request.Category, error)
		Getcategories(ctx context.Context, params request.QueryParam) ([]entity.Category, int, error)
		GetcategoryByID(ctx context.Context, id *uuid.UUID) (*entity.Category, error)
		UpdateCategory(ctx context.Context, obj *request.Category) (*request.Category, error)
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
