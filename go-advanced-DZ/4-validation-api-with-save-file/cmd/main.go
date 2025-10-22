package main

import (
	// "6-Architecture/configs"

	"4-validation-api-with-save-file/configs"
	"4-validation-api-with-save-file/internal/verify"
	"fmt"
	"log"
	"net/http"
)

/*
TODO
1. Создать структуру payload входящего email и валидировать ее соответственно
2. Создать структуру json файла для хранения ожидающих валидации email
3. Создать структуру json файлов для хранения валидированных email
4. Создать функции создания, чтения файлов
*/
func main() {

	conf := configs.LoadConfig()

	// router := http.NewServeMux()
	// hello.NewHelloHandler(router)

	authrouter := http.NewServeMux()
	verify.NewVerifyHandler(authrouter, verify.EmailHandler{
		EmailValidationConfig: conf,
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
