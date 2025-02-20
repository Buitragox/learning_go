package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 2, 4)
	num := copy(y, x)
	fmt.Printf("y: %v, num: %d\n", y, num)

	a := []string{"a", "b", "c", "d", "f"}
	b := make([]string, 10)
	num2 := copy(b, a[2:])
	fmt.Printf("\nb: %v, num: %d\n", b, num2)
}
