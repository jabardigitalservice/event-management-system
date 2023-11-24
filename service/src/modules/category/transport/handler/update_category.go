package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	helperresponse "github.com/jabardigitalservice/super-app-services/event/src/helper/http/response"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/category/transport/handler/http/request"
	httpresponse "github.com/jabardigitalservice/super-app-services/event/src/response/http"
)

func (h *Handler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		RenderError(w, r, h, err, nil)
		return
	}

	var obj request.Category
	if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {
		RenderError(w, r, h, err, nil)
		return
	}

	obj.ID = id

	ctx := r.Context()

	updatedObj, err := h.endpoint.UpdateCategory(ctx, &obj)
	if err != nil {
		log.Printf("error : %v", err)
		RenderError(w, r, h, err, nil)
		return
	}

	helperresponse.Render(w, r, h.responseMapping, httpresponse.Created, updatedObj, nil)
}
