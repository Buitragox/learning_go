package main

import "fmt"

type book struct {
	title  string
	pages  int
	author string
}

func main() {
	learningGo := book{
		title:  "Learning Go",
		author: "Jon Bodner",
	}

	fmt.Println(learningGo)

	phone := struct {
		brand string
		color string
	}{
		brand: "samsung",
		color: "blue",
	}

	fmt.Println(phone)
}
