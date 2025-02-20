package main

import "fmt"

type person struct {
	age int
}

func modifyValue(p person, n int) {
	n *= 2
	p.age = n
	fmt.Println(p, n)
}

func main() {
	p := person{25}
	n := 40
	modifyValue(p, n) // modify `n` and `p`
	fmt.Println(p, n) // still prints 25, 40
}
