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

	validSortOptions := []string{"name", "address", "created_at", "updated_at"}

	isOptionValid, _ := helperutils.InArray(queryParam.SortBy, validSortOptions)
	if !isOptionValid {
		queryParam.SortBy = ""
	}

	responseObjects, count, err := h.endpoint.GetObjects(ctx, queryParam)
	if err != nil {
		RenderError(w, r, h, err, responseObjects)
		return
	}

	totalData := uint64(*count)

	pageSize := uint64(queryParam.PageSize)

	totalPages := (totalData + pageSize - 1) / pageSize

	page := uint64(queryParam.Page)

	startIndex := (page - 1) * pageSize
	endIndex := startIndex + pageSize

	if endIndex > totalData {
		endIndex = totalData
	}

	response := map[string]interface{}{
		"data": responseObjects[startIndex:endIndex],
		"meta": map[string]interface{}{
			"page":        page,
			"page_size":   pageSize,
			"total_data":  totalData,
			"total_pages": totalPages,
		},
	}

	helperresponse.Render(w, r, h.responseMapping, httpresponse.Success, response, nil)
}
