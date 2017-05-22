package main

import (
	"fmt"
)

type Person struct {
	name    string
	age     int
	friends []string
}

func main() {
	friends := []string{"Yoa", "Aoy", "Oay"}
	person1 := Person{"Bertioulli", 19, friends}
	fmt.Println(person1)
}
