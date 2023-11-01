package handler

import (
	"net/http"

	"github.com/jabardigitalservice/super-app-services/event/src/app"
	_errors "github.com/jabardigitalservice/super-app-services/event/src/error"
	helperresponse "github.com/jabardigitalservice/super-app-services/event/src/helper/http/response"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/endpoint"
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

func RenderError(w http.ResponseWriter, r *http.Request, h *Handler, err error, result interface{}) {
	switch err {
	case _errors.ErrPayloadValidation:
		helperresponse.Render(w, r, h.responseMapping, httpresponse.UnprocessableEntity, nil, result)
	case _errors.ErrDB:
		helperresponse.Render(w, r, h.responseMapping, httpresponse.InternalServerErrorDB, nil, result)
	case _errors.ErrNotFound:
		helperresponse.Render(w, r, h.responseMapping, httpresponse.NotFound, nil, result)
	default:
		helperresponse.Render(w, r, h.responseMapping, httpresponse.BadRequest, nil, nil)
	}
}
