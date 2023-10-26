package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	_errors "github.com/jabardigitalservice/super-app-services/event/src/error"
	helperresponse "github.com/jabardigitalservice/super-app-services/event/src/helper/http/response"
	httpresponse "github.com/jabardigitalservice/super-app-services/event/src/response/http"
)

func (h *Handler) GetObjectByID(w http.ResponseWriter, r *http.Request) {
	// Parse and validate the object ID from the request
	idStr := chi.URLParam(r, "id")
	objectID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	// Create a context for the request
	ctx := r.Context()

	// Call the endpoint to get the object by ID
	obj, err := h.endpoint.GetObjectByID(ctx, objectID)
	if err != nil {
		// Check if the error is related to the object not being found
		if errors.Is(err, _errors.ErrNotFound) {
			// Object not found, respond with a 404 status
			http.Error(w, "Object not found", http.StatusNotFound)
			return
		}

		// Handle other internal server errors with a 500 status
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Serialize and send the object as an HTTP response
	helperresponse.Render(w, r, h.responseMapping, httpresponse.Success, obj, nil)
}
