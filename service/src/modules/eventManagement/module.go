package eventManagement

import (
	"github.com/fazpass/goliath/v3/module"
	"github.com/go-chi/chi"
	"github.com/jabardigitalservice/super-app-services/event/src/app"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/eventManagement/endpoint"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/eventManagement/repository"
	transporthttp "github.com/jabardigitalservice/super-app-services/event/src/modules/eventManagement/transport/http"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/eventManagement/usecase"
)

type Module struct {
	usecase    usecase.UsecaseInterface
	endpoint   endpoint.EndpointInterface
	repository *repository.Repository
	httpRouter *chi.Mux
}

func Init(app *app.App) module.ModuleInterface {
	var (
		repository = repository.Init(app)
		usecase    = usecase.Init(app, repository)
		endpoint   = endpoint.Init(app, usecase)
		httpRouter = transporthttp.Init(app, endpoint)
	)

	return &Module{
		repository: repository,
		usecase:    usecase,
		endpoint:   endpoint,
		httpRouter: httpRouter,
	}
}

func (module *Module) GetHttpRouter() *chi.Mux {
	return module.httpRouter
}
