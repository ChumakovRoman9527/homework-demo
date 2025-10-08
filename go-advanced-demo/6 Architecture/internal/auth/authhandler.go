package auth

import (
	"fmt"
	"net/http"
)

type authHandler struct{}

func NewAuthHandler(router *http.ServeMux) {
	handler := &authHandler{}
	router.HandleFunc("/auth/register", handler.Register())
	router.HandleFunc("/auth/login", handler.Login())
}

func (handler *authHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		///Не нравится мне что в каждом хендлере надо писать проверку на метод
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", "POST")
			http.Error(w, "Метод запрещен", http.StatusMethodNotAllowed)
			return
		}
		fmt.Println("register !!!!")
		w.Write([]byte("register !!!"))
	}
}

func (handler *authHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		///Не нравится мне что в каждом хендлере надо писать проверку на метод
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", "POST")
			http.Error(w, "Метод запрещен", http.StatusMethodNotAllowed)
			return
		}
		fmt.Println("login !!!!")
		w.Write([]byte("login !!!"))
	}

}
