package main

import "fmt"

type Doubler interface {
	Double()
}

type DoubleInt int

func (d *DoubleInt) Double() {
	*d = *d * 2
}

type DoubleIntSlice []int

func (dis DoubleIntSlice) Double() {
	for i := range dis {
		dis[i] = dis[i] * 2
	}
}

func DoublerCompare(d1, d2 Doubler) bool {
	return d1 == d2
}

func main() {
	var (
		di       DoubleInt      = 14
		di2      DoubleInt      = 41
		diSlice  DoubleIntSlice = DoubleIntSlice{1, 2, 3}
		diSlice2 DoubleIntSlice = DoubleIntSlice{4, 5, 6}
	)

	// Not di and di2 are values, they do not implement the interface
	// fmt.Println(DoublerCompare(di, di2))

	// You can compare pointer instances, but their values (the address that they point to) are different
	// So it prints false
	fmt.Println(DoublerCompare(&di, &di2))

	// Same thing as before, the address don't match. Prints false
	fmt.Println(DoublerCompare(&diSlice, &diSlice2))

	// Types don't match, prints false
	fmt.Println(DoublerCompare(&di, diSlice))

	// Types match so it tries to compare the data
	// but a slice is not comparable. The program panics.
	fmt.Println(DoublerCompare(diSlice, diSlice2))
}
