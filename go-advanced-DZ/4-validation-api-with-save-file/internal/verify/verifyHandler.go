package verify

import (
	"4-validation-api-with-save-file/configs"
	"4-validation-api-with-save-file/pkg/res"
	"strings"

	"fmt"
	"net/http"
)

type EmailHandler struct {
	*configs.EmailValidationConfig
}

type EmailResponse struct {
	statusCode int
	statusText string
}

func NewVerifyHandler(router *http.ServeMux, deps EmailHandler) {
	handler := &EmailHandler{
		EmailValidationConfig: deps.EmailValidationConfig,
	}
	router.HandleFunc("POST /send", handler.VerifyPost())
	router.HandleFunc("GET /verify/", handler.VerifyGet())
}

func (handler *EmailHandler) VerifyPost() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("r.Body = ", r.Body)
		emailVerifyResponse := SendVerifyEmail(*handler, r)
		// fmt.Println("Вот тут надо отправлять email !!!!")
		data := emailVerifyResponse.statusText
		status := emailVerifyResponse.statusCode
		res.Json(w, data, status)
	}
}

func (handler *EmailHandler) VerifyGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		VerifyURI, err := r.URL.Parse(r.RequestURI)
		if err != nil {
			fmt.Println("!!Ошибка обработки URI!!")
		}

		sHash := strings.Split(VerifyURI.String(), "/")[2]

		HashValid := Hash_Check(*handler, sHash)
		data := HashValid.statusText
		status := HashValid.statusCode

		res.Json(w, data, status)

	}
}
