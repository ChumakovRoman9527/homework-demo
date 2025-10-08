package auth

import (
	"6-Architecture/configs"
	"6-Architecture/pkg/res"
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
		fmt.Println("login !!!!")
		data := LoginResponse{
			TOKEN: "9999", //handler.Config.Auth.Secret,
		}
		res.Json(w, data, http.StatusAccepted)
	}
}

func (handler *authHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(handler.Config.Auth.Secret)
		fmt.Println("register !!!!")
	}
}
