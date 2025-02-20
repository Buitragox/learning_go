package main

import "fmt"

func main() {
	var total int
	for i := 0; i < 10; i++ {
		total := total + 1 // shadowing a variable by accident causes weird behavior
		fmt.Println(total)
	}
}
