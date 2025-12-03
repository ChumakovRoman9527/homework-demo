package auth

import (
	"6-order-api-cart/configs"
	"6-order-api-cart/pkg/jwt"
	"6-order-api-cart/pkg/req"
	"6-order-api-cart/pkg/res"
	"fmt"
	"net/http"
)

type authHandler struct {
	*configs.Config
	*AuthService
}

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &authHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/register", handler.Register())
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/loginphone", handler.LoginPhone())
	router.HandleFunc("POST /auth/loginSMS", handler.LoginSMS())
}

func (handler *authHandler) LoginPhone() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := req.HandleBody[LoginPhoneRequest](&w, r)
		if err != nil {
			return
		}

		sessionId, err := handler.AuthService.SessionGenerate(body.Phone)
		if err != nil {
			res.Json(w, err.Error(), http.StatusUnauthorized)
			return
		}

		data := LoginPhoneResponse{
			SessionID: sessionId,
		}
		res.Json(w, data, http.StatusOK)
	}
}

func (handler *authHandler) LoginSMS() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := req.HandleBody[LoginSMSRequest](&w, r)
		if err != nil {
			return
		}
		phone, err := handler.AuthService.CodeCheck(body.SessionID, body.Code)
		if err != nil {
			res.Json(w, err.Error(), http.StatusUnauthorized)
			return
		}

		j := jwt.NewJWT(
			handler.Auth.Secret,
		)

		token, err := j.Create(phone)
		if err != nil {
			res.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = handler.AuthService.ClearOldSession(body.SessionID)
		if err != nil {
			res.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := LoginSMSResponse{
			TOKEN: token,
		}

		res.Json(w, data, http.StatusOK)

	}
}

func (handler *authHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		fmt.Println(body)

		data := LoginResponse{
			TOKEN: "9999", //handler.Config.Auth.Secret,
		}
		res.Json(w, data, http.StatusOK)
	}
}

func (handler *authHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := req.HandleBody[RegisterRequest](&w, r)
		fmt.Println(body)

		data := RegisterResponse{
			TOKEN: "9999", //handler.Config.Auth.Secret,
		}
		res.Json(w, data, http.StatusOK)
	}
}
