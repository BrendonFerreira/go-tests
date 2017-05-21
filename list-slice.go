package main

import "fmt"

func main() {
	numbersList := []float64{152232, 1298, -3123.918, 0, -1.2231}
	numbersSlice := numbersList[1:5]
	fmt.Println(numbersSlice)
}
