package main

func cocktailSort(numbers []int) []int {
	start := 0
	end := len(numbers) - 1
	swapped := true

	for {
		swapped = false

		// ->
		for i := start; i < end; i++ {
			if numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
				swapped = true
			}
		}

		/* Range Sintax
		for i := range numbers[start:end] {
			i = start + i
			if numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
				swapped = true
			}
		}
		*/

		if !swapped {
			break
		}

		end--

		// reverse loop
		// <-
		for i := end; i > start-1; i-- {
			if numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
				swapped = true
			}
		}

		/* Range Sintax
		for i := range numbers[start:end] {
			j := end - i - 1
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
				swapped = true
			}
		}
		*/

		start++
	}

	return numbers
}
