package main

import (
	"fmt"
	"net/http"
)

func main() {
	code := make(chan int)
	for i := 0; i <= 10; i++ {
		go getHttpCode(code)
	}
	<-code
}

func getHttpCode(codeCh chan int) {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Ошибка !!!!")
	}
	fmt.Println("Код ответа:", resp.StatusCode)
	codeCh <- resp.StatusCode
}
