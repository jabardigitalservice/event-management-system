package request

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type (
	SocialMedia struct {
		Name string `json:"name"`
		Link string `json:"link"`
	}

	Status string

	Object struct {
		ID          uuid.UUID      `json:"id"`
		Name        string         `json:"name" validate:"required"`
		Address     string         `json:"address" validate:"required"`
		Description string         `json:"description" validate:"required"`
		Banner      pq.StringArray `json:"banner"`
		Logo        string         `json:"logo"`
		SocialMedia []SocialMedia  `json:"social_media"`
		Organizer   string         `json:"organizer" validate:"required"`
		Status      Status         `json:"status" validate:"required"`
		CreatedAt   time.Time      `json:"created_at"`
		UpdatedAt   time.Time      `json:"updated_at"`
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

	if startDate == "" || endDate == "" {
		today := time.Now()
		dateFormat := "2006-01-02 15:04:05"
		beginningOfDay := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
		startDate = time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location()).Format(dateFormat)
		endDate = beginningOfDay.Add(24 * time.Hour).Add(-time.Second).Format(dateFormat)
	}

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
