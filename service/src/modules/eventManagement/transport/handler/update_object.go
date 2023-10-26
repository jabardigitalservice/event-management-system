package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi" // Import the chi router package for URL parameters
	_errors "github.com/jabardigitalservice/super-app-services/event/src/error"
	helperresponse "github.com/jabardigitalservice/super-app-services/event/src/helper/http/response"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/eventManagement/entity"
	httpresponse "github.com/jabardigitalservice/super-app-services/event/src/response/http"
)

func (h *Handler) UpdateObject(w http.ResponseWriter, r *http.Request) {
	// Parse the object ID from the URL parameter
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	// Parse the request body to obtain the object data to be updated
	var obj entity.Object
	if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set the ID of the object to be updated
	obj.ID = id

	// Create a context for the request
	ctx := r.Context()

	// Call the endpoint to update the object
	updatedObj, err := h.endpoint.UpdateObject(ctx, obj)
	if err != nil {
		// Check if the error is related to the object not being found
		if errors.Is(err, _errors.ErrNotFound) {
			// Object not found, respond with a 404 status
			http.Error(w, "Object not found", http.StatusNotFound)
			return
		}

		// Handle other internal server errors
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If the update was successful, you can respond with the updated object
	helperresponse.Render(w, r, h.responseMapping, httpresponse.Success, updatedObj, nil)
}
