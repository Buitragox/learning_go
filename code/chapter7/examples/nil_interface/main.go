package main

import "fmt"

type Whateverer interface {
	Whatever()
}

type Something struct {
	X int
}

func (s Something) Whatever() {
	fmt.Println(s.X)
}

func main() {
	var w Whateverer
	fmt.Println("w nil?", w == nil)

	var s *Something
	fmt.Println("s nil?", s == nil)

	w = s
	fmt.Println("w nil?", w == nil)
}
