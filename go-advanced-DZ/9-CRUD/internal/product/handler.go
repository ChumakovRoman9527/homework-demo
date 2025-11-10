package product

import (
	"9-CRUD_ORDER_API/pkg/req"
	"9-CRUD_ORDER_API/pkg/res"
	"net/http"
)

type productsHandler struct {
	ProductRepository *ProductRepository
}

type ProductHandlerDeps struct {
	ProductRepository *ProductRepository
}

func ProductsHandler(router *http.ServeMux, deps ProductHandlerDeps) {
	handler := &productsHandler{ProductRepository: deps.ProductRepository}
	router.HandleFunc("POST /product", handler.Create())
	// router.HandleFunc("PATCH /product/{id}", handler.Update())
	// router.HandleFunc("DELETE /product/{id}", handler.Delete())
	// router.HandleFunc("GET /{hash}", handler.GoTo())
}

func (handler *productsHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[ProductCreateRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		product := newProduct(*body)

		createdProduct, err := handler.ProductRepository.Create(product)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, createdProduct, http.StatusCreated)

	}
}
