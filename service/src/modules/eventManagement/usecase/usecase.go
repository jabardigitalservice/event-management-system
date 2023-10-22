package usecase

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/app"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/eventManagement/entity"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/eventManagement/repository"
)

type (
	UsecaseInterface interface {
		CreateOrganizations(ctx context.Context, orgData entity.Organization) (uint64, entity.Organization, error)
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
