package handler

import (
	"net/http"

	helperresponse "github.com/jabardigitalservice/super-app-services/event/src/helper/http/response"
	httpresponse "github.com/jabardigitalservice/super-app-services/event/src/response/http"
)

func (h *Handler) GetHealthCheck(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	var result interface{}
	result, err := h.endpoint.GetHealthCheck(ctx)
	if err != nil {
		helperresponse.Render(w, r, h.responseMapping, httpresponse.BadRequest, nil, nil)
		return
	}

	helperresponse.Render(w, r, h.responseMapping, httpresponse.Success, result, nil)
}
