package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("in doUpdateWrong:", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("in doUpdateRight:", c.String())
}

func main() {
	var c Counter
	fmt.Println(c.String())
	c.Increment() // Automatically converted to pointer
	fmt.Println(c.String())

	doUpdateWrong(c)                    // Passes copy of c to function. So the original value is not updated.
	fmt.Println("in main:", c.String()) // Prints original value
	doUpdateRight(&c)
	fmt.Println("in main:", c.String()) // Print updated value
}
