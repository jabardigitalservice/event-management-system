package usecase

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/app"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/entity"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/repository"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/transport/handler/http/request"
)

type (
	UsecaseInterface interface {
		CreateOrganization(ctx context.Context, obj entity.Organization) (*entity.Organization, error)
		GetOrganizations(ctx context.Context, params request.QueryParam) ([]entity.Organization, int, error)
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
