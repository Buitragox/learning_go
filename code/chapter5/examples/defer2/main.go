package main

import "fmt"

func main() {

	a := 10
	defer func(n int) {
		fmt.Println("first:", n)
	}(a)

	a = 20
	defer func(n int) {
		fmt.Println("second:", n)
	}(a)

	a = 30
	fmt.Printf("third: %v\n", a)
}
