package main

import (
	"fmt"
)

type Score int
type HighScore Score

func printScore(s Score) {
	fmt.Println("The score is", s)
}

func main() {
	var s Score = 10
	printScore(s)

	// var hs HighScore = 50
	// printScore(hs) // Compile error
	//

	// var hs2 HighScore = s // Compile error, cannot assign Score to HighScore
	var hs3 = HighScore(s) // Use type conversion
	fmt.Println("High score:", hs3)

	s2 := s + 100 // `:=` makes `s2` a Score since `s` is a Score
	printScore(s2)
}
