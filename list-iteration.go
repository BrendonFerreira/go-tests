package main

import "fmt"

func main() {
	var numbersList [5]float64

	numbersList[0] = 152232
	numbersList[1] = 1298
	numbersList[2] = -3123.918
	numbersList[3] = 0
	numbersList[4] = -1.2231

	for index, value := range numbersList {
		fmt.Println(index, value)
	}
}
