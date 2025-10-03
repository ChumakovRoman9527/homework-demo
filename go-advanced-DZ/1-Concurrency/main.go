package main

import (
	"fmt"
	"math/rand/v2"
)

func generateNum(numChan chan int, max int, min int) {
	res := -1

	res = min + rand.IntN(max-min)

	numChan <- res
}

func squareNum(num int) int {
	return num * num
}

func main() {
	chanBuff := 10
	numChan := make(chan int, chanBuff)

	for i := 0; i < chanBuff; i++ {
		go generateNum(numChan, 10, 0)
	}

	for range chanBuff {
		fmt.Printf("%d ", squareNum(<-numChan))
	}

}
