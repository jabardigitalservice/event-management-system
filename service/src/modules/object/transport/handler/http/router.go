package http

import (
	"github.com/go-chi/chi"
	"github.com/jabardigitalservice/golog/http/middleware"
	gologlogger "github.com/jabardigitalservice/golog/logger"
	"github.com/jabardigitalservice/super-app-services/event/src/app"
	"github.com/jabardigitalservice/super-app-services/event/src/constant"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/endpoint"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler"
	"github.com/newrelic/go-agent/v3/newrelic"
)

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

	router.Post(newrelic.WrapHandleFunc(app.GetNewRelic(), "/", h.CreateObject))
	router.Get(newrelic.WrapHandleFunc(app.GetNewRelic(), "/", h.GetObjects))
	router.Get(newrelic.WrapHandleFunc(app.GetNewRelic(), "/{id}", h.GetObjectByID))
	router.Put(newrelic.WrapHandleFunc(app.GetNewRelic(), "/{id}", h.UpdateObject))
	router.Patch(newrelic.WrapHandleFunc(app.GetNewRelic(), "/{id}", h.UpdateObjectStatus))
	router.Delete(newrelic.WrapHandleFunc(app.GetNewRelic(), "/{id}", h.DeleteObject))

	return router
}
