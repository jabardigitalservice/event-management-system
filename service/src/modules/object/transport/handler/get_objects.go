package handler

import (
	"net/http"
	"strconv"

	helperresponse "github.com/jabardigitalservice/super-app-services/event/src/helper/http/response"
	httpresponse "github.com/jabardigitalservice/super-app-services/event/src/response/http"
)

func (h *Handler) GetObjects(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	perPageStr := r.URL.Query().Get("per_page")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		http.Error(w, "Invalid page parameter", http.StatusBadRequest)
		return
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil {
		http.Error(w, "Invalid per_page parameter", http.StatusBadRequest)
		return
	}
	ctx := r.Context()

	objects, err := h.endpoint.GetObjects(ctx, page, perPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helperresponse.Render(w, r, h.responseMapping, httpresponse.Success, objects, nil)
}
