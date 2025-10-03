package main

import (
	"fmt"
)

func arrGen(arr []int, numberArrElements int) []int {
	for i := 1; i <= numberArrElements; i++ {
		arr = append(arr, i)
	}

	return arr
}

func main() {
	var arr []int
	numberArrElements := 12
	arr = arrGen(arr, numberArrElements)

	numberGoroutines := 7
	sumChan := make(chan int, numberGoroutines)
	var wholeSum int

	var partlen int
	partlen = len(arr) / numberGoroutines
	if len(arr)%numberGoroutines > 0 {
		partlen = partlen + 1
	}

	for i := 0; i < numberGoroutines; i++ {
		index1 := i * partlen

		if index1 > len(arr) {
			continue
		}
		index2 := index1 + partlen
		if index2 > len(arr) {
			index2 = len(arr)
		}
		fmt.Println(index1, index2)

		go partSum(sumChan, arr[index1:index2])

	}
	// go func() {
	// 	wg.Wait()
	// 	close(sumChan)
	// }()
	// for ite := range sumChan {
	// 	wholeSum = wholeSum + ite
	// }
	for i := 0; i < numberGoroutines; i++ {
		wholeSum += <-sumChan
	}
	fmt.Printf("Сумма = %d", wholeSum)

}

func partSum(sumChan chan int, partarr []int) {
	fmt.Println(partarr)
	partSum := 0
	for _, arr_var := range partarr {
		partSum = partSum + arr_var
	}

	sumChan <- partSum
}
