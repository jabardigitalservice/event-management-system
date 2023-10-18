package http

import (
	"github.com/go-chi/chi"
	"github.com/jabardigitalservice/super-app-services/event/src/app"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/ping/endpoint"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/ping/transport/http/handler"
)

func Init(app *app.App, endpoint endpoint.EndpointInterface) *chi.Mux {
	var (
		router = chi.NewRouter()
		h      = handler.InitHandler(app, endpoint)
	)

	router.Get("/", h.GetHealthCheck)

	return router
}
