package request

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type (
	Category struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"  validate:"required"`
	}
	QueryParam struct {
		Page      uint64
		PageSize  uint64
		Offset    uint64
		SortBy    string
		SortOrder string
		StartDate string
		EndDate   string
		Keyword   string
		Filters   map[string]interface{}
	}
)

func GetRequestParam(r *http.Request) QueryParam {
	var query = r.URL.Query()

	page, _ := strconv.ParseInt(query.Get("page"), 10, 64)
	limit, _ := strconv.ParseInt(query.Get("pageSize"), 10, 64)
	sortOrder := strings.ToUpper(query.Get("sortOrder"))
	startDate := query.Get("startDate")
	endDate := query.Get("endDate")
	keyword := query.Get("q")

	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	if sortOrder != "ASC" {
		sortOrder = "DESC"
	}

	offset := (page - 1) * limit

	request := QueryParam{
		Page:      uint64(page),
		PageSize:  uint64(limit),
		Offset:    uint64(offset),
		SortOrder: sortOrder,
		StartDate: startDate,
		EndDate:   endDate,
		Keyword:   keyword,
		Filters:   make(map[string]interface{}),
	}

	return request
}
