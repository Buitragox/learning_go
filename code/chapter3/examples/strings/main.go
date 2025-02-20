package main

import "fmt"

func main() {
	s := "Hello 🌎"
	s2 := s[4:7]
	s3 := s[6:]

	fmt.Printf("s: %v, len: %d\n", s, len(s))
	fmt.Printf("s2: %v, len: %d\n", s2, len(s2))
	fmt.Printf("s3: %v, len: %d\n", s3, len(s3))

	bs := []byte(s)
	rs := []rune(s)
	fmt.Printf("byte slice: %v, len: %d\n", bs, len(bs))
	fmt.Printf("rune slice: %v, len: %d\n", rs, len(rs))
}
