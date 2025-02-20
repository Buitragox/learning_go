package main

import "fmt"

func main() {
	n := 10
	pointerToN := &n
	fmt.Println(n, *pointerToN, pointerToN)

	// illegal := &"hello"
	// illegal := &1260

	x := new(int)
	fmt.Println(x, *x)

}
