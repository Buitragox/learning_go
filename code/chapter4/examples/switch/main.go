package main

import "fmt"

func main() {
	words := []string{"bird", "potato", "monitor", "motherboard"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is short")
		case 5, 6:
			fmt.Println(word, "is nice")
		case 7, 8: //do nothing with this cases
		default:
			fmt.Println(word, "is long")
		}
	}

	for _, word := range words {
		switch size := len(word); {
		case size <= 4:
			fmt.Println(word, "is short")
		case size == 5 || size == 6:
			fmt.Println(word, "is nice")
		case size == 7 || size == 8: //do nothing with this cases
		default:
			fmt.Println(word, "is long")
		}
	}
}
