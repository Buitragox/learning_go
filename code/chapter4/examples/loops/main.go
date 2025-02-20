package main

import "fmt"

func main() {
	samples := []string{"hello", "apple_π!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}

	words := []string{"potato", "rendezvous"}
outer:
	for _, word := range words {
		for i, r := range word {
			if r == 'z' {
				fmt.Println("Found z at index", i)
				break outer
			}
		}
		fmt.Println("No z in", word)
	}
}
