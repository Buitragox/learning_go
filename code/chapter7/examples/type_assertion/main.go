package main

import "fmt"

type MyInt int

func main() {
	var a any
	var mi MyInt = 100

	a = mi

	b := a.(MyInt)
	fmt.Println(b + 20)

	// Code panics
	// c := a.(string)
	// fmt.Print(c)

	// Code panics
	// d := a.(int)
	// fmt.Println(d)

	e, ok := a.(bool)
	if !ok {
		fmt.Println("The underlying type of `a` is not a bool")
	}
	fmt.Println("Value of e:", e)
}
