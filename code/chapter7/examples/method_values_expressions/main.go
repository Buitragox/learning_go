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
	return fmt.Sprintf("%d, %v", c.total, c.lastUpdated)
}

func main() {
	var c Counter
	f1 := c.Increment // method value
	f1()
	fmt.Println(c)

	f2 := (*Counter).Increment // method expression with pointer
	f2(&c)

	f3 := Counter.String // method expression with value
	fmt.Println(f3(c))
}
