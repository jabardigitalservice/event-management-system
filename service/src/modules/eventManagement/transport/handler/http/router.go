package http

import (
	"github.com/go-chi/chi"
	"github.com/jabardigitalservice/golog/http/middleware"
	gologlogger "github.com/jabardigitalservice/golog/logger"
	"github.com/jabardigitalservice/super-app-services/event/src/app"
	"github.com/jabardigitalservice/super-app-services/event/src/constant"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/eventManagement/endpoint"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/eventManagement/transport/handler"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func Init(app *app.App, endpoint endpoint.EndpointInterface) *chi.Mux {

	var (
		router = chi.NewRouter()
		h      = handler.InitHandler(app, endpoint)
	)

	router.Use(middleware.Logger(app.GetLogger(), &gologlogger.LoggerData{
		Service: constant.ServiceName,
		Module:  constant.ModuleNameEventManagement,
		Version: app.GetVersion(),
	}, false))

	router.Post(newrelic.WrapHandleFunc(app.GetNewRelic(), "/organization", h.CreateOrganizations))
	router.Post(newrelic.WrapHandleFunc(app.GetNewRelic(), "/object", h.CreateObject))
	router.Get(newrelic.WrapHandleFunc(app.GetNewRelic(), "/object", h.GetObjects))
	router.Get(newrelic.WrapHandleFunc(app.GetNewRelic(), "/object/{id}", h.GetObjectByID))
	router.Put(newrelic.WrapHandleFunc(app.GetNewRelic(), "/object/{id}", h.UpdateObject))
	router.Delete(newrelic.WrapHandleFunc(app.GetNewRelic(), "/object/{id}", h.DeleteObject))
	return router
}
