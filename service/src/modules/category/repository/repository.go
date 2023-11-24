// repository/organization_repository.go
package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jabardigitalservice/super-app-services/event/src/app"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/category/entity"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/category/transport/handler/http/request"
)

type (
	RepositoryInterface interface {
		Createcategory(ctx context.Context, obj request.Category) (*request.Category, error)
		Getcategories(ctx context.Context, params request.QueryParam) ([]entity.Category, error)
		CountFilteredCategories(ctx context.Context, params request.QueryParam) (int, error)
		GetCategoryByID(ctx context.Context, id *uuid.UUID) (*entity.Category, error)
		UpdateCategory(ctx context.Context, obj *request.Category) (*request.Category, error)
		DeleteCategory(ctx context.Context, id *uuid.UUID) error
	}
	Repository struct {
		app *app.App
		db  *app.DB
	}
)

func Init(app *app.App) *Repository {
	return &Repository{
		app: app,
		db:  app.GetDB(),
	}
}
