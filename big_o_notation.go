package main

import "fmt"

// O(1)
func oOne(numbers []int) {
	fmt.Println(numbers[0])
}

// O(log(n))
func oLogN(n float64) {
	if n <= 1 {
		return
	}

	fmt.Println(n)
	oLogN(n / 2)
}

// O(n)
func oN(numbers []int) {
	for _, i := range numbers {
		fmt.Println(i)
	}
}

// O(n * log(n))
func oNLongN(n int) {
	for i := 0; i < n; i++ {
		fmt.Print(i)
	}

	fmt.Println()
	if n <= 1 {
		return
	}

	oNLongN(n / 2)
}

// O(n**2)
func oN2(numbers []int) {
	for _, i := range numbers {
		for _, j := range numbers {
			fmt.Println(i, j)
		}
		fmt.Println()
	}
}
