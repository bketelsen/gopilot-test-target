package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println(Greet("World"))

	http.HandleFunc("/health", healthHandler())
}

func Greet(name string) string {
	return "Hello, " + name + "!"
}

func Farewell(name string) string {
	return "Goodbye, " + name + "!"
}
