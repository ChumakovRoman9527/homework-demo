package auth

import (
	"fmt"
	"net/http"
)

type authHandler struct{}

func NewAuthHandler(router *http.ServeMux) {
	handler := &authHandler{}
	router.HandleFunc("POST /auth/register", handler.Register())
	router.HandleFunc("POST /auth/login", handler.Login())
}

func (handler *authHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("register !!!!")
		w.Write([]byte("register !!!"))
	}
}

func (handler *authHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("login !!!!")
		w.Write([]byte("login !!!"))
	}

}
