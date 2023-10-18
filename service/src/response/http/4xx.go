package http

import (
	"net/http"

	helperresponse "github.com/jabardigitalservice/super-app-services/event/src/helper/http/response"
)

var ResponseBadRequest = &helperresponse.Response{
	HttpStatusCode: http.StatusBadRequest,
	Status:         false,
	Error:          nil,
	Message: map[string]string{
		"en": "Bad Request",
		"id": "Terdapat Kesalahan Request",
	},
}

var ResponseBadRequestBody = &helperresponse.Response{
	HttpStatusCode: http.StatusBadRequest,
	Status:         false,
	Error:          nil,
	Message: map[string]string{
		"en": "Bad Request (Body)",
		"id": "Terdapat Kesalahan Request Body",
	},
}

var ResponseUnauthorized = &helperresponse.Response{
	HttpStatusCode: http.StatusUnauthorized,
	Status:         false,
	Error:          nil,
	Message: map[string]string{
		"en": "Unauthorized",
		"id": "Autorisasi Salah",
	},
}

var ResponseTokenExpired = &helperresponse.Response{
	HttpStatusCode: http.StatusUnauthorized,
	Status:         false,
	Error:          nil,
	Message: map[string]string{
		"en": "Token Expired",
		"id": "Token Kedaluwarsa",
	},
}

var ResponseNotFound = &helperresponse.Response{
	HttpStatusCode: http.StatusNotFound,
	Status:         false,
	Error:          nil,
	Message: map[string]string{
		"en": "Not found",
		"id": "Tidak ditemukan",
	},
}

var ResponseConflict = &helperresponse.Response{
	HttpStatusCode: http.StatusConflict,
	Status:         false,
	Error:          nil,
	Message: map[string]string{
		"en": "Conflict",
		"id": "Data konflik",
	},
}

var ResponseUnprocessableEntity = &helperresponse.Response{
	HttpStatusCode: http.StatusUnprocessableEntity,
	Status:         false,
	Error:          nil,
	Message: map[string]string{
		"en": "Unprocessable Entity",
		"id": "Terdapat Kesalahan Validasi",
	},
}

var ResponseUnprocessableEntityDB = &helperresponse.Response{
	HttpStatusCode: http.StatusUnprocessableEntity,
	Status:         false,
	Error:          nil,
	Message: map[string]string{
		"en": "Unprocessable Entity",
		"id": "Terdapat Kesalahan Validasi",
	},
}

var ResponseDailyQuotaExceeded = &helperresponse.Response{
	HttpStatusCode: http.StatusUnprocessableEntity,
	Status:         false,
	Error:          nil,
	Message: map[string]string{
		"en": "Your order exceeds the daily quota",
		"id": "Quota harian anda sudah habis",
	},
}

var ResponseOutOfQuota = &helperresponse.Response{
	HttpStatusCode: http.StatusUnprocessableEntity,
	Status:         false,
	Error:          nil,
	Message: map[string]string{
		"en": "Session quota sold out",
		"id": "Quota sesi sudah habis",
	},
}

var ResponseOverPaymentPeriod = &helperresponse.Response{
	HttpStatusCode: http.StatusUnprocessableEntity,
	Status:         false,
	Error:          nil,
	Message: map[string]string{
		"en": "Payment period has ended",
		"id": "Sudah melewati batas waktu pembayaran",
	},
}

var ResponseInvalidStatus = &helperresponse.Response{
	HttpStatusCode: http.StatusBadRequest,
	Status:         false,
	Error:          nil,
	Message: map[string]string{
		"en": "Invalid status",
		"id": "Status tidak sesuai",
	},
}

var ResponseExceedMaxOrderQuantity = &helperresponse.Response{
	HttpStatusCode: http.StatusBadRequest,
	Status:         false,
	Error:          nil,
	Message: map[string]string{
		"en": "Ticket orders exceed the maximum limit",
		"id": "Pesanan tiket melebihi batas maksimal",
	},
}

var ResponseExceedDailyHourLimit = &helperresponse.Response{
	HttpStatusCode: http.StatusBadRequest,
	Status:         false,
	Error:          nil,
	Message: map[string]string{
		"en": "Today's orders exceed closing hours",
		"id": "Pesanan hari ini melebihi jam tutup",
	},
}

var ResponseExceedAllowedOrderDays = &helperresponse.Response{
	HttpStatusCode: http.StatusBadRequest,
	Status:         false,
	Error:          nil,
	Message: map[string]string{
		"en": "Orders exceed the specified day limit",
		"id": "Pesanan melebihi batas hari yang ditentukan",
	},
}

var ResponseClosingDay = &helperresponse.Response{
	HttpStatusCode: http.StatusBadRequest,
	Status:         false,
	Error:          nil,
	Message: map[string]string{
		"en": "Cannot order on closing day",
		"id": "Tidak dapat memesan pada hari libur",
	},
}

var ResponseOrderBeforeToday = &helperresponse.Response{
	HttpStatusCode: http.StatusBadRequest,
	Status:         false,
	Error:          nil,
	Message: map[string]string{
		"en": "Unable to book on a date before today",
		"id": "Tidak dapat memesan pada tanggal sebelum hari ini",
	},
}

var ResponseOrderNotFound = &helperresponse.Response{
	HttpStatusCode: http.StatusNotFound,
	Status:         false,
	Error:          nil,
	Message: map[string]string{
		"en": "Order not found",
		"id": "Pesanan tidak ditemukan",
	},
}

var ResponseOrderTicketNotFound = &helperresponse.Response{
	HttpStatusCode: http.StatusNotFound,
	Status:         false,
	Error:          nil,
	Message: map[string]string{
		"en": "Order ticket not found",
		"id": "Tiket tidak ditemukan",
	},
}

var ResponseExternalServiceError = &helperresponse.Response{
	HttpStatusCode: http.StatusInternalServerError,
	Status:         false,
	Error:          nil,
	Message: map[string]string{
		"en": "The system is currently unavailable",
		"id": "Sistem sedang mengalami gangguan",
	},
}

var ResponseInvalidPaymentMethod = &helperresponse.Response{
	HttpStatusCode: http.StatusBadRequest,
	Status:         false,
	Error:          nil,
	Message: map[string]string{
		"en": "Invalid payment method",
		"id": "Metode pembayaran tidak valid",
	},
}

var ResponseNotImplemented = &helperresponse.Response{
	HttpStatusCode: http.StatusNotImplemented,
	Status:         false,
	Error:          nil,
	Message: map[string]string{
		"en": "The feature not implemented yet",
		"id": "Layanan masih dalam tahap pengembangan",
	},
}

var ResponseEmptyUserName = &helperresponse.Response{
	HttpStatusCode: http.StatusBadRequest,
	Status:         false,
	Error:          nil,
	Message: map[string]string{
		"en": "Please fill in your name in the profile",
		"id": "Mohon isi nama Anda di menu profil",
	},
}

var ResponseInvalidQRString = &helperresponse.Response{
	HttpStatusCode: http.StatusBadRequest,
	Status:         false,
	Error:          nil,
	Message: map[string]string{
		"en": "Unknown QR",
		"id": "QR tidak dikenal",
	},
}

var ResponseTicketSoldOut = &helperresponse.Response{
	HttpStatusCode: http.StatusBadRequest,
	Status:         false,
	Error:          nil,
	Message: map[string]string{
		"en": "Ticket sold out",
		"id": "Tiket habis",
	},
}
