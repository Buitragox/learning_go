package main

import "fmt"

func addValue(numbers []int) {
	numbers = append(numbers, 6)
	fmt.Println(numbers)
}

func main() {
	numbers := make([]int, 0, 10)
	for i := 1; i <= 5; i++ {
		numbers = append(numbers, i)
	}

	addValue(numbers)
	fmt.Println(numbers)
}
