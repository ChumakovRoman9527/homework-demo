package middleware

import (
	"5-order-api-auth/configs"
	"5-order-api-auth/pkg/jwt"
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type key string

const (
	ConstPhoneKey key = "ConstPhoneKey"
)

type AuthDeps struct {
	Config *configs.Config
}

func writeUnAuthed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func (authDeps AuthDeps) IsAuthed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			log.Println("Bearer токена нет !!!")
			writeUnAuthed(w)
			return
		}
		if !strings.HasPrefix(authorization, "Bearer ") {
			writeUnAuthed(w)
			return
		}
		token := strings.TrimPrefix(authorization, "Bearer ")
		fmt.Println(authDeps)
		isValid, data := jwt.NewJWT(authDeps.Config.Auth.Secret).Parse(token)
		if !isValid {
			writeUnAuthed(w)
			return
		}
		ctx := context.WithValue(r.Context(), ConstPhoneKey, data.Phone)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
