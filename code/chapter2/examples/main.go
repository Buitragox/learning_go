package main

import "fmt"

func main() {
	var x int = 10
	var y float64 = 30.8

	// Explicit type conversion is required
	var sumFloat float64 = float64(x) + y
	var sumInt int = x + int(y)

	fmt.Println(sumFloat, sumInt)

	var (
		a    int
		b    = 20
		z    int
		d, e = 30, "hey"
		f, g string
	)

	fmt.Println(a, b, z, d, e, f, g)

	k := 2
	fmt.Println(k)
}
