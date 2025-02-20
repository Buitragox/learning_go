package main

import "fmt"

func main() {
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	// Use a full slice expression to limit the capacity of the subslices
	y := x[:2:2]
	z := x[2:4:4]

	fmt.Println("Before append")
	fmt.Printf("\tx: %v, cap: %d\n", x, cap(x))
	fmt.Printf("\ty: %v, cap: %d\n", y, cap(y))
	fmt.Printf("\tz: %v, cap: %d\n", z, cap(z))

	x = append(x, "x")

	// Since len == cap, a new memory allocation will take place, so they don't share memory anymore.
	y = append(y, "i", "j")
	z = append(z, "z")

	fmt.Println("\nAfter append")
	fmt.Printf("\tx: %v, cap: %d\n", x, cap(x))
	fmt.Printf("\ty: %v, cap: %d\n", y, cap(y))
	fmt.Printf("\tz: %v, cap: %d\n", z, cap(z))

}
