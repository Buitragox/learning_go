package main

import "fmt"

type MyParams struct {
	firstName string
	lastName  string
	age       int
}

func myFunc(opts MyParams) {
	fmt.Println(opts)
}

func main() {
	myFunc(MyParams{
		firstName: "Gob",
		age:       42,
	})
}
