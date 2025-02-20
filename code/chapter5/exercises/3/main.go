package main

import "fmt"

func prefixer(prefix string) func(string) string {
	return func(s string) string {
		return fmt.Sprintf("%s %s", prefix, s)
		// return prefix + " " + s
	}
}

func main() {
	helloPrefix := prefixer("hello")
	fmt.Println(helloPrefix("world"))
}
