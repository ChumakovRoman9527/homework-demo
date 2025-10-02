package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func ping(url string, RespChan chan int, ErrChan chan error) {
	fmt.Println("Начали ", url)
	resp, err := http.Get(url)
	if err != nil {
		ErrChan <- err
		return
	}
	RespChan <- resp.StatusCode
	// fmt.Printf("Закончили %s\n", url)
}

func main() {
	path := flag.String("file", "url.txt", "path to urls file")
	flag.Parse()
	file, err := os.ReadFile(*path)
	if err != nil {
		panic(err.Error())
	}
	urlSlice := strings.Split(string(file), "\n")

	respChan := make(chan int)
	errChan := make(chan error)

	for _, url_ := range urlSlice {
		go ping(url_, respChan, errChan)
	}

	for i := 0; i < len(urlSlice); i++ {
		res := <-respChan
		errRes := <-errChan
		fmt.Println("ОК  \n", res)
		fmt.Println("Ошибка !!! \n", errRes)
	}

}
