package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	value1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}

	operator := os.Args[2]

	value2, err := strconv.Atoi(os.Args[3])
	if err != nil {
		return
	}

	var result int

	switch operator {
	case "+":
		result = value1 + value2
	case "-":
		result = value1 - value2
	case "*":
		result = value1 * value2
	case "/":
		if value2 == 0 {
			fmt.Println("No division by 0")
			return
		}
		result = value1 / value2
	case "%":
		if value2 == 0 {
			fmt.Println("No modulo by 0")
			return
		}
		result = value1 % value2
	default:
		return
	}

	fmt.Println(result)
}
