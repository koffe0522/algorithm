package main

func bubbleSort(numbers []int) []int {
	lenNumbers := len(numbers)
	for i := range numbers {
		for j := range numbers[:lenNumbers-1-i] {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	return numbers
}
