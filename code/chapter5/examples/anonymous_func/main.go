package main

import "fmt"

func main() {
	func(x int) {
		fmt.Println("Printing", x)
	}(5)
}
