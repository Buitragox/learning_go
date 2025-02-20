package main

import (
	"errors"
	"fmt"
)

func divAndRemainder(num, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		err = errors.New("Cannot divide by zero")
		return result, remainder, err
	}

	result, remainder = num/denom, num%denom
	return result, remainder, err
}

func divAndRemainderV2(num, denom int) (_ int, _ int, err error) {
	if denom == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}

	return num / denom, num % denom, err
}

func main() {
	x, y, err := divAndRemainder(5, 2)
	fmt.Println(x, y, err)

	x2, y2, err2 := divAndRemainder(5, 2)
	fmt.Println(x2, y2, err2)
}
