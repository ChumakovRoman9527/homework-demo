package main

import (
	"fmt"
	"time"
)

func main() {
	go printHi()
	go fmt.Println("привет из main")
	go fmt.Println("привет из main2")
	time.Sleep(time.Second)

}

func printHi() {
	fmt.Println("привет из gr")
}
