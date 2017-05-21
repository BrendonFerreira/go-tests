package main

import "fmt"

func main(){

	yourAge := 16

	switch yourAge {
		case 16 : fmt.Println("Go to jail")
		case 18 : fmt.Println("Go drive")
		default: fmt.Println("Go have fun")
	}

}