package handler

import (
	"net/http"

	"github.com/jabardigitalservice/super-app-services/event/src/app"
	_error "github.com/jabardigitalservice/super-app-services/event/src/error"
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
	case _error.ErrPayloadValidation:
		helperresponse.Render(w, r, h.responseMapping, httpresponse.UnprocessableEntity, nil, result)
	case _error.ErrDB:
		helperresponse.Render(w, r, h.responseMapping, httpresponse.InternalServerErrorDB, nil, result)
	case _error.ErrInvalidToken:
		helperresponse.Render(w, r, h.responseMapping, httpresponse.Unauthorized, nil, result)
	case _error.ErrTokenExpired:
		helperresponse.Render(w, r, h.responseMapping, httpresponse.TokenExpired, nil, result)
	case _error.ErrNotFound:
		helperresponse.Render(w, r, h.responseMapping, httpresponse.NotFound, nil, result)
	case _error.ErrExceedMaxOrderQuantity:
		helperresponse.Render(w, r, h.responseMapping, httpresponse.ExceedMaxOrderQuantity, nil, result)
	case _error.ErrExceedDailyHourLimit:
		helperresponse.Render(w, r, h.responseMapping, httpresponse.ExceedDailyHourLimit, nil, result)
	case _error.ErrExceedAllowedOrderDays:
		helperresponse.Render(w, r, h.responseMapping, httpresponse.ExceedAllowedOrderDays, nil, result)
	case _error.ErrClosingDay:
		helperresponse.Render(w, r, h.responseMapping, httpresponse.ClosingDay, nil, result)
	case _error.ErrOrderBeforeToday:
		helperresponse.Render(w, r, h.responseMapping, httpresponse.OrderBeforeToday, nil, result)
	case _error.ErrOrderNotFound:
		helperresponse.Render(w, r, h.responseMapping, httpresponse.OrderNotFound, nil, result)
	case _error.ErrOrderTicketNotFound:
		helperresponse.Render(w, r, h.responseMapping, httpresponse.OrderTicketNotFound, nil, result)
	case _error.ErrExternal:
		helperresponse.Render(w, r, h.responseMapping, httpresponse.ExternalServiceError, nil, result)
	case _error.ErrInvalidPaymentMethod:
		helperresponse.Render(w, r, h.responseMapping, httpresponse.InvalidPaymentMethod, nil, result)
	case _error.ErrNotImplemented:
		helperresponse.Render(w, r, h.responseMapping, httpresponse.NotImplemented, nil, result)
	case _error.ErrEmptyUserName:
		helperresponse.Render(w, r, h.responseMapping, httpresponse.EmptyUserName, nil, result)
	case _error.ErrInvalidConfig:
		helperresponse.Render(w, r, h.responseMapping, httpresponse.InternalServerError, nil, result)
	case _error.ErrInvalidQRString:
		helperresponse.Render(w, r, h.responseMapping, httpresponse.InvalidQRString, nil, result)
	default:
		helperresponse.Render(w, r, h.responseMapping, httpresponse.BadRequest, nil, nil)
	}
}
