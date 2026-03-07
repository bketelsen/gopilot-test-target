package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println(Greet("World"))

	http.HandleFunc("/health", HealthHandler())

	log.Println("starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func Greet(name string) string {
	return "Hello, " + name + "!"
}

func Farewell(name string) string {
	return "Goodbye, " + name + "!"
}
