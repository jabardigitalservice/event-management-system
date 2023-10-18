package http

import (
	"net/http"

	helperresponse "github.com/jabardigitalservice/super-app-services/event/src/helper/http/response"
)

var ResponseInternalServerError = &helperresponse.Response{
	HttpStatusCode: http.StatusInternalServerError,
	Status:         false,
	Error:          nil,
	Message: map[string]string{
		"en": "Internal Server Error",
		"id": "Terdapat Kesalahan Internal",
	},
}

var ResponseInternalServerErrorDB = &helperresponse.Response{
	HttpStatusCode: http.StatusInternalServerError,
	Status:         false,
	Error:          nil,
	Message: map[string]string{
		"en": "Internal Server Error",
		"id": "Terdapat Kesalahan Internal",
	},
}

var ResponseInvalidFCMToken = &helperresponse.Response{
	HttpStatusCode: http.StatusInternalServerError,
	Status:         false,
	Error:          nil,
	Message: map[string]string{
		"en": "Invalid FCM Token",
		"id": "FCM Token Tidak Valid",
	},
}

var ResponseFailedCreateMessaging = &helperresponse.Response{
	HttpStatusCode: http.StatusInternalServerError,
	Status:         false,
	Error:          nil,
	Message: map[string]string{
		"en": "Create Message Notification Failed",
		"id": "Gagal Membuat Pesan Notifikasi",
	},
}

var ResponseFailedSendNotification = &helperresponse.Response{
	HttpStatusCode: http.StatusInternalServerError,
	Status:         false,
	Error:          nil,
	Message: map[string]string{
		"en": "Send Notification Failed",
		"id": "Gagal Mengirim Pesan",
	},
}
