package main

import (
	"fmt"
)

func main() {
	user := map[string]string{
		"name":     "bertioulli",
		"username": "beercules",
	}

	for index, value := range user {
		fmt.Println(index+":", value)
	}

}
