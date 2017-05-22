package main

import (
	"fmt"
	"reflect"
)

func main() {
	stringNotTyped := "Hello World"
	fmt.Println(reflect.TypeOf(stringNotTyped))
	stringNotTyped = "Hello World!"
	fmt.Println(reflect.TypeOf(stringNotTyped))
	// stringNotTyped := 32231.32312312 -> Return an assignment error
}
