package auth

import (
	"9-CRUD_ORDER_API/configs"
	"9-CRUD_ORDER_API/pkg/req"
	"9-CRUD_ORDER_API/pkg/res"
	"fmt"
	"net/http"
)

type authHandler struct {
	*configs.Config
}

type AuthHandlerDeps struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &authHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/register", handler.Register())
	router.HandleFunc("POST /auth/login", handler.Login())
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
