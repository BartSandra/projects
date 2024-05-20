package main

import (
	"fmt"
	"time"
)

func sleepSort(numbers []int) <-chan int {
	result := make(chan int)

	for _, number := range numbers {
		go func(n int) {
			time.Sleep(time.Duration(n) * time.Second)
			result <- n
		}(number)
	}

	go func() {
		time.Sleep(time.Duration(maxValue(numbers)+1) * time.Second)
		close(result)
	}()

	return result
}

func maxValue(numbers []int) int {
	maxNumber := numbers[0]
	for _, n := range numbers[1:] {
		if n > maxNumber {
			maxNumber = n
		}
	}
	return maxNumber
}

func main() {
	numbers := []int{1, 8, 2, 3, 9, 7, 4, 6}
	sorted := sleepSort(numbers)
	for n := range sorted {
		fmt.Println(n)
	}
}
