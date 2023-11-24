package endpoint

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/app"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/category/transport/handler/http/request"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/category/transport/handler/http/response"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/category/usecase"
)

type (
	EndpointInterface interface {
		Createcategory(ctx context.Context, catData request.Category) (interface{}, error)
		Getcategories(ctx context.Context, params request.QueryParam) ([]response.Category, int, error)
		UpdateCategory(ctx context.Context, obj *request.Category) (interface{}, error)
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
