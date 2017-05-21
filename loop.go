package main

import "fmt"

func main() {
	//var pi float64 = 3.14159265
	//fmt.Printf("pi is: %.3f\n", pi)
	loop()
}

func loop() {
	i := 0

	for i < 10 {
		fmt.Println(i)
		i++
	}
}
