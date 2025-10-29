package main

import (
	// "6-Architecture/configs"
	"8-DockerPGSQL/configs"
	"8-DockerPGSQL/internal/auth"
	"8-DockerPGSQL/pkg/db"
	"fmt"
	"log"
	"net/http"
)

func main() {

	conf := configs.LoadConfig()

	// router := http.NewServeMux()
	// hello.NewHelloHandler(router)
	_ = db.NewDb(conf)
	authrouter := http.NewServeMux()
	auth.NewAuthHandler(authrouter, auth.AuthHandlerDeps{
		Config: conf,
	})

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
