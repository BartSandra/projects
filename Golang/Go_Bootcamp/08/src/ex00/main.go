package main

import (
	"errors"
	"fmt"
)

func getElement(arr []int, idx int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("slice is empty")
	}
	if idx < 0 {
		return 0, errors.New("index is negative")
	}
	if idx >= len(arr) {
		return 0, errors.New("index is out of bounds")
	}

	for i := 0; i < idx; i++ {
		arr = arr[1:]
	}

	return arr[0], nil
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	idx := 4

	element, err := getElement(arr, idx)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Element at index %d is %d\n", idx, element)
	}
}
