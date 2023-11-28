package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jabardigitalservice/super-app-services/event/src/app"
	"github.com/jabardigitalservice/super-app-services/event/src/constant"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/entity"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/transport/handler/http/request"
)

type (
	RepositoryInterface interface {
		CreateOrganization(ctx context.Context, obj entity.Organization, method string) (*entity.Organization, error)
		GetOrganizations(ctx context.Context, params request.QueryParam, method string) ([]entity.Organization, error)
		CountFilteredOrganizations(ctx context.Context, params request.QueryParam, method string) (int, error)
		GetOrganizationByID(ctx context.Context, id *uuid.UUID, method string) (*entity.Organization, error)
		UpdateOrganization(ctx context.Context, obj *request.Organization, method string) (*request.Organization, error)
		DeleteOrganization(ctx context.Context, id *uuid.UUID, method string) error
	}
	Repository struct {
		app    *app.App
		db     *app.DB
		logger *app.AppLogger
	}
)

func Init(app *app.App) *Repository {
	return &Repository{
		app:    app,
		db:     app.GetDB(),
		logger: app.GetAppLogger(),
	}
}

func (r *Repository) Log(ctx context.Context) *app.AppLogger {
	return r.logger.Log(ctx, constant.LogCategoryApp, constant.ModuleNameOrganization)
}
