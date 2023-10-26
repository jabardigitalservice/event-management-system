package usecase

import (
	"context"
	"time"

	"github.com/jabardigitalservice/super-app-services/event/src/app"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/eventManagement/entity"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/eventManagement/repository"
)

type (
	UsecaseInterface interface {
		CreateOrganizations(ctx context.Context, orgData entity.Organization) (uint64, entity.Organization, error)
		CreateObject(ctx context.Context, objData entity.Object) (uint64, entity.Object, time.Time, time.Time, error)
		GetObjects(ctx context.Context, page int, perPage int) ([]entity.Object, error)
		GetObjectByID(ctx context.Context, objectID uint64) (entity.Object, error)
		UpdateObject(ctx context.Context, obj entity.Object) (entity.Object, error)
		DeleteObject(ctx context.Context, objectID uint64) error
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
