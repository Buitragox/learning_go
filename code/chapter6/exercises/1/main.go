package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName string, lastName string, age int) Person {
	return Person{firstName, lastName, age}
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	return &Person{firstName, lastName, age}
}

func main() {
	myperson := MakePerson("Gob", "Gobleson", 20)
	// myperson escapes to heap
	// This is because the parameter to fmt.Println are ...any.
	// The Go compiler moves to the heap any value that is passed in to a function
	// via a parameter that is of an interface type.
	fmt.Println(myperson)
	pointer := MakePersonPointer("Gob", "Gobleson", 20)
	fmt.Println(pointer)
}
