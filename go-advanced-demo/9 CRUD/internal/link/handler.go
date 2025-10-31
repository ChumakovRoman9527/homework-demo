package link

import (
	"net/http"
)

type linkHandler struct {
}

type LinkHandlerDeps struct {
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &linkHandler{}
	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /linke/{id}", handler.Delete())
	router.HandleFunc("GET /{alias}", handler.GoTo())

}

func (handler *linkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *linkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *linkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *linkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
