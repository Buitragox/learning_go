package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

func (e Employee) String() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

func main() {
	m := Manager{
		Employee: Employee{
			Name: "Gorb",
			ID:   "45678",
		},
		Reports: []Employee{},
	}

	fmt.Println("ID", m.ID) // Able to access the fields of Employee directly
	fmt.Println(m.String()) // Able to use the methods of Employee
}
