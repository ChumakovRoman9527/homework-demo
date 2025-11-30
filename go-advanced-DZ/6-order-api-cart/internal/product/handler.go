package product

import (
	"6-order-api-cart/configs"
	"6-order-api-cart/pkg/middleware"
	"6-order-api-cart/pkg/req"
	"6-order-api-cart/pkg/res"
	"fmt"
	"net/http"
	"strconv"
)

type productsHandler struct {
	ProductRepository *ProductRepository
	Config            *configs.Config
}

type ProductHandlerDeps struct {
	ProductRepository *ProductRepository
	Config            *configs.Config
}

func ProductsHandler(router *http.ServeMux, deps ProductHandlerDeps) {

	handler := &productsHandler{
		ProductRepository: deps.ProductRepository,
		Config:            deps.Config,
	}
	authDeps := middleware.AuthDeps{Config: deps.Config}
	fmt.Println("ProductsHandler", deps.Config)
	stack := middleware.Chain(
		authDeps.IsAuthed,
	)
	router.Handle("POST /product", stack(handler.Create()))
	router.Handle("PATCH /product/{id}", stack(handler.Update()))
	router.Handle("DELETE /product/{id}", stack(handler.Delete()))
	router.Handle("GET /product/{id}", stack(handler.Get()))
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
