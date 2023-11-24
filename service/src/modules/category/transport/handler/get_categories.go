package handler

import (
	"net/http"

	helperresponse "github.com/jabardigitalservice/super-app-services/event/src/helper/http/response"
	helperutils "github.com/jabardigitalservice/super-app-services/event/src/helper/utils"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/category/transport/handler/http/request"
	httpresponse "github.com/jabardigitalservice/super-app-services/event/src/response/http"
)

func (h *Handler) GetCategories(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queryParam := request.GetRequestParam(r)

	validSortOptions := []string{"name"}

	isOptionValid, _ := helperutils.InArray(queryParam.SortBy, validSortOptions)
	if !isOptionValid {
		queryParam.SortBy = ""
	}

	responseCategories, count, err := h.endpoint.Getcategories(ctx, queryParam)
	if err != nil {
		RenderError(w, r, h, err, responseCategories)
		return
	}

	totalData := uint64(count)

	pageSize := uint64(queryParam.PageSize)

	totalPages := (totalData + pageSize - 1) / pageSize

	page := uint64(queryParam.Page)

	response := map[string]interface{}{
		"data": responseCategories,
		"meta": map[string]interface{}{
			"page":        page,
			"page_size":   pageSize,
			"total_data":  totalData,
			"total_pages": totalPages,
		},
	}

	helperresponse.Render(w, r, h.responseMapping, httpresponse.Success, response, nil)
}
