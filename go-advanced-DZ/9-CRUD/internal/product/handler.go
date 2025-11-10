package product

import (
	"9-CRUD_ORDER_API/pkg/req"
	"9-CRUD_ORDER_API/pkg/res"
	"net/http"
	"strconv"
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
	router.HandleFunc("PATCH /product/{id}", handler.Update())
	router.HandleFunc("DELETE /product/{id}", handler.Delete())
	router.HandleFunc("GET /product/{id}", handler.Get())
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

func (handler *productsHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[ProductUpdateRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		newoldProduct, err := handler.ProductRepository.GetByID(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		changed := false

		if newoldProduct.Name != body.Name {
			newoldProduct.Name = body.Name
			changed = true
		}

		if newoldProduct.Description != body.Description {
			newoldProduct.Description = body.Description
			changed = true
		}

		newoldProduct.Images = body.Images

		if !changed {
			http.Error(w, "нет изменений !", http.StatusBadRequest)
			return
		}

		updatedProduct, err := handler.ProductRepository.Update(newoldProduct)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, updatedProduct, http.StatusCreated)

	}
}

func (handler *productsHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = handler.ProductRepository.GetByID(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = handler.ProductRepository.Delete(uint(id))

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, "Продукт удален", http.StatusCreated)

	}
}

func (handler *productsHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		Product, err := handler.ProductRepository.GetByID(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, Product, http.StatusCreated)

	}
}
