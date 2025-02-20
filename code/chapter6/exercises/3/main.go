// Compile this program and time the execution with the GOGC flag
// set to different values [50 100 200 500 1000 off]
// Then change the program to specify the capacity and compare again.
// You will see that allocating the memory first is much faster

package main

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

const size int = 10_000_000

func main() {
	persons := []Person{}
	//persons := make([]Person, 0, size)
	for i := 0; i < size; i++ {
		persons = append(persons, Person{"Name", "LastName", 50})
	}
}
