package handler

import (
	"net/http"

	helperresponse "github.com/jabardigitalservice/super-app-services/event/src/helper/http/response"
	helperutils "github.com/jabardigitalservice/super-app-services/event/src/helper/utils"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/request"
	httpresponse "github.com/jabardigitalservice/super-app-services/event/src/response/http"
)

func (h *Handler) GetObjects(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queryParam := request.GetRequestParam(r)

	sortByQuery := queryParam.SortBy

	validSortOptions := []string{"name", "address", "created_at", "updated_at"}

	isOptionValid, _ := helperutils.InArray(sortByQuery, validSortOptions)

	if !isOptionValid {
		queryParam.SortBy = ""
	}

	tmpResult, err := h.endpoint.GetObjects(ctx, queryParam)
	if err != nil {
		RenderError(w, r, h, err, tmpResult)
		return
	}

	result := tmpResult

	helperresponse.Render(w, r, h.responseMapping, httpresponse.Success, result, nil)
}
