package main

import "fmt"

func main() {
	message := "Hi 👩 and 👨"
	runeSlice := []rune(message)
	fmt.Printf("message: '%v', len: %d\n", message, len(message))
	fmt.Println(string(runeSlice[3]))
}
