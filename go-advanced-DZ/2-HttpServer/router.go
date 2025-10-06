package main

import (
	"2-HttpServerDZ/internal"
	"fmt"
	"net/http"
)

type HomeWorkHandler struct{}

func NewHelloHandler(router *http.ServeMux) {
	handler := &HomeWorkHandler{}
	router.HandleFunc("/getRND", handler.getRND())
}

func (handler *HomeWorkHandler) getRND() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(internal.GetRND())
	}
}
