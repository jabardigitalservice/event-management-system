package endpoint

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/app"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/ping/usecase"
)

type (
	EndpointInterface interface {
		GetHealthCheck(context.Context) (map[string]interface{}, error)
	}

	Endpoint struct {
		app     *app.App
		usecase usecase.UseCaseInterface
	}
)

func Init(app *app.App, usecase usecase.UseCaseInterface) EndpointInterface {
	return &Endpoint{
		app:     app,
		usecase: usecase,
	}
}
