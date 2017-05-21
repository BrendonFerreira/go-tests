package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/beego/mux"
)

func main() {
	router := mux.New()
	router.Get("/", mainPage)
	router.Get("/user/:id", consultUser)
	fmt.Println("Server listening in port", 3000)
	log.Fatal(http.ListenAndServe(":3000", router))
}

func mainPage(res http.ResponseWriter, req *http.Request) {
	message := []byte("Hello World")
	res.Write(message)
}

func consultUser(res http.ResponseWriter, req *http.Request) {
	id := mux.Param(req, ":id")
	// Write in res the formated print
	fmt.Fprintf(res, "hello, user page %s", id)
}
