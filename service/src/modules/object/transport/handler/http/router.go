package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jabardigitalservice/golog/http/middleware"
	gologlogger "github.com/jabardigitalservice/golog/logger"
	"github.com/jabardigitalservice/super-app-services/event/src/app"
	"github.com/jabardigitalservice/super-app-services/event/src/constant"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/endpoint"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func NewRelicMiddleware(app *app.App, handlerFunc func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	_, wrappedHandler := newrelic.WrapHandleFunc(app.GetNewRelic().Application, "object", handlerFunc)
	return wrappedHandler
}

func Init(app *app.App, endpoint endpoint.EndpointInterface) *chi.Mux {
	var (
		router = chi.NewRouter()
		h      = handler.InitHandler(app, endpoint)
	)

	router.Use(middleware.Logger(app.GetLogger(), &gologlogger.LoggerData{
		Service: constant.ServiceName,
		Module:  constant.ModuleNameobject,
		Version: app.GetVersion(),
	}, false))

	router.Post("/", NewRelicMiddleware(app, h.CreateObject))
	router.Get("/", NewRelicMiddleware(app, h.GetObjects))
	router.Get("/{id}", NewRelicMiddleware(app, h.GetObjectByID))
	router.Put("/{id}", NewRelicMiddleware(app, h.UpdateObject))
	router.Patch("/{id}", NewRelicMiddleware(app, h.UpdateObjectStatus))
	router.Delete("/{id}", NewRelicMiddleware(app, h.DeleteObject))

	return router
}
