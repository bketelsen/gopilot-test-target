package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	startTime = time.Now()
	fmt.Println(Greet("World"))

	http.HandleFunc("/health", healthHandler())
	http.ListenAndServe(":8080", nil)
}

func Greet(name string) string {
	return "Hello, " + name + "!"
}

func Farewell(name string) string {
	return "Goodbye, " + name + "!"
}
