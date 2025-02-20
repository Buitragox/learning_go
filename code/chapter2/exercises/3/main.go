package main

import "fmt"

func main() {
	// Assign max values
	var (
		b      byte   = 255
		smallI int32  = 2_147_483_647
		bigI   uint64 = 18_446_744_073_709_551_615
	)

	b += 1
	smallI += 1
	bigI += 1

	// Overflow
	fmt.Println(b, smallI, bigI)
}
