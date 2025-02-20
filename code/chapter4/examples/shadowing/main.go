package main

import "fmt"

var y int = 100

func main() {
	fmt.Println(100)
	y := 50 // Shadows the package variable y
	fmt.Println(y)

	x := 99
	if x > 10 {
		fmt.Println(x)
		x := 5 // shadows the local variable x
		fmt.Println(x)
	}
	fmt.Println(x)

	// You can shadow universe block identifiers
	true := 1
	fmt.Println(true)

	// You can shadow packages
	// fmt := 5
	// fmt.Printf("fmt: %v\n", fmt)
}
