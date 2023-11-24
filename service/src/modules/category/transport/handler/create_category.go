package handler

import (
	"encoding/json"
	"net/http"

	helperresponse "github.com/jabardigitalservice/super-app-services/event/src/helper/http/response"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/category/transport/handler/http/request"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/category/transport/handler/http/response"
	httpresponse "github.com/jabardigitalservice/super-app-services/event/src/response/http"
)

func (h *Handler) Createcategory(w http.ResponseWriter, r *http.Request) {

	var respCat response.Category
	if err := json.NewDecoder(r.Body).Decode(&respCat); err != nil {
		RenderError(w, r, h, err, nil)
		return
	}

	reqCatJSON, err := json.Marshal(respCat)
	if err != nil {
		RenderError(w, r, h, err, nil)
		return
	}

	var reqCat request.Category
	if err := json.Unmarshal(reqCatJSON, &reqCat); err != nil {
		RenderError(w, r, h, err, nil)
		return
	}

	ctx := r.Context()

	createdCat, err := h.endpoint.Createcategory(ctx, reqCat)
	if err != nil {
		RenderError(w, r, h, err, nil)
		return
	}
	helperresponse.Render(w, r, h.responseMapping, httpresponse.Created, createdCat, nil)
}
