package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println(Greet("World"))

	http.HandleFunc("/health", healthHandler())

	log.Println("starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Greet(name string) string {
	return "Hello, " + name + "!"
}

func Farewell(name string) string {
	return "Goodbye, " + name + "!"
}
