package main

import "fmt"

func main() {
	ages := make(map[string]int, 5)
	fmt.Printf("ages: %v, len: %d\n", ages, len(ages))

	ages["Marco"] = 30
	ages["Hanna"] = 29
	ages["Juan"] = 25
	fmt.Printf("ages: %v, len: %d\n", ages, len(ages))

	fmt.Println("Juan:", ages["Juan"])  // 25
	fmt.Println("Maria", ages["Maria"]) // 0

	ages["Maria"]++
	fmt.Println("Maria:", ages["Maria"]) // 1

	v, ok := ages["Ryan"] // 0, false
	fmt.Println(v, ok)

	delete(ages, "Hanna")
	fmt.Printf("After delete\n\tages: %v, len: %d\n", ages, len(ages))

	clear(ages)
	fmt.Printf("After clear\n\tages: %v, len: %d\n", ages, len(ages))

}
