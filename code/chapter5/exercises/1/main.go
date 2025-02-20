package main

import (
	"errors"
	"fmt"
	"strconv"
)

func add(x, y int) (int, error) {
	return x + y, nil
}

func sub(x, y int) (int, error) {
	return x - y, nil
}

func mult(x, y int) (int, error) {
	return x * y, nil
}

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Cannot divide by 0")
	}
	return x / y, nil
}

type Op func(int, int) (int, error)

var opMap map[string]Op = map[string]Op{
	"+": add,
	"-": sub,
	"*": mult,
	"/": div,
}

func main() {
	inputs := [][]string{
		{"5", "+", "10"},
		{"10", "-", "25"},
		{"50", "*", "7"},
		{"60", "/", "12"},
		{"40"},
		{"hello", "+", "10"},
		{"10", "sum", "40"},
		{"50", "*", "bye"},
		{"9", "/", "0"},
	}

	for _, input := range inputs {
		if len(input) != 3 {
			fmt.Println("3 values are needed")
			continue
		}

		x, err := strconv.Atoi(input[0])
		if err != nil {
			fmt.Println("First number is invalid")
			continue
		}

		op, ok := opMap[input[1]]
		if !ok {
			fmt.Println("Invalid operator")
			continue
		}

		y, err := strconv.Atoi(input[2])
		if err != nil {
			fmt.Println("Second number is invalid")
			continue
		}

		result, err := op(x, y)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("Result:", result)
	}
}
