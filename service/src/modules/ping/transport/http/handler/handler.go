package handler

import (
	"github.com/jabardigitalservice/super-app-services/event/src/app"
	helperresponse "github.com/jabardigitalservice/super-app-services/event/src/helper/http/response"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/ping/endpoint"
	httpresponse "github.com/jabardigitalservice/super-app-services/event/src/response/http"
)

type Handler struct {
	app             *app.App
	endpoint        endpoint.EndpointInterface
	responseMapping map[string]*helperresponse.Response
}

func InitHandler(app *app.App, endpoint endpoint.EndpointInterface) *Handler {
	handler := &Handler{
		app:             app,
		endpoint:        endpoint,
		responseMapping: httpresponse.GetResponses(),
	}

	return handler
}
