package main

import (
	"fmt"
)

func main() {
	factorialOfFive := factorial(5)
	fmt.Println(factorialOfFive)
}

func factorial(value int) int {
	if value <= 1 {
		return 1
	}
	return value * factorial((value - 1))
}
