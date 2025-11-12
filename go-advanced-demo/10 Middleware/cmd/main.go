package main

import (
	// "6-Architecture/configs"
	"10-Middleware/configs"
	"10-Middleware/internal/auth"
	"10-Middleware/internal/link"
	"10-Middleware/pkg/db"
	"10-Middleware/pkg/middleware"
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
	//Repositories
	linkRepository := link.NewLinkRepository(db)

	//Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})
	// /auth/login
	// /auth/register

	server := http.Server{
		Addr:    ":8081",
		Handler: middleware.Logging(router),
	}

	fmt.Println("Server is listening on port 8081")
	//http.ListenAndServe(":8081", nil)
	err := server.ListenAndServe()
	if err != nil {
		log.Panic("ошибка запуска сервера")
	}

}
