package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	miwhale := Employee{"Miwhale", "Johnson", 12345}

	james := Employee{
		firstName: "Nibor",
		lastName:  "Fourthero",
		id:        56789,
	}

	var gorble Employee
	gorble.firstName = "Gorble"
	gorble.lastName = "Gorbolson"
	gorble.id = 67890

	fmt.Printf("miwhale: %v\n", miwhale)
	fmt.Printf("james: %v\n", james)
	fmt.Printf("gorble: %v\n", gorble)
}
