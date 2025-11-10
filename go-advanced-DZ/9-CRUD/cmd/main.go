package main

import (
	// "6-Architecture/configs"
	"9-CRUD_ORDER_API/configs"
	"9-CRUD_ORDER_API/internal/auth"
	"9-CRUD_ORDER_API/internal/product"
	"9-CRUD_ORDER_API/pkg/db"
	"fmt"
	"log"
	"net/http"
)

func main() {

	conf := configs.LoadConfig()

	// router := http.NewServeMux()
	// hello.NewHelloHandler(router)
	db := db.NewDb(conf)
	router := http.NewServeMux()

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	productRepository := product.NewProductRepository(db)

	product.ProductsHandler(router, product.ProductHandlerDeps{ProductRepository: productRepository})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	//http.ListenAndServe(":8081", nil)
	err := server.ListenAndServe()
	if err != nil {
		log.Panic("ошибка запуска сервера")
	}

}
