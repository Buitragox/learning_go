package main

import "fmt"

type NumberPrinter interface {
	NumberPrint(int)
}

type PrintFunc func(int)

func (f PrintFunc) NumberPrint(i int) {
	f(i)
}

func SquarePrint(i int) {
	fmt.Println(i * i)
}

func PrinterCall(fp NumberPrinter, i int) {
	fp.NumberPrint(i)
}

func main() {
	var doublePrint PrintFunc = func(i int) {
		fmt.Println(i * 2)
	}

	doublePrint.NumberPrint(5)
	PrintFunc(SquarePrint).NumberPrint(5)

	PrinterCall(doublePrint, 10)
	PrinterCall(PrintFunc(SquarePrint), 10)
}
