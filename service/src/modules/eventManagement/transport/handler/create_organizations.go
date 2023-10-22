package handler

import (
	"encoding/json"
	"net/http"

	helperresponse "github.com/jabardigitalservice/super-app-services/event/src/helper/http/response"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/eventManagement/entity"
	httpresponse "github.com/jabardigitalservice/super-app-services/event/src/response/http"
)

func (h *Handler) CreateOrganizations(w http.ResponseWriter, r *http.Request) {
	var org entity.Organization
	if err := json.NewDecoder(r.Body).Decode(&org); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a context for the request
	ctx := r.Context()

	_, orgData, err := h.endpoint.CreateOrganizations(ctx, org)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with a success message using orgData
	helperresponse.Render(w, r, h.responseMapping, httpresponse.Created, orgData, nil)
}
