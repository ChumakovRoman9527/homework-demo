package main

import "fmt"

func double(x int) int {
	return x * 2
}

func main() {
	x := 10
	y := double(x)
	fmt.Println(y)

}
