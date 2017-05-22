package main

import (
	"fmt"
)

type Person struct {
	name    string
	age     int
	friends []string
}

func (p Person) printFriends() {
	fmt.Println(p.friends)
}

func (p Person) getName() {
	return p.name
}

func main() {
	friends := []string{"Yoa", "Aoy", "Oay"}
	person1 := Person{"Bertioulli", 19, friends}
	person1.printFriends()
}
