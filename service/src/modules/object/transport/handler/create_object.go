package handler

import (
	"encoding/json"
	"net/http"

	helperresponse "github.com/jabardigitalservice/super-app-services/event/src/helper/http/response"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/request"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/response"
	httpresponse "github.com/jabardigitalservice/super-app-services/event/src/response/http"
)

func (h *Handler) CreateObject(w http.ResponseWriter, r *http.Request) {
	var respObj response.Object
	if err := json.NewDecoder(r.Body).Decode(&respObj); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Encode the response object as JSON
	reqObjJSON, err := json.Marshal(respObj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Decode the JSON data into the request object
	var reqObj request.Object
	if err := json.Unmarshal(reqObjJSON, &reqObj); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a context for the request
	ctx := r.Context()

	createdObj, err := h.endpoint.CreateObject(ctx, reqObj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helperresponse.Render(w, r, h.responseMapping, httpresponse.Created, createdObj, nil)
}
