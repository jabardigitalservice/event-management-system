package http

import (
	"net/http"

	helperresponse "github.com/jabardigitalservice/super-app-services/event/src/helper/http/response"
)

var ResponseSuccess = &helperresponse.Response{
	HttpStatusCode: http.StatusOK,
	Status:         true,
	Error:          nil,
	Message: map[string]string{
		"en": "Success",
		"id": "Berhasil",
	},
}

var ResponseSuccessEmptyActiveOrders = &helperresponse.Response{
	HttpStatusCode: http.StatusOK,
	Status:         true,
	Error:          nil,
	Message: map[string]string{
		"en": "No active order was found",
		"id": "Tidak ada order aktif yang ditemukan",
	},
}

var ResponseCreated = &helperresponse.Response{
	HttpStatusCode: http.StatusCreated,
	Status:         true,
	Error:          nil,
	Message: map[string]string{
		"en": "Created",
		"id": "Berhasil menyimpan data",
	},
}
