package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi" // Import the chi router package for URL parameters
	helperresponse "github.com/jabardigitalservice/super-app-services/event/src/helper/http/response"
	httpresponse "github.com/jabardigitalservice/super-app-services/event/src/response/http"
)

func (h *Handler) DeleteObject(w http.ResponseWriter, r *http.Request) {
	// Parse the object ID from the URL parameter
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	// Create a context for the request
	ctx := r.Context()

	// Call the DeleteObject method in the endpoint to delete the object
	err = h.endpoint.DeleteObject(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If the deletion was successful, you can respond with a success message
	helperresponse.Render(w, r, h.responseMapping, httpresponse.Success, "Object deleted", nil)
}
