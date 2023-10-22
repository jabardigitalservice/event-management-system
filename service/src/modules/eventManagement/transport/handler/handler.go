package handler

import (
	"github.com/jabardigitalservice/super-app-services/event/src/app"
	helperresponse "github.com/jabardigitalservice/super-app-services/event/src/helper/http/response"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/eventManagement/endpoint"
	httpresponse "github.com/jabardigitalservice/super-app-services/event/src/response/http"
)

type Handler struct {
	app             *app.App
	endpoint        endpoint.EndpointInterface
	responseMapping map[string]*helperresponse.Response
}

func InitHandler(app *app.App, endpoint endpoint.EndpointInterface) *Handler {
	return &Handler{
		app:             app,
		endpoint:        endpoint,
		responseMapping: httpresponse.GetResponses(),
	}
}
