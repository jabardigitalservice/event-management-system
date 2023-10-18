package response

import (
	"net/http"

	helperresponse "github.com/fazpass/goliath/v3/helper/http/response"
	"github.com/go-chi/render"
)

type Response struct {
	HttpStatusCode int
	Status         bool
	Error          error
	Message        map[string]string
	Code           string
}

func Render(
	w http.ResponseWriter,
	r *http.Request,
	responses map[string]*Response,
	code string,
	data interface{},
	err interface{},
) {
	var (
		response = responses[code]
		lang     = r.Header.Get("X-localization")
	)

	if lang == "" {
		lang = "en"
	}

	render.Status(r, response.HttpStatusCode)
	_ = render.Render(w, r, helperresponse.Build(&helperresponse.Response{
		Status:  response.Status,
		Errors:  err,
		Message: GetMessage(response.Message, lang),
		Code:    code,
		Data:    data,
	}))
}

func GetMessage(messages map[string]string, lang string) string {
	var value string
	var ok bool

	value, ok = messages[lang]
	if ok {
		return value
	}

	return messages["en"]
}
