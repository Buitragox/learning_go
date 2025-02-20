package main

import "fmt"

func UpdateSlice(strings []string, s string) {
	size := len(strings)
	strings[size-1] = s
	fmt.Println("UpdateSlice:", strings)
}

func GrowSlice(strings []string, s string) {
	strings = append(strings, s)
	fmt.Println("GrowSlice:", strings)
}

func main() {
	strings := []string{"Hello", "Gob", "Goodbye"}
	fmt.Println(strings)
	UpdateSlice(strings, "Potato")
	fmt.Println(strings)
	GrowSlice(strings, "Tomato")
	fmt.Println(strings)
}
