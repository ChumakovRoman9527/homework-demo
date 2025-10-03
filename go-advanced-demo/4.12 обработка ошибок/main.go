package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func ping(url string, RespChan chan string, ErrChan chan error) {
	// fmt.Println("Начали ", url)
	resp, err := http.Get(url)
	if err != nil {
		ErrChan <- err
		return
	}
	RespChan <- fmt.Sprintf("%s -> %d", url, resp.StatusCode)
	// fmt.Printf("Закончили %s\n", url)
}

func main() {
	path := flag.String("file", "url.txt", "path to urls file")
	flag.Parse()
	file, err := os.ReadFile(*path)
	if err != nil {
		panic(err.Error())
	}
	urlSlice := strings.Split(string(file), "\r\n")

	respChan := make(chan string)
	errChan := make(chan error)

	for _, url_ := range urlSlice {
		go ping(url_, respChan, errChan)
	}

	for range urlSlice {
		select {
		case errRes := <-errChan:
			fmt.Println(" Ошибка !!! ", errRes)
		case res := <-respChan:
			fmt.Println(" ОК ", res)
		}

	}

}
