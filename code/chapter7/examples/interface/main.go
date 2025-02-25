package main

import (
	"fmt"
	"time"
)

type Incrementer interface {
	Increment()
}

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
	var myStringer fmt.Stringer
	var myIncrementer Incrementer

	pointer := &Counter{}
	value := Counter{}

	myStringer = pointer
	myStringer = value
	fmt.Println(myStringer.String())

	myIncrementer = pointer

	// Value instance does not have pointer receiver methods in the method set.
	// myIncrementer = value // Compile-time error
	myIncrementer.Increment()
}
