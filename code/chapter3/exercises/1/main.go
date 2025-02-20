package main

import "fmt"

func main() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	first2 := greetings[:2]
	middle3 := greetings[1:4]
	last2 := greetings[3:]

	fmt.Printf("greetings: %v\n", greetings)
	fmt.Printf("firstTwo: %v\n", first2)
	fmt.Printf("middle: %v\n", middle3)
	fmt.Printf("end: %v\n", last2)
}
