package endpoint

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/app"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/transport/handler/http/request"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/transport/handler/http/response"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/usecase"
)

type (
	EndpointInterface interface {
		CreateOrganization(ctx context.Context, objDatsa request.Organization) (interface{}, error)
		GetOrganizations(ctx context.Context, params request.QueryParam) ([]response.Organization, int, error)
	}

	Endpoint struct {
		app     *app.App
		usecase usecase.UsecaseInterface
	}
)

func Init(app *app.App, usecase usecase.UsecaseInterface) EndpointInterface {
	return &Endpoint{
		app:     app,
		usecase: usecase,
	}
}
