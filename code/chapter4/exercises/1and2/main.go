package main

import (
	"fmt"
	"math/rand"
)

func main() {
	r := rand.New(rand.NewSource(42))

	numbers := make([]int, 0, 100)
	for i := 0; i < 10; i++ {
		numbers = append(numbers, r.Intn(101))
	}
	fmt.Println(len(numbers), numbers)

	for _, n := range numbers {
		switch {
		//case n%2 == 0 && n%3 == 0:
		case n%6 == 0:
			fmt.Println("Six!")
		case n%2 == 0:
			fmt.Println("Two!")
		case n%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}
}
