package main

import (
	"fmt"
	"net/http"

	"5.2-SimpleHTTPServer/configs"
	"5.2-SimpleHTTPServer/internal/hello"
)

func main() {

	conf := configs.LoadConfig()

	router := http.NewServeMux()
	hello.NewHelloHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	//http.ListenAndServe(":8081", nil)
	server.ListenAndServe()

}
