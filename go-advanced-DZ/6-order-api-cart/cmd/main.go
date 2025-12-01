package main

import (
	// "6-Architecture/configs"
	"6-order-api-cart/configs"
	"6-order-api-cart/internal/auth"
	"6-order-api-cart/internal/orders"
	"6-order-api-cart/internal/product"
	"6-order-api-cart/pkg/db"
	"6-order-api-cart/pkg/middleware"
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

	productRepository := product.NewProductRepository(db)
	phoneAuthRepository := auth.NewPhoneAuthRepository(db)
	orderRepository := orders.NewOrderRepository(db)
	//Services
	AuthService := auth.NewAuthService(phoneAuthRepository)

	product.ProductsHandler(router, product.ProductHandlerDeps{
		ProductRepository: productRepository,
		Config:            conf,
	})

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: AuthService,
	})

	orders.OrderHandler(router, orders.OrderHandlerDeps{
		OrderRepository: orderRepository,
		Config:          conf,
	})

	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	fmt.Println("Server is listening on port 8081")
	//http.ListenAndServe(":8081", nil)
	err := server.ListenAndServe()
	if err != nil {
		log.Panic("ошибка запуска сервера")
	}

}
