package main

import (
	// "6-Architecture/configs"
	"6-Architecture/internal/auth"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// _ := configs.LoadConfig()

	// router := http.NewServeMux()
	// hello.NewHelloHandler(router)

	authrouter := http.NewServeMux()
	auth.NewAuthHandler(authrouter)

	// /auth/login
	// /auth/register

	server := http.Server{
		Addr:    ":8081",
		Handler: authrouter,
	}

	fmt.Println("Server is listening on port 8081")
	//http.ListenAndServe(":8081", nil)
	err := server.ListenAndServe()
	if err != nil {
		log.Panic("ошибка запуска сервера")
	}

}
