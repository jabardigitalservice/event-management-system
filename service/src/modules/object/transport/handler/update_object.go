package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	_errors "github.com/jabardigitalservice/super-app-services/event/src/error"
	helperresponse "github.com/jabardigitalservice/super-app-services/event/src/helper/http/response"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/request"
	httpresponse "github.com/jabardigitalservice/super-app-services/event/src/response/http"
)

func (h *Handler) UpdateObject(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	var obj request.Object
	if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	obj.ID = id

	ctx := r.Context()

	updatedObj, err := h.endpoint.UpdateObject(ctx, obj)
	if err != nil {
		if errors.Is(err, _errors.ErrNotFound) {
			http.Error(w, "Object not found", http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helperresponse.Render(w, r, h.responseMapping, httpresponse.Created, updatedObj, nil)
}
