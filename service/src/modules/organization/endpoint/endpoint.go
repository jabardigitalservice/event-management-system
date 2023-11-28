package endpoint

import (
	"context"

	"github.com/google/uuid"
	"github.com/jabardigitalservice/super-app-services/event/src/app"
	"github.com/jabardigitalservice/super-app-services/event/src/constant"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/transport/handler/http/request"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/transport/handler/http/response"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/usecase"
)

type (
	EndpointInterface interface {
		CreateOrganization(ctx context.Context, objDatsa request.Organization) (interface{}, error)
		GetOrganizations(ctx context.Context, params request.QueryParam) ([]response.Organization, int, error)
		GetOrganizationByID(ctx context.Context, id *uuid.UUID) (interface{}, error)
		UpdateOrganization(ctx context.Context, obj *request.Organization) (interface{}, error)
		DeleteOrganization(ctx context.Context, id *uuid.UUID) error
	}

	Endpoint struct {
		app     *app.App
		usecase usecase.UsecaseInterface
		logger  *app.AppLogger
	}
)

func Init(app *app.App, usecase usecase.UsecaseInterface) EndpointInterface {
	return &Endpoint{
		app:     app,
		usecase: usecase,
		logger:  app.GetAppLogger(),
	}
}

func (e *Endpoint) Log(ctx context.Context, category string) *app.AppLogger {
	return e.logger.Log(ctx, category, constant.ModuleNameOrganization)
}
