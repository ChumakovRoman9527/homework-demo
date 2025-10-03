package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

//10 запросов GET по адресу google.com
// StatusCode вывести в консоле

func main() {
	t := time.Now()
	var wg sync.WaitGroup
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func() {
			getHttpCode(i)
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println(time.Since(t))
}

func getHttpCode(i int) int64 {

	t := time.Now()
	var curr_t int64
	fmt.Println("Начало запроса - ", i)
	rasp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Ошибка запроса!!! ", i)
		return 0
	}
	fmt.Println("Результат запроса - ", i, rasp.StatusCode)
	fmt.Println(time.Since(t))

	curr_t = time.Since(t).Milliseconds()
	return curr_t
}
