package main

import "fmt"

func main() {
	const pi float64 = 3.14159265
	const myName string = "Brendon Ferreira"
	const stringSize float64 = float64(len(myName))
	fmt.Println(pi * stringSize)
}
