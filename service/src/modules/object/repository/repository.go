// repository/organization_repository.go
package repository

import (
	"context"
	"time"

	"github.com/jabardigitalservice/super-app-services/event/src/app"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/entity"
)

type (
	RepositoryInterface interface {
		CreateObject(context.Context, entity.Object) (uint64, time.Time, time.Time, error)
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
