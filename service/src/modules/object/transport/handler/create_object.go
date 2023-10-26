package handler

import (
	"encoding/json"
	"net/http"

	helperresponse "github.com/jabardigitalservice/super-app-services/event/src/helper/http/response"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/entity"
	httpresponse "github.com/jabardigitalservice/super-app-services/event/src/response/http"
)

func (h *Handler) CreateObject(w http.ResponseWriter, r *http.Request) {
	var obj entity.Object
	if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a context for the request
	ctx := r.Context()

	createdObj, err := h.endpoint.CreateObject(ctx, obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helperresponse.Render(w, r, h.responseMapping, httpresponse.Created, createdObj, nil)
}
