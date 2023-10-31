// repository/organization_repository.go
package repository

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/app"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/entity"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/request"
)

type (
	RepositoryInterface interface {
		CreateObject(ctx context.Context, obj request.Object) (request.Object, error)
		GetObjects(ctx context.Context, params request.QueryParam) ([]entity.Object, error)
		UpdateObject(ctx context.Context, obj request.Object) (request.Object, error)
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
