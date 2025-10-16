package auth

import (
	"7-RequestAndValidation/configs"
	"7-RequestAndValidation/pkg/res"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
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
		var payload LoginRequest
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			res.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		if payload.Email == "" || payload.Password == "" {
			if payload.Email == "" && payload.Password != "" {
				res.Json(w, "e-mail required", http.StatusBadRequest)
			}
			if payload.Email != "" && payload.Password == "" {
				res.Json(w, "password required", http.StatusBadRequest)
			}
			if payload.Email == "" && payload.Password == "" {
				res.Json(w, "password and email required", http.StatusBadRequest)
			}
			return
		}
		reg, _ := regexp.Compile(`[A-Za-z0-9\._%+\-]+@[A-Za-z0-9\._%+\-]+\.[A-Za-z]{2,}`)
		if !reg.MatchString(payload.Email) {
			res.Json(w, "e-mail is wrong !!!", http.StatusBadRequest)
			return
		}
		fmt.Println("r.Body:", payload)

		data := LoginResponse{
			TOKEN: "9999", //handler.Config.Auth.Secret,
		}
		res.Json(w, data, http.StatusOK)
	}
}

func (handler *authHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(handler.Config.Auth.Secret)
		fmt.Println("register !!!!")
	}
}
