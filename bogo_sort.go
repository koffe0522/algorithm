package main

import (
	"fmt"
)

func isOrder(numbers []int) bool {
	for i := range numbers[:len(numbers)-1] {
		if numbers[i] > numbers[i+1] {
			return false
		}
	}

	return true
}

func bogoSort(numbers []int) []int {
	for {
		if isOrder(numbers) {
			break
		}
		shuffle(numbers)
	}

	fmt.Println()

	return numbers
}
