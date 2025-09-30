package main

import (
	"fmt"
	"net/http"
)

func main() {
	code := make(chan int)
	for i := 0; i < 10; i++ {
		go getHttpCode(code)
	}
	for res := range code {
		fmt.Println("Код ответа:", res)
	}
}

func getHttpCode(codeCh chan int) {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Ошибка !!!!")
	}

	codeCh <- resp.StatusCode
}
