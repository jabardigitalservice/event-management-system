package http

import helperresponse "github.com/jabardigitalservice/super-app-services/event/src/helper/http/response"

const (
	Success                            = "2000800"
	SuccessEmptyActiveOrders           = "2000801"
	Created                            = "2010800"
	BadRequest                         = "4000800"
	BadRequestBody                     = "4000801"
	BadRequestInvalidStatus            = "4000802"
	ExceedMaxOrderQuantity             = "4000803"
	ExceedDailyHourLimit               = "4000804"
	ExceedAllowedOrderDays             = "4000805"
	ClosingDay                         = "4000806"
	OrderBeforeToday                   = "4000807"
	InvalidPaymentMethod               = "4000808"
	EmptyUserName                      = "4000809"
	InvalidQRString                    = "4000810"
	TicketSoldOut                      = "4000811"
	Unauthorized                       = "4010800"
	TokenExpired                       = "4010801"
	NotFound                           = "4040800"
	OrderNotFound                      = "4040801"
	OrderTicketNotFound                = "4040802"
	Conflict                           = "4090800"
	UnprocessableEntity                = "4220800"
	UnprocessableEntityDB              = "4220801"
	UnprocessableEntityExternalService = "4220802"
	DailyQuotaExceeded                 = "4220803"
	OutOfQuota                         = "4220804"
	OverPaymentPeriod                  = "4220805"
	InternalServerError                = "5000800"
	InternalServerErrorDB              = "5000801"
	ExternalServiceError               = "5000802"
	InvalidFCMToken                    = "5000803"
	FailedCreateMessaging              = "5000804"
	FailedSendNotification             = "5000805"
	NotImplemented                     = "5010800"
)

var responses = map[string]*helperresponse.Response{
	"2000800": ResponseSuccess,
	"2000801": ResponseSuccessEmptyActiveOrders,
	"2010800": ResponseCreated,
	"4000800": ResponseBadRequest,
	"4000801": ResponseBadRequestBody,
	"4000802": ResponseInvalidStatus,
	"4000803": ResponseExceedMaxOrderQuantity,
	"4000804": ResponseExceedDailyHourLimit,
	"4000805": ResponseExceedAllowedOrderDays,
	"4000806": ResponseClosingDay,
	"4000807": ResponseOrderBeforeToday,
	"4000808": ResponseInvalidPaymentMethod,
	"4000809": ResponseEmptyUserName,
	"4000810": ResponseInvalidQRString,
	"4000811": ResponseTicketSoldOut,
	"4010800": ResponseUnauthorized,
	"4010801": ResponseTokenExpired,
	"4040800": ResponseNotFound,
	"4040801": ResponseOrderNotFound,
	"4040802": ResponseOrderTicketNotFound,
	"4090800": ResponseConflict,
	"4220800": ResponseUnprocessableEntity,
	"4220801": ResponseUnprocessableEntityDB,
	"4220803": ResponseDailyQuotaExceeded,
	"4220804": ResponseOutOfQuota,
	"4220805": ResponseOverPaymentPeriod,
	"5000800": ResponseInternalServerError,
	"5000801": ResponseInternalServerErrorDB,
	"5000802": ResponseExternalServiceError,
	"5000803": ResponseInvalidFCMToken,
	"5000804": ResponseFailedCreateMessaging,
	"5000805": ResponseFailedSendNotification,
	"5010800": ResponseNotImplemented,
}

func GetResponses() map[string]*helperresponse.Response {
	return responses
}
