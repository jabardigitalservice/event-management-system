// repository/organization_repository.go
package repository

import (
	"context"
	"time"

	"github.com/jabardigitalservice/super-app-services/event/src/app"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/eventManagement/entity"
)

type (
	RepositoryInterface interface {
		CreateOrganization(ctx context.Context, org entity.Organization) (uint64, error)
		CreateObject(context.Context, entity.Object) (uint64, time.Time, time.Time, error)
		GetObjects(ctx context.Context, page int, perPage int) ([]entity.Object, error)
		GetObjectByID(ctx context.Context, id uint64) (entity.Object, error)
		UpdateObject(ctx context.Context, obj entity.Object) (entity.Object, error)
		DeleteObject(ctx context.Context, objectID uint64) error
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
