package request

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

type (
	Organization struct {
		Id           uint64    `json:"id"`
		Name         string    `json:"name" required:"true"`
		Email        string    `json:"email" required:"true"`
		Address      string    `json:"address" required:"true"`
		Phone_number string    `json:"phone_number" required:"true"`
		Description  string    `json:"description" required:"true"`
		Logo         string    `json:"logo" required:"true"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
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
		Status    string
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
	status := query.Get("status")

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
		Status:    status,
		Filters:   make(map[string]interface{}),
	}

	return request
}
