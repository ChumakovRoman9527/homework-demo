package main

import (
	"fmt"
	"math/rand/v2"
)

func generateNum(max int, min int) int {

	res := min + rand.IntN(max-min)

	return res
}

func generateSlice(size int, max int, min int, numChan chan int) {
	defer close(numChan)
	res := make([]int, size)

	for i := 0; i < size; i++ {
		res[i] = generateNum(max, min)
	}

	for _, val := range res {

		numChan <- val
	}

}

func consoleSquare(numChan chan int, squareChan chan int) {

	for iter := range numChan {

		squareChan <- squareNum(iter)

	}
	close(squareChan)

}

func squareNum(num int) int {
	return num * num
}

func main() {
	//Определяем размер каналов, масимум и минимум для случайных чисел
	chanBuff := 10
	max := 100
	min := 0
	// Создаем каналы
	numChan := make(chan int, chanBuff)
	squareChan := make(chan int, chanBuff)
	//Запускаем горутину генерации случайных чисел, создания слайса
	go generateSlice(chanBuff, max, min, numChan)
	//Запускаем горутину возведения в квадрат случайных чисел ищ горутины 1
	go consoleSquare(numChan, squareChan)
	//ожидаем когда канал квадратных чисел закроется, а пока ждем последовательно выводим в консоль
	//важно ! не понял почему если сделать слепой range а в теле цикла выводить извлекая из канала значение, канал будет перебран неполностью ! вероятно из за рассинхронизации
	for val := range squareChan {

		fmt.Printf("%d ", val)
	}

}
