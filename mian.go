package main

import (
	"fmt"
)

func main() {
	numbers := []int{20, 64, 18, 4, 98}
	shuffle(numbers)
	bogoSort(numbers)
	fmt.Println(numbers)
}
