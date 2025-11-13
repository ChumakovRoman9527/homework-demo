package main

import (
	// "6-Architecture/configs"
	"10-Middleware-order-api/configs"
	"10-Middleware-order-api/internal/auth"
	"10-Middleware-order-api/internal/product"
	"10-Middleware-order-api/pkg/db"
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
