package main

func cocktailSort(numbers []int) []int {
	start := 0
	end := len(numbers) - 1
	swapped := true

	for {
		swapped = false
		// ->
		for i := range numbers[start:end] {
			if numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
				swapped = true
			}
		}

		if !swapped {
			break
		}

		end--

		// reverse loop
		// ->
		for i := range numbers[start:end] {
			i = end - i
			if numbers[i-1] > numbers[i] {
				numbers[i], numbers[i-1] = numbers[i-1], numbers[i]
				swapped = true
			}
		}

		if !swapped {
			break
		}

		start++
	}

	return numbers
}
