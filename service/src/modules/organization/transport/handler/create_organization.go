package handler

import (
	"encoding/json"
	"net/http"

	helperresponse "github.com/jabardigitalservice/super-app-services/event/src/helper/http/response"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/transport/handler/http/request"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/transport/handler/http/response"
	httpresponse "github.com/jabardigitalservice/super-app-services/event/src/response/http"
)

func (h *Handler) CreateOrganization(w http.ResponseWriter, r *http.Request) {
	var respObj response.Organization
	if err := json.NewDecoder(r.Body).Decode(&respObj); err != nil {
		RenderError(w, r, h, err, nil)
		return
	}

	reqObjJSON, err := json.Marshal(respObj)
	if err != nil {
		RenderError(w, r, h, err, nil)
		return
	}

	var reqObj request.Organization
	if err := json.Unmarshal(reqObjJSON, &reqObj); err != nil {
		RenderError(w, r, h, err, nil)
		return
	}

	ctx := r.Context()

	createdObj, err := h.endpoint.CreateOrganization(ctx, reqObj)
	if err != nil {
		RenderError(w, r, h, err, nil)
		return
	}

	helperresponse.Render(w, r, h.responseMapping, httpresponse.Created, createdObj, nil)
}
