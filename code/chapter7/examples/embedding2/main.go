package main

import "fmt"

type Inner struct {
	X int
}

func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner IntPrinter: %d", val)
}

func (i Inner) DoublePrinter() string {
	return i.IntPrinter(i.X * 2)
}

type Outer struct {
	Inner
	X int
}

func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer IntPrinter: %d", val)
}

func main() {
	o := Outer{
		Inner: Inner{
			X: 325,
		},
		X: 2,
	}

	fmt.Println(o.X)

	// Access fields and methods of the same name by accessing the embedded field first.
	fmt.Println(o.Inner.X)

	// DoublePrinter will use Inner IntPrinter.
	fmt.Println(o.DoublePrinter())
}
