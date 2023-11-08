package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jabardigitalservice/super-app-services/event/src/app"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/entity"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/transport/handler/http/request"
)

type (
	RepositoryInterface interface {
		CreateOrganization(ctx context.Context, obj entity.Organization) (*entity.Organization, error)
		GetOrganizations(ctx context.Context, params request.QueryParam) ([]entity.Organization, error)
		CountFilteredOrganizations(ctx context.Context, params request.QueryParam) (int, error)
		GetOrganizationByID(ctx context.Context, id *uuid.UUID) (*entity.Organization, error)
		UpdateOrganization(ctx context.Context, obj *request.Organization) (*request.Organization, error)
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
