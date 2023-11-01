package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	helperresponse "github.com/jabardigitalservice/super-app-services/event/src/helper/http/response"
	httpresponse "github.com/jabardigitalservice/super-app-services/event/src/response/http"
)

func (h *Handler) GetObjectByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		RenderError(w, r, h, err, nil)
		return
	}

	ctx := r.Context()

	object, err := h.endpoint.GetObjectByID(ctx, &id)
	if err != nil {
		RenderError(w, r, h, err, nil)
		return
	}

	helperresponse.Render(w, r, h.responseMapping, httpresponse.Success, object, nil)
}
